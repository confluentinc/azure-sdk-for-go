//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DatabasesClient contains the methods for the Databases group.
// Don't use this type directly, use NewDatabasesClient() instead.
type DatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabasesClient creates a new instance of DatabasesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabasesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// AddPrincipals - Add Database principals permissions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - databasePrincipalsToAdd - List of database principals to add.
//   - options - DatabasesClientAddPrincipalsOptions contains the optional parameters for the DatabasesClient.AddPrincipals method.
func (client *DatabasesClient) AddPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest, options *DatabasesClientAddPrincipalsOptions) (DatabasesClientAddPrincipalsResponse, error) {
	var err error
	const operationName = "DatabasesClient.AddPrincipals"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.addPrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToAdd, options)
	if err != nil {
		return DatabasesClientAddPrincipalsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabasesClientAddPrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabasesClientAddPrincipalsResponse{}, err
	}
	resp, err := client.addPrincipalsHandleResponse(httpResp)
	return resp, err
}

// addPrincipalsCreateRequest creates the AddPrincipals request.
func (client *DatabasesClient) addPrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest, options *DatabasesClientAddPrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/addPrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, databasePrincipalsToAdd); err != nil {
		return nil, err
	}
	return req, nil
}

// addPrincipalsHandleResponse handles the AddPrincipals response.
func (client *DatabasesClient) addPrincipalsHandleResponse(resp *http.Response) (DatabasesClientAddPrincipalsResponse, error) {
	result := DatabasesClientAddPrincipalsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesClientAddPrincipalsResponse{}, err
	}
	return result, nil
}

// CheckNameAvailability - Checks that the databases resource name is valid and is not already in use.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - resourceName - The name of the resource.
//   - options - DatabasesClientCheckNameAvailabilityOptions contains the optional parameters for the DatabasesClient.CheckNameAvailability
//     method.
func (client *DatabasesClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest, options *DatabasesClientCheckNameAvailabilityOptions) (DatabasesClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "DatabasesClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, clusterName, resourceName, options)
	if err != nil {
		return DatabasesClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabasesClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabasesClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *DatabasesClient) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest, options *DatabasesClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/checkNameAvailability"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resourceName); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *DatabasesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (DatabasesClientCheckNameAvailabilityResponse, error) {
	result := DatabasesClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameResult); err != nil {
		return DatabasesClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Creates or updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - parameters - The database parameters supplied to the CreateOrUpdate operation.
//   - options - DatabasesClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabasesClient.BeginCreateOrUpdate
//     method.
func (client *DatabasesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DatabasesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabasesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabasesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
func (client *DatabasesClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabasesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabasesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	if options != nil && options.CallerRole != nil {
		reqQP.Set("callerRole", string(*options.CallerRole))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the database with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - options - DatabasesClientBeginDeleteOptions contains the optional parameters for the DatabasesClient.BeginDelete method.
func (client *DatabasesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*runtime.Poller[DatabasesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, databaseName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabasesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabasesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the database with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
func (client *DatabasesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabasesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
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
func (client *DatabasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns a database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - options - DatabasesClientGetOptions contains the optional parameters for the DatabasesClient.Get method.
func (client *DatabasesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientGetOptions) (DatabasesClientGetResponse, error) {
	var err error
	const operationName = "DatabasesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabasesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabasesClient) getHandleResponse(resp *http.Response) (DatabasesClientGetResponse, error) {
	result := DatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return DatabasesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByClusterPager - Returns the list of databases of the given Kusto cluster.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - options - DatabasesClientListByClusterOptions contains the optional parameters for the DatabasesClient.NewListByClusterPager
//     method.
func (client *DatabasesClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *DatabasesClientListByClusterOptions) *runtime.Pager[DatabasesClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabasesClientListByClusterResponse]{
		More: func(page DatabasesClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatabasesClientListByClusterResponse) (DatabasesClientListByClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DatabasesClient.NewListByClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			}, nil)
			if err != nil {
				return DatabasesClientListByClusterResponse{}, err
			}
			return client.listByClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *DatabasesClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *DatabasesClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *DatabasesClient) listByClusterHandleResponse(resp *http.Response) (DatabasesClientListByClusterResponse, error) {
	result := DatabasesClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseListResult); err != nil {
		return DatabasesClientListByClusterResponse{}, err
	}
	return result, nil
}

// NewListPrincipalsPager - Returns a list of database principals of the given Kusto cluster and database.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - options - DatabasesClientListPrincipalsOptions contains the optional parameters for the DatabasesClient.NewListPrincipalsPager
//     method.
func (client *DatabasesClient) NewListPrincipalsPager(resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientListPrincipalsOptions) *runtime.Pager[DatabasesClientListPrincipalsResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabasesClientListPrincipalsResponse]{
		More: func(page DatabasesClientListPrincipalsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *DatabasesClientListPrincipalsResponse) (DatabasesClientListPrincipalsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DatabasesClient.NewListPrincipalsPager")
			req, err := client.listPrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
			if err != nil {
				return DatabasesClientListPrincipalsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DatabasesClientListPrincipalsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatabasesClientListPrincipalsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listPrincipalsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listPrincipalsCreateRequest creates the ListPrincipals request.
func (client *DatabasesClient) listPrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientListPrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/listPrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listPrincipalsHandleResponse handles the ListPrincipals response.
func (client *DatabasesClient) listPrincipalsHandleResponse(resp *http.Response) (DatabasesClientListPrincipalsResponse, error) {
	result := DatabasesClientListPrincipalsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesClientListPrincipalsResponse{}, err
	}
	return result, nil
}

// RemovePrincipals - Remove Database principals permissions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - databasePrincipalsToRemove - List of database principals to remove.
//   - options - DatabasesClientRemovePrincipalsOptions contains the optional parameters for the DatabasesClient.RemovePrincipals
//     method.
func (client *DatabasesClient) RemovePrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest, options *DatabasesClientRemovePrincipalsOptions) (DatabasesClientRemovePrincipalsResponse, error) {
	var err error
	const operationName = "DatabasesClient.RemovePrincipals"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.removePrincipalsCreateRequest(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToRemove, options)
	if err != nil {
		return DatabasesClientRemovePrincipalsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabasesClientRemovePrincipalsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabasesClientRemovePrincipalsResponse{}, err
	}
	resp, err := client.removePrincipalsHandleResponse(httpResp)
	return resp, err
}

// removePrincipalsCreateRequest creates the RemovePrincipals request.
func (client *DatabasesClient) removePrincipalsCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest, options *DatabasesClientRemovePrincipalsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/removePrincipals"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, databasePrincipalsToRemove); err != nil {
		return nil, err
	}
	return req, nil
}

// removePrincipalsHandleResponse handles the RemovePrincipals response.
func (client *DatabasesClient) removePrincipalsHandleResponse(resp *http.Response) (DatabasesClientRemovePrincipalsResponse, error) {
	result := DatabasesClientRemovePrincipalsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabasePrincipalListResult); err != nil {
		return DatabasesClientRemovePrincipalsResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - parameters - The database parameters supplied to the Update operation.
//   - options - DatabasesClientBeginUpdateOptions contains the optional parameters for the DatabasesClient.BeginUpdate method.
func (client *DatabasesClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginUpdateOptions) (*runtime.Poller[DatabasesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabasesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabasesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates a database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-15
func (client *DatabasesClient) update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabasesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *DatabasesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseClassification, options *DatabasesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-15")
	if options != nil && options.CallerRole != nil {
		reqQP.Set("callerRole", string(*options.CallerRole))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}