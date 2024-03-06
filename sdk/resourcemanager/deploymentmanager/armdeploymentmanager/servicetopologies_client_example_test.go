//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeploymentmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_createorupdate.json
func ExampleServiceTopologiesClient_CreateOrUpdate_createATopologyWithArtifactSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServiceTopologiesClient().CreateOrUpdate(ctx, "myResourceGroup", "myTopology", armdeploymentmanager.ServiceTopologyResource{
		Location: to.Ptr("centralus"),
		Tags:     map[string]*string{},
		Properties: &armdeploymentmanager.ServiceTopologyResourceProperties{
			ArtifactSourceID: to.Ptr("Microsoft.DeploymentManager/artifactSources/myArtifactSource"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_createorupdate_noartifactsource.json
func ExampleServiceTopologiesClient_CreateOrUpdate_createATopologyWithoutArtifactSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServiceTopologiesClient().CreateOrUpdate(ctx, "myResourceGroup", "myTopology", armdeploymentmanager.ServiceTopologyResource{
		Location:   to.Ptr("centralus"),
		Tags:       map[string]*string{},
		Properties: &armdeploymentmanager.ServiceTopologyResourceProperties{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_get.json
func ExampleServiceTopologiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceTopologiesClient().Get(ctx, "myResourceGroup", "myTopology", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceTopologyResource = armdeploymentmanager.ServiceTopologyResource{
	// 	Name: to.Ptr("myTopology"),
	// 	Type: to.Ptr("Microsoft.DeploymentManager/serviceTopologies"),
	// 	Location: to.Ptr("centralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armdeploymentmanager.ServiceTopologyResourceProperties{
	// 		ArtifactSourceID: to.Ptr("Microsoft.DeploymentManager/artifactSources/myArtifactSource"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_delete.json
func ExampleServiceTopologiesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServiceTopologiesClient().Delete(ctx, "myResourceGroup", "myTopology", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopologies_list.json
func ExampleServiceTopologiesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentmanager.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceTopologiesClient().List(ctx, "myResourceGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServiceTopologyResourceArray = []*armdeploymentmanager.ServiceTopologyResource{
	// 	{
	// 		Name: to.Ptr("ContosoSvc1Topology"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/serviceTopologies"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.ServiceTopologyResourceProperties{
	// 			ArtifactSourceID: to.Ptr("Microsoft.DeploymentManager/artifactSources/contoso1ArtifactSource"),
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("ContosoSvc2Topology"),
	// 		Type: to.Ptr("Microsoft.DeploymentManager/serviceTopologies"),
	// 		Location: to.Ptr("centralus"),
	// 		Tags: map[string]*string{
	// 		},
	// 		Properties: &armdeploymentmanager.ServiceTopologyResourceProperties{
	// 			ArtifactSourceID: to.Ptr("Microsoft.DeploymentManager/artifactSources/contoso2ArtifactSource"),
	// 		},
	// }}
}
