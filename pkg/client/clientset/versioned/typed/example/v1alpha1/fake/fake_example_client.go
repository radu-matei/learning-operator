/*
This code is auto-generated by k8s.io/code-generator
*/package fake

import (
	v1alpha1 "github.com/radu-matei/learning-operator/pkg/client/clientset/versioned/typed/example/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeExampleV1alpha1 struct {
	*testing.Fake
}

func (c *FakeExampleV1alpha1) SimpleExamples(namespace string) v1alpha1.SimpleExampleInterface {
	return &FakeSimpleExamples{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeExampleV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
