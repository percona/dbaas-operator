//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSchedule) DeepCopyInto(out *BackupSchedule) {
	*out = *in
	if in.CompressionLevel != nil {
		in, out := &in.CompressionLevel, &out.CompressionLevel
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSchedule.
func (in *BackupSchedule) DeepCopy() *BackupSchedule {
	if in == nil {
		return nil
	}
	out := new(BackupSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSource) DeepCopyInto(out *BackupSource) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(BackupStorageProviderSpec)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(BackupStorageProviderSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSource.
func (in *BackupSource) DeepCopy() *BackupSource {
	if in == nil {
		return nil
	}
	out := new(BackupSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]BackupSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make(map[string]*BackupStorageSpec, len(*in))
		for key, val := range *in {
			var outVal *BackupStorageSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(BackupStorageSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageProviderSpec) DeepCopyInto(out *BackupStorageProviderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageProviderSpec.
func (in *BackupStorageProviderSpec) DeepCopy() *BackupStorageProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageSpec) DeepCopyInto(out *BackupStorageSpec) {
	*out = *in
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(VolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageProvider != nil {
		in, out := &in.StorageProvider, &out.StorageProvider
		*out = new(BackupStorageProviderSpec)
		**out = **in
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.VerifyTLS != nil {
		in, out := &in.VerifyTLS, &out.VerifyTLS
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageSpec.
func (in *BackupStorageSpec) DeepCopy() *BackupStorageSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBInstanceSpec) DeepCopyInto(out *DBInstanceSpec) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	out.DiskSize = in.DiskSize.DeepCopy()
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBInstanceSpec.
func (in *DBInstanceSpec) DeepCopy() *DBInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(DBInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseCluster) DeepCopyInto(out *DatabaseCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseCluster.
func (in *DatabaseCluster) DeepCopy() *DatabaseCluster {
	if in == nil {
		return nil
	}
	out := new(DatabaseCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterBackup) DeepCopyInto(out *DatabaseClusterBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterBackup.
func (in *DatabaseClusterBackup) DeepCopy() *DatabaseClusterBackup {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseClusterBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterBackupList) DeepCopyInto(out *DatabaseClusterBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseClusterBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterBackupList.
func (in *DatabaseClusterBackupList) DeepCopy() *DatabaseClusterBackupList {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseClusterBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterBackupSpec) DeepCopyInto(out *DatabaseClusterBackupSpec) {
	*out = *in
	in.BackupSource.DeepCopyInto(&out.BackupSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterBackupSpec.
func (in *DatabaseClusterBackupSpec) DeepCopy() *DatabaseClusterBackupSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterBackupStatus) DeepCopyInto(out *DatabaseClusterBackupStatus) {
	*out = *in
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterBackupStatus.
func (in *DatabaseClusterBackupStatus) DeepCopy() *DatabaseClusterBackupStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterList) DeepCopyInto(out *DatabaseClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterList.
func (in *DatabaseClusterList) DeepCopy() *DatabaseClusterList {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterRestore) DeepCopyInto(out *DatabaseClusterRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterRestore.
func (in *DatabaseClusterRestore) DeepCopy() *DatabaseClusterRestore {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseClusterRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterRestoreList) DeepCopyInto(out *DatabaseClusterRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseClusterRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterRestoreList.
func (in *DatabaseClusterRestoreList) DeepCopy() *DatabaseClusterRestoreList {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseClusterRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterRestoreSpec) DeepCopyInto(out *DatabaseClusterRestoreSpec) {
	*out = *in
	if in.BackupSource != nil {
		in, out := &in.BackupSource, &out.BackupSource
		*out = new(BackupSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterRestoreSpec.
func (in *DatabaseClusterRestoreSpec) DeepCopy() *DatabaseClusterRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterRestoreStatus) DeepCopyInto(out *DatabaseClusterRestoreStatus) {
	*out = *in
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
	if in.LastScheduled != nil {
		in, out := &in.LastScheduled, &out.LastScheduled
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterRestoreStatus.
func (in *DatabaseClusterRestoreStatus) DeepCopy() *DatabaseClusterRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseClusterStatus) DeepCopyInto(out *DatabaseClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseClusterStatus.
func (in *DatabaseClusterStatus) DeepCopy() *DatabaseClusterStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseEngine) DeepCopyInto(out *DatabaseEngine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseEngine.
func (in *DatabaseEngine) DeepCopy() *DatabaseEngine {
	if in == nil {
		return nil
	}
	out := new(DatabaseEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseEngine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseEngineList) DeepCopyInto(out *DatabaseEngineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseEngine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseEngineList.
func (in *DatabaseEngineList) DeepCopy() *DatabaseEngineList {
	if in == nil {
		return nil
	}
	out := new(DatabaseEngineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseEngineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseEngineSpec) DeepCopyInto(out *DatabaseEngineSpec) {
	*out = *in
	if in.EngineVersions != nil {
		in, out := &in.EngineVersions, &out.EngineVersions
		*out = make(map[string]*Component, len(*in))
		for key, val := range *in {
			var outVal *Component
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Component)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseEngineSpec.
func (in *DatabaseEngineSpec) DeepCopy() *DatabaseEngineSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseEngineStatus) DeepCopyInto(out *DatabaseEngineStatus) {
	*out = *in
	in.AvailableVersions.DeepCopyInto(&out.AvailableVersions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseEngineStatus.
func (in *DatabaseEngineStatus) DeepCopy() *DatabaseEngineStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseEngineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	in.DBInstance.DeepCopyInto(&out.DBInstance)
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(BackupSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DataSource != nil {
		in, out := &in.DataSource, &out.DataSource
		*out = new(BackupSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSpec) DeepCopyInto(out *MonitoringSpec) {
	*out = *in
	if in.PMM != nil {
		in, out := &in.PMM, &out.PMM
		*out = new(PMMSpec)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSpec.
func (in *MonitoringSpec) DeepCopy() *MonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PMMSpec) DeepCopyInto(out *PMMSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PMMSpec.
func (in *PMMSpec) DeepCopy() *PMMSpec {
	if in == nil {
		return nil
	}
	out := new(PMMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Versions) DeepCopyInto(out *Versions) {
	*out = *in
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = make(map[string]*Component, len(*in))
		for key, val := range *in {
			var outVal *Component
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Component)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = make(map[string]*Component, len(*in))
		for key, val := range *in {
			var outVal *Component
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Component)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = make(map[string]map[string]*Component, len(*in))
		for key, val := range *in {
			var outVal map[string]*Component
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]*Component, len(*in))
				for key, val := range *in {
					var outVal *Component
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(Component)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Tools != nil {
		in, out := &in.Tools, &out.Tools
		*out = make(map[string]map[string]*Component, len(*in))
		for key, val := range *in {
			var outVal map[string]*Component
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]*Component, len(*in))
				for key, val := range *in {
					var outVal *Component
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(Component)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Versions.
func (in *Versions) DeepCopy() *Versions {
	if in == nil {
		return nil
	}
	out := new(Versions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(corev1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(corev1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
