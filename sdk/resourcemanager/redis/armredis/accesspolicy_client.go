//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

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

// AccessPolicyClient contains the methods for the AccessPolicy group.
// Don't use this type directly, use NewAccessPolicyClient() instead.
type AccessPolicyClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAccessPolicyClient creates a new instance of AccessPolicyClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessPolicyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessPolicyClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessPolicyClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateUpdate - Adds an access policy to the redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - accessPolicyName - The name of the access policy that is being added to the Redis cache.
//   - parameters - Parameters supplied to the Create Update Access Policy operation.
//   - options - AccessPolicyClientBeginCreateUpdateOptions contains the optional parameters for the AccessPolicyClient.BeginCreateUpdate
//     method.
func (client *AccessPolicyClient) BeginCreateUpdate(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, parameters CacheAccessPolicy, options *AccessPolicyClientBeginCreateUpdateOptions) (*runtime.Poller[AccessPolicyClientCreateUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createUpdate(ctx, resourceGroupName, cacheName, accessPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccessPolicyClientCreateUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccessPolicyClientCreateUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateUpdate - Adds an access policy to the redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
func (client *AccessPolicyClient) createUpdate(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, parameters CacheAccessPolicy, options *AccessPolicyClientBeginCreateUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AccessPolicyClient.BeginCreateUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createUpdateCreateRequest(ctx, resourceGroupName, cacheName, accessPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createUpdateCreateRequest creates the CreateUpdate request.
func (client *AccessPolicyClient) createUpdateCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, parameters CacheAccessPolicy, options *AccessPolicyClientBeginCreateUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicies/{accessPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the access policy from a redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - accessPolicyName - The name of the access policy that is being added to the Redis cache.
//   - options - AccessPolicyClientBeginDeleteOptions contains the optional parameters for the AccessPolicyClient.BeginDelete
//     method.
func (client *AccessPolicyClient) BeginDelete(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, options *AccessPolicyClientBeginDeleteOptions) (*runtime.Poller[AccessPolicyClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, cacheName, accessPolicyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AccessPolicyClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AccessPolicyClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the access policy from a redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
func (client *AccessPolicyClient) deleteOperation(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, options *AccessPolicyClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "AccessPolicyClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cacheName, accessPolicyName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AccessPolicyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, options *AccessPolicyClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicies/{accessPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the detailed information about an access policy of a redis cache
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - accessPolicyName - The name of the access policy that is being added to the Redis cache.
//   - options - AccessPolicyClientGetOptions contains the optional parameters for the AccessPolicyClient.Get method.
func (client *AccessPolicyClient) Get(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, options *AccessPolicyClientGetOptions) (AccessPolicyClientGetResponse, error) {
	var err error
	const operationName = "AccessPolicyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, cacheName, accessPolicyName, options)
	if err != nil {
		return AccessPolicyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AccessPolicyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AccessPolicyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AccessPolicyClient) getCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, accessPolicyName string, options *AccessPolicyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicies/{accessPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if accessPolicyName == "" {
		return nil, errors.New("parameter accessPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessPolicyName}", url.PathEscape(accessPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AccessPolicyClient) getHandleResponse(resp *http.Response) (AccessPolicyClientGetResponse, error) {
	result := AccessPolicyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CacheAccessPolicy); err != nil {
		return AccessPolicyClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of access policies associated with this redis cache
//
// Generated from API version 2023-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - options - AccessPolicyClientListOptions contains the optional parameters for the AccessPolicyClient.NewListPager method.
func (client *AccessPolicyClient) NewListPager(resourceGroupName string, cacheName string, options *AccessPolicyClientListOptions) *runtime.Pager[AccessPolicyClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessPolicyClientListResponse]{
		More: func(page AccessPolicyClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessPolicyClientListResponse) (AccessPolicyClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccessPolicyClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, cacheName, options)
			}, nil)
			if err != nil {
				return AccessPolicyClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AccessPolicyClient) listCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *AccessPolicyClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/accessPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessPolicyClient) listHandleResponse(resp *http.Response) (AccessPolicyClientListResponse, error) {
	result := AccessPolicyClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CacheAccessPolicyList); err != nil {
		return AccessPolicyClientListResponse{}, err
	}
	return result, nil
}