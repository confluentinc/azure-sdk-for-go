//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azeventgrid

// AcknowledgeCloudEventsOptions contains the optional parameters for the Client.AcknowledgeCloudEvents method.
type AcknowledgeCloudEventsOptions struct {
	// placeholder for future optional parameters
}

// PublishCloudEventOptions contains the optional parameters for the Client.PublishCloudEvent method.
type PublishCloudEventOptions struct {
	// placeholder for future optional parameters
}

// PublishCloudEventsOptions contains the optional parameters for the Client.PublishCloudEvents method.
type PublishCloudEventsOptions struct {
	// placeholder for future optional parameters
}

// ReceiveCloudEventsOptions contains the optional parameters for the Client.ReceiveCloudEvents method.
type ReceiveCloudEventsOptions struct {
	// Max Events count to be received. Minimum value is 1, while maximum value is 100 events. If not specified, the default value
	// is 1.
	MaxEvents *int32

	// Max wait time value for receive operation in Seconds. It is the time in seconds that the server approximately waits for
	// the availability of an event and responds to the request. If an event is
	// available, the broker responds immediately to the client. Minimum value is 10 seconds, while maximum value is 120 seconds.
	// If not specified, the default value is 60 seconds.
	MaxWaitTime *int32
}

// RejectCloudEventsOptions contains the optional parameters for the Client.RejectCloudEvents method.
type RejectCloudEventsOptions struct {
	// placeholder for future optional parameters
}

// ReleaseCloudEventsOptions contains the optional parameters for the Client.ReleaseCloudEvents method.
type ReleaseCloudEventsOptions struct {
	// Release cloud events with the specified delay in seconds.
	ReleaseDelayInSeconds *ReleaseDelay
}

// RenewCloudEventLocksOptions contains the optional parameters for the Client.RenewCloudEventLocks method.
type RenewCloudEventLocksOptions struct {
	// placeholder for future optional parameters
}