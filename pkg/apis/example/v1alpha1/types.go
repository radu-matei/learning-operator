package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SimpleExample is a specification for a SimpleExample resource
type SimpleExample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SimpleExampleSpec `json:"spec"`
}

// SimpleExampleSpec is the spec for a SimpleExample resource
type SimpleExampleSpec struct {
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SimpleExampleList is a list of SimpleExample resources
type SimpleExampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SimpleExampleSpec `json:"items"`
}
