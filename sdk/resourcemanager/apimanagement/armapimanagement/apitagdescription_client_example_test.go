//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListApiTagDescriptions.json
func ExampleAPITagDescriptionClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPITagDescriptionClient().NewListByServicePager("rg1", "apimService1", "57d2ef278aa04f0888cba3f3", &armapimanagement.APITagDescriptionClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
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
		// page.TagDescriptionCollection = armapimanagement.TagDescriptionCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.TagDescriptionContract{
		// 		{
		// 			Name: to.Ptr("5600b539c53f5b0062060002"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/tags"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tagDescriptions/5600b539c53f5b0062060002"),
		// 			Properties: &armapimanagement.TagDescriptionContractProperties{
		// 				ExternalDocsDescription: to.Ptr("some additional info"),
		// 				ExternalDocsURL: to.Ptr("http://some_url.com"),
		// 				DisplayName: to.Ptr("tag1"),
		// 				TagID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/5600b539c53f5b0062060002"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiTagDescription.json
func ExampleAPITagDescriptionClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAPITagDescriptionClient().GetEntityTag(ctx, "rg1", "apimService1", "59d6bb8f1f7fab13dc67ec9b", "59306a29e4bbd510dc24e5f9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiTagDescription.json
func ExampleAPITagDescriptionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPITagDescriptionClient().Get(ctx, "rg1", "apimService1", "59d6bb8f1f7fab13dc67ec9b", "59306a29e4bbd510dc24e5f9", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagDescriptionContract = armapimanagement.TagDescriptionContract{
	// 	Name: to.Ptr("59306a29e4bbd510dc24e5f9"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/tagDescriptions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/59d6bb8f1f7fab13dc67ec9b/tagDescriptions/59306a29e4bbd510dc24e5f9"),
	// 	Properties: &armapimanagement.TagDescriptionContractProperties{
	// 		ExternalDocsDescription: to.Ptr("some additional info"),
	// 		ExternalDocsURL: to.Ptr("http://some_url.com"),
	// 		DisplayName: to.Ptr("tag1"),
	// 		TagID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/59306a29e4bbd510dc24e5f9"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiTagDescription.json
func ExampleAPITagDescriptionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPITagDescriptionClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5931a75ae4bbd512a88c680b", "tagId1", armapimanagement.TagDescriptionCreateParameters{
		Properties: &armapimanagement.TagDescriptionBaseProperties{
			Description:             to.Ptr("Some description that will be displayed for operation's tag if the tag is assigned to operation of the API"),
			ExternalDocsDescription: to.Ptr("Description of the external docs resource"),
			ExternalDocsURL:         to.Ptr("http://some.url/additionaldoc"),
		},
	}, &armapimanagement.APITagDescriptionClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TagDescriptionContract = armapimanagement.TagDescriptionContract{
	// 	Name: to.Ptr("tagId1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/tagDescriptions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/5931a75ae4bbd512a88c680b/tagDescriptions/tagId1"),
	// 	Properties: &armapimanagement.TagDescriptionContractProperties{
	// 		Description: to.Ptr("Some description that will be displayed for operation's tag if the tag is assigned to operation of the API"),
	// 		ExternalDocsDescription: to.Ptr("some additional info"),
	// 		ExternalDocsURL: to.Ptr("http://some_url.com"),
	// 		DisplayName: to.Ptr("tag1"),
	// 		TagID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/tags/tagId1"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteApiTagDescription.json
func ExampleAPITagDescriptionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAPITagDescriptionClient().Delete(ctx, "rg1", "apimService1", "59d5b28d1f7fab116c282650", "59d5b28e1f7fab116402044e", "*", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}