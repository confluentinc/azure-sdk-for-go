//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a5e7ff51c8af3781e7f6dd3b82a3fc922e2f6f93/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/ScriptPackages_List.json
func ExampleScriptPackagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScriptPackagesClient().NewListPager("group1", "cloud1", nil)
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
		// page.ScriptPackagesList = armavs.ScriptPackagesList{
		// 	Value: []*armavs.ScriptPackage{
		// 		{
		// 			Name: to.Ptr("Microsoft.AVS.Management@3.0.48"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/scriptPackages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/Microsoft.AVS.Management@3.0.48"),
		// 			Properties: &armavs.ScriptPackageProperties{
		// 				Description: to.Ptr("Various cmdlets for elevated access to Private Cloud administrative functions"),
		// 				Company: to.Ptr("Microsoft"),
		// 				URI: to.Ptr("https://microsoft.com"),
		// 				Version: to.Ptr("3.0.48"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("JSDR.Configuration@1.0.0"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/scriptPackages"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/JSDR.Configuration@1.0.0"),
		// 			Properties: &armavs.ScriptPackageProperties{
		// 				Description: to.Ptr("Various cmdlets by Jetstream for Private Cloud administration"),
		// 				Company: to.Ptr("Jetstream Software"),
		// 				URI: to.Ptr("https://www.jetstreamsoft.com/about/support/"),
		// 				Version: to.Ptr("1.0.0"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a5e7ff51c8af3781e7f6dd3b82a3fc922e2f6f93/specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/ScriptPackages_Get.json
func ExampleScriptPackagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptPackagesClient().Get(ctx, "group1", "cloud1", "Microsoft.AVS.Management@3.0.48", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScriptPackage = armavs.ScriptPackage{
	// 	Name: to.Ptr("Microsoft.AVS.Management@3.0.48"),
	// 	Type: to.Ptr("Microsoft.AVS/privateClouds/scriptPackages"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/Microsoft.AVS.Management@3.0.48"),
	// 	Properties: &armavs.ScriptPackageProperties{
	// 		Description: to.Ptr("Various cmdlets for elevated access to Private Cloud administrative functions"),
	// 		Company: to.Ptr("Microsoft"),
	// 		URI: to.Ptr("https://microsoft.com"),
	// 		Version: to.Ptr("3.0.48"),
	// 	},
	// }
}
