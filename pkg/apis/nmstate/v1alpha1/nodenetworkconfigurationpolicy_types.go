package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkConfigurationPolicyList contains a list of NodeNetworkConfigurationPolicy
type NodeNetworkConfigurationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeNetworkConfigurationPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkConfigurationPolicy is the Schema for the nodenetworkconfigurationpolicies API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodenetworkconfigurationpolicies,shortName=nncp,scope=Cluster
type NodeNetworkConfigurationPolicy struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec NodeNetworkConfigurationPolicySpec `json:"spec,omitempty"`
	// +optional
	Status NodeNetworkConfigurationPolicyStatus `json:"status,omitempty"`
}

// NodeNetworkConfigurationPolicySpec defines the desired state of NodeNetworkConfigurationPolicy
// +k8s:openapi-gen=true
type NodeNetworkConfigurationPolicySpec struct {
	// NodeSelector is a selector which must be true for the policy to be applied to the node.
	// Selector which must match a node's labels for the policy to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// The desired configuration of the policy
	// +optional
	DesiredState State `json:"desiredState,omitempty"`
}

// NodeNetworkConfigurationPolicyStatus defines the observed state of NodeNetworkConfigurationPolicy
// +k8s:openapi-gen=true
type NodeNetworkConfigurationPolicyStatus struct {
	// +optional
	Conditions ConditionList `json:"conditions,omitempty"`
}

const (
	NodeNetworkConfigurationPolicyConditionAvailable ConditionType = "Available"
	NodeNetworkConfigurationPolicyConditionDegraded  ConditionType = "Degraded"
)

var NodeNetworkConfigurationPolicyConditionTypes = [...]ConditionType{
	NodeNetworkConfigurationPolicyConditionAvailable,
	NodeNetworkConfigurationPolicyConditionDegraded,
}

const (
	NodeNetworkConfigurationPolicyConditionFailedToConfigure           ConditionReason = "FailedToConfigure"
	NodeNetworkConfigurationPolicyConditionSuccessfullyConfigured      ConditionReason = "SuccessfullyConfigured"
	NodeNetworkConfigurationPolicyConditionConfigurationProgressing    ConditionReason = "ConfigurationProgressing"
	NodeNetworkConfigurationPolicyConditionConfigurationNoMatchingNode ConditionReason = "NoMatchingNode"
)

func init() {
	SchemeBuilder.Register(&NodeNetworkConfigurationPolicy{}, &NodeNetworkConfigurationPolicyList{})
}
