//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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

// DisksClient contains the methods for the Disks group.
// Don't use this type directly, use NewDisksClient() instead.
type DisksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDisksClient creates a new instance of DisksClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDisksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DisksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DisksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginAttach - Attach and create the lease of the disk to the virtual machine. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the disk.
//   - attachDiskProperties - Properties of the disk to attach.
//   - options - DisksClientBeginAttachOptions contains the optional parameters for the DisksClient.BeginAttach method.
func (client *DisksClient) BeginAttach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties, options *DisksClientBeginAttachOptions) (*runtime.Poller[DisksClientAttachResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.attach(ctx, resourceGroupName, labName, userName, name, attachDiskProperties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DisksClientAttachResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DisksClientAttachResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Attach - Attach and create the lease of the disk to the virtual machine. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *DisksClient) attach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties, options *DisksClientBeginAttachOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginAttach"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.attachCreateRequest(ctx, resourceGroupName, labName, userName, name, attachDiskProperties, options)
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

// attachCreateRequest creates the Attach request.
func (client *DisksClient) attachCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties AttachDiskProperties, options *DisksClientBeginAttachOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}/attach"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, attachDiskProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - Create or replace an existing disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the disk.
//   - disk - A Disk.
//   - options - DisksClientBeginCreateOrUpdateOptions contains the optional parameters for the DisksClient.BeginCreateOrUpdate
//     method.
func (client *DisksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*runtime.Poller[DisksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, userName, name, disk, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DisksClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DisksClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or replace an existing disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *DisksClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, userName, name, disk, options)
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
func (client *DisksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk Disk, options *DisksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, disk); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the disk.
//   - options - DisksClientBeginDeleteOptions contains the optional parameters for the DisksClient.BeginDelete method.
func (client *DisksClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientBeginDeleteOptions) (*runtime.Poller[DisksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, userName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DisksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DisksClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete disk. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *DisksClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
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
func (client *DisksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDetach - Detach and break the lease of the disk attached to the virtual machine. This operation can take a while to
// complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the disk.
//   - detachDiskProperties - Properties of the disk to detach.
//   - options - DisksClientBeginDetachOptions contains the optional parameters for the DisksClient.BeginDetach method.
func (client *DisksClient) BeginDetach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties, options *DisksClientBeginDetachOptions) (*runtime.Poller[DisksClientDetachResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.detach(ctx, resourceGroupName, labName, userName, name, detachDiskProperties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DisksClientDetachResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DisksClientDetachResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Detach - Detach and break the lease of the disk attached to the virtual machine. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *DisksClient) detach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties, options *DisksClientBeginDetachOptions) (*http.Response, error) {
	var err error
	const operationName = "DisksClient.BeginDetach"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.detachCreateRequest(ctx, resourceGroupName, labName, userName, name, detachDiskProperties, options)
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

// detachCreateRequest creates the Detach request.
func (client *DisksClient) detachCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties DetachDiskProperties, options *DisksClientBeginDetachOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}/detach"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, detachDiskProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get disk.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the disk.
//   - options - DisksClientGetOptions contains the optional parameters for the DisksClient.Get method.
func (client *DisksClient) Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientGetOptions) (DisksClientGetResponse, error) {
	var err error
	const operationName = "DisksClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, userName, name, options)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DisksClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *DisksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DisksClient) getHandleResponse(resp *http.Response) (DisksClientGetResponse, error) {
	result := DisksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Disk); err != nil {
		return DisksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List disks in a given user profile.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - options - DisksClientListOptions contains the optional parameters for the DisksClient.NewListPager method.
func (client *DisksClient) NewListPager(resourceGroupName string, labName string, userName string, options *DisksClientListOptions) *runtime.Pager[DisksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DisksClientListResponse]{
		More: func(page DisksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DisksClientListResponse) (DisksClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DisksClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, labName, userName, options)
			}, nil)
			if err != nil {
				return DisksClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DisksClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, options *DisksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DisksClient) listHandleResponse(resp *http.Response) (DisksClientListResponse, error) {
	result := DisksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiskList); err != nil {
		return DisksClientListResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of disks. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - userName - The name of the user profile.
//   - name - The name of the disk.
//   - disk - A Disk.
//   - options - DisksClientUpdateOptions contains the optional parameters for the DisksClient.Update method.
func (client *DisksClient) Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk DiskFragment, options *DisksClientUpdateOptions) (DisksClientUpdateResponse, error) {
	var err error
	const operationName = "DisksClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, userName, name, disk, options)
	if err != nil {
		return DisksClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DisksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DisksClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DisksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk DiskFragment, options *DisksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{userName}/disks/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if userName == "" {
		return nil, errors.New("parameter userName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userName}", url.PathEscape(userName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, disk); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DisksClient) updateHandleResponse(resp *http.Response) (DisksClientUpdateResponse, error) {
	result := DisksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Disk); err != nil {
		return DisksClientUpdateResponse{}, err
	}
	return result, nil
}