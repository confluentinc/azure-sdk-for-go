//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// AdaptiveApplicationControlsClient contains the methods for the AdaptiveApplicationControls group.
// Don't use this type directly, use NewAdaptiveApplicationControlsClient() instead.
type AdaptiveApplicationControlsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAdaptiveApplicationControlsClient creates a new instance of AdaptiveApplicationControlsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAdaptiveApplicationControlsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AdaptiveApplicationControlsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AdaptiveApplicationControlsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Delete an application control machine group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - groupName - Name of an application control machine group
//   - options - AdaptiveApplicationControlsClientDeleteOptions contains the optional parameters for the AdaptiveApplicationControlsClient.Delete
//     method.
func (client *AdaptiveApplicationControlsClient) Delete(ctx context.Context, ascLocation string, groupName string, options *AdaptiveApplicationControlsClientDeleteOptions) (AdaptiveApplicationControlsClientDeleteResponse, error) {
	var err error
	const operationName = "AdaptiveApplicationControlsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, ascLocation, groupName, options)
	if err != nil {
		return AdaptiveApplicationControlsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdaptiveApplicationControlsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AdaptiveApplicationControlsClientDeleteResponse{}, err
	}
	return AdaptiveApplicationControlsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AdaptiveApplicationControlsClient) deleteCreateRequest(ctx context.Context, ascLocation string, groupName string, options *AdaptiveApplicationControlsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/applicationWhitelistings/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an application control VM/server group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - groupName - Name of an application control machine group
//   - options - AdaptiveApplicationControlsClientGetOptions contains the optional parameters for the AdaptiveApplicationControlsClient.Get
//     method.
func (client *AdaptiveApplicationControlsClient) Get(ctx context.Context, ascLocation string, groupName string, options *AdaptiveApplicationControlsClientGetOptions) (AdaptiveApplicationControlsClientGetResponse, error) {
	var err error
	const operationName = "AdaptiveApplicationControlsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, ascLocation, groupName, options)
	if err != nil {
		return AdaptiveApplicationControlsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdaptiveApplicationControlsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AdaptiveApplicationControlsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AdaptiveApplicationControlsClient) getCreateRequest(ctx context.Context, ascLocation string, groupName string, options *AdaptiveApplicationControlsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/applicationWhitelistings/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AdaptiveApplicationControlsClient) getHandleResponse(resp *http.Response) (AdaptiveApplicationControlsClientGetResponse, error) {
	result := AdaptiveApplicationControlsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdaptiveApplicationControlGroup); err != nil {
		return AdaptiveApplicationControlsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of application control machine groups for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - options - AdaptiveApplicationControlsClientListOptions contains the optional parameters for the AdaptiveApplicationControlsClient.List
//     method.
func (client *AdaptiveApplicationControlsClient) List(ctx context.Context, options *AdaptiveApplicationControlsClientListOptions) (AdaptiveApplicationControlsClientListResponse, error) {
	var err error
	const operationName = "AdaptiveApplicationControlsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return AdaptiveApplicationControlsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdaptiveApplicationControlsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AdaptiveApplicationControlsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *AdaptiveApplicationControlsClient) listCreateRequest(ctx context.Context, options *AdaptiveApplicationControlsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/applicationWhitelistings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	if options != nil && options.IncludePathRecommendations != nil {
		reqQP.Set("includePathRecommendations", strconv.FormatBool(*options.IncludePathRecommendations))
	}
	if options != nil && options.Summary != nil {
		reqQP.Set("summary", strconv.FormatBool(*options.Summary))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AdaptiveApplicationControlsClient) listHandleResponse(resp *http.Response) (AdaptiveApplicationControlsClientListResponse, error) {
	result := AdaptiveApplicationControlsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdaptiveApplicationControlGroups); err != nil {
		return AdaptiveApplicationControlsClientListResponse{}, err
	}
	return result, nil
}

// Put - Update an application control machine group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - groupName - Name of an application control machine group
//   - options - AdaptiveApplicationControlsClientPutOptions contains the optional parameters for the AdaptiveApplicationControlsClient.Put
//     method.
func (client *AdaptiveApplicationControlsClient) Put(ctx context.Context, ascLocation string, groupName string, body AdaptiveApplicationControlGroup, options *AdaptiveApplicationControlsClientPutOptions) (AdaptiveApplicationControlsClientPutResponse, error) {
	var err error
	const operationName = "AdaptiveApplicationControlsClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, ascLocation, groupName, body, options)
	if err != nil {
		return AdaptiveApplicationControlsClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AdaptiveApplicationControlsClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AdaptiveApplicationControlsClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *AdaptiveApplicationControlsClient) putCreateRequest(ctx context.Context, ascLocation string, groupName string, body AdaptiveApplicationControlGroup, options *AdaptiveApplicationControlsClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/applicationWhitelistings/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *AdaptiveApplicationControlsClient) putHandleResponse(resp *http.Response) (AdaptiveApplicationControlsClientPutResponse, error) {
	result := AdaptiveApplicationControlsClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdaptiveApplicationControlGroup); err != nil {
		return AdaptiveApplicationControlsClientPutResponse{}, err
	}
	return result, nil
}