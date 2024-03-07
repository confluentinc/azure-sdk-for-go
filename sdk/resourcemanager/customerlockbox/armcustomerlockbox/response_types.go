//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerlockbox

// GetClientTenantOptedInResponse contains the response from method GetClient.TenantOptedIn.
type GetClientTenantOptedInResponse struct {
	// TenantOptIn Response object
	TenantOptInResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Customer Lockbox operations. It contains a list of operations.
	OperationListResult
}

// PostClientDisableLockboxResponse contains the response from method PostClient.DisableLockbox.
type PostClientDisableLockboxResponse struct {
	// placeholder for future response values
}

// PostClientEnableLockboxResponse contains the response from method PostClient.EnableLockbox.
type PostClientEnableLockboxResponse struct {
	// placeholder for future response values
}

// RequestsClientGetResponse contains the response from method RequestsClient.Get.
type RequestsClientGetResponse struct {
	// A Lockbox request response object, containing all information associated with the request.
	LockboxRequestResponse
}

// RequestsClientListResponse contains the response from method RequestsClient.NewListPager.
type RequestsClientListResponse struct {
	// Object containing a list of streaming jobs.
	RequestListResult
}

// RequestsClientUpdateStatusResponse contains the response from method RequestsClient.UpdateStatus.
type RequestsClientUpdateStatusResponse struct {
	// Request content object, in the use of Approve or Deny a Lockbox request.
	Approval
}