//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfluidrelay

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

// ContainersClient contains the methods for the FluidRelayContainers group.
// Don't use this type directly, use NewContainersClient() instead.
type ContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainersClient creates a new instance of ContainersClient with the specified values.
//   - subscriptionID - The subscription id (GUID) for this resource.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Delete a Fluid Relay container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroup - The resource group containing the resource.
//   - fluidRelayServerName - The Fluid Relay server resource name.
//   - fluidRelayContainerName - The Fluid Relay container resource name.
//   - options - ContainersClientDeleteOptions contains the optional parameters for the ContainersClient.Delete method.
func (client *ContainersClient) Delete(ctx context.Context, resourceGroup string, fluidRelayServerName string, fluidRelayContainerName string, options *ContainersClientDeleteOptions) (ContainersClientDeleteResponse, error) {
	var err error
	const operationName = "ContainersClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroup, fluidRelayServerName, fluidRelayContainerName, options)
	if err != nil {
		return ContainersClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ContainersClientDeleteResponse{}, err
	}
	return ContainersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContainersClient) deleteCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, fluidRelayContainerName string, options *ContainersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}/fluidRelayContainers/{fluidRelayContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	if fluidRelayContainerName == "" {
		return nil, errors.New("parameter fluidRelayContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayContainerName}", url.PathEscape(fluidRelayContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Fluid Relay container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-06-01
//   - resourceGroup - The resource group containing the resource.
//   - fluidRelayServerName - The Fluid Relay server resource name.
//   - fluidRelayContainerName - The Fluid Relay container resource name.
//   - options - ContainersClientGetOptions contains the optional parameters for the ContainersClient.Get method.
func (client *ContainersClient) Get(ctx context.Context, resourceGroup string, fluidRelayServerName string, fluidRelayContainerName string, options *ContainersClientGetOptions) (ContainersClientGetResponse, error) {
	var err error
	const operationName = "ContainersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroup, fluidRelayServerName, fluidRelayContainerName, options)
	if err != nil {
		return ContainersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContainersClient) getCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, fluidRelayContainerName string, options *ContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}/fluidRelayContainers/{fluidRelayContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	if fluidRelayContainerName == "" {
		return nil, errors.New("parameter fluidRelayContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayContainerName}", url.PathEscape(fluidRelayContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainersClient) getHandleResponse(resp *http.Response) (ContainersClientGetResponse, error) {
	result := ContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Container); err != nil {
		return ContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFluidRelayServersPager - List all Fluid Relay containers which are children of a given Fluid Relay server.
//
// Generated from API version 2022-06-01
//   - resourceGroup - The resource group containing the resource.
//   - fluidRelayServerName - The Fluid Relay server resource name.
//   - options - ContainersClientListByFluidRelayServersOptions contains the optional parameters for the ContainersClient.NewListByFluidRelayServersPager
//     method.
func (client *ContainersClient) NewListByFluidRelayServersPager(resourceGroup string, fluidRelayServerName string, options *ContainersClientListByFluidRelayServersOptions) *runtime.Pager[ContainersClientListByFluidRelayServersResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainersClientListByFluidRelayServersResponse]{
		More: func(page ContainersClientListByFluidRelayServersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainersClientListByFluidRelayServersResponse) (ContainersClientListByFluidRelayServersResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainersClient.NewListByFluidRelayServersPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFluidRelayServersCreateRequest(ctx, resourceGroup, fluidRelayServerName, options)
			}, nil)
			if err != nil {
				return ContainersClientListByFluidRelayServersResponse{}, err
			}
			return client.listByFluidRelayServersHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFluidRelayServersCreateRequest creates the ListByFluidRelayServers request.
func (client *ContainersClient) listByFluidRelayServersCreateRequest(ctx context.Context, resourceGroup string, fluidRelayServerName string, options *ContainersClientListByFluidRelayServersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.FluidRelay/fluidRelayServers/{fluidRelayServerName}/fluidRelayContainers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if fluidRelayServerName == "" {
		return nil, errors.New("parameter fluidRelayServerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fluidRelayServerName}", url.PathEscape(fluidRelayServerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFluidRelayServersHandleResponse handles the ListByFluidRelayServers response.
func (client *ContainersClient) listByFluidRelayServersHandleResponse(resp *http.Response) (ContainersClientListByFluidRelayServersResponse, error) {
	result := ContainersClientListByFluidRelayServersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerList); err != nil {
		return ContainersClientListByFluidRelayServersResponse{}, err
	}
	return result, nil
}