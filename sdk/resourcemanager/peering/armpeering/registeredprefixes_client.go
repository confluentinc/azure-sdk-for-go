//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// RegisteredPrefixesClient contains the methods for the RegisteredPrefixes group.
// Don't use this type directly, use NewRegisteredPrefixesClient() instead.
type RegisteredPrefixesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRegisteredPrefixesClient creates a new instance of RegisteredPrefixesClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegisteredPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RegisteredPrefixesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegisteredPrefixesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new registered prefix with the specified name under the given subscription, resource group and
// peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - registeredPrefixName - The name of the registered prefix.
//   - registeredPrefix - The properties needed to create a registered prefix.
//   - options - RegisteredPrefixesClientCreateOrUpdateOptions contains the optional parameters for the RegisteredPrefixesClient.CreateOrUpdate
//     method.
func (client *RegisteredPrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, registeredPrefix RegisteredPrefix, options *RegisteredPrefixesClientCreateOrUpdateOptions) (RegisteredPrefixesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "RegisteredPrefixesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringName, registeredPrefixName, registeredPrefix, options)
	if err != nil {
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RegisteredPrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, registeredPrefix RegisteredPrefix, options *RegisteredPrefixesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredPrefixName == "" {
		return nil, errors.New("parameter registeredPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredPrefixName}", url.PathEscape(registeredPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, registeredPrefix); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *RegisteredPrefixesClient) createOrUpdateHandleResponse(resp *http.Response) (RegisteredPrefixesClientCreateOrUpdateResponse, error) {
	result := RegisteredPrefixesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredPrefix); err != nil {
		return RegisteredPrefixesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing registered prefix with the specified name under the given subscription, resource group and
// peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - registeredPrefixName - The name of the registered prefix.
//   - options - RegisteredPrefixesClientDeleteOptions contains the optional parameters for the RegisteredPrefixesClient.Delete
//     method.
func (client *RegisteredPrefixesClient) Delete(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientDeleteOptions) (RegisteredPrefixesClientDeleteResponse, error) {
	var err error
	const operationName = "RegisteredPrefixesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringName, registeredPrefixName, options)
	if err != nil {
		return RegisteredPrefixesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegisteredPrefixesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RegisteredPrefixesClientDeleteResponse{}, err
	}
	return RegisteredPrefixesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegisteredPrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredPrefixName == "" {
		return nil, errors.New("parameter registeredPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredPrefixName}", url.PathEscape(registeredPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing registered prefix with the specified name under the given subscription, resource group and peering.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - registeredPrefixName - The name of the registered prefix.
//   - options - RegisteredPrefixesClientGetOptions contains the optional parameters for the RegisteredPrefixesClient.Get method.
func (client *RegisteredPrefixesClient) Get(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientGetOptions) (RegisteredPrefixesClientGetResponse, error) {
	var err error
	const operationName = "RegisteredPrefixesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringName, registeredPrefixName, options)
	if err != nil {
		return RegisteredPrefixesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegisteredPrefixesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RegisteredPrefixesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RegisteredPrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, registeredPrefixName string, options *RegisteredPrefixesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes/{registeredPrefixName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if registeredPrefixName == "" {
		return nil, errors.New("parameter registeredPrefixName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registeredPrefixName}", url.PathEscape(registeredPrefixName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegisteredPrefixesClient) getHandleResponse(resp *http.Response) (RegisteredPrefixesClientGetResponse, error) {
	result := RegisteredPrefixesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredPrefix); err != nil {
		return RegisteredPrefixesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPeeringPager - Lists all registered prefixes under the given subscription, resource group and peering.
//
// Generated from API version 2022-01-01
//   - resourceGroupName - The name of the resource group.
//   - peeringName - The name of the peering.
//   - options - RegisteredPrefixesClientListByPeeringOptions contains the optional parameters for the RegisteredPrefixesClient.NewListByPeeringPager
//     method.
func (client *RegisteredPrefixesClient) NewListByPeeringPager(resourceGroupName string, peeringName string, options *RegisteredPrefixesClientListByPeeringOptions) *runtime.Pager[RegisteredPrefixesClientListByPeeringResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegisteredPrefixesClientListByPeeringResponse]{
		More: func(page RegisteredPrefixesClientListByPeeringResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegisteredPrefixesClientListByPeeringResponse) (RegisteredPrefixesClientListByPeeringResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RegisteredPrefixesClient.NewListByPeeringPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByPeeringCreateRequest(ctx, resourceGroupName, peeringName, options)
			}, nil)
			if err != nil {
				return RegisteredPrefixesClientListByPeeringResponse{}, err
			}
			return client.listByPeeringHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByPeeringCreateRequest creates the ListByPeering request.
func (client *RegisteredPrefixesClient) listByPeeringCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, options *RegisteredPrefixesClientListByPeeringOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/registeredPrefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPeeringHandleResponse handles the ListByPeering response.
func (client *RegisteredPrefixesClient) listByPeeringHandleResponse(resp *http.Response) (RegisteredPrefixesClientListByPeeringResponse, error) {
	result := RegisteredPrefixesClientListByPeeringResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegisteredPrefixListResult); err != nil {
		return RegisteredPrefixesClientListByPeeringResponse{}, err
	}
	return result, nil
}