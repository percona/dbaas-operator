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
	"context"
	"fmt"
	"time"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	dbaasv1 "github.com/percona/dbaas-operator/api/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"

	psmdbv1 "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
	pxcv1 "github.com/percona/percona-xtradb-cluster-operator/pkg/apis/pxc/v1"
	"github.com/pkg/errors"
)

const (
	pxcRestoreKind   = "PerconaXtraDBClusterRestore"
	pxcRestoreAPI    = "pxc.percona.com/v1"
	psmdbRestoreKind = "PerconaServerMongoDBRestore"
	psmdbRestoreAPI  = "psmdb.percona.com/v1"
)

// DatabaseClusterRestoreReconciler reconciles a DatabaseClusterRestore object
type DatabaseClusterRestoreReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=dbaas.percona.com,resources=databaseclusterrestores,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbaas.percona.com,resources=databaseclusterrestores/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=dbaas.percona.com,resources=databaseclusterrestores/finalizers,verbs=update
func (r *DatabaseClusterRestoreReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("reconciling", "request", req)

	cr := &dbaasv1.DatabaseClusterRestore{}
	err := r.Get(ctx, req.NamespacedName, cr)
	if err != nil {
		// NotFound cannot be fixed by requeuing so ignore it. During background
		// deletion, we receive delete events from cluster's dependents after
		// cluster is deleted.
		if err = client.IgnoreNotFound(err); err != nil {
			logger.Error(err, "unable to fetch DatabaseClusterRestore")
		}
		return reconcile.Result{}, err
	}
	if cr.Status.State != "" {
		return reconcile.Result{}, nil
	}

	cluster := &dbaasv1.DatabaseCluster{}
	err = r.Get(ctx, types.NamespacedName{Name: cr.Spec.DatabaseCluster, Namespace: cr.Namespace}, cluster)
	if err != nil {
		logger.Error(err, "unable to get DatabaseCluster")
		return reconcile.Result{}, err
	}
	orig := cluster.DeepCopy()

	err = r.stopCluster(cluster.DeepCopy())
	if err != nil {
		logger.Error(err, "unable to stop DatabaseCluster")
		return reconcile.Result{}, err
	}
	if cr.Spec.DatabaseType == dbaasv1.PXCEngine {
		pxc := &pxcv1.PerconaXtraDBCluster{}
		err = r.Get(ctx, types.NamespacedName{Name: cr.Spec.DatabaseCluster, Namespace: cr.Namespace}, pxc)
		if err != nil {
			return reconcile.Result{}, err
		}
		if cr.Spec.BackupSource != nil && cr.Spec.BackupSource.Image == "" {
			cr.Spec.BackupSource.Image = fmt.Sprintf(pxcBackupImageTmpl, pxc.Spec.CRVersion)
		}
		if err := r.restorePXC(cr); err != nil {
			logger.Error(err, "unable to restore PXC Cluster")
			return reconcile.Result{}, err
		}
	}
	if cr.Spec.DatabaseType == dbaasv1.PSMDBEngine {
		psmdb := &psmdbv1.PerconaServerMongoDB{}
		err = r.Get(ctx, types.NamespacedName{Name: cr.Spec.DatabaseCluster, Namespace: cr.Namespace}, psmdb)
		if err != nil {
			return reconcile.Result{}, err
		}
		if cr.Spec.BackupSource != nil && cr.Spec.BackupSource.Image == "" {
			cr.Spec.BackupSource.Image = fmt.Sprintf(psmdbBackupImageTmpl, psmdb.Spec.CRVersion)
		}
		if err := r.restorePSMDB(cr); err != nil {
			logger.Error(err, "unable to restore PXC Cluster")
			return reconcile.Result{}, err
		}
	}

	err = r.startCluster(orig)
	if err != nil {
		logger.Error(err, "failed to start DatabaseCluster")
		return reconcile.Result{}, err
	}
	pxcCR := &pxcv1.PerconaXtraDBClusterRestore{}
	err = r.Get(context.Background(), types.NamespacedName{Name: pxcCR.Name, Namespace: pxcCR.Namespace}, pxcCR)
	if err != nil {
		return reconcile.Result{}, err
	}
	cr.Status.State = dbaasv1.RestoreState(pxcCR.Status.State)
	cr.Status.CompletedAt = pxcCR.Status.CompletedAt
	cr.Status.LastScheduled = pxcCR.Status.LastScheduled
	cr.Status.Message = pxcCR.Status.Comments
	r.Status().Update(context.Background(), cr)

	return ctrl.Result{}, nil
}
func (r *DatabaseClusterRestoreReconciler) restorePSMDB(restore *dbaasv1.DatabaseClusterRestore) error {
	psmdbCR := &psmdbv1.PerconaServerMongoDBRestore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      restore.Name,
			Namespace: restore.Namespace,
		},
	}
	if err := controllerutil.SetControllerReference(restore, psmdbCR, r.Client.Scheme()); err != nil {
		return err
	}
	_, err := controllerutil.CreateOrUpdate(context.Background(), r.Client, psmdbCR, func() error {
		psmdbCR.TypeMeta = metav1.TypeMeta{
			APIVersion: psmdbRestoreAPI,
			Kind:       psmdbRestoreKind,
		}
		psmdbCR.Spec.ClusterName = restore.Spec.DatabaseCluster
		if restore.Spec.BackupName == "" && restore.Spec.BackupSource == nil {
			return errors.New("specify either backupName or backupSource")
		}
		if restore.Spec.BackupName != "" {
			psmdbCR.Spec.BackupName = restore.Spec.BackupName
		}
		if restore.Spec.BackupSource != nil {
			psmdbCR.Spec.BackupSource = &psmdbv1.PerconaServerMongoDBBackupStatus{
				Destination: restore.Spec.BackupSource.Destination,
				StorageName: restore.Spec.BackupSource.StorageName,
			}
			switch restore.Spec.BackupSource.StorageType {
			case dbaasv1.BackupStorageS3:
				psmdbCR.Spec.BackupSource.S3 = &psmdbv1.BackupStorageS3Spec{
					Bucket:            restore.Spec.BackupSource.S3.Bucket,
					CredentialsSecret: restore.Spec.BackupSource.S3.CredentialsSecret,
					Region:            restore.Spec.BackupSource.S3.Region,
					EndpointURL:       restore.Spec.BackupSource.S3.EndpointURL,
				}
			case dbaasv1.BackupStorageAzure:
				psmdbCR.Spec.BackupSource.Azure = &psmdbv1.BackupStorageAzureSpec{
					CredentialsSecret: restore.Spec.BackupSource.Azure.CredentialsSecret,
					Container:         restore.Spec.BackupSource.Azure.ContainerName,
					Prefix:            restore.Spec.BackupSource.Azure.Prefix,
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	// Waiting for the job to complete
	for {
		time.Sleep(time.Second)

		cr := &psmdbv1.PerconaServerMongoDBRestore{}
		err := r.Get(context.Background(), types.NamespacedName{Name: psmdbCR.Name, Namespace: psmdbCR.Namespace}, cr)
		if err != nil {
			return err
		}
		restore.Status.State = dbaasv1.RestoreState(cr.Status.State)
		restore.Status.Message = cr.Status.Error
		restore.Status.CompletedAt = cr.Status.CompletedAt
		r.Status().Update(context.Background(), restore)
		if cr.Status.State == psmdbv1.RestoreStateReady || cr.Status.State == psmdbv1.RestoreStateError {
			return nil
		}

	}

}
func (r *DatabaseClusterRestoreReconciler) restorePXC(restore *dbaasv1.DatabaseClusterRestore) error {
	pxcCR := &pxcv1.PerconaXtraDBClusterRestore{
		ObjectMeta: metav1.ObjectMeta{
			Name:      restore.Name,
			Namespace: restore.Namespace,
		},
	}
	if err := controllerutil.SetControllerReference(restore, pxcCR, r.Client.Scheme()); err != nil {
		return err
	}
	_, err := controllerutil.CreateOrUpdate(context.Background(), r.Client, pxcCR, func() error {
		pxcCR.TypeMeta = metav1.TypeMeta{
			APIVersion: pxcRestoreAPI,
			Kind:       pxcRestoreKind,
		}
		pxcCR.Spec.PXCCluster = restore.Spec.DatabaseCluster
		if restore.Spec.BackupName == "" && restore.Spec.BackupSource == nil {
			return errors.New("specify either backupName or backupSource")
		}
		if restore.Spec.BackupName != "" {
			pxcCR.Spec.BackupName = restore.Spec.BackupName
		}
		if restore.Spec.BackupSource != nil {
			pxcCR.Spec.BackupSource = &pxcv1.PXCBackupStatus{
				Destination: restore.Spec.BackupSource.Destination,
				StorageName: restore.Spec.BackupSource.StorageName,
			}
			switch restore.Spec.BackupSource.StorageType {
			case dbaasv1.BackupStorageS3:
				pxcCR.Spec.BackupSource.S3 = &pxcv1.BackupStorageS3Spec{
					Bucket:            restore.Spec.BackupSource.S3.Bucket,
					CredentialsSecret: restore.Spec.BackupSource.S3.CredentialsSecret,
					Region:            restore.Spec.BackupSource.S3.Region,
					EndpointURL:       restore.Spec.BackupSource.S3.EndpointURL,
				}
			case dbaasv1.BackupStorageAzure:
				pxcCR.Spec.BackupSource.Azure = &pxcv1.BackupStorageAzureSpec{
					CredentialsSecret: restore.Spec.BackupSource.Azure.CredentialsSecret,
					ContainerPath:     restore.Spec.BackupSource.Azure.ContainerName,
					Endpoint:          restore.Spec.BackupSource.Azure.EndpointURL,
					StorageClass:      restore.Spec.BackupSource.Azure.StorageClass,
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	// Waiting for the job to complete
	for {
		time.Sleep(time.Second)

		cr := &pxcv1.PerconaXtraDBClusterRestore{}
		err := r.Get(context.Background(), types.NamespacedName{Name: pxcCR.Name, Namespace: pxcCR.Namespace}, cr)
		if err != nil {
			return err
		}
		restore.Status.State = dbaasv1.RestoreState(cr.Status.State)
		restore.Status.CompletedAt = cr.Status.CompletedAt
		restore.Status.LastScheduled = cr.Status.LastScheduled
		restore.Status.Message = cr.Status.Comments
		r.Status().Update(context.Background(), restore)
		jobName := "restore-job-" + cr.Name + "-" + cr.Spec.PXCCluster
		checkJob := batchv1.Job{}
		err = r.Get(context.Background(), types.NamespacedName{Name: jobName, Namespace: pxcCR.Namespace}, &checkJob)
		if err != nil && !k8serrors.IsNotFound(err) {
			return errors.Wrap(err, "get job status")
		}
		for _, cond := range checkJob.Status.Conditions {
			if cond.Status != corev1.ConditionTrue {
				continue
			}
			switch cond.Type {
			case batchv1.JobComplete:
				return nil
			case batchv1.JobFailed:
				return errors.New(cond.Message)
			}
		}

	}

}

func (r *DatabaseClusterRestoreReconciler) startCluster(cluster *dbaasv1.DatabaseCluster) error {
	fmt.Println("starting cluster")
	current := &dbaasv1.DatabaseCluster{}
	if err := r.Get(context.Background(), types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, current); err != nil {
		return err
	}
	current.Spec = cluster.Spec
	current.Spec.Pause = false

	if err := r.Update(context.Background(), current); err != nil {
		return err
	}
	timeout, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()
	for {
		time.Sleep(time.Second)
		current := &dbaasv1.DatabaseCluster{}
		if err := r.Get(context.Background(), types.NamespacedName{Name: cluster.Name, Namespace: cluster.Namespace}, current); err != nil {
			return err
		}
		if current.Status.State == dbaasv1.AppStateReady {
			break
		}
		select {
		case <-timeout.Done():
			return errors.New("timeout waiting for a condition")
		default:
		}

	}
	return nil
}

func (r *DatabaseClusterRestoreReconciler) stopCluster(cluster *dbaasv1.DatabaseCluster) error {
	cluster.Spec.Pause = true
	return r.Update(context.Background(), cluster)
}

func (r *DatabaseClusterRestoreReconciler) addPXCKnownTypes(scheme *runtime.Scheme) error {
	pxcSchemeGroupVersion := schema.GroupVersion{Group: "pxc.percona.com", Version: "v1"}
	scheme.AddKnownTypes(pxcSchemeGroupVersion,
		&pxcv1.PerconaXtraDBClusterRestore{}, &pxcv1.PerconaXtraDBClusterRestoreList{},
	)

	metav1.AddToGroupVersion(scheme, pxcSchemeGroupVersion)
	return nil
}
func (r *DatabaseClusterRestoreReconciler) addPSMDBKnownTypes(scheme *runtime.Scheme) error {
	pxcSchemeGroupVersion := schema.GroupVersion{Group: "psmdb.percona.com", Version: "v1"}
	scheme.AddKnownTypes(pxcSchemeGroupVersion,
		&psmdbv1.PerconaServerMongoDBRestore{}, &psmdbv1.PerconaServerMongoDBRestoreList{},
	)

	metav1.AddToGroupVersion(scheme, pxcSchemeGroupVersion)
	return nil
}

func (r *DatabaseClusterRestoreReconciler) addPXCToScheme(scheme *runtime.Scheme) error {
	builder := runtime.NewSchemeBuilder(r.addPXCKnownTypes)
	return builder.AddToScheme(scheme)
}
func (r *DatabaseClusterRestoreReconciler) addPSMDBToScheme(scheme *runtime.Scheme) error {
	builder := runtime.NewSchemeBuilder(r.addPSMDBKnownTypes)
	return builder.AddToScheme(scheme)
}

// SetupWithManager sets up the controller with the Manager.
func (r *DatabaseClusterRestoreReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := r.addPXCToScheme(r.Scheme); err != nil {
		return err
	}
	if err := r.addPSMDBToScheme(r.Scheme); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&dbaasv1.DatabaseClusterRestore{}).
		Complete(r)
}
