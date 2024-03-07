//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift

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

// OpenShiftClustersClient contains the methods for the OpenShiftClusters group.
// Don't use this type directly, use NewOpenShiftClustersClient() instead.
type OpenShiftClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOpenShiftClustersClient creates a new instance of OpenShiftClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOpenShiftClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OpenShiftClustersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OpenShiftClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - parameters - The OpenShift cluster resource.
//   - options - OpenShiftClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the OpenShiftClustersClient.BeginCreateOrUpdate
//     method.
func (client *OpenShiftClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftCluster, options *OpenShiftClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[OpenShiftClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OpenShiftClustersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OpenShiftClustersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
func (client *OpenShiftClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftCluster, options *OpenShiftClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "OpenShiftClustersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
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
func (client *OpenShiftClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftCluster, options *OpenShiftClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation returns nothing.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - options - OpenShiftClustersClientBeginDeleteOptions contains the optional parameters for the OpenShiftClustersClient.BeginDelete
//     method.
func (client *OpenShiftClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientBeginDeleteOptions) (*runtime.Poller[OpenShiftClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OpenShiftClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OpenShiftClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation returns nothing.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
func (client *OpenShiftClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "OpenShiftClustersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
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
func (client *OpenShiftClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - options - OpenShiftClustersClientGetOptions contains the optional parameters for the OpenShiftClustersClient.Get method.
func (client *OpenShiftClustersClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientGetOptions) (OpenShiftClustersClientGetResponse, error) {
	var err error
	const operationName = "OpenShiftClustersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return OpenShiftClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OpenShiftClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OpenShiftClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OpenShiftClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OpenShiftClustersClient) getHandleResponse(resp *http.Response) (OpenShiftClustersClientGetResponse, error) {
	result := OpenShiftClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftCluster); err != nil {
		return OpenShiftClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The operation returns properties of each OpenShift cluster.
//
// Generated from API version 2023-09-04
//   - options - OpenShiftClustersClientListOptions contains the optional parameters for the OpenShiftClustersClient.NewListPager
//     method.
func (client *OpenShiftClustersClient) NewListPager(options *OpenShiftClustersClientListOptions) *runtime.Pager[OpenShiftClustersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OpenShiftClustersClientListResponse]{
		More: func(page OpenShiftClustersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OpenShiftClustersClientListResponse) (OpenShiftClustersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OpenShiftClustersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return OpenShiftClustersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *OpenShiftClustersClient) listCreateRequest(ctx context.Context, options *OpenShiftClustersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.RedHatOpenShift/openShiftClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OpenShiftClustersClient) listHandleResponse(resp *http.Response) (OpenShiftClustersClientListResponse, error) {
	result := OpenShiftClustersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftClusterList); err != nil {
		return OpenShiftClustersClientListResponse{}, err
	}
	return result, nil
}

// ListAdminCredentials - The operation returns the admin kubeconfig.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - options - OpenShiftClustersClientListAdminCredentialsOptions contains the optional parameters for the OpenShiftClustersClient.ListAdminCredentials
//     method.
func (client *OpenShiftClustersClient) ListAdminCredentials(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientListAdminCredentialsOptions) (OpenShiftClustersClientListAdminCredentialsResponse, error) {
	var err error
	const operationName = "OpenShiftClustersClient.ListAdminCredentials"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAdminCredentialsCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return OpenShiftClustersClientListAdminCredentialsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OpenShiftClustersClientListAdminCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OpenShiftClustersClientListAdminCredentialsResponse{}, err
	}
	resp, err := client.listAdminCredentialsHandleResponse(httpResp)
	return resp, err
}

// listAdminCredentialsCreateRequest creates the ListAdminCredentials request.
func (client *OpenShiftClustersClient) listAdminCredentialsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientListAdminCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}/listAdminCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAdminCredentialsHandleResponse handles the ListAdminCredentials response.
func (client *OpenShiftClustersClient) listAdminCredentialsHandleResponse(resp *http.Response) (OpenShiftClustersClientListAdminCredentialsResponse, error) {
	result := OpenShiftClustersClientListAdminCredentialsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftClusterAdminKubeconfig); err != nil {
		return OpenShiftClustersClientListAdminCredentialsResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - The operation returns properties of each OpenShift cluster.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - OpenShiftClustersClientListByResourceGroupOptions contains the optional parameters for the OpenShiftClustersClient.NewListByResourceGroupPager
//     method.
func (client *OpenShiftClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *OpenShiftClustersClientListByResourceGroupOptions) *runtime.Pager[OpenShiftClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[OpenShiftClustersClientListByResourceGroupResponse]{
		More: func(page OpenShiftClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OpenShiftClustersClientListByResourceGroupResponse) (OpenShiftClustersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OpenShiftClustersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return OpenShiftClustersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OpenShiftClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OpenShiftClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OpenShiftClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (OpenShiftClustersClientListByResourceGroupResponse, error) {
	result := OpenShiftClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftClusterList); err != nil {
		return OpenShiftClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListCredentials - The operation returns the credentials.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - options - OpenShiftClustersClientListCredentialsOptions contains the optional parameters for the OpenShiftClustersClient.ListCredentials
//     method.
func (client *OpenShiftClustersClient) ListCredentials(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientListCredentialsOptions) (OpenShiftClustersClientListCredentialsResponse, error) {
	var err error
	const operationName = "OpenShiftClustersClient.ListCredentials"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCredentialsCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return OpenShiftClustersClientListCredentialsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OpenShiftClustersClientListCredentialsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OpenShiftClustersClientListCredentialsResponse{}, err
	}
	resp, err := client.listCredentialsHandleResponse(httpResp)
	return resp, err
}

// listCredentialsCreateRequest creates the ListCredentials request.
func (client *OpenShiftClustersClient) listCredentialsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OpenShiftClustersClientListCredentialsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}/listCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listCredentialsHandleResponse handles the ListCredentials response.
func (client *OpenShiftClustersClient) listCredentialsHandleResponse(resp *http.Response) (OpenShiftClustersClientListCredentialsResponse, error) {
	result := OpenShiftClustersClientListCredentialsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OpenShiftClusterCredentials); err != nil {
		return OpenShiftClustersClientListCredentialsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the OpenShift cluster resource.
//   - parameters - The OpenShift cluster resource.
//   - options - OpenShiftClustersClientBeginUpdateOptions contains the optional parameters for the OpenShiftClustersClient.BeginUpdate
//     method.
func (client *OpenShiftClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftClusterUpdate, options *OpenShiftClustersClientBeginUpdateOptions) (*runtime.Poller[OpenShiftClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OpenShiftClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OpenShiftClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - The operation returns properties of a OpenShift cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-04
func (client *OpenShiftClustersClient) update(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftClusterUpdate, options *OpenShiftClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "OpenShiftClustersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *OpenShiftClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters OpenShiftClusterUpdate, options *OpenShiftClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RedHatOpenShift/openShiftClusters/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-04")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}