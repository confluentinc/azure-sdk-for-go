//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcommerce

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
	"strconv"
	"strings"
	"time"
)

// UsageAggregatesClient contains the methods for the UsageAggregates group.
// Don't use this type directly, use NewUsageAggregatesClient() instead.
type UsageAggregatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUsageAggregatesClient creates a new instance of UsageAggregatesClient with the specified values.
//   - subscriptionID - It uniquely identifies Microsoft Azure subscription. The subscription ID forms part of the URI for every
//     service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUsageAggregatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsageAggregatesClient, error) {
	cl, err := arm.NewClient(internal.ModuleName+"/armcommerce.UsageAggregatesClient", internal.ModuleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UsageAggregatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Query aggregated Azure subscription consumption data for a date range.
//
// Generated from API version 2015-06-01-preview
//   - reportedStartTime - The start of the time range to retrieve data for.
//   - reportedEndTime - The end of the time range to retrieve data for.
//   - options - UsageAggregatesClientListOptions contains the optional parameters for the UsageAggregatesClient.NewListPager
//     method.
func (client *UsageAggregatesClient) NewListPager(reportedStartTime time.Time, reportedEndTime time.Time, options *UsageAggregatesClientListOptions) *runtime.Pager[UsageAggregatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[UsageAggregatesClientListResponse]{
		More: func(page UsageAggregatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UsageAggregatesClientListResponse) (UsageAggregatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, reportedStartTime, reportedEndTime, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return UsageAggregatesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return UsageAggregatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return UsageAggregatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *UsageAggregatesClient) listCreateRequest(ctx context.Context, reportedStartTime time.Time, reportedEndTime time.Time, options *UsageAggregatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Commerce/UsageAggregates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("reportedStartTime", reportedStartTime.Format(time.RFC3339Nano))
	reqQP.Set("reportedEndTime", reportedEndTime.Format(time.RFC3339Nano))
	if options != nil && options.ShowDetails != nil {
		reqQP.Set("showDetails", strconv.FormatBool(*options.ShowDetails))
	}
	if options != nil && options.AggregationGranularity != nil {
		reqQP.Set("aggregationGranularity", string(*options.AggregationGranularity))
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	reqQP.Set("api-version", "2015-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsageAggregatesClient) listHandleResponse(resp *http.Response) (UsageAggregatesClientListResponse, error) {
	result := UsageAggregatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageAggregationListResult); err != nil {
		return UsageAggregatesClientListResponse{}, err
	}
	return result, nil
}
