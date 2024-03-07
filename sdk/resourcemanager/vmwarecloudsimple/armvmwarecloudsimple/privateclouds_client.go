//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// PrivateCloudsClient contains the methods for the PrivateClouds group.
// Don't use this type directly, use NewPrivateCloudsClient() instead.
type PrivateCloudsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateCloudsClient creates a new instance of PrivateCloudsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateCloudsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateCloudsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateCloudsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns private cloud by its name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-04-01
//   - pcName - The private cloud name
//   - regionID - The region Id (westus, eastus)
//   - options - PrivateCloudsClientGetOptions contains the optional parameters for the PrivateCloudsClient.Get method.
func (client *PrivateCloudsClient) Get(ctx context.Context, pcName string, regionID string, options *PrivateCloudsClientGetOptions) (PrivateCloudsClientGetResponse, error) {
	var err error
	const operationName = "PrivateCloudsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, pcName, regionID, options)
	if err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateCloudsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateCloudsClient) getCreateRequest(ctx context.Context, pcName string, regionID string, options *PrivateCloudsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateCloudsClient) getHandleResponse(resp *http.Response) (PrivateCloudsClientGetResponse, error) {
	result := PrivateCloudsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloud); err != nil {
		return PrivateCloudsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of private clouds in particular region
//
// Generated from API version 2019-04-01
//   - regionID - The region Id (westus, eastus)
//   - options - PrivateCloudsClientListOptions contains the optional parameters for the PrivateCloudsClient.NewListPager method.
func (client *PrivateCloudsClient) NewListPager(regionID string, options *PrivateCloudsClientListOptions) *runtime.Pager[PrivateCloudsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateCloudsClientListResponse]{
		More: func(page PrivateCloudsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateCloudsClientListResponse) (PrivateCloudsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateCloudsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, regionID, options)
			}, nil)
			if err != nil {
				return PrivateCloudsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PrivateCloudsClient) listCreateRequest(ctx context.Context, regionID string, options *PrivateCloudsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateCloudsClient) listHandleResponse(resp *http.Response) (PrivateCloudsClientListResponse, error) {
	result := PrivateCloudsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateCloudList); err != nil {
		return PrivateCloudsClientListResponse{}, err
	}
	return result, nil
}