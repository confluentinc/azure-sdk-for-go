//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_List_MaximumSet_Gen.json
func ExamplePreRulesClient_NewListPager_preRulesListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPreRulesClient().NewListPager("lrs1", nil)
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
		// page.PreRulesResourceListResult = armpanngfw.PreRulesResourceListResult{
		// 	Value: []*armpanngfw.PreRulesResource{
		// 		{
		// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 			SystemData: &armpanngfw.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				CreatedBy: to.Ptr("praval"),
		// 				CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("praval"),
		// 				LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 			},
		// 			Properties: &armpanngfw.RuleEntry{
		// 				Description: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				ActionType: to.Ptr(armpanngfw.ActionEnumAllow),
		// 				Applications: []*string{
		// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		// 					AuditComment: to.Ptr("aaa"),
		// 					Category: &armpanngfw.Category{
		// 						Feeds: []*string{
		// 							to.Ptr("aaaaaaaaaaaa")},
		// 							URLCustom: []*string{
		// 								to.Ptr("aaaaa")},
		// 							},
		// 							DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
		// 							Destination: &armpanngfw.DestinationAddr{
		// 								Cidrs: []*string{
		// 									to.Ptr("aaaaaaa")},
		// 									Countries: []*string{
		// 										to.Ptr("aaaaaaaaaaaaaa")},
		// 										Feeds: []*string{
		// 											to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa")},
		// 											FqdnLists: []*string{
		// 												to.Ptr("aaaaaaaaaaaaa")},
		// 												PrefixLists: []*string{
		// 													to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
		// 												},
		// 												EnableLogging: to.Ptr(armpanngfw.StateEnumDISABLED),
		// 												Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 												InboundInspectionCertificate: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 												NegateDestination: to.Ptr(armpanngfw.BooleanEnumTRUE),
		// 												NegateSource: to.Ptr(armpanngfw.BooleanEnumTRUE),
		// 												Priority: to.Ptr[int32](24),
		// 												ProtocolPortList: []*string{
		// 													to.Ptr("aaaaaaaaaaaa")},
		// 													ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
		// 													RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 													RuleState: to.Ptr(armpanngfw.StateEnumDISABLED),
		// 													Source: &armpanngfw.SourceAddr{
		// 														Cidrs: []*string{
		// 															to.Ptr("aaa")},
		// 															Countries: []*string{
		// 																to.Ptr("aaaaa")},
		// 																Feeds: []*string{
		// 																	to.Ptr("aaaaaaaaaaaaaaaaaaa")},
		// 																	PrefixLists: []*string{
		// 																		to.Ptr("aaaaaaaaaaaaaaaaaaaa")},
		// 																	},
		// 																	Tags: []*armpanngfw.TagInfo{
		// 																		{
		// 																			Key: to.Ptr("keyName"),
		// 																			Value: to.Ptr("value"),
		// 																	}},
		// 																	Protocol: to.Ptr("aaaa"),
		// 																},
		// 														}},
		// 													}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_List_MinimumSet_Gen.json
func ExamplePreRulesClient_NewListPager_preRulesListMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPreRulesClient().NewListPager("lrs1", nil)
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
		// page.PreRulesResourceListResult = armpanngfw.PreRulesResourceListResult{
		// 	Value: []*armpanngfw.PreRulesResource{
		// 		{
		// 			ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/lrs1/prerules/1"),
		// 			Properties: &armpanngfw.RuleEntry{
		// 				RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_Get_MaximumSet_Gen.json
func ExamplePreRulesClient_Get_preRulesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPreRulesClient().Get(ctx, "lrs1", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PreRulesResource = armpanngfw.PreRulesResource{
	// 	Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 	SystemData: &armpanngfw.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		CreatedBy: to.Ptr("praval"),
	// 		CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("praval"),
	// 		LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 	},
	// 	Properties: &armpanngfw.RuleEntry{
	// 		Description: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		ActionType: to.Ptr(armpanngfw.ActionEnumAllow),
	// 		Applications: []*string{
	// 			to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 			AuditComment: to.Ptr("aaa"),
	// 			Category: &armpanngfw.Category{
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaaaaaaa")},
	// 					URLCustom: []*string{
	// 						to.Ptr("aaaaa")},
	// 					},
	// 					DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
	// 					Destination: &armpanngfw.DestinationAddr{
	// 						Cidrs: []*string{
	// 							to.Ptr("aaaaaaa")},
	// 							Countries: []*string{
	// 								to.Ptr("aaaaaaaaaaaaaa")},
	// 								Feeds: []*string{
	// 									to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 									FqdnLists: []*string{
	// 										to.Ptr("aaaaaaaaaaaaa")},
	// 										PrefixLists: []*string{
	// 											to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 										},
	// 										EnableLogging: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 										Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 										InboundInspectionCertificate: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 										NegateDestination: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 										NegateSource: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 										Priority: to.Ptr[int32](24),
	// 										ProtocolPortList: []*string{
	// 											to.Ptr("aaaaaaaaaaaa")},
	// 											ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
	// 											RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 											RuleState: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 											Source: &armpanngfw.SourceAddr{
	// 												Cidrs: []*string{
	// 													to.Ptr("aaa")},
	// 													Countries: []*string{
	// 														to.Ptr("aaaaa")},
	// 														Feeds: []*string{
	// 															to.Ptr("aaaaaaaaaaaaaaaaaaa")},
	// 															PrefixLists: []*string{
	// 																to.Ptr("aaaaaaaaaaaaaaaaaaaa")},
	// 															},
	// 															Tags: []*armpanngfw.TagInfo{
	// 																{
	// 																	Key: to.Ptr("keyName"),
	// 																	Value: to.Ptr("value"),
	// 															}},
	// 															Protocol: to.Ptr("aaaa"),
	// 														},
	// 													}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_Get_MinimumSet_Gen.json
func ExamplePreRulesClient_Get_preRulesGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPreRulesClient().Get(ctx, "lrs1", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PreRulesResource = armpanngfw.PreRulesResource{
	// 	ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/lrs1/prerules/1"),
	// 	Properties: &armpanngfw.RuleEntry{
	// 		RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_CreateOrUpdate_MaximumSet_Gen.json
func ExamplePreRulesClient_BeginCreateOrUpdate_preRulesCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPreRulesClient().BeginCreateOrUpdate(ctx, "lrs1", "1", armpanngfw.PreRulesResource{
		Properties: &armpanngfw.RuleEntry{
			Description: to.Ptr("description of pre rule"),
			ActionType:  to.Ptr(armpanngfw.ActionEnumAllow),
			Applications: []*string{
				to.Ptr("app1")},
			AuditComment: to.Ptr("example comment"),
			Category: &armpanngfw.Category{
				Feeds: []*string{
					to.Ptr("feed")},
				URLCustom: []*string{
					to.Ptr("https://microsoft.com")},
			},
			DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
			Destination: &armpanngfw.DestinationAddr{
				Cidrs: []*string{
					to.Ptr("1.0.0.1/10")},
				Countries: []*string{
					to.Ptr("India")},
				Feeds: []*string{
					to.Ptr("feed")},
				FqdnLists: []*string{
					to.Ptr("FQDN1")},
				PrefixLists: []*string{
					to.Ptr("PL1")},
			},
			EnableLogging:                to.Ptr(armpanngfw.StateEnumDISABLED),
			Etag:                         to.Ptr("c18e6eef-ba3e-49ee-8a85-2b36c863a9d0"),
			InboundInspectionCertificate: to.Ptr("cert1"),
			NegateDestination:            to.Ptr(armpanngfw.BooleanEnumTRUE),
			NegateSource:                 to.Ptr(armpanngfw.BooleanEnumTRUE),
			ProtocolPortList: []*string{
				to.Ptr("80")},
			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
			RuleName:          to.Ptr("preRule1"),
			RuleState:         to.Ptr(armpanngfw.StateEnumDISABLED),
			Source: &armpanngfw.SourceAddr{
				Cidrs: []*string{
					to.Ptr("1.0.0.1/10")},
				Countries: []*string{
					to.Ptr("India")},
				Feeds: []*string{
					to.Ptr("feed")},
				PrefixLists: []*string{
					to.Ptr("PL1")},
			},
			Tags: []*armpanngfw.TagInfo{
				{
					Key:   to.Ptr("keyName"),
					Value: to.Ptr("value"),
				}},
			Protocol: to.Ptr("HTTP"),
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
	// res.PreRulesResource = armpanngfw.PreRulesResource{
	// 	Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 	SystemData: &armpanngfw.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		CreatedBy: to.Ptr("praval"),
	// 		CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("praval"),
	// 		LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 	},
	// 	Properties: &armpanngfw.RuleEntry{
	// 		Description: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		ActionType: to.Ptr(armpanngfw.ActionEnumAllow),
	// 		Applications: []*string{
	// 			to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 			AuditComment: to.Ptr("aaa"),
	// 			Category: &armpanngfw.Category{
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaaaaaaa")},
	// 					URLCustom: []*string{
	// 						to.Ptr("aaaaa")},
	// 					},
	// 					DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
	// 					Destination: &armpanngfw.DestinationAddr{
	// 						Cidrs: []*string{
	// 							to.Ptr("aaaaaaa")},
	// 							Countries: []*string{
	// 								to.Ptr("aaaaaaaaaaaaaa")},
	// 								Feeds: []*string{
	// 									to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 									FqdnLists: []*string{
	// 										to.Ptr("aaaaaaaaaaaaa")},
	// 										PrefixLists: []*string{
	// 											to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 										},
	// 										EnableLogging: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 										Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 										InboundInspectionCertificate: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 										NegateDestination: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 										NegateSource: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 										Priority: to.Ptr[int32](24),
	// 										ProtocolPortList: []*string{
	// 											to.Ptr("aaaaaaaaaaaa")},
	// 											ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
	// 											RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 											RuleState: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 											Source: &armpanngfw.SourceAddr{
	// 												Cidrs: []*string{
	// 													to.Ptr("aaa")},
	// 													Countries: []*string{
	// 														to.Ptr("aaaaa")},
	// 														Feeds: []*string{
	// 															to.Ptr("aaaaaaaaaaaaaaaaaaa")},
	// 															PrefixLists: []*string{
	// 																to.Ptr("aaaaaaaaaaaaaaaaaaaa")},
	// 															},
	// 															Tags: []*armpanngfw.TagInfo{
	// 																{
	// 																	Key: to.Ptr("keyName"),
	// 																	Value: to.Ptr("value"),
	// 															}},
	// 															Protocol: to.Ptr("aaaa"),
	// 														},
	// 													}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_CreateOrUpdate_MinimumSet_Gen.json
func ExamplePreRulesClient_BeginCreateOrUpdate_preRulesCreateOrUpdateMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPreRulesClient().BeginCreateOrUpdate(ctx, "lrs1", "1", armpanngfw.PreRulesResource{
		Properties: &armpanngfw.RuleEntry{
			RuleName: to.Ptr("preRule1"),
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
	// res.PreRulesResource = armpanngfw.PreRulesResource{
	// 	ID: to.Ptr("/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks/lrs1/prerules/1"),
	// 	Properties: &armpanngfw.RuleEntry{
	// 		RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_Delete_MaximumSet_Gen.json
func ExamplePreRulesClient_BeginDelete_preRulesDeleteMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPreRulesClient().BeginDelete(ctx, "lrs1", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_Delete_MinimumSet_Gen.json
func ExamplePreRulesClient_BeginDelete_preRulesDeleteMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPreRulesClient().BeginDelete(ctx, "lrs1", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_getCounters_MaximumSet_Gen.json
func ExamplePreRulesClient_GetCounters_preRulesGetCountersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPreRulesClient().GetCounters(ctx, "lrs1", "1", &armpanngfw.PreRulesClientGetCountersOptions{FirewallName: to.Ptr("firewall1")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RuleCounter = armpanngfw.RuleCounter{
	// 	AppSeen: &armpanngfw.AppSeenData{
	// 		AppSeenList: []*armpanngfw.AppSeenInfo{
	// 			{
	// 				Category: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 				Risk: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				StandardPorts: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 				SubCategory: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 				Tag: to.Ptr("aaaaaaaaaa"),
	// 				Technology: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 				Title: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		}},
	// 		Count: to.Ptr[int32](13),
	// 	},
	// 	FirewallName: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 	HitCount: to.Ptr[int32](20),
	// 	LastUpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 	Priority: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 	RequestTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 	RuleListName: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 	RuleName: to.Ptr("aaaa"),
	// 	RuleStackName: to.Ptr("aaaaaaaaaaaaaaaaa"),
	// 	Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_getCounters_MinimumSet_Gen.json
func ExamplePreRulesClient_GetCounters_preRulesGetCountersMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPreRulesClient().GetCounters(ctx, "lrs1", "1", &armpanngfw.PreRulesClientGetCountersOptions{FirewallName: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RuleCounter = armpanngfw.RuleCounter{
	// 	Priority: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 	RuleName: to.Ptr("aaaa"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_refreshCounters_MaximumSet_Gen.json
func ExamplePreRulesClient_RefreshCounters_preRulesRefreshCountersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPreRulesClient().RefreshCounters(ctx, "lrs1", "1", &armpanngfw.PreRulesClientRefreshCountersOptions{FirewallName: to.Ptr("firewall1")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_refreshCounters_MinimumSet_Gen.json
func ExamplePreRulesClient_RefreshCounters_preRulesRefreshCountersMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPreRulesClient().RefreshCounters(ctx, "lrs1", "1", &armpanngfw.PreRulesClientRefreshCountersOptions{FirewallName: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_resetCounters_MaximumSet_Gen.json
func ExamplePreRulesClient_ResetCounters_preRulesResetCountersMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPreRulesClient().ResetCounters(ctx, "lrs1", "1", &armpanngfw.PreRulesClientResetCountersOptions{FirewallName: to.Ptr("firewall1")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RuleCounterReset = armpanngfw.RuleCounterReset{
	// 	FirewallName: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 	Priority: to.Ptr("aaaaaaa"),
	// 	RuleListName: to.Ptr("aaaaa"),
	// 	RuleName: to.Ptr("aaaaa"),
	// 	RuleStackName: to.Ptr("aa"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_resetCounters_MinimumSet_Gen.json
func ExamplePreRulesClient_ResetCounters_preRulesResetCountersMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPreRulesClient().ResetCounters(ctx, "lrs1", "1", &armpanngfw.PreRulesClientResetCountersOptions{FirewallName: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RuleCounterReset = armpanngfw.RuleCounterReset{
	// }
}