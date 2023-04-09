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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/apis/ako/v1alpha2"
	scheme "github.com/vmware/load-balancer-and-ingress-services-for-kubernetes/pkg/client/v1alpha2/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OAuthSamlConfigsGetter has a method to return a OAuthSamlConfigInterface.
// A group's client should implement this interface.
type OAuthSamlConfigsGetter interface {
	OAuthSamlConfigs(namespace string) OAuthSamlConfigInterface
}

// OAuthSamlConfigInterface has methods to work with OAuthSamlConfig resources.
type OAuthSamlConfigInterface interface {
	Create(ctx context.Context, oAuthSamlConfig *v1alpha2.OAuthSamlConfig, opts v1.CreateOptions) (*v1alpha2.OAuthSamlConfig, error)
	Update(ctx context.Context, oAuthSamlConfig *v1alpha2.OAuthSamlConfig, opts v1.UpdateOptions) (*v1alpha2.OAuthSamlConfig, error)
	UpdateStatus(ctx context.Context, oAuthSamlConfig *v1alpha2.OAuthSamlConfig, opts v1.UpdateOptions) (*v1alpha2.OAuthSamlConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.OAuthSamlConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.OAuthSamlConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.OAuthSamlConfig, err error)
	OAuthSamlConfigExpansion
}

// oAuthSamlConfigs implements OAuthSamlConfigInterface
type oAuthSamlConfigs struct {
	client rest.Interface
	ns     string
}

// newOAuthSamlConfigs returns a OAuthSamlConfigs
func newOAuthSamlConfigs(c *AkoV1alpha2Client, namespace string) *oAuthSamlConfigs {
	return &oAuthSamlConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the oAuthSamlConfig, and returns the corresponding oAuthSamlConfig object, and an error if there is any.
func (c *oAuthSamlConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.OAuthSamlConfig, err error) {
	result = &v1alpha2.OAuthSamlConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OAuthSamlConfigs that match those selectors.
func (c *oAuthSamlConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.OAuthSamlConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.OAuthSamlConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oAuthSamlConfigs.
func (c *oAuthSamlConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a oAuthSamlConfig and creates it.  Returns the server's representation of the oAuthSamlConfig, and an error, if there is any.
func (c *oAuthSamlConfigs) Create(ctx context.Context, oAuthSamlConfig *v1alpha2.OAuthSamlConfig, opts v1.CreateOptions) (result *v1alpha2.OAuthSamlConfig, err error) {
	result = &v1alpha2.OAuthSamlConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oAuthSamlConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a oAuthSamlConfig and updates it. Returns the server's representation of the oAuthSamlConfig, and an error, if there is any.
func (c *oAuthSamlConfigs) Update(ctx context.Context, oAuthSamlConfig *v1alpha2.OAuthSamlConfig, opts v1.UpdateOptions) (result *v1alpha2.OAuthSamlConfig, err error) {
	result = &v1alpha2.OAuthSamlConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		Name(oAuthSamlConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oAuthSamlConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *oAuthSamlConfigs) UpdateStatus(ctx context.Context, oAuthSamlConfig *v1alpha2.OAuthSamlConfig, opts v1.UpdateOptions) (result *v1alpha2.OAuthSamlConfig, err error) {
	result = &v1alpha2.OAuthSamlConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		Name(oAuthSamlConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oAuthSamlConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the oAuthSamlConfig and deletes it. Returns an error if one occurs.
func (c *oAuthSamlConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oAuthSamlConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched oAuthSamlConfig.
func (c *oAuthSamlConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.OAuthSamlConfig, err error) {
	result = &v1alpha2.OAuthSamlConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("oauthsamlconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
