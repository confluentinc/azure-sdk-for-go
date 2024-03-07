//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

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

// TemplateSpecVersionsClient contains the methods for the TemplateSpecVersions group.
// Don't use this type directly, use NewTemplateSpecVersionsClient() instead.
type TemplateSpecVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTemplateSpecVersionsClient creates a new instance of TemplateSpecVersionsClient with the specified values.
//   - subscriptionID - Subscription Id which forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTemplateSpecVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TemplateSpecVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TemplateSpecVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Template Spec version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - templateSpecVersion - The version of the Template Spec.
//   - templateSpecVersionModel - Template Spec Version supplied to the operation.
//   - options - TemplateSpecVersionsClientCreateOrUpdateOptions contains the optional parameters for the TemplateSpecVersionsClient.CreateOrUpdate
//     method.
func (client *TemplateSpecVersionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, templateSpecVersionModel TemplateSpecVersion, options *TemplateSpecVersionsClientCreateOrUpdateOptions) (TemplateSpecVersionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TemplateSpecVersionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, templateSpecName, templateSpecVersion, templateSpecVersionModel, options)
	if err != nil {
		return TemplateSpecVersionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TemplateSpecVersionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TemplateSpecVersionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TemplateSpecVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, templateSpecVersionModel TemplateSpecVersion, options *TemplateSpecVersionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}/versions/{templateSpecVersion}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	if templateSpecVersion == "" {
		return nil, errors.New("parameter templateSpecVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecVersion}", url.PathEscape(templateSpecVersion))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, templateSpecVersionModel); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TemplateSpecVersionsClient) createOrUpdateHandleResponse(resp *http.Response) (TemplateSpecVersionsClientCreateOrUpdateResponse, error) {
	result := TemplateSpecVersionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecVersion); err != nil {
		return TemplateSpecVersionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a specific version from a Template Spec. When operation completes, status code 200 returned without content.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - templateSpecVersion - The version of the Template Spec.
//   - options - TemplateSpecVersionsClientDeleteOptions contains the optional parameters for the TemplateSpecVersionsClient.Delete
//     method.
func (client *TemplateSpecVersionsClient) Delete(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientDeleteOptions) (TemplateSpecVersionsClientDeleteResponse, error) {
	var err error
	const operationName = "TemplateSpecVersionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, templateSpecName, templateSpecVersion, options)
	if err != nil {
		return TemplateSpecVersionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TemplateSpecVersionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TemplateSpecVersionsClientDeleteResponse{}, err
	}
	return TemplateSpecVersionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TemplateSpecVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}/versions/{templateSpecVersion}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	if templateSpecVersion == "" {
		return nil, errors.New("parameter templateSpecVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecVersion}", url.PathEscape(templateSpecVersion))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Template Spec version from a specific Template Spec.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - templateSpecVersion - The version of the Template Spec.
//   - options - TemplateSpecVersionsClientGetOptions contains the optional parameters for the TemplateSpecVersionsClient.Get
//     method.
func (client *TemplateSpecVersionsClient) Get(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientGetOptions) (TemplateSpecVersionsClientGetResponse, error) {
	var err error
	const operationName = "TemplateSpecVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, templateSpecName, templateSpecVersion, options)
	if err != nil {
		return TemplateSpecVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TemplateSpecVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TemplateSpecVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TemplateSpecVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}/versions/{templateSpecVersion}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	if templateSpecVersion == "" {
		return nil, errors.New("parameter templateSpecVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecVersion}", url.PathEscape(templateSpecVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TemplateSpecVersionsClient) getHandleResponse(resp *http.Response) (TemplateSpecVersionsClientGetResponse, error) {
	result := TemplateSpecVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecVersion); err != nil {
		return TemplateSpecVersionsClientGetResponse{}, err
	}
	return result, nil
}

// GetBuiltIn - Gets a Template Spec version from a specific built-in Template Spec.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - templateSpecName - Name of the Template Spec.
//   - templateSpecVersion - The version of the Template Spec.
//   - options - TemplateSpecVersionsClientGetBuiltInOptions contains the optional parameters for the TemplateSpecVersionsClient.GetBuiltIn
//     method.
func (client *TemplateSpecVersionsClient) GetBuiltIn(ctx context.Context, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientGetBuiltInOptions) (TemplateSpecVersionsClientGetBuiltInResponse, error) {
	var err error
	const operationName = "TemplateSpecVersionsClient.GetBuiltIn"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBuiltInCreateRequest(ctx, templateSpecName, templateSpecVersion, options)
	if err != nil {
		return TemplateSpecVersionsClientGetBuiltInResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TemplateSpecVersionsClientGetBuiltInResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TemplateSpecVersionsClientGetBuiltInResponse{}, err
	}
	resp, err := client.getBuiltInHandleResponse(httpResp)
	return resp, err
}

// getBuiltInCreateRequest creates the GetBuiltIn request.
func (client *TemplateSpecVersionsClient) getBuiltInCreateRequest(ctx context.Context, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientGetBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/builtInTemplateSpecs/{templateSpecName}/versions/{templateSpecVersion}"
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	if templateSpecVersion == "" {
		return nil, errors.New("parameter templateSpecVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecVersion}", url.PathEscape(templateSpecVersion))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBuiltInHandleResponse handles the GetBuiltIn response.
func (client *TemplateSpecVersionsClient) getBuiltInHandleResponse(resp *http.Response) (TemplateSpecVersionsClientGetBuiltInResponse, error) {
	result := TemplateSpecVersionsClientGetBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecVersion); err != nil {
		return TemplateSpecVersionsClientGetBuiltInResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the Template Spec versions in the specified Template Spec.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - options - TemplateSpecVersionsClientListOptions contains the optional parameters for the TemplateSpecVersionsClient.NewListPager
//     method.
func (client *TemplateSpecVersionsClient) NewListPager(resourceGroupName string, templateSpecName string, options *TemplateSpecVersionsClientListOptions) *runtime.Pager[TemplateSpecVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TemplateSpecVersionsClientListResponse]{
		More: func(page TemplateSpecVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TemplateSpecVersionsClientListResponse) (TemplateSpecVersionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TemplateSpecVersionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, templateSpecName, options)
			}, nil)
			if err != nil {
				return TemplateSpecVersionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TemplateSpecVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, options *TemplateSpecVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TemplateSpecVersionsClient) listHandleResponse(resp *http.Response) (TemplateSpecVersionsClientListResponse, error) {
	result := TemplateSpecVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecVersionsListResult); err != nil {
		return TemplateSpecVersionsClientListResponse{}, err
	}
	return result, nil
}

// NewListBuiltInsPager - Lists all the Template Spec versions in the specified built-in Template Spec.
//
// Generated from API version 2022-02-01
//   - templateSpecName - Name of the Template Spec.
//   - options - TemplateSpecVersionsClientListBuiltInsOptions contains the optional parameters for the TemplateSpecVersionsClient.NewListBuiltInsPager
//     method.
func (client *TemplateSpecVersionsClient) NewListBuiltInsPager(templateSpecName string, options *TemplateSpecVersionsClientListBuiltInsOptions) *runtime.Pager[TemplateSpecVersionsClientListBuiltInsResponse] {
	return runtime.NewPager(runtime.PagingHandler[TemplateSpecVersionsClientListBuiltInsResponse]{
		More: func(page TemplateSpecVersionsClientListBuiltInsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TemplateSpecVersionsClientListBuiltInsResponse) (TemplateSpecVersionsClientListBuiltInsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TemplateSpecVersionsClient.NewListBuiltInsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBuiltInsCreateRequest(ctx, templateSpecName, options)
			}, nil)
			if err != nil {
				return TemplateSpecVersionsClientListBuiltInsResponse{}, err
			}
			return client.listBuiltInsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBuiltInsCreateRequest creates the ListBuiltIns request.
func (client *TemplateSpecVersionsClient) listBuiltInsCreateRequest(ctx context.Context, templateSpecName string, options *TemplateSpecVersionsClientListBuiltInsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Resources/builtInTemplateSpecs/{templateSpecName}/versions"
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBuiltInsHandleResponse handles the ListBuiltIns response.
func (client *TemplateSpecVersionsClient) listBuiltInsHandleResponse(resp *http.Response) (TemplateSpecVersionsClientListBuiltInsResponse, error) {
	result := TemplateSpecVersionsClientListBuiltInsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecVersionsListResult); err != nil {
		return TemplateSpecVersionsClientListBuiltInsResponse{}, err
	}
	return result, nil
}

// Update - Updates Template Spec Version tags with specified values.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - templateSpecName - Name of the Template Spec.
//   - templateSpecVersion - The version of the Template Spec.
//   - options - TemplateSpecVersionsClientUpdateOptions contains the optional parameters for the TemplateSpecVersionsClient.Update
//     method.
func (client *TemplateSpecVersionsClient) Update(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientUpdateOptions) (TemplateSpecVersionsClientUpdateResponse, error) {
	var err error
	const operationName = "TemplateSpecVersionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, templateSpecName, templateSpecVersion, options)
	if err != nil {
		return TemplateSpecVersionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TemplateSpecVersionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TemplateSpecVersionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TemplateSpecVersionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, templateSpecName string, templateSpecVersion string, options *TemplateSpecVersionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Resources/templateSpecs/{templateSpecName}/versions/{templateSpecVersion}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if templateSpecName == "" {
		return nil, errors.New("parameter templateSpecName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecName}", url.PathEscape(templateSpecName))
	if templateSpecVersion == "" {
		return nil, errors.New("parameter templateSpecVersion cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{templateSpecVersion}", url.PathEscape(templateSpecVersion))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.TemplateSpecVersionUpdateModel != nil {
		if err := runtime.MarshalAsJSON(req, *options.TemplateSpecVersionUpdateModel); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TemplateSpecVersionsClient) updateHandleResponse(resp *http.Response) (TemplateSpecVersionsClientUpdateResponse, error) {
	result := TemplateSpecVersionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TemplateSpecVersion); err != nil {
		return TemplateSpecVersionsClientUpdateResponse{}, err
	}
	return result, nil
}