package v1

import "k8s.io/apimachinery/pkg/runtime"

// DeepCopyInto copies all properties of this object into another object of the
// same type that is provided as a pointer.
func (in *Kubeconfig) DeepCopyInto(out *Kubeconfig) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = KubeconfigSpec{
		ServiceAccount: in.Spec.ServiceAccount,
	}
}

// DeepCopyObject returns a generically typed copy of an object
func (in *Kubeconfig) DeepCopyObject() runtime.Object {
	out := Kubeconfig{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject returns a generically typed copy of an object
func (in *KubeconfigList) DeepCopyObject() runtime.Object {
	out := KubeconfigList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Kubeconfig, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}