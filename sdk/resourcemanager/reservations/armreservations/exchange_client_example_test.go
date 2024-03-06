//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armreservations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-11-01/examples/Exchange.json
func ExampleExchangeClient_BeginPost() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armreservations.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExchangeClient().BeginPost(ctx, armreservations.ExchangeRequest{
		Properties: &armreservations.ExchangeRequestProperties{
			SessionID: to.Ptr("66e2ac8f-439e-4345-8235-6fef07608081"),
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
	// res.ExchangeOperationResultResponse = armreservations.ExchangeOperationResultResponse{
	// 	Name: to.Ptr("4e2ffff7-b331-4fcb-ab11-b5fa49368188"),
	// 	ID: to.Ptr("/providers/microsoft.capacity/operationResults/4e2ffff7-b331-4fcb-ab11-b5fa49368188"),
	// 	Properties: &armreservations.ExchangeResponseProperties{
	// 		NetPayable: &armreservations.Price{
	// 			Amount: to.Ptr[float64](149254.8),
	// 			CurrencyCode: to.Ptr("USD"),
	// 		},
	// 		PolicyResult: &armreservations.ExchangePolicyErrors{
	// 		},
	// 		PurchasesTotal: &armreservations.Price{
	// 			Amount: to.Ptr[float64](153214.8),
	// 			CurrencyCode: to.Ptr("USD"),
	// 		},
	// 		RefundsTotal: &armreservations.Price{
	// 			Amount: to.Ptr[float64](3960),
	// 			CurrencyCode: to.Ptr("USD"),
	// 		},
	// 		ReservationsToExchange: []*armreservations.ReservationToReturnForExchange{
	// 			{
	// 				BillingInformation: &armreservations.BillingInformation{
	// 					BillingCurrencyProratedAmount: &armreservations.Price{
	// 						Amount: to.Ptr[float64](149254.8),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					BillingCurrencyRemainingCommitmentAmount: &armreservations.Price{
	// 						Amount: to.Ptr[float64](0),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 					BillingCurrencyTotalPaidAmount: &armreservations.Price{
	// 						Amount: to.Ptr[float64](153214.8),
	// 						CurrencyCode: to.Ptr("USD"),
	// 					},
	// 				},
	// 				BillingRefundAmount: &armreservations.Price{
	// 					Amount: to.Ptr[float64](3960),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 				Quantity: to.Ptr[int32](1),
	// 				ReservationID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1f14354c-dc12-4c8d-8090-6f295a3a34aa/reservations/c8c926bd-fc5d-4e29-9d43-b68340ac23a6"),
	// 				Status: to.Ptr(armreservations.OperationStatusSucceeded),
	// 		}},
	// 		ReservationsToPurchase: []*armreservations.ReservationToPurchaseExchange{
	// 			{
	// 				BillingCurrencyTotal: &armreservations.Price{
	// 					Amount: to.Ptr[float64](19800),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 				Properties: &armreservations.PurchaseRequest{
	// 					Location: to.Ptr("westus"),
	// 					Properties: &armreservations.PurchaseRequestProperties{
	// 						AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeShared),
	// 						BillingPlan: to.Ptr(armreservations.ReservationBillingPlanUpfront),
	// 						BillingScopeID: to.Ptr("/subscriptions/ed3a1871-612d-abcd-a849-c2542a68be83"),
	// 						DisplayName: to.Ptr("testDisplayName"),
	// 						Quantity: to.Ptr[int32](1),
	// 						Renew: to.Ptr(false),
	// 						ReservedResourceProperties: &armreservations.PurchaseRequestPropertiesReservedResourceProperties{
	// 							InstanceFlexibility: to.Ptr(armreservations.InstanceFlexibilityOn),
	// 						},
	// 						ReservedResourceType: to.Ptr(armreservations.ReservedResourceTypeVirtualMachines),
	// 						Term: to.Ptr(armreservations.ReservationTermP1Y),
	// 					},
	// 					SKU: &armreservations.SKUName{
	// 						Name: to.Ptr("Standard_B1ls"),
	// 					},
	// 				},
	// 				ReservationID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1e85c519-b815-4169-8d79-62fc460c608f/reservations/0c80fceb-305c-40a8-b5a6-11445807bbb3"),
	// 				ReservationOrderID: to.Ptr("/providers/microsoft.capacity/reservationOrders/1e85c519-b815-4169-8d79-62fc460c608f"),
	// 				Status: to.Ptr(armreservations.OperationStatusSucceeded),
	// 		}},
	// 		SavingsPlansToPurchase: []*armreservations.SavingsPlanToPurchaseExchange{
	// 			{
	// 				BillingCurrencyTotal: &armreservations.Price{
	// 					Amount: to.Ptr[float64](133414.8),
	// 					CurrencyCode: to.Ptr("USD"),
	// 				},
	// 				Properties: &armreservations.SavingsPlanPurchaseRequest{
	// 					Properties: &armreservations.SavingsPlanPurchaseRequestProperties{
	// 						AppliedScopeProperties: &armreservations.AppliedScopeProperties{
	// 							ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
	// 						},
	// 						AppliedScopeType: to.Ptr(armreservations.AppliedScopeTypeSingle),
	// 						BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
	// 						Commitment: &armreservations.Commitment{
	// 							Amount: to.Ptr[float64](15.23),
	// 							CurrencyCode: to.Ptr("USD"),
	// 							Grain: to.Ptr(armreservations.CommitmentGrainHourly),
	// 						},
	// 						DisplayName: to.Ptr("ComputeSavingsPlan"),
	// 						Term: to.Ptr(armreservations.SavingsPlanTermP1Y),
	// 					},
	// 					SKU: &armreservations.SKUName{
	// 						Name: to.Ptr("Compute_Savings_Plan"),
	// 					},
	// 				},
	// 				SavingsPlanID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
	// 				SavingsPlanOrderID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000"),
	// 				Status: to.Ptr(armreservations.OperationStatusSucceeded),
	// 		}},
	// 		SessionID: to.Ptr("66e2ac8f-439e-4345-8235-6fef07608081"),
	// 	},
	// 	Status: to.Ptr(armreservations.ExchangeOperationResultStatusSucceeded),
	// }
}
