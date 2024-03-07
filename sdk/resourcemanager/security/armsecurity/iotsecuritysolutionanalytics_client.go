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
	"strings"
)

// IotSecuritySolutionAnalyticsClient contains the methods for the IotSecuritySolutionAnalytics group.
// Don't use this type directly, use NewIotSecuritySolutionAnalyticsClient() instead.
type IotSecuritySolutionAnalyticsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIotSecuritySolutionAnalyticsClient creates a new instance of IotSecuritySolutionAnalyticsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIotSecuritySolutionAnalyticsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IotSecuritySolutionAnalyticsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IotSecuritySolutionAnalyticsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Use this method to get IoT Security Analytics metrics.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - solutionName - The name of the IoT Security solution.
//   - options - IotSecuritySolutionAnalyticsClientGetOptions contains the optional parameters for the IotSecuritySolutionAnalyticsClient.Get
//     method.
func (client *IotSecuritySolutionAnalyticsClient) Get(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsClientGetOptions) (IotSecuritySolutionAnalyticsClientGetResponse, error) {
	var err error
	const operationName = "IotSecuritySolutionAnalyticsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return IotSecuritySolutionAnalyticsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IotSecuritySolutionAnalyticsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IotSecuritySolutionAnalyticsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IotSecuritySolutionAnalyticsClient) getCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IotSecuritySolutionAnalyticsClient) getHandleResponse(resp *http.Response) (IotSecuritySolutionAnalyticsClientGetResponse, error) {
	result := IotSecuritySolutionAnalyticsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionAnalyticsModel); err != nil {
		return IotSecuritySolutionAnalyticsClientGetResponse{}, err
	}
	return result, nil
}

// List - Use this method to get IoT security Analytics metrics in an array.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-08-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - solutionName - The name of the IoT Security solution.
//   - options - IotSecuritySolutionAnalyticsClientListOptions contains the optional parameters for the IotSecuritySolutionAnalyticsClient.List
//     method.
func (client *IotSecuritySolutionAnalyticsClient) List(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsClientListOptions) (IotSecuritySolutionAnalyticsClientListResponse, error) {
	var err error
	const operationName = "IotSecuritySolutionAnalyticsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, solutionName, options)
	if err != nil {
		return IotSecuritySolutionAnalyticsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IotSecuritySolutionAnalyticsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IotSecuritySolutionAnalyticsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *IotSecuritySolutionAnalyticsClient) listCreateRequest(ctx context.Context, resourceGroupName string, solutionName string, options *IotSecuritySolutionAnalyticsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/iotSecuritySolutions/{solutionName}/analyticsModels"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if solutionName == "" {
		return nil, errors.New("parameter solutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{solutionName}", url.PathEscape(solutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IotSecuritySolutionAnalyticsClient) listHandleResponse(resp *http.Response) (IotSecuritySolutionAnalyticsClientListResponse, error) {
	result := IotSecuritySolutionAnalyticsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IoTSecuritySolutionAnalyticsModelList); err != nil {
		return IotSecuritySolutionAnalyticsClientListResponse{}, err
	}
	return result, nil
}