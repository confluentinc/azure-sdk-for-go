//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArmTemplates_List.json
func ExampleArmTemplatesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArmTemplatesClient().NewListPager("resourceGroupName", "{labName}", "{artifactSourceName}", &armdevtestlabs.ArmTemplatesClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.ArmTemplateList = armdevtestlabs.ArmTemplateList{
		// 	Value: []*armdevtestlabs.ArmTemplate{
		// 		{
		// 			Name: to.Ptr("Template1"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/armTemplates"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/{artifactSourceName}/armTemplates/{armTemplateName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"MyTag": to.Ptr("MyValue"),
		// 			},
		// 			Properties: &armdevtestlabs.ArmTemplateProperties{
		// 				Contents: map[string]any{
		// 					"$schema": "http://schema.management.azure.com/schemas/2014-04-01-preview/deploymentTemplate.json#",
		// 					"contentVersion": "1.0.0.0",
		// 					"parameters":map[string]any{
		// 					},
		// 					"resources":[]any{
		// 					},
		// 					"variables":map[string]any{
		// 						"hostingPlanName": "[toLower(concat(variables('resourceNamePrefix'), '-', take(uniqueString(resourceGroup().id), 6), '-sp'))]",
		// 						"resourceNamePrefix": "[take(uniqueString(resourceGroup().id), 3)]",
		// 						"siteName": "[toLower(concat(variables('resourceNamePrefix'), '-', take(uniqueString(resourceGroup().id), 6)))]",
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("Template1"),
		// 				Enabled: to.Ptr(true),
		// 				ParametersValueFilesInfo: []*armdevtestlabs.ParametersValueFileInfo{
		// 				},
		// 				Publisher: to.Ptr("Microsoft"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArmTemplates_Get.json
func ExampleArmTemplatesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArmTemplatesClient().Get(ctx, "resourceGroupName", "{labName}", "{artifactSourceName}", "{armTemplateName}", &armdevtestlabs.ArmTemplatesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArmTemplate = armdevtestlabs.ArmTemplate{
	// 	Name: to.Ptr("Template1"),
	// 	Type: to.Ptr("Microsoft.DevTestLab/labs/artifactSources/armTemplates"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/artifactSources/{artifactSourceName}/armTemplates/{armTemplateName}"),
	// 	Location: to.Ptr("{location}"),
	// 	Tags: map[string]*string{
	// 		"MyTag": to.Ptr("MyValue"),
	// 	},
	// 	Properties: &armdevtestlabs.ArmTemplateProperties{
	// 		Contents: map[string]any{
	// 			"$schema": "http://schema.management.azure.com/schemas/2014-04-01-preview/deploymentTemplate.json#",
	// 			"contentVersion": "1.0.0.0",
	// 			"parameters":map[string]any{
	// 			},
	// 			"resources":[]any{
	// 			},
	// 			"variables":map[string]any{
	// 				"hostingPlanName": "[toLower(concat(variables('resourceNamePrefix'), '-', take(uniqueString(resourceGroup().id), 6), '-sp'))]",
	// 				"resourceNamePrefix": "[take(uniqueString(resourceGroup().id), 3)]",
	// 				"siteName": "[toLower(concat(variables('resourceNamePrefix'), '-', take(uniqueString(resourceGroup().id), 6)))]",
	// 			},
	// 		},
	// 		DisplayName: to.Ptr("Template1"),
	// 		Enabled: to.Ptr(true),
	// 		ParametersValueFilesInfo: []*armdevtestlabs.ParametersValueFileInfo{
	// 		},
	// 		Publisher: to.Ptr("Microsoft"),
	// 	},
	// }
}