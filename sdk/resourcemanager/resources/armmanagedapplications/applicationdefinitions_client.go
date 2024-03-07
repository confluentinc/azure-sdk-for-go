//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

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

// ApplicationDefinitionsClient contains the methods for the ApplicationDefinitions group.
// Don't use this type directly, use NewApplicationDefinitionsClient() instead.
type ApplicationDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationDefinitionsClient creates a new instance of ApplicationDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - parameters - Parameters supplied to the create or update an managed application definition.
//   - options - ApplicationDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ApplicationDefinitionsClient.BeginCreateOrUpdate
//     method.
func (client *ApplicationDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ApplicationDefinitionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationDefinitionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ApplicationDefinitionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a new managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *ApplicationDefinitionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationDefinitionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdateByID - Creates a new managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - parameters - Parameters supplied to the create or update a managed application definition.
//   - options - ApplicationDefinitionsClientBeginCreateOrUpdateByIDOptions contains the optional parameters for the ApplicationDefinitionsClient.BeginCreateOrUpdateByID
//     method.
func (client *ApplicationDefinitionsClient) BeginCreateOrUpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateByIDOptions) (*runtime.Poller[ApplicationDefinitionsClientCreateOrUpdateByIDResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateByID(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationDefinitionsClientCreateOrUpdateByIDResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ApplicationDefinitionsClientCreateOrUpdateByIDResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdateByID - Creates a new managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *ApplicationDefinitionsClient) createOrUpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateByIDOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationDefinitionsClient.BeginCreateOrUpdateByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
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

// createOrUpdateByIDCreateRequest creates the CreateOrUpdateByID request.
func (client *ApplicationDefinitionsClient) createOrUpdateByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientBeginCreateOrUpdateByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition to delete.
//   - options - ApplicationDefinitionsClientBeginDeleteOptions contains the optional parameters for the ApplicationDefinitionsClient.BeginDelete
//     method.
func (client *ApplicationDefinitionsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteOptions) (*runtime.Poller[ApplicationDefinitionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, applicationDefinitionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationDefinitionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ApplicationDefinitionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *ApplicationDefinitionsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationDefinitionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeleteByID - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - options - ApplicationDefinitionsClientBeginDeleteByIDOptions contains the optional parameters for the ApplicationDefinitionsClient.BeginDeleteByID
//     method.
func (client *ApplicationDefinitionsClient) BeginDeleteByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteByIDOptions) (*runtime.Poller[ApplicationDefinitionsClientDeleteByIDResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteByID(ctx, resourceGroupName, applicationDefinitionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ApplicationDefinitionsClientDeleteByIDResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ApplicationDefinitionsClientDeleteByIDResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DeleteByID - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *ApplicationDefinitionsClient) deleteByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteByIDOptions) (*http.Response, error) {
	var err error
	const operationName = "ApplicationDefinitionsClient.BeginDeleteByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *ApplicationDefinitionsClient) deleteByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientBeginDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - options - ApplicationDefinitionsClientGetOptions contains the optional parameters for the ApplicationDefinitionsClient.Get
//     method.
func (client *ApplicationDefinitionsClient) Get(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetOptions) (ApplicationDefinitionsClientGetResponse, error) {
	var err error
	const operationName = "ApplicationDefinitionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplicationDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationDefinitionsClient) getHandleResponse(resp *http.Response) (ApplicationDefinitionsClientGetResponse, error) {
	result := ApplicationDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetByID - Gets the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - options - ApplicationDefinitionsClientGetByIDOptions contains the optional parameters for the ApplicationDefinitionsClient.GetByID
//     method.
func (client *ApplicationDefinitionsClient) GetByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetByIDOptions) (ApplicationDefinitionsClientGetByIDResponse, error) {
	var err error
	const operationName = "ApplicationDefinitionsClient.GetByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	resp, err := client.getByIDHandleResponse(httpResp)
	return resp, err
}

// getByIDCreateRequest creates the GetByID request.
func (client *ApplicationDefinitionsClient) getByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *ApplicationDefinitionsClient) getByIDHandleResponse(resp *http.Response) (ApplicationDefinitionsClientGetByIDResponse, error) {
	result := ApplicationDefinitionsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists the managed application definitions in a resource group.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ApplicationDefinitionsClientListByResourceGroupOptions contains the optional parameters for the ApplicationDefinitionsClient.NewListByResourceGroupPager
//     method.
func (client *ApplicationDefinitionsClient) NewListByResourceGroupPager(resourceGroupName string, options *ApplicationDefinitionsClientListByResourceGroupOptions) *runtime.Pager[ApplicationDefinitionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationDefinitionsClientListByResourceGroupResponse]{
		More: func(page ApplicationDefinitionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationDefinitionsClientListByResourceGroupResponse) (ApplicationDefinitionsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ApplicationDefinitionsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationDefinitionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationDefinitionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationDefinitionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationDefinitionsClientListByResourceGroupResponse, error) {
	result := ApplicationDefinitionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinitionListResult); err != nil {
		return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}