//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/createCompilationJob.json
func ExampleDscCompilationJobClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDscCompilationJobClient().BeginCreate(ctx, "rg", "myAutomationAccount33", "TestCompilationJob", armautomation.DscCompilationJobCreateParameters{
		Properties: &armautomation.DscCompilationJobCreateProperties{
			Configuration: &armautomation.DscConfigurationAssociationProperty{
				Name: to.Ptr("SetupServer"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getCompilationJob.json
func ExampleDscCompilationJobClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscCompilationJobClient().Get(ctx, "rg", "myAutomationAccount33", "TestCompilationJob", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DscCompilationJob = armautomation.DscCompilationJob{
	// 	Name: to.Ptr("TestCompilationJob"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/compilationjobs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/TestCompilationJob"),
	// 	Properties: &armautomation.DscCompilationJobProperties{
	// 		Configuration: &armautomation.DscConfigurationAssociationProperty{
	// 			Name: to.Ptr("SetupServer"),
	// 		},
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903Z"); return t}()),
	// 		JobID: to.Ptr("ce6fe3e3-9db3-4096-a6b4-82bfb4c10a9a"),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903Z"); return t}()),
	// 		LastStatusModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-28T23:14:26.903Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armautomation.JobProvisioningStateSucceeded),
	// 		Status: to.Ptr(armautomation.JobStatusNew),
	// 		StatusDetails: to.Ptr("None"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCompilationJobsByAutomationAccount.json
func ExampleDscCompilationJobClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDscCompilationJobClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", &armautomation.DscCompilationJobClientListByAutomationAccountOptions{Filter: nil})
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
		// page.DscCompilationJobListResult = armautomation.DscCompilationJobListResult{
		// 	Value: []*armautomation.DscCompilationJob{
		// 		{
		// 			Name: to.Ptr("CompilationConfiguration1"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/compilationjobs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/"),
		// 			Properties: &armautomation.DscCompilationJobProperties{
		// 				Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 					Name: to.Ptr("TestDscConfiguration"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T19:45:24.590Z"); return t}()),
		// 				JobID: to.Ptr("e6e7fbab-183c-405a-afe6-9eb5db97921a"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T19:45:58.593Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-17T19:45:52.983Z"); return t}()),
		// 				Status: to.Ptr(armautomation.JobStatusSuspended),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CompilationConfiguration2"),
		// 			Type: to.Ptr("Microsoft.Automation/AutomationAccounts/compilationjobs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/"),
		// 			Properties: &armautomation.DscCompilationJobProperties{
		// 				Configuration: &armautomation.DscConfigurationAssociationProperty{
		// 					Name: to.Ptr("NewDscConfiguration"),
		// 				},
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:29:07.740Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.600Z"); return t}()),
		// 				JobID: to.Ptr("111d4e06-2d88-46b4-8500-7febd4906838"),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.600Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:26.480Z"); return t}()),
		// 				Status: to.Ptr(armautomation.JobStatusCompleted),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/compilationJobStreamByJobStreamId.json
func ExampleDscCompilationJobClient_GetStream() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDscCompilationJobClient().GetStream(ctx, "rg", "myAutomationAccount33", "836d4e06-2d88-46b4-8500-7febd4906838", "836d4e06-2d88-46b4-8500-7febd4906838_00636481062421684835_00000000000000000008", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobStream = armautomation.JobStream{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/compilationjobs/836d4e06-2d88-46b4-8500-7febd4906838/streams/836d4e06-2d88-46b4-8500-7febd4906838_00636481062421684835_00000000000000000008"),
	// 	Properties: &armautomation.JobStreamProperties{
	// 		JobStreamID: to.Ptr("836d4e06-2d88-46b4-8500-7febd4906838:00636481062421684835:00000000000000000001"),
	// 		StreamText: to.Ptr(""),
	// 		StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
	// 		Summary: to.Ptr(""),
	// 		Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-12-05T21:30:42.168Z"); return t}()),
	// 		Value: map[string]any{
	// 			"PSComputerName": "localhost",
	// 			"PSShowComputerName": true,
	// 			"PSSourceJobInstanceId": "836d4e06-2d88-46b4-8500-7febd4906838",
	// 			"value": "",
	// 		},
	// 	},
	// }
}