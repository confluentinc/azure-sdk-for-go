//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/ListTargets.json
func ExampleTargetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTargetsClient().NewListPager("exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", &armchaos.TargetsClientListOptions{ContinuationToken: nil})
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
		// page.TargetListResult = armchaos.TargetListResult{
		// 	Value: []*armchaos.Target{
		// 		{
		// 			Name: to.Ptr("Microsoft-Agent"),
		// 			Type: to.Ptr("Microsoft.Chaos/targets"),
		// 			ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-Agent"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: map[string]any{
		// 				"agentProfileId": "ac4e8251-fdc9-4277-8e87-dc57fe5794cf",
		// 				"identities": []any{
		// 					map[string]any{
		// 						"type": "CertificateSubjectIssuer",
		// 						"subject": "CN=example.subject",
		// 					},
		// 				},
		// 			},
		// 			SystemData: &armchaos.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/GetTarget.json
func ExampleTargetsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTargetsClient().Get(ctx, "exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", "Microsoft-Agent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Target = armchaos.Target{
	// 	Name: to.Ptr("Microsoft-Agent"),
	// 	Type: to.Ptr("Microsoft.Chaos/targets"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-Agent"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: map[string]any{
	// 		"agentProfileId": "ac4e8251-fdc9-4277-8e87-dc57fe5794cf",
	// 		"identities": []any{
	// 			map[string]any{
	// 				"type": "CertificateSubjectIssuer",
	// 				"subject": "CN=example.subject",
	// 			},
	// 		},
	// 	},
	// 	SystemData: &armchaos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/DeleteTarget.json
func ExampleTargetsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTargetsClient().Delete(ctx, "exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", "Microsoft-Agent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/chaos/resource-manager/Microsoft.Chaos/stable/2023-11-01/examples/CreateUpdateTarget.json
func ExampleTargetsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTargetsClient().CreateOrUpdate(ctx, "exampleRG", "Microsoft.Compute", "virtualMachines", "exampleVM", "Microsoft-Agent", armchaos.Target{
		Properties: map[string]any{
			"identities": []any{
				map[string]any{
					"type":    "CertificateSubjectIssuer",
					"subject": "CN=example.subject",
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Target = armchaos.Target{
	// 	Name: to.Ptr("Microsoft-Agent"),
	// 	Type: to.Ptr("Microsoft.Chaos/targets"),
	// 	ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Compute/virtualMachines/exampleVM/providers/Microsoft.Chaos/targets/Microsoft-Agent"),
	// 	Location: to.Ptr("centraluseuap"),
	// 	Properties: map[string]any{
	// 		"agentProfileId": "ac4e8251-fdc9-4277-8e87-dc57fe5794cf",
	// 		"identities": []any{
	// 			map[string]any{
	// 				"type": "CertificateSubjectIssuer",
	// 				"subject": "CN=example.subject",
	// 			},
	// 		},
	// 	},
	// 	SystemData: &armchaos.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.000Z"); return t}()),
	// 	},
	// }
}