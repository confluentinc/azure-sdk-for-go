//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	// Data Lake Analytics account name availability result information.
	NameAvailabilityInformation
}

// AccountsClientCreateResponse contains the response from method AccountsClient.BeginCreate.
type AccountsClientCreateResponse struct {
	// A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	// A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	// Data Lake Analytics account list information.
	AccountListResult
}

// AccountsClientListResponse contains the response from method AccountsClient.NewListPager.
type AccountsClientListResponse struct {
	// Data Lake Analytics account list information.
	AccountListResult
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	// A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
	Account
}

// ComputePoliciesClientCreateOrUpdateResponse contains the response from method ComputePoliciesClient.CreateOrUpdate.
type ComputePoliciesClientCreateOrUpdateResponse struct {
	// Data Lake Analytics compute policy information.
	ComputePolicy
}

// ComputePoliciesClientDeleteResponse contains the response from method ComputePoliciesClient.Delete.
type ComputePoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ComputePoliciesClientGetResponse contains the response from method ComputePoliciesClient.Get.
type ComputePoliciesClientGetResponse struct {
	// Data Lake Analytics compute policy information.
	ComputePolicy
}

// ComputePoliciesClientListByAccountResponse contains the response from method ComputePoliciesClient.NewListByAccountPager.
type ComputePoliciesClientListByAccountResponse struct {
	// The list of compute policies in the account.
	ComputePolicyListResult
}

// ComputePoliciesClientUpdateResponse contains the response from method ComputePoliciesClient.Update.
type ComputePoliciesClientUpdateResponse struct {
	// Data Lake Analytics compute policy information.
	ComputePolicy
}

// DataLakeStoreAccountsClientAddResponse contains the response from method DataLakeStoreAccountsClient.Add.
type DataLakeStoreAccountsClientAddResponse struct {
	// placeholder for future response values
}

// DataLakeStoreAccountsClientDeleteResponse contains the response from method DataLakeStoreAccountsClient.Delete.
type DataLakeStoreAccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataLakeStoreAccountsClientGetResponse contains the response from method DataLakeStoreAccountsClient.Get.
type DataLakeStoreAccountsClientGetResponse struct {
	// Data Lake Store account information.
	DataLakeStoreAccountInformation
}

// DataLakeStoreAccountsClientListByAccountResponse contains the response from method DataLakeStoreAccountsClient.NewListByAccountPager.
type DataLakeStoreAccountsClientListByAccountResponse struct {
	// Data Lake Store account list information.
	DataLakeStoreAccountInformationListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.CreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	// Data Lake Analytics firewall rule information.
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.Delete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	// Data Lake Analytics firewall rule information.
	FirewallRule
}

// FirewallRulesClientListByAccountResponse contains the response from method FirewallRulesClient.NewListByAccountPager.
type FirewallRulesClientListByAccountResponse struct {
	// Data Lake Analytics firewall rule list information.
	FirewallRuleListResult
}

// FirewallRulesClientUpdateResponse contains the response from method FirewallRulesClient.Update.
type FirewallRulesClientUpdateResponse struct {
	// Data Lake Analytics firewall rule information.
	FirewallRule
}

// LocationsClientGetCapabilityResponse contains the response from method LocationsClient.GetCapability.
type LocationsClientGetCapabilityResponse struct {
	// Subscription-level properties and limits for Data Lake Analytics.
	CapabilityInformation
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// The list of available operations for Data Lake Analytics.
	OperationListResult
}

// StorageAccountsClientAddResponse contains the response from method StorageAccountsClient.Add.
type StorageAccountsClientAddResponse struct {
	// placeholder for future response values
}

// StorageAccountsClientDeleteResponse contains the response from method StorageAccountsClient.Delete.
type StorageAccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageAccountsClientGetResponse contains the response from method StorageAccountsClient.Get.
type StorageAccountsClientGetResponse struct {
	// Azure Storage account information.
	StorageAccountInformation
}

// StorageAccountsClientGetStorageContainerResponse contains the response from method StorageAccountsClient.GetStorageContainer.
type StorageAccountsClientGetStorageContainerResponse struct {
	// Azure Storage blob container information.
	StorageContainer
}

// StorageAccountsClientListByAccountResponse contains the response from method StorageAccountsClient.NewListByAccountPager.
type StorageAccountsClientListByAccountResponse struct {
	// Azure Storage account list information.
	StorageAccountInformationListResult
}

// StorageAccountsClientListSasTokensResponse contains the response from method StorageAccountsClient.NewListSasTokensPager.
type StorageAccountsClientListSasTokensResponse struct {
	// The SAS response that contains the storage account, container and associated SAS token for connection use.
	SasTokenInformationListResult
}

// StorageAccountsClientListStorageContainersResponse contains the response from method StorageAccountsClient.NewListStorageContainersPager.
type StorageAccountsClientListStorageContainersResponse struct {
	// The list of blob containers associated with the storage account attached to the Data Lake Analytics account.
	StorageContainerListResult
}

// StorageAccountsClientUpdateResponse contains the response from method StorageAccountsClient.Update.
type StorageAccountsClientUpdateResponse struct {
	// placeholder for future response values
}