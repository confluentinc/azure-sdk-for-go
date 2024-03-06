//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragepool

import "time"

// ACL - Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
type ACL struct {
	// REQUIRED; iSCSI initiator IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:client".
	InitiatorIqn *string

	// REQUIRED; List of LUN names mapped to the ACL.
	MappedLuns []*string
}

// Disk - Azure Managed Disk to attach to the Disk Pool.
type Disk struct {
	// REQUIRED; Unique Azure Resource ID of the Managed Disk.
	ID *string
}

// DiskPool - Response for Disk Pool request.
type DiskPool struct {
	// REQUIRED; The geo-location where the resource lives.
	Location *string

	// REQUIRED; Properties of Disk Pool.
	Properties *DiskPoolProperties

	// Determines the SKU of the Disk pool
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string

	// READ-ONLY; List of Azure resource ids that manage this resource.
	ManagedByExtended []*string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Resource metadata required by ARM RPC
	SystemData *SystemMetadata

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// DiskPoolCreate - Request payload for create or update Disk Pool request.
type DiskPoolCreate struct {
	// REQUIRED; The geo-location where the resource lives.
	Location *string

	// REQUIRED; Properties for Disk Pool create request.
	Properties *DiskPoolCreateProperties

	// REQUIRED; Determines the SKU of the Disk Pool
	SKU *SKU

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// DiskPoolCreateProperties - Properties for Disk Pool create or update request.
type DiskPoolCreateProperties struct {
	// REQUIRED; Azure Resource ID of a Subnet for the Disk Pool.
	SubnetID *string

	// List of additional capabilities for a Disk Pool.
	AdditionalCapabilities []*string

	// Logical zone for Disk Pool resource; example: ["1"].
	AvailabilityZones []*string

	// List of Azure Managed Disks to attach to a Disk Pool.
	Disks []*Disk
}

// DiskPoolListResult - List of Disk Pools
type DiskPoolListResult struct {
	// REQUIRED; An array of Disk pool objects.
	Value []*DiskPool

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// DiskPoolProperties - Disk Pool response properties.
type DiskPoolProperties struct {
	// REQUIRED; Logical zone for Disk Pool resource; example: ["1"].
	AvailabilityZones []*string

	// REQUIRED; Operational status of the Disk Pool.
	Status *OperationalStatus

	// REQUIRED; Azure Resource ID of a Subnet for the Disk Pool.
	SubnetID *string

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates

	// List of additional capabilities for Disk Pool.
	AdditionalCapabilities []*string

	// List of Azure Managed Disks to attach to a Disk Pool.
	Disks []*Disk
}

// DiskPoolUpdate - Request payload for Update Disk Pool request.
type DiskPoolUpdate struct {
	// REQUIRED; Properties for Disk Pool update request.
	Properties *DiskPoolUpdateProperties

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string

	// Determines the SKU of the Disk Pool
	SKU *SKU

	// Resource tags.
	Tags map[string]*string
}

// DiskPoolUpdateProperties - Properties for Disk Pool update request.
type DiskPoolUpdateProperties struct {
	// List of Azure Managed Disks to attach to a Disk Pool.
	Disks []*Disk
}

// DiskPoolZoneInfo - Disk Pool SKU Details
type DiskPoolZoneInfo struct {
	// READ-ONLY; List of additional capabilities for Disk Pool.
	AdditionalCapabilities []*string

	// READ-ONLY; Logical zone for Disk Pool resource; example: ["1"].
	AvailabilityZones []*string

	// READ-ONLY; Determines the SKU of VM deployed for Disk Pool
	SKU *SKU
}

// DiskPoolZoneListResult - List Disk Pool skus operation response.
type DiskPoolZoneListResult struct {
	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string

	// READ-ONLY; The list of Disk Pool Skus.
	Value []*DiskPoolZoneInfo
}

// EndpointDependency - A domain name that a service is reached at, including details of the current connection status.
type EndpointDependency struct {
	// The domain name of the dependency.
	DomainName *string

	// The IP Addresses and Ports used when connecting to DomainName.
	EndpointDetails []*EndpointDetail
}

// EndpointDetail - Current TCP connectivity information from the App Service Environment to a single endpoint.
type EndpointDetail struct {
	// An IP Address that Domain Name currently resolves to.
	IPAddress *string

	// Whether it is possible to create a TCP connection from the App Service Environment to this IpAddress at this Port.
	IsAccessible *bool

	// The time in milliseconds it takes for a TCP connection to be created from the App Service Environment to this IpAddress
	// at this Port.
	Latency *float64

	// The port an endpoint is connected to.
	Port *int32
}

// Error - The resource management error response.
type Error struct {
	// RP error response.
	Error *ErrorResponse
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorResponse - The resource management error response.
type ErrorResponse struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorResponse

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// IscsiLun - LUN to expose the Azure Managed Disk.
type IscsiLun struct {
	// REQUIRED; Azure Resource ID of the Managed Disk.
	ManagedDiskAzureResourceID *string

	// REQUIRED; User defined name for iSCSI LUN; example: "lun0"
	Name *string

	// READ-ONLY; Specifies the Logical Unit Number of the iSCSI LUN.
	Lun *int32
}

// IscsiTarget - Response for iSCSI Target requests.
type IscsiTarget struct {
	// REQUIRED; Properties for iSCSI Target operations.
	Properties *IscsiTargetProperties

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string

	// READ-ONLY; List of Azure resource ids that manage this resource.
	ManagedByExtended []*string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Resource metadata required by ARM RPC
	SystemData *SystemMetadata

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// IscsiTargetCreate - Payload for iSCSI Target create or update requests.
type IscsiTargetCreate struct {
	// REQUIRED; Properties for iSCSI Target create request.
	Properties *IscsiTargetCreateProperties

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// IscsiTargetCreateProperties - Properties for iSCSI Target create or update request.
type IscsiTargetCreateProperties struct {
	// REQUIRED; Mode for Target connectivity.
	ACLMode *IscsiTargetACLMode

	// List of LUNs to be exposed through iSCSI Target.
	Luns []*IscsiLun

	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticACLs []*ACL

	// iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn *string
}

// IscsiTargetList - List of iSCSI Targets.
type IscsiTargetList struct {
	// REQUIRED; An array of iSCSI Targets in a Disk Pool.
	Value []*IscsiTarget

	// READ-ONLY; URI to fetch the next section of the paginated response.
	NextLink *string
}

// IscsiTargetProperties - Response properties for iSCSI Target operations.
type IscsiTargetProperties struct {
	// REQUIRED; Mode for Target connectivity.
	ACLMode *IscsiTargetACLMode

	// REQUIRED; Operational status of the iSCSI Target.
	Status *OperationalStatus

	// REQUIRED; iSCSI Target IQN (iSCSI Qualified Name); example: "iqn.2005-03.org.iscsi:server".
	TargetIqn *string

	// READ-ONLY; State of the operation on the resource.
	ProvisioningState *ProvisioningStates

	// List of private IPv4 addresses to connect to the iSCSI Target.
	Endpoints []*string

	// List of LUNs to be exposed through iSCSI Target.
	Luns []*IscsiLun

	// The port used by iSCSI Target portal group.
	Port *int32

	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticACLs []*ACL

	// READ-ONLY; List of identifiers for active sessions on the iSCSI target
	Sessions []*string
}

// IscsiTargetUpdate - Payload for iSCSI Target update requests.
type IscsiTargetUpdate struct {
	// REQUIRED; Properties for iSCSI Target update request.
	Properties *IscsiTargetUpdateProperties

	// Azure resource id. Indicates if this resource is managed by another Azure resource.
	ManagedBy *string

	// List of Azure resource ids that manage this resource.
	ManagedByExtended []*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// IscsiTargetUpdateProperties - Properties for iSCSI Target update request.
type IscsiTargetUpdateProperties struct {
	// List of LUNs to be exposed through iSCSI Target.
	Luns []*IscsiLun

	// Access Control List (ACL) for an iSCSI Target; defines LUN masking policy
	StaticACLs []*ACL
}

// OperationDisplay - Metadata about an operation.
type OperationDisplay struct {
	// REQUIRED; Localized friendly description for the operation, as it should be shown to the user.
	Description *string

	// REQUIRED; Localized friendly name for the operation, as it should be shown to the user.
	Operation *string

	// REQUIRED; Localized friendly form of the resource provider name.
	Provider *string

	// REQUIRED; Localized friendly form of the resource type related to this action/operation.
	Resource *string
}

// OperationListResult - List of operations supported by the RP.
type OperationListResult struct {
	// REQUIRED; An array of operations supported by the StoragePool RP.
	Value []*RPOperation

	// URI to fetch the next section of the paginated response.
	NextLink *string
}

// OutboundEnvironmentEndpoint - Endpoints accessed for a common purpose that the App Service Environment requires outbound
// network access to.
type OutboundEnvironmentEndpoint struct {
	// The type of service accessed by the App Service Environment, e.g., Azure Storage, Azure SQL Database, and Azure Active
	// Directory.
	Category *string

	// The endpoints that the App Service Environment reaches the service at.
	Endpoints []*EndpointDependency
}

// OutboundEnvironmentEndpointList - Collection of Outbound Environment Endpoints
type OutboundEnvironmentEndpointList struct {
	// REQUIRED; Collection of resources.
	Value []*OutboundEnvironmentEndpoint

	// READ-ONLY; Link to next page of resources.
	NextLink *string
}

// ProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required location
// and tags
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// RPOperation - Description of a StoragePool RP Operation
type RPOperation struct {
	// REQUIRED; Additional metadata about RP operation.
	Display *OperationDisplay

	// REQUIRED; Indicates whether the operation applies to data-plane.
	IsDataAction *bool

	// REQUIRED; The name of the operation being performed on this particular object
	Name *string

	// Indicates the action type.
	ActionType *string

	// The intended executor of the operation; governs the display of the operation in the RBAC UX and the audit logs UX.
	Origin *string
}

// Resource - ARM resource model definition.
type Resource struct {
	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// ResourceSKUCapability - Capability a resource SKU has.
type ResourceSKUCapability struct {
	// READ-ONLY; Capability name
	Name *string

	// READ-ONLY; Capability value
	Value *string
}

// ResourceSKUInfo - Resource SKU Details
type ResourceSKUInfo struct {
	// READ-ONLY; StoragePool RP API version
	APIVersion *string

	// READ-ONLY; List of additional capabilities for StoragePool resource.
	Capabilities []*ResourceSKUCapability

	// READ-ONLY; Zones and zone capabilities in those locations where the SKU is available.
	LocationInfo *ResourceSKULocationInfo

	// READ-ONLY; Sku name
	Name *string

	// READ-ONLY; StoragePool resource type
	ResourceType *string

	// READ-ONLY; The restrictions because of which SKU cannot be used. This is empty if there are no restrictions.
	Restrictions []*ResourceSKURestrictions

	// READ-ONLY; Sku tier
	Tier *string
}

// ResourceSKUListResult - List Disk Pool skus operation response.
type ResourceSKUListResult struct {
	// URI to fetch the next section of the paginated response.
	NextLink *string

	// The list of StoragePool resource skus.
	Value []*ResourceSKUInfo
}

// ResourceSKULocationInfo - Zone and capability info for resource sku
type ResourceSKULocationInfo struct {
	// READ-ONLY; Location of the SKU
	Location *string

	// READ-ONLY; Details of capabilities available to a SKU in specific zones.
	ZoneDetails []*ResourceSKUZoneDetails

	// READ-ONLY; List of availability zones where the SKU is supported.
	Zones []*string
}

// ResourceSKURestrictionInfo - Describes an available Compute SKU Restriction Information.
type ResourceSKURestrictionInfo struct {
	// READ-ONLY; Locations where the SKU is restricted
	Locations []*string

	// READ-ONLY; List of availability zones where the SKU is restricted.
	Zones []*string
}

// ResourceSKURestrictions - Describes scaling information of a SKU.
type ResourceSKURestrictions struct {
	// READ-ONLY; The reason for restriction.
	ReasonCode *ResourceSKURestrictionsReasonCode

	// READ-ONLY; The information about the restriction where the SKU cannot be used.
	RestrictionInfo *ResourceSKURestrictionInfo

	// READ-ONLY; The type of restrictions.
	Type *ResourceSKURestrictionsType

	// READ-ONLY; The value of restrictions. If the restriction type is set to location. This would be different locations where
	// the SKU is restricted.
	Values []*string
}

// ResourceSKUZoneDetails - Describes The zonal capabilities of a SKU.
type ResourceSKUZoneDetails struct {
	// READ-ONLY; A list of capabilities that are available for the SKU in the specified list of zones.
	Capabilities []*ResourceSKUCapability

	// READ-ONLY; The set of zones that the SKU is available in with the specified capabilities.
	Name []*string
}

// SKU - Sku for ARM resource
type SKU struct {
	// REQUIRED; Sku name
	Name *string

	// Sku tier
	Tier *string
}

// SystemMetadata - Metadata pertaining to creation and last modification of the resource.
type SystemMetadata struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The type of identity that last modified the resource.
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TrackedResource - The resource model definition for a ARM tracked top level resource.
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives.
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}
