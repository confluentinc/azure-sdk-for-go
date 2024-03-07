//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/AdministratorAdd.json
func ExampleAdministratorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAdministratorsClient().BeginCreate(ctx, "testrg", "testserver", "oooooooo-oooo-oooo-oooo-oooooooooooo", armpostgresqlflexibleservers.ActiveDirectoryAdministratorAdd{
		Properties: &armpostgresqlflexibleservers.AdministratorPropertiesForAdd{
			PrincipalName: to.Ptr("testuser1@microsoft.com"),
			PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeUser),
			TenantID:      to.Ptr("tttttttt-tttt-tttt-tttt-tttttttttttt"),
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
	// res.ActiveDirectoryAdministrator = armpostgresqlflexibleservers.ActiveDirectoryAdministrator{
	// 	Name: to.Ptr("testuser1@microsoft.com"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/administrators/oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 	Properties: &armpostgresqlflexibleservers.AdministratorProperties{
	// 		ObjectID: to.Ptr("oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 		PrincipalName: to.Ptr("testuser1@microsoft.com"),
	// 		PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/AdministratorDelete.json
func ExampleAdministratorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAdministratorsClient().BeginDelete(ctx, "testrg", "testserver", "oooooooo-oooo-oooo-oooo-oooooooooooo", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/AdministratorGet.json
func ExampleAdministratorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdministratorsClient().Get(ctx, "testrg", "pgtestsvc1", "oooooooo-oooo-oooo-oooo-oooooooooooo", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ActiveDirectoryAdministrator = armpostgresqlflexibleservers.ActiveDirectoryAdministrator{
	// 	Name: to.Ptr("testuser1@microsoft.com"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc1/administrators/oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 	Properties: &armpostgresqlflexibleservers.AdministratorProperties{
	// 		ObjectID: to.Ptr("oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 		PrincipalName: to.Ptr("testuer1@microsoft.com"),
	// 		PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeUser),
	// 		TenantID: to.Ptr("tttttttt-tttt-tttt-tttt-tttttttttttt"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/AdministratorsListByServer.json
func ExampleAdministratorsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAdministratorsClient().NewListByServerPager("testrg", "pgtestsvc1", nil)
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
		// page.AdministratorListResult = armpostgresqlflexibleservers.AdministratorListResult{
		// 	Value: []*armpostgresqlflexibleservers.ActiveDirectoryAdministrator{
		// 		{
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc1/administrators/oooooooo-oooo-oooo-oooo-oooooooooooo"),
		// 			Properties: &armpostgresqlflexibleservers.AdministratorProperties{
		// 				ObjectID: to.Ptr("oooooooo-oooo-oooo-oooo-oooooooooooo"),
		// 				PrincipalName: to.Ptr("testuer1@microsoft.com"),
		// 				PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeUser),
		// 				TenantID: to.Ptr("tttttttt-tttt-tttt-tttt-tttttttttttt"),
		// 			},
		// 		},
		// 		{
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc1/administrators/gggggggg-gggg-gggg-gggg-gggggggggggg"),
		// 			Properties: &armpostgresqlflexibleservers.AdministratorProperties{
		// 				ObjectID: to.Ptr("gggggggg-gggg-gggg-gggg-gggggggggggg"),
		// 				PrincipalName: to.Ptr("testgroup1@microsoft.com"),
		// 				PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeGroup),
		// 				TenantID: to.Ptr("tttttttt-tttt-tttt-tttt-tttttttttttt"),
		// 			},
		// 	}},
		// }
	}
}