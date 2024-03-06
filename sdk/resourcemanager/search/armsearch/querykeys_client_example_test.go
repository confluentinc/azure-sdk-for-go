//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7e29dd59eef13ef347d09e41a63f2585be77b3ca/specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchCreateQueryKey.json
func ExampleQueryKeysClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueryKeysClient().Create(ctx, "rg1", "mysearchservice", "Query key for browser-based clients", &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryKey = armsearch.QueryKey{
	// 	Name: to.Ptr("Query key for browser-based clients"),
	// 	Key: to.Ptr("<a query API key>"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7e29dd59eef13ef347d09e41a63f2585be77b3ca/specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchListQueryKeysBySearchService.json
func ExampleQueryKeysClient_NewListBySearchServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQueryKeysClient().NewListBySearchServicePager("rg1", "mysearchservice", &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
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
		// page.ListQueryKeysResult = armsearch.ListQueryKeysResult{
		// 	Value: []*armsearch.QueryKey{
		// 		{
		// 			Name: to.Ptr("Query key for browser-based clients"),
		// 			Key: to.Ptr("<a query API key>"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Query key for mobile clients"),
		// 			Key: to.Ptr("<another query API key>"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7e29dd59eef13ef347d09e41a63f2585be77b3ca/specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/SearchDeleteQueryKey.json
func ExampleQueryKeysClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewQueryKeysClient().Delete(ctx, "rg1", "mysearchservice", "<a query API key>", &armsearch.SearchManagementRequestOptions{ClientRequestID: nil}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
