//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// ProtectionPoliciesClient contains the methods for the ProtectionPolicies group.
// Don't use this type directly, use NewProtectionPoliciesClient() instead.
type ProtectionPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProtectionPoliciesClient creates a new instance of ProtectionPoliciesClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProtectionPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProtectionPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProtectionPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or modifies a backup policy. This is an asynchronous operation. Status of the operation can be
// fetched using GetPolicyOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - policyName - Backup policy to be created.
//   - parameters - resource backup policy
//   - options - ProtectionPoliciesClientCreateOrUpdateOptions contains the optional parameters for the ProtectionPoliciesClient.CreateOrUpdate
//     method.
func (client *ProtectionPoliciesClient) CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, policyName string, parameters ProtectionPolicyResource, options *ProtectionPoliciesClientCreateOrUpdateOptions) (ProtectionPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ProtectionPoliciesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, vaultName, resourceGroupName, policyName, parameters, options)
	if err != nil {
		return ProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProtectionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, policyName string, parameters ProtectionPolicyResource, options *ProtectionPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProtectionPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ProtectionPoliciesClientCreateOrUpdateResponse, error) {
	result := ProtectionPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionPolicyResource); err != nil {
		return ProtectionPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes specified backup policy from your Recovery Services Vault. This is an asynchronous operation. Status
// of the operation can be fetched using GetProtectionPolicyOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - policyName - Backup policy to be deleted.
//   - options - ProtectionPoliciesClientBeginDeleteOptions contains the optional parameters for the ProtectionPoliciesClient.BeginDelete
//     method.
func (client *ProtectionPoliciesClient) BeginDelete(ctx context.Context, vaultName string, resourceGroupName string, policyName string, options *ProtectionPoliciesClientBeginDeleteOptions) (*runtime.Poller[ProtectionPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, vaultName, resourceGroupName, policyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ProtectionPoliciesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ProtectionPoliciesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes specified backup policy from your Recovery Services Vault. This is an asynchronous operation. Status of
// the operation can be fetched using GetProtectionPolicyOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *ProtectionPoliciesClient) deleteOperation(ctx context.Context, vaultName string, resourceGroupName string, policyName string, options *ProtectionPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ProtectionPoliciesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, vaultName, resourceGroupName, policyName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProtectionPoliciesClient) deleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, policyName string, options *ProtectionPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Provides the details of the backup policies associated to Recovery Services Vault. This is an asynchronous operation.
// Status of the operation can be fetched using GetPolicyOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - policyName - Backup policy information to be fetched.
//   - options - ProtectionPoliciesClientGetOptions contains the optional parameters for the ProtectionPoliciesClient.Get method.
func (client *ProtectionPoliciesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, options *ProtectionPoliciesClientGetOptions) (ProtectionPoliciesClientGetResponse, error) {
	var err error
	const operationName = "ProtectionPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, policyName, options)
	if err != nil {
		return ProtectionPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProtectionPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProtectionPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProtectionPoliciesClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, policyName string, options *ProtectionPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies/{policyName}"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if policyName == "" {
		return nil, errors.New("parameter policyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProtectionPoliciesClient) getHandleResponse(resp *http.Response) (ProtectionPoliciesClientGetResponse, error) {
	result := ProtectionPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionPolicyResource); err != nil {
		return ProtectionPoliciesClientGetResponse{}, err
	}
	return result, nil
}