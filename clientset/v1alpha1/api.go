package v1alpha1

import (
	. "github.com/mark8s/crd-demo/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ProjectInterface interface {
	List(opts metav1.ListOptions) (*ProjectList, error)
	Get(name string, options metav1.GetOptions) (*Project, error)
	Create(*Project) (*Project, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type ExampleVAlphaInterface interface {
	Projects(namespace string) ProjectInterface
}

type ExampleVAlphaClient struct {
	restClient rest.Interface
}

type ProjectClient struct {
	restClient rest.Interface
	ns         string
}

func NewForConfig(config *rest.Config) (*ExampleVAlphaClient, error) {
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: GroupName, Version: GroupVersion}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&crdConfig)
	if err != nil {
		return nil, err
	}

	return &ExampleVAlphaClient{restClient: client}, nil
}

func (c *ExampleVAlphaClient) Projects(namespace string) ProjectInterface {
	return &ProjectClient{restClient: c.restClient, ns: namespace}
}

func (c *ProjectClient) List(opts metav1.ListOptions) (*ProjectList, error) {
	result := ProjectList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(KindPluralName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *ProjectClient) Get(name string, opts metav1.GetOptions) (*Project, error) {
	result := Project{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(KindPluralName).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *ProjectClient) Create(project *Project) (*Project, error) {
	result := Project{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource(KindPluralName).
		Body(project).
		Do().
		Into(&result)

	return &result, err
}

func (c *ProjectClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource(KindPluralName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
