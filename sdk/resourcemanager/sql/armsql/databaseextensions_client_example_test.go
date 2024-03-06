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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/GetDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDatabaseExtensionsClient().Get(ctx, "rg_a1f9d6f8-30d5-4228-9504-8a364361bca3", "srv_65858e0f-b1d1-4bdc-8351-a7da86ca4939", "11aa6c5e-58ed-4693-b303-3b8e3131deaa", "polybaseimport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/CreateOrUpdateDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDatabaseExtensionsClient().BeginCreateOrUpdate(ctx, "rg_20cbe0f0-c2d9-4522-9177-5469aad53029", "srv_1ffd1cf8-9951-47fb-807d-a9c384763849", "878e303f-1ea0-4f17-aa3d-a22ac5e9da08", "polybaseimport", armsql.DatabaseExtensions{
		Properties: &armsql.DatabaseExtensionsProperties{
			OperationMode:  to.Ptr(armsql.OperationModePolybaseImport),
			StorageKey:     to.Ptr("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
			StorageKeyType: to.Ptr(armsql.StorageKeyTypeStorageAccessKey),
			StorageURI:     to.Ptr("https://teststorage.blob.core.windows.net/testcontainer/Manifest.xml"),
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
	// res.ImportExportExtensionsOperationResult = armsql.ImportExportExtensionsOperationResult{
	// 	Name: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/extensions"),
	// 	ID: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 	Properties: &armsql.ImportExportExtensionsOperationResultProperties{
	// 		DatabaseName: to.Ptr("878e303f-1ea0-4f17-aa3d-a22ac5e9da08"),
	// 		LastModifiedTime: to.Ptr("lastModifiedTime"),
	// 		RequestID: to.Ptr("10000000-0000-0000-0000-000000000002"),
	// 		RequestType: to.Ptr("PolybaseImport"),
	// 		ServerName: to.Ptr("srv_1ffd1cf8-9951-47fb-807d-a9c384763849"),
	// 		Status: to.Ptr("succeeded"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabaseExtensionsClient().NewListByDatabasePager("rg_4007c5a9-b3b0-41e1-bd46-9eef38768a4a", "srv_3b67ec2a-519b-43a7-8533-fb62dce3431e", "719d8fa4-bf0f-48fc-8cd3-ef40fe6ba1fe", nil)
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
		// page.ImportExportExtensionsOperationListResult = armsql.ImportExportExtensionsOperationListResult{
		// 	Value: []*armsql.ImportExportExtensionsOperationResult{
		// 	},
		// }
	}
}
