//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSectionsListByBillingProfile.json
func ExampleInvoiceSectionsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInvoiceSectionsClient().NewListByBillingProfilePager("{billingAccountName}", "{billingProfileName}", nil)
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
		// page.InvoiceSectionListResult = armbilling.InvoiceSectionListResult{
		// 	Value: []*armbilling.InvoiceSection{
		// 		{
		// 			Name: to.Ptr("22000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/22000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.InvoiceSectionProperties{
		// 				DisplayName: to.Ptr("invoiceSection1"),
		// 				Labels: map[string]*string{
		// 					"costCategory": to.Ptr("Support"),
		// 					"pcCode": to.Ptr("A123456"),
		// 				},
		// 				State: to.Ptr(armbilling.InvoiceSectionStateActive),
		// 				SystemID: to.Ptr("XX1X-XXAA-XXX-ZZZ"),
		// 				Tags: map[string]*string{
		// 					"costCategory": to.Ptr("Support"),
		// 					"pcCode": to.Ptr("A123456"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("22000000-0000-0000-0000-000000000011"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/22000000-0000-0000-0000-000000000011"),
		// 			Properties: &armbilling.InvoiceSectionProperties{
		// 				DisplayName: to.Ptr("invoiceSection2"),
		// 				Labels: map[string]*string{
		// 					"costCategory": to.Ptr("Marketing"),
		// 					"pcCode": to.Ptr("Z223456"),
		// 				},
		// 				State: to.Ptr(armbilling.InvoiceSectionStateRestricted),
		// 				SystemID: to.Ptr("YY1X-BBAA-XXX-ZZZ"),
		// 				Tags: map[string]*string{
		// 					"costCategory": to.Ptr("Marketing"),
		// 					"pcCode": to.Ptr("Z223456"),
		// 				},
		// 				TargetCloud: to.Ptr(armbilling.TargetCloudUSSec),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSection.json
func ExampleInvoiceSectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInvoiceSectionsClient().Get(ctx, "{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InvoiceSection = armbilling.InvoiceSection{
	// 	Name: to.Ptr("{invoiceSectionName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"),
	// 	Properties: &armbilling.InvoiceSectionProperties{
	// 		DisplayName: to.Ptr("invoiceSection1"),
	// 		Labels: map[string]*string{
	// 			"costCategory": to.Ptr("Support"),
	// 			"pcCode": to.Ptr("A123456"),
	// 		},
	// 		State: to.Ptr(armbilling.InvoiceSectionStateActive),
	// 		SystemID: to.Ptr("XX1X-XXAA-XXX-ZZZ"),
	// 		Tags: map[string]*string{
	// 			"costCategory": to.Ptr("Support"),
	// 			"pcCode": to.Ptr("A123456"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutInvoiceSection.json
func ExampleInvoiceSectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInvoiceSectionsClient().BeginCreateOrUpdate(ctx, "{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", armbilling.InvoiceSection{
		Properties: &armbilling.InvoiceSectionProperties{
			DisplayName: to.Ptr("invoiceSection1"),
			Labels: map[string]*string{
				"costCategory": to.Ptr("Support"),
				"pcCode":       to.Ptr("A123456"),
			},
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
	// res.InvoiceSection = armbilling.InvoiceSection{
	// 	Name: to.Ptr("{invoiceSectionName}"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}"),
	// 	Properties: &armbilling.InvoiceSectionProperties{
	// 		DisplayName: to.Ptr("invoiceSection1"),
	// 		Labels: map[string]*string{
	// 			"costCategory": to.Ptr("Support"),
	// 			"pcCode": to.Ptr("A123456"),
	// 		},
	// 		SystemID: to.Ptr("XX1X-XXAA-XXX-ZZZ"),
	// 	},
	// }
}
