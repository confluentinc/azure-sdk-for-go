//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// GraphQLAPIResolverPolicyClient contains the methods for the GraphQLAPIResolverPolicy group.
// Don't use this type directly, use NewGraphQLAPIResolverPolicyClient() instead.
type GraphQLAPIResolverPolicyClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGraphQLAPIResolverPolicyClient creates a new instance of GraphQLAPIResolverPolicyClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGraphQLAPIResolverPolicyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GraphQLAPIResolverPolicyClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GraphQLAPIResolverPolicyClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates policy configuration for the GraphQL API Resolver level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - parameters - The policy contents to apply.
//   - options - GraphQLAPIResolverPolicyClientCreateOrUpdateOptions contains the optional parameters for the GraphQLAPIResolverPolicyClient.CreateOrUpdate
//     method.
func (client *GraphQLAPIResolverPolicyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, parameters PolicyContract, options *GraphQLAPIResolverPolicyClientCreateOrUpdateOptions) (GraphQLAPIResolverPolicyClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverPolicyClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, policyID, parameters, options)
	if err != nil {
		return GraphQLAPIResolverPolicyClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverPolicyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverPolicyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GraphQLAPIResolverPolicyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, parameters PolicyContract, options *GraphQLAPIResolverPolicyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GraphQLAPIResolverPolicyClient) createOrUpdateHandleResponse(resp *http.Response) (GraphQLAPIResolverPolicyClientCreateOrUpdateResponse, error) {
	result := GraphQLAPIResolverPolicyClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return GraphQLAPIResolverPolicyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the policy configuration at the GraphQL Api Resolver.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - GraphQLAPIResolverPolicyClientDeleteOptions contains the optional parameters for the GraphQLAPIResolverPolicyClient.Delete
//     method.
func (client *GraphQLAPIResolverPolicyClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, ifMatch string, options *GraphQLAPIResolverPolicyClientDeleteOptions) (GraphQLAPIResolverPolicyClientDeleteResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverPolicyClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, policyID, ifMatch, options)
	if err != nil {
		return GraphQLAPIResolverPolicyClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverPolicyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverPolicyClientDeleteResponse{}, err
	}
	return GraphQLAPIResolverPolicyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GraphQLAPIResolverPolicyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, ifMatch string, options *GraphQLAPIResolverPolicyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the policy configuration at the GraphQL API Resolver level.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - options - GraphQLAPIResolverPolicyClientGetOptions contains the optional parameters for the GraphQLAPIResolverPolicyClient.Get
//     method.
func (client *GraphQLAPIResolverPolicyClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, options *GraphQLAPIResolverPolicyClientGetOptions) (GraphQLAPIResolverPolicyClientGetResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverPolicyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, policyID, options)
	if err != nil {
		return GraphQLAPIResolverPolicyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverPolicyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverPolicyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GraphQLAPIResolverPolicyClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, options *GraphQLAPIResolverPolicyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Format != nil {
		reqQP.Set("format", string(*options.Format))
	}
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GraphQLAPIResolverPolicyClient) getHandleResponse(resp *http.Response) (GraphQLAPIResolverPolicyClientGetResponse, error) {
	result := GraphQLAPIResolverPolicyClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyContract); err != nil {
		return GraphQLAPIResolverPolicyClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the GraphQL API resolver policy specified by its identifier.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - policyID - The identifier of the Policy.
//   - options - GraphQLAPIResolverPolicyClientGetEntityTagOptions contains the optional parameters for the GraphQLAPIResolverPolicyClient.GetEntityTag
//     method.
func (client *GraphQLAPIResolverPolicyClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, options *GraphQLAPIResolverPolicyClientGetEntityTagOptions) (GraphQLAPIResolverPolicyClientGetEntityTagResponse, error) {
	var err error
	const operationName = "GraphQLAPIResolverPolicyClient.GetEntityTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, policyID, options)
	if err != nil {
		return GraphQLAPIResolverPolicyClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GraphQLAPIResolverPolicyClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GraphQLAPIResolverPolicyClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *GraphQLAPIResolverPolicyClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, policyID PolicyIDName, options *GraphQLAPIResolverPolicyClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}/policies/{policyId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if policyID == "" {
		return nil, errors.New("parameter policyID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyId}", url.PathEscape(string(policyID)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *GraphQLAPIResolverPolicyClient) getEntityTagHandleResponse(resp *http.Response) (GraphQLAPIResolverPolicyClientGetEntityTagResponse, error) {
	result := GraphQLAPIResolverPolicyClientGetEntityTagResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// NewListByResolverPager - Get the list of policy configuration at the GraphQL API Resolver level.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API revision identifier. Must be unique in the current API Management service instance. Non-current revision has
//     ;rev=n as a suffix where n is the revision number.
//   - resolverID - Resolver identifier within a GraphQL API. Must be unique in the current API Management service instance.
//   - options - GraphQLAPIResolverPolicyClientListByResolverOptions contains the optional parameters for the GraphQLAPIResolverPolicyClient.NewListByResolverPager
//     method.
func (client *GraphQLAPIResolverPolicyClient) NewListByResolverPager(resourceGroupName string, serviceName string, apiID string, resolverID string, options *GraphQLAPIResolverPolicyClientListByResolverOptions) *runtime.Pager[GraphQLAPIResolverPolicyClientListByResolverResponse] {
	return runtime.NewPager(runtime.PagingHandler[GraphQLAPIResolverPolicyClientListByResolverResponse]{
		More: func(page GraphQLAPIResolverPolicyClientListByResolverResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GraphQLAPIResolverPolicyClientListByResolverResponse) (GraphQLAPIResolverPolicyClientListByResolverResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GraphQLAPIResolverPolicyClient.NewListByResolverPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResolverCreateRequest(ctx, resourceGroupName, serviceName, apiID, resolverID, options)
			}, nil)
			if err != nil {
				return GraphQLAPIResolverPolicyClientListByResolverResponse{}, err
			}
			return client.listByResolverHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResolverCreateRequest creates the ListByResolver request.
func (client *GraphQLAPIResolverPolicyClient) listByResolverCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, resolverID string, options *GraphQLAPIResolverPolicyClientListByResolverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/resolvers/{resolverId}/policies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if resolverID == "" {
		return nil, errors.New("parameter resolverID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resolverId}", url.PathEscape(resolverID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResolverHandleResponse handles the ListByResolver response.
func (client *GraphQLAPIResolverPolicyClient) listByResolverHandleResponse(resp *http.Response) (GraphQLAPIResolverPolicyClientListByResolverResponse, error) {
	result := GraphQLAPIResolverPolicyClientListByResolverResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyCollection); err != nil {
		return GraphQLAPIResolverPolicyClientListByResolverResponse{}, err
	}
	return result, nil
}