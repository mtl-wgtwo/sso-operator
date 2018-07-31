// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/jenkins-x/sso-operator/pkg/client/clientset/versioned"
	jenkinsv1 "github.com/jenkins-x/sso-operator/pkg/client/clientset/versioned/typed/jenkins.io/v1"
	fakejenkinsv1 "github.com/jenkins-x/sso-operator/pkg/client/clientset/versioned/typed/jenkins.io/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// JenkinsV1 retrieves the JenkinsV1Client
func (c *Clientset) JenkinsV1() jenkinsv1.JenkinsV1Interface {
	return &fakejenkinsv1.FakeJenkinsV1{Fake: &c.Fake}
}

// Jenkins retrieves the JenkinsV1Client
func (c *Clientset) Jenkins() jenkinsv1.JenkinsV1Interface {
	return &fakejenkinsv1.FakeJenkinsV1{Fake: &c.Fake}
}