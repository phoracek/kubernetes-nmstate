package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NodeNetworkStateStatus is the status of the NodeNetworkState of a specific node
// +k8s:openapi-gen=true
type NodeNetworkStateStatus struct {
	// +optional
	CurrentState *State `json:"currentState,omitempty"`
	// +optional
	LastSuccessfulUpdateTime *metav1.Time `json:"lastSuccessfulUpdateTime,omitempty"`
	// +optional
	Conditions ConditionList `json:"conditions,omitempty"`
}

const (
	NodeNetworkStateConditionAvailable ConditionType = "Available"
	NodeNetworkStateConditionDegraded  ConditionType = "Degraded"
)

const (
	NodeNetworkStateConditionFailedToObtainCurrentState        ConditionReason = "FailedToObtainCurrentState"
	NodeNetworkStateConditionSuccessfullyObtainedCurrentStatus ConditionReason = "SuccessfullyObtainedCurrentStatus"
)

var NodeNetworkStateConditionTypes = [...]ConditionType{
	NodeNetworkStateConditionAvailable,
	NodeNetworkStateConditionDegraded,
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkState is the Schema for the nodenetworkstates API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodenetworkstates,shortName=nns,scope=Cluster
type NodeNetworkState struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status NodeNetworkStateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkStateList contains a list of NodeNetworkState
type NodeNetworkStateList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeNetworkState `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeNetworkState{}, &NodeNetworkStateList{})
}
