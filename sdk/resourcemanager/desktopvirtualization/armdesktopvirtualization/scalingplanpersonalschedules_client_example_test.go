//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Get.json
func ExampleScalingPlanPersonalSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScalingPlanPersonalSchedulesClient().Get(ctx, "resourceGroup1", "PersonalScalingPlan1", "PersonalScalingPlanSchedule", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScalingPlanPersonalSchedule = armdesktopvirtualization.ScalingPlanPersonalSchedule{
	// 	Name: to.Ptr("PersonalScalingPlanSchedule"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/personalSchedules"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/personalSchedules/PersonalScalingPlanScheduleWeekdays1"),
	// 	Properties: &armdesktopvirtualization.ScalingPlanPersonalScheduleProperties{
	// 		DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
	// 			OffPeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			OffPeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			OffPeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			OffPeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			OffPeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](20),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			OffPeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			PeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			PeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			PeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			PeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			PeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](8),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			PeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			RampDownActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampDownActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			RampDownMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			RampDownMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			RampDownStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](18),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			RampUpActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampUpActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampUpAutoStartHosts: to.Ptr(armdesktopvirtualization.StartupBehaviorAll),
	// 			RampUpMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			RampUpMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			RampUpStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](6),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampUpStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 		},
	// 		SystemData: &armdesktopvirtualization.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Create.json
func ExampleScalingPlanPersonalSchedulesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScalingPlanPersonalSchedulesClient().Create(ctx, "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", armdesktopvirtualization.ScalingPlanPersonalSchedule{
		Properties: &armdesktopvirtualization.ScalingPlanPersonalScheduleProperties{
			DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
				to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
				to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
				to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
				to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
				to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
			OffPeakActionOnDisconnect:        to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			OffPeakActionOnLogoff:            to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
			OffPeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
			OffPeakMinutesToWaitOnLogoff:     to.Ptr[int32](10),
			OffPeakStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](20),
				Minute: to.Ptr[int32](0),
			},
			OffPeakStartVMOnConnect:       to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
			PeakActionOnDisconnect:        to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			PeakActionOnLogoff:            to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
			PeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
			PeakMinutesToWaitOnLogoff:     to.Ptr[int32](10),
			PeakStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](8),
				Minute: to.Ptr[int32](0),
			},
			PeakStartVMOnConnect:              to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
			RampDownActionOnDisconnect:        to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			RampDownActionOnLogoff:            to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
			RampDownMinutesToWaitOnDisconnect: to.Ptr[int32](10),
			RampDownMinutesToWaitOnLogoff:     to.Ptr[int32](10),
			RampDownStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](18),
				Minute: to.Ptr[int32](0),
			},
			RampDownStartVMOnConnect:        to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
			RampUpActionOnDisconnect:        to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			RampUpActionOnLogoff:            to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			RampUpAutoStartHosts:            to.Ptr(armdesktopvirtualization.StartupBehaviorAll),
			RampUpMinutesToWaitOnDisconnect: to.Ptr[int32](10),
			RampUpMinutesToWaitOnLogoff:     to.Ptr[int32](10),
			RampUpStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](6),
				Minute: to.Ptr[int32](0),
			},
			RampUpStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScalingPlanPersonalSchedule = armdesktopvirtualization.ScalingPlanPersonalSchedule{
	// 	Name: to.Ptr("scalingPlanScheduleWeekdays1"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/personalSchedules"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/personalSchedules/scalingPlanScheduleWeekdays1"),
	// 	Properties: &armdesktopvirtualization.ScalingPlanPersonalScheduleProperties{
	// 		DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
	// 			OffPeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			OffPeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			OffPeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			OffPeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			OffPeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](20),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			OffPeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			PeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			PeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			PeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			PeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			PeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](8),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			PeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			RampDownActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampDownActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			RampDownMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			RampDownMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			RampDownStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](18),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			RampUpActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampUpActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampUpAutoStartHosts: to.Ptr(armdesktopvirtualization.StartupBehaviorAll),
	// 			RampUpMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			RampUpMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			RampUpStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](6),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		SystemData: &armdesktopvirtualization.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Delete.json
func ExampleScalingPlanPersonalSchedulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScalingPlanPersonalSchedulesClient().Delete(ctx, "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_Update.json
func ExampleScalingPlanPersonalSchedulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScalingPlanPersonalSchedulesClient().Update(ctx, "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", &armdesktopvirtualization.ScalingPlanPersonalSchedulesClientUpdateOptions{ScalingPlanSchedule: &armdesktopvirtualization.ScalingPlanPersonalSchedulePatch{
		Properties: &armdesktopvirtualization.ScalingPlanPersonalScheduleProperties{
			OffPeakActionOnDisconnect:        to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			OffPeakActionOnLogoff:            to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
			OffPeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
			OffPeakMinutesToWaitOnLogoff:     to.Ptr[int32](10),
			OffPeakStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](20),
				Minute: to.Ptr[int32](0),
			},
			OffPeakStartVMOnConnect:       to.Ptr(armdesktopvirtualization.SetStartVMOnConnectDisable),
			PeakActionOnDisconnect:        to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			PeakActionOnLogoff:            to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
			PeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
			PeakMinutesToWaitOnLogoff:     to.Ptr[int32](10),
			PeakStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](8),
				Minute: to.Ptr[int32](0),
			},
			RampDownActionOnDisconnect:        to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
			RampDownActionOnLogoff:            to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
			RampDownMinutesToWaitOnDisconnect: to.Ptr[int32](10),
			RampDownMinutesToWaitOnLogoff:     to.Ptr[int32](10),
			RampDownStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](18),
				Minute: to.Ptr[int32](0),
			},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScalingPlanPersonalSchedule = armdesktopvirtualization.ScalingPlanPersonalSchedule{
	// 	Name: to.Ptr("scalingPlanScheduleWeekdays1"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/personalSchedules"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/personalSchedules/scalingPlanScheduleWeekdays1"),
	// 	Properties: &armdesktopvirtualization.ScalingPlanPersonalScheduleProperties{
	// 		DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
	// 			OffPeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			OffPeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			OffPeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			OffPeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			OffPeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](20),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			OffPeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectDisable),
	// 			PeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			PeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			PeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			PeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			PeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](8),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			PeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			RampDownActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampDownActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
	// 			RampDownMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			RampDownMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			RampDownStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](18),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 			RampUpActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampUpActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
	// 			RampUpAutoStartHosts: to.Ptr(armdesktopvirtualization.StartupBehaviorAll),
	// 			RampUpMinutesToWaitOnDisconnect: to.Ptr[int32](10),
	// 			RampUpMinutesToWaitOnLogoff: to.Ptr[int32](10),
	// 			RampUpStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](6),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampUpStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
	// 		},
	// 		SystemData: &armdesktopvirtualization.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ScalingPlanPersonalSchedule_List.json
func ExampleScalingPlanPersonalSchedulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScalingPlanPersonalSchedulesClient().NewListPager("resourceGroup1", "scalingPlan", &armdesktopvirtualization.ScalingPlanPersonalSchedulesClientListOptions{PageSize: to.Ptr[int32](10),
		IsDescending: to.Ptr(true),
		InitialSkip:  to.Ptr[int32](0),
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
		// page.ScalingPlanPersonalScheduleList = armdesktopvirtualization.ScalingPlanPersonalScheduleList{
		// 	Value: []*armdesktopvirtualization.ScalingPlanPersonalSchedule{
		// 		{
		// 			Name: to.Ptr("scalingPlanScheduleWeekday"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/personalSchedules"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/personalSchedules/scalingPlanScheduleWeekdays1"),
		// 			Properties: &armdesktopvirtualization.ScalingPlanPersonalScheduleProperties{
		// 				DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
		// 					OffPeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
		// 					OffPeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 					OffPeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 					OffPeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 					OffPeakStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](20),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 					OffPeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 					PeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
		// 					PeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 					PeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 					PeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 					PeakStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](8),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 					PeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 					RampDownActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
		// 					RampDownActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 					RampDownMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 					RampDownMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 					RampDownStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](18),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 					RampDownStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 					RampUpActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
		// 					RampUpActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
		// 					RampUpAutoStartHosts: to.Ptr(armdesktopvirtualization.StartupBehaviorAll),
		// 					RampUpMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 					RampUpMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 					RampUpStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](6),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 					RampUpStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 				},
		// 				SystemData: &armdesktopvirtualization.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("scalingPlanScheduleWeekend"),
		// 				Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/personalSchedules"),
		// 				ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/personalSchedules/scalingPlanScheduleWeekends1"),
		// 				Properties: &armdesktopvirtualization.ScalingPlanPersonalScheduleProperties{
		// 					DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
		// 						to.Ptr(armdesktopvirtualization.DayOfWeekSaturday),
		// 						to.Ptr(armdesktopvirtualization.DayOfWeekSunday)},
		// 						OffPeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 						OffPeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 						OffPeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 						OffPeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 						OffPeakStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](16),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 						OffPeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 						PeakActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 						PeakActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 						PeakMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 						PeakMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 						PeakStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](10),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 						PeakStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 						RampDownActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 						RampDownActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationDeallocate),
		// 						RampDownMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 						RampDownMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 						RampDownStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](14),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 						RampDownStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 						RampUpActionOnDisconnect: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
		// 						RampUpActionOnLogoff: to.Ptr(armdesktopvirtualization.SessionHandlingOperationNone),
		// 						RampUpAutoStartHosts: to.Ptr(armdesktopvirtualization.StartupBehaviorAll),
		// 						RampUpMinutesToWaitOnDisconnect: to.Ptr[int32](10),
		// 						RampUpMinutesToWaitOnLogoff: to.Ptr[int32](10),
		// 						RampUpStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](6),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 						RampUpStartVMOnConnect: to.Ptr(armdesktopvirtualization.SetStartVMOnConnectEnable),
		// 					},
		// 					SystemData: &armdesktopvirtualization.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 						CreatedBy: to.Ptr("user1"),
		// 						CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("user2"),
		// 						LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 					},
		// 			}},
		// 		}
	}
}