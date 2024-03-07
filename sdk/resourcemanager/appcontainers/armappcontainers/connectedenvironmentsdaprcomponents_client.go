//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// ConnectedEnvironmentsDaprComponentsClient contains the methods for the ConnectedEnvironmentsDaprComponents group.
// Don't use this type directly, use NewConnectedEnvironmentsDaprComponentsClient() instead.
type ConnectedEnvironmentsDaprComponentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectedEnvironmentsDaprComponentsClient creates a new instance of ConnectedEnvironmentsDaprComponentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedEnvironmentsDaprComponentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedEnvironmentsDaprComponentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectedEnvironmentsDaprComponentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Dapr Component in a connected environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the connected environment.
//   - componentName - Name of the Dapr Component.
//   - daprComponentEnvelope - Configuration details of the Dapr Component.
//   - options - ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateOptions contains the optional parameters for the ConnectedEnvironmentsDaprComponentsClient.CreateOrUpdate
//     method.
func (client *ConnectedEnvironmentsDaprComponentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, daprComponentEnvelope DaprComponent, options *ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateOptions) (ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsDaprComponentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, componentName, daprComponentEnvelope, options)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectedEnvironmentsDaprComponentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, daprComponentEnvelope DaprComponent, options *ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprComponents/{componentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, daprComponentEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectedEnvironmentsDaprComponentsClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse, error) {
	result := ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponent); err != nil {
		return ConnectedEnvironmentsDaprComponentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Dapr Component from a connected environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the connected environment.
//   - componentName - Name of the Dapr Component.
//   - options - ConnectedEnvironmentsDaprComponentsClientDeleteOptions contains the optional parameters for the ConnectedEnvironmentsDaprComponentsClient.Delete
//     method.
func (client *ConnectedEnvironmentsDaprComponentsClient) Delete(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, options *ConnectedEnvironmentsDaprComponentsClientDeleteOptions) (ConnectedEnvironmentsDaprComponentsClientDeleteResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsDaprComponentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, componentName, options)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsDaprComponentsClientDeleteResponse{}, err
	}
	return ConnectedEnvironmentsDaprComponentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedEnvironmentsDaprComponentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, options *ConnectedEnvironmentsDaprComponentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprComponents/{componentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a dapr component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the connected environment.
//   - componentName - Name of the Dapr Component.
//   - options - ConnectedEnvironmentsDaprComponentsClientGetOptions contains the optional parameters for the ConnectedEnvironmentsDaprComponentsClient.Get
//     method.
func (client *ConnectedEnvironmentsDaprComponentsClient) Get(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, options *ConnectedEnvironmentsDaprComponentsClientGetOptions) (ConnectedEnvironmentsDaprComponentsClientGetResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsDaprComponentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, componentName, options)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsDaprComponentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectedEnvironmentsDaprComponentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, options *ConnectedEnvironmentsDaprComponentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprComponents/{componentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedEnvironmentsDaprComponentsClient) getHandleResponse(resp *http.Response) (ConnectedEnvironmentsDaprComponentsClientGetResponse, error) {
	result := ConnectedEnvironmentsDaprComponentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponent); err != nil {
		return ConnectedEnvironmentsDaprComponentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Dapr Components for a connected environment.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the connected environment.
//   - options - ConnectedEnvironmentsDaprComponentsClientListOptions contains the optional parameters for the ConnectedEnvironmentsDaprComponentsClient.NewListPager
//     method.
func (client *ConnectedEnvironmentsDaprComponentsClient) NewListPager(resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsDaprComponentsClientListOptions) *runtime.Pager[ConnectedEnvironmentsDaprComponentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedEnvironmentsDaprComponentsClientListResponse]{
		More: func(page ConnectedEnvironmentsDaprComponentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedEnvironmentsDaprComponentsClientListResponse) (ConnectedEnvironmentsDaprComponentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectedEnvironmentsDaprComponentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, options)
			}, nil)
			if err != nil {
				return ConnectedEnvironmentsDaprComponentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ConnectedEnvironmentsDaprComponentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, options *ConnectedEnvironmentsDaprComponentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprComponents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedEnvironmentsDaprComponentsClient) listHandleResponse(resp *http.Response) (ConnectedEnvironmentsDaprComponentsClientListResponse, error) {
	result := ConnectedEnvironmentsDaprComponentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponentsCollection); err != nil {
		return ConnectedEnvironmentsDaprComponentsClientListResponse{}, err
	}
	return result, nil
}

// ListSecrets - List secrets for a dapr component
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectedEnvironmentName - Name of the connected environment.
//   - componentName - Name of the Dapr Component.
//   - options - ConnectedEnvironmentsDaprComponentsClientListSecretsOptions contains the optional parameters for the ConnectedEnvironmentsDaprComponentsClient.ListSecrets
//     method.
func (client *ConnectedEnvironmentsDaprComponentsClient) ListSecrets(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, options *ConnectedEnvironmentsDaprComponentsClientListSecretsOptions) (ConnectedEnvironmentsDaprComponentsClientListSecretsResponse, error) {
	var err error
	const operationName = "ConnectedEnvironmentsDaprComponentsClient.ListSecrets"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, connectedEnvironmentName, componentName, options)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientListSecretsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectedEnvironmentsDaprComponentsClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectedEnvironmentsDaprComponentsClientListSecretsResponse{}, err
	}
	resp, err := client.listSecretsHandleResponse(httpResp)
	return resp, err
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *ConnectedEnvironmentsDaprComponentsClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, connectedEnvironmentName string, componentName string, options *ConnectedEnvironmentsDaprComponentsClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/connectedEnvironments/{connectedEnvironmentName}/daprComponents/{componentName}/listSecrets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectedEnvironmentName == "" {
		return nil, errors.New("parameter connectedEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedEnvironmentName}", url.PathEscape(connectedEnvironmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *ConnectedEnvironmentsDaprComponentsClient) listSecretsHandleResponse(resp *http.Response) (ConnectedEnvironmentsDaprComponentsClientListSecretsResponse, error) {
	result := ConnectedEnvironmentsDaprComponentsClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSecretsCollection); err != nil {
		return ConnectedEnvironmentsDaprComponentsClientListSecretsResponse{}, err
	}
	return result, nil
}