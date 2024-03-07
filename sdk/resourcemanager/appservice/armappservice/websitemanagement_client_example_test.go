//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListCustomHostNameSites.json
func ExampleWebSiteManagementClient_NewListCustomHostNameSitesPager_getCustomHostnamesUnderSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebSiteManagementClient().NewListCustomHostNameSitesPager(&armappservice.WebSiteManagementClientListCustomHostNameSitesOptions{Hostname: nil})
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
		// page.CustomHostnameSitesCollection = armappservice.CustomHostnameSitesCollection{
		// 	Value: []*armappservice.CustomHostnameSites{
		// 		{
		// 			Name: to.Ptr("mywebapp.azurewebsites.net"),
		// 			Type: to.Ptr("Microsoft.Web/customhostnameSites"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Web/customhostnameSites/mywebapp.azurewebsites.net"),
		// 			Properties: &armappservice.CustomHostnameSitesProperties{
		// 				CustomHostname: to.Ptr("mywebapp.azurewebsites.net"),
		// 				Region: to.Ptr("West US"),
		// 				SiteResourceIDs: []*armappservice.Identifier{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus-rg/providers/Microsoft.Web/sites/mywebapp"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("www.example.com"),
		// 			Type: to.Ptr("Microsoft.Web/customhostnameSites"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Web/customhostnameSites/www.example.com"),
		// 			Properties: &armappservice.CustomHostnameSitesProperties{
		// 				CustomHostname: to.Ptr("www.example.com"),
		// 				Region: to.Ptr("West US 2"),
		// 				SiteResourceIDs: []*armappservice.Identifier{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app2"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app3"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListCustomSpecificHostNameSites.json
func ExampleWebSiteManagementClient_NewListCustomHostNameSitesPager_getSpecificCustomHostnameUnderSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebSiteManagementClient().NewListCustomHostNameSitesPager(&armappservice.WebSiteManagementClientListCustomHostNameSitesOptions{Hostname: to.Ptr("www.example.com")})
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
		// page.CustomHostnameSitesCollection = armappservice.CustomHostnameSitesCollection{
		// 	Value: []*armappservice.CustomHostnameSites{
		// 		{
		// 			Name: to.Ptr("www.example.com"),
		// 			Type: to.Ptr("Microsoft.Web/customhostnameSites"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Web/customhostnameSites/www.example.com"),
		// 			Properties: &armappservice.CustomHostnameSitesProperties{
		// 				CustomHostname: to.Ptr("www.example.com"),
		// 				Region: to.Ptr("West US 2"),
		// 				SiteResourceIDs: []*armappservice.Identifier{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app1"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app2"),
		// 					},
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/westus2-rg/providers/Microsoft.Web/sites/westus2app3"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/ListAseRegions.json
func ExampleWebSiteManagementClient_NewListAseRegionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebSiteManagementClient().NewListAseRegionsPager(nil)
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
		// page.AseRegionCollection = armappservice.AseRegionCollection{
		// 	Value: []*armappservice.AseRegion{
		// 		{
		// 			Type: to.Ptr("Microsoft.Web/aseRegions"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/aseRegions"),
		// 			Properties: &armappservice.AseRegionProperties{
		// 				AvailableOS: []*string{
		// 					to.Ptr("Windows"),
		// 					to.Ptr("Linux"),
		// 					to.Ptr("HyperV")},
		// 					AvailableSKU: []*string{
		// 						to.Ptr("I1v2"),
		// 						to.Ptr("I2v2"),
		// 						to.Ptr("I3v2")},
		// 						DedicatedHost: to.Ptr(true),
		// 						DisplayName: to.Ptr("southcentralus"),
		// 						Standard: to.Ptr(true),
		// 						ZoneRedundant: to.Ptr(true),
		// 					},
		// 				},
		// 				{
		// 					Type: to.Ptr("Microsoft.Web/aseRegions"),
		// 					ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/aseRegions"),
		// 					Properties: &armappservice.AseRegionProperties{
		// 						AvailableOS: []*string{
		// 							to.Ptr("Windows"),
		// 							to.Ptr("Linux")},
		// 							AvailableSKU: []*string{
		// 								to.Ptr("I1v4"),
		// 								to.Ptr("I2v2"),
		// 								to.Ptr("I3v2")},
		// 								DedicatedHost: to.Ptr(true),
		// 								DisplayName: to.Ptr("northcentralus"),
		// 								Standard: to.Ptr(true),
		// 								ZoneRedundant: to.Ptr(true),
		// 							},
		// 					}},
		// 				}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/VerifyHostingEnvironmentVnet.json
func ExampleWebSiteManagementClient_VerifyHostingEnvironmentVnet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebSiteManagementClient().VerifyHostingEnvironmentVnet(ctx, armappservice.VnetParameters{
		Properties: &armappservice.VnetParametersProperties{
			VnetName:          to.Ptr("vNet123"),
			VnetResourceGroup: to.Ptr("vNet123rg"),
			VnetSubnetName:    to.Ptr("vNet123SubNet"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VnetValidationFailureDetails = armappservice.VnetValidationFailureDetails{
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/verifyHostingEnvironmentVnet"),
	// 	Properties: &armappservice.VnetValidationFailureDetailsProperties{
	// 		Failed: to.Ptr(false),
	// 		FailedTests: []*armappservice.VnetValidationTestFailure{
	// 		},
	// 	},
	// }
}