//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get properties of a private link resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentType - The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace name or
//     namespace name).
//   - privateLinkResourceName - The name of private link resource will be either topic, domain, partnerNamespace or namespace.
//   - options - PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get
//     method.
func (client *PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateLinkResourceName string, options *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error) {
	var err error
	const operationName = "PrivateLinkResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, parentType, parentName, privateLinkResourceName, options)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, parentType string, parentName string, privateLinkResourceName string, options *PrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateLinkResources/{privateLinkResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentType == "" {
		return nil, errors.New("parameter parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(parentType))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	if privateLinkResourceName == "" {
		return nil, errors.New("parameter privateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkResourceName}", url.PathEscape(privateLinkResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (PrivateLinkResourcesClientGetResponse, error) {
	result := PrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourcePager - List all the private link resources under a topic, domain, or partner namespace or namespace.
//
// Generated from API version 2023-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - parentType - The type of the parent resource. This can be either \'topics\', \'domains\', or \'partnerNamespaces\' or \'namespaces\'.
//   - parentName - The name of the parent resource (namely, either, the topic name, domain name, or partner namespace or namespace
//     name).
//   - options - PrivateLinkResourcesClientListByResourceOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByResourcePager
//     method.
func (client *PrivateLinkResourcesClient) NewListByResourcePager(resourceGroupName string, parentType string, parentName string, options *PrivateLinkResourcesClientListByResourceOptions) *runtime.Pager[PrivateLinkResourcesClientListByResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientListByResourceResponse]{
		More: func(page PrivateLinkResourcesClientListByResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientListByResourceResponse) (PrivateLinkResourcesClientListByResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinkResourcesClient.NewListByResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceCreateRequest(ctx, resourceGroupName, parentType, parentName, options)
			}, nil)
			if err != nil {
				return PrivateLinkResourcesClientListByResourceResponse{}, err
			}
			return client.listByResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *PrivateLinkResourcesClient) listByResourceCreateRequest(ctx context.Context, resourceGroupName string, parentType string, parentName string, options *PrivateLinkResourcesClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/{parentType}/{parentName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if parentType == "" {
		return nil, errors.New("parameter parentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentType}", url.PathEscape(parentType))
	if parentName == "" {
		return nil, errors.New("parameter parentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{parentName}", url.PathEscape(parentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-15-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *PrivateLinkResourcesClient) listByResourceHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByResourceResponse, error) {
	result := PrivateLinkResourcesClientListByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourcesListResult); err != nil {
		return PrivateLinkResourcesClientListByResourceResponse{}, err
	}
	return result, nil
}