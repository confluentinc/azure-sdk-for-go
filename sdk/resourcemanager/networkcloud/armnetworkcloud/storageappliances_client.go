//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

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

// StorageAppliancesClient contains the methods for the StorageAppliances group.
// Don't use this type directly, use NewStorageAppliancesClient() instead.
type StorageAppliancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageAppliancesClient creates a new instance of StorageAppliancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageAppliancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageAppliancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageAppliancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new storage appliance or update the properties of the existing one. All customer initiated
// requests will be rejected as the life cycle of this resource is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageApplianceName - The name of the storage appliance.
//   - storageApplianceParameters - The request body.
//   - options - StorageAppliancesClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageAppliancesClient.BeginCreateOrUpdate
//     method.
func (client *StorageAppliancesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, storageApplianceName string, storageApplianceParameters StorageAppliance, options *StorageAppliancesClientBeginCreateOrUpdateOptions) (*runtime.Poller[StorageAppliancesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, storageApplianceName, storageApplianceParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageAppliancesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[StorageAppliancesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a new storage appliance or update the properties of the existing one. All customer initiated requests
// will be rejected as the life cycle of this resource is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *StorageAppliancesClient) createOrUpdate(ctx context.Context, resourceGroupName string, storageApplianceName string, storageApplianceParameters StorageAppliance, options *StorageAppliancesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "StorageAppliancesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageApplianceName, storageApplianceParameters, options)
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
func (client *StorageAppliancesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageApplianceName string, storageApplianceParameters StorageAppliance, options *StorageAppliancesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/storageAppliances/{storageApplianceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageApplianceName == "" {
		return nil, errors.New("parameter storageApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageApplianceName}", url.PathEscape(storageApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, storageApplianceParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the provided storage appliance. All customer initiated requests will be rejected as the life cycle
// of this resource is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageApplianceName - The name of the storage appliance.
//   - options - StorageAppliancesClientBeginDeleteOptions contains the optional parameters for the StorageAppliancesClient.BeginDelete
//     method.
func (client *StorageAppliancesClient) BeginDelete(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginDeleteOptions) (*runtime.Poller[StorageAppliancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storageApplianceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageAppliancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[StorageAppliancesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the provided storage appliance. All customer initiated requests will be rejected as the life cycle of this
// resource is managed by the system.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *StorageAppliancesClient) deleteOperation(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "StorageAppliancesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageApplianceName, options)
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
func (client *StorageAppliancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/storageAppliances/{storageApplianceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageApplianceName == "" {
		return nil, errors.New("parameter storageApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageApplianceName}", url.PathEscape(storageApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDisableRemoteVendorManagement - Disable remote vendor management of the provided storage appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageApplianceName - The name of the storage appliance.
//   - options - StorageAppliancesClientBeginDisableRemoteVendorManagementOptions contains the optional parameters for the StorageAppliancesClient.BeginDisableRemoteVendorManagement
//     method.
func (client *StorageAppliancesClient) BeginDisableRemoteVendorManagement(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginDisableRemoteVendorManagementOptions) (*runtime.Poller[StorageAppliancesClientDisableRemoteVendorManagementResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.disableRemoteVendorManagement(ctx, resourceGroupName, storageApplianceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageAppliancesClientDisableRemoteVendorManagementResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[StorageAppliancesClientDisableRemoteVendorManagementResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DisableRemoteVendorManagement - Disable remote vendor management of the provided storage appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *StorageAppliancesClient) disableRemoteVendorManagement(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginDisableRemoteVendorManagementOptions) (*http.Response, error) {
	var err error
	const operationName = "StorageAppliancesClient.BeginDisableRemoteVendorManagement"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.disableRemoteVendorManagementCreateRequest(ctx, resourceGroupName, storageApplianceName, options)
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

// disableRemoteVendorManagementCreateRequest creates the DisableRemoteVendorManagement request.
func (client *StorageAppliancesClient) disableRemoteVendorManagementCreateRequest(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginDisableRemoteVendorManagementOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/storageAppliances/{storageApplianceName}/disableRemoteVendorManagement"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageApplianceName == "" {
		return nil, errors.New("parameter storageApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageApplianceName}", url.PathEscape(storageApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginEnableRemoteVendorManagement - Enable remote vendor management of the provided storage appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageApplianceName - The name of the storage appliance.
//   - options - StorageAppliancesClientBeginEnableRemoteVendorManagementOptions contains the optional parameters for the StorageAppliancesClient.BeginEnableRemoteVendorManagement
//     method.
func (client *StorageAppliancesClient) BeginEnableRemoteVendorManagement(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginEnableRemoteVendorManagementOptions) (*runtime.Poller[StorageAppliancesClientEnableRemoteVendorManagementResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.enableRemoteVendorManagement(ctx, resourceGroupName, storageApplianceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageAppliancesClientEnableRemoteVendorManagementResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[StorageAppliancesClientEnableRemoteVendorManagementResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// EnableRemoteVendorManagement - Enable remote vendor management of the provided storage appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *StorageAppliancesClient) enableRemoteVendorManagement(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginEnableRemoteVendorManagementOptions) (*http.Response, error) {
	var err error
	const operationName = "StorageAppliancesClient.BeginEnableRemoteVendorManagement"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.enableRemoteVendorManagementCreateRequest(ctx, resourceGroupName, storageApplianceName, options)
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

// enableRemoteVendorManagementCreateRequest creates the EnableRemoteVendorManagement request.
func (client *StorageAppliancesClient) enableRemoteVendorManagementCreateRequest(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientBeginEnableRemoteVendorManagementOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/storageAppliances/{storageApplianceName}/enableRemoteVendorManagement"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageApplianceName == "" {
		return nil, errors.New("parameter storageApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageApplianceName}", url.PathEscape(storageApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.StorageApplianceEnableRemoteVendorManagementParameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.StorageApplianceEnableRemoteVendorManagementParameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Get - Get properties of the provided storage appliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageApplianceName - The name of the storage appliance.
//   - options - StorageAppliancesClientGetOptions contains the optional parameters for the StorageAppliancesClient.Get method.
func (client *StorageAppliancesClient) Get(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientGetOptions) (StorageAppliancesClientGetResponse, error) {
	var err error
	const operationName = "StorageAppliancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageApplianceName, options)
	if err != nil {
		return StorageAppliancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAppliancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageAppliancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StorageAppliancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageApplianceName string, options *StorageAppliancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/storageAppliances/{storageApplianceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageApplianceName == "" {
		return nil, errors.New("parameter storageApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageApplianceName}", url.PathEscape(storageApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageAppliancesClient) getHandleResponse(resp *http.Response) (StorageAppliancesClientGetResponse, error) {
	result := StorageAppliancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAppliance); err != nil {
		return StorageAppliancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of storage appliances in the provided resource group.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - StorageAppliancesClientListByResourceGroupOptions contains the optional parameters for the StorageAppliancesClient.NewListByResourceGroupPager
//     method.
func (client *StorageAppliancesClient) NewListByResourceGroupPager(resourceGroupName string, options *StorageAppliancesClientListByResourceGroupOptions) *runtime.Pager[StorageAppliancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAppliancesClientListByResourceGroupResponse]{
		More: func(page StorageAppliancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAppliancesClientListByResourceGroupResponse) (StorageAppliancesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StorageAppliancesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return StorageAppliancesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *StorageAppliancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *StorageAppliancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/storageAppliances"
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
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *StorageAppliancesClient) listByResourceGroupHandleResponse(resp *http.Response) (StorageAppliancesClientListByResourceGroupResponse, error) {
	result := StorageAppliancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageApplianceList); err != nil {
		return StorageAppliancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of storage appliances in the provided subscription.
//
// Generated from API version 2023-07-01
//   - options - StorageAppliancesClientListBySubscriptionOptions contains the optional parameters for the StorageAppliancesClient.NewListBySubscriptionPager
//     method.
func (client *StorageAppliancesClient) NewListBySubscriptionPager(options *StorageAppliancesClientListBySubscriptionOptions) *runtime.Pager[StorageAppliancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAppliancesClientListBySubscriptionResponse]{
		More: func(page StorageAppliancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAppliancesClientListBySubscriptionResponse) (StorageAppliancesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StorageAppliancesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return StorageAppliancesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *StorageAppliancesClient) listBySubscriptionCreateRequest(ctx context.Context, options *StorageAppliancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/storageAppliances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *StorageAppliancesClient) listBySubscriptionHandleResponse(resp *http.Response) (StorageAppliancesClientListBySubscriptionResponse, error) {
	result := StorageAppliancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageApplianceList); err != nil {
		return StorageAppliancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update properties of the provided storage appliance, or update tags associated with the storage appliance
// Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageApplianceName - The name of the storage appliance.
//   - storageApplianceUpdateParameters - The request body.
//   - options - StorageAppliancesClientBeginUpdateOptions contains the optional parameters for the StorageAppliancesClient.BeginUpdate
//     method.
func (client *StorageAppliancesClient) BeginUpdate(ctx context.Context, resourceGroupName string, storageApplianceName string, storageApplianceUpdateParameters StorageAppliancePatchParameters, options *StorageAppliancesClientBeginUpdateOptions) (*runtime.Poller[StorageAppliancesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, storageApplianceName, storageApplianceUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageAppliancesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[StorageAppliancesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update properties of the provided storage appliance, or update tags associated with the storage appliance Properties
// and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *StorageAppliancesClient) update(ctx context.Context, resourceGroupName string, storageApplianceName string, storageApplianceUpdateParameters StorageAppliancePatchParameters, options *StorageAppliancesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "StorageAppliancesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageApplianceName, storageApplianceUpdateParameters, options)
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
func (client *StorageAppliancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageApplianceName string, storageApplianceUpdateParameters StorageAppliancePatchParameters, options *StorageAppliancesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/storageAppliances/{storageApplianceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageApplianceName == "" {
		return nil, errors.New("parameter storageApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageApplianceName}", url.PathEscape(storageApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, storageApplianceUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}