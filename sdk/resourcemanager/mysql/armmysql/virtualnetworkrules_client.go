//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// VirtualNetworkRulesClient contains the methods for the VirtualNetworkRules group.
// Don't use this type directly, use NewVirtualNetworkRulesClient() instead.
type VirtualNetworkRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworkRulesClient creates a new instance of VirtualNetworkRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworkRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworkRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an existing virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - virtualNetworkRuleName - The name of the virtual network rule.
//   - parameters - The requested virtual Network Rule Resource state.
//   - options - VirtualNetworkRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkRulesClient.BeginCreateOrUpdate
//     method.
func (client *VirtualNetworkRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworkRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, virtualNetworkRuleName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworkRulesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualNetworkRulesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates an existing virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
func (client *VirtualNetworkRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualNetworkRulesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters VirtualNetworkRule, options *VirtualNetworkRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the virtual network rule with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - virtualNetworkRuleName - The name of the virtual network rule.
//   - options - VirtualNetworkRulesClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkRulesClient.BeginDelete
//     method.
func (client *VirtualNetworkRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworkRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworkRulesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualNetworkRulesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the virtual network rule with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
func (client *VirtualNetworkRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualNetworkRulesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
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
func (client *VirtualNetworkRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a virtual network rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - virtualNetworkRuleName - The name of the virtual network rule.
//   - options - VirtualNetworkRulesClientGetOptions contains the optional parameters for the VirtualNetworkRulesClient.Get method.
func (client *VirtualNetworkRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientGetOptions) (VirtualNetworkRulesClientGetResponse, error) {
	var err error
	const operationName = "VirtualNetworkRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, virtualNetworkRuleName, options)
	if err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *VirtualNetworkRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules/{virtualNetworkRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualNetworkRuleName == "" {
		return nil, errors.New("parameter virtualNetworkRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkRuleName}", url.PathEscape(virtualNetworkRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkRulesClient) getHandleResponse(resp *http.Response) (VirtualNetworkRulesClientGetResponse, error) {
	result := VirtualNetworkRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRule); err != nil {
		return VirtualNetworkRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of virtual network rules in a server.
//
// Generated from API version 2017-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - VirtualNetworkRulesClientListByServerOptions contains the optional parameters for the VirtualNetworkRulesClient.NewListByServerPager
//     method.
func (client *VirtualNetworkRulesClient) NewListByServerPager(resourceGroupName string, serverName string, options *VirtualNetworkRulesClientListByServerOptions) *runtime.Pager[VirtualNetworkRulesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkRulesClientListByServerResponse]{
		More: func(page VirtualNetworkRulesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkRulesClientListByServerResponse) (VirtualNetworkRulesClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualNetworkRulesClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return VirtualNetworkRulesClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *VirtualNetworkRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *VirtualNetworkRulesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/virtualNetworkRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *VirtualNetworkRulesClient) listByServerHandleResponse(resp *http.Response) (VirtualNetworkRulesClientListByServerResponse, error) {
	result := VirtualNetworkRulesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkRuleListResult); err != nil {
		return VirtualNetworkRulesClientListByServerResponse{}, err
	}
	return result, nil
}