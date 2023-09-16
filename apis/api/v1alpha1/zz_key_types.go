/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EnvironmentObservation struct {

	// The unique identifier for the environment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EnvironmentParameters struct {

	// The unique identifier for the environment.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type KeyObservation struct {

	// A free-form description of the API key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Defaults to `false`.
	DisableWaitForReady *bool `json:"disableWaitForReady,omitempty" tf:"disable_wait_for_ready,omitempty"`

	// A human-readable name for the API key.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource associated with this object. The only resource that is supported is 'cmk.v2.Cluster', 'srcm.v2.Cluster'.
	ManagedResource []ManagedResourceObservation `json:"managedResource,omitempty" tf:"managed_resource,omitempty"`

	// The owner to which the API Key belongs. The owner can be one of 'iam.v2.User', 'iam.v2.ServiceAccount'.
	Owner []OwnerObservation `json:"owner,omitempty" tf:"owner,omitempty"`
}

type KeyParameters struct {

	// A free-form description of the API key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Defaults to `false`.
	// +kubebuilder:validation:Optional
	DisableWaitForReady *bool `json:"disableWaitForReady,omitempty" tf:"disable_wait_for_ready,omitempty"`

	// A human-readable name for the API key.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The resource associated with this object. The only resource that is supported is 'cmk.v2.Cluster', 'srcm.v2.Cluster'.
	// +kubebuilder:validation:Optional
	ManagedResource []ManagedResourceParameters `json:"managedResource,omitempty" tf:"managed_resource,omitempty"`

	// The owner to which the API Key belongs. The owner can be one of 'iam.v2.User', 'iam.v2.ServiceAccount'.
	// +kubebuilder:validation:Optional
	Owner []OwnerParameters `json:"owner,omitempty" tf:"owner,omitempty"`
}

type ManagedResourceObservation struct {

	// The API version of the referred owner.
	APIVersion *string `json:"apiVersion,omitempty" tf:"api_version,omitempty"`

	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	Environment []EnvironmentObservation `json:"environment,omitempty" tf:"environment,omitempty"`

	// The unique identifier for the referred resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The kind of the referred resource.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`
}

type ManagedResourceParameters struct {

	// The API version of the referred owner.
	// +kubebuilder:validation:Required
	APIVersion *string `json:"apiVersion" tf:"api_version,omitempty"`

	// Environment objects represent an isolated namespace for your Confluent resources for organizational purposes.
	// +kubebuilder:validation:Required
	Environment []EnvironmentParameters `json:"environment" tf:"environment,omitempty"`

	// The unique identifier for the referred resource.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// The kind of the referred resource.
	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`
}

type OwnerObservation struct {

	// The API version of the referred owner.
	APIVersion *string `json:"apiVersion,omitempty" tf:"api_version,omitempty"`

	// The unique identifier for the referred owner.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The kind of the referred owner.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`
}

type OwnerParameters struct {

	// The API version of the referred owner.
	// +kubebuilder:validation:Required
	APIVersion *string `json:"apiVersion" tf:"api_version,omitempty"`

	// The unique identifier for the referred owner.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// The kind of the referred owner.
	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`
}

// KeySpec defines the desired state of Key
type KeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeyParameters `json:"forProvider"`
}

// KeyStatus defines the observed state of Key.
type KeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Key is the Schema for the Keys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.owner)",message="owner is a required parameter"
	Spec   KeySpec   `json:"spec"`
	Status KeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyList contains a list of Keys
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Key `json:"items"`
}

// Repository type metadata.
var (
	Key_Kind             = "Key"
	Key_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Key_Kind}.String()
	Key_KindAPIVersion   = Key_Kind + "." + CRDGroupVersion.String()
	Key_GroupVersionKind = CRDGroupVersion.WithKind(Key_Kind)
)

func init() {
	SchemeBuilder.Register(&Key{}, &KeyList{})
}
