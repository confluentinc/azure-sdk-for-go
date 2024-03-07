//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Create.json
func ExampleMoveCollectionsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMoveCollectionsClient().Create(ctx, "rg1", "movecollection1", &armresourcemover.MoveCollectionsClientCreateOptions{Body: &armresourcemover.MoveCollection{
		Identity: &armresourcemover.Identity{
			Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
		},
		Location: to.Ptr("eastus2"),
		Properties: &armresourcemover.MoveCollectionProperties{
			MoveType:     to.Ptr(armresourcemover.MoveTypeRegionToRegion),
			SourceRegion: to.Ptr("eastus"),
			TargetRegion: to.Ptr("westus"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MoveCollection = armresourcemover.MoveCollection{
	// 	Name: to.Ptr("movecollection1"),
	// 	Type: to.Ptr("Microsoft.Migrate/MoveCollections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1"),
	// 	Identity: &armresourcemover.Identity{
	// 		Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("7488fa50-1c8e-42c4-a117-38c8d05d8b72"),
	// 		TenantID: to.Ptr("cc7e5736-dbba-4059-a621-664eab9c1d80"),
	// 	},
	// 	Location: to.Ptr("United States"),
	// 	Properties: &armresourcemover.MoveCollectionProperties{
	// 		ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
	// 		SourceRegion: to.Ptr("eastus"),
	// 		TargetRegion: to.Ptr("westus"),
	// 		Version: to.Ptr("V1"),
	// 	},
	// 	SystemData: &armresourcemover.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Update.json
func ExampleMoveCollectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMoveCollectionsClient().Update(ctx, "rg1", "movecollection1", &armresourcemover.MoveCollectionsClientUpdateOptions{Body: &armresourcemover.UpdateMoveCollectionRequest{
		Identity: &armresourcemover.Identity{
			Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
		},
		Tags: map[string]*string{
			"key1": to.Ptr("mc1"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MoveCollection = armresourcemover.MoveCollection{
	// 	Name: to.Ptr("movecollection1"),
	// 	Type: to.Ptr("Microsoft.Migrate/MoveCollections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1"),
	// 	Identity: &armresourcemover.Identity{
	// 		Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("7488fa50-1c8e-42c4-a117-38c8d05d8b72"),
	// 		TenantID: to.Ptr("cc7e5736-dbba-4059-a621-664eab9c1d80"),
	// 	},
	// 	Location: to.Ptr("United States"),
	// 	Properties: &armresourcemover.MoveCollectionProperties{
	// 		MoveType: to.Ptr(armresourcemover.MoveTypeRegionToRegion),
	// 		ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
	// 		SourceRegion: to.Ptr("eastus"),
	// 		TargetRegion: to.Ptr("westus"),
	// 		Version: to.Ptr("V1"),
	// 	},
	// 	SystemData: &armresourcemover.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T15:06:54.275Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T17:16:24.364Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Delete.json
func ExampleMoveCollectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveCollectionsClient().BeginDelete(ctx, "rg1", "movecollection1", nil)
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Get.json
func ExampleMoveCollectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMoveCollectionsClient().Get(ctx, "rg1", "movecollection1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MoveCollection = armresourcemover.MoveCollection{
	// 	Name: to.Ptr("movecollection1"),
	// 	Type: to.Ptr("Microsoft.Migrate/MoveCollections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1"),
	// 	Identity: &armresourcemover.Identity{
	// 		Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("7488fa50-1c8e-42c4-a117-38c8d05d8b72"),
	// 		TenantID: to.Ptr("cc7e5736-dbba-4059-a621-664eab9c1d80"),
	// 	},
	// 	Location: to.Ptr("United States"),
	// 	Properties: &armresourcemover.MoveCollectionProperties{
	// 		MoveType: to.Ptr(armresourcemover.MoveTypeRegionToRegion),
	// 		ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
	// 		SourceRegion: to.Ptr("eastus"),
	// 		TargetRegion: to.Ptr("westus"),
	// 		Version: to.Ptr("V1"),
	// 	},
	// 	SystemData: &armresourcemover.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-01T15:06:54.275Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@microsoft.com"),
	// 		CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-27T17:16:24.364Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Prepare.json
func ExampleMoveCollectionsClient_BeginPrepare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveCollectionsClient().BeginPrepare(ctx, "rg1", "movecollection1", &armresourcemover.MoveCollectionsClientBeginPrepareOptions{Body: &armresourcemover.PrepareRequest{
		MoveResources: []*string{
			to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")},
		ValidateOnly: to.Ptr(false),
	},
	})
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_InitiateMove.json
func ExampleMoveCollectionsClient_BeginInitiateMove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveCollectionsClient().BeginInitiateMove(ctx, "rg1", "movecollection1", &armresourcemover.MoveCollectionsClientBeginInitiateMoveOptions{Body: &armresourcemover.ResourceMoveRequest{
		MoveResources: []*string{
			to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")},
		ValidateOnly: to.Ptr(false),
	},
	})
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Commit.json
func ExampleMoveCollectionsClient_BeginCommit() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveCollectionsClient().BeginCommit(ctx, "rg1", "movecollection1", &armresourcemover.MoveCollectionsClientBeginCommitOptions{Body: &armresourcemover.CommitRequest{
		MoveResources: []*string{
			to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")},
		ValidateOnly: to.Ptr(false),
	},
	})
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_Discard.json
func ExampleMoveCollectionsClient_BeginDiscard() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveCollectionsClient().BeginDiscard(ctx, "rg1", "movecollection1", &armresourcemover.MoveCollectionsClientBeginDiscardOptions{Body: &armresourcemover.DiscardRequest{
		MoveResources: []*string{
			to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")},
		ValidateOnly: to.Ptr(false),
	},
	})
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ResolveDependencies.json
func ExampleMoveCollectionsClient_BeginResolveDependencies() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveCollectionsClient().BeginResolveDependencies(ctx, "rg1", "movecollection1", nil)
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_BulkRemove.json
func ExampleMoveCollectionsClient_BeginBulkRemove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMoveCollectionsClient().BeginBulkRemove(ctx, "rg1", "movecollection1", &armresourcemover.MoveCollectionsClientBeginBulkRemoveOptions{Body: &armresourcemover.BulkRemoveRequest{
		MoveResources: []*string{
			to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1/MoveResources/moveresource1")},
		ValidateOnly: to.Ptr(false),
	},
	})
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
	// res.OperationStatus = armresourcemover.OperationStatus{
	// 	Name: to.Ptr("1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	EndTime: to.Ptr("6/17/2020 6:45:56 AM"),
	// 	ID: to.Ptr("/subscriptions/e80eb9fa-c996-4435-aa32-5af6f3d3077c/resourceGroups/RegionMoveRG-southcentralus-southeastasia/providers/Microsoft.Migrate/MoveCollections/MoveCollection-southcentralus-southeastasia/operations/1e4193c3-206e-4916-b124-1da16175eb0e"),
	// 	Properties: map[string]any{
	// 	},
	// 	StartTime: to.Ptr("6/17/2020 6:45:55 AM"),
	// 	Status: to.Ptr("Succeeded"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ListMoveCollectionsBySubscription.json
func ExampleMoveCollectionsClient_NewListMoveCollectionsBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMoveCollectionsClient().NewListMoveCollectionsBySubscriptionPager(nil)
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
		// page.MoveCollectionResultList = armresourcemover.MoveCollectionResultList{
		// 	Value: []*armresourcemover.MoveCollection{
		// 		{
		// 			Name: to.Ptr("movecollection1"),
		// 			Type: to.Ptr("Microsoft.Migrate/MoveCollections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1"),
		// 			Identity: &armresourcemover.Identity{
		// 				Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("7488fa50-1c8e-42c4-a117-38c8d05d8b72"),
		// 				TenantID: to.Ptr("cc7e5736-dbba-4059-a621-664eab9c1d80"),
		// 			},
		// 			Location: to.Ptr("United States"),
		// 			Properties: &armresourcemover.MoveCollectionProperties{
		// 				MoveType: to.Ptr(armresourcemover.MoveTypeRegionToRegion),
		// 				ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
		// 				SourceRegion: to.Ptr("eastus"),
		// 				TargetRegion: to.Ptr("westus"),
		// 				Version: to.Ptr("V1"),
		// 			},
		// 			SystemData: &armresourcemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/MoveCollections_ListMoveCollectionsByResourceGroup.json
func ExampleMoveCollectionsClient_NewListMoveCollectionsByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMoveCollectionsClient().NewListMoveCollectionsByResourceGroupPager("rg1", nil)
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
		// page.MoveCollectionResultList = armresourcemover.MoveCollectionResultList{
		// 	Value: []*armresourcemover.MoveCollection{
		// 		{
		// 			Name: to.Ptr("movecollection1"),
		// 			Type: to.Ptr("Microsoft.Migrate/MoveCollections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Migrate/MoveCollections/movecollection1"),
		// 			Identity: &armresourcemover.Identity{
		// 				Type: to.Ptr(armresourcemover.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("7488fa50-1c8e-42c4-a117-38c8d05d8b72"),
		// 				TenantID: to.Ptr("cc7e5736-dbba-4059-a621-664eab9c1d80"),
		// 			},
		// 			Location: to.Ptr("United States"),
		// 			Properties: &armresourcemover.MoveCollectionProperties{
		// 				MoveType: to.Ptr(armresourcemover.MoveTypeRegionToRegion),
		// 				ProvisioningState: to.Ptr(armresourcemover.ProvisioningStateSucceeded),
		// 				SourceRegion: to.Ptr("eastus"),
		// 				TargetRegion: to.Ptr("westus"),
		// 				Version: to.Ptr("V1"),
		// 			},
		// 			SystemData: &armresourcemover.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-29T15:06:54.275Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armresourcemover.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/RequiredFor_Get.json
func ExampleMoveCollectionsClient_ListRequiredFor() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresourcemover.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMoveCollectionsClient().ListRequiredFor(ctx, "rg1", "movecollection1", "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/nic1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RequiredForResourcesCollection = armresourcemover.RequiredForResourcesCollection{
	// 	SourceIDs: []*string{
	// 		to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/nic1")},
	// 	}
}