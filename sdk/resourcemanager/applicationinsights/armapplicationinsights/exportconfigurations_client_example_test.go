//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationsList.json
func ExampleExportConfigurationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportConfigurationsClient().List(ctx, "my-resource-group", "my-component", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentExportConfigurationArray = []*armapplicationinsights.ComponentExportConfiguration{
	// 	{
	// 		ApplicationName: to.Ptr("my-component"),
	// 		ContainerName: to.Ptr("mystorageblob"),
	// 		DestinationAccountID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorageblob"),
	// 		DestinationStorageLocationID: to.Ptr("eastasia"),
	// 		DestinationStorageSubscriptionID: to.Ptr("subidc"),
	// 		DestinationType: to.Ptr("Blob"),
	// 		ExportID: to.Ptr("uGOoki0jQsyEs3IdQ83Q4QsNr4="),
	// 		ExportStatus: to.Ptr("Preparing"),
	// 		InstrumentationKey: to.Ptr("8330b4a4-0b8e-40cf-a643-bbaf60d375c9"),
	// 		IsUserEnabled: to.Ptr("False"),
	// 		LastGapTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 		LastSuccessTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 		LastUserUpdate: to.Ptr("2017-06-05T06:34:26.957Z"),
	// 		NotificationQueueEnabled: to.Ptr("False"),
	// 		PermanentErrorReason: to.Ptr("None"),
	// 		RecordTypes: to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
	// 		ResourceGroup: to.Ptr("2"),
	// 		StorageName: to.Ptr("mystorageblob"),
	// 		SubscriptionID: to.Ptr("subid"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationsPost.json
func ExampleExportConfigurationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportConfigurationsClient().Create(ctx, "my-resource-group", "my-component", armapplicationinsights.ComponentExportRequest{
		DestinationAccountID:             to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.ClassicStorage/storageAccounts/mystorageblob"),
		DestinationAddress:               to.Ptr("https://mystorageblob.blob.core.windows.net/testexport?sv=2015-04-05&sr=c&sig=token"),
		DestinationStorageLocationID:     to.Ptr("eastus"),
		DestinationStorageSubscriptionID: to.Ptr("subid"),
		DestinationType:                  to.Ptr("Blob"),
		IsEnabled:                        to.Ptr("true"),
		NotificationQueueEnabled:         to.Ptr("false"),
		NotificationQueueURI:             to.Ptr(""),
		RecordTypes:                      to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentExportConfigurationArray = []*armapplicationinsights.ComponentExportConfiguration{
	// 	{
	// 		ApplicationName: to.Ptr("my-component"),
	// 		ContainerName: to.Ptr("mystorageblob"),
	// 		DestinationAccountID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorageblob"),
	// 		DestinationStorageLocationID: to.Ptr("eastasia"),
	// 		DestinationStorageSubscriptionID: to.Ptr("subidc"),
	// 		DestinationType: to.Ptr("Blob"),
	// 		ExportID: to.Ptr("uGOoki0jQsyEs3IdQ83Q4QsNr4="),
	// 		ExportStatus: to.Ptr("Preparing"),
	// 		InstrumentationKey: to.Ptr("8330b4a4-0b8e-40cf-a643-bbaf60d375c9"),
	// 		IsUserEnabled: to.Ptr("False"),
	// 		LastGapTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 		LastSuccessTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 		LastUserUpdate: to.Ptr("2017-06-05T06:34:26.957Z"),
	// 		NotificationQueueEnabled: to.Ptr("False"),
	// 		PermanentErrorReason: to.Ptr("None"),
	// 		RecordTypes: to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
	// 		ResourceGroup: to.Ptr("2"),
	// 		StorageName: to.Ptr("mystorageblob"),
	// 		SubscriptionID: to.Ptr("subid"),
	// }}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationDelete.json
func ExampleExportConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportConfigurationsClient().Delete(ctx, "my-resource-group", "my-component", "uGOoki0jQsyEs3IdQ83Q4QsNr4=", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentExportConfiguration = armapplicationinsights.ComponentExportConfiguration{
	// 	ApplicationName: to.Ptr("my-component"),
	// 	ContainerName: to.Ptr("mystorageblob"),
	// 	DestinationAccountID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorageblob"),
	// 	DestinationStorageLocationID: to.Ptr("eastasia"),
	// 	DestinationStorageSubscriptionID: to.Ptr("subidc"),
	// 	DestinationType: to.Ptr("Blob"),
	// 	ExportID: to.Ptr("uGOoki0jQsyEs3IdQ83Q4QsNr4="),
	// 	ExportStatus: to.Ptr("Preparing"),
	// 	InstrumentationKey: to.Ptr("8330b4a4-0b8e-40cf-a643-bbaf60d375c9"),
	// 	IsUserEnabled: to.Ptr("False"),
	// 	LastGapTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 	LastSuccessTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 	LastUserUpdate: to.Ptr("2017-06-05T06:34:26.957Z"),
	// 	NotificationQueueEnabled: to.Ptr("False"),
	// 	PermanentErrorReason: to.Ptr("None"),
	// 	RecordTypes: to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
	// 	ResourceGroup: to.Ptr("2"),
	// 	StorageName: to.Ptr("mystorageblob"),
	// 	SubscriptionID: to.Ptr("subid"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationGet.json
func ExampleExportConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportConfigurationsClient().Get(ctx, "my-resource-group", "my-component", "uGOoki0jQsyEs3IdQ83Q4QsNr4=", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentExportConfiguration = armapplicationinsights.ComponentExportConfiguration{
	// 	ApplicationName: to.Ptr("my-component"),
	// 	ContainerName: to.Ptr("mystorageblob"),
	// 	DestinationAccountID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorageblob"),
	// 	DestinationStorageLocationID: to.Ptr("eastasia"),
	// 	DestinationStorageSubscriptionID: to.Ptr("subidc"),
	// 	DestinationType: to.Ptr("Blob"),
	// 	ExportID: to.Ptr("uGOoki0jQsyEs3IdQ83Q4QsNr4="),
	// 	ExportStatus: to.Ptr("Preparing"),
	// 	InstrumentationKey: to.Ptr("8330b4a4-0b8e-40cf-a643-bbaf60d375c9"),
	// 	IsUserEnabled: to.Ptr("False"),
	// 	LastGapTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 	LastSuccessTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 	LastUserUpdate: to.Ptr("2017-06-05T06:34:26.957Z"),
	// 	NotificationQueueEnabled: to.Ptr("False"),
	// 	PermanentErrorReason: to.Ptr("None"),
	// 	RecordTypes: to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
	// 	ResourceGroup: to.Ptr("2"),
	// 	StorageName: to.Ptr("mystorageblob"),
	// 	SubscriptionID: to.Ptr("subid"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationUpdate.json
func ExampleExportConfigurationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExportConfigurationsClient().Update(ctx, "my-resource-group", "my-component", "uGOoki0jQsyEs3IdQ83Q4QsNr4=", armapplicationinsights.ComponentExportRequest{
		DestinationAccountID:             to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.ClassicStorage/storageAccounts/mystorageblob"),
		DestinationAddress:               to.Ptr("https://mystorageblob.blob.core.windows.net/fchentest?sv=2015-04-05&sr=c&sig=token"),
		DestinationStorageLocationID:     to.Ptr("eastus"),
		DestinationStorageSubscriptionID: to.Ptr("subid"),
		DestinationType:                  to.Ptr("Blob"),
		IsEnabled:                        to.Ptr("true"),
		NotificationQueueEnabled:         to.Ptr("false"),
		NotificationQueueURI:             to.Ptr(""),
		RecordTypes:                      to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentExportConfiguration = armapplicationinsights.ComponentExportConfiguration{
	// 	ApplicationName: to.Ptr("my-component"),
	// 	ContainerName: to.Ptr("mystorageblob"),
	// 	DestinationAccountID: to.Ptr("/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Storage/storageAccounts/mystorageblob"),
	// 	DestinationStorageLocationID: to.Ptr("eastasia"),
	// 	DestinationStorageSubscriptionID: to.Ptr("subidc"),
	// 	DestinationType: to.Ptr("Blob"),
	// 	ExportID: to.Ptr("uGOoki0jQsyEs3IdQ83Q4QsNr4="),
	// 	ExportStatus: to.Ptr("Preparing"),
	// 	InstrumentationKey: to.Ptr("8330b4a4-0b8e-40cf-a643-bbaf60d375c9"),
	// 	IsUserEnabled: to.Ptr("False"),
	// 	LastGapTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 	LastSuccessTime: to.Ptr("9999-12-31T23:59:59.999Z"),
	// 	LastUserUpdate: to.Ptr("2017-06-05T06:34:26.957Z"),
	// 	NotificationQueueEnabled: to.Ptr("False"),
	// 	PermanentErrorReason: to.Ptr("None"),
	// 	RecordTypes: to.Ptr("Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability"),
	// 	ResourceGroup: to.Ptr("2"),
	// 	StorageName: to.Ptr("mystorageblob"),
	// 	SubscriptionID: to.Ptr("subid"),
	// }
}
