//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch

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

// LocationClient contains the methods for the Location group.
// Don't use this type directly, use NewLocationClient() instead.
type LocationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLocationClient creates a new instance of LocationClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LocationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Checks whether the Batch account name is available in the specified region.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - locationName - The desired region for the name check.
//   - parameters - Properties needed to check the availability of a name.
//   - options - LocationClientCheckNameAvailabilityOptions contains the optional parameters for the LocationClient.CheckNameAvailability
//     method.
func (client *LocationClient) CheckNameAvailability(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters, options *LocationClientCheckNameAvailabilityOptions) (LocationClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "LocationClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return LocationClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *LocationClient) checkNameAvailabilityCreateRequest(ctx context.Context, locationName string, parameters CheckNameAvailabilityParameters, options *LocationClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/checkNameAvailability"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *LocationClient) checkNameAvailabilityHandleResponse(resp *http.Response) (LocationClientCheckNameAvailabilityResponse, error) {
	result := LocationClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return LocationClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// GetQuotas - Gets the Batch service quotas for the specified subscription at the given location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - locationName - The region for which to retrieve Batch service quotas.
//   - options - LocationClientGetQuotasOptions contains the optional parameters for the LocationClient.GetQuotas method.
func (client *LocationClient) GetQuotas(ctx context.Context, locationName string, options *LocationClientGetQuotasOptions) (LocationClientGetQuotasResponse, error) {
	var err error
	const operationName = "LocationClient.GetQuotas"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getQuotasCreateRequest(ctx, locationName, options)
	if err != nil {
		return LocationClientGetQuotasResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LocationClientGetQuotasResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LocationClientGetQuotasResponse{}, err
	}
	resp, err := client.getQuotasHandleResponse(httpResp)
	return resp, err
}

// getQuotasCreateRequest creates the GetQuotas request.
func (client *LocationClient) getQuotasCreateRequest(ctx context.Context, locationName string, options *LocationClientGetQuotasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/quotas"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getQuotasHandleResponse handles the GetQuotas response.
func (client *LocationClient) getQuotasHandleResponse(resp *http.Response) (LocationClientGetQuotasResponse, error) {
	result := LocationClientGetQuotasResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationQuota); err != nil {
		return LocationClientGetQuotasResponse{}, err
	}
	return result, nil
}

// NewListSupportedCloudServiceSKUsPager - Gets the list of Batch supported Cloud Service VM sizes available at the given
// location.
//
// Generated from API version 2023-11-01
//   - locationName - The region for which to retrieve Batch service supported SKUs.
//   - options - LocationClientListSupportedCloudServiceSKUsOptions contains the optional parameters for the LocationClient.NewListSupportedCloudServiceSKUsPager
//     method.
func (client *LocationClient) NewListSupportedCloudServiceSKUsPager(locationName string, options *LocationClientListSupportedCloudServiceSKUsOptions) *runtime.Pager[LocationClientListSupportedCloudServiceSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocationClientListSupportedCloudServiceSKUsResponse]{
		More: func(page LocationClientListSupportedCloudServiceSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationClientListSupportedCloudServiceSKUsResponse) (LocationClientListSupportedCloudServiceSKUsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LocationClient.NewListSupportedCloudServiceSKUsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSupportedCloudServiceSKUsCreateRequest(ctx, locationName, options)
			}, nil)
			if err != nil {
				return LocationClientListSupportedCloudServiceSKUsResponse{}, err
			}
			return client.listSupportedCloudServiceSKUsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSupportedCloudServiceSKUsCreateRequest creates the ListSupportedCloudServiceSKUs request.
func (client *LocationClient) listSupportedCloudServiceSKUsCreateRequest(ctx context.Context, locationName string, options *LocationClientListSupportedCloudServiceSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/cloudServiceSkus"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSupportedCloudServiceSKUsHandleResponse handles the ListSupportedCloudServiceSKUs response.
func (client *LocationClient) listSupportedCloudServiceSKUsHandleResponse(resp *http.Response) (LocationClientListSupportedCloudServiceSKUsResponse, error) {
	result := LocationClientListSupportedCloudServiceSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SupportedSKUsResult); err != nil {
		return LocationClientListSupportedCloudServiceSKUsResponse{}, err
	}
	return result, nil
}

// NewListSupportedVirtualMachineSKUsPager - Gets the list of Batch supported Virtual Machine VM sizes available at the given
// location.
//
// Generated from API version 2023-11-01
//   - locationName - The region for which to retrieve Batch service supported SKUs.
//   - options - LocationClientListSupportedVirtualMachineSKUsOptions contains the optional parameters for the LocationClient.NewListSupportedVirtualMachineSKUsPager
//     method.
func (client *LocationClient) NewListSupportedVirtualMachineSKUsPager(locationName string, options *LocationClientListSupportedVirtualMachineSKUsOptions) *runtime.Pager[LocationClientListSupportedVirtualMachineSKUsResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocationClientListSupportedVirtualMachineSKUsResponse]{
		More: func(page LocationClientListSupportedVirtualMachineSKUsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationClientListSupportedVirtualMachineSKUsResponse) (LocationClientListSupportedVirtualMachineSKUsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LocationClient.NewListSupportedVirtualMachineSKUsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listSupportedVirtualMachineSKUsCreateRequest(ctx, locationName, options)
			}, nil)
			if err != nil {
				return LocationClientListSupportedVirtualMachineSKUsResponse{}, err
			}
			return client.listSupportedVirtualMachineSKUsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listSupportedVirtualMachineSKUsCreateRequest creates the ListSupportedVirtualMachineSKUs request.
func (client *LocationClient) listSupportedVirtualMachineSKUsCreateRequest(ctx context.Context, locationName string, options *LocationClientListSupportedVirtualMachineSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Batch/locations/{locationName}/virtualMachineSkus"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Maxresults != nil {
		reqQP.Set("maxresults", strconv.FormatInt(int64(*options.Maxresults), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSupportedVirtualMachineSKUsHandleResponse handles the ListSupportedVirtualMachineSKUs response.
func (client *LocationClient) listSupportedVirtualMachineSKUsHandleResponse(resp *http.Response) (LocationClientListSupportedVirtualMachineSKUsResponse, error) {
	result := LocationClientListSupportedVirtualMachineSKUsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SupportedSKUsResult); err != nil {
		return LocationClientListSupportedVirtualMachineSKUsResponse{}, err
	}
	return result, nil
}