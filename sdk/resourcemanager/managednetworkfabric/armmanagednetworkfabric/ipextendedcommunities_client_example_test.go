//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetworkfabric/armmanagednetworkfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Create_MaximumSet_Gen.json
func ExampleIPExtendedCommunitiesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPExtendedCommunitiesClient().BeginCreate(ctx, "example-rg", "example-ipExtendedCommunity", armmanagednetworkfabric.IPExtendedCommunity{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"keyID": to.Ptr("KeyValue"),
		},
		Properties: &armmanagednetworkfabric.IPExtendedCommunityProperties{
			Annotation: to.Ptr("annotation"),
			IPExtendedCommunityRules: []*armmanagednetworkfabric.IPExtendedCommunityRule{
				{
					Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
					RouteTargets: []*string{
						to.Ptr("1234:2345")},
					SequenceNumber: to.Ptr[int64](4155123341),
				}},
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
	// res.IPExtendedCommunity = armmanagednetworkfabric.IPExtendedCommunity{
	// 	Name: to.Ptr("example-ipExtendedCommunity"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/ipExtendedCommunities"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.IPExtendedCommunityProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		IPExtendedCommunityRules: []*armmanagednetworkfabric.IPExtendedCommunityRule{
	// 			{
	// 				Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 				RouteTargets: []*string{
	// 					to.Ptr("1234:2345")},
	// 					SequenceNumber: to.Ptr[int64](4155123341),
	// 			}},
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Get_MaximumSet_Gen.json
func ExampleIPExtendedCommunitiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIPExtendedCommunitiesClient().Get(ctx, "example-rg", "example-ipExtendedCommunity", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IPExtendedCommunity = armmanagednetworkfabric.IPExtendedCommunity{
	// 	Name: to.Ptr("example-ipExtendedCommunity"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/ipExtendedCommunities"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.IPExtendedCommunityProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		IPExtendedCommunityRules: []*armmanagednetworkfabric.IPExtendedCommunityRule{
	// 			{
	// 				Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 				RouteTargets: []*string{
	// 					to.Ptr("1234:2345")},
	// 					SequenceNumber: to.Ptr[int64](4155123341),
	// 			}},
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Update_MaximumSet_Gen.json
func ExampleIPExtendedCommunitiesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPExtendedCommunitiesClient().BeginUpdate(ctx, "example-rg", "example-ipExtendedCommunity", armmanagednetworkfabric.IPExtendedCommunityPatch{
		Tags: map[string]*string{
			"keyID": to.Ptr("KeyValue"),
		},
		Properties: &armmanagednetworkfabric.IPExtendedCommunityPatchProperties{
			IPExtendedCommunityRules: []*armmanagednetworkfabric.IPExtendedCommunityRule{
				{
					Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
					RouteTargets: []*string{
						to.Ptr("1234:2345")},
					SequenceNumber: to.Ptr[int64](4155123341),
				}},
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
	// res.IPExtendedCommunity = armmanagednetworkfabric.IPExtendedCommunity{
	// 	Name: to.Ptr("example-ipExtendedCommunity"),
	// 	Type: to.Ptr("microsoft.managednetworkfabric/ipExtendedCommunities"),
	// 	ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
	// 	SystemData: &armmanagednetworkfabric.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
	// 		CreatedBy: to.Ptr("email@address.com"),
	// 		CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@mail.com"),
	// 		LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"keyID": to.Ptr("KeyValue"),
	// 	},
	// 	Properties: &armmanagednetworkfabric.IPExtendedCommunityProperties{
	// 		Annotation: to.Ptr("annotation"),
	// 		IPExtendedCommunityRules: []*armmanagednetworkfabric.IPExtendedCommunityRule{
	// 			{
	// 				Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
	// 				RouteTargets: []*string{
	// 					to.Ptr("1234:2345")},
	// 					SequenceNumber: to.Ptr[int64](4155123341),
	// 			}},
	// 			AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
	// 			ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
	// 			ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Delete_MaximumSet_Gen.json
func ExampleIPExtendedCommunitiesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewIPExtendedCommunitiesClient().BeginDelete(ctx, "example-rg", "example-ipExtendedCommunity", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_ListByResourceGroup_MaximumSet_Gen.json
func ExampleIPExtendedCommunitiesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPExtendedCommunitiesClient().NewListByResourceGroupPager("example-rg", nil)
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
		// page.IPExtendedCommunityListResult = armmanagednetworkfabric.IPExtendedCommunityListResult{
		// 	Value: []*armmanagednetworkfabric.IPExtendedCommunity{
		// 		{
		// 			Name: to.Ptr("example-ipExtendedCommunity"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/ipExtendedCommunities"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@mail.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"keyID": to.Ptr("KeyValue"),
		// 			},
		// 			Properties: &armmanagednetworkfabric.IPExtendedCommunityProperties{
		// 				Annotation: to.Ptr("annotation"),
		// 				IPExtendedCommunityRules: []*armmanagednetworkfabric.IPExtendedCommunityRule{
		// 					{
		// 						Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
		// 						RouteTargets: []*string{
		// 							to.Ptr("1234:2345")},
		// 							SequenceNumber: to.Ptr[int64](4155123341),
		// 					}},
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/925ba149e17454ce91ecd3f9f4134effb2f97844/specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_ListBySubscription_MaximumSet_Gen.json
func ExampleIPExtendedCommunitiesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagednetworkfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIPExtendedCommunitiesClient().NewListBySubscriptionPager(nil)
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
		// page.IPExtendedCommunityListResult = armmanagednetworkfabric.IPExtendedCommunityListResult{
		// 	Value: []*armmanagednetworkfabric.IPExtendedCommunity{
		// 		{
		// 			Name: to.Ptr("example-ipExtendedCommunity"),
		// 			Type: to.Ptr("microsoft.managednetworkfabric/ipExtendedCommunities"),
		// 			ID: to.Ptr("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/ipExtendedCommunities/example-ipExtendedCommunity"),
		// 			SystemData: &armmanagednetworkfabric.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
		// 				CreatedBy: to.Ptr("email@address.com"),
		// 				CreatedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-11T18:49:33.904Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@mail.com"),
		// 				LastModifiedByType: to.Ptr(armmanagednetworkfabric.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"keyID": to.Ptr("KeyValue"),
		// 			},
		// 			Properties: &armmanagednetworkfabric.IPExtendedCommunityProperties{
		// 				Annotation: to.Ptr("annotation"),
		// 				IPExtendedCommunityRules: []*armmanagednetworkfabric.IPExtendedCommunityRule{
		// 					{
		// 						Action: to.Ptr(armmanagednetworkfabric.CommunityActionTypesPermit),
		// 						RouteTargets: []*string{
		// 							to.Ptr("1234:2345")},
		// 							SequenceNumber: to.Ptr[int64](4155123341),
		// 					}},
		// 					AdministrativeState: to.Ptr(armmanagednetworkfabric.AdministrativeStateEnabled),
		// 					ConfigurationState: to.Ptr(armmanagednetworkfabric.ConfigurationStateSucceeded),
		// 					ProvisioningState: to.Ptr(armmanagednetworkfabric.ProvisioningStateSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}