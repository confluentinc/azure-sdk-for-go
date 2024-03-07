//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsiot

import "time"

// DeviceService - The description of the Windows IoT Device Service.
type DeviceService struct {
	// The Etag field is not required. If it is provided in the response body, it must also be provided as a header per the normal
	// ETag convention.
	Etag *string

	// The Azure Region where the resource lives
	Location *string

	// The properties of a Windows IoT Device Service.
	Properties *DeviceServiceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// DeviceServiceCheckNameAvailabilityParameters - Input values.
type DeviceServiceCheckNameAvailabilityParameters struct {
	// REQUIRED; The name of the Windows IoT Device Service to check.
	Name *string
}

// DeviceServiceDescriptionListResult - The JSON-serialized array of DeviceService objects with a next link.
type DeviceServiceDescriptionListResult struct {
	// The array of DeviceService objects.
	Value []*DeviceService

	// READ-ONLY; The next link.
	NextLink *string
}

// DeviceServiceNameAvailabilityInfo - The properties indicating whether a given Windows IoT Device Service name is available.
type DeviceServiceNameAvailabilityInfo struct {
	// The detailed reason message.
	Message *string

	// READ-ONLY; The value which indicates whether the provided name is available.
	NameAvailable *bool

	// READ-ONLY; The reason for unavailability.
	Reason *ServiceNameUnavailabilityReason
}

// DeviceServiceProperties - The properties of a Windows IoT Device Service.
type DeviceServiceProperties struct {
	// Windows IoT Device Service OEM AAD domain
	AdminDomainName *string

	// Windows IoT Device Service ODM AAD domain
	BillingDomainName *string

	// Windows IoT Device Service notes.
	Notes *string

	// Windows IoT Device Service device allocation,
	Quantity *int64

	// READ-ONLY; Windows IoT Device Service start date,
	StartDate *time.Time
}

// ErrorDetails - The details of the error.
type ErrorDetails struct {
	// The error object.
	Error *ErrorDetailsError
}

// ErrorDetailsError - The error object.
type ErrorDetailsError struct {
	// One of a server-defined set of error codes.
	Code *string

	// A human-readable representation of the error's details.
	Details *string

	// A human-readable representation of the error.
	Message *string

	// The target of the particular error.
	Target *string
}

// OperationDisplayInfo - The operation supported by Azure Data Catalog Service.
type OperationDisplayInfo struct {
	// The description of the operation.
	Description *string

	// The action that users can perform, based on their permission level.
	Operation *string

	// Service provider: Azure Data Catalog Service.
	Provider *string

	// Resource on which the operation is performed.
	Resource *string
}

// OperationEntity - The operation supported by Azure Data Catalog Service.
type OperationEntity struct {
	// The operation supported by Azure Data Catalog Service.
	Display *OperationDisplayInfo

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name: {provider}/{resource}/{operation}.
	Name *string

	// Indicates the executor of the operation.
	Origin *string
}

// OperationListResult - Result of the request to list Windows IoT Device Service operations. It contains a list of operations
// and a URL link to get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string

	// READ-ONLY; List of Windows IoT Device Service operations supported by the Microsoft.WindowsIoT resource provider.
	Value []*OperationEntity
}

// ProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required location
// and tags
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource Id for the resource
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// Resource - The core properties of ARM resources
type Resource struct {
	// READ-ONLY; Fully qualified resource Id for the resource
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// TrackedResource - The resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// The Azure Region where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}