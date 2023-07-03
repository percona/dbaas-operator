/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Storage struct {
	// Size is the size of the persistent volume claim
	Size string `json:"size"`
	// Class is the storage class to use for the persistent volume claim
	Class string `json:"class,omitempty"`
}

type Resources struct {
	// CPU is the CPU resource requirements
	CPU    string `json:"cpu,omitempty"`
	// Memory is the memory resource requirements
	Memory string `json:"memory,omitempty"`
}

type Engine struct {
	// Type is the engine type
	Type     string `json:"type"`
	// Version is the engine version
	Version  string `json:"version,omitempty"`
	// Replicas is the number of engine replicas
	Replicas int32  `json:"replicas"`
	// Storage is the engine storage configuration
	Storage  Storage `json:"storage"`
	// Resources is the resource requirements for the engine container
	Resources Resources `json:"resources,omitempty"`
	// Config is the engine configuration
	Config   string `json:"config,omitempty"`
	// UserSecretsName is the name of the secret containing the user secrets
	UserSecretsName string `json:"userSecretsName,omitempty"`
}

type Expose struct {
	// Type is the expose type, can be Internal, External or InternetFacing
	Type   string `json:"type"`
	// IPSourceRanges is the list of IP source ranges to allow access from
	IPSourceRanges []string `json:"ipSourceRanges,omitempty"`
}

type Proxy struct {
	// Type is the proxy type
	Type     string `json:"type,omitempty"`
	// Replicas is the number of proxy replicas
	Replicas int32  `json:"replicas,omitempty"`
	// Config is the proxy configuration
	Config   string `json:"config,omitempty"`
	// Expose is the proxy expose configuration
	Expose   Expose `json:"expose,omitempty"`
	// Resources is the resource requirements for the proxy container
	Resources Resources `json:"resources,omitempty"`
}

type DataSource struct {
	// BackupName is the name of the backup to use
	BackupName          string `json:"backupName"`
	// ObjectStorageName is the name of the ObjectStorage CR that defines the
	// storage location
	ObjectStorageName   string `json:"objectStorageName"`
}

type BackupSchedule struct {
	// Enabled is a flag to enable the schedule
	Enabled          bool   `json:"enabled"`
	// Name is the name of the schedule
	Name             string `json:"name"`
	// RetentionCopies is the number of backup copies to retain
	RetentionCopies  int32  `json:"retentionCopies,omitempty"`
	// Schedule is the cron schedule
	Schedule         string `json:"schedule"`
	// ObjectStorageName is the name of the ObjectStorage CR that defines the
	// storage location
	ObjectStorageName   string `json:"objectStorageName"`
}

type Backup struct {
	// Enabled is a flag to enable backups
	Enabled   bool             `json:"enabled"`
	// Schedules is a list of backup schedules
	Schedules []BackupSchedule `json:"schedules,omitempty"`
}

type Monitoring struct {
	// Enabled is a flag to enable monitoring
	Enabled             bool   `json:"enabled"`
	// MonitoringConfigName is the name of the MonitoringConfig CR that defines
	// the monitoring configuration
	MonitoringConfigName string `json:"monitoringConfigName,omitempty"`
	// Resources is the resource requirements for the monitoring container
	Resources Resources `json:"resources,omitempty"`
}

// DatabaseClusterSpec defines the desired state of DatabaseCluster
type DatabaseClusterSpec struct {
	// Paused is a flag to stop the cluster
	Paused bool `json:"paused,omitempty"`
	// Engine is the database engine specification
	Engine Engine `json:"engine"`
	// Proxy is the proxy specification
	Proxy Proxy `json:"proxy,omitempty"`
	// DataSource defines a data source for bootstraping a new cluster
	DataSource DataSource `json:"dataSource,omitempty"`
	// Backup is the backup specification
	Backup Backup `json:"backup,omitempty"`
	// Monitoring is the monitoring specification
	Monitoring Monitoring `json:"monitoring,omitempty"`
}

// DatabaseClusterStatus defines the observed state of DatabaseCluster
type DatabaseClusterStatus struct {
	// Status is the status of the cluster
	Status string `json:"status,omitempty"`
	// Hostname is the hostname where the cluster can be reached
	Hostname string `json:"hostname,omitempty"`
	// Port is the port where the cluster can be reached
	Port int32 `json:"port,omitempty"`
	// Ready is the number of ready pods
	Ready string `json:"ready,omitempty"`
	// Size is the total number of pods
	Size int32 `json:"size,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DatabaseCluster is the Schema for the databaseclusters API
type DatabaseCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseClusterSpec   `json:"spec,omitempty"`
	Status DatabaseClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DatabaseClusterList contains a list of DatabaseCluster
type DatabaseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseCluster{}, &DatabaseClusterList{})
}
