package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type KubeconfigSpec struct {
		ServiceAccount string `json:"serviceAccount"`
}

type Kubeconfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubeconfigSpec `json:"spec"`
}

type KubeconfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Kubeconfig `json:"items"`
}