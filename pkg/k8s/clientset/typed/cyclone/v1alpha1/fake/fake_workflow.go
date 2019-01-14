/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/caicloud/cyclone/pkg/apis/cyclone/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeWorkflows implements WorkflowInterface
type FakeWorkflows struct {
	Fake *FakeCycloneV1alpha1
	ns   string
}

var workflowsResource = schema.GroupVersionResource{Group: "cyclone.io", Version: "v1alpha1", Resource: "workflows"}

var workflowsKind = schema.GroupVersionKind{Group: "cyclone.io", Version: "v1alpha1", Kind: "Workflow"}

// Get takes name of the workflow, and returns the corresponding workflow object, and an error if there is any.
func (c *FakeWorkflows) Get(name string, options v1.GetOptions) (result *v1alpha1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workflowsResource, c.ns, name), &v1alpha1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workflow), err
}

// List takes label and field selectors, and returns the list of Workflows that match those selectors.
func (c *FakeWorkflows) List(opts v1.ListOptions) (result *v1alpha1.WorkflowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workflowsResource, workflowsKind, c.ns, opts), &v1alpha1.WorkflowList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkflowList{}
	for _, item := range obj.(*v1alpha1.WorkflowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workflows.
func (c *FakeWorkflows) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workflowsResource, c.ns, opts))

}

// Create takes the representation of a workflow and creates it.  Returns the server's representation of the workflow, and an error, if there is any.
func (c *FakeWorkflows) Create(workflow *v1alpha1.Workflow) (result *v1alpha1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workflowsResource, c.ns, workflow), &v1alpha1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workflow), err
}

// Update takes the representation of a workflow and updates it. Returns the server's representation of the workflow, and an error, if there is any.
func (c *FakeWorkflows) Update(workflow *v1alpha1.Workflow) (result *v1alpha1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workflowsResource, c.ns, workflow), &v1alpha1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workflow), err
}

// Delete takes name of the workflow and deletes it. Returns an error if one occurs.
func (c *FakeWorkflows) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workflowsResource, c.ns, name), &v1alpha1.Workflow{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkflows) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workflowsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkflowList{})
	return err
}

// Patch applies the patch and returns the patched workflow.
func (c *FakeWorkflows) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Workflow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workflowsResource, c.ns, name, data, subresources...), &v1alpha1.Workflow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workflow), err
}