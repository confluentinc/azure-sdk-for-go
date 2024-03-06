//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/NetworkInterfaceLoadBalancerList.json
func ExampleInterfaceLoadBalancersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInterfaceLoadBalancersClient().NewListPager("testrg", "nic1", nil)
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
		// page.InterfaceLoadBalancerListResult = armnetwork.InterfaceLoadBalancerListResult{
		// 	Value: []*armnetwork.LoadBalancer{
		// 		{
		// 			Name: to.Ptr("lbname1"),
		// 			Type: to.Ptr("Microsoft.Network/loadBalancers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 			Properties: &armnetwork.LoadBalancerPropertiesFormat{
		// 				BackendAddressPools: []*armnetwork.BackendAddressPool{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/backendAddressPools/bepool1"),
		// 						Name: to.Ptr("bepool1"),
		// 						Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 						Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
		// 							BackendIPConfigurations: []*armnetwork.InterfaceIPConfiguration{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
		// 							}},
		// 							LoadBalancingRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/loadBalancingRules/rule1"),
		// 							}},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/frontendIPConfigurations/lbfrontend"),
		// 						Name: to.Ptr("lbfrontend"),
		// 						Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 						Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
		// 							InboundNatRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/inboundNatRules/inbound1"),
		// 							}},
		// 							LoadBalancingRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/loadBalancingRules/rule1"),
		// 							}},
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							PublicIPAddress: &armnetwork.PublicIPAddress{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/publicIPAddresses/myDynamicPublicIP"),
		// 							},
		// 						},
		// 				}},
		// 				InboundNatPools: []*armnetwork.InboundNatPool{
		// 				},
		// 				InboundNatRules: []*armnetwork.InboundNatRule{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/inboundNatRules/inbound1"),
		// 						Name: to.Ptr("inbound1"),
		// 						Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 						Properties: &armnetwork.InboundNatRulePropertiesFormat{
		// 							BackendIPConfiguration: &armnetwork.InterfaceIPConfiguration{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/nic1/ipConfigurations/ipconfig1"),
		// 							},
		// 							BackendPort: to.Ptr[int32](3389),
		// 							EnableFloatingIP: to.Ptr(false),
		// 							FrontendIPConfiguration: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/frontendIPConfigurations/lbfrontend"),
		// 							},
		// 							FrontendPort: to.Ptr[int32](3389),
		// 							IdleTimeoutInMinutes: to.Ptr[int32](15),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 						},
		// 				}},
		// 				LoadBalancingRules: []*armnetwork.LoadBalancingRule{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/loadBalancingRules/rule1"),
		// 						Name: to.Ptr("rule1"),
		// 						Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 						Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
		// 							BackendAddressPool: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/backendAddressPools/bepool1"),
		// 							},
		// 							BackendPort: to.Ptr[int32](80),
		// 							EnableFloatingIP: to.Ptr(false),
		// 							FrontendIPConfiguration: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/frontendIPConfigurations/lbfrontend"),
		// 							},
		// 							FrontendPort: to.Ptr[int32](80),
		// 							IdleTimeoutInMinutes: to.Ptr[int32](15),
		// 							LoadDistribution: to.Ptr(armnetwork.LoadDistributionDefault),
		// 							Probe: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/probes/probe1"),
		// 							},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 						},
		// 				}},
		// 				OutboundRules: []*armnetwork.OutboundRule{
		// 				},
		// 				Probes: []*armnetwork.Probe{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/probes/probe1"),
		// 						Name: to.Ptr("probe1"),
		// 						Etag: to.Ptr("W/\\\"00000000-0000-0000-0000-000000000000\\\""),
		// 						Properties: &armnetwork.ProbePropertiesFormat{
		// 							IntervalInSeconds: to.Ptr[int32](15),
		// 							LoadBalancingRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/loadBalancers/lbname1/loadBalancingRules/rule1"),
		// 							}},
		// 							NumberOfProbes: to.Ptr[int32](2),
		// 							Port: to.Ptr[int32](80),
		// 							ProbeThreshold: to.Ptr[int32](1),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							RequestPath: to.Ptr("healthcheck.aspx"),
		// 							Protocol: to.Ptr(armnetwork.ProbeProtocolHTTP),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			},
		// 	}},
		// }
	}
}
