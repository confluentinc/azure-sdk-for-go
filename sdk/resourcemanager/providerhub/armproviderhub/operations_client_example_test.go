//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armproviderhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OperationsDefinitionArrayResponseWithContinuation = armproviderhub.OperationsDefinitionArrayResponseWithContinuation{
		// 	Value: []*armproviderhub.OperationsDefinition{
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/register/action"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Registers the specified subscription with Microsoft.ProviderHub resource provider"),
		// 				Operation: to.Ptr("Register for Microsoft.ProviderHub"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("register"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/write"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Creates or Updates any rollout"),
		// 				Operation: to.Ptr("Create or Update rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/read"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Reads any rollout"),
		// 				Operation: to.Ptr("Read rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/delete"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Deletes any rollout"),
		// 				Operation: to.Ptr("Delete rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/defaultRollouts/stop/action"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Deletes any rollout"),
		// 				Operation: to.Ptr("Delete rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("defaultRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/customRollouts/write"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Creates or Updates any rollout"),
		// 				Operation: to.Ptr("Create or Update rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("customRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ProviderHub/customRollouts/read"),
		// 			Display: &armproviderhub.OperationsDefinitionDisplay{
		// 				Description: to.Ptr("Reads any rollout"),
		// 				Operation: to.Ptr("Read rollout"),
		// 				Provider: to.Ptr("Microsoft ProviderHub"),
		// 				Resource: to.Ptr("customRollouts"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_ListByProviderRegistration.json
func ExampleOperationsClient_ListByProviderRegistration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().ListByProviderRegistration(ctx, "Microsoft.Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationsDefinitionArray = []*armproviderhub.OperationsDefinition{
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Read"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Read employees"),
	// 			Operation: to.Ptr("Gets/List employee resources"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(false),
	// 	},
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Write"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Writes employees"),
	// 			Operation: to.Ptr("Create/update employee resources"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(false),
	// 	},
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Delete"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Deletes employees"),
	// 			Operation: to.Ptr("Deletes employee resource"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(false),
	// 		Origin: to.Ptr(armproviderhub.OperationsDefinitionOriginUser),
	// 	},
	// 	{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Action"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Writes employees"),
	// 			Operation: to.Ptr("Create/update employee resources"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 		IsDataAction: to.Ptr(true),
	// 		Origin: to.Ptr(armproviderhub.OperationsDefinitionOriginSystem),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_CreateOrUpdate.json
func ExampleOperationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationsClient().CreateOrUpdate(ctx, "Microsoft.Contoso", armproviderhub.OperationsPutContent{
		Contents: []*armproviderhub.OperationsDefinition{
			{
				Name: to.Ptr("Microsoft.Contoso/Employees/Read"),
				Display: &armproviderhub.OperationsDefinitionDisplay{
					Description: to.Ptr("Read employees"),
					Operation:   to.Ptr("Gets/List employee resources"),
					Provider:    to.Ptr("Microsoft.Contoso"),
					Resource:    to.Ptr("Employees"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationsContent = armproviderhub.OperationsContent{
	// 	Properties: &armproviderhub.OperationsDefinition{
	// 		Name: to.Ptr("Microsoft.Contoso/Employees/Read"),
	// 		Display: &armproviderhub.OperationsDefinitionDisplay{
	// 			Description: to.Ptr("Read employees"),
	// 			Operation: to.Ptr("Gets/List employee resources"),
	// 			Provider: to.Ptr("Microsoft.Contoso"),
	// 			Resource: to.Ptr("Employees"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/Operations_Delete.json
func ExampleOperationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armproviderhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewOperationsClient().Delete(ctx, "Microsoft.Contoso", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}