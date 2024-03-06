//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsiot

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientCheckDeviceServiceNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckDeviceServiceNameAvailability
// method.
type ServicesClientCheckDeviceServiceNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientCreateOrUpdateOptions contains the optional parameters for the ServicesClient.CreateOrUpdate method.
type ServicesClientCreateOrUpdateOptions struct {
	// ETag of the Windows IoT Device Service. Do not specify for creating a new Windows IoT Device Service. Required to update
	// an existing Windows IoT Device Service.
	IfMatch *string
}

// ServicesClientDeleteOptions contains the optional parameters for the ServicesClient.Delete method.
type ServicesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
type ServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.NewListByResourceGroupPager
// method.
type ServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListOptions contains the optional parameters for the ServicesClient.NewListPager method.
type ServicesClientListOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientUpdateOptions contains the optional parameters for the ServicesClient.Update method.
type ServicesClientUpdateOptions struct {
	// ETag of the Windows IoT Device Service. Do not specify for creating a brand new Windows IoT Device Service. Required to
	// update an existing Windows IoT Device Service.
	IfMatch *string
}
