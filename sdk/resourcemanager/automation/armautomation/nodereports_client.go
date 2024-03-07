//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// NodeReportsClient contains the methods for the NodeReports group.
// Don't use this type directly, use NewNodeReportsClient() instead.
type NodeReportsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNodeReportsClient creates a new instance of NodeReportsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNodeReportsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NodeReportsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NodeReportsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve the Dsc node report data by node id and report id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeID - The Dsc node id.
//   - reportID - The report id.
//   - options - NodeReportsClientGetOptions contains the optional parameters for the NodeReportsClient.Get method.
func (client *NodeReportsClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string, options *NodeReportsClientGetOptions) (NodeReportsClientGetResponse, error) {
	var err error
	const operationName = "NodeReportsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, reportID, options)
	if err != nil {
		return NodeReportsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NodeReportsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NodeReportsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NodeReportsClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string, options *NodeReportsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if reportID == "" {
		return nil, errors.New("parameter reportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportId}", url.PathEscape(reportID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NodeReportsClient) getHandleResponse(resp *http.Response) (NodeReportsClientGetResponse, error) {
	result := NodeReportsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNodeReport); err != nil {
		return NodeReportsClientGetResponse{}, err
	}
	return result, nil
}

// GetContent - Retrieve the Dsc node reports by node id and report id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeID - The Dsc node id.
//   - reportID - The report id.
//   - options - NodeReportsClientGetContentOptions contains the optional parameters for the NodeReportsClient.GetContent method.
func (client *NodeReportsClient) GetContent(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string, options *NodeReportsClientGetContentOptions) (NodeReportsClientGetContentResponse, error) {
	var err error
	const operationName = "NodeReportsClient.GetContent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getContentCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, reportID, options)
	if err != nil {
		return NodeReportsClientGetContentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NodeReportsClientGetContentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NodeReportsClientGetContentResponse{}, err
	}
	resp, err := client.getContentHandleResponse(httpResp)
	return resp, err
}

// getContentCreateRequest creates the GetContent request.
func (client *NodeReportsClient) getContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, reportID string, options *NodeReportsClientGetContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports/{reportId}/content"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if reportID == "" {
		return nil, errors.New("parameter reportID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{reportId}", url.PathEscape(reportID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getContentHandleResponse handles the GetContent response.
func (client *NodeReportsClient) getContentHandleResponse(resp *http.Response) (NodeReportsClientGetContentResponse, error) {
	result := NodeReportsClientGetContentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Interface); err != nil {
		return NodeReportsClientGetContentResponse{}, err
	}
	return result, nil
}

// NewListByNodePager - Retrieve the Dsc node report list by node id.
//
// Generated from API version 2020-01-13-preview
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - nodeID - The parameters supplied to the list operation.
//   - options - NodeReportsClientListByNodeOptions contains the optional parameters for the NodeReportsClient.NewListByNodePager
//     method.
func (client *NodeReportsClient) NewListByNodePager(resourceGroupName string, automationAccountName string, nodeID string, options *NodeReportsClientListByNodeOptions) *runtime.Pager[NodeReportsClientListByNodeResponse] {
	return runtime.NewPager(runtime.PagingHandler[NodeReportsClientListByNodeResponse]{
		More: func(page NodeReportsClientListByNodeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NodeReportsClientListByNodeResponse) (NodeReportsClientListByNodeResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NodeReportsClient.NewListByNodePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNodeCreateRequest(ctx, resourceGroupName, automationAccountName, nodeID, options)
			}, nil)
			if err != nil {
				return NodeReportsClientListByNodeResponse{}, err
			}
			return client.listByNodeHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNodeCreateRequest creates the ListByNode request.
func (client *NodeReportsClient) listByNodeCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, nodeID string, options *NodeReportsClientListByNodeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/nodes/{nodeId}/reports"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if nodeID == "" {
		return nil, errors.New("parameter nodeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeId}", url.PathEscape(nodeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNodeHandleResponse handles the ListByNode response.
func (client *NodeReportsClient) listByNodeHandleResponse(resp *http.Response) (NodeReportsClientListByNodeResponse, error) {
	result := NodeReportsClientListByNodeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DscNodeReportListResult); err != nil {
		return NodeReportsClientListByNodeResponse{}, err
	}
	return result, nil
}