//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhdinsight_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsight/armhdinsight"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetAllPrivateLinkResourcesInCluster.json
func ExamplePrivateLinkResourcesClient_ListByCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().ListByCluster(ctx, "rg1", "cluster1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourceListResult = armhdinsight.PrivateLinkResourceListResult{
	// 	Value: []*armhdinsight.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("gateway"),
	// 			Type: to.Ptr("Microsoft.HDInsight/clusters/privateLinkResources"),
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1/privateLinkResources/gateway"),
	// 			Properties: &armhdinsight.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("gateway"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("gateway")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.azurehdinsight.net")},
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("headnode"),
	// 					Type: to.Ptr("Microsoft.HDInsight/clusters/privateLinkResources"),
	// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1/privateLinkResources/headnode"),
	// 					Properties: &armhdinsight.PrivateLinkResourceProperties{
	// 						GroupID: to.Ptr("headnode"),
	// 						RequiredMembers: []*string{
	// 							to.Ptr("headnode")},
	// 							RequiredZoneNames: []*string{
	// 								to.Ptr("privatelink.azurehdinsight.net")},
	// 							},
	// 					}},
	// 				}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetPrivateLinkResource.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsight.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "rg1", "cluster1", "gateway", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armhdinsight.PrivateLinkResource{
	// 	Name: to.Ptr("gateway"),
	// 	Type: to.Ptr("Microsoft.HDInsight/clusters/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HDInsight/clusters/cluster1/privateLinkResources/gateway"),
	// 	Properties: &armhdinsight.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("gateway"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("gateway")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.azurehdinsight.net")},
	// 			},
	// 		}
}