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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
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
func (in *BackupStorageAzureSpec) DeepCopyInto(out *BackupStorageAzureSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageAzureSpec.
func (in *BackupStorageAzureSpec) DeepCopy() *BackupStorageAzureSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageAzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageGCSSpec) DeepCopyInto(out *BackupStorageGCSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageGCSSpec.
func (in *BackupStorageGCSSpec) DeepCopy() *BackupStorageGCSSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageGCSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageS3Spec) DeepCopyInto(out *BackupStorageS3Spec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageS3Spec.
func (in *BackupStorageS3Spec) DeepCopy() *BackupStorageS3Spec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageS3Spec)
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
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(BackupStorageS3Spec)
		**out = **in
	}
	if in.GCS != nil {
		in, out := &in.GCS, &out.GCS
		*out = new(BackupStorageGCSSpec)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(BackupStorageAzureSpec)
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
func (in *DBInstanceSpec) DeepCopyInto(out *DBInstanceSpec) {
	*out = *in
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
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	in.LoadBalancer.DeepCopyInto(&out.LoadBalancer)
	in.Monitoring.DeepCopyInto(&out.Monitoring)
	out.DBInstance = in.DBInstance
	in.Backup.DeepCopyInto(&out.Backup)
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
	out.PMM = in.PMM
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
