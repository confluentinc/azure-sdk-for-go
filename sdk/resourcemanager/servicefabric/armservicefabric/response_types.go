//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabric

// ApplicationTypeVersionsClientCreateOrUpdateResponse contains the response from method ApplicationTypeVersionsClient.BeginCreateOrUpdate.
type ApplicationTypeVersionsClientCreateOrUpdateResponse struct {
	// An application type version resource for the specified application type name resource.
	ApplicationTypeVersionResource
}

// ApplicationTypeVersionsClientDeleteResponse contains the response from method ApplicationTypeVersionsClient.BeginDelete.
type ApplicationTypeVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationTypeVersionsClientGetResponse contains the response from method ApplicationTypeVersionsClient.Get.
type ApplicationTypeVersionsClientGetResponse struct {
	// An application type version resource for the specified application type name resource.
	ApplicationTypeVersionResource
}

// ApplicationTypeVersionsClientListResponse contains the response from method ApplicationTypeVersionsClient.NewListPager.
type ApplicationTypeVersionsClientListResponse struct {
	// The list of application type version resources for the specified application type name resource.
	ApplicationTypeVersionResourceList
}

// ApplicationTypesClientCreateOrUpdateResponse contains the response from method ApplicationTypesClient.CreateOrUpdate.
type ApplicationTypesClientCreateOrUpdateResponse struct {
	// The application type name resource
	ApplicationTypeResource
}

// ApplicationTypesClientDeleteResponse contains the response from method ApplicationTypesClient.BeginDelete.
type ApplicationTypesClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationTypesClientGetResponse contains the response from method ApplicationTypesClient.Get.
type ApplicationTypesClientGetResponse struct {
	// The application type name resource
	ApplicationTypeResource
}

// ApplicationTypesClientListResponse contains the response from method ApplicationTypesClient.NewListPager.
type ApplicationTypesClientListResponse struct {
	// The list of application type names.
	ApplicationTypeResourceList
}

// ApplicationsClientCreateOrUpdateResponse contains the response from method ApplicationsClient.BeginCreateOrUpdate.
type ApplicationsClientCreateOrUpdateResponse struct {
	// The application resource.
	ApplicationResource
}

// ApplicationsClientDeleteResponse contains the response from method ApplicationsClient.BeginDelete.
type ApplicationsClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationsClientGetResponse contains the response from method ApplicationsClient.Get.
type ApplicationsClientGetResponse struct {
	// The application resource.
	ApplicationResource
}

// ApplicationsClientListResponse contains the response from method ApplicationsClient.NewListPager.
type ApplicationsClientListResponse struct {
	// The list of application resources.
	ApplicationResourceList
}

// ApplicationsClientUpdateResponse contains the response from method ApplicationsClient.BeginUpdate.
type ApplicationsClientUpdateResponse struct {
	// The application resource.
	ApplicationResource
}

// ClusterVersionsClientGetByEnvironmentResponse contains the response from method ClusterVersionsClient.GetByEnvironment.
type ClusterVersionsClientGetByEnvironmentResponse struct {
	// The list results of the Service Fabric runtime versions.
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientGetResponse contains the response from method ClusterVersionsClient.Get.
type ClusterVersionsClientGetResponse struct {
	// The list results of the Service Fabric runtime versions.
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientListByEnvironmentResponse contains the response from method ClusterVersionsClient.ListByEnvironment.
type ClusterVersionsClientListByEnvironmentResponse struct {
	// The list results of the Service Fabric runtime versions.
	ClusterCodeVersionsListResult
}

// ClusterVersionsClientListResponse contains the response from method ClusterVersionsClient.List.
type ClusterVersionsClientListResponse struct {
	// The list results of the Service Fabric runtime versions.
	ClusterCodeVersionsListResult
}

// ClustersClientCreateOrUpdateResponse contains the response from method ClustersClient.BeginCreateOrUpdate.
type ClustersClientCreateOrUpdateResponse struct {
	// The cluster resource
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.Delete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	// The cluster resource
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.NewListByResourceGroupPager.
type ClustersClientListByResourceGroupResponse struct {
	// Cluster list results
	ClusterListResult
}

// ClustersClientListResponse contains the response from method ClustersClient.NewListPager.
type ClustersClientListResponse struct {
	// Cluster list results
	ClusterListResult
}

// ClustersClientListUpgradableVersionsResponse contains the response from method ClustersClient.ListUpgradableVersions.
type ClustersClientListUpgradableVersionsResponse struct {
	// The list of intermediate cluster code versions for an upgrade or downgrade. Or minimum and maximum upgradable version if
	// no target was given
	UpgradableVersionPathResult
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.BeginUpdate.
type ClustersClientUpdateResponse struct {
	// The cluster resource
	Cluster
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Describes the result of the request to list Service Fabric resource provider operations.
	OperationListResult
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// The service resource.
	ServiceResource
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.BeginDelete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// The service resource.
	ServiceResource
}

// ServicesClientListResponse contains the response from method ServicesClient.NewListPager.
type ServicesClientListResponse struct {
	// The list of service resources.
	ServiceResourceList
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.BeginUpdate.
type ServicesClientUpdateResponse struct {
	// The service resource.
	ServiceResource
}