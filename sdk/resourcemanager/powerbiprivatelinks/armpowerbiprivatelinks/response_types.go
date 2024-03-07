//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PowerBIResourcesClientCreateResponse contains the response from method PowerBIResourcesClient.Create.
type PowerBIResourcesClientCreateResponse struct {
	// TenantResource
	TenantResource
}

// PowerBIResourcesClientDeleteResponse contains the response from method PowerBIResourcesClient.Delete.
type PowerBIResourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// PowerBIResourcesClientListByResourceNameResponse contains the response from method PowerBIResourcesClient.ListByResourceName.
type PowerBIResourcesClientListByResourceNameResponse struct {
	// Array of TenantResource
	TenantResourceArray []*TenantResource
}

// PowerBIResourcesClientUpdateResponse contains the response from method PowerBIResourcesClient.Update.
type PowerBIResourcesClientUpdateResponse struct {
	// TenantResource
	TenantResource
}

// PrivateEndpointConnectionsClientCreateResponse contains the response from method PrivateEndpointConnectionsClient.Create.
type PrivateEndpointConnectionsClientCreateResponse struct {
	// PrivateEndpointConnection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// PrivateEndpointConnection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByResourceResponse contains the response from method PrivateEndpointConnectionsClient.NewListByResourcePager.
type PrivateEndpointConnectionsClientListByResourceResponse struct {
	// List of private endpoint connections.
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// A private link resource
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByResourceResponse contains the response from method PrivateLinkResourcesClient.NewListByResourcePager.
type PrivateLinkResourcesClientListByResourceResponse struct {
	// Specifies list of the private link resource.
	PrivateLinkResourcesListResult
}

// PrivateLinkServiceResourceOperationResultsClientGetResponse contains the response from method PrivateLinkServiceResourceOperationResultsClient.BeginGet.
type PrivateLinkServiceResourceOperationResultsClientGetResponse struct {
	// AsyncOperationDetail
	AsyncOperationDetail
}

// PrivateLinkServicesClientListByResourceGroupResponse contains the response from method PrivateLinkServicesClient.ListByResourceGroup.
type PrivateLinkServicesClientListByResourceGroupResponse struct {
	// Array of TenantResource
	TenantResourceArray []*TenantResource
}

// PrivateLinkServicesForPowerBIClientListBySubscriptionIDResponse contains the response from method PrivateLinkServicesForPowerBIClient.ListBySubscriptionID.
type PrivateLinkServicesForPowerBIClientListBySubscriptionIDResponse struct {
	// Array of TenantResource
	TenantResourceArray []*TenantResource
}