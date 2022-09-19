/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProjectSSHKeyObservation struct {
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type ProjectSSHKeyParameters struct {

	// The name of the SSH key for identification
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of parent project
	// +kubebuilder:validation:Required
	ProjectID *string `json:"projectId" tf:"project_id,omitempty"`

	// The public key. If this is a file, it
	// +kubebuilder:validation:Required
	PublicKey *string `json:"publicKey" tf:"public_key,omitempty"`
}

// ProjectSSHKeySpec defines the desired state of ProjectSSHKey
type ProjectSSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectSSHKeyParameters `json:"forProvider"`
}

// ProjectSSHKeyStatus defines the observed state of ProjectSSHKey.
type ProjectSSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectSSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectSSHKey is the Schema for the ProjectSSHKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinixjet}
type ProjectSSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSSHKeySpec   `json:"spec"`
	Status            ProjectSSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectSSHKeyList contains a list of ProjectSSHKeys
type ProjectSSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectSSHKey `json:"items"`
}

// Repository type metadata.
var (
	ProjectSSHKey_Kind             = "ProjectSSHKey"
	ProjectSSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectSSHKey_Kind}.String()
	ProjectSSHKey_KindAPIVersion   = ProjectSSHKey_Kind + "." + CRDGroupVersion.String()
	ProjectSSHKey_GroupVersionKind = CRDGroupVersion.WithKind(ProjectSSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectSSHKey{}, &ProjectSSHKeyList{})
}
