//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

// ActiveDirectoryConnectorsClientCreateResponse contains the response from method ActiveDirectoryConnectorsClient.BeginCreate.
type ActiveDirectoryConnectorsClientCreateResponse struct {
	// Active directory connector resource
	ActiveDirectoryConnectorResource
}

// ActiveDirectoryConnectorsClientDeleteResponse contains the response from method ActiveDirectoryConnectorsClient.BeginDelete.
type ActiveDirectoryConnectorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ActiveDirectoryConnectorsClientGetResponse contains the response from method ActiveDirectoryConnectorsClient.Get.
type ActiveDirectoryConnectorsClientGetResponse struct {
	// Active directory connector resource
	ActiveDirectoryConnectorResource
}

// ActiveDirectoryConnectorsClientListResponse contains the response from method ActiveDirectoryConnectorsClient.NewListPager.
type ActiveDirectoryConnectorsClientListResponse struct {
	// A list of active directory connectors
	ActiveDirectoryConnectorListResult
}

// DataControllersClientDeleteDataControllerResponse contains the response from method DataControllersClient.BeginDeleteDataController.
type DataControllersClientDeleteDataControllerResponse struct {
	// placeholder for future response values
}

// DataControllersClientGetDataControllerResponse contains the response from method DataControllersClient.GetDataController.
type DataControllersClientGetDataControllerResponse struct {
	// Data controller resource
	DataControllerResource
}

// DataControllersClientListInGroupResponse contains the response from method DataControllersClient.NewListInGroupPager.
type DataControllersClientListInGroupResponse struct {
	// A list of data controllers.
	PageOfDataControllerResource
}

// DataControllersClientListInSubscriptionResponse contains the response from method DataControllersClient.NewListInSubscriptionPager.
type DataControllersClientListInSubscriptionResponse struct {
	// A list of data controllers.
	PageOfDataControllerResource
}

// DataControllersClientPatchDataControllerResponse contains the response from method DataControllersClient.BeginPatchDataController.
type DataControllersClientPatchDataControllerResponse struct {
	// Data controller resource
	DataControllerResource
}

// DataControllersClientPutDataControllerResponse contains the response from method DataControllersClient.BeginPutDataController.
type DataControllersClientPutDataControllerResponse struct {
	// Data controller resource
	DataControllerResource
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Azure Data Services on Azure Arc operations.
	OperationListResult
}

// PostgresInstancesClientCreateResponse contains the response from method PostgresInstancesClient.BeginCreate.
type PostgresInstancesClientCreateResponse struct {
	// A Postgres Instance.
	PostgresInstance
}

// PostgresInstancesClientDeleteResponse contains the response from method PostgresInstancesClient.BeginDelete.
type PostgresInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// PostgresInstancesClientGetResponse contains the response from method PostgresInstancesClient.Get.
type PostgresInstancesClientGetResponse struct {
	// A Postgres Instance.
	PostgresInstance
}

// PostgresInstancesClientListByResourceGroupResponse contains the response from method PostgresInstancesClient.NewListByResourceGroupPager.
type PostgresInstancesClientListByResourceGroupResponse struct {
	// A list of PostgresInstance.
	PostgresInstanceListResult
}

// PostgresInstancesClientListResponse contains the response from method PostgresInstancesClient.NewListPager.
type PostgresInstancesClientListResponse struct {
	// A list of PostgresInstance.
	PostgresInstanceListResult
}

// PostgresInstancesClientUpdateResponse contains the response from method PostgresInstancesClient.Update.
type PostgresInstancesClientUpdateResponse struct {
	// A Postgres Instance.
	PostgresInstance
}

// SQLManagedInstancesClientCreateResponse contains the response from method SQLManagedInstancesClient.BeginCreate.
type SQLManagedInstancesClientCreateResponse struct {
	// A SqlManagedInstance.
	SQLManagedInstance
}

// SQLManagedInstancesClientDeleteResponse contains the response from method SQLManagedInstancesClient.BeginDelete.
type SQLManagedInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// SQLManagedInstancesClientGetResponse contains the response from method SQLManagedInstancesClient.Get.
type SQLManagedInstancesClientGetResponse struct {
	// A SqlManagedInstance.
	SQLManagedInstance
}

// SQLManagedInstancesClientListByResourceGroupResponse contains the response from method SQLManagedInstancesClient.NewListByResourceGroupPager.
type SQLManagedInstancesClientListByResourceGroupResponse struct {
	// A list of SqlManagedInstance.
	SQLManagedInstanceListResult
}

// SQLManagedInstancesClientListResponse contains the response from method SQLManagedInstancesClient.NewListPager.
type SQLManagedInstancesClientListResponse struct {
	// A list of SqlManagedInstance.
	SQLManagedInstanceListResult
}

// SQLManagedInstancesClientUpdateResponse contains the response from method SQLManagedInstancesClient.Update.
type SQLManagedInstancesClientUpdateResponse struct {
	// A SqlManagedInstance.
	SQLManagedInstance
}

// SQLServerInstancesClientCreateResponse contains the response from method SQLServerInstancesClient.BeginCreate.
type SQLServerInstancesClientCreateResponse struct {
	// A SqlServerInstance.
	SQLServerInstance
}

// SQLServerInstancesClientDeleteResponse contains the response from method SQLServerInstancesClient.BeginDelete.
type SQLServerInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// SQLServerInstancesClientGetResponse contains the response from method SQLServerInstancesClient.Get.
type SQLServerInstancesClientGetResponse struct {
	// A SqlServerInstance.
	SQLServerInstance
}

// SQLServerInstancesClientListByResourceGroupResponse contains the response from method SQLServerInstancesClient.NewListByResourceGroupPager.
type SQLServerInstancesClientListByResourceGroupResponse struct {
	// A list of SqlServerInstance.
	SQLServerInstanceListResult
}

// SQLServerInstancesClientListResponse contains the response from method SQLServerInstancesClient.NewListPager.
type SQLServerInstancesClientListResponse struct {
	// A list of SqlServerInstance.
	SQLServerInstanceListResult
}

// SQLServerInstancesClientUpdateResponse contains the response from method SQLServerInstancesClient.Update.
type SQLServerInstancesClientUpdateResponse struct {
	// A SqlServerInstance.
	SQLServerInstance
}