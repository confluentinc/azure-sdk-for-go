//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

import "time"

// AccessKeys - Namespace/EventHub Connection String
type AccessKeys struct {
	// READ-ONLY; Primary connection string of the alias if GEO DR is enabled
	AliasPrimaryConnectionString *string

	// READ-ONLY; Secondary connection string of the alias if GEO DR is enabled
	AliasSecondaryConnectionString *string

	// READ-ONLY; A string that describes the AuthorizationRule.
	KeyName *string

	// READ-ONLY; Primary connection string of the created namespace AuthorizationRule.
	PrimaryConnectionString *string

	// READ-ONLY; A base64-encoded 256-bit primary key for signing and validating the SAS token.
	PrimaryKey *string

	// READ-ONLY; Secondary connection string of the created namespace AuthorizationRule.
	SecondaryConnectionString *string

	// READ-ONLY; A base64-encoded 256-bit primary key for signing and validating the SAS token.
	SecondaryKey *string
}

// ApplicationGroup - The Application Group object
type ApplicationGroup struct {
	Properties *ApplicationGroupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// ApplicationGroupListResult - The response from the List Application Groups operation.
type ApplicationGroupListResult struct {
	// Result of the List Application Groups operation.
	Value []*ApplicationGroup

	// READ-ONLY; Link to the next set of results. Not empty if Value contains an incomplete list of Authorization Rules
	NextLink *string
}

// ApplicationGroupPolicy - Properties of the Application Group policy
type ApplicationGroupPolicy struct {
	// REQUIRED; The Name of this policy
	Name *string

	// REQUIRED; Application Group Policy types
	Type *ApplicationGroupPolicyType
}

// GetApplicationGroupPolicy implements the ApplicationGroupPolicyClassification interface for type ApplicationGroupPolicy.
func (a *ApplicationGroupPolicy) GetApplicationGroupPolicy() *ApplicationGroupPolicy { return a }

type ApplicationGroupProperties struct {
	// REQUIRED; The Unique identifier for application group.Supports SAS(SASKeyName=KeyName) or AAD(AADAppID=Guid)
	ClientAppGroupIdentifier *string

	// Determines if Application Group is allowed to create connection with namespace or not. Once the isEnabled is set to false,
	// all the existing connections of application group gets dropped and no new
	// connections will be allowed
	IsEnabled *bool

	// List of group policies that define the behavior of application group. The policies can support resource governance scenarios
	// such as limiting ingress or egress traffic.
	Policies []ApplicationGroupPolicyClassification
}

// ArmDisasterRecovery - Single item in List or Get Alias(Disaster Recovery configuration) operation
type ArmDisasterRecovery struct {
	// Properties required to the Create Or Update Alias(Disaster Recovery configurations)
	Properties *ArmDisasterRecoveryProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// ArmDisasterRecoveryListResult - The result of the List Alias(Disaster Recovery configuration) operation.
type ArmDisasterRecoveryListResult struct {
	// List of Alias(Disaster Recovery configurations)
	Value []*ArmDisasterRecovery

	// READ-ONLY; Link to the next set of results. Not empty if Value contains incomplete list of Alias(Disaster Recovery configuration)
	NextLink *string
}

// ArmDisasterRecoveryProperties - Properties required to the Create Or Update Alias(Disaster Recovery configurations)
type ArmDisasterRecoveryProperties struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string

	// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
	PartnerNamespace *string

	// READ-ONLY; Number of entities pending to be replicated.
	PendingReplicationOperationsCount *int64

	// READ-ONLY; Provisioning state of the Alias(Disaster Recovery configuration) - possible values 'Accepted' or 'Succeeded'
	// or 'Failed'
	ProvisioningState *ProvisioningStateDR

	// READ-ONLY; role of namespace in GEO DR - possible values 'Primary' or 'PrimaryNotReplicating' or 'Secondary'
	Role *RoleDisasterRecovery
}

// AuthorizationRule - Single item in a List or Get AuthorizationRule operation
type AuthorizationRule struct {
	// Properties supplied to create or update AuthorizationRule
	Properties *AuthorizationRuleProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// AuthorizationRuleListResult - The response from the List namespace operation.
type AuthorizationRuleListResult struct {
	// Link to the next set of results. Not empty if Value contains an incomplete list of Authorization Rules
	NextLink *string

	// Result of the List Authorization Rules operation.
	Value []*AuthorizationRule
}

// AuthorizationRuleProperties - Properties supplied to create or update AuthorizationRule
type AuthorizationRuleProperties struct {
	// REQUIRED; The rights associated with the rule.
	Rights []*AccessRights
}

// AvailableCluster - Pre-provisioned and readily available Event Hubs Cluster count per region.
type AvailableCluster struct {
	// Location fo the Available Cluster
	Location *string
}

// AvailableClustersList - The response of the List Available Clusters operation.
type AvailableClustersList struct {
	// The count of readily available and pre-provisioned Event Hubs Clusters per region.
	Value []*AvailableCluster
}

// CaptureDescription - Properties to configure capture description for eventhub
type CaptureDescription struct {
	// Properties of Destination where capture will be stored. (Storage Account, Blob Names)
	Destination *Destination

	// A value that indicates whether capture description is enabled.
	Enabled *bool

	// Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be deprecated in
	// New API Version
	Encoding *EncodingCaptureDescription

	// The time window allows you to set the frequency with which the capture to Azure Blobs will happen, value should between
	// 60 to 900 seconds
	IntervalInSeconds *int32

	// The size window defines the amount of data built up in your Event Hub before an capture operation, value should be between
	// 10485760 to 524288000 bytes
	SizeLimitInBytes *int32

	// A value that indicates whether to Skip Empty Archives
	SkipEmptyArchives *bool
}

// CheckNameAvailabilityParameter - Parameter supplied to check Namespace name availability operation
type CheckNameAvailabilityParameter struct {
	// REQUIRED; Name to check the namespace name availability
	Name *string
}

// CheckNameAvailabilityResult - The Result of the CheckNameAvailability operation
type CheckNameAvailabilityResult struct {
	// Value indicating Namespace is availability, true if the Namespace is available; otherwise, false.
	NameAvailable *bool

	// The reason for unavailability of a Namespace.
	Reason *UnavailableReason

	// READ-ONLY; The detailed info regarding the reason associated with the Namespace.
	Message *string
}

// Cluster - Single Event Hubs Cluster resource in List or Get operations.
type Cluster struct {
	// Resource location.
	Location *string

	// Event Hubs Cluster properties supplied in responses in List or Get operations.
	Properties *ClusterProperties

	// Properties of the cluster SKU.
	SKU *ClusterSKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ClusterListResult - The response of the List Event Hubs Clusters operation.
type ClusterListResult struct {
	// Link to the next set of results. Empty unless the value parameter contains an incomplete list of Event Hubs Clusters.
	NextLink *string

	// The Event Hubs Clusters present in the List Event Hubs operation results.
	Value []*Cluster
}

// ClusterProperties - Event Hubs Cluster properties supplied in responses in List or Get operations.
type ClusterProperties struct {
	// A value that indicates whether Scaling is Supported.
	SupportsScaling *bool

	// READ-ONLY; The UTC time when the Event Hubs Cluster was created.
	CreatedAt *string

	// READ-ONLY; The metric ID of the cluster resource. Provided by the service and not modifiable by the user.
	MetricID *string

	// READ-ONLY; Status of the Cluster resource
	Status *string

	// READ-ONLY; The UTC time when the Event Hubs Cluster was last updated.
	UpdatedAt *string
}

// ClusterQuotaConfigurationProperties - Contains all settings for the cluster.
type ClusterQuotaConfigurationProperties struct {
	// All possible Cluster settings - a collection of key/value paired settings which apply to quotas and configurations imposed
	// on the cluster.
	Settings map[string]*string
}

// ClusterSKU - SKU parameters particular to a cluster instance.
type ClusterSKU struct {
	// REQUIRED; Name of this SKU.
	Name *ClusterSKUName

	// The quantity of Event Hubs Cluster Capacity Units contained in this cluster.
	Capacity *int32
}

// ConnectionState information.
type ConnectionState struct {
	// Description of the connection state.
	Description *string

	// Status of the connection.
	Status *PrivateLinkConnectionStatus
}

// ConsumerGroup - Single item in List or Get Consumer group operation
type ConsumerGroup struct {
	// Single item in List or Get Consumer group operation
	Properties *ConsumerGroupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// ConsumerGroupListResult - The result to the List Consumer Group operation.
type ConsumerGroupListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of Consumer Group
	NextLink *string

	// Result of the List Consumer Group operation.
	Value []*ConsumerGroup
}

// ConsumerGroupProperties - Single item in List or Get Consumer group operation
type ConsumerGroupProperties struct {
	// User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be used to store
	// descriptive data, such as list of teams and their contact information also
	// user-defined configuration settings can be stored.
	UserMetadata *string

	// READ-ONLY; Exact time the message was created.
	CreatedAt *time.Time

	// READ-ONLY; The exact time the message was updated.
	UpdatedAt *time.Time
}

// Destination - Capture storage details for capture description
type Destination struct {
	// Name for capture destination
	Name *string

	// Properties describing the storage account, blob container and archive name format for capture destination
	Properties *DestinationProperties
}

// DestinationProperties - Properties describing the storage account, blob container and archive name format for capture destination
type DestinationProperties struct {
	// Blob naming convention for archive, e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}.
	// Here all the parameters (Namespace,EventHub .. etc) are mandatory
	// irrespective of order
	ArchiveNameFormat *string

	// Blob container Name
	BlobContainer *string

	// The Azure Data Lake Store name for the captured events
	DataLakeAccountName *string

	// The destination folder path for the captured events
	DataLakeFolderPath *string

	// Subscription Id of Azure Data Lake Store
	DataLakeSubscriptionID *string

	// Resource id of the storage account to be used to create the blobs
	StorageAccountResourceID *string
}

// EHNamespace - Single Namespace item in List or Get Operation
type EHNamespace struct {
	// Properties of BYOK Identity description
	Identity *Identity

	// Resource location.
	Location *string

	// Namespace properties supplied for create namespace operation.
	Properties *EHNamespaceProperties

	// Properties of sku resource
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EHNamespaceIDContainer - The full ARM ID of an Event Hubs Namespace
type EHNamespaceIDContainer struct {
	// id parameter
	ID *string
}

// EHNamespaceIDListResult - The response of the List Namespace IDs operation
type EHNamespaceIDListResult struct {
	// Result of the List Namespace IDs operation
	Value []*EHNamespaceIDContainer
}

// EHNamespaceListResult - The response of the List Namespace operation
type EHNamespaceListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of namespaces.
	NextLink *string

	// Result of the List Namespace operation
	Value []*EHNamespace
}

// EHNamespaceProperties - Namespace properties supplied for create namespace operation.
type EHNamespaceProperties struct {
	// Alternate name specified when alias and namespace names are same.
	AlternateName *string

	// Cluster ARM ID of the Namespace.
	ClusterArmID *string

	// This property disables SAS authentication for the Event Hubs namespace.
	DisableLocalAuth *bool

	// Properties of BYOK Encryption description
	Encryption *Encryption

	// Value that indicates whether AutoInflate is enabled for eventhub namespace.
	IsAutoInflateEnabled *bool

	// Value that indicates whether Kafka is enabled for eventhub namespace.
	KafkaEnabled *bool

	// Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20 throughput units. ( '0' if
	// AutoInflateEnabled = true)
	MaximumThroughputUnits *int32

	// The minimum TLS version for the cluster to support, e.g. '1.2'
	MinimumTLSVersion *TLSVersion

	// List of private endpoint connections.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// This determines if traffic is allowed over public network. By default it is enabled.
	PublicNetworkAccess *PublicNetworkAccess

	// Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.
	ZoneRedundant *bool

	// READ-ONLY; The time the Namespace was created.
	CreatedAt *time.Time

	// READ-ONLY; Identifier for Azure Insights metrics.
	MetricID *string

	// READ-ONLY; Provisioning state of the Namespace.
	ProvisioningState *string

	// READ-ONLY; Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string

	// READ-ONLY; Status of the Namespace.
	Status *string

	// READ-ONLY; The time the Namespace was updated.
	UpdatedAt *time.Time
}

// Encryption - Properties to configure Encryption
type Encryption struct {
	// Enumerates the possible value of keySource for Encryption
	KeySource *string

	// Properties of KeyVault
	KeyVaultProperties []*KeyVaultProperties

	// Enable Infrastructure Encryption (Double Encryption)
	RequireInfrastructureEncryption *bool
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

// ErrorResponse - Error response indicates Event Hub service is not able to process the incoming request. The reason is provided
// in the error message.
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// Eventhub - Single item in List or Get Event Hub operation
type Eventhub struct {
	// Properties supplied to the Create Or Update Event Hub operation.
	Properties *Properties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// Identity - Properties to configure Identity for Bring your Own Keys
type Identity struct {
	// Type of managed service identity.
	Type *ManagedServiceIdentityType

	// Properties for User Assigned Identities
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; ObjectId from the KeyVault
	PrincipalID *string

	// READ-ONLY; TenantId from the KeyVault
	TenantID *string
}

// KeyVaultProperties - Properties to configure keyVault Properties
type KeyVaultProperties struct {
	Identity *UserAssignedIdentityProperties

	// Name of the Key from KeyVault
	KeyName *string

	// Uri of KeyVault
	KeyVaultURI *string

	// Key Version
	KeyVersion *string
}

// ListResult - The result of the List EventHubs operation.
type ListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of EventHubs.
	NextLink *string

	// Result of the List EventHubs operation.
	Value []*Eventhub
}

// NWRuleSetIPRules - The response from the List namespace operation.
type NWRuleSetIPRules struct {
	// The IP Filter Action
	Action *NetworkRuleIPAction

	// IP Mask
	IPMask *string
}

// NWRuleSetVirtualNetworkRules - The response from the List namespace operation.
type NWRuleSetVirtualNetworkRules struct {
	// Value that indicates whether to ignore missing Vnet Service Endpoint
	IgnoreMissingVnetServiceEndpoint *bool

	// Subnet properties
	Subnet *Subnet
}

// NetworkRuleSet - Description of topic resource.
type NetworkRuleSet struct {
	// NetworkRuleSet properties
	Properties *NetworkRuleSetProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// NetworkRuleSetListResult - The response of the List NetworkRuleSet operation
type NetworkRuleSetListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of NetworkRuleSet.
	NextLink *string

	// Result of the List NetworkRuleSet operation
	Value []*NetworkRuleSet
}

// NetworkRuleSetProperties - NetworkRuleSet properties
type NetworkRuleSetProperties struct {
	// Default Action for Network Rule Set
	DefaultAction *DefaultAction

	// List of IpRules
	IPRules []*NWRuleSetIPRules

	// This determines if traffic is allowed over public network. By default it is enabled. If value is SecuredByPerimeter then
	// Inbound and Outbound communication is controlled by the network security
	// perimeter and profile's access rules.
	PublicNetworkAccess *PublicNetworkAccessFlag

	// Value that indicates whether Trusted Service Access is Enabled or not.
	TrustedServiceAccessEnabled *bool

	// List VirtualNetwork Rules
	VirtualNetworkRules []*NWRuleSetVirtualNetworkRules
}

// NetworkSecurityPerimeter related information
type NetworkSecurityPerimeter struct {
	// Fully qualified identifier of the resource
	ID *string

	// Location of the resource
	Location *string

	// Guid of the resource
	PerimeterGUID *string
}

// NetworkSecurityPerimeterConfiguration - Network Security Perimeter related configurations of a given namespace
type NetworkSecurityPerimeterConfiguration struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Properties of the Network Security Perimeter Configuration
	Properties *NetworkSecurityPerimeterConfigurationProperties

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// NetworkSecurityPerimeterConfigurationList - Result of the List NetworkSecurityPerimeterConfiguration operation.
type NetworkSecurityPerimeterConfigurationList struct {
	// READ-ONLY; A collection of NetworkSecurityPerimeterConfigurations
	Value []*NetworkSecurityPerimeterConfiguration
}

// NetworkSecurityPerimeterConfigurationProperties - Properties of NetworkSecurityPerimeterConfiguration
type NetworkSecurityPerimeterConfigurationProperties struct {
	// List of Provisioning Issues if any
	ProvisioningIssues []*ProvisioningIssue

	// Provisioning state of NetworkSecurityPerimeter configuration propagation
	ProvisioningState *NetworkSecurityPerimeterConfigurationProvisioningState

	// READ-ONLY; NetworkSecurityPerimeter related information
	NetworkSecurityPerimeter *NetworkSecurityPerimeter

	// READ-ONLY; Information about current network profile
	Profile *NetworkSecurityPerimeterConfigurationPropertiesProfile

	// READ-ONLY; Information about resource association
	ResourceAssociation *NetworkSecurityPerimeterConfigurationPropertiesResourceAssociation
}

// NetworkSecurityPerimeterConfigurationPropertiesProfile - Information about current network profile
type NetworkSecurityPerimeterConfigurationPropertiesProfile struct {
	// List of Access Rules
	AccessRules []*NspAccessRule

	// Current access rules version
	AccessRulesVersion *string

	// Name of the resource
	Name *string
}

// NetworkSecurityPerimeterConfigurationPropertiesResourceAssociation - Information about resource association
type NetworkSecurityPerimeterConfigurationPropertiesResourceAssociation struct {
	// Access Mode of the resource association
	AccessMode *ResourceAssociationAccessMode

	// Name of the resource association
	Name *string
}

// NspAccessRule - Information of Access Rule in Network Profile
type NspAccessRule struct {
	// Fully qualified identifier of the resource
	ID *string

	// Name of the resource
	Name *string

	// Type of the resource
	Type *string

	// READ-ONLY; Properties of Access Rule
	Properties *NspAccessRuleProperties
}

// NspAccessRuleProperties - Properties of Access Rule
type NspAccessRuleProperties struct {
	// Address prefixes in the CIDR format for inbound rules
	AddressPrefixes []*string

	// Direction of Access Rule
	Direction *NspAccessRuleDirection

	// Subscriptions for inbound rules
	Subscriptions []*NspAccessRulePropertiesSubscriptionsItem

	// READ-ONLY; FQDN for outbound rules
	FullyQualifiedDomainNames []*string

	// READ-ONLY; NetworkSecurityPerimeters for inbound rules
	NetworkSecurityPerimeters []*NetworkSecurityPerimeter
}

// NspAccessRulePropertiesSubscriptionsItem - Subscription for inbound rule
type NspAccessRulePropertiesSubscriptionsItem struct {
	// Fully qualified identifier of subscription
	ID *string
}

// Operation - A Event Hub REST API operation
type Operation struct {
	// Display of the operation
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Origin of the operation
	Origin *string

	// Properties of the operation
	Properties any

	// READ-ONLY; Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - Operation display payload
type OperationDisplay struct {
	// READ-ONLY; Localized friendly description for the operation
	Description *string

	// READ-ONLY; Localized friendly name for the operation
	Operation *string

	// READ-ONLY; Resource provider of the operation
	Provider *string

	// READ-ONLY; Resource of the operation
	Resource *string
}

// OperationListResult - Result of the request to list Event Hub operations. It contains a list of operations and a URL link
// to get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results if there are any.
	NextLink *string

	// READ-ONLY; List of Event Hub operations supported by the Microsoft.EventHub resource provider.
	Value []*Operation
}

// PrivateEndpoint information.
type PrivateEndpoint struct {
	// The ARM identifier for Private Endpoint.
	ID *string
}

// PrivateEndpointConnection - Properties of the PrivateEndpointConnection.
type PrivateEndpointConnection struct {
	// Properties of the PrivateEndpointConnection.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// PrivateEndpointConnectionListResult - Result of the list of all private endpoint connections operation.
type PrivateEndpointConnectionListResult struct {
	// A link for the next page of private endpoint connection resources.
	NextLink *string

	// A collection of private endpoint connection resources.
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of the private endpoint connection resource.
type PrivateEndpointConnectionProperties struct {
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint *PrivateEndpoint

	// Details about the state of the connection.
	PrivateLinkServiceConnectionState *ConnectionState

	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *EndPointProvisioningState
}

// PrivateLinkResource - Information of the private link resource.
type PrivateLinkResource struct {
	// Fully qualified identifier of the resource.
	ID *string

	// Name of the resource
	Name *string

	// Properties of the private link resource.
	Properties *PrivateLinkResourceProperties

	// Type of the resource
	Type *string
}

// PrivateLinkResourceProperties - Properties of PrivateLinkResource
type PrivateLinkResourceProperties struct {
	// The private link resource group id.
	GroupID *string

	// The private link resource required member names.
	RequiredMembers []*string

	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string
}

// PrivateLinkResourcesListResult - Result of the List private link resources operation.
type PrivateLinkResourcesListResult struct {
	// A link for the next page of private link resources.
	NextLink *string

	// A collection of private link resources
	Value []*PrivateLinkResource
}

// Properties supplied to the Create Or Update Event Hub operation.
type Properties struct {
	// Properties of capture description
	CaptureDescription *CaptureDescription

	// Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *int64

	// Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *int64

	// Event Hub retention settings
	RetentionDescription *RetentionDescription

	// Enumerates the possible values for the status of the Event Hub.
	Status *EntityStatus

	// READ-ONLY; Exact time the Event Hub was created.
	CreatedAt *time.Time

	// READ-ONLY; Current number of shards on the Event Hub.
	PartitionIDs []*string

	// READ-ONLY; The exact time the message was updated.
	UpdatedAt *time.Time
}

// ProvisioningIssue - Describes Provisioning issue for given NetworkSecurityPerimeterConfiguration
type ProvisioningIssue struct {
	// Name of the issue
	Name *string

	// READ-ONLY; Properties of Provisioning Issue
	Properties *ProvisioningIssueProperties
}

// ProvisioningIssueProperties - Properties of Provisioning Issue
type ProvisioningIssueProperties struct {
	// Description of the issue
	Description *string

	// Type of Issue
	IssueType *string
}

// ProxyResource - Common fields that are returned in the response for all Azure Resource Manager resources
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// RegenerateAccessKeyParameters - Parameters supplied to the Regenerate Authorization Rule operation, specifies which key
// needs to be reset.
type RegenerateAccessKeyParameters struct {
	// REQUIRED; The access key to regenerate.
	KeyType *KeyType

	// Optional, if the key value provided, is set for KeyType or autogenerated Key value set for keyType
	Key *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RetentionDescription - Properties to configure retention settings for the eventhub
type RetentionDescription struct {
	// Enumerates the possible values for cleanup policy
	CleanupPolicy *CleanupPolicyRetentionDescription

	// Number of hours to retain the events for this Event Hub. This value is only used when cleanupPolicy is Delete. If cleanupPolicy
	// is Compaction the returned value of this property is Long.MaxValue
	RetentionTimeInHours *int64

	// Number of hours to retain the tombstone markers of a compacted Event Hub. This value is only used when cleanupPolicy is
	// Compaction. Consumer must complete reading the tombstone marker within this
	// specified amount of time if consumer begins from starting offset to ensure they get a valid snapshot for the specific key
	// described by the tombstone marker within the compacted Event Hub
	TombstoneRetentionTimeInHours *int32
}

// SKU parameters supplied to the create namespace operation
type SKU struct {
	// REQUIRED; Name of this SKU.
	Name *SKUName

	// The Event Hubs throughput units for Basic or Standard tiers, where value should be 0 to 20 throughput units. The Event
	// Hubs premium units for Premium tier, where value should be 0 to 10 premium units.
	Capacity *int32

	// The billing tier of this particular SKU.
	Tier *SKUTier
}

// SchemaGroup - Single item in List or Get Schema Group operation
type SchemaGroup struct {
	Properties *SchemaGroupProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string
}

// SchemaGroupListResult - The result of the List SchemaGroup operation.
type SchemaGroupListResult struct {
	// Link to the next set of results. Not empty if Value contains incomplete list of Schema Groups.
	NextLink *string

	// Result of the List SchemaGroups operation.
	Value []*SchemaGroup
}

type SchemaGroupProperties struct {
	// dictionary object for SchemaGroup group properties
	GroupProperties     map[string]*string
	SchemaCompatibility *SchemaCompatibility
	SchemaType          *SchemaType

	// READ-ONLY; Exact time the Schema Group was created.
	CreatedAtUTC *time.Time

	// READ-ONLY; The ETag value.
	ETag *string

	// READ-ONLY; Exact time the Schema Group was updated
	UpdatedAtUTC *time.Time
}

// Subnet - Properties supplied for Subnet
type Subnet struct {
	// Resource ID of Virtual Network Subnet
	ID *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
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

// ThrottlingPolicy - Properties of the throttling policy
type ThrottlingPolicy struct {
	// REQUIRED; Metric Id on which the throttle limit should be set, MetricId can be discovered by hovering over Metric in the
	// Metrics section of Event Hub Namespace inside Azure Portal
	MetricID *MetricID

	// REQUIRED; The Name of this policy
	Name *string

	// REQUIRED; The Threshold limit above which the application group will be throttled.Rate limit is always per second.
	RateLimitThreshold *int64

	// REQUIRED; Application Group Policy types
	Type *ApplicationGroupPolicyType
}

// GetApplicationGroupPolicy implements the ApplicationGroupPolicyClassification interface for type ThrottlingPolicy.
func (t *ThrottlingPolicy) GetApplicationGroupPolicy() *ApplicationGroupPolicy {
	return &ApplicationGroupPolicy{
		Name: t.Name,
		Type: t.Type,
	}
}

// TrackedResource - Definition of resource.
type TrackedResource struct {
	// Resource location.
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UserAssignedIdentity - Recognized Dictionary value.
type UserAssignedIdentity struct {
	// READ-ONLY; Client Id of user assigned identity
	ClientID *string

	// READ-ONLY; Principal Id of user assigned identity
	PrincipalID *string
}

type UserAssignedIdentityProperties struct {
	// ARM ID of user Identity selected for encryption
	UserAssignedIdentity *string
}