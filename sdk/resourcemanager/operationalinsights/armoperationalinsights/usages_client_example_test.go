//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesListUsages.json
func ExampleUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListPager("rg1", "TestLinkWS", nil)
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
		// page.WorkspaceListUsagesResult = armoperationalinsights.WorkspaceListUsagesResult{
		// 	Value: []*armoperationalinsights.UsageMetric{
		// 		{
		// 			Name: &armoperationalinsights.MetricName{
		// 				LocalizedValue: to.Ptr("Data Analyzed"),
		// 				Value: to.Ptr("DataAnalyzed"),
		// 			},
		// 			CurrentValue: to.Ptr[float64](0),
		// 			Limit: to.Ptr[float64](524288000),
		// 			NextResetTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-03T00:00:00.000Z"); return t}()),
		// 			QuotaPeriod: to.Ptr("P1D"),
		// 			Unit: to.Ptr("Bytes"),
		// 	}},
		// }
	}
}
