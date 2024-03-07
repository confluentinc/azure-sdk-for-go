//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlinks

// Operation - Microsoft.Resources operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string

	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.Resources
	Provider *string

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string
}

// OperationListResult - Result of the request to list Microsoft.Resources operations. It contains a list of operations and
// a URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of Microsoft.Resources operations.
	Value []*Operation
}

// ResourceLink - The resource link.
type ResourceLink struct {
	// Properties for resource link.
	Properties *ResourceLinkProperties

	// READ-ONLY; The fully qualified ID of the resource link.
	ID *string

	// READ-ONLY; The name of the resource link.
	Name *string

	// READ-ONLY; The resource link object.
	Type any
}

// ResourceLinkFilter - Resource link filter.
type ResourceLinkFilter struct {
	// REQUIRED; The ID of the target resource.
	TargetID *string
}

// ResourceLinkProperties - The resource link properties.
type ResourceLinkProperties struct {
	// REQUIRED; The fully qualified ID of the target resource in the link.
	TargetID *string

	// Notes about the resource link.
	Notes *string

	// READ-ONLY; The fully qualified ID of the source resource in the link.
	SourceID *string
}

// ResourceLinkResult - List of resource links.
type ResourceLinkResult struct {
	// REQUIRED; An array of resource links.
	Value []*ResourceLink

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}