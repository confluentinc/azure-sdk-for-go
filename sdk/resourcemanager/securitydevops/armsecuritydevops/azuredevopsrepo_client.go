//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecuritydevops

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

// AzureDevOpsRepoClient contains the methods for the AzureDevOpsRepo group.
// Don't use this type directly, use NewAzureDevOpsRepoClient() instead.
type AzureDevOpsRepoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureDevOpsRepoClient creates a new instance of AzureDevOpsRepoClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureDevOpsRepoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureDevOpsRepoClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureDevOpsRepoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Updates an Azure DevOps Repo.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureDevOpsConnectorName - Name of the AzureDevOps Connector.
//   - azureDevOpsOrgName - Name of the AzureDevOps Org.
//   - azureDevOpsProjectName - Name of the AzureDevOps Project.
//   - azureDevOpsRepoName - Name of the AzureDevOps Repo.
//   - azureDevOpsRepo - Azure DevOps Repo resource payload.
//   - options - AzureDevOpsRepoClientBeginCreateOrUpdateOptions contains the optional parameters for the AzureDevOpsRepoClient.BeginCreateOrUpdate
//     method.
func (client *AzureDevOpsRepoClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, azureDevOpsRepo AzureDevOpsRepo, options *AzureDevOpsRepoClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureDevOpsRepoClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsOrgName, azureDevOpsProjectName, azureDevOpsRepoName, azureDevOpsRepo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureDevOpsRepoClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureDevOpsRepoClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Updates an Azure DevOps Repo.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *AzureDevOpsRepoClient) createOrUpdate(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, azureDevOpsRepo AzureDevOpsRepo, options *AzureDevOpsRepoClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureDevOpsRepoClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsOrgName, azureDevOpsProjectName, azureDevOpsRepoName, azureDevOpsRepo, options)
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
func (client *AzureDevOpsRepoClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, azureDevOpsRepo AzureDevOpsRepo, options *AzureDevOpsRepoClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs/{azureDevOpsOrgName}/projects/{azureDevOpsProjectName}/repos/{azureDevOpsRepoName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	if azureDevOpsOrgName == "" {
		return nil, errors.New("parameter azureDevOpsOrgName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsOrgName}", url.PathEscape(azureDevOpsOrgName))
	if azureDevOpsProjectName == "" {
		return nil, errors.New("parameter azureDevOpsProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsProjectName}", url.PathEscape(azureDevOpsProjectName))
	if azureDevOpsRepoName == "" {
		return nil, errors.New("parameter azureDevOpsRepoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsRepoName}", url.PathEscape(azureDevOpsRepoName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, azureDevOpsRepo); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Returns a monitored AzureDevOps Project resource for a given ID.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureDevOpsConnectorName - Name of the AzureDevOps Connector.
//   - azureDevOpsOrgName - Name of the AzureDevOps Org.
//   - azureDevOpsProjectName - Name of the AzureDevOps Project.
//   - azureDevOpsRepoName - Name of the AzureDevOps Repo.
//   - options - AzureDevOpsRepoClientGetOptions contains the optional parameters for the AzureDevOpsRepoClient.Get method.
func (client *AzureDevOpsRepoClient) Get(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, options *AzureDevOpsRepoClientGetOptions) (AzureDevOpsRepoClientGetResponse, error) {
	var err error
	const operationName = "AzureDevOpsRepoClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsOrgName, azureDevOpsProjectName, azureDevOpsRepoName, options)
	if err != nil {
		return AzureDevOpsRepoClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AzureDevOpsRepoClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AzureDevOpsRepoClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AzureDevOpsRepoClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, options *AzureDevOpsRepoClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs/{azureDevOpsOrgName}/projects/{azureDevOpsProjectName}/repos/{azureDevOpsRepoName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	if azureDevOpsOrgName == "" {
		return nil, errors.New("parameter azureDevOpsOrgName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsOrgName}", url.PathEscape(azureDevOpsOrgName))
	if azureDevOpsProjectName == "" {
		return nil, errors.New("parameter azureDevOpsProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsProjectName}", url.PathEscape(azureDevOpsProjectName))
	if azureDevOpsRepoName == "" {
		return nil, errors.New("parameter azureDevOpsRepoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsRepoName}", url.PathEscape(azureDevOpsRepoName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureDevOpsRepoClient) getHandleResponse(resp *http.Response) (AzureDevOpsRepoClientGetResponse, error) {
	result := AzureDevOpsRepoClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsRepo); err != nil {
		return AzureDevOpsRepoClientGetResponse{}, err
	}
	return result, nil
}

//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureDevOpsConnectorName - Name of the AzureDevOps Connector.
//   - azureDevOpsOrgName - Name of the AzureDevOps Org.
//   - azureDevOpsProjectName - Name of the AzureDevOps Project.
//   - options - AzureDevOpsRepoClientListOptions contains the optional parameters for the AzureDevOpsRepoClient.NewListPager
//     method.
func (client *AzureDevOpsRepoClient) NewListPager(resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, options *AzureDevOpsRepoClientListOptions) *runtime.Pager[AzureDevOpsRepoClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureDevOpsRepoClientListResponse]{
		More: func(page AzureDevOpsRepoClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureDevOpsRepoClientListResponse) (AzureDevOpsRepoClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureDevOpsRepoClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsOrgName, azureDevOpsProjectName, options)
			}, nil)
			if err != nil {
				return AzureDevOpsRepoClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AzureDevOpsRepoClient) listCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, options *AzureDevOpsRepoClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs/{azureDevOpsOrgName}/projects/{azureDevOpsProjectName}/repos"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	if azureDevOpsOrgName == "" {
		return nil, errors.New("parameter azureDevOpsOrgName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsOrgName}", url.PathEscape(azureDevOpsOrgName))
	if azureDevOpsProjectName == "" {
		return nil, errors.New("parameter azureDevOpsProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsProjectName}", url.PathEscape(azureDevOpsProjectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureDevOpsRepoClient) listHandleResponse(resp *http.Response) (AzureDevOpsRepoClientListResponse, error) {
	result := AzureDevOpsRepoClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsRepoListResponse); err != nil {
		return AzureDevOpsRepoClientListResponse{}, err
	}
	return result, nil
}

//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureDevOpsConnectorName - Name of the AzureDevOps Connector.
//   - options - AzureDevOpsRepoClientListByConnectorOptions contains the optional parameters for the AzureDevOpsRepoClient.NewListByConnectorPager
//     method.
func (client *AzureDevOpsRepoClient) NewListByConnectorPager(resourceGroupName string, azureDevOpsConnectorName string, options *AzureDevOpsRepoClientListByConnectorOptions) *runtime.Pager[AzureDevOpsRepoClientListByConnectorResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureDevOpsRepoClientListByConnectorResponse]{
		More: func(page AzureDevOpsRepoClientListByConnectorResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureDevOpsRepoClientListByConnectorResponse) (AzureDevOpsRepoClientListByConnectorResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureDevOpsRepoClient.NewListByConnectorPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByConnectorCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, options)
			}, nil)
			if err != nil {
				return AzureDevOpsRepoClientListByConnectorResponse{}, err
			}
			return client.listByConnectorHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByConnectorCreateRequest creates the ListByConnector request.
func (client *AzureDevOpsRepoClient) listByConnectorCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, options *AzureDevOpsRepoClientListByConnectorOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/repos"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByConnectorHandleResponse handles the ListByConnector response.
func (client *AzureDevOpsRepoClient) listByConnectorHandleResponse(resp *http.Response) (AzureDevOpsRepoClientListByConnectorResponse, error) {
	result := AzureDevOpsRepoClientListByConnectorResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureDevOpsRepoListResponse); err != nil {
		return AzureDevOpsRepoClientListByConnectorResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update monitored AzureDevOps Project details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - azureDevOpsConnectorName - Name of the AzureDevOps Connector.
//   - azureDevOpsOrgName - Name of the AzureDevOps Org.
//   - azureDevOpsProjectName - Name of the AzureDevOps Project.
//   - azureDevOpsRepoName - Name of the AzureDevOps Repo.
//   - azureDevOpsRepo - Azure DevOps Org resource payload.
//   - options - AzureDevOpsRepoClientBeginUpdateOptions contains the optional parameters for the AzureDevOpsRepoClient.BeginUpdate
//     method.
func (client *AzureDevOpsRepoClient) BeginUpdate(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, azureDevOpsRepo AzureDevOpsRepo, options *AzureDevOpsRepoClientBeginUpdateOptions) (*runtime.Poller[AzureDevOpsRepoClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsOrgName, azureDevOpsProjectName, azureDevOpsRepoName, azureDevOpsRepo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AzureDevOpsRepoClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[AzureDevOpsRepoClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update monitored AzureDevOps Project details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *AzureDevOpsRepoClient) update(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, azureDevOpsRepo AzureDevOpsRepo, options *AzureDevOpsRepoClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "AzureDevOpsRepoClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureDevOpsConnectorName, azureDevOpsOrgName, azureDevOpsProjectName, azureDevOpsRepoName, azureDevOpsRepo, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *AzureDevOpsRepoClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureDevOpsConnectorName string, azureDevOpsOrgName string, azureDevOpsProjectName string, azureDevOpsRepoName string, azureDevOpsRepo AzureDevOpsRepo, options *AzureDevOpsRepoClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs/{azureDevOpsOrgName}/projects/{azureDevOpsProjectName}/repos/{azureDevOpsRepoName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureDevOpsConnectorName == "" {
		return nil, errors.New("parameter azureDevOpsConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsConnectorName}", url.PathEscape(azureDevOpsConnectorName))
	if azureDevOpsOrgName == "" {
		return nil, errors.New("parameter azureDevOpsOrgName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsOrgName}", url.PathEscape(azureDevOpsOrgName))
	if azureDevOpsProjectName == "" {
		return nil, errors.New("parameter azureDevOpsProjectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsProjectName}", url.PathEscape(azureDevOpsProjectName))
	if azureDevOpsRepoName == "" {
		return nil, errors.New("parameter azureDevOpsRepoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureDevOpsRepoName}", url.PathEscape(azureDevOpsRepoName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, azureDevOpsRepo); err != nil {
		return nil, err
	}
	return req, nil
}