package v1

import (
	"github.com/asdfsx/k8s-example-controller/pkg/api/types/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type KubeconfigInterface interface {
	List(opts metav1.ListOptions) (*v1.KubeconfigList, error)
	Get(name string, options metav1.GetOptions) (*v1.Kubeconfig, error)
	Create(*v1.Kubeconfig) (*v1.Kubeconfig, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type kubeConfigClient struct {
	restClient rest.Interface
	ns         string
}

func (c *kubeConfigClient) List(opts metav1.ListOptions) (*v1.KubeconfigList, error) {
	result := v1.KubeconfigList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("kubeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *kubeConfigClient) Get(name string, opts metav1.GetOptions) (*v1.Kubeconfig, error) {
	result := v1.Kubeconfig{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("kubeconfigs").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *kubeConfigClient) Create(project *v1.Kubeconfig) (*v1.Kubeconfig, error) {
	result := v1.Kubeconfig{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("kubeconfigs").
		Body(project).
		Do().
		Into(&result)

	return &result, err
}

func (c *kubeConfigClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("kubeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}