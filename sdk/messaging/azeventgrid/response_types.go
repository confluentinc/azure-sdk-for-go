//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azeventgrid

// AcknowledgeCloudEventsResponse contains the response from method Client.AcknowledgeCloudEvents.
type AcknowledgeCloudEventsResponse struct {
	// The result of the Acknowledge operation.
	AcknowledgeResult
}

// PublishCloudEventResponse contains the response from method Client.PublishCloudEvent.
type PublishCloudEventResponse struct {
	// placeholder for future response values
}

// PublishCloudEventsResponse contains the response from method Client.PublishCloudEvents.
type PublishCloudEventsResponse struct {
	// placeholder for future response values
}

// ReceiveCloudEventsResponse contains the response from method Client.ReceiveCloudEvents.
type ReceiveCloudEventsResponse struct {
	// Details of the Receive operation response.
	ReceiveResult
}

// RejectCloudEventsResponse contains the response from method Client.RejectCloudEvents.
type RejectCloudEventsResponse struct {
	// The result of the Reject operation.
	RejectResult
}

// ReleaseCloudEventsResponse contains the response from method Client.ReleaseCloudEvents.
type ReleaseCloudEventsResponse struct {
	// The result of the Release operation.
	ReleaseResult
}

// RenewCloudEventLocksResponse contains the response from method Client.RenewCloudEventLocks.
type RenewCloudEventLocksResponse struct {
	// The result of the RenewLock operation.
	RenewCloudEventLocksResult
}
