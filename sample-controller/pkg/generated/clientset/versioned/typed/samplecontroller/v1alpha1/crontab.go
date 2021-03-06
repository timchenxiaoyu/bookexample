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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
	scheme "k8s.io/sample-controller/pkg/generated/clientset/versioned/scheme"
)

// CrontabsGetter has a method to return a CrontabInterface.
// A group's client should implement this interface.
type CrontabsGetter interface {
	Crontabs(namespace string) CrontabInterface
}

// CrontabInterface has methods to work with Crontab resources.
type CrontabInterface interface {
	Create(*v1alpha1.Crontab) (*v1alpha1.Crontab, error)
	Update(*v1alpha1.Crontab) (*v1alpha1.Crontab, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Crontab, error)
	List(opts v1.ListOptions) (*v1alpha1.CrontabList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Crontab, err error)
	CrontabExpansion
}

// crontabs implements CrontabInterface
type crontabs struct {
	client rest.Interface
	ns     string
}

// newCrontabs returns a Crontabs
func newCrontabs(c *SamplecontrollerV1alpha1Client, namespace string) *crontabs {
	return &crontabs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the crontab, and returns the corresponding crontab object, and an error if there is any.
func (c *crontabs) Get(name string, options v1.GetOptions) (result *v1alpha1.Crontab, err error) {
	result = &v1alpha1.Crontab{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("crontabs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Crontabs that match those selectors.
func (c *crontabs) List(opts v1.ListOptions) (result *v1alpha1.CrontabList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CrontabList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("crontabs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested crontabs.
func (c *crontabs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("crontabs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a crontab and creates it.  Returns the server's representation of the crontab, and an error, if there is any.
func (c *crontabs) Create(crontab *v1alpha1.Crontab) (result *v1alpha1.Crontab, err error) {
	result = &v1alpha1.Crontab{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("crontabs").
		Body(crontab).
		Do().
		Into(result)
	return
}

// Update takes the representation of a crontab and updates it. Returns the server's representation of the crontab, and an error, if there is any.
func (c *crontabs) Update(crontab *v1alpha1.Crontab) (result *v1alpha1.Crontab, err error) {
	result = &v1alpha1.Crontab{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("crontabs").
		Name(crontab.Name).
		Body(crontab).
		Do().
		Into(result)
	return
}

// Delete takes name of the crontab and deletes it. Returns an error if one occurs.
func (c *crontabs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("crontabs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *crontabs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("crontabs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched crontab.
func (c *crontabs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Crontab, err error) {
	result = &v1alpha1.Crontab{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("crontabs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
