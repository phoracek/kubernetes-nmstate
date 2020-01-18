package v1alpha1

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkConfigurationEnactmentList contains a list of NodeNetworkConfigurationEnactment
type NodeNetworkConfigurationEnactmentList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeNetworkConfigurationEnactment `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeNetworkConfigurationEnactment is the Schema for the nodenetworkconfigurationenactments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodenetworkconfigurationenactments,shortName=nnce,scope=Cluster
type NodeNetworkConfigurationEnactment struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status NodeNetworkConfigurationEnactmentStatus `json:"status,omitempty"`
}

// NodeNetworkConfigurationEnactmentStatus defines the observed state of NodeNetworkConfigurationEnactment
// +k8s:openapi-gen=true
type NodeNetworkConfigurationEnactmentStatus struct {

	// The desired state rendered for the enactment's node using
	// the policy desiredState as template
	// +optional
	DesiredState State `json:"desiredState,omitempty"`

	// +optional
	Conditions ConditionList `json:"conditions,omitempty"`
}

const (
	EnactmentPolicyLabel                                                = "nmstate.io/policy"
	NodeNetworkConfigurationEnactmentConditionAvailable   ConditionType = "Available"
	NodeNetworkConfigurationEnactmentConditionFailing     ConditionType = "Failing"
	NodeNetworkConfigurationEnactmentConditionProgressing ConditionType = "Progressing"
	NodeNetworkConfigurationEnactmentConditionMatching    ConditionType = "Matching"
)

var NodeNetworkConfigurationEnactmentConditionTypes = [...]ConditionType{
	NodeNetworkConfigurationEnactmentConditionAvailable,
	NodeNetworkConfigurationEnactmentConditionFailing,
	NodeNetworkConfigurationEnactmentConditionProgressing,
	NodeNetworkConfigurationEnactmentConditionMatching,
}

const (
	NodeNetworkConfigurationEnactmentConditionFailedToConfigure                ConditionReason = "FailedToConfigure"
	NodeNetworkConfigurationEnactmentConditionSuccessfullyConfigured           ConditionReason = "SuccessfullyConfigured"
	NodeNetworkConfigurationEnactmentConditionConfigurationProgressing         ConditionReason = "ConfigurationProgressing"
	NodeNetworkConfigurationEnactmentConditionNodeSelectorNotMatching          ConditionReason = "NodeSelectorNotMatching"
	NodeNetworkConfigurationEnactmentConditionNodeSelectorAllSelectorsMatching ConditionReason = "AllSelectorsMatching"
)

func EnactmentKey(node string, policy string) types.NamespacedName {
	return types.NamespacedName{Name: fmt.Sprintf("%s.%s", node, policy)}
}

func NewEnactment(nodeName string, policy NodeNetworkConfigurationPolicy) NodeNetworkConfigurationEnactment {
	enactment := NodeNetworkConfigurationEnactment{
		ObjectMeta: metav1.ObjectMeta{
			Name: EnactmentKey(nodeName, policy.Name).Name,
			OwnerReferences: []metav1.OwnerReference{
				{Name: policy.Name, Kind: policy.TypeMeta.Kind, APIVersion: policy.TypeMeta.APIVersion, UID: policy.UID},
			},
			// Associate policy with the enactment using labels
			Labels: map[string]string{
				EnactmentPolicyLabel: policy.Name,
			},
		},
		Status: NodeNetworkConfigurationEnactmentStatus{
			DesiredState: NewState(""),
			Conditions:   ConditionList{},
		},
	}

	for _, conditionType := range NodeNetworkConfigurationEnactmentConditionTypes {
		enactment.Status.Conditions.Set(conditionType, corev1.ConditionUnknown, "", "")
	}
	return enactment
}

func init() {
	SchemeBuilder.Register(&NodeNetworkConfigurationEnactment{}, &NodeNetworkConfigurationEnactmentList{})
}
