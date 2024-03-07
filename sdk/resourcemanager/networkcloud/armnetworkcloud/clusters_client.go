//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

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

// ClustersClient contains the methods for the Clusters group.
// Don't use this type directly, use NewClustersClient() instead.
type ClustersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClustersClient creates a new instance of ClustersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClustersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClustersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClustersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new cluster or update the properties of the cluster if it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - clusterParameters - The request body.
//   - options - ClustersClientBeginCreateOrUpdateOptions contains the optional parameters for the ClustersClient.BeginCreateOrUpdate
//     method.
func (client *ClustersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, clusterParameters Cluster, options *ClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClustersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, clusterParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClustersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a new cluster or update the properties of the cluster if it exists.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ClustersClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, clusterParameters Cluster, options *ClustersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ClustersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, clusterParameters, options)
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
func (client *ClustersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, clusterParameters Cluster, options *ClustersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ClustersClientBeginDeleteOptions contains the optional parameters for the ClustersClient.BeginDelete method.
func (client *ClustersClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*runtime.Poller[ClustersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClustersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ClustersClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ClustersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, options)
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
func (client *ClustersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDeploy - Deploy the cluster to the provided rack.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ClustersClientBeginDeployOptions contains the optional parameters for the ClustersClient.BeginDeploy method.
func (client *ClustersClient) BeginDeploy(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeployOptions) (*runtime.Poller[ClustersClientDeployResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deploy(ctx, resourceGroupName, clusterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientDeployResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClustersClientDeployResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Deploy - Deploy the cluster to the provided rack.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ClustersClient) deploy(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeployOptions) (*http.Response, error) {
	var err error
	const operationName = "ClustersClient.BeginDeploy"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deployCreateRequest(ctx, resourceGroupName, clusterName, options)
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

// deployCreateRequest creates the Deploy request.
func (client *ClustersClient) deployCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientBeginDeployOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}/deploy"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.ClusterDeployParameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.ClusterDeployParameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// Get - Get properties of the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - ClustersClientGetOptions contains the optional parameters for the ClustersClient.Get method.
func (client *ClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientGetOptions) (ClustersClientGetResponse, error) {
	var err error
	const operationName = "ClustersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClustersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClustersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ClustersClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ClustersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClustersClient) getHandleResponse(resp *http.Response) (ClustersClientGetResponse, error) {
	result := ClustersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Cluster); err != nil {
		return ClustersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of clusters in the provided resource group.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ClustersClientListByResourceGroupOptions contains the optional parameters for the ClustersClient.NewListByResourceGroupPager
//     method.
func (client *ClustersClient) NewListByResourceGroupPager(resourceGroupName string, options *ClustersClientListByResourceGroupOptions) *runtime.Pager[ClustersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListByResourceGroupResponse]{
		More: func(page ClustersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListByResourceGroupResponse) (ClustersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ClustersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ClustersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ClustersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClustersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters"
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
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ClustersClient) listByResourceGroupHandleResponse(resp *http.Response) (ClustersClientListByResourceGroupResponse, error) {
	result := ClustersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterList); err != nil {
		return ClustersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of clusters in the provided subscription.
//
// Generated from API version 2023-07-01
//   - options - ClustersClientListBySubscriptionOptions contains the optional parameters for the ClustersClient.NewListBySubscriptionPager
//     method.
func (client *ClustersClient) NewListBySubscriptionPager(options *ClustersClientListBySubscriptionOptions) *runtime.Pager[ClustersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClustersClientListBySubscriptionResponse]{
		More: func(page ClustersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClustersClientListBySubscriptionResponse) (ClustersClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ClustersClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ClustersClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ClustersClient) listBySubscriptionCreateRequest(ctx context.Context, options *ClustersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/clusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ClustersClient) listBySubscriptionHandleResponse(resp *http.Response) (ClustersClientListBySubscriptionResponse, error) {
	result := ClustersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterList); err != nil {
		return ClustersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch the properties of the provided cluster, or update the tags associated with the cluster. Properties
// and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - clusterUpdateParameters - The request body.
//   - options - ClustersClientBeginUpdateOptions contains the optional parameters for the ClustersClient.BeginUpdate method.
func (client *ClustersClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateParameters ClusterPatchParameters, options *ClustersClientBeginUpdateOptions) (*runtime.Poller[ClustersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, clusterUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClustersClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch the properties of the provided cluster, or update the tags associated with the cluster. Properties and tag
// updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ClustersClient) update(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateParameters ClusterPatchParameters, options *ClustersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ClustersClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, clusterUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *ClustersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateParameters ClusterPatchParameters, options *ClustersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateVersion - Update the version of the provided cluster to one of the available supported versions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - clusterUpdateVersionParameters - The request body.
//   - options - ClustersClientBeginUpdateVersionOptions contains the optional parameters for the ClustersClient.BeginUpdateVersion
//     method.
func (client *ClustersClient) BeginUpdateVersion(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateVersionParameters ClusterUpdateVersionParameters, options *ClustersClientBeginUpdateVersionOptions) (*runtime.Poller[ClustersClientUpdateVersionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateVersion(ctx, resourceGroupName, clusterName, clusterUpdateVersionParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClustersClientUpdateVersionResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClustersClientUpdateVersionResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateVersion - Update the version of the provided cluster to one of the available supported versions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *ClustersClient) updateVersion(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateVersionParameters ClusterUpdateVersionParameters, options *ClustersClientBeginUpdateVersionOptions) (*http.Response, error) {
	var err error
	const operationName = "ClustersClient.BeginUpdateVersion"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateVersionCreateRequest(ctx, resourceGroupName, clusterName, clusterUpdateVersionParameters, options)
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

// updateVersionCreateRequest creates the UpdateVersion request.
func (client *ClustersClient) updateVersionCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateVersionParameters ClusterUpdateVersionParameters, options *ClustersClientBeginUpdateVersionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}/updateVersion"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterUpdateVersionParameters); err != nil {
		return nil, err
	}
	return req, nil
}