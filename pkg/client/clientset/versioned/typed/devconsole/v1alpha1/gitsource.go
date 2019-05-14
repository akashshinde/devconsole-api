/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/redhat-developer/devconsole-api/pkg/apis/devconsole/v1alpha1"
	scheme "github.com/redhat-developer/devconsole-api/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GitSourcesGetter has a method to return a GitSourceInterface.
// A group's client should implement this interface.
type GitSourcesGetter interface {
	GitSources(namespace string) GitSourceInterface
}

// GitSourceInterface has methods to work with GitSource resources.
type GitSourceInterface interface {
	Create(*v1alpha1.GitSource) (*v1alpha1.GitSource, error)
	Update(*v1alpha1.GitSource) (*v1alpha1.GitSource, error)
	UpdateStatus(*v1alpha1.GitSource) (*v1alpha1.GitSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GitSource, error)
	List(opts v1.ListOptions) (*v1alpha1.GitSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitSource, err error)
	GitSourceExpansion
}

// gitSources implements GitSourceInterface
type gitSources struct {
	client rest.Interface
	ns     string
}

// newGitSources returns a GitSources
func newGitSources(c *DevconsoleV1alpha1Client, namespace string) *gitSources {
	return &gitSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gitSource, and returns the corresponding gitSource object, and an error if there is any.
func (c *gitSources) Get(name string, options v1.GetOptions) (result *v1alpha1.GitSource, err error) {
	result = &v1alpha1.GitSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GitSources that match those selectors.
func (c *gitSources) List(opts v1.ListOptions) (result *v1alpha1.GitSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GitSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gitSources.
func (c *gitSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gitsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gitSource and creates it.  Returns the server's representation of the gitSource, and an error, if there is any.
func (c *gitSources) Create(gitSource *v1alpha1.GitSource) (result *v1alpha1.GitSource, err error) {
	result = &v1alpha1.GitSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gitsources").
		Body(gitSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gitSource and updates it. Returns the server's representation of the gitSource, and an error, if there is any.
func (c *gitSources) Update(gitSource *v1alpha1.GitSource) (result *v1alpha1.GitSource, err error) {
	result = &v1alpha1.GitSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitsources").
		Name(gitSource.Name).
		Body(gitSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gitSources) UpdateStatus(gitSource *v1alpha1.GitSource) (result *v1alpha1.GitSource, err error) {
	result = &v1alpha1.GitSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitsources").
		Name(gitSource.Name).
		SubResource("status").
		Body(gitSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the gitSource and deletes it. Returns an error if one occurs.
func (c *gitSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gitSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gitSource.
func (c *gitSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitSource, err error) {
	result = &v1alpha1.GitSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gitsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}