//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MachinePoolsClient contains the methods for the MachinePools group.
// Don't use this type directly, use NewMachinePoolsClient() instead.
type MachinePoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMachinePoolsClient creates a new instance of MachinePoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMachinePoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MachinePoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MachinePoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - The operation returns properties of a MachinePool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the MachinePool resource.
//   - parameters - The MachinePool resource.
//   - options - MachinePoolsClientCreateOrUpdateOptions contains the optional parameters for the MachinePoolsClient.CreateOrUpdate
//     method.
func (client *MachinePoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters MachinePool, options *MachinePoolsClientCreateOrUpdateOptions) (MachinePoolsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "MachinePoolsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return MachinePoolsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinePoolsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return MachinePoolsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MachinePoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters MachinePool, options *MachinePoolsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/machinePool/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MachinePoolsClient) createOrUpdateHandleResponse(resp *http.Response) (MachinePoolsClientCreateOrUpdateResponse, error) {
	result := MachinePoolsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachinePool); err != nil {
		return MachinePoolsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - The operation returns nothing.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the MachinePool resource.
//   - options - MachinePoolsClientDeleteOptions contains the optional parameters for the MachinePoolsClient.Delete method.
func (client *MachinePoolsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *MachinePoolsClientDeleteOptions) (MachinePoolsClientDeleteResponse, error) {
	var err error
	const operationName = "MachinePoolsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return MachinePoolsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinePoolsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MachinePoolsClientDeleteResponse{}, err
	}
	return MachinePoolsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MachinePoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *MachinePoolsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/machinePool/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation returns properties of a MachinePool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the MachinePool resource.
//   - options - MachinePoolsClientGetOptions contains the optional parameters for the MachinePoolsClient.Get method.
func (client *MachinePoolsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *MachinePoolsClientGetOptions) (MachinePoolsClientGetResponse, error) {
	var err error
	const operationName = "MachinePoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, options)
	if err != nil {
		return MachinePoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinePoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachinePoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MachinePoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *MachinePoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/machinePool/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MachinePoolsClient) getHandleResponse(resp *http.Response) (MachinePoolsClientGetResponse, error) {
	result := MachinePoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachinePool); err != nil {
		return MachinePoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation returns properties of each MachinePool.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - options - MachinePoolsClientListOptions contains the optional parameters for the MachinePoolsClient.NewListPager method.
func (client *MachinePoolsClient) NewListPager(resourceGroupName string, resourceName string, options *MachinePoolsClientListOptions) *runtime.Pager[MachinePoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MachinePoolsClientListResponse]{
		More: func(page MachinePoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MachinePoolsClientListResponse) (MachinePoolsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MachinePoolsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return MachinePoolsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MachinePoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *MachinePoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftCluster/{resourceName}/machinePools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MachinePoolsClient) listHandleResponse(resp *http.Response) (MachinePoolsClientListResponse, error) {
	result := MachinePoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachinePoolList); err != nil {
		return MachinePoolsClientListResponse{}, err
	}
	return result, nil
}

// Update - The operation returns properties of a MachinePool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - childResourceName - The name of the MachinePool resource.
//   - parameters - The MachinePool resource.
//   - options - MachinePoolsClientUpdateOptions contains the optional parameters for the MachinePoolsClient.Update method.
func (client *MachinePoolsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters MachinePoolUpdate, options *MachinePoolsClientUpdateOptions) (MachinePoolsClientUpdateResponse, error) {
	var err error
	const operationName = "MachinePoolsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, childResourceName, parameters, options)
	if err != nil {
		return MachinePoolsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MachinePoolsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MachinePoolsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *MachinePoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters MachinePoolUpdate, options *MachinePoolsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openshiftclusters/{resourceName}/machinePool/{childResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if childResourceName == "" {
		return nil, errors.New("parameter childResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{childResourceName}", url.PathEscape(childResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MachinePoolsClient) updateHandleResponse(resp *http.Response) (MachinePoolsClientUpdateResponse, error) {
	result := MachinePoolsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MachinePool); err != nil {
		return MachinePoolsClientUpdateResponse{}, err
	}
	return result, nil
}