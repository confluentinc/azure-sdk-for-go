//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

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

// MarketplaceAgreementsClient contains the methods for the MarketplaceAgreements group.
// Don't use this type directly, use NewMarketplaceAgreementsClient() instead.
type MarketplaceAgreementsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMarketplaceAgreementsClient creates a new instance of MarketplaceAgreementsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMarketplaceAgreementsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MarketplaceAgreementsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MarketplaceAgreementsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create Confluent Marketplace agreement in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-22
//   - options - MarketplaceAgreementsClientCreateOptions contains the optional parameters for the MarketplaceAgreementsClient.Create
//     method.
func (client *MarketplaceAgreementsClient) Create(ctx context.Context, options *MarketplaceAgreementsClientCreateOptions) (MarketplaceAgreementsClientCreateResponse, error) {
	var err error
	const operationName = "MarketplaceAgreementsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, options)
	if err != nil {
		return MarketplaceAgreementsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceAgreementsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MarketplaceAgreementsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *MarketplaceAgreementsClient) createCreateRequest(ctx context.Context, options *MarketplaceAgreementsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/agreements/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *MarketplaceAgreementsClient) createHandleResponse(resp *http.Response) (MarketplaceAgreementsClientCreateResponse, error) {
	result := MarketplaceAgreementsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementResource); err != nil {
		return MarketplaceAgreementsClientCreateResponse{}, err
	}
	return result, nil
}

// NewListPager - List Confluent marketplace agreements in the subscription.
//
// Generated from API version 2023-08-22
//   - options - MarketplaceAgreementsClientListOptions contains the optional parameters for the MarketplaceAgreementsClient.NewListPager
//     method.
func (client *MarketplaceAgreementsClient) NewListPager(options *MarketplaceAgreementsClientListOptions) *runtime.Pager[MarketplaceAgreementsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MarketplaceAgreementsClientListResponse]{
		More: func(page MarketplaceAgreementsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MarketplaceAgreementsClientListResponse) (MarketplaceAgreementsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MarketplaceAgreementsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return MarketplaceAgreementsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MarketplaceAgreementsClient) listCreateRequest(ctx context.Context, options *MarketplaceAgreementsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Confluent/agreements"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-22")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplaceAgreementsClient) listHandleResponse(resp *http.Response) (MarketplaceAgreementsClientListResponse, error) {
	result := MarketplaceAgreementsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementResourceListResponse); err != nil {
		return MarketplaceAgreementsClientListResponse{}, err
	}
	return result, nil
}