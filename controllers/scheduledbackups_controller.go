// dbaas-operator
// Copyright (C) 2022 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	dbaasv1 "github.com/percona/dbaas-operator/api/v1"
	psmdbv1 "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
	pxcv1 "github.com/percona/percona-xtradb-cluster-operator/pkg/apis/pxc/v1"
	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// ScheduledBackupReconciler reconciles a DatabaseEngine object.
type ScheduledBackupReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=batchv1,resources=job,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ScheduledBackupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling", "request", req)
	job := &batchv1.Job{}
	err := r.Get(ctx, req.NamespacedName, job)
	if err != nil {
		return ctrl.Result{}, err
	}
	if job.Status.Active > 0 {
		err := r.backupSecrets(ctx, job)
		if err != nil {
			logger.Error(err, "failed backing up secrets")
			return ctrl.Result{}, err
		}

	}
	logger.Info("Reconciled", "request", req)

	return ctrl.Result{}, nil

}
func (r *ScheduledBackupReconciler) backupSecrets(ctx context.Context, job *batchv1.Job) error {
	clusterName, ok := job.Spec.Selector.MatchLabels["cluster"]
	if clusterName == "" || !ok {
		return errors.New("can't obtain clusterName from labels")
	}
	jobName, ok := job.Spec.Selector.MatchLabels["job-name"]
	if jobName == "" || !ok {
		return errors.New("can't obtain jobName from labels")
	}
	jobType, ok := job.Spec.Selector.MatchLabels["type"]
	if jobType == "" || !ok {
		return errors.New("can't obtain jobType from labels")
	}

	dbCluster := &dbaasv1.DatabaseCluster{}
	err := r.Get(ctx, types.NamespacedName{
		Name:      clusterName,
		Namespace: job.Namespace,
	}, dbCluster)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed getting the cluster spec for %s", clusterName))
	}

	secret := &corev1.Secret{}
	err = r.Get(ctx, types.NamespacedName{
		Name:      dbCluster.Spec.SecretsName,
		Namespace: job.Namespace,
	}, secret)

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed getting secret for the cluster %s", clusterName))
	}
	var storageProvider *dbaasv1.BackupStorageProviderSpec
	if jobType == "xtrabackup" {
		backup := &pxcv1.PerconaXtraDBClusterBackup{}
		err = r.Get(ctx, types.NamespacedName{
			Name:      jobName,
			Namespace: job.Namespace,
		}, backup)
		if err != nil {
			return err
		}
		storageName, ok := dbCluster.Spec.Backup.Storages[backup.Spec.StorageName]
		if !ok {
			return errors.New("storage does not exist in the backup spec")
		}
		storageProvider = storageName.StorageProvider
	}

	backupSecret := &corev1.Secret{}

	err = r.Get(ctx, types.NamespacedName{
		Name:      storageProvider.CredentialsSecret,
		Namespace: job.Namespace,
	}, backupSecret)
	if err != nil {
		return errors.Wrap(err, "failed getting secret with credentials to the cloud provider")
	}

	aws_access_key, ok := backupSecret.Data["AWS_ACCESS_KEY_ID"]
	if len(aws_access_key) == 0 || !ok {
		return errors.Wrap(err, "empty AWS_ACCESS_KEY_ID")
	}
	aws_secret_key, ok := backupSecret.Data["AWS_SECRET_ACCESS_KEY"]
	if len(aws_secret_key) == 0 || !ok {
		return errors.Wrap(err, "empty AWS_SECRET_ACCESS_KEY")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(storageProvider.Region),
		Credentials: credentials.NewStaticCredentials(
			string(aws_access_key),
			string(aws_secret_key),
			""),
	})
	if err != nil {
		return errors.Wrap(err, "failed creating aws session with provided credentials")
	}
	s3Client := s3.New(sess)

	data, err := json.Marshal(secret)
	if err != nil {
		return errors.Wrap(err, "failed marshaling secret data")
	}

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(storageProvider.Bucket),
		Body:   aws.ReadSeekCloser(bytes.NewBuffer(data)),
		Key:    aws.String(fmt.Sprintf("%s-secrets", clusterName)),
	})

	return err
}
func (r *ScheduledBackupReconciler) addPXCKnownTypes(scheme *runtime.Scheme) error {
	pxcSchemeGroupVersion := schema.GroupVersion{Group: "pxc.percona.com", Version: "v1"}
	scheme.AddKnownTypes(pxcSchemeGroupVersion,
		&pxcv1.PerconaXtraDBClusterBackup{}, &pxcv1.PerconaXtraDBClusterBackupList{})

	metav1.AddToGroupVersion(scheme, pxcSchemeGroupVersion)
	return nil
}

func (r *ScheduledBackupReconciler) addPSMDBKnownTypes(scheme *runtime.Scheme) error {
	pxcSchemeGroupVersion := schema.GroupVersion{Group: "psmdb.percona.com", Version: "v1"}
	scheme.AddKnownTypes(pxcSchemeGroupVersion,
		&psmdbv1.PerconaServerMongoDBBackup{}, &psmdbv1.PerconaServerMongoDBBackupList{})

	metav1.AddToGroupVersion(scheme, pxcSchemeGroupVersion)
	return nil
}

func (r *ScheduledBackupReconciler) addPXCToScheme(scheme *runtime.Scheme) error {
	builder := runtime.NewSchemeBuilder(r.addPXCKnownTypes)
	return builder.AddToScheme(scheme)
}

func (r *ScheduledBackupReconciler) addPSMDBToScheme(scheme *runtime.Scheme) error {
	builder := runtime.NewSchemeBuilder(r.addPSMDBKnownTypes)
	return builder.AddToScheme(scheme)
}

// SetupWithManager sets up the controller with the Manager.
func (r *ScheduledBackupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	unstructuredResource := &unstructured.Unstructured{}
	unstructuredResource.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Kind:    "CustomResourceDefinition",
		Version: "v1",
	})
	err := r.Get(context.Background(), types.NamespacedName{Name: pxcRestoreCRDName}, unstructuredResource)
	if err == nil {
		if err := r.addPXCToScheme(r.Scheme); err != nil {
			return err
		}
	}

	err = r.Get(context.Background(), types.NamespacedName{Name: psmdbRestoreCRDName}, unstructuredResource)
	if err == nil {
		if err := r.addPSMDBToScheme(r.Scheme); err != nil {
			return err
		}
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.Job{}).
		WithEventFilter(scheduledBackupPredicate()).
		Complete(r)
}

func scheduledBackupPredicate() predicate.Predicate {
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			references := e.ObjectNew.GetOwnerReferences()
			for _, reference := range references {
				return reference.Kind == "PerconaXtraDBClusterBackup" || reference.Kind == "PerconaServerMongoDBBackup"
			}
			return false
		},
		CreateFunc: func(e event.CreateEvent) bool {
			return false
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return false
		},
	}
}
