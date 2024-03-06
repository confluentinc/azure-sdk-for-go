//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2023-05-01/examples/OperationsList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkubernetesconfiguration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.ResourceProviderOperationList = armkubernetesconfiguration.ResourceProviderOperationList{
		// 	Value: []*armkubernetesconfiguration.ResourceProviderOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations/write"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Creates or updates source control configuration."),
		// 				Operation: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations/write"),
		// 				Provider: to.Ptr("Microsoft Kubernetes Configuration"),
		// 				Resource: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations/read"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Gets source control configuration."),
		// 				Operation: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations/read"),
		// 				Provider: to.Ptr("Microsoft Kubernetes Configuration"),
		// 				Resource: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations/delete"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Deletes source control configuration."),
		// 				Operation: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations/delete"),
		// 				Provider: to.Ptr("Microsoft Kubernetes Configuration"),
		// 				Resource: to.Ptr("Microsoft.KubernetesConfiguration/sourceControlConfigurations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/extensions/read"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Get the specified Extension."),
		// 				Operation: to.Ptr("Get extension"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("extensions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/extensions/write"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Create the Extension specified."),
		// 				Operation: to.Ptr("Create a Extension"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("extensions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/extensions/delete"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Delete the specified Extension."),
		// 				Operation: to.Ptr("Delete Extension"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("extensions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/extensions/operations/read"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Get the Status of the Extension Async Operation."),
		// 				Operation: to.Ptr("Get Extension Async Operation Status"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("extensions"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations/read"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Get the specified Flux Configuration."),
		// 				Operation: to.Ptr("Get fluxConfiguration"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("fluxConfigurations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations/write"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Create the Flux Configuration specified."),
		// 				Operation: to.Ptr("Create a fluxConfiguration"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("fluxConfigurations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations/delete"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Delete the specified Flux Configuration."),
		// 				Operation: to.Ptr("Delete Flux Configuration"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("fluxConfigurations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/fluxConfigurations/operations/read"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Get the Status of the Flux Configuration Async Operation."),
		// 				Operation: to.Ptr("Get Flux Configuration Async Operation Status"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("fluxConfigurations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/register/action"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Registers subscription to Microsoft.KubernetesConfiguration resource provider."),
		// 				Operation: to.Ptr("Microsoft.KubernetesConfiguration/register/action"),
		// 				Provider: to.Ptr("Microsoft Kubernetes Configuration"),
		// 				Resource: to.Ptr("Register"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.KubernetesConfiguration/extensionTypes/read"),
		// 			Display: &armkubernetesconfiguration.ResourceProviderOperationDisplay{
		// 				Description: to.Ptr("Get Kubernetes Configuration Available Extensions"),
		// 				Operation: to.Ptr("Microsoft.KubernetesConfiguration/extensionTypes/read"),
		// 				Provider: to.Ptr("Microsoft KubernetesConfiguration"),
		// 				Resource: to.Ptr("extensionTypes"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}
