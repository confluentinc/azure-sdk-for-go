//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// IPv6FirewallRulesClient contains the methods for the IPv6FirewallRules group.
// Don't use this type directly, use NewIPv6FirewallRulesClient() instead.
type IPv6FirewallRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIPv6FirewallRulesClient creates a new instance of IPv6FirewallRulesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIPv6FirewallRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IPv6FirewallRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IPv6FirewallRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an IPv6 firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the firewall rule.
//   - parameters - The required parameters for creating or updating an IPv6 firewall rule.
//   - options - IPv6FirewallRulesClientCreateOrUpdateOptions contains the optional parameters for the IPv6FirewallRulesClient.CreateOrUpdate
//     method.
func (client *IPv6FirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters IPv6FirewallRule, options *IPv6FirewallRulesClientCreateOrUpdateOptions) (IPv6FirewallRulesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "IPv6FirewallRulesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, parameters, options)
	if err != nil {
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IPv6FirewallRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters IPv6FirewallRule, options *IPv6FirewallRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *IPv6FirewallRulesClient) createOrUpdateHandleResponse(resp *http.Response) (IPv6FirewallRulesClientCreateOrUpdateResponse, error) {
	result := IPv6FirewallRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPv6FirewallRule); err != nil {
		return IPv6FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an IPv6 firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the firewall rule.
//   - options - IPv6FirewallRulesClientDeleteOptions contains the optional parameters for the IPv6FirewallRulesClient.Delete
//     method.
func (client *IPv6FirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientDeleteOptions) (IPv6FirewallRulesClientDeleteResponse, error) {
	var err error
	const operationName = "IPv6FirewallRulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, options)
	if err != nil {
		return IPv6FirewallRulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IPv6FirewallRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IPv6FirewallRulesClientDeleteResponse{}, err
	}
	return IPv6FirewallRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IPv6FirewallRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets an IPv6 firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the firewall rule.
//   - options - IPv6FirewallRulesClientGetOptions contains the optional parameters for the IPv6FirewallRulesClient.Get method.
func (client *IPv6FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientGetOptions) (IPv6FirewallRulesClientGetResponse, error) {
	var err error
	const operationName = "IPv6FirewallRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, options)
	if err != nil {
		return IPv6FirewallRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IPv6FirewallRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IPv6FirewallRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IPv6FirewallRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *IPv6FirewallRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules/{firewallRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IPv6FirewallRulesClient) getHandleResponse(resp *http.Response) (IPv6FirewallRulesClientGetResponse, error) {
	result := IPv6FirewallRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPv6FirewallRule); err != nil {
		return IPv6FirewallRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of IPv6 firewall rules.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - IPv6FirewallRulesClientListByServerOptions contains the optional parameters for the IPv6FirewallRulesClient.NewListByServerPager
//     method.
func (client *IPv6FirewallRulesClient) NewListByServerPager(resourceGroupName string, serverName string, options *IPv6FirewallRulesClientListByServerOptions) *runtime.Pager[IPv6FirewallRulesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[IPv6FirewallRulesClientListByServerResponse]{
		More: func(page IPv6FirewallRulesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IPv6FirewallRulesClientListByServerResponse) (IPv6FirewallRulesClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IPv6FirewallRulesClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return IPv6FirewallRulesClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *IPv6FirewallRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *IPv6FirewallRulesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/ipv6FirewallRules"
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
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *IPv6FirewallRulesClient) listByServerHandleResponse(resp *http.Response) (IPv6FirewallRulesClientListByServerResponse, error) {
	result := IPv6FirewallRulesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IPv6FirewallRuleListResult); err != nil {
		return IPv6FirewallRulesClientListByServerResponse{}, err
	}
	return result, nil
}