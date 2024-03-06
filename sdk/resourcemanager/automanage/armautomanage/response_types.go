//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomanage

// BestPracticesClientGetResponse contains the response from method BestPracticesClient.Get.
type BestPracticesClientGetResponse struct {
	// Definition of the Automanage best practice.
	BestPractice
}

// BestPracticesClientListByTenantResponse contains the response from method BestPracticesClient.NewListByTenantPager.
type BestPracticesClientListByTenantResponse struct {
	// The response of the list best practice operation.
	BestPracticeList
}

// BestPracticesVersionsClientGetResponse contains the response from method BestPracticesVersionsClient.Get.
type BestPracticesVersionsClientGetResponse struct {
	// Definition of the Automanage best practice.
	BestPractice
}

// BestPracticesVersionsClientListByTenantResponse contains the response from method BestPracticesVersionsClient.NewListByTenantPager.
type BestPracticesVersionsClientListByTenantResponse struct {
	// The response of the list best practice operation.
	BestPracticeList
}

// ConfigurationProfileAssignmentsClientCreateOrUpdateResponse contains the response from method ConfigurationProfileAssignmentsClient.CreateOrUpdate.
type ConfigurationProfileAssignmentsClientCreateOrUpdateResponse struct {
	// Configuration profile assignment is an association between a VM and automanage profile configuration.
	ConfigurationProfileAssignment
}

// ConfigurationProfileAssignmentsClientDeleteResponse contains the response from method ConfigurationProfileAssignmentsClient.Delete.
type ConfigurationProfileAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConfigurationProfileAssignmentsClientGetResponse contains the response from method ConfigurationProfileAssignmentsClient.Get.
type ConfigurationProfileAssignmentsClientGetResponse struct {
	// Configuration profile assignment is an association between a VM and automanage profile configuration.
	ConfigurationProfileAssignment
}

// ConfigurationProfileAssignmentsClientListByClusterNameResponse contains the response from method ConfigurationProfileAssignmentsClient.NewListByClusterNamePager.
type ConfigurationProfileAssignmentsClientListByClusterNameResponse struct {
	// The response of the list configuration profile assignment operation.
	ConfigurationProfileAssignmentList
}

// ConfigurationProfileAssignmentsClientListByMachineNameResponse contains the response from method ConfigurationProfileAssignmentsClient.NewListByMachineNamePager.
type ConfigurationProfileAssignmentsClientListByMachineNameResponse struct {
	// The response of the list configuration profile assignment operation.
	ConfigurationProfileAssignmentList
}

// ConfigurationProfileAssignmentsClientListBySubscriptionResponse contains the response from method ConfigurationProfileAssignmentsClient.NewListBySubscriptionPager.
type ConfigurationProfileAssignmentsClientListBySubscriptionResponse struct {
	// The response of the list configuration profile assignment operation.
	ConfigurationProfileAssignmentList
}

// ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse contains the response from method ConfigurationProfileAssignmentsClient.NewListByVirtualMachinesPager.
type ConfigurationProfileAssignmentsClientListByVirtualMachinesResponse struct {
	// The response of the list configuration profile assignment operation.
	ConfigurationProfileAssignmentList
}

// ConfigurationProfileAssignmentsClientListResponse contains the response from method ConfigurationProfileAssignmentsClient.NewListPager.
type ConfigurationProfileAssignmentsClientListResponse struct {
	// The response of the list configuration profile assignment operation.
	ConfigurationProfileAssignmentList
}

// ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse contains the response from method ConfigurationProfileHCIAssignmentsClient.CreateOrUpdate.
type ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse struct {
	// Configuration profile assignment is an association between a VM and automanage profile configuration.
	ConfigurationProfileAssignment
}

// ConfigurationProfileHCIAssignmentsClientDeleteResponse contains the response from method ConfigurationProfileHCIAssignmentsClient.Delete.
type ConfigurationProfileHCIAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConfigurationProfileHCIAssignmentsClientGetResponse contains the response from method ConfigurationProfileHCIAssignmentsClient.Get.
type ConfigurationProfileHCIAssignmentsClientGetResponse struct {
	// Configuration profile assignment is an association between a VM and automanage profile configuration.
	ConfigurationProfileAssignment
}

// ConfigurationProfileHCRPAssignmentsClientCreateOrUpdateResponse contains the response from method ConfigurationProfileHCRPAssignmentsClient.CreateOrUpdate.
type ConfigurationProfileHCRPAssignmentsClientCreateOrUpdateResponse struct {
	// Configuration profile assignment is an association between a VM and automanage profile configuration.
	ConfigurationProfileAssignment
}

// ConfigurationProfileHCRPAssignmentsClientDeleteResponse contains the response from method ConfigurationProfileHCRPAssignmentsClient.Delete.
type ConfigurationProfileHCRPAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConfigurationProfileHCRPAssignmentsClientGetResponse contains the response from method ConfigurationProfileHCRPAssignmentsClient.Get.
type ConfigurationProfileHCRPAssignmentsClientGetResponse struct {
	// Configuration profile assignment is an association between a VM and automanage profile configuration.
	ConfigurationProfileAssignment
}

// ConfigurationProfilesClientCreateOrUpdateResponse contains the response from method ConfigurationProfilesClient.CreateOrUpdate.
type ConfigurationProfilesClientCreateOrUpdateResponse struct {
	// Definition of the configuration profile.
	ConfigurationProfile
}

// ConfigurationProfilesClientDeleteResponse contains the response from method ConfigurationProfilesClient.Delete.
type ConfigurationProfilesClientDeleteResponse struct {
	// placeholder for future response values
}

// ConfigurationProfilesClientGetResponse contains the response from method ConfigurationProfilesClient.Get.
type ConfigurationProfilesClientGetResponse struct {
	// Definition of the configuration profile.
	ConfigurationProfile
}

// ConfigurationProfilesClientListByResourceGroupResponse contains the response from method ConfigurationProfilesClient.NewListByResourceGroupPager.
type ConfigurationProfilesClientListByResourceGroupResponse struct {
	// The response of the list configuration profile operation.
	ConfigurationProfileList
}

// ConfigurationProfilesClientListBySubscriptionResponse contains the response from method ConfigurationProfilesClient.NewListBySubscriptionPager.
type ConfigurationProfilesClientListBySubscriptionResponse struct {
	// The response of the list configuration profile operation.
	ConfigurationProfileList
}

// ConfigurationProfilesClientUpdateResponse contains the response from method ConfigurationProfilesClient.Update.
type ConfigurationProfilesClientUpdateResponse struct {
	// Definition of the configuration profile.
	ConfigurationProfile
}

// ConfigurationProfilesVersionsClientCreateOrUpdateResponse contains the response from method ConfigurationProfilesVersionsClient.CreateOrUpdate.
type ConfigurationProfilesVersionsClientCreateOrUpdateResponse struct {
	// Definition of the configuration profile.
	ConfigurationProfile
}

// ConfigurationProfilesVersionsClientDeleteResponse contains the response from method ConfigurationProfilesVersionsClient.Delete.
type ConfigurationProfilesVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ConfigurationProfilesVersionsClientGetResponse contains the response from method ConfigurationProfilesVersionsClient.Get.
type ConfigurationProfilesVersionsClientGetResponse struct {
	// Definition of the configuration profile.
	ConfigurationProfile
}

// ConfigurationProfilesVersionsClientListChildResourcesResponse contains the response from method ConfigurationProfilesVersionsClient.NewListChildResourcesPager.
type ConfigurationProfilesVersionsClientListChildResourcesResponse struct {
	// The response of the list configuration profile operation.
	ConfigurationProfileList
}

// HCIReportsClientGetResponse contains the response from method HCIReportsClient.Get.
type HCIReportsClientGetResponse struct {
	// Definition of the report.
	Report
}

// HCIReportsClientListByConfigurationProfileAssignmentsResponse contains the response from method HCIReportsClient.NewListByConfigurationProfileAssignmentsPager.
type HCIReportsClientListByConfigurationProfileAssignmentsResponse struct {
	// The response of the list report operation.
	ReportList
}

// HCRPReportsClientGetResponse contains the response from method HCRPReportsClient.Get.
type HCRPReportsClientGetResponse struct {
	// Definition of the report.
	Report
}

// HCRPReportsClientListByConfigurationProfileAssignmentsResponse contains the response from method HCRPReportsClient.NewListByConfigurationProfileAssignmentsPager.
type HCRPReportsClientListByConfigurationProfileAssignmentsResponse struct {
	// The response of the list report operation.
	ReportList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// ReportsClientGetResponse contains the response from method ReportsClient.Get.
type ReportsClientGetResponse struct {
	// Definition of the report.
	Report
}

// ReportsClientListByConfigurationProfileAssignmentsResponse contains the response from method ReportsClient.NewListByConfigurationProfileAssignmentsPager.
type ReportsClientListByConfigurationProfileAssignmentsResponse struct {
	// The response of the list report operation.
	ReportList
}

// ServicePrincipalsClientGetResponse contains the response from method ServicePrincipalsClient.Get.
type ServicePrincipalsClientGetResponse struct {
	// The Service Principal Id for the subscription.
	ServicePrincipal
}

// ServicePrincipalsClientListBySubscriptionResponse contains the response from method ServicePrincipalsClient.NewListBySubscriptionPager.
type ServicePrincipalsClientListBySubscriptionResponse struct {
	// The list of ServicePrincipals.
	ServicePrincipalListResult
}
