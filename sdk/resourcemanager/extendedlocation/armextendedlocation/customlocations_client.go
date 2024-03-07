//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armextendedlocation

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

// CustomLocationsClient contains the methods for the CustomLocations group.
// Don't use this type directly, use NewCustomLocationsClient() instead.
type CustomLocationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomLocationsClient creates a new instance of CustomLocationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomLocationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomLocationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Custom Location in the specified Subscription and Resource Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - parameters - Parameters supplied to create or update a Custom Location.
//   - options - CustomLocationsClientBeginCreateOrUpdateOptions contains the optional parameters for the CustomLocationsClient.BeginCreateOrUpdate
//     method.
func (client *CustomLocationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[CustomLocationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CustomLocationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CustomLocationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a Custom Location in the specified Subscription and Resource Group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
func (client *CustomLocationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "CustomLocationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
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
func (client *CustomLocationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocation, options *CustomLocationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - options - CustomLocationsClientBeginDeleteOptions contains the optional parameters for the CustomLocationsClient.BeginDelete
//     method.
func (client *CustomLocationsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientBeginDeleteOptions) (*runtime.Poller[CustomLocationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[CustomLocationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[CustomLocationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the Custom Location with the specified Resource Name, Resource Group, and Subscription Id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
func (client *CustomLocationsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "CustomLocationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomLocationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// FindTargetResourceGroup - Returns the target resource group associated with the resource sync rules of the Custom Location
// that match the rules passed in with the Find Target Resource Group Request.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - parameters - Parameters of the find target resource group request.
//   - options - CustomLocationsClientFindTargetResourceGroupOptions contains the optional parameters for the CustomLocationsClient.FindTargetResourceGroup
//     method.
func (client *CustomLocationsClient) FindTargetResourceGroup(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocationFindTargetResourceGroupProperties, options *CustomLocationsClientFindTargetResourceGroupOptions) (CustomLocationsClientFindTargetResourceGroupResponse, error) {
	var err error
	const operationName = "CustomLocationsClient.FindTargetResourceGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.findTargetResourceGroupCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return CustomLocationsClientFindTargetResourceGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomLocationsClientFindTargetResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CustomLocationsClientFindTargetResourceGroupResponse{}, err
	}
	resp, err := client.findTargetResourceGroupHandleResponse(httpResp)
	return resp, err
}

// findTargetResourceGroupCreateRequest creates the FindTargetResourceGroup request.
func (client *CustomLocationsClient) findTargetResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters CustomLocationFindTargetResourceGroupProperties, options *CustomLocationsClientFindTargetResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/findTargetResourceGroup"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// findTargetResourceGroupHandleResponse handles the FindTargetResourceGroup response.
func (client *CustomLocationsClient) findTargetResourceGroupHandleResponse(resp *http.Response) (CustomLocationsClientFindTargetResourceGroupResponse, error) {
	result := CustomLocationsClientFindTargetResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationFindTargetResourceGroupResult); err != nil {
		return CustomLocationsClientFindTargetResourceGroupResponse{}, err
	}
	return result, nil
}

// Get - Gets the details of the customLocation with a specified resource group and name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - options - CustomLocationsClientGetOptions contains the optional parameters for the CustomLocationsClient.Get method.
func (client *CustomLocationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientGetOptions) (CustomLocationsClientGetResponse, error) {
	var err error
	const operationName = "CustomLocationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return CustomLocationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomLocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CustomLocationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CustomLocationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomLocationsClient) getHandleResponse(resp *http.Response) (CustomLocationsClientGetResponse, error) {
	result := CustomLocationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocation); err != nil {
		return CustomLocationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of Custom Locations in the specified subscription and resource group. The operation
// returns properties of each Custom Location.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - CustomLocationsClientListByResourceGroupOptions contains the optional parameters for the CustomLocationsClient.NewListByResourceGroupPager
//     method.
func (client *CustomLocationsClient) NewListByResourceGroupPager(resourceGroupName string, options *CustomLocationsClientListByResourceGroupOptions) *runtime.Pager[CustomLocationsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomLocationsClientListByResourceGroupResponse]{
		More: func(page CustomLocationsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomLocationsClientListByResourceGroupResponse) (CustomLocationsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomLocationsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return CustomLocationsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomLocationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomLocationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations"
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
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomLocationsClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomLocationsClientListByResourceGroupResponse, error) {
	result := CustomLocationsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationListResult); err != nil {
		return CustomLocationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of Custom Locations in the specified subscription. The operation returns properties
// of each Custom Location
//
// Generated from API version 2021-08-31-preview
//   - options - CustomLocationsClientListBySubscriptionOptions contains the optional parameters for the CustomLocationsClient.NewListBySubscriptionPager
//     method.
func (client *CustomLocationsClient) NewListBySubscriptionPager(options *CustomLocationsClientListBySubscriptionOptions) *runtime.Pager[CustomLocationsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomLocationsClientListBySubscriptionResponse]{
		More: func(page CustomLocationsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomLocationsClientListBySubscriptionResponse) (CustomLocationsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomLocationsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CustomLocationsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomLocationsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomLocationsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ExtendedLocation/customLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomLocationsClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomLocationsClientListBySubscriptionResponse, error) {
	result := CustomLocationsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationListResult); err != nil {
		return CustomLocationsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// NewListEnabledResourceTypesPager - Gets the list of the Enabled Resource Types.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - options - CustomLocationsClientListEnabledResourceTypesOptions contains the optional parameters for the CustomLocationsClient.NewListEnabledResourceTypesPager
//     method.
func (client *CustomLocationsClient) NewListEnabledResourceTypesPager(resourceGroupName string, resourceName string, options *CustomLocationsClientListEnabledResourceTypesOptions) *runtime.Pager[CustomLocationsClientListEnabledResourceTypesResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomLocationsClientListEnabledResourceTypesResponse]{
		More: func(page CustomLocationsClientListEnabledResourceTypesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomLocationsClientListEnabledResourceTypesResponse) (CustomLocationsClientListEnabledResourceTypesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomLocationsClient.NewListEnabledResourceTypesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listEnabledResourceTypesCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return CustomLocationsClientListEnabledResourceTypesResponse{}, err
			}
			return client.listEnabledResourceTypesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listEnabledResourceTypesCreateRequest creates the ListEnabledResourceTypes request.
func (client *CustomLocationsClient) listEnabledResourceTypesCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *CustomLocationsClientListEnabledResourceTypesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}/enabledResourceTypes"
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
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listEnabledResourceTypesHandleResponse handles the ListEnabledResourceTypes response.
func (client *CustomLocationsClient) listEnabledResourceTypesHandleResponse(resp *http.Response) (CustomLocationsClientListEnabledResourceTypesResponse, error) {
	result := CustomLocationsClientListEnabledResourceTypesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnabledResourceTypesListResult); err != nil {
		return CustomLocationsClientListEnabledResourceTypesResponse{}, err
	}
	return result, nil
}

// NewListOperationsPager - Lists all available Custom Locations operations.
//
// Generated from API version 2021-08-31-preview
//   - options - CustomLocationsClientListOperationsOptions contains the optional parameters for the CustomLocationsClient.NewListOperationsPager
//     method.
func (client *CustomLocationsClient) NewListOperationsPager(options *CustomLocationsClientListOperationsOptions) *runtime.Pager[CustomLocationsClientListOperationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomLocationsClientListOperationsResponse]{
		More: func(page CustomLocationsClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomLocationsClientListOperationsResponse) (CustomLocationsClientListOperationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CustomLocationsClient.NewListOperationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listOperationsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return CustomLocationsClientListOperationsResponse{}, err
			}
			return client.listOperationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *CustomLocationsClient) listOperationsCreateRequest(ctx context.Context, options *CustomLocationsClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ExtendedLocation/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *CustomLocationsClient) listOperationsHandleResponse(resp *http.Response) (CustomLocationsClientListOperationsResponse, error) {
	result := CustomLocationsClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocationOperationsList); err != nil {
		return CustomLocationsClientListOperationsResponse{}, err
	}
	return result, nil
}

// Update - Updates a Custom Location with the specified Resource Name in the specified Resource Group and Subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-31-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - Custom Locations name.
//   - parameters - The updatable fields of an existing Custom Location.
//   - options - CustomLocationsClientUpdateOptions contains the optional parameters for the CustomLocationsClient.Update method.
func (client *CustomLocationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableCustomLocations, options *CustomLocationsClientUpdateOptions) (CustomLocationsClientUpdateResponse, error) {
	var err error
	const operationName = "CustomLocationsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return CustomLocationsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomLocationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CustomLocationsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *CustomLocationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters PatchableCustomLocations, options *CustomLocationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ExtendedLocation/customLocations/{resourceName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *CustomLocationsClient) updateHandleResponse(resp *http.Response) (CustomLocationsClientUpdateResponse, error) {
	result := CustomLocationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomLocation); err != nil {
		return CustomLocationsClientUpdateResponse{}, err
	}
	return result, nil
}