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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BGPConfigObservation struct {

	// The maximum number of route filters allowed per server.
	// The maximum number of route filters allowed per server
	MaxPrefix *float64 `json:"maxPrefix,omitempty" tf:"max_prefix,omitempty"`

	// status of BGP configuration in the project.
	// Status of BGP configuration in the project
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BGPConfigParameters struct {

	// Autonomous System Number for local BGP deployment.
	// Autonomous System Number for local BGP deployment
	// +kubebuilder:validation:Required
	Asn *float64 `json:"asn" tf:"asn,omitempty"`

	// private or public, the private is likely to be usable immediately, the
	// public will need to be reviewed by Equinix Metal engineers.
	// "local" or "global", the local is likely to be usable immediately, the global will need to be review by Equinix Metal engineers
	// +kubebuilder:validation:Required
	DeploymentType *string `json:"deploymentType" tf:"deployment_type,omitempty"`

	// Password for BGP session in plaintext (not a checksum).
	// Password for BGP session in plaintext (not a checksum)
	// +kubebuilder:validation:Optional
	Md5SecretRef *v1.SecretKeySelector `json:"md5SecretRef,omitempty" tf:"-"`
}

type ProjectObservation struct {

	// Optional BGP settings. Refer to Equinix Metal guide for BGP.
	// Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/)
	// +kubebuilder:validation:Optional
	BGPConfig []BGPConfigObservation `json:"bgpConfig,omitempty" tf:"bgp_config,omitempty"`

	// The timestamp for when the project was created.
	// The timestamp for when the project was created
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The unique ID of the project.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp for the last time the project was updated.
	// The timestamp for the last time the project was updated
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type ProjectParameters struct {

	// Optional BGP settings. Refer to Equinix Metal guide for BGP.
	// Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/)
	// +kubebuilder:validation:Optional
	BGPConfig []BGPConfigParameters `json:"bgpConfig,omitempty" tf:"bgp_config,omitempty"`

	// Enable or disable Backend Transfer, default is false.
	// Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
	// +kubebuilder:validation:Optional
	BackendTransfer *bool `json:"backendTransfer,omitempty" tf:"backend_transfer,omitempty"`

	// The name of the project.
	// The name of the project
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The UUID of organization under which you want to create the project. If you
	// leave it out, the project will be create under your the default organization of your account.
	// The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account
	// +crossplane:generate:reference:type=Organization
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Reference to a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDRef *v1.Reference `json:"organizationIdRef,omitempty" tf:"-"`

	// Selector for a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// The UUID of payment method for this project. The payment method and the
	// project need to belong to the same organization (passed with organization_id, or default).
	// The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with organization_id, or default)
	// +kubebuilder:validation:Optional
	PaymentMethodID *string `json:"paymentMethodId,omitempty" tf:"payment_method_id,omitempty"`
}

// ProjectSpec defines the desired state of Project
type ProjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectParameters `json:"forProvider"`
}

// ProjectStatus defines the observed state of Project.
type ProjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Project is the Schema for the Projects API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectSpec   `json:"spec"`
	Status            ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Projects
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Repository type metadata.
var (
	Project_Kind             = "Project"
	Project_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Project_Kind}.String()
	Project_KindAPIVersion   = Project_Kind + "." + CRDGroupVersion.String()
	Project_GroupVersionKind = CRDGroupVersion.WithKind(Project_Kind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}
