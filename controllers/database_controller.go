// dbaas-operator
// Copyright (C) 2022 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package controllers

import (
	"context"
	"time"

	psmdbAPIs "github.com/percona/percona-server-mongodb-operator/pkg/apis"
	psmdbv1 "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
	pxcAPIs "github.com/percona/percona-xtradb-cluster-operator/pkg/apis"
	pxcv1 "github.com/percona/percona-xtradb-cluster-operator/pkg/apis/pxc/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	dbaasv1 "github.com/gen1us2k/dbaas-operator/api/v1"
)

const (
	pxcDefaultImage          = "percona/percona-xtradb-cluster:8.0.27-18.1"
	PerconaXtraDBClusterKind = "PerconaXtraDBCluster"
)

// DatabaseReconciler reconciles a Database object
type DatabaseReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=dbaas.percona.com,resources=databases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=dbaas.percona.com,resources=databases/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=dbaas.percona.com,resources=databases/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *DatabaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling", "request", req)

	database := &dbaasv1.Database{}

	err := r.Get(ctx, req.NamespacedName, database)
	if err != nil {
		// NotFound cannot be fixed by requeuing so ignore it. During background
		// deletion, we receive delete events from cluster's dependents after
		// cluster is deleted.
		if err = client.IgnoreNotFound(err); err != nil {
			logger.Error(err, "unable to fetch DatabaseCluster")
		}
		return reconcile.Result{}, err
	}
	logger.Info("Reconciled", "request", req)
	if database.Spec.Database == "pxc" {
		err := r.reconcilePXC(ctx, req, database)
		return reconcile.Result{}, err
	}
	if database.Spec.Database == "psmdb" {
		err := r.reconcilePSMDB(ctx, req, database)
		return reconcile.Result{}, err
	}
	return reconcile.Result{RequeueAfter: 5 * time.Second}, nil
}
func (r *DatabaseReconciler) reconcilePSMDB(ctx context.Context, req ctrl.Request, database *dbaasv1.Database) error {
	diskSize, err := resource.ParseQuantity(database.Spec.DBInstance.DiskSize)
	if err != nil {
		return err
	}
	psmdb := &psmdbv1.PerconaServerMongoDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:       database.Name,
			Namespace:  database.Namespace,
			Finalizers: []string{"delete-psmdb-pvc"},
		},
	}
	if err := controllerutil.SetControllerReference(database, psmdb, r.Client.Scheme()); err != nil {
		return err
	}
	_, err = controllerutil.CreateOrUpdate(ctx, r.Client, psmdb, func() error {
		psmdb.TypeMeta = metav1.TypeMeta{
			APIVersion: "psmdb.percona.com/v1-12-0",
			Kind:       "PerconaServerMongoDB",
		}
		psmdb.Spec = psmdbv1.PerconaServerMongoDBSpec{
			CRVersion:  "1.12.0",
			UnsafeConf: true,
			Image:      "percona/percona-server-mongodb:5.0.11-10",
			Secrets: &psmdbv1.SecretsSpec{
				Users: "minimal-cluster-secrets",
			},
			UpgradeOptions: psmdbv1.UpgradeOptions{
				Apply:    "disabled",
				Schedule: "0 4 * * *",
			},
			Replsets: []*psmdbv1.ReplsetSpec{
				{
					Name:          "rs0",
					Configuration: psmdbv1.MongoConfiguration(database.Spec.DatabaseConfig),
					Size:          1,
					VolumeSpec: &psmdbv1.VolumeSpec{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimSpec{
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceStorage: diskSize,
								},
							},
						},
					},
				},
			},
			Sharding: psmdbv1.Sharding{
				Enabled: true,
				ConfigsvrReplSet: &psmdbv1.ReplsetSpec{
					Size: 1,
					VolumeSpec: &psmdbv1.VolumeSpec{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimSpec{
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceStorage: diskSize,
								},
							},
						},
					},
				},
				Mongos: &psmdbv1.MongosSpec{
					Size: 1,
				},
			},
		}

		return nil
	})
	if err != nil {
		return err
	}
	psmdb = &psmdbv1.PerconaServerMongoDB{}
	if err := r.Get(ctx, req.NamespacedName, psmdb); err != nil {
		return err
	}
	database.Status.Host = psmdb.Status.Host
	database.Status.State = dbaasv1.AppState(psmdb.Status.State)
	if err := r.Status().Update(ctx, database); err != nil {
		return err
	}
	return nil
}
func (r *DatabaseReconciler) reconcilePXC(ctx context.Context, req ctrl.Request, database *dbaasv1.Database) error {
	diskSize, err := resource.ParseQuantity(database.Spec.DBInstance.DiskSize)
	if err != nil {
		return err
	}
	pxc := &pxcv1.PerconaXtraDBCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:       database.Name,
			Namespace:  database.Namespace,
			Finalizers: []string{"delete-proxysql-pvc", "delete-pxc-pvc"},
		},
	}
	if err := controllerutil.SetControllerReference(database, pxc, r.Client.Scheme()); err != nil {
		return err
	}
	_, err = controllerutil.CreateOrUpdate(ctx, r.Client, pxc, func() error {
		pxc.TypeMeta = metav1.TypeMeta{
			APIVersion: "pxc.percona.com/v1-11-0",
			Kind:       PerconaXtraDBClusterKind,
		}
		pxc.Spec = pxcv1.PerconaXtraDBClusterSpec{
			CRVersion:         "1.11.0",
			AllowUnsafeConfig: true,
			SecretsName:       "minimal-cluster-secrets-trololo",
			UpgradeOptions: pxcv1.UpgradeOptions{
				Apply:    "8.0-recommended",
				Schedule: "0 4 * * *",
			},
			PXC: &pxcv1.PXCSpec{
				PodSpec: &pxcv1.PodSpec{
					Configuration: database.Spec.DatabaseConfig,
					ServiceType:   corev1.ServiceTypeClusterIP,
					Size:          database.Spec.ClusterSize,
					Image:         pxcDefaultImage,
					VolumeSpec: &pxcv1.VolumeSpec{
						PersistentVolumeClaim: &corev1.PersistentVolumeClaimSpec{
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceStorage: diskSize,
								},
							},
						},
					},
				},
			},
			HAProxy: &pxcv1.HAProxySpec{
				PodSpec: pxcv1.PodSpec{
					Size:        database.Spec.ClusterSize,
					ServiceType: corev1.ServiceTypeClusterIP,
					Enabled:     true,
					Image:       "percona/percona-xtradb-cluster-operator:1.11.0-haproxy",
				},
			},
		}
		return nil
	})
	if err != nil {
		return err
	}
	cPXC := &pxcv1.PerconaXtraDBCluster{}
	if err := r.Get(ctx, req.NamespacedName, cPXC); err != nil {
		return err
	}
	database.Status.Host = cPXC.Status.Host
	database.Status.State = dbaasv1.AppState(cPXC.Status.Status)
	if err := r.Status().Update(ctx, database); err != nil {
		return err
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DatabaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := pxcAPIs.AddToScheme(r.Scheme); err != nil {
		return err
	}
	if err := psmdbAPIs.AddToScheme(r.Scheme); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&dbaasv1.Database{}).
		Complete(r)
}
