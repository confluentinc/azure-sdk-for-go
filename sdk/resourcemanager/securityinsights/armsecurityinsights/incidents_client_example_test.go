//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/GetIncidents.json
func ExampleIncidentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIncidentsClient().NewListPager("myRg", "myWorkspace", &armsecurityinsights.IncidentsClientListOptions{Filter: nil,
		Orderby:   to.Ptr("properties/createdTimeUtc desc"),
		Top:       to.Ptr[int32](1),
		SkipToken: nil,
	})
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
		// page.IncidentList = armsecurityinsights.IncidentList{
		// 	Value: []*armsecurityinsights.Incident{
		// 		{
		// 			Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Type: to.Ptr("Microsoft.SecurityInsights/incidents"),
		// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 			Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		// 			Properties: &armsecurityinsights.IncidentProperties{
		// 				Description: to.Ptr("This is a demo incident"),
		// 				AdditionalData: &armsecurityinsights.IncidentAdditionalData{
		// 					AlertProductNames: []*string{
		// 					},
		// 					AlertsCount: to.Ptr[int32](0),
		// 					BookmarksCount: to.Ptr[int32](0),
		// 					CommentsCount: to.Ptr[int32](3),
		// 					Tactics: []*armsecurityinsights.AttackTactic{
		// 						to.Ptr(armsecurityinsights.AttackTacticPersistence)},
		// 					},
		// 					Classification: to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
		// 					ClassificationComment: to.Ptr("Not a malicious activity"),
		// 					ClassificationReason: to.Ptr(armsecurityinsights.IncidentClassificationReasonIncorrectAlertLogic),
		// 					CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
		// 					FirstActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30.000Z"); return t}()),
		// 					IncidentNumber: to.Ptr[int32](3177),
		// 					IncidentURL: to.Ptr("https://portal.azure.com/#asset/Microsoft_Azure_Security_Insights/Incident/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
		// 					Labels: []*armsecurityinsights.IncidentLabel{
		// 					},
		// 					LastActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30.000Z"); return t}()),
		// 					LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
		// 					Owner: &armsecurityinsights.IncidentOwnerInfo{
		// 						AssignedTo: to.Ptr("john doe"),
		// 						Email: to.Ptr("john.doe@contoso.com"),
		// 						ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
		// 						UserPrincipalName: to.Ptr("john@contoso.com"),
		// 					},
		// 					RelatedAnalyticRuleIDs: []*string{
		// 						to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/fab3d2d4-747f-46a7-8ef0-9c0be8112bf7"),
		// 						to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/8deb8303-e94d-46ff-96e0-5fd94b33df1a")},
		// 						Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
		// 						Status: to.Ptr(armsecurityinsights.IncidentStatusClosed),
		// 						Title: to.Ptr("My incident"),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/GetIncidentById.json
func ExampleIncidentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().Get(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Incident = armsecurityinsights.Incident{
	// 	Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/incidents"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 	Properties: &armsecurityinsights.IncidentProperties{
	// 		Description: to.Ptr("This is a demo incident"),
	// 		AdditionalData: &armsecurityinsights.IncidentAdditionalData{
	// 			AlertProductNames: []*string{
	// 			},
	// 			AlertsCount: to.Ptr[int32](0),
	// 			BookmarksCount: to.Ptr[int32](0),
	// 			CommentsCount: to.Ptr[int32](3),
	// 			Tactics: []*armsecurityinsights.AttackTactic{
	// 				to.Ptr(armsecurityinsights.AttackTacticInitialAccess),
	// 				to.Ptr(armsecurityinsights.AttackTacticPersistence)},
	// 			},
	// 			Classification: to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
	// 			ClassificationComment: to.Ptr("Not a malicious activity"),
	// 			ClassificationReason: to.Ptr(armsecurityinsights.IncidentClassificationReasonInaccurateData),
	// 			CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
	// 			FirstActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30.000Z"); return t}()),
	// 			IncidentNumber: to.Ptr[int32](3177),
	// 			IncidentURL: to.Ptr("https://portal.azure.com/#asset/Microsoft_Azure_Security_Insights/Incident/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 			Labels: []*armsecurityinsights.IncidentLabel{
	// 			},
	// 			LastActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30.000Z"); return t}()),
	// 			LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
	// 			Owner: &armsecurityinsights.IncidentOwnerInfo{
	// 				AssignedTo: to.Ptr("john doe"),
	// 				Email: to.Ptr("john.doe@contoso.com"),
	// 				ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 				UserPrincipalName: to.Ptr("john@contoso.com"),
	// 			},
	// 			RelatedAnalyticRuleIDs: []*string{
	// 				to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/fab3d2d4-747f-46a7-8ef0-9c0be8112bf7")},
	// 				Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
	// 				Status: to.Ptr(armsecurityinsights.IncidentStatusClosed),
	// 				Title: to.Ptr("My incident"),
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/CreateIncident.json
func ExampleIncidentsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", armsecurityinsights.Incident{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Properties: &armsecurityinsights.IncidentProperties{
			Description:           to.Ptr("This is a demo incident"),
			Classification:        to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
			ClassificationComment: to.Ptr("Not a malicious activity"),
			ClassificationReason:  to.Ptr(armsecurityinsights.IncidentClassificationReasonIncorrectAlertLogic),
			FirstActivityTimeUTC:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30.000Z"); return t }()),
			LastActivityTimeUTC:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30.000Z"); return t }()),
			Owner: &armsecurityinsights.IncidentOwnerInfo{
				ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
			},
			Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
			Status:   to.Ptr(armsecurityinsights.IncidentStatusClosed),
			Title:    to.Ptr("My incident"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Incident = armsecurityinsights.Incident{
	// 	Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Type: to.Ptr("Microsoft.SecurityInsights/incidents"),
	// 	ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 	Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0001\""),
	// 	Properties: &armsecurityinsights.IncidentProperties{
	// 		Description: to.Ptr("This is a demo incident"),
	// 		AdditionalData: &armsecurityinsights.IncidentAdditionalData{
	// 			AlertProductNames: []*string{
	// 			},
	// 			AlertsCount: to.Ptr[int32](0),
	// 			BookmarksCount: to.Ptr[int32](0),
	// 			CommentsCount: to.Ptr[int32](3),
	// 			Tactics: []*armsecurityinsights.AttackTactic{
	// 			},
	// 		},
	// 		Classification: to.Ptr(armsecurityinsights.IncidentClassificationFalsePositive),
	// 		ClassificationComment: to.Ptr("Not a malicious activity"),
	// 		ClassificationReason: to.Ptr(armsecurityinsights.IncidentClassificationReasonIncorrectAlertLogic),
	// 		CreatedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
	// 		FirstActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:00:30.000Z"); return t}()),
	// 		IncidentNumber: to.Ptr[int32](3177),
	// 		IncidentURL: to.Ptr("https://portal.azure.com/#asset/Microsoft_Azure_Security_Insights/Incident/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/incidents/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Labels: []*armsecurityinsights.IncidentLabel{
	// 		},
	// 		LastActivityTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:05:30.000Z"); return t}()),
	// 		LastModifiedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-01-01T13:15:30.000Z"); return t}()),
	// 		Owner: &armsecurityinsights.IncidentOwnerInfo{
	// 			AssignedTo: to.Ptr("john doe"),
	// 			Email: to.Ptr("john.doe@contoso.com"),
	// 			ObjectID: to.Ptr("2046feea-040d-4a46-9e2b-91c2941bfa70"),
	// 			UserPrincipalName: to.Ptr("john@contoso.com"),
	// 		},
	// 		RelatedAnalyticRuleIDs: []*string{
	// 		},
	// 		Severity: to.Ptr(armsecurityinsights.IncidentSeverityHigh),
	// 		Status: to.Ptr(armsecurityinsights.IncidentStatusClosed),
	// 		Title: to.Ptr("My incident"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/DeleteIncident.json
func ExampleIncidentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIncidentsClient().Delete(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/GetAllIncidentAlerts.json
func ExampleIncidentsClient_ListAlerts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().ListAlerts(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IncidentAlertList = armsecurityinsights.IncidentAlertList{
	// 	Value: []*armsecurityinsights.SecurityAlert{
	// 		{
	// 			Name: to.Ptr("baa8a239-6fde-4ab7-a093-d09f7b75c58c"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/bd794837-4d29-4647-9105-6339bfdb4e6a/resourceGroups/myRG/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Entities/baa8a239-6fde-4ab7-a093-d09f7b75c58c"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindEnumSecurityAlert),
	// 			Properties: &armsecurityinsights.SecurityAlertProperties{
	// 				AdditionalData: map[string]any{
	// 					"AlertMessageEnqueueTime": "2020-07-20T18:21:57.304Z",
	// 				},
	// 				FriendlyName: to.Ptr("myAlert"),
	// 				AlertDisplayName: to.Ptr("myAlert"),
	// 				AlertType: to.Ptr("myAlert"),
	// 				ConfidenceLevel: to.Ptr(armsecurityinsights.ConfidenceLevelUnknown),
	// 				EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				ProcessingEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				ProductName: to.Ptr("Azure Security Center"),
	// 				ResourceIdentifiers: []any{
	// 					map[string]any{
	// 						"type": "LogAnalytics",
	// 						"resourceGroup": "myRG",
	// 						"subscriptionId": "bd794837-4d29-4647-9105-6339bfdb4e6a",
	// 						"workspaceId": "c8c99641-985d-4e4e-8e91-fb3466cd0e5b",
	// 				}},
	// 				Severity: to.Ptr(armsecurityinsights.AlertSeverityLow),
	// 				StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				Status: to.Ptr(armsecurityinsights.AlertStatusNew),
	// 				SystemAlertID: to.Ptr("baa8a239-6fde-4ab7-a093-d09f7b75c58c"),
	// 				Tactics: []*armsecurityinsights.AttackTactic{
	// 				},
	// 				TimeGenerated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-07-20T18:21:53.615Z"); return t}()),
	// 				VendorName: to.Ptr("Microsoft"),
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/GetAllIncidentBookmarks.json
func ExampleIncidentsClient_ListBookmarks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().ListBookmarks(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IncidentBookmarkList = armsecurityinsights.IncidentBookmarkList{
	// 	Value: []*armsecurityinsights.HuntingBookmark{
	// 		{
	// 			Name: to.Ptr("afbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/afbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindEnumBookmark),
	// 			Properties: &armsecurityinsights.HuntingBookmarkProperties{
	// 				AdditionalData: map[string]any{
	// 					"ETag": "\"3b00acab-0000-0d00-0000-5f15e4ed0000\"",
	// 					"EntityId": "afbd324f-6c48-459c-8710-8d1e1cd03812",
	// 				},
	// 				FriendlyName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				CreatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("b03ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 				DisplayName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				Labels: []*string{
	// 				},
	// 				Query: to.Ptr("SecurityEvent\r\n| take 1\n"),
	// 				QueryResult: to.Ptr("{\"TimeGenerated\":\"2020-05-24T01:24:25.67Z\",\"Account\":\"\\\\ADMINISTRATOR\",\"AccountType\":\"User\",\"Computer\":\"SecurityEvents\",\"EventSourceName\":\"Microsoft-Windows-Security-Auditing\",\"Channel\":\"Security\",\"Task\":12544,\"Level\":\"16\",\"EventID\":4625,\"Activity\":\"4625 - An account failed to log on.\",\"AuthenticationPackageName\":\"NTLM\",\"FailureReason\":\"%%2313\",\"IpAddress\":\"176.113.115.73\",\"IpPort\":\"0\",\"LmPackageName\":\"-\",\"LogonProcessName\":\"NtLmSsp \",\"LogonType\":3,\"LogonTypeName\":\"3 - Network\",\"Process\":\"-\",\"ProcessId\":\"0x0\",\"__entityMapping\":{\"\\\\ADMINISTRATOR\":\"Account\",\"SecurityEvents\":\"Host\"}}"),
	// 				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				UpdatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("b03ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("bbbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/bookmarks/bbbd324f-6c48-459c-8710-8d1e1cd03812"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindEnumBookmark),
	// 			Properties: &armsecurityinsights.HuntingBookmarkProperties{
	// 				AdditionalData: map[string]any{
	// 					"ETag": "\"3b00acab-0000-0d00-0000-5f15e4ed0000\"",
	// 					"EntityId": "afbd324f-6c48-459c-8710-8d1e1cd03812",
	// 				},
	// 				FriendlyName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				CreatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("303ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 				DisplayName: to.Ptr("SecurityEvent - 868f40f4698d"),
	// 				EventTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				Labels: []*string{
	// 				},
	// 				Query: to.Ptr("SecurityEvent\r\n| take 1\n"),
	// 				QueryResult: to.Ptr("{\"TimeGenerated\":\"2020-05-24T01:24:25.67Z\",\"Account\":\"\\\\ADMINISTRATOR\",\"AccountType\":\"User\",\"Computer\":\"SecurityEvents\",\"EventSourceName\":\"Microsoft-Windows-Security-Auditing\",\"Channel\":\"Security\",\"Task\":12544,\"Level\":\"16\",\"EventID\":4625,\"Activity\":\"4625 - An account failed to log on.\",\"AuthenticationPackageName\":\"NTLM\",\"FailureReason\":\"%%2313\",\"IpAddress\":\"176.113.115.73\",\"IpPort\":\"0\",\"LmPackageName\":\"-\",\"LogonProcessName\":\"NtLmSsp \",\"LogonType\":3,\"LogonTypeName\":\"3 - Network\",\"Process\":\"-\",\"ProcessId\":\"0x0\",\"__entityMapping\":{\"\\\\ADMINISTRATOR\":\"Account\",\"SecurityEvents\":\"Host\"}}"),
	// 				Updated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-06-17T15:34:01.426Z"); return t}()),
	// 				UpdatedBy: &armsecurityinsights.UserInfo{
	// 					Name: to.Ptr("user"),
	// 					Email: to.Ptr("user@microsoft.com"),
	// 					ObjectID: to.Ptr("b03ca914-5eb6-45e5-9417-fe0797c372fd"),
	// 				},
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2021-10-01/examples/incidents/GetAllIncidentEntities.json
func ExampleIncidentsClient_ListEntities() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIncidentsClient().ListEntities(ctx, "myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IncidentEntitiesResponse = armsecurityinsights.IncidentEntitiesResponse{
	// 	Entities: []armsecurityinsights.EntityClassification{
	// 		&armsecurityinsights.AccountEntity{
	// 			Name: to.Ptr("e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 			Type: to.Ptr("Microsoft.SecurityInsights/Entities"),
	// 			ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/Entities/e1d3d618-e11f-478b-98e3-bb381539a8e1"),
	// 			Kind: to.Ptr(armsecurityinsights.EntityKindEnumAccount),
	// 			Properties: &armsecurityinsights.AccountEntityProperties{
	// 				FriendlyName: to.Ptr("administrator"),
	// 				AccountName: to.Ptr("administrator"),
	// 				NtDomain: to.Ptr("domain"),
	// 			},
	// 	}},
	// 	MetaData: []*armsecurityinsights.IncidentEntitiesResultsMetadata{
	// 		{
	// 			Count: to.Ptr[int32](1),
	// 			EntityKind: to.Ptr(armsecurityinsights.EntityKindEnumAccount),
	// 	}},
	// }
}