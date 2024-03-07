//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

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

// DiskPoolsClient contains the methods for the DiskPools group.
// Don't use this type directly, use NewDiskPoolsClient() instead.
type DiskPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDiskPoolsClient creates a new instance of DiskPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDiskPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiskPoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DiskPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update Disk pool. This create or update operation can take 15 minutes to complete. This
// is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - diskPoolCreatePayload - Request payload for Disk Pool create operation
//   - options - DiskPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the DiskPoolsClient.BeginCreateOrUpdate
//     method.
func (client *DiskPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolCreatePayload DiskPoolCreate, options *DiskPoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DiskPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, diskPoolName, diskPoolCreatePayload, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiskPoolsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DiskPoolsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or Update Disk pool. This create or update operation can take 15 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *DiskPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolCreatePayload DiskPoolCreate, options *DiskPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DiskPoolsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, diskPoolName, diskPoolCreatePayload, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiskPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolCreatePayload DiskPoolCreate, options *DiskPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, diskPoolCreatePayload); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeallocate - Shuts down the Disk Pool and releases the compute resources. You are not billed for the compute resources
// that this Disk Pool uses. This operation can take 10 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - options - DiskPoolsClientBeginDeallocateOptions contains the optional parameters for the DiskPoolsClient.BeginDeallocate
//     method.
func (client *DiskPoolsClient) BeginDeallocate(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeallocateOptions) (*runtime.Poller[DiskPoolsClientDeallocateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deallocate(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiskPoolsClientDeallocateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DiskPoolsClientDeallocateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Deallocate - Shuts down the Disk Pool and releases the compute resources. You are not billed for the compute resources
// that this Disk Pool uses. This operation can take 10 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *DiskPoolsClient) deallocate(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeallocateOptions) (*http.Response, error) {
	var err error
	const operationName = "DiskPoolsClient.BeginDeallocate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deallocateCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deallocateCreateRequest creates the Deallocate request.
func (client *DiskPoolsClient) deallocateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeallocateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/deallocate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Delete a Disk pool; attached disks are not affected. This delete operation can take 10 minutes to complete.
// This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - options - DiskPoolsClientBeginDeleteOptions contains the optional parameters for the DiskPoolsClient.BeginDelete method.
func (client *DiskPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeleteOptions) (*runtime.Poller[DiskPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiskPoolsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DiskPoolsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a Disk pool; attached disks are not affected. This delete operation can take 10 minutes to complete. This
// is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *DiskPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DiskPoolsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, diskPoolName, options)
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
func (client *DiskPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Disk pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - options - DiskPoolsClientGetOptions contains the optional parameters for the DiskPoolsClient.Get method.
func (client *DiskPoolsClient) Get(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientGetOptions) (DiskPoolsClientGetResponse, error) {
	var err error
	const operationName = "DiskPoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return DiskPoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DiskPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DiskPoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DiskPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiskPoolsClient) getHandleResponse(resp *http.Response) (DiskPoolsClientGetResponse, error) {
	result := DiskPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskPool); err != nil {
		return DiskPoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of DiskPools in a resource group.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DiskPoolsClientListByResourceGroupOptions contains the optional parameters for the DiskPoolsClient.NewListByResourceGroupPager
//     method.
func (client *DiskPoolsClient) NewListByResourceGroupPager(resourceGroupName string, options *DiskPoolsClientListByResourceGroupOptions) *runtime.Pager[DiskPoolsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DiskPoolsClientListByResourceGroupResponse]{
		More: func(page DiskPoolsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskPoolsClientListByResourceGroupResponse) (DiskPoolsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DiskPoolsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DiskPoolsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DiskPoolsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DiskPoolsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools"
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
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DiskPoolsClient) listByResourceGroupHandleResponse(resp *http.Response) (DiskPoolsClientListByResourceGroupResponse, error) {
	result := DiskPoolsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskPoolListResult); err != nil {
		return DiskPoolsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of Disk Pools in a subscription
//
// Generated from API version 2021-08-01
//   - options - DiskPoolsClientListBySubscriptionOptions contains the optional parameters for the DiskPoolsClient.NewListBySubscriptionPager
//     method.
func (client *DiskPoolsClient) NewListBySubscriptionPager(options *DiskPoolsClientListBySubscriptionOptions) *runtime.Pager[DiskPoolsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DiskPoolsClientListBySubscriptionResponse]{
		More: func(page DiskPoolsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskPoolsClientListBySubscriptionResponse) (DiskPoolsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DiskPoolsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DiskPoolsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DiskPoolsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DiskPoolsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.StoragePool/diskPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DiskPoolsClient) listBySubscriptionHandleResponse(resp *http.Response) (DiskPoolsClientListBySubscriptionResponse, error) {
	result := DiskPoolsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskPoolListResult); err != nil {
		return DiskPoolsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListOutboundNetworkDependenciesEndpointsPager - Gets the network endpoints of all outbound dependencies of a Disk Pool
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - options - DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the DiskPoolsClient.NewListOutboundNetworkDependenciesEndpointsPager
//     method.
func (client *DiskPoolsClient) NewListOutboundNetworkDependenciesEndpointsPager(resourceGroupName string, diskPoolName string, options *DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions) *runtime.Pager[DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse] {
	return runtime.NewPager(runtime.PagingHandler[DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse]{
		More: func(page DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse) (DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DiskPoolsClient.NewListOutboundNetworkDependenciesEndpointsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listOutboundNetworkDependenciesEndpointsCreateRequest(ctx, resourceGroupName, diskPoolName, options)
			}, nil)
			if err != nil {
				return DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}, err
			}
			return client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listOutboundNetworkDependenciesEndpointsCreateRequest creates the ListOutboundNetworkDependenciesEndpoints request.
func (client *DiskPoolsClient) listOutboundNetworkDependenciesEndpointsCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientListOutboundNetworkDependenciesEndpointsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/outboundNetworkDependenciesEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOutboundNetworkDependenciesEndpointsHandleResponse handles the ListOutboundNetworkDependenciesEndpoints response.
func (client *DiskPoolsClient) listOutboundNetworkDependenciesEndpointsHandleResponse(resp *http.Response) (DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse, error) {
	result := DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEnvironmentEndpointList); err != nil {
		return DiskPoolsClientListOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	return result, nil
}

// BeginStart - The operation to start a Disk Pool. This start operation can take 10 minutes to complete. This is expected
// service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - options - DiskPoolsClientBeginStartOptions contains the optional parameters for the DiskPoolsClient.BeginStart method.
func (client *DiskPoolsClient) BeginStart(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginStartOptions) (*runtime.Poller[DiskPoolsClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiskPoolsClientStartResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DiskPoolsClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - The operation to start a Disk Pool. This start operation can take 10 minutes to complete. This is expected service
// behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *DiskPoolsClient) start(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "DiskPoolsClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *DiskPoolsClient) startCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginUpdate - Update a Disk pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - diskPoolUpdatePayload - Request payload for Disk Pool update operation.
//   - options - DiskPoolsClientBeginUpdateOptions contains the optional parameters for the DiskPoolsClient.BeginUpdate method.
func (client *DiskPoolsClient) BeginUpdate(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolUpdatePayload DiskPoolUpdate, options *DiskPoolsClientBeginUpdateOptions) (*runtime.Poller[DiskPoolsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, diskPoolName, diskPoolUpdatePayload, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiskPoolsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DiskPoolsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update a Disk pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *DiskPoolsClient) update(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolUpdatePayload DiskPoolUpdate, options *DiskPoolsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DiskPoolsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, diskPoolName, diskPoolUpdatePayload, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *DiskPoolsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, diskPoolUpdatePayload DiskPoolUpdate, options *DiskPoolsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, diskPoolUpdatePayload); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpgrade - Upgrade replaces the underlying virtual machine hosts one at a time. This operation can take 10-15 minutes
// to complete. This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - diskPoolName - The name of the Disk Pool.
//   - options - DiskPoolsClientBeginUpgradeOptions contains the optional parameters for the DiskPoolsClient.BeginUpgrade method.
func (client *DiskPoolsClient) BeginUpgrade(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginUpgradeOptions) (*runtime.Poller[DiskPoolsClientUpgradeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.upgrade(ctx, resourceGroupName, diskPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DiskPoolsClientUpgradeResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DiskPoolsClientUpgradeResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Upgrade - Upgrade replaces the underlying virtual machine hosts one at a time. This operation can take 10-15 minutes to
// complete. This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
func (client *DiskPoolsClient) upgrade(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginUpgradeOptions) (*http.Response, error) {
	var err error
	const operationName = "DiskPoolsClient.BeginUpgrade"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.upgradeCreateRequest(ctx, resourceGroupName, diskPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// upgradeCreateRequest creates the Upgrade request.
func (client *DiskPoolsClient) upgradeCreateRequest(ctx context.Context, resourceGroupName string, diskPoolName string, options *DiskPoolsClientBeginUpgradeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StoragePool/diskPools/{diskPoolName}/upgrade"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if diskPoolName == "" {
		return nil, errors.New("parameter diskPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{diskPoolName}", url.PathEscape(diskPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}