//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// OrdersClient contains the methods for the Orders group.
// Don't use this type directly, use NewOrdersClient() instead.
type OrdersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOrdersClient creates a new instance of OrdersClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOrdersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrdersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OrdersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an order.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The order details of a device.
//   - resourceGroupName - The resource group name.
//   - order - The order to be created or updated.
//   - options - OrdersClientBeginCreateOrUpdateOptions contains the optional parameters for the OrdersClient.BeginCreateOrUpdate
//     method.
func (client *OrdersClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, resourceGroupName string, order Order, options *OrdersClientBeginCreateOrUpdateOptions) (*runtime.Poller[OrdersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, resourceGroupName, order, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrdersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrdersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an order.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *OrdersClient) createOrUpdate(ctx context.Context, deviceName string, resourceGroupName string, order Order, options *OrdersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "OrdersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, resourceGroupName, order, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *OrdersClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, order Order, options *OrdersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, order); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the order related to the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - OrdersClientBeginDeleteOptions contains the optional parameters for the OrdersClient.BeginDelete method.
func (client *OrdersClient) BeginDelete(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientBeginDeleteOptions) (*runtime.Poller[OrdersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrdersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrdersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the order related to the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *OrdersClient) deleteOperation(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "OrdersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, deviceName, resourceGroupName, options)
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
func (client *OrdersClient) deleteCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a specific order by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - OrdersClientGetOptions contains the optional parameters for the OrdersClient.Get method.
func (client *OrdersClient) Get(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientGetOptions) (OrdersClientGetResponse, error) {
	var err error
	const operationName = "OrdersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return OrdersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrdersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrdersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OrdersClient) getCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrdersClient) getHandleResponse(resp *http.Response) (OrdersClientGetResponse, error) {
	result := OrdersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Order); err != nil {
		return OrdersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataBoxEdgeDevicePager - Lists all the orders related to a Data Box Edge/Data Box Gateway device.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - OrdersClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the OrdersClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *OrdersClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *OrdersClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[OrdersClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrdersClientListByDataBoxEdgeDeviceResponse]{
		More: func(page OrdersClientListByDataBoxEdgeDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrdersClientListByDataBoxEdgeDeviceResponse) (OrdersClientListByDataBoxEdgeDeviceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OrdersClient.NewListByDataBoxEdgeDevicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			}, nil)
			if err != nil {
				return OrdersClientListByDataBoxEdgeDeviceResponse{}, err
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *OrdersClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *OrdersClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (OrdersClientListByDataBoxEdgeDeviceResponse, error) {
	result := OrdersClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderList); err != nil {
		return OrdersClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}

// ListDCAccessCode - Gets the DCAccess Code
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name
//   - resourceGroupName - The resource group name.
//   - options - OrdersClientListDCAccessCodeOptions contains the optional parameters for the OrdersClient.ListDCAccessCode method.
func (client *OrdersClient) ListDCAccessCode(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientListDCAccessCodeOptions) (OrdersClientListDCAccessCodeResponse, error) {
	var err error
	const operationName = "OrdersClient.ListDCAccessCode"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listDCAccessCodeCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return OrdersClientListDCAccessCodeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrdersClientListDCAccessCodeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrdersClientListDCAccessCodeResponse{}, err
	}
	resp, err := client.listDCAccessCodeHandleResponse(httpResp)
	return resp, err
}

// listDCAccessCodeCreateRequest creates the ListDCAccessCode request.
func (client *OrdersClient) listDCAccessCodeCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *OrdersClientListDCAccessCodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/orders/default/listDCAccessCode"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listDCAccessCodeHandleResponse handles the ListDCAccessCode response.
func (client *OrdersClient) listDCAccessCodeHandleResponse(resp *http.Response) (OrdersClientListDCAccessCodeResponse, error) {
	result := OrdersClientListDCAccessCodeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DCAccessCode); err != nil {
		return OrdersClientListDCAccessCodeResponse{}, err
	}
	return result, nil
}
