//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/DeletedBackupInstanceOperations/ListDeletedBackupInstances.json
func ExampleDeletedBackupInstancesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedBackupInstancesClient().NewListPager("000pikumar", "PratikPrivatePreviewVault1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DeletedBackupInstanceResourceList = armdataprotection.DeletedBackupInstanceResourceList{
		// 	Value: []*armdataprotection.DeletedBackupInstanceResource{
		// 		{
		// 			Name: to.Ptr("testInstance1"),
		// 			Type: to.Ptr("Microsoft.DataProtection/backupVaults/deletedBackupInstances"),
		// 			ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PratikPrivatePreviewVault1/deletedBackupInstances/testInstance1"),
		// 			Properties: &armdataprotection.DeletedBackupInstance{
		// 				DataSourceInfo: &armdataprotection.Datasource{
		// 					DatasourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
		// 					ObjectType: to.Ptr("Datasource"),
		// 					ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
		// 					ResourceLocation: to.Ptr(""),
		// 					ResourceName: to.Ptr("testdb"),
		// 					ResourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
		// 					ResourceURI: to.Ptr(""),
		// 				},
		// 				DataSourceSetInfo: &armdataprotection.DatasourceSet{
		// 					DatasourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
		// 					ObjectType: to.Ptr("DatasourceSet"),
		// 					ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
		// 					ResourceLocation: to.Ptr(""),
		// 					ResourceName: to.Ptr("viveksipgtest"),
		// 					ResourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
		// 					ResourceURI: to.Ptr(""),
		// 				},
		// 				FriendlyName: to.Ptr("testInstance1"),
		// 				ObjectType: to.Ptr("DeletedBackupInstance"),
		// 				PolicyInfo: &armdataprotection.PolicyInfo{
		// 					PolicyID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PratikPrivatePreviewVault1/backupPolicies/PratikPolicy1"),
		// 				},
		// 				ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
		// 					Status: to.Ptr(armdataprotection.StatusSoftDeleted),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				DeletionInfo: &armdataprotection.DeletionInfo{
		// 					BillingEndDate: to.Ptr("2022-05-06T00:00:36.6660445Z"),
		// 					DeleteActivityID: to.Ptr("1e9ec790-d198-4efb-bbd7-e4669d5351a4"),
		// 					DeletionTime: to.Ptr("2022-05-04T00:00:36.6660445Z"),
		// 					ScheduledPurgeTime: to.Ptr("2022-05-20T00:00:36.6660445Z"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/DeletedBackupInstanceOperations/GetDeletedBackupInstance.json
func ExampleDeletedBackupInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedBackupInstancesClient().Get(ctx, "000pikumar", "PratikPrivatePreviewVault1", "testInstance1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedBackupInstanceResource = armdataprotection.DeletedBackupInstanceResource{
	// 	Name: to.Ptr("testInstance1"),
	// 	Type: to.Ptr("Microsoft.DataProtection/backupVaults/deletedBackupInstances"),
	// 	ID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PratikPrivatePreviewVault1/deletedBackupInstances/testInstance1"),
	// 	Properties: &armdataprotection.DeletedBackupInstance{
	// 		DataSourceInfo: &armdataprotection.Datasource{
	// 			DatasourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
	// 			ObjectType: to.Ptr("Datasource"),
	// 			ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb"),
	// 			ResourceLocation: to.Ptr(""),
	// 			ResourceName: to.Ptr("testdb"),
	// 			ResourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
	// 			ResourceURI: to.Ptr(""),
	// 		},
	// 		DataSourceSetInfo: &armdataprotection.DatasourceSet{
	// 			DatasourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases"),
	// 			ObjectType: to.Ptr("DatasourceSet"),
	// 			ResourceID: to.Ptr("/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest"),
	// 			ResourceLocation: to.Ptr(""),
	// 			ResourceName: to.Ptr("viveksipgtest"),
	// 			ResourceType: to.Ptr("Microsoft.DBforPostgreSQL/servers"),
	// 			ResourceURI: to.Ptr(""),
	// 		},
	// 		FriendlyName: to.Ptr("testInstance1"),
	// 		ObjectType: to.Ptr("DeletedBackupInstance"),
	// 		PolicyInfo: &armdataprotection.PolicyInfo{
	// 			PolicyID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/000pikumar/providers/Microsoft.DataProtection/backupVaults/PratikPrivatePreviewVault1/backupPolicies/PratikPolicy1"),
	// 		},
	// 		ProtectionStatus: &armdataprotection.ProtectionStatusDetails{
	// 			Status: to.Ptr(armdataprotection.StatusSoftDeleted),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		DeletionInfo: &armdataprotection.DeletionInfo{
	// 			BillingEndDate: to.Ptr("2022-05-06T00:00:36.6660445Z"),
	// 			DeleteActivityID: to.Ptr("1e9ec790-d198-4efb-bbd7-e4669d5351a4"),
	// 			DeletionTime: to.Ptr("2022-05-04T00:00:36.6660445Z"),
	// 			ScheduledPurgeTime: to.Ptr("2022-05-20T00:00:36.6660445Z"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-11-01/examples/DeletedBackupInstanceOperations/UndeleteDeletedBackupInstance.json
func ExampleDeletedBackupInstancesClient_BeginUndelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataprotection.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeletedBackupInstancesClient().BeginUndelete(ctx, "testrg", "testvault", "testbi", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}