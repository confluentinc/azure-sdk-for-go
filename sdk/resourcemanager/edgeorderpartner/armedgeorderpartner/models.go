//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorderpartner

import "time"

type AdditionalErrorInfo struct {
	// Anything
	Info any
	Type *string
}

// AdditionalInventoryDetails - Contains additional data about inventory in dictionary format
type AdditionalInventoryDetails struct {
	// READ-ONLY; Additional Data
	AdditionalData map[string]*string
}

// AdditionalOrderItemDetails - Contains additional order item details
type AdditionalOrderItemDetails struct {
	// READ-ONLY; Order item status
	Status *StageDetails

	// READ-ONLY; Subscription details
	Subscription *SubscriptionDetails
}

// BillingDetails - Contains billing details for the inventory
type BillingDetails struct {
	// READ-ONLY; Billing type for the inventory
	BillingType *string

	// READ-ONLY; Billing status for the inventory
	Status *string
}

// ConfigurationData - Contains information about inventory configuration
type ConfigurationData struct {
	// READ-ONLY; Configuration identifier of inventory
	ConfigurationIdentifier *string

	// READ-ONLY; Configuration identifier on device - this is used in case of any mismatch between actual configuration on inventory
	// and configuration stored in service
	ConfigurationIdentifierOnDevice *string

	// READ-ONLY; Family identifier of inventory
	FamilyIdentifier *string

	// READ-ONLY; Product identifier of inventory
	ProductIdentifier *string

	// READ-ONLY; Product Line identifier of inventory
	ProductLineIdentifier *string
}

// ConfigurationDetails - Contains additional configuration details about inventory
type ConfigurationDetails struct {
	// READ-ONLY; Collection of specification details about the inventory
	Specifications []*SpecificationDetails
}

// ConfigurationOnDevice - Configuration parameters for ManageInventoryMetadata call
type ConfigurationOnDevice struct {
	// REQUIRED; Configuration identifier on device
	ConfigurationIdentifier *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// InventoryAdditionalDetails - Represents additional details about the partner inventory
type InventoryAdditionalDetails struct {
	// Represents additional details about the order item
	OrderItem *AdditionalOrderItemDetails

	// READ-ONLY; Represents additional details about billing for the inventory
	Billing *BillingDetails

	// READ-ONLY; Represents additional details about the configuration
	Configuration *ConfigurationDetails

	// READ-ONLY; Represents additional data about the inventory
	Inventory *AdditionalInventoryDetails

	// READ-ONLY; Contains inventory metadata
	InventoryMetadata *string

	// READ-ONLY; Represents secrets on the inventory
	InventorySecrets map[string]*string
}

// InventoryData - Contains basic information about inventory
type InventoryData struct {
	// READ-ONLY; Inventory location
	Location *string

	// READ-ONLY; Boolean flag to indicate if registration is allowed
	RegistrationAllowed *bool

	// READ-ONLY; Inventory status
	Status *string
}

// InventoryProperties - Represents inventory properties
type InventoryProperties struct {
	// READ-ONLY; Represents basic configuration data.
	Configuration *ConfigurationData

	// READ-ONLY; Represents additional details of inventory
	Details *InventoryAdditionalDetails

	// READ-ONLY; Represents basic inventory data.
	Inventory *InventoryData

	// READ-ONLY; Location of inventory
	Location *string

	// READ-ONLY; Represents management resource data associated with inventory.
	ManagementResource *ManagementResourceData

	// READ-ONLY; Represents basic order item data.
	OrderItem *OrderItemData

	// READ-ONLY; Serial number of the device.
	SerialNumber *string
}

// ManageInventoryMetadataRequest - Request body for ManageInventoryMetadata call
type ManageInventoryMetadataRequest struct {
	// REQUIRED; Inventory metadata to be updated
	InventoryMetadata *string

	// Inventory configuration to be updated
	ConfigurationOnDevice *ConfigurationOnDevice
}

// ManageLinkRequest - Request body for ManageLink call
type ManageLinkRequest struct {
	// REQUIRED; Arm Id of the management resource to which inventory is to be linked For unlink operation, enter empty string
	ManagementResourceArmID *string

	// REQUIRED; Operation to be performed - Link, Unlink, Relink
	Operation *ManageLinkOperation

	// REQUIRED; Tenant ID of management resource associated with inventory
	TenantID *string
}

// ManagementResourceData - Contains information about management resource
type ManagementResourceData struct {
	// READ-ONLY; Arm ID of management resource associated with inventory
	ArmID *string

	// READ-ONLY; Tenant ID of management resource associated with inventory
	TenantID *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// OrderItemData - Contains information about the order item to which inventory belongs
type OrderItemData struct {
	// READ-ONLY; Arm ID of order item
	ArmID *string

	// READ-ONLY; Order item type - purchase or rental
	OrderItemType *OrderItemType
}

// PartnerInventory - Represents partner inventory contract
type PartnerInventory struct {
	// READ-ONLY; Inventory properties
	Properties *InventoryProperties
}

// PartnerInventoryList - Represents the list of partner inventories
type PartnerInventoryList struct {
	// Link for the next set of partner inventories.
	NextLink *string

	// READ-ONLY; List of partner inventories
	Value []*PartnerInventory
}

// SearchInventoriesRequest - Request body for SearchInventories call
type SearchInventoriesRequest struct {
	// REQUIRED; Family identifier for inventory
	FamilyIdentifier *string

	// REQUIRED; Serial number of the inventory
	SerialNumber *string
}

// SpecificationDetails - Specification details for the inventory
type SpecificationDetails struct {
	// READ-ONLY; Name of the specification property
	Name *string

	// READ-ONLY; Value of the specification property
	Value *string
}

// StageDetails - Resource stage details.
type StageDetails struct {
	// READ-ONLY; Display name of the resource stage.
	DisplayName *string

	// READ-ONLY; Stage name
	StageName *StageName

	// READ-ONLY; Stage status.
	StageStatus *StageStatus

	// READ-ONLY; Stage start time
	StartTime *time.Time
}

// SubscriptionDetails - Contains subscription details
type SubscriptionDetails struct {
	// READ-ONLY; Subscription Id
	ID *string

	// READ-ONLY; Subscription QuotaId
	QuotaID *string

	// READ-ONLY; Subscription State
	State *string
}
