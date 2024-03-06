//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvirtualmachineimagebuilder

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list REST API operations. It contains a list of operations and a URL nextLink to get the next
	// set of results.
	OperationListResult
}

// TriggersClientCreateOrUpdateResponse contains the response from method TriggersClient.BeginCreateOrUpdate.
type TriggersClientCreateOrUpdateResponse struct {
	// Represents a trigger that can invoke an image template build.
	Trigger
}

// TriggersClientDeleteResponse contains the response from method TriggersClient.BeginDelete.
type TriggersClientDeleteResponse struct {
	// placeholder for future response values
}

// TriggersClientGetResponse contains the response from method TriggersClient.Get.
type TriggersClientGetResponse struct {
	// Represents a trigger that can invoke an image template build.
	Trigger
}

// TriggersClientListByImageTemplateResponse contains the response from method TriggersClient.NewListByImageTemplatePager.
type TriggersClientListByImageTemplateResponse struct {
	// The result of List triggers operation
	TriggerCollection
}

// VirtualMachineImageTemplatesClientCancelResponse contains the response from method VirtualMachineImageTemplatesClient.BeginCancel.
type VirtualMachineImageTemplatesClientCancelResponse struct {
	// placeholder for future response values
}

// VirtualMachineImageTemplatesClientCreateOrUpdateResponse contains the response from method VirtualMachineImageTemplatesClient.BeginCreateOrUpdate.
type VirtualMachineImageTemplatesClientCreateOrUpdateResponse struct {
	// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
	ImageTemplate
}

// VirtualMachineImageTemplatesClientDeleteResponse contains the response from method VirtualMachineImageTemplatesClient.BeginDelete.
type VirtualMachineImageTemplatesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineImageTemplatesClientGetResponse contains the response from method VirtualMachineImageTemplatesClient.Get.
type VirtualMachineImageTemplatesClientGetResponse struct {
	// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
	ImageTemplate
}

// VirtualMachineImageTemplatesClientGetRunOutputResponse contains the response from method VirtualMachineImageTemplatesClient.GetRunOutput.
type VirtualMachineImageTemplatesClientGetRunOutputResponse struct {
	// Represents an output that was created by running an image template.
	RunOutput
}

// VirtualMachineImageTemplatesClientListByResourceGroupResponse contains the response from method VirtualMachineImageTemplatesClient.NewListByResourceGroupPager.
type VirtualMachineImageTemplatesClientListByResourceGroupResponse struct {
	// The result of List image templates operation
	ImageTemplateListResult
}

// VirtualMachineImageTemplatesClientListResponse contains the response from method VirtualMachineImageTemplatesClient.NewListPager.
type VirtualMachineImageTemplatesClientListResponse struct {
	// The result of List image templates operation
	ImageTemplateListResult
}

// VirtualMachineImageTemplatesClientListRunOutputsResponse contains the response from method VirtualMachineImageTemplatesClient.NewListRunOutputsPager.
type VirtualMachineImageTemplatesClientListRunOutputsResponse struct {
	// The result of List run outputs operation
	RunOutputCollection
}

// VirtualMachineImageTemplatesClientRunResponse contains the response from method VirtualMachineImageTemplatesClient.BeginRun.
type VirtualMachineImageTemplatesClientRunResponse struct {
	// placeholder for future response values
}

// VirtualMachineImageTemplatesClientUpdateResponse contains the response from method VirtualMachineImageTemplatesClient.BeginUpdate.
type VirtualMachineImageTemplatesClientUpdateResponse struct {
	// Image template is an ARM resource managed by Microsoft.VirtualMachineImages provider
	ImageTemplate
}
