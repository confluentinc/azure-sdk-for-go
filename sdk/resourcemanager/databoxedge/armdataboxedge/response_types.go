//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import "encoding/json"

// AddonsClientCreateOrUpdateResponse contains the response from method AddonsClient.BeginCreateOrUpdate.
type AddonsClientCreateOrUpdateResponse struct {
	// Role Addon
	AddonClassification
}

// MarshalJSON implements the json.Marshaller interface for type AddonsClientCreateOrUpdateResponse.
func (a AddonsClientCreateOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.AddonClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddonsClientCreateOrUpdateResponse.
func (a *AddonsClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAddonClassification(data)
	if err != nil {
		return err
	}
	a.AddonClassification = res
	return nil
}

// AddonsClientDeleteResponse contains the response from method AddonsClient.BeginDelete.
type AddonsClientDeleteResponse struct {
	// placeholder for future response values
}

// AddonsClientGetResponse contains the response from method AddonsClient.Get.
type AddonsClientGetResponse struct {
	// Role Addon
	AddonClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddonsClientGetResponse.
func (a *AddonsClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalAddonClassification(data)
	if err != nil {
		return err
	}
	a.AddonClassification = res
	return nil
}

// AddonsClientListByRoleResponse contains the response from method AddonsClient.NewListByRolePager.
type AddonsClientListByRoleResponse struct {
	// Collection of all the Role addon on the Azure Stack Edge device.
	AddonList
}

// AlertsClientGetResponse contains the response from method AlertsClient.Get.
type AlertsClientGetResponse struct {
	// Alert on the data box edge/gateway device.
	Alert
}

// AlertsClientListByDataBoxEdgeDeviceResponse contains the response from method AlertsClient.NewListByDataBoxEdgeDevicePager.
type AlertsClientListByDataBoxEdgeDeviceResponse struct {
	// Collection of alerts.
	AlertList
}

// AvailableSKUsClientListResponse contains the response from method AvailableSKUsClient.NewListPager.
type AvailableSKUsClientListResponse struct {
	// List of SKU Information objects.
	SKUList
}

// BandwidthSchedulesClientCreateOrUpdateResponse contains the response from method BandwidthSchedulesClient.BeginCreateOrUpdate.
type BandwidthSchedulesClientCreateOrUpdateResponse struct {
	// The bandwidth schedule details.
	BandwidthSchedule
}

// BandwidthSchedulesClientDeleteResponse contains the response from method BandwidthSchedulesClient.BeginDelete.
type BandwidthSchedulesClientDeleteResponse struct {
	// placeholder for future response values
}

// BandwidthSchedulesClientGetResponse contains the response from method BandwidthSchedulesClient.Get.
type BandwidthSchedulesClientGetResponse struct {
	// The bandwidth schedule details.
	BandwidthSchedule
}

// BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse contains the response from method BandwidthSchedulesClient.NewListByDataBoxEdgeDevicePager.
type BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse struct {
	// The collection of bandwidth schedules.
	BandwidthSchedulesList
}

// ContainersClientCreateOrUpdateResponse contains the response from method ContainersClient.BeginCreateOrUpdate.
type ContainersClientCreateOrUpdateResponse struct {
	// Represents a container on the Data Box Edge/Gateway device.
	Container
}

// ContainersClientDeleteResponse contains the response from method ContainersClient.BeginDelete.
type ContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// ContainersClientGetResponse contains the response from method ContainersClient.Get.
type ContainersClientGetResponse struct {
	// Represents a container on the Data Box Edge/Gateway device.
	Container
}

// ContainersClientListByStorageAccountResponse contains the response from method ContainersClient.NewListByStorageAccountPager.
type ContainersClientListByStorageAccountResponse struct {
	// Collection of all the containers on the Data Box Edge/Gateway device.
	ContainerList
}

// ContainersClientRefreshResponse contains the response from method ContainersClient.BeginRefresh.
type ContainersClientRefreshResponse struct {
	// placeholder for future response values
}

// DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse contains the response from method DeviceCapacityCheckClient.BeginCheckResourceCreationFeasibility.
type DeviceCapacityCheckClientCheckResourceCreationFeasibilityResponse struct {
	// placeholder for future response values
}

// DeviceCapacityInfoClientGetDeviceCapacityInfoResponse contains the response from method DeviceCapacityInfoClient.GetDeviceCapacityInfo.
type DeviceCapacityInfoClientGetDeviceCapacityInfoResponse struct {
	// Object for Capturing DeviceCapacityInfo
	DeviceCapacityInfo
}

// DevicesClientCreateOrUpdateResponse contains the response from method DevicesClient.CreateOrUpdate.
type DevicesClientCreateOrUpdateResponse struct {
	// The Data Box Edge/Gateway device.
	Device
}

// DevicesClientCreateOrUpdateSecuritySettingsResponse contains the response from method DevicesClient.BeginCreateOrUpdateSecuritySettings.
type DevicesClientCreateOrUpdateSecuritySettingsResponse struct {
	// placeholder for future response values
}

// DevicesClientDeleteResponse contains the response from method DevicesClient.BeginDelete.
type DevicesClientDeleteResponse struct {
	// placeholder for future response values
}

// DevicesClientDownloadUpdatesResponse contains the response from method DevicesClient.BeginDownloadUpdates.
type DevicesClientDownloadUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientGenerateCertificateResponse contains the response from method DevicesClient.GenerateCertificate.
type DevicesClientGenerateCertificateResponse struct {
	// Used in activation key generation flow.
	GenerateCertResponse
}

// DevicesClientGetExtendedInformationResponse contains the response from method DevicesClient.GetExtendedInformation.
type DevicesClientGetExtendedInformationResponse struct {
	// The extended Info of the Data Box Edge/Gateway device.
	DeviceExtendedInfo
}

// DevicesClientGetNetworkSettingsResponse contains the response from method DevicesClient.GetNetworkSettings.
type DevicesClientGetNetworkSettingsResponse struct {
	// The network settings of a device.
	NetworkSettings
}

// DevicesClientGetResponse contains the response from method DevicesClient.Get.
type DevicesClientGetResponse struct {
	// The Data Box Edge/Gateway device.
	Device
}

// DevicesClientGetUpdateSummaryResponse contains the response from method DevicesClient.GetUpdateSummary.
type DevicesClientGetUpdateSummaryResponse struct {
	// Details about ongoing updates and availability of updates on the device.
	UpdateSummary
}

// DevicesClientInstallUpdatesResponse contains the response from method DevicesClient.BeginInstallUpdates.
type DevicesClientInstallUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientListByResourceGroupResponse contains the response from method DevicesClient.NewListByResourceGroupPager.
type DevicesClientListByResourceGroupResponse struct {
	// The collection of Data Box Edge/Gateway devices.
	DeviceList
}

// DevicesClientListBySubscriptionResponse contains the response from method DevicesClient.NewListBySubscriptionPager.
type DevicesClientListBySubscriptionResponse struct {
	// The collection of Data Box Edge/Gateway devices.
	DeviceList
}

// DevicesClientScanForUpdatesResponse contains the response from method DevicesClient.BeginScanForUpdates.
type DevicesClientScanForUpdatesResponse struct {
	// placeholder for future response values
}

// DevicesClientUpdateExtendedInformationResponse contains the response from method DevicesClient.UpdateExtendedInformation.
type DevicesClientUpdateExtendedInformationResponse struct {
	// The extended Info of the Data Box Edge/Gateway device.
	DeviceExtendedInfo
}

// DevicesClientUpdateResponse contains the response from method DevicesClient.Update.
type DevicesClientUpdateResponse struct {
	// The Data Box Edge/Gateway device.
	Device
}

// DevicesClientUploadCertificateResponse contains the response from method DevicesClient.UploadCertificate.
type DevicesClientUploadCertificateResponse struct {
	// The upload registration certificate response.
	UploadCertificateResponse
}

// DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse contains the response from method DiagnosticSettingsClient.GetDiagnosticProactiveLogCollectionSettings.
type DiagnosticSettingsClientGetDiagnosticProactiveLogCollectionSettingsResponse struct {
	// The diagnostic proactive log collection settings of a device.
	DiagnosticProactiveLogCollectionSettings
}

// DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse contains the response from method DiagnosticSettingsClient.GetDiagnosticRemoteSupportSettings.
type DiagnosticSettingsClientGetDiagnosticRemoteSupportSettingsResponse struct {
	// The remote support settings of a device.
	DiagnosticRemoteSupportSettings
}

// DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse contains the response from method DiagnosticSettingsClient.BeginUpdateDiagnosticProactiveLogCollectionSettings.
type DiagnosticSettingsClientUpdateDiagnosticProactiveLogCollectionSettingsResponse struct {
	// The diagnostic proactive log collection settings of a device.
	DiagnosticProactiveLogCollectionSettings
}

// DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse contains the response from method DiagnosticSettingsClient.BeginUpdateDiagnosticRemoteSupportSettings.
type DiagnosticSettingsClientUpdateDiagnosticRemoteSupportSettingsResponse struct {
	// The remote support settings of a device.
	DiagnosticRemoteSupportSettings
}

// JobsClientGetResponse contains the response from method JobsClient.Get.
type JobsClientGetResponse struct {
	// A device job.
	Job
}

// MonitoringConfigClientCreateOrUpdateResponse contains the response from method MonitoringConfigClient.BeginCreateOrUpdate.
type MonitoringConfigClientCreateOrUpdateResponse struct {
	// The metric setting details for the role
	MonitoringMetricConfiguration
}

// MonitoringConfigClientDeleteResponse contains the response from method MonitoringConfigClient.BeginDelete.
type MonitoringConfigClientDeleteResponse struct {
	// placeholder for future response values
}

// MonitoringConfigClientGetResponse contains the response from method MonitoringConfigClient.Get.
type MonitoringConfigClientGetResponse struct {
	// The metric setting details for the role
	MonitoringMetricConfiguration
}

// MonitoringConfigClientListResponse contains the response from method MonitoringConfigClient.NewListPager.
type MonitoringConfigClientListResponse struct {
	// Collection of metric configurations.
	MonitoringMetricConfigurationList
}

// NodesClientListByDataBoxEdgeDeviceResponse contains the response from method NodesClient.NewListByDataBoxEdgeDevicePager.
type NodesClientListByDataBoxEdgeDeviceResponse struct {
	// Collection of Nodes.
	NodeList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The list of operations used for the discovery of available provider operations.
	OperationsList
}

// OperationsStatusClientGetResponse contains the response from method OperationsStatusClient.Get.
type OperationsStatusClientGetResponse struct {
	// A device job.
	Job
}

// OrdersClientCreateOrUpdateResponse contains the response from method OrdersClient.BeginCreateOrUpdate.
type OrdersClientCreateOrUpdateResponse struct {
	// The order details.
	Order
}

// OrdersClientDeleteResponse contains the response from method OrdersClient.BeginDelete.
type OrdersClientDeleteResponse struct {
	// placeholder for future response values
}

// OrdersClientGetResponse contains the response from method OrdersClient.Get.
type OrdersClientGetResponse struct {
	// The order details.
	Order
}

// OrdersClientListByDataBoxEdgeDeviceResponse contains the response from method OrdersClient.NewListByDataBoxEdgeDevicePager.
type OrdersClientListByDataBoxEdgeDeviceResponse struct {
	// List of order entities.
	OrderList
}

// OrdersClientListDCAccessCodeResponse contains the response from method OrdersClient.ListDCAccessCode.
type OrdersClientListDCAccessCodeResponse struct {
	// DC Access code in the case of Self Managed Shipping.
	DCAccessCode
}

// RolesClientCreateOrUpdateResponse contains the response from method RolesClient.BeginCreateOrUpdate.
type RolesClientCreateOrUpdateResponse struct {
	// Compute role.
	RoleClassification
}

// MarshalJSON implements the json.Marshaller interface for type RolesClientCreateOrUpdateResponse.
func (r RolesClientCreateOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.RoleClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RolesClientCreateOrUpdateResponse.
func (r *RolesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalRoleClassification(data)
	if err != nil {
		return err
	}
	r.RoleClassification = res
	return nil
}

// RolesClientDeleteResponse contains the response from method RolesClient.BeginDelete.
type RolesClientDeleteResponse struct {
	// placeholder for future response values
}

// RolesClientGetResponse contains the response from method RolesClient.Get.
type RolesClientGetResponse struct {
	// Compute role.
	RoleClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RolesClientGetResponse.
func (r *RolesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalRoleClassification(data)
	if err != nil {
		return err
	}
	r.RoleClassification = res
	return nil
}

// RolesClientListByDataBoxEdgeDeviceResponse contains the response from method RolesClient.NewListByDataBoxEdgeDevicePager.
type RolesClientListByDataBoxEdgeDeviceResponse struct {
	// Collection of all the roles on the Data Box Edge device.
	RoleList
}

// SharesClientCreateOrUpdateResponse contains the response from method SharesClient.BeginCreateOrUpdate.
type SharesClientCreateOrUpdateResponse struct {
	// Represents a share on the Data Box Edge/Gateway device.
	Share
}

// SharesClientDeleteResponse contains the response from method SharesClient.BeginDelete.
type SharesClientDeleteResponse struct {
	// placeholder for future response values
}

// SharesClientGetResponse contains the response from method SharesClient.Get.
type SharesClientGetResponse struct {
	// Represents a share on the Data Box Edge/Gateway device.
	Share
}

// SharesClientListByDataBoxEdgeDeviceResponse contains the response from method SharesClient.NewListByDataBoxEdgeDevicePager.
type SharesClientListByDataBoxEdgeDeviceResponse struct {
	// Collection of all the shares on the Data Box Edge/Gateway device.
	ShareList
}

// SharesClientRefreshResponse contains the response from method SharesClient.BeginRefresh.
type SharesClientRefreshResponse struct {
	// placeholder for future response values
}

// StorageAccountCredentialsClientCreateOrUpdateResponse contains the response from method StorageAccountCredentialsClient.BeginCreateOrUpdate.
type StorageAccountCredentialsClientCreateOrUpdateResponse struct {
	// The storage account credential.
	StorageAccountCredential
}

// StorageAccountCredentialsClientDeleteResponse contains the response from method StorageAccountCredentialsClient.BeginDelete.
type StorageAccountCredentialsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageAccountCredentialsClientGetResponse contains the response from method StorageAccountCredentialsClient.Get.
type StorageAccountCredentialsClientGetResponse struct {
	// The storage account credential.
	StorageAccountCredential
}

// StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse contains the response from method StorageAccountCredentialsClient.NewListByDataBoxEdgeDevicePager.
type StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse struct {
	// The collection of storage account credentials.
	StorageAccountCredentialList
}

// StorageAccountsClientCreateOrUpdateResponse contains the response from method StorageAccountsClient.BeginCreateOrUpdate.
type StorageAccountsClientCreateOrUpdateResponse struct {
	// Represents a Storage Account on the Data Box Edge/Gateway device.
	StorageAccount
}

// StorageAccountsClientDeleteResponse contains the response from method StorageAccountsClient.BeginDelete.
type StorageAccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageAccountsClientGetResponse contains the response from method StorageAccountsClient.Get.
type StorageAccountsClientGetResponse struct {
	// Represents a Storage Account on the Data Box Edge/Gateway device.
	StorageAccount
}

// StorageAccountsClientListByDataBoxEdgeDeviceResponse contains the response from method StorageAccountsClient.NewListByDataBoxEdgeDevicePager.
type StorageAccountsClientListByDataBoxEdgeDeviceResponse struct {
	// Collection of all the Storage Accounts on the Data Box Edge/Gateway device.
	StorageAccountList
}

// SupportPackagesClientTriggerSupportPackageResponse contains the response from method SupportPackagesClient.BeginTriggerSupportPackage.
type SupportPackagesClientTriggerSupportPackageResponse struct {
	// placeholder for future response values
}

// TriggersClientCreateOrUpdateResponse contains the response from method TriggersClient.BeginCreateOrUpdate.
type TriggersClientCreateOrUpdateResponse struct {
	// Trigger details.
	TriggerClassification
}

// MarshalJSON implements the json.Marshaller interface for type TriggersClientCreateOrUpdateResponse.
func (t TriggersClientCreateOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.TriggerClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientCreateOrUpdateResponse.
func (t *TriggersClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}

// TriggersClientDeleteResponse contains the response from method TriggersClient.BeginDelete.
type TriggersClientDeleteResponse struct {
	// placeholder for future response values
}

// TriggersClientGetResponse contains the response from method TriggersClient.Get.
type TriggersClientGetResponse struct {
	// Trigger details.
	TriggerClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TriggersClientGetResponse.
func (t *TriggersClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalTriggerClassification(data)
	if err != nil {
		return err
	}
	t.TriggerClassification = res
	return nil
}

// TriggersClientListByDataBoxEdgeDeviceResponse contains the response from method TriggersClient.NewListByDataBoxEdgeDevicePager.
type TriggersClientListByDataBoxEdgeDeviceResponse struct {
	// Collection of all trigger on the data box edge device.
	TriggerList
}

// UsersClientCreateOrUpdateResponse contains the response from method UsersClient.BeginCreateOrUpdate.
type UsersClientCreateOrUpdateResponse struct {
	// Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
	User
}

// UsersClientDeleteResponse contains the response from method UsersClient.BeginDelete.
type UsersClientDeleteResponse struct {
	// placeholder for future response values
}

// UsersClientGetResponse contains the response from method UsersClient.Get.
type UsersClientGetResponse struct {
	// Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
	User
}

// UsersClientListByDataBoxEdgeDeviceResponse contains the response from method UsersClient.NewListByDataBoxEdgeDevicePager.
type UsersClientListByDataBoxEdgeDeviceResponse struct {
	// Collection of users.
	UserList
}
