//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridcontainerservice

// AgentPoolClientCreateOrUpdateResponse contains the response from method AgentPoolClient.BeginCreateOrUpdate.
type AgentPoolClientCreateOrUpdateResponse struct {
	// The agentPool resource definition
	AgentPool
}

// AgentPoolClientDeleteResponse contains the response from method AgentPoolClient.BeginDelete.
type AgentPoolClientDeleteResponse struct {
	// placeholder for future response values
}

// AgentPoolClientGetResponse contains the response from method AgentPoolClient.Get.
type AgentPoolClientGetResponse struct {
	// The agentPool resource definition
	AgentPool
}

// AgentPoolClientListByProvisionedClusterResponse contains the response from method AgentPoolClient.NewListByProvisionedClusterPager.
type AgentPoolClientListByProvisionedClusterResponse struct {
	// List of all agent pool resources associated with the provisioned cluster.
	AgentPoolListResult
}

// ClientDeleteKubernetesVersionsResponse contains the response from method Client.BeginDeleteKubernetesVersions.
type ClientDeleteKubernetesVersionsResponse struct {
	// placeholder for future response values
}

// ClientDeleteVMSKUsResponse contains the response from method Client.BeginDeleteVMSKUs.
type ClientDeleteVMSKUsResponse struct {
	// placeholder for future response values
}

// ClientGetKubernetesVersionsResponse contains the response from method Client.GetKubernetesVersions.
type ClientGetKubernetesVersionsResponse struct {
	// The supported kubernetes versions.
	KubernetesVersionProfile
}

// ClientGetVMSKUsResponse contains the response from method Client.GetVMSKUs.
type ClientGetVMSKUsResponse struct {
	// The list of supported VM SKUs.
	VMSKUProfile
}

// ClientPutKubernetesVersionsResponse contains the response from method Client.BeginPutKubernetesVersions.
type ClientPutKubernetesVersionsResponse struct {
	// The supported kubernetes versions.
	KubernetesVersionProfile
}

// ClientPutVMSKUsResponse contains the response from method Client.BeginPutVMSKUs.
type ClientPutVMSKUsResponse struct {
	// The list of supported VM SKUs.
	VMSKUProfile
}

// HybridIdentityMetadataClientDeleteResponse contains the response from method HybridIdentityMetadataClient.BeginDelete.
type HybridIdentityMetadataClientDeleteResponse struct {
	// placeholder for future response values
}

// HybridIdentityMetadataClientGetResponse contains the response from method HybridIdentityMetadataClient.Get.
type HybridIdentityMetadataClientGetResponse struct {
	// Defines the hybridIdentityMetadata.
	HybridIdentityMetadata
}

// HybridIdentityMetadataClientListByClusterResponse contains the response from method HybridIdentityMetadataClient.NewListByClusterPager.
type HybridIdentityMetadataClientListByClusterResponse struct {
	// List of hybridIdentityMetadata.
	HybridIdentityMetadataList
}

// HybridIdentityMetadataClientPutResponse contains the response from method HybridIdentityMetadataClient.Put.
type HybridIdentityMetadataClientPutResponse struct {
	// Defines the hybridIdentityMetadata.
	HybridIdentityMetadata
}

// KubernetesVersionsClientListResponse contains the response from method KubernetesVersionsClient.NewListPager.
type KubernetesVersionsClientListResponse struct {
	// List of supported kubernetes versions.
	KubernetesVersionProfileList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// ProvisionedClusterInstancesClientCreateOrUpdateResponse contains the response from method ProvisionedClusterInstancesClient.BeginCreateOrUpdate.
type ProvisionedClusterInstancesClientCreateOrUpdateResponse struct {
	// The provisioned cluster resource definition.
	ProvisionedCluster
}

// ProvisionedClusterInstancesClientDeleteResponse contains the response from method ProvisionedClusterInstancesClient.BeginDelete.
type ProvisionedClusterInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProvisionedClusterInstancesClientGetResponse contains the response from method ProvisionedClusterInstancesClient.Get.
type ProvisionedClusterInstancesClientGetResponse struct {
	// The provisioned cluster resource definition.
	ProvisionedCluster
}

// ProvisionedClusterInstancesClientGetUpgradeProfileResponse contains the response from method ProvisionedClusterInstancesClient.GetUpgradeProfile.
type ProvisionedClusterInstancesClientGetUpgradeProfileResponse struct {
	// The list of available kubernetes version upgrades for the provisioned cluster.
	ProvisionedClusterUpgradeProfile
}

// ProvisionedClusterInstancesClientListAdminKubeconfigResponse contains the response from method ProvisionedClusterInstancesClient.BeginListAdminKubeconfig.
type ProvisionedClusterInstancesClientListAdminKubeconfigResponse struct {
	// The list kubeconfig result response.
	ListCredentialResponse
}

// ProvisionedClusterInstancesClientListResponse contains the response from method ProvisionedClusterInstancesClient.NewListPager.
type ProvisionedClusterInstancesClientListResponse struct {
	// Lists the ProvisionedClusterInstance resource associated with the ConnectedCluster.
	ProvisionedClusterListResult
}

// ProvisionedClusterInstancesClientListUserKubeconfigResponse contains the response from method ProvisionedClusterInstancesClient.BeginListUserKubeconfig.
type ProvisionedClusterInstancesClientListUserKubeconfigResponse struct {
	// The list kubeconfig result response.
	ListCredentialResponse
}

// VMSKUsClientListResponse contains the response from method VMSKUsClient.NewListPager.
type VMSKUsClientListResponse struct {
	// The list of supported VM SKUs.
	VMSKUProfileList
}

// VirtualNetworksClientCreateOrUpdateResponse contains the response from method VirtualNetworksClient.BeginCreateOrUpdate.
type VirtualNetworksClientCreateOrUpdateResponse struct {
	// The Virtual Network resource definition.
	VirtualNetwork
}

// VirtualNetworksClientDeleteResponse contains the response from method VirtualNetworksClient.BeginDelete.
type VirtualNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworksClientListByResourceGroupResponse contains the response from method VirtualNetworksClient.NewListByResourceGroupPager.
type VirtualNetworksClientListByResourceGroupResponse struct {
	// A list of virtual network resources.
	VirtualNetworksListResult
}

// VirtualNetworksClientListBySubscriptionResponse contains the response from method VirtualNetworksClient.NewListBySubscriptionPager.
type VirtualNetworksClientListBySubscriptionResponse struct {
	// A list of virtual network resources.
	VirtualNetworksListResult
}

// VirtualNetworksClientRetrieveResponse contains the response from method VirtualNetworksClient.Retrieve.
type VirtualNetworksClientRetrieveResponse struct {
	// The Virtual Network resource definition.
	VirtualNetwork
}

// VirtualNetworksClientUpdateResponse contains the response from method VirtualNetworksClient.BeginUpdate.
type VirtualNetworksClientUpdateResponse struct {
	// The Virtual Network resource definition.
	VirtualNetwork
}