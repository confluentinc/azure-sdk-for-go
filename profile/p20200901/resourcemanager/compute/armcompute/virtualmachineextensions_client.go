//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineExtensionsClient contains the methods for the VirtualMachineExtensions group.
// Don't use this type directly, use NewVirtualMachineExtensionsClient() instead.
type VirtualMachineExtensionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualMachineExtensionsClient creates a new instance of VirtualMachineExtensionsClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualMachineExtensionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineExtensionsClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armcompute.VirtualMachineExtensionsClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualMachineExtensionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine where the extension should be created or updated.
//   - vmExtensionName - The name of the virtual machine extension.
//   - extensionParameters - Parameters supplied to the Create Virtual Machine Extension operation.
//   - options - VirtualMachineExtensionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualMachineExtensionsClient.BeginCreateOrUpdate
//     method.
func (client *VirtualMachineExtensionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineExtensionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualMachineExtensionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineExtensionsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineExtensionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - The operation to create or update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
func (client *VirtualMachineExtensionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineExtensionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualMachineExtensionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtension, options *VirtualMachineExtensionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}

// BeginDelete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine where the extension should be deleted.
//   - vmExtensionName - The name of the virtual machine extension.
//   - options - VirtualMachineExtensionsClientBeginDeleteOptions contains the optional parameters for the VirtualMachineExtensionsClient.BeginDelete
//     method.
func (client *VirtualMachineExtensionsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsClientBeginDeleteOptions) (*runtime.Poller[VirtualMachineExtensionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vmName, vmExtensionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineExtensionsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineExtensionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
func (client *VirtualMachineExtensionsClient) deleteOperation(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualMachineExtensionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - The operation to get the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine containing the extension.
//   - vmExtensionName - The name of the virtual machine extension.
//   - options - VirtualMachineExtensionsClientGetOptions contains the optional parameters for the VirtualMachineExtensionsClient.Get
//     method.
func (client *VirtualMachineExtensionsClient) Get(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsClientGetOptions) (VirtualMachineExtensionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, options)
	if err != nil {
		return VirtualMachineExtensionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineExtensionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, options *VirtualMachineExtensionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualMachineExtensionsClient) getHandleResponse(resp *http.Response) (VirtualMachineExtensionsClientGetResponse, error) {
	result := VirtualMachineExtensionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtension); err != nil {
		return VirtualMachineExtensionsClientGetResponse{}, err
	}
	return result, nil
}

// List - The operation to get all extensions of a Virtual Machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine containing the extension.
//   - options - VirtualMachineExtensionsClientListOptions contains the optional parameters for the VirtualMachineExtensionsClient.List
//     method.
func (client *VirtualMachineExtensionsClient) List(ctx context.Context, resourceGroupName string, vmName string, options *VirtualMachineExtensionsClientListOptions) (VirtualMachineExtensionsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, vmName, options)
	if err != nil {
		return VirtualMachineExtensionsClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineExtensionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineExtensionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmName string, options *VirtualMachineExtensionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineExtensionsClient) listHandleResponse(resp *http.Response) (VirtualMachineExtensionsClientListResponse, error) {
	result := VirtualMachineExtensionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineExtensionsListResult); err != nil {
		return VirtualMachineExtensionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - vmName - The name of the virtual machine where the extension should be updated.
//   - vmExtensionName - The name of the virtual machine extension.
//   - extensionParameters - Parameters supplied to the Update Virtual Machine Extension operation.
//   - options - VirtualMachineExtensionsClientBeginUpdateOptions contains the optional parameters for the VirtualMachineExtensionsClient.BeginUpdate
//     method.
func (client *VirtualMachineExtensionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineExtensionsClientBeginUpdateOptions) (*runtime.Poller[VirtualMachineExtensionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualMachineExtensionsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualMachineExtensionsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update the extension.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
func (client *VirtualMachineExtensionsClient) update(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineExtensionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vmName, vmExtensionName, extensionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualMachineExtensionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vmName string, vmExtensionName string, extensionParameters VirtualMachineExtensionUpdate, options *VirtualMachineExtensionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmName == "" {
		return nil, errors.New("parameter vmName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmName}", url.PathEscape(vmName))
	if vmExtensionName == "" {
		return nil, errors.New("parameter vmExtensionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmExtensionName}", url.PathEscape(vmExtensionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, extensionParameters)
}