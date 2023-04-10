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

package v1

import (
	"encoding/json"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// BackupStorageFilesystem represents file system storage type.
	BackupStorageFilesystem BackupStorageType = "filesystem"
	// BackupStorageS3 represents s3 storage.
	BackupStorageS3 BackupStorageType = "s3"
	// BackupStorageGCS represents Google Cloud storage.
	BackupStorageGCS BackupStorageType = "gcs"
	// BackupStorageAzure represents azure storage.
	BackupStorageAzure BackupStorageType = "azure"
)

type (
	// RestoreState represents state of restoration.
	RestoreState string
	// BackupStorageType represents backup storage type.
	BackupStorageType string

	// DatabaseClusterRestoreSpec defines the desired state of DatabaseClusterRestore.
	DatabaseClusterRestoreSpec struct {
		DatabaseCluster string        `json:"databaseCluster"`
		DatabaseType    EngineType    `json:"databaseType"`
		BackupName      string        `json:"backupName,omitempty"`
		BackupSource    *BackupSource `json:"backupSource,omitempty"`
		PITR            *PITR         `json:"pitr,omitempty"`
	}
	// PITR represents a specification to configure point in time recovery for a database backup/restore.
	PITR struct {
		BackupSource *BackupSource `json:"backupSource,omitempty"`
		Type         string        `json:"type"`
		Date         *RestoreDate  `json:"date,omitempty"`
		GTID         string        `json:"gtid,omitempty"`
	}
	// BackupSource represents settings of a source where to get a backup to run restoration.
	BackupSource struct {
		Destination           string                     `json:"destination,omitempty"`
		StorageName           string                     `json:"storageName,omitempty"`
		S3                    *BackupStorageProviderSpec `json:"s3,omitempty"`
		Azure                 *BackupStorageProviderSpec `json:"azure,omitempty"`
		StorageType           BackupStorageType          `json:"storageType"`
		Image                 string                     `json:"image,omitempty"`
		SSLSecretName         string                     `json:"sslSecretName,omitempty"`
		SSLInternalSecretName string                     `json:"sslInternalSecretName,omitempty"`
		VaultSecretName       string                     `json:"vaultSecretName,omitempty"`
	}
)

// RestoreDate is a data type for better time.Time support.
// +kubebuilder:validation:Type=string
type RestoreDate struct {
	metav1.Time `json:",inline"`
}

// OpenAPISchemaType returns a schema type for OperAPI specification.
func (RestoreDate) OpenAPISchemaType() []string { return []string{"string"} }

// OpenAPISchemaFormat returns a format for OperAPI specification.
func (RestoreDate) OpenAPISchemaFormat() string { return "" }

// UnmarshalJSON unmarshals JSON.
func (t *RestoreDate) UnmarshalJSON(b []byte) error {
	if len(b) == 4 && string(b) == "null" {
		t.Time = metav1.NewTime(time.Time{})
		return nil
	}

	var str string

	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	pt, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}

	t.Time = metav1.NewTime(pt)

	return nil
}

// DatabaseClusterRestoreStatus defines the observed state of DatabaseClusterRestore.
type DatabaseClusterRestoreStatus struct {
	State         RestoreState       `json:"state,omitempty"`
	CompletedAt   *metav1.Time       `json:"completed,omitempty"`
	LastScheduled *metav1.Time       `json:"lastscheduled,omitempty"`
	Conditions    []metav1.Condition `json:"conditions,omitempty"`
	Message       string             `json:"message,omitempty"`
	Destination   string             `json:"destination,omitempty"`
	StorageName   string             `json:"storageName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName="dbc-restore";"dbc-restores"
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".spec.databaseCluster",description="Cluster name"
// +kubebuilder:printcolumn:name="Storage",type="string",JSONPath=".status.storageName",description="Storage name from pxc spec"
// +kubebuilder:printcolumn:name="Destination",type="string",JSONPath=".status.destination",description="Backup destination"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state",description="Job status"
// +kubebuilder:printcolumn:name="Completed",type="date",JSONPath=".status.completed",description="Completed time"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// DatabaseClusterRestore is the Schema for the databaseclusterrestores API.
type DatabaseClusterRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseClusterRestoreSpec   `json:"spec,omitempty"`
	Status DatabaseClusterRestoreStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DatabaseClusterRestoreList contains a list of DatabaseClusterRestore.
type DatabaseClusterRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseClusterRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseClusterRestore{}, &DatabaseClusterRestoreList{})
}
