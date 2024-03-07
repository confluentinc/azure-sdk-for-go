//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationServiceEnvironmentManagedApisClient contains the methods for the IntegrationServiceEnvironmentManagedApis group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentManagedApisClient() instead.
type IntegrationServiceEnvironmentManagedApisClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationServiceEnvironmentManagedApisClient creates a new instance of IntegrationServiceEnvironmentManagedApisClient with the specified values.
//   - subscriptionID - The subscription id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationServiceEnvironmentManagedApisClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationServiceEnvironmentManagedApisClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationServiceEnvironmentManagedApisClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDelete - Deletes the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - apiName - The api name.
//   - options - IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.BeginDelete
//     method.
func (client *IntegrationServiceEnvironmentManagedApisClient) BeginDelete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions) (*runtime.Poller[IntegrationServiceEnvironmentManagedApisClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IntegrationServiceEnvironmentManagedApisClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IntegrationServiceEnvironmentManagedApisClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *IntegrationServiceEnvironmentManagedApisClient) deleteOperation(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentManagedApisClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationServiceEnvironmentManagedApisClient) deleteCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group name.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - apiName - The api name.
//   - options - IntegrationServiceEnvironmentManagedApisClientGetOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.Get
//     method.
func (client *IntegrationServiceEnvironmentManagedApisClient) Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientGetOptions) (IntegrationServiceEnvironmentManagedApisClientGetResponse, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentManagedApisClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, options)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationServiceEnvironmentManagedApisClient) getCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, options *IntegrationServiceEnvironmentManagedApisClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationServiceEnvironmentManagedApisClient) getHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentManagedApisClientGetResponse, error) {
	result := IntegrationServiceEnvironmentManagedApisClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentManagedAPI); err != nil {
		return IntegrationServiceEnvironmentManagedApisClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the integration service environment managed Apis.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - options - IntegrationServiceEnvironmentManagedApisClientListOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.NewListPager
//     method.
func (client *IntegrationServiceEnvironmentManagedApisClient) NewListPager(resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentManagedApisClientListOptions) *runtime.Pager[IntegrationServiceEnvironmentManagedApisClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[IntegrationServiceEnvironmentManagedApisClientListResponse]{
		More: func(page IntegrationServiceEnvironmentManagedApisClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *IntegrationServiceEnvironmentManagedApisClientListResponse) (IntegrationServiceEnvironmentManagedApisClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "IntegrationServiceEnvironmentManagedApisClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
			}, nil)
			if err != nil {
				return IntegrationServiceEnvironmentManagedApisClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *IntegrationServiceEnvironmentManagedApisClient) listCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentManagedApisClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationServiceEnvironmentManagedApisClient) listHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentManagedApisClientListResponse, error) {
	result := IntegrationServiceEnvironmentManagedApisClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentManagedAPIListResult); err != nil {
		return IntegrationServiceEnvironmentManagedApisClientListResponse{}, err
	}
	return result, nil
}

// BeginPut - Puts the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
//   - resourceGroup - The resource group name.
//   - integrationServiceEnvironmentName - The integration service environment name.
//   - apiName - The api name.
//   - integrationServiceEnvironmentManagedAPI - The integration service environment managed api.
//   - options - IntegrationServiceEnvironmentManagedApisClientBeginPutOptions contains the optional parameters for the IntegrationServiceEnvironmentManagedApisClient.BeginPut
//     method.
func (client *IntegrationServiceEnvironmentManagedApisClient) BeginPut(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions) (*runtime.Poller[IntegrationServiceEnvironmentManagedApisClientPutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.put(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, integrationServiceEnvironmentManagedAPI, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[IntegrationServiceEnvironmentManagedApisClientPutResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[IntegrationServiceEnvironmentManagedApisClientPutResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Put - Puts the integration service environment managed Api.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-05-01
func (client *IntegrationServiceEnvironmentManagedApisClient) put(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions) (*http.Response, error) {
	var err error
	const operationName = "IntegrationServiceEnvironmentManagedApisClient.BeginPut"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, apiName, integrationServiceEnvironmentManagedAPI, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// putCreateRequest creates the Put request.
func (client *IntegrationServiceEnvironmentManagedApisClient) putCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, apiName string, integrationServiceEnvironmentManagedAPI IntegrationServiceEnvironmentManagedAPI, options *IntegrationServiceEnvironmentManagedApisClientBeginPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/managedApis/{apiName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	if apiName == "" {
		return nil, errors.New("parameter apiName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiName}", url.PathEscape(apiName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, integrationServiceEnvironmentManagedAPI); err != nil {
		return nil, err
	}
	return req, nil
}