//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armorbital

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

// AvailableGroundStationsClient contains the methods for the AvailableGroundStations group.
// Don't use this type directly, use NewAvailableGroundStationsClient() instead.
type AvailableGroundStationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAvailableGroundStationsClient creates a new instance of AvailableGroundStationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailableGroundStationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableGroundStationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableGroundStationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByCapabilityPager - Returns list of available ground stations.
//
// Generated from API version 2022-11-01
//   - capability - Ground Station Capability.
//   - options - AvailableGroundStationsClientListByCapabilityOptions contains the optional parameters for the AvailableGroundStationsClient.NewListByCapabilityPager
//     method.
func (client *AvailableGroundStationsClient) NewListByCapabilityPager(capability CapabilityParameter, options *AvailableGroundStationsClientListByCapabilityOptions) *runtime.Pager[AvailableGroundStationsClientListByCapabilityResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailableGroundStationsClientListByCapabilityResponse]{
		More: func(page AvailableGroundStationsClientListByCapabilityResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableGroundStationsClientListByCapabilityResponse) (AvailableGroundStationsClientListByCapabilityResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AvailableGroundStationsClient.NewListByCapabilityPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByCapabilityCreateRequest(ctx, capability, options)
			}, nil)
			if err != nil {
				return AvailableGroundStationsClientListByCapabilityResponse{}, err
			}
			return client.listByCapabilityHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByCapabilityCreateRequest creates the ListByCapability request.
func (client *AvailableGroundStationsClient) listByCapabilityCreateRequest(ctx context.Context, capability CapabilityParameter, options *AvailableGroundStationsClientListByCapabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Orbital/availableGroundStations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	reqQP.Set("capability", string(capability))
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCapabilityHandleResponse handles the ListByCapability response.
func (client *AvailableGroundStationsClient) listByCapabilityHandleResponse(resp *http.Response) (AvailableGroundStationsClientListByCapabilityResponse, error) {
	result := AvailableGroundStationsClientListByCapabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableGroundStationListResult); err != nil {
		return AvailableGroundStationsClientListByCapabilityResponse{}, err
	}
	return result, nil
}