//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

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

// LtrBackupOperationsClient contains the methods for the LtrBackupOperations group.
// Don't use this type directly, use NewLtrBackupOperationsClient() instead.
type LtrBackupOperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLtrBackupOperationsClient creates a new instance of LtrBackupOperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLtrBackupOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LtrBackupOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LtrBackupOperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the result of the give long term retention backup operation for the flexible server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - backupName - The name of the backup.
//   - options - LtrBackupOperationsClientGetOptions contains the optional parameters for the LtrBackupOperationsClient.Get method.
func (client *LtrBackupOperationsClient) Get(ctx context.Context, resourceGroupName string, serverName string, backupName string, options *LtrBackupOperationsClientGetOptions) (LtrBackupOperationsClientGetResponse, error) {
	var err error
	const operationName = "LtrBackupOperationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, backupName, options)
	if err != nil {
		return LtrBackupOperationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LtrBackupOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LtrBackupOperationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LtrBackupOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, backupName string, options *LtrBackupOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/ltrBackupOperations/{backupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if backupName == "" {
		return nil, errors.New("parameter backupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupName}", url.PathEscape(backupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LtrBackupOperationsClient) getHandleResponse(resp *http.Response) (LtrBackupOperationsClientGetResponse, error) {
	result := LtrBackupOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LtrServerBackupOperation); err != nil {
		return LtrBackupOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets the result of the give long term retention backup operations for the flexible server.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - LtrBackupOperationsClientListByServerOptions contains the optional parameters for the LtrBackupOperationsClient.NewListByServerPager
//     method.
func (client *LtrBackupOperationsClient) NewListByServerPager(resourceGroupName string, serverName string, options *LtrBackupOperationsClientListByServerOptions) *runtime.Pager[LtrBackupOperationsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[LtrBackupOperationsClientListByServerResponse]{
		More: func(page LtrBackupOperationsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LtrBackupOperationsClientListByServerResponse) (LtrBackupOperationsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LtrBackupOperationsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return LtrBackupOperationsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *LtrBackupOperationsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *LtrBackupOperationsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/ltrBackupOperations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *LtrBackupOperationsClient) listByServerHandleResponse(resp *http.Response) (LtrBackupOperationsClientListByServerResponse, error) {
	result := LtrBackupOperationsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LtrServerBackupOperationList); err != nil {
		return LtrBackupOperationsClientListByServerResponse{}, err
	}
	return result, nil
}
