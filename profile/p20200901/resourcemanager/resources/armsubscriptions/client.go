//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsubscriptions

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the Subscriptions group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal *arm.Client
}

// NewClient creates a new instance of Client with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armsubscriptions.Client", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		internal: cl,
	}
	return client, nil
}

// CheckZonePeers - Compares a subscriptions logical zone mapping
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-06-01
//   - subscriptionID - The ID of the target subscription.
//   - parameters - Parameters for checking zone peers.
//   - options - ClientCheckZonePeersOptions contains the optional parameters for the Client.CheckZonePeers method.
func (client *Client) CheckZonePeers(ctx context.Context, subscriptionID string, parameters CheckZonePeersRequest, options *ClientCheckZonePeersOptions) (ClientCheckZonePeersResponse, error) {
	req, err := client.checkZonePeersCreateRequest(ctx, subscriptionID, parameters, options)
	if err != nil {
		return ClientCheckZonePeersResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCheckZonePeersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientCheckZonePeersResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkZonePeersHandleResponse(resp)
}

// checkZonePeersCreateRequest creates the CheckZonePeers request.
func (client *Client) checkZonePeersCreateRequest(ctx context.Context, subscriptionID string, parameters CheckZonePeersRequest, options *ClientCheckZonePeersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Resources/checkZonePeers/"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// checkZonePeersHandleResponse handles the CheckZonePeers response.
func (client *Client) checkZonePeersHandleResponse(resp *http.Response) (ClientCheckZonePeersResponse, error) {
	result := ClientCheckZonePeersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckZonePeersResult); err != nil {
		return ClientCheckZonePeersResponse{}, err
	}
	return result, nil
}

// Get - Gets details about a specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2016-06-01
//   - subscriptionID - The ID of the target subscription.
//   - options - ClientGetOptions contains the optional parameters for the Client.Get method.
func (client *Client) Get(ctx context.Context, subscriptionID string, options *ClientGetOptions) (ClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return ClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *Client) getCreateRequest(ctx context.Context, subscriptionID string, options *ClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *Client) getHandleResponse(resp *http.Response) (ClientGetResponse, error) {
	result := ClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Subscription); err != nil {
		return ClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all subscriptions for a tenant.
//
// Generated from API version 2016-06-01
//   - options - ClientListOptions contains the optional parameters for the Client.NewListPager method.
func (client *Client) NewListPager(options *ClientListOptions) *runtime.Pager[ClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListResponse]{
		More: func(page ClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListResponse) (ClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *Client) listCreateRequest(ctx context.Context, options *ClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *Client) listHandleResponse(resp *http.Response) (ClientListResponse, error) {
	result := ClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionListResult); err != nil {
		return ClientListResponse{}, err
	}
	return result, nil
}

// NewListLocationsPager - This operation provides all the locations that are available for resource providers; however, each
// resource provider may support a subset of this list.
//
// Generated from API version 2016-06-01
//   - subscriptionID - The ID of the target subscription.
//   - options - ClientListLocationsOptions contains the optional parameters for the Client.NewListLocationsPager method.
func (client *Client) NewListLocationsPager(subscriptionID string, options *ClientListLocationsOptions) *runtime.Pager[ClientListLocationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListLocationsResponse]{
		More: func(page ClientListLocationsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ClientListLocationsResponse) (ClientListLocationsResponse, error) {
			req, err := client.listLocationsCreateRequest(ctx, subscriptionID, options)
			if err != nil {
				return ClientListLocationsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientListLocationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListLocationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listLocationsHandleResponse(resp)
		},
	})
}

// listLocationsCreateRequest creates the ListLocations request.
func (client *Client) listLocationsCreateRequest(ctx context.Context, subscriptionID string, options *ClientListLocationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/locations"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listLocationsHandleResponse handles the ListLocations response.
func (client *Client) listLocationsHandleResponse(resp *http.Response) (ClientListLocationsResponse, error) {
	result := ClientListLocationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationListResult); err != nil {
		return ClientListLocationsResponse{}, err
	}
	return result, nil
}
