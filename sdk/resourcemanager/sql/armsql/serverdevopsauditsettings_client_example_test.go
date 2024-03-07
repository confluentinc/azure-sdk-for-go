//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditList.json
func ExampleServerDevOpsAuditSettingsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerDevOpsAuditSettingsClient().NewListByServerPager("devAuditTestRG", "devOpsAuditTestSvr", nil)
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
		// page.ServerDevOpsAuditSettingsListResult = armsql.ServerDevOpsAuditSettingsListResult{
		// 	Value: []*armsql.ServerDevOpsAuditingSettings{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/devOpsAuditingSettings"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/devAuditTestRG/providers/Microsoft.Sql/servers/devOpsAuditTestSvr/devOpsAuditingSettings/default"),
		// 			Properties: &armsql.ServerDevOpsAuditSettingsProperties{
		// 				IsAzureMonitorTargetEnabled: to.Ptr(false),
		// 				IsManagedIdentityInUse: to.Ptr(false),
		// 				State: to.Ptr(armsql.BlobAuditingPolicyStateDisabled),
		// 				StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditGet.json
func ExampleServerDevOpsAuditSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerDevOpsAuditSettingsClient().Get(ctx, "devAuditTestRG", "devOpsAuditTestSvr", armsql.DevOpsAuditingSettingsNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerDevOpsAuditingSettings = armsql.ServerDevOpsAuditingSettings{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/devOpsAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/devAuditTestRG/providers/Microsoft.Sql/servers/devOpsAuditTestSvr/devOpsAuditingSettings/default"),
	// 	Properties: &armsql.ServerDevOpsAuditSettingsProperties{
	// 		IsAzureMonitorTargetEnabled: to.Ptr(false),
	// 		IsManagedIdentityInUse: to.Ptr(false),
	// 		State: to.Ptr(armsql.BlobAuditingPolicyStateDisabled),
	// 		StorageAccountSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditCreateMax.json
func ExampleServerDevOpsAuditSettingsClient_BeginCreateOrUpdate_updateAServersDevOpsAuditSettingsWithAllParams() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerDevOpsAuditSettingsClient().BeginCreateOrUpdate(ctx, "devAuditTestRG", "devOpsAuditTestSvr", armsql.DevOpsAuditingSettingsNameDefault, armsql.ServerDevOpsAuditingSettings{
		Properties: &armsql.ServerDevOpsAuditSettingsProperties{
			IsAzureMonitorTargetEnabled:  to.Ptr(true),
			State:                        to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey:      to.Ptr("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
			StorageEndpoint:              to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerDevOpsAuditingSettings = armsql.ServerDevOpsAuditingSettings{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/devOpsAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/devAuditTestRG/providers/Microsoft.Sql/servers/devOpsAuditTestSvr/devOpsAuditingSettings/default"),
	// 	Properties: &armsql.ServerDevOpsAuditSettingsProperties{
	// 		IsAzureMonitorTargetEnabled: to.Ptr(true),
	// 		State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 		StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 		StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2022-02-01-preview/examples/ServerDevOpsAuditCreateMin.json
func ExampleServerDevOpsAuditSettingsClient_BeginCreateOrUpdate_updateAServersDevOpsAuditSettingsWithMinimalInput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServerDevOpsAuditSettingsClient().BeginCreateOrUpdate(ctx, "devAuditTestRG", "devOpsAuditTestSvr", armsql.DevOpsAuditingSettingsNameDefault, armsql.ServerDevOpsAuditingSettings{
		Properties: &armsql.ServerDevOpsAuditSettingsProperties{
			State:                   to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
			StorageAccountAccessKey: to.Ptr("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
			StorageEndpoint:         to.Ptr("https://mystorage.blob.core.windows.net"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerDevOpsAuditingSettings = armsql.ServerDevOpsAuditingSettings{
	// 	Name: to.Ptr("Default"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/devOpsAuditingSettings"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/devAuditTestRG/providers/Microsoft.Sql/servers/devOpsAuditTestSvr/devOpsAuditingSettings/default"),
	// 	Properties: &armsql.ServerDevOpsAuditSettingsProperties{
	// 		State: to.Ptr(armsql.BlobAuditingPolicyStateEnabled),
	// 		StorageAccountSubscriptionID: to.Ptr("00000000-1234-0000-5678-000000000000"),
	// 		StorageEndpoint: to.Ptr("https://mystorage.blob.core.windows.net"),
	// 	},
	// }
}