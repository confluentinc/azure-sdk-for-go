//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// PortalSettingsClient contains the methods for the PortalSettings group.
// Don't use this type directly, use NewPortalSettingsClient() instead.
type PortalSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPortalSettingsClient creates a new instance of PortalSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPortalSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PortalSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PortalSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// ListByService - Lists a collection of portalsettings defined within a service instance..
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - PortalSettingsClientListByServiceOptions contains the optional parameters for the PortalSettingsClient.ListByService
//     method.
func (client *PortalSettingsClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, options *PortalSettingsClientListByServiceOptions) (PortalSettingsClientListByServiceResponse, error) {
	var err error
	const operationName = "PortalSettingsClient.ListByService"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return PortalSettingsClientListByServiceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PortalSettingsClientListByServiceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PortalSettingsClientListByServiceResponse{}, err
	}
	resp, err := client.listByServiceHandleResponse(httpResp)
	return resp, err
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PortalSettingsClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PortalSettingsClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *PortalSettingsClient) listByServiceHandleResponse(resp *http.Response) (PortalSettingsClientListByServiceResponse, error) {
	result := PortalSettingsClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalSettingsCollection); err != nil {
		return PortalSettingsClientListByServiceResponse{}, err
	}
	return result, nil
}