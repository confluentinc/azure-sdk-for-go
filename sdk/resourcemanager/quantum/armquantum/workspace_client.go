//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

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

// WorkspaceClient contains the methods for the Workspace group.
// Don't use this type directly, use NewWorkspaceClient() instead.
type WorkspaceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceClient creates a new instance of WorkspaceClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Check the availability of the resource name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-10-preview
//   - locationName - Location.
//   - checkNameAvailabilityParameters - The name and type of the resource.
//   - options - WorkspaceClientCheckNameAvailabilityOptions contains the optional parameters for the WorkspaceClient.CheckNameAvailability
//     method.
func (client *WorkspaceClient) CheckNameAvailability(ctx context.Context, locationName string, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *WorkspaceClientCheckNameAvailabilityOptions) (WorkspaceClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "WorkspaceClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, locationName, checkNameAvailabilityParameters, options)
	if err != nil {
		return WorkspaceClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *WorkspaceClient) checkNameAvailabilityCreateRequest(ctx context.Context, locationName string, checkNameAvailabilityParameters CheckNameAvailabilityParameters, options *WorkspaceClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Quantum/locations/{locationName}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkNameAvailabilityParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *WorkspaceClient) checkNameAvailabilityHandleResponse(resp *http.Response) (WorkspaceClientCheckNameAvailabilityResponse, error) {
	result := WorkspaceClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return WorkspaceClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}