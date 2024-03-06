//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthbot

// BotsClientCreateResponse contains the response from method BotsClient.BeginCreate.
type BotsClientCreateResponse struct {
	// Azure Health Bot resource definition
	HealthBot
}

// BotsClientDeleteResponse contains the response from method BotsClient.BeginDelete.
type BotsClientDeleteResponse struct {
	// placeholder for future response values
}

// BotsClientGetResponse contains the response from method BotsClient.Get.
type BotsClientGetResponse struct {
	// Azure Health Bot resource definition
	HealthBot
}

// BotsClientListByResourceGroupResponse contains the response from method BotsClient.NewListByResourceGroupPager.
type BotsClientListByResourceGroupResponse struct {
	// The list of Azure Health Bot operation response.
	BotResponseList
}

// BotsClientListResponse contains the response from method BotsClient.NewListPager.
type BotsClientListResponse struct {
	// The list of Azure Health Bot operation response.
	BotResponseList
}

// BotsClientUpdateResponse contains the response from method BotsClient.Update.
type BotsClientUpdateResponse struct {
	// Azure Health Bot resource definition
	HealthBot
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Available operations of the service
	AvailableOperations
}
