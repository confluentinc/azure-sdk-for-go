//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// LookingGlassClient contains the methods for the LookingGlass group.
// Don't use this type directly, use NewLookingGlassClient() instead.
type LookingGlassClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLookingGlassClient creates a new instance of LookingGlassClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLookingGlassClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LookingGlassClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LookingGlassClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Invoke - Run looking glass functionality
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01
//   - command - The command to be executed: ping, traceroute, bgpRoute.
//   - sourceType - The type of the source: Edge site or Azure Region.
//   - sourceLocation - The location of the source.
//   - destinationIP - The IP address of the destination.
//   - options - LookingGlassClientInvokeOptions contains the optional parameters for the LookingGlassClient.Invoke method.
func (client *LookingGlassClient) Invoke(ctx context.Context, command LookingGlassCommand, sourceType LookingGlassSourceType, sourceLocation string, destinationIP string, options *LookingGlassClientInvokeOptions) (LookingGlassClientInvokeResponse, error) {
	var err error
	const operationName = "LookingGlassClient.Invoke"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.invokeCreateRequest(ctx, command, sourceType, sourceLocation, destinationIP, options)
	if err != nil {
		return LookingGlassClientInvokeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LookingGlassClientInvokeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LookingGlassClientInvokeResponse{}, err
	}
	resp, err := client.invokeHandleResponse(httpResp)
	return resp, err
}

// invokeCreateRequest creates the Invoke request.
func (client *LookingGlassClient) invokeCreateRequest(ctx context.Context, command LookingGlassCommand, sourceType LookingGlassSourceType, sourceLocation string, destinationIP string, options *LookingGlassClientInvokeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/lookingGlass"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("command", string(command))
	reqQP.Set("sourceType", string(sourceType))
	reqQP.Set("sourceLocation", sourceLocation)
	reqQP.Set("destinationIP", destinationIP)
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// invokeHandleResponse handles the Invoke response.
func (client *LookingGlassClient) invokeHandleResponse(resp *http.Response) (LookingGlassClientInvokeResponse, error) {
	result := LookingGlassClientInvokeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LookingGlassOutput); err != nil {
		return LookingGlassClientInvokeResponse{}, err
	}
	return result, nil
}