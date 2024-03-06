//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhealthcareapis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/privatelink/WorkspaceListPrivateEndpointConnections.json
func ExampleWorkspacePrivateEndpointConnectionsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacePrivateEndpointConnectionsClient().NewListByWorkspacePager("testRG", "workspace1", nil)
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
		// page.PrivateEndpointConnectionListResultDescription = armhealthcareapis.PrivateEndpointConnectionListResultDescription{
		// 	Value: []*armhealthcareapis.PrivateEndpointConnectionDescription{
		// 		{
		// 			Name: to.Ptr("myConnection"),
		// 			Type: to.Ptr("Microsoft.HealthcareApis/workspaces/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/workspaces/workspace1/privateEndpointConnections/myConnection"),
		// 			Properties: &armhealthcareapis.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armhealthcareapis.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/peexample01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armhealthcareapis.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armhealthcareapis.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/privatelink/WorkspaceGetPrivateEndpointConnection.json
func ExampleWorkspacePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacePrivateEndpointConnectionsClient().Get(ctx, "testRG", "workspace1", "myConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionDescription = armhealthcareapis.PrivateEndpointConnectionDescription{
	// 	Name: to.Ptr("myConnection"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/workspaces/workspace1/privateEndpointConnections/myConnection"),
	// 	Properties: &armhealthcareapis.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armhealthcareapis.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/peexample01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armhealthcareapis.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armhealthcareapis.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/privatelink/WorkspaceCreatePrivateEndpointConnection.json
func ExampleWorkspacePrivateEndpointConnectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacePrivateEndpointConnectionsClient().BeginCreateOrUpdate(ctx, "testRG", "workspace1", "myConnection", armhealthcareapis.PrivateEndpointConnectionDescription{
		Properties: &armhealthcareapis.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armhealthcareapis.PrivateLinkServiceConnectionState{
				Description: to.Ptr("Auto-Approved"),
				Status:      to.Ptr(armhealthcareapis.PrivateEndpointServiceConnectionStatusApproved),
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
	// res.PrivateEndpointConnectionDescription = armhealthcareapis.PrivateEndpointConnectionDescription{
	// 	Name: to.Ptr("myConnection"),
	// 	Type: to.Ptr("Microsoft.HealthcareApis/workspaces/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.HealthcareApis/workspaces/workspace1/privateEndpointConnections/myConnection"),
	// 	Properties: &armhealthcareapis.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armhealthcareapis.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/peexample01"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armhealthcareapis.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("Auto-Approved"),
	// 			ActionsRequired: to.Ptr("None"),
	// 			Status: to.Ptr(armhealthcareapis.PrivateEndpointServiceConnectionStatusApproved),
	// 		},
	// 		ProvisioningState: to.Ptr(armhealthcareapis.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/privatelink/WorkspaceDeletePrivateEndpointConnection.json
func ExampleWorkspacePrivateEndpointConnectionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhealthcareapis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacePrivateEndpointConnectionsClient().BeginDelete(ctx, "testRG", "workspace1", "myConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
