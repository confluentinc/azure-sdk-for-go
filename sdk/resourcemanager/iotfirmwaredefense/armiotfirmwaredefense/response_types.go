//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotfirmwaredefense

// FirmwareClientCreateResponse contains the response from method FirmwareClient.Create.
type FirmwareClientCreateResponse struct {
	// Firmware definition
	Firmware
}

// FirmwareClientDeleteResponse contains the response from method FirmwareClient.Delete.
type FirmwareClientDeleteResponse struct {
	// placeholder for future response values
}

// FirmwareClientGenerateBinaryHardeningDetailsResponse contains the response from method FirmwareClient.GenerateBinaryHardeningDetails.
type FirmwareClientGenerateBinaryHardeningDetailsResponse struct {
	// Binary hardening of a firmware.
	BinaryHardening
}

// FirmwareClientGenerateBinaryHardeningSummaryResponse contains the response from method FirmwareClient.GenerateBinaryHardeningSummary.
type FirmwareClientGenerateBinaryHardeningSummaryResponse struct {
	// Binary hardening summary percentages.
	BinaryHardeningSummary
}

// FirmwareClientGenerateComponentDetailsResponse contains the response from method FirmwareClient.GenerateComponentDetails.
type FirmwareClientGenerateComponentDetailsResponse struct {
	// Component of a firmware.
	Component
}

// FirmwareClientGenerateCryptoCertificateSummaryResponse contains the response from method FirmwareClient.GenerateCryptoCertificateSummary.
type FirmwareClientGenerateCryptoCertificateSummaryResponse struct {
	// Cryptographic certificate summary values.
	CryptoCertificateSummary
}

// FirmwareClientGenerateCryptoKeySummaryResponse contains the response from method FirmwareClient.GenerateCryptoKeySummary.
type FirmwareClientGenerateCryptoKeySummaryResponse struct {
	// Cryptographic key summary values.
	CryptoKeySummary
}

// FirmwareClientGenerateCveSummaryResponse contains the response from method FirmwareClient.GenerateCveSummary.
type FirmwareClientGenerateCveSummaryResponse struct {
	// CVE summary values.
	CveSummary
}

// FirmwareClientGenerateDownloadURLResponse contains the response from method FirmwareClient.GenerateDownloadURL.
type FirmwareClientGenerateDownloadURLResponse struct {
	// Url data for creating or accessing a blob file.
	URLToken
}

// FirmwareClientGenerateFilesystemDownloadURLResponse contains the response from method FirmwareClient.GenerateFilesystemDownloadURL.
type FirmwareClientGenerateFilesystemDownloadURLResponse struct {
	// Url data for creating or accessing a blob file.
	URLToken
}

// FirmwareClientGenerateSummaryResponse contains the response from method FirmwareClient.GenerateSummary.
type FirmwareClientGenerateSummaryResponse struct {
	// Summary result after scanning the firmware.
	FirmwareSummary
}

// FirmwareClientGetResponse contains the response from method FirmwareClient.Get.
type FirmwareClientGetResponse struct {
	// Firmware definition
	Firmware
}

// FirmwareClientListByWorkspaceResponse contains the response from method FirmwareClient.NewListByWorkspacePager.
type FirmwareClientListByWorkspaceResponse struct {
	// List of firmwares
	FirmwareList
}

// FirmwareClientListGenerateBinaryHardeningListResponse contains the response from method FirmwareClient.NewListGenerateBinaryHardeningListPager.
type FirmwareClientListGenerateBinaryHardeningListResponse struct {
	// List result for binary hardening
	BinaryHardeningList
}

// FirmwareClientListGenerateComponentListResponse contains the response from method FirmwareClient.NewListGenerateComponentListPager.
type FirmwareClientListGenerateComponentListResponse struct {
	// List result for components
	ComponentList
}

// FirmwareClientListGenerateCryptoCertificateListResponse contains the response from method FirmwareClient.NewListGenerateCryptoCertificateListPager.
type FirmwareClientListGenerateCryptoCertificateListResponse struct {
	// Crypto certificates list
	CryptoCertificateList
}

// FirmwareClientListGenerateCryptoKeyListResponse contains the response from method FirmwareClient.NewListGenerateCryptoKeyListPager.
type FirmwareClientListGenerateCryptoKeyListResponse struct {
	// Crypto keys list
	CryptoKeyList
}

// FirmwareClientListGenerateCveListResponse contains the response from method FirmwareClient.NewListGenerateCveListPager.
type FirmwareClientListGenerateCveListResponse struct {
	// List result for CVE
	CveList
}

// FirmwareClientListGeneratePasswordHashListResponse contains the response from method FirmwareClient.NewListGeneratePasswordHashListPager.
type FirmwareClientListGeneratePasswordHashListResponse struct {
	// Password hashes list
	PasswordHashList
}

// FirmwareClientUpdateResponse contains the response from method FirmwareClient.Update.
type FirmwareClientUpdateResponse struct {
	// Firmware definition
	Firmware
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// WorkspacesClientCreateResponse contains the response from method WorkspacesClient.Create.
type WorkspacesClientCreateResponse struct {
	// Firmware analysis workspace.
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.Delete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGenerateUploadURLResponse contains the response from method WorkspacesClient.GenerateUploadURL.
type WorkspacesClientGenerateUploadURLResponse struct {
	// Url data for creating or accessing a blob file.
	URLToken
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	// Firmware analysis workspace.
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.NewListByResourceGroupPager.
type WorkspacesClientListByResourceGroupResponse struct {
	// Return a list of firmware analysis workspaces.
	WorkspaceList
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.NewListBySubscriptionPager.
type WorkspacesClientListBySubscriptionResponse struct {
	// Return a list of firmware analysis workspaces.
	WorkspaceList
}

// WorkspacesClientUpdateResponse contains the response from method WorkspacesClient.Update.
type WorkspacesClientUpdateResponse struct {
	// Firmware analysis workspace.
	Workspace
}