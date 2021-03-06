/*
Copyright 2018 Openstorage.org

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
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterPairsGetter has a method to return a ClusterPairInterface.
// A group's client should implement this interface.
type ClusterPairsGetter interface {
	ClusterPairs() ClusterPairInterface
}

// ClusterPairInterface has methods to work with ClusterPair resources.
type ClusterPairInterface interface {
	Create(*v1alpha1.ClusterPair) (*v1alpha1.ClusterPair, error)
	Update(*v1alpha1.ClusterPair) (*v1alpha1.ClusterPair, error)
	UpdateStatus(*v1alpha1.ClusterPair) (*v1alpha1.ClusterPair, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterPair, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterPairList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterPair, err error)
	ClusterPairExpansion
}

// clusterPairs implements ClusterPairInterface
type clusterPairs struct {
	client rest.Interface
}

// newClusterPairs returns a ClusterPairs
func newClusterPairs(c *StorkV1alpha1Client) *clusterPairs {
	return &clusterPairs{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterPair, and returns the corresponding clusterPair object, and an error if there is any.
func (c *clusterPairs) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterPair, err error) {
	result = &v1alpha1.ClusterPair{}
	err = c.client.Get().
		Resource("clusterpairs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterPairs that match those selectors.
func (c *clusterPairs) List(opts v1.ListOptions) (result *v1alpha1.ClusterPairList, err error) {
	result = &v1alpha1.ClusterPairList{}
	err = c.client.Get().
		Resource("clusterpairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterPairs.
func (c *clusterPairs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("clusterpairs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterPair and creates it.  Returns the server's representation of the clusterPair, and an error, if there is any.
func (c *clusterPairs) Create(clusterPair *v1alpha1.ClusterPair) (result *v1alpha1.ClusterPair, err error) {
	result = &v1alpha1.ClusterPair{}
	err = c.client.Post().
		Resource("clusterpairs").
		Body(clusterPair).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterPair and updates it. Returns the server's representation of the clusterPair, and an error, if there is any.
func (c *clusterPairs) Update(clusterPair *v1alpha1.ClusterPair) (result *v1alpha1.ClusterPair, err error) {
	result = &v1alpha1.ClusterPair{}
	err = c.client.Put().
		Resource("clusterpairs").
		Name(clusterPair.Name).
		Body(clusterPair).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterPairs) UpdateStatus(clusterPair *v1alpha1.ClusterPair) (result *v1alpha1.ClusterPair, err error) {
	result = &v1alpha1.ClusterPair{}
	err = c.client.Put().
		Resource("clusterpairs").
		Name(clusterPair.Name).
		SubResource("status").
		Body(clusterPair).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterPair and deletes it. Returns an error if one occurs.
func (c *clusterPairs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterpairs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterPairs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("clusterpairs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterPair.
func (c *clusterPairs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterPair, err error) {
	result = &v1alpha1.ClusterPair{}
	err = c.client.Patch(pt).
		Resource("clusterpairs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
