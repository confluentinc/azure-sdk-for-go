//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageimportexport

// BitLockerKeysClientListResponse contains the response from method BitLockerKeysClient.NewListPager.
type BitLockerKeysClientListResponse struct {
	// GetBitLockerKeys response
	GetBitLockerKeysResponse
}

// JobsClientCreateResponse contains the response from method JobsClient.Create.
type JobsClientCreateResponse struct {
	// Contains the job information.
	JobResponse
}

// JobsClientDeleteResponse contains the response from method JobsClient.Delete.
type JobsClientDeleteResponse struct {
	// placeholder for future response values
}

// JobsClientGetResponse contains the response from method JobsClient.Get.
type JobsClientGetResponse struct {
	// Contains the job information.
	JobResponse
}

// JobsClientListByResourceGroupResponse contains the response from method JobsClient.NewListByResourceGroupPager.
type JobsClientListByResourceGroupResponse struct {
	// List jobs response
	ListJobsResponse
}

// JobsClientListBySubscriptionResponse contains the response from method JobsClient.NewListBySubscriptionPager.
type JobsClientListBySubscriptionResponse struct {
	// List jobs response
	ListJobsResponse
}

// JobsClientUpdateResponse contains the response from method JobsClient.Update.
type JobsClientUpdateResponse struct {
	// Contains the job information.
	JobResponse
}

// LocationsClientGetResponse contains the response from method LocationsClient.Get.
type LocationsClientGetResponse struct {
	// Provides information about an Azure data center location.
	Location
}

// LocationsClientListResponse contains the response from method LocationsClient.NewListPager.
type LocationsClientListResponse struct {
	// Locations response
	LocationsResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// List operations response
	ListOperationsResponse
}