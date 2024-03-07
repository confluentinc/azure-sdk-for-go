//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmixedreality

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

// SpatialAnchorsAccountsClient contains the methods for the SpatialAnchorsAccounts group.
// Don't use this type directly, use NewSpatialAnchorsAccountsClient() instead.
type SpatialAnchorsAccountsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSpatialAnchorsAccountsClient creates a new instance of SpatialAnchorsAccountsClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSpatialAnchorsAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SpatialAnchorsAccountsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SpatialAnchorsAccountsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creating or Updating a Spatial Anchors Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - spatialAnchorsAccount - Spatial Anchors Account parameter.
//   - options - SpatialAnchorsAccountsClientCreateOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Create
//     method.
func (client *SpatialAnchorsAccountsClient) Create(ctx context.Context, resourceGroupName string, accountName string, spatialAnchorsAccount SpatialAnchorsAccount, options *SpatialAnchorsAccountsClientCreateOptions) (SpatialAnchorsAccountsClientCreateResponse, error) {
	var err error
	const operationName = "SpatialAnchorsAccountsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, spatialAnchorsAccount, options)
	if err != nil {
		return SpatialAnchorsAccountsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpatialAnchorsAccountsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SpatialAnchorsAccountsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SpatialAnchorsAccountsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, spatialAnchorsAccount SpatialAnchorsAccount, options *SpatialAnchorsAccountsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, spatialAnchorsAccount); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SpatialAnchorsAccountsClient) createHandleResponse(resp *http.Response) (SpatialAnchorsAccountsClientCreateResponse, error) {
	result := SpatialAnchorsAccountsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpatialAnchorsAccount); err != nil {
		return SpatialAnchorsAccountsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Spatial Anchors Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - options - SpatialAnchorsAccountsClientDeleteOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Delete
//     method.
func (client *SpatialAnchorsAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, options *SpatialAnchorsAccountsClientDeleteOptions) (SpatialAnchorsAccountsClientDeleteResponse, error) {
	var err error
	const operationName = "SpatialAnchorsAccountsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return SpatialAnchorsAccountsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpatialAnchorsAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SpatialAnchorsAccountsClientDeleteResponse{}, err
	}
	return SpatialAnchorsAccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SpatialAnchorsAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *SpatialAnchorsAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve a Spatial Anchors Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - options - SpatialAnchorsAccountsClientGetOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Get
//     method.
func (client *SpatialAnchorsAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, options *SpatialAnchorsAccountsClientGetOptions) (SpatialAnchorsAccountsClientGetResponse, error) {
	var err error
	const operationName = "SpatialAnchorsAccountsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return SpatialAnchorsAccountsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpatialAnchorsAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SpatialAnchorsAccountsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SpatialAnchorsAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *SpatialAnchorsAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SpatialAnchorsAccountsClient) getHandleResponse(resp *http.Response) (SpatialAnchorsAccountsClientGetResponse, error) {
	result := SpatialAnchorsAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpatialAnchorsAccount); err != nil {
		return SpatialAnchorsAccountsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List Resources by Resource Group
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - options - SpatialAnchorsAccountsClientListByResourceGroupOptions contains the optional parameters for the SpatialAnchorsAccountsClient.NewListByResourceGroupPager
//     method.
func (client *SpatialAnchorsAccountsClient) NewListByResourceGroupPager(resourceGroupName string, options *SpatialAnchorsAccountsClientListByResourceGroupOptions) *runtime.Pager[SpatialAnchorsAccountsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpatialAnchorsAccountsClientListByResourceGroupResponse]{
		More: func(page SpatialAnchorsAccountsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpatialAnchorsAccountsClientListByResourceGroupResponse) (SpatialAnchorsAccountsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SpatialAnchorsAccountsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SpatialAnchorsAccountsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SpatialAnchorsAccountsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SpatialAnchorsAccountsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SpatialAnchorsAccountsClient) listByResourceGroupHandleResponse(resp *http.Response) (SpatialAnchorsAccountsClientListByResourceGroupResponse, error) {
	result := SpatialAnchorsAccountsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpatialAnchorsAccountPage); err != nil {
		return SpatialAnchorsAccountsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List Spatial Anchors Accounts by Subscription
//
// Generated from API version 2021-03-01-preview
//   - options - SpatialAnchorsAccountsClientListBySubscriptionOptions contains the optional parameters for the SpatialAnchorsAccountsClient.NewListBySubscriptionPager
//     method.
func (client *SpatialAnchorsAccountsClient) NewListBySubscriptionPager(options *SpatialAnchorsAccountsClientListBySubscriptionOptions) *runtime.Pager[SpatialAnchorsAccountsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpatialAnchorsAccountsClientListBySubscriptionResponse]{
		More: func(page SpatialAnchorsAccountsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpatialAnchorsAccountsClientListBySubscriptionResponse) (SpatialAnchorsAccountsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SpatialAnchorsAccountsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SpatialAnchorsAccountsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SpatialAnchorsAccountsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SpatialAnchorsAccountsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MixedReality/spatialAnchorsAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SpatialAnchorsAccountsClient) listBySubscriptionHandleResponse(resp *http.Response) (SpatialAnchorsAccountsClientListBySubscriptionResponse, error) {
	result := SpatialAnchorsAccountsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpatialAnchorsAccountPage); err != nil {
		return SpatialAnchorsAccountsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListKeys - List Both of the 2 Keys of a Spatial Anchors Account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - options - SpatialAnchorsAccountsClientListKeysOptions contains the optional parameters for the SpatialAnchorsAccountsClient.ListKeys
//     method.
func (client *SpatialAnchorsAccountsClient) ListKeys(ctx context.Context, resourceGroupName string, accountName string, options *SpatialAnchorsAccountsClientListKeysOptions) (SpatialAnchorsAccountsClientListKeysResponse, error) {
	var err error
	const operationName = "SpatialAnchorsAccountsClient.ListKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return SpatialAnchorsAccountsClientListKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpatialAnchorsAccountsClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SpatialAnchorsAccountsClientListKeysResponse{}, err
	}
	resp, err := client.listKeysHandleResponse(httpResp)
	return resp, err
}

// listKeysCreateRequest creates the ListKeys request.
func (client *SpatialAnchorsAccountsClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *SpatialAnchorsAccountsClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *SpatialAnchorsAccountsClient) listKeysHandleResponse(resp *http.Response) (SpatialAnchorsAccountsClientListKeysResponse, error) {
	result := SpatialAnchorsAccountsClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return SpatialAnchorsAccountsClientListKeysResponse{}, err
	}
	return result, nil
}

// RegenerateKeys - Regenerate specified Key of a Spatial Anchors Account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - regenerate - Required information for key regeneration.
//   - options - SpatialAnchorsAccountsClientRegenerateKeysOptions contains the optional parameters for the SpatialAnchorsAccountsClient.RegenerateKeys
//     method.
func (client *SpatialAnchorsAccountsClient) RegenerateKeys(ctx context.Context, resourceGroupName string, accountName string, regenerate AccountKeyRegenerateRequest, options *SpatialAnchorsAccountsClientRegenerateKeysOptions) (SpatialAnchorsAccountsClientRegenerateKeysResponse, error) {
	var err error
	const operationName = "SpatialAnchorsAccountsClient.RegenerateKeys"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.regenerateKeysCreateRequest(ctx, resourceGroupName, accountName, regenerate, options)
	if err != nil {
		return SpatialAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpatialAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SpatialAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	resp, err := client.regenerateKeysHandleResponse(httpResp)
	return resp, err
}

// regenerateKeysCreateRequest creates the RegenerateKeys request.
func (client *SpatialAnchorsAccountsClient) regenerateKeysCreateRequest(ctx context.Context, resourceGroupName string, accountName string, regenerate AccountKeyRegenerateRequest, options *SpatialAnchorsAccountsClientRegenerateKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}/regenerateKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, regenerate); err != nil {
		return nil, err
	}
	return req, nil
}

// regenerateKeysHandleResponse handles the RegenerateKeys response.
func (client *SpatialAnchorsAccountsClient) regenerateKeysHandleResponse(resp *http.Response) (SpatialAnchorsAccountsClientRegenerateKeysResponse, error) {
	result := SpatialAnchorsAccountsClientRegenerateKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccountKeys); err != nil {
		return SpatialAnchorsAccountsClientRegenerateKeysResponse{}, err
	}
	return result, nil
}

// Update - Updating a Spatial Anchors Account
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-01-preview
//   - resourceGroupName - Name of an Azure resource group.
//   - accountName - Name of an Mixed Reality Account.
//   - spatialAnchorsAccount - Spatial Anchors Account parameter.
//   - options - SpatialAnchorsAccountsClientUpdateOptions contains the optional parameters for the SpatialAnchorsAccountsClient.Update
//     method.
func (client *SpatialAnchorsAccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, spatialAnchorsAccount SpatialAnchorsAccount, options *SpatialAnchorsAccountsClientUpdateOptions) (SpatialAnchorsAccountsClientUpdateResponse, error) {
	var err error
	const operationName = "SpatialAnchorsAccountsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, spatialAnchorsAccount, options)
	if err != nil {
		return SpatialAnchorsAccountsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpatialAnchorsAccountsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SpatialAnchorsAccountsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SpatialAnchorsAccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, spatialAnchorsAccount SpatialAnchorsAccount, options *SpatialAnchorsAccountsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MixedReality/spatialAnchorsAccounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, spatialAnchorsAccount); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SpatialAnchorsAccountsClient) updateHandleResponse(resp *http.Response) (SpatialAnchorsAccountsClientUpdateResponse, error) {
	result := SpatialAnchorsAccountsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpatialAnchorsAccount); err != nil {
		return SpatialAnchorsAccountsClientUpdateResponse{}, err
	}
	return result, nil
}