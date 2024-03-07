//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

import "time"

// CheckNameAvailabilityParameters - Input values.
type CheckNameAvailabilityParameters struct {
	// REQUIRED; The name of the service instance to check.
	Name *string

	// REQUIRED; The fully qualified resource type which includes provider namespace.
	Type *string
}

// CorsConfiguration - The settings for the CORS configuration of the service instance.
type CorsConfiguration struct {
	// If credentials are allowed via CORS.
	AllowCredentials *bool

	// The headers to be allowed via CORS.
	Headers []*string

	// The max age to be allowed via CORS.
	MaxAge *int32

	// The methods to be allowed via CORS.
	Methods []*string

	// The origins to be allowed via CORS.
	Origins []*string
}

// DicomService - The description of Dicom Service
type DicomService struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity

	// The resource location.
	Location *string

	// Dicom Service configuration.
	Properties *DicomServiceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// DicomServiceAuthenticationConfiguration - Authentication configuration information
type DicomServiceAuthenticationConfiguration struct {
	// READ-ONLY; The audiences for the service
	Audiences []*string

	// READ-ONLY; The authority url for the service
	Authority *string
}

// DicomServiceCollection - The collection of Dicom Services.
type DicomServiceCollection struct {
	// The link used to get the next page of Dicom Services.
	NextLink *string

	// The list of Dicom Services.
	Value []*DicomService
}

// DicomServicePatchResource - Dicom Service patch properties
type DicomServicePatchResource struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity

	// Resource tags.
	Tags map[string]*string
}

// DicomServiceProperties - Dicom Service properties.
type DicomServiceProperties struct {
	// Dicom Service authentication configuration.
	AuthenticationConfiguration *DicomServiceAuthenticationConfiguration

	// Dicom Service Cors configuration.
	CorsConfiguration *CorsConfiguration

	// The encryption settings of the DICOM service
	Encryption *Encryption

	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess

	// READ-ONLY; DICOM Service event support status.
	EventState *ServiceEventState

	// READ-ONLY; The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The url of the Dicom Services.
	ServiceURL *string
}

// Encryption - Settings to encrypt a service
type Encryption struct {
	// The encryption settings for the customer-managed key
	CustomerManagedKeyEncryption *EncryptionCustomerManagedKeyEncryption
}

// EncryptionCustomerManagedKeyEncryption - The encryption settings for the customer-managed key
type EncryptionCustomerManagedKeyEncryption struct {
	// The URL of the key to use for encryption
	KeyEncryptionKeyURL *string
}

// Error details.
type Error struct {
	// Error details
	Error *ErrorDetailsInternal
}

// ErrorDetails - Error details.
type ErrorDetails struct {
	// Error details
	Error *ErrorDetailsInternal
}

// ErrorDetailsInternal - Error details.
type ErrorDetailsInternal struct {
	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The target of the particular error.
	Target *string
}

// FhirService - The description of Fhir Service
type FhirService struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity

	// The kind of the service.
	Kind *FhirServiceKind

	// The resource location.
	Location *string

	// Fhir Service configuration.
	Properties *FhirServiceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// FhirServiceAcrConfiguration - Azure container registry configuration information
type FhirServiceAcrConfiguration struct {
	// The list of the Azure container registry login servers.
	LoginServers []*string

	// The list of Open Container Initiative (OCI) artifacts.
	OciArtifacts []*ServiceOciArtifactEntry
}

// FhirServiceAuthenticationConfiguration - Authentication configuration information
type FhirServiceAuthenticationConfiguration struct {
	// The audience url for the service
	Audience *string

	// The authority url for the service
	Authority *string

	// If the SMART on FHIR proxy is enabled
	SmartProxyEnabled *bool
}

// FhirServiceCollection - A collection of Fhir services.
type FhirServiceCollection struct {
	// The link used to get the next page of Fhir Services.
	NextLink *string

	// The list of Fhir Services.
	Value []*FhirService
}

// FhirServiceCorsConfiguration - The settings for the CORS configuration of the service instance.
type FhirServiceCorsConfiguration struct {
	// If credentials are allowed via CORS.
	AllowCredentials *bool

	// The headers to be allowed via CORS.
	Headers []*string

	// The max age to be allowed via CORS.
	MaxAge *int32

	// The methods to be allowed via CORS.
	Methods []*string

	// The origins to be allowed via CORS.
	Origins []*string
}

// FhirServiceExportConfiguration - Export operation configuration information
type FhirServiceExportConfiguration struct {
	// The name of the default export storage account.
	StorageAccountName *string
}

// FhirServiceImportConfiguration - Import operation configuration information
type FhirServiceImportConfiguration struct {
	// If the import operation is enabled.
	Enabled *bool

	// If the FHIR service is in InitialImportMode.
	InitialImportMode *bool

	// The name of the default integration storage account.
	IntegrationDataStore *string
}

// FhirServicePatchResource - FhirService patch properties
type FhirServicePatchResource struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity

	// Resource tags.
	Tags map[string]*string
}

// FhirServiceProperties - Fhir Service properties.
type FhirServiceProperties struct {
	// Fhir Service Azure container registry configuration.
	AcrConfiguration *FhirServiceAcrConfiguration

	// Fhir Service authentication configuration.
	AuthenticationConfiguration *FhirServiceAuthenticationConfiguration

	// Fhir Service Cors configuration.
	CorsConfiguration *FhirServiceCorsConfiguration

	// The encryption settings of the FHIR service
	Encryption *Encryption

	// Fhir Service export configuration.
	ExportConfiguration *FhirServiceExportConfiguration

	// Implementation Guides configuration.
	ImplementationGuidesConfiguration *ImplementationGuidesConfiguration

	// Fhir Service import configuration.
	ImportConfiguration *FhirServiceImportConfiguration

	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess

	// Determines tracking of history for resources.
	ResourceVersionPolicyConfiguration *ResourceVersionPolicyConfiguration

	// READ-ONLY; Fhir Service event support status.
	EventState *ServiceEventState

	// READ-ONLY; The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState
}

// ImplementationGuidesConfiguration - The settings for Implementation Guides - defining capabilities for national standards,
// vendor consortiums, clinical societies, etc.
type ImplementationGuidesConfiguration struct {
	// If US Core Missing Data requirement is enabled.
	UsCoreMissingData *bool
}

// IotConnector - IoT Connector definition.
type IotConnector struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity

	// The resource location.
	Location *string

	// IoT Connector configuration.
	Properties *IotConnectorProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// IotConnectorCollection - A collection of IoT Connectors.
type IotConnectorCollection struct {
	// The link used to get the next page of IoT Connectors.
	NextLink *string

	// The list of IoT Connectors.
	Value []*IotConnector
}

// IotConnectorPatchResource - Iot Connector patch properties
type IotConnectorPatchResource struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity

	// Resource tags.
	Tags map[string]*string
}

// IotConnectorProperties - IoT Connector properties.
type IotConnectorProperties struct {
	// Device Mappings.
	DeviceMapping *IotMappingProperties

	// Source configuration.
	IngestionEndpointConfiguration *IotEventHubIngestionEndpointConfiguration

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState
}

// IotDestinationProperties - Common IoT Connector destination properties.
type IotDestinationProperties struct {
	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState
}

// IotEventHubIngestionEndpointConfiguration - Event Hub ingestion endpoint configuration
type IotEventHubIngestionEndpointConfiguration struct {
	// Consumer group of the event hub to connected to.
	ConsumerGroup *string

	// Event Hub name to connect to.
	EventHubName *string

	// Fully qualified namespace of the Event Hub to connect to.
	FullyQualifiedEventHubNamespace *string
}

// IotFhirDestination - IoT Connector FHIR destination definition.
type IotFhirDestination struct {
	// REQUIRED; IoT FHIR Destination settings.
	Properties *IotFhirDestinationProperties

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// The resource location.
	Location *string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// IotFhirDestinationCollection - A collection of IoT Connector FHIR destinations.
type IotFhirDestinationCollection struct {
	// The link used to get the next page of IoT FHIR destinations.
	NextLink *string

	// The list of IoT Connector FHIR destinations.
	Value []*IotFhirDestination
}

// IotFhirDestinationProperties - IoT Connector destination properties for an Azure FHIR service.
type IotFhirDestinationProperties struct {
	// REQUIRED; FHIR Mappings
	FhirMapping *IotMappingProperties

	// REQUIRED; Fully qualified resource id of the FHIR service to connect to.
	FhirServiceResourceID *string

	// REQUIRED; Determines how resource identity is resolved on the destination.
	ResourceIdentityResolutionType *IotIdentityResolutionType

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState
}

// IotMappingProperties - The mapping content.
type IotMappingProperties struct {
	// The mapping.
	Content any
}

// ListOperations - Available operations of the service
type ListOperations struct {
	// URL client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// READ-ONLY; Collection of available operation details
	Value []*OperationDetail
}

// LocationBasedResource - The common properties for any location based resource, tracked or proxy.
type LocationBasedResource struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// The resource location.
	Location *string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// LogSpecification - Specifications of the Log for Azure Monitoring
type LogSpecification struct {
	// Blob duration of the log
	BlobDuration *string

	// Localized friendly display name of the log
	DisplayName *string

	// Name of the log
	Name *string
}

// MetricDimension - Specifications of the Dimension of metrics
type MetricDimension struct {
	// Localized friendly display name of the dimension
	DisplayName *string

	// Name of the dimension
	Name *string

	// Whether this dimension should be included for the Shoebox export scenario
	ToBeExportedForShoebox *bool
}

// MetricSpecification - Specifications of the Metrics for Azure Monitoring
type MetricSpecification struct {
	// Only provide one value for this field. Valid values: Average, Minimum, Maximum, Total, Count.
	AggregationType *string

	// Name of the metric category that the metric belongs to. A metric can only belong to a single category.
	Category *string

	// Dimensions of the metric
	Dimensions []*MetricDimension

	// Localized friendly description of the metric
	DisplayDescription *string

	// Localized friendly display name of the metric
	DisplayName *string

	// Whether regional MDM account enabled.
	EnableRegionalMdmAccount *bool

	// Optional. If set to true, then zero will be returned for time duration where no metric is emitted/published.
	FillGapWithZero *bool

	// Whether the metric is internal.
	IsInternal *bool

	// Pattern for the filter of the metric.
	MetricFilterPattern *string

	// Name of the metric
	Name *string

	// The resource Id dimension name override.
	ResourceIDDimensionNameOverride *string

	// The source MDM account.
	SourceMdmAccount *string

	// The source MDM namespace.
	SourceMdmNamespace *string

	// Supported aggregation types
	SupportedAggregationTypes []*string

	// Supported time grain types
	SupportedTimeGrainTypes []*string

	// Unit that makes sense for the metric
	Unit *string
}

// OperationDetail - Service REST API operation.
type OperationDetail struct {
	// Display of the operation
	Display *OperationDisplay

	// Properties of the operation
	Properties *OperationProperties

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; Name of the operation
	Name *string

	// READ-ONLY; Default value is 'user,system'.
	Origin *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// READ-ONLY; Friendly description for the operation,
	Description *string

	// READ-ONLY; Name of the operation
	Operation *string

	// READ-ONLY; Service provider: Microsoft.HealthcareApis
	Provider *string

	// READ-ONLY; Resource Type: Services
	Resource *string
}

// OperationProperties - Extra Operation properties
type OperationProperties struct {
	// Service specifications of the operation
	ServiceSpecification *ServiceSpecification
}

// OperationResultsDescription - The properties indicating the operation result of an operation on a service.
type OperationResultsDescription struct {
	// Additional properties of the operation result.
	Properties any

	// READ-ONLY; The time that the operation finished.
	EndTime *string

	// READ-ONLY; The ID of the operation returned.
	ID *string

	// READ-ONLY; The name of the operation result.
	Name *string

	// READ-ONLY; The time that the operation was started.
	StartTime *string

	// READ-ONLY; The status of the operation being performed.
	Status *OperationResultStatus
}

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string
}

// PrivateEndpointConnection - The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionDescription - The Private Endpoint Connection resource.
type PrivateEndpointConnectionDescription struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - List of private endpoint connection associated with the specified storage account
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionListResultDescription - List of private endpoint connection associated with the specified storage
// account
type PrivateEndpointConnectionListResultDescription struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnectionDescription
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceDescription - The Private Endpoint Connection resource.
type PrivateLinkResourceDescription struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceListResultDescription - A list of private link resources
type PrivateLinkResourceListResultDescription struct {
	// Array of private link resources
	Value []*PrivateLinkResourceDescription
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string

	// The reason for approval/rejection of the connection.
	Description *string

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus
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

// ResourceCore - The common properties for any resource, tracked or proxy.
type ResourceCore struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// ResourceTags - List of key value pairs that describe the resource. This will overwrite the existing tags.
type ResourceTags struct {
	// Resource tags.
	Tags map[string]*string
}

// ResourceVersionPolicyConfiguration - The settings for history tracking for FHIR resources.
type ResourceVersionPolicyConfiguration struct {
	// The default value for tracking history across all resources.
	Default *FhirResourceVersionPolicy

	// A list of FHIR Resources and their version policy overrides.
	ResourceTypeOverrides map[string]*FhirResourceVersionPolicy
}

// ServiceAccessPolicyEntry - An access policy entry.
type ServiceAccessPolicyEntry struct {
	// REQUIRED; An Azure AD object ID (User or Apps) that is allowed access to the FHIR service.
	ObjectID *string
}

// ServiceAcrConfigurationInfo - Azure container registry configuration information
type ServiceAcrConfigurationInfo struct {
	// The list of the ACR login servers.
	LoginServers []*string

	// The list of Open Container Initiative (OCI) artifacts.
	OciArtifacts []*ServiceOciArtifactEntry
}

// ServiceAuthenticationConfigurationInfo - Authentication configuration information
type ServiceAuthenticationConfigurationInfo struct {
	// The audience url for the service
	Audience *string

	// The authority url for the service
	Authority *string

	// If the SMART on FHIR proxy is enabled
	SmartProxyEnabled *bool
}

// ServiceCorsConfigurationInfo - The settings for the CORS configuration of the service instance.
type ServiceCorsConfigurationInfo struct {
	// If credentials are allowed via CORS.
	AllowCredentials *bool

	// The headers to be allowed via CORS.
	Headers []*string

	// The max age to be allowed via CORS.
	MaxAge *int32

	// The methods to be allowed via CORS.
	Methods []*string

	// The origins to be allowed via CORS.
	Origins []*string
}

// ServiceCosmosDbConfigurationInfo - The settings for the Cosmos DB database backing the service.
type ServiceCosmosDbConfigurationInfo struct {
	// The multi-tenant application id used to enable CMK access for services in a data sovereign region.
	CrossTenantCmkApplicationID *string

	// The URI of the customer-managed key for the backing database.
	KeyVaultKeyURI *string

	// The provisioned throughput for the backing database.
	OfferThroughput *int32
}

// ServiceExportConfigurationInfo - Export operation configuration information
type ServiceExportConfigurationInfo struct {
	// The name of the default export storage account.
	StorageAccountName *string
}

// ServiceImportConfigurationInfo - Import operation configuration information
type ServiceImportConfigurationInfo struct {
	// If the import operation is enabled.
	Enabled *bool

	// If the FHIR service is in InitialImportMode.
	InitialImportMode *bool

	// The name of the default integration storage account.
	IntegrationDataStore *string
}

// ServiceManagedIdentity - Managed service identity (system assigned and/or user assigned identities)
type ServiceManagedIdentity struct {
	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServiceManagedIdentityIdentity
}

// ServiceManagedIdentityIdentity - Setting indicating whether the service has a managed identity associated with it.
type ServiceManagedIdentityIdentity struct {
	// REQUIRED; Type of identity being specified, currently SystemAssigned and None are allowed.
	Type *ServiceManagedIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// ServiceOciArtifactEntry - An Open Container Initiative (OCI) artifact.
type ServiceOciArtifactEntry struct {
	// The artifact digest.
	Digest *string

	// The artifact name.
	ImageName *string

	// The Azure Container Registry login server.
	LoginServer *string
}

// ServiceSpecification - Service specification payload
type ServiceSpecification struct {
	// Specifications of the Log for Azure Monitoring
	LogSpecifications []*LogSpecification

	// Specifications of the Metrics for Azure Monitoring
	MetricSpecifications []*MetricSpecification
}

// ServicesDescription - The description of the service.
type ServicesDescription struct {
	// REQUIRED; The kind of the service.
	Kind *Kind

	// REQUIRED; The resource location.
	Location *string

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceIdentity

	// The common properties of a service.
	Properties *ServicesProperties

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// ServicesDescriptionListResult - A list of service description objects with a next link.
type ServicesDescriptionListResult struct {
	// The link used to get the next page of service description objects.
	NextLink *string

	// A list of service description objects.
	Value []*ServicesDescription
}

// ServicesNameAvailabilityInfo - The properties indicating whether a given service name is available.
type ServicesNameAvailabilityInfo struct {
	// The detailed reason message.
	Message *string

	// READ-ONLY; The value which indicates whether the provided name is available.
	NameAvailable *bool

	// READ-ONLY; The reason for unavailability.
	Reason *ServiceNameUnavailabilityReason
}

// ServicesPatchDescription - The description of the service.
type ServicesPatchDescription struct {
	// The properties for updating a service instance.
	Properties *ServicesPropertiesUpdateParameters

	// Instance tags
	Tags map[string]*string
}

// ServicesProperties - The properties of a service instance.
type ServicesProperties struct {
	// The access policies of the service instance.
	AccessPolicies []*ServiceAccessPolicyEntry

	// The azure container registry settings used for convert data operation of the service instance.
	AcrConfiguration *ServiceAcrConfigurationInfo

	// The authentication configuration for the service instance.
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfo

	// The settings for the CORS configuration of the service instance.
	CorsConfiguration *ServiceCorsConfigurationInfo

	// The settings for the Cosmos DB database backing the service.
	CosmosDbConfiguration *ServiceCosmosDbConfigurationInfo

	// The settings for the export operation of the service instance.
	ExportConfiguration *ServiceExportConfigurationInfo

	// The settings for the import operation of the service instance.
	ImportConfiguration *ServiceImportConfigurationInfo

	// The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState
}

// ServicesPropertiesUpdateParameters - The properties for updating a service instance.
type ServicesPropertiesUpdateParameters struct {
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess
}

// ServicesResource - The common properties of a service.
type ServicesResource struct {
	// REQUIRED; The kind of the service.
	Kind *Kind

	// REQUIRED; The resource location.
	Location *string

	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// Setting indicating whether the service has a managed identity associated with it.
	Identity *ServicesResourceIdentity

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// ServicesResourceIdentity - Setting indicating whether the service has a managed identity associated with it.
type ServicesResourceIdentity struct {
	// Type of identity being specified, currently SystemAssigned and None are allowed.
	Type *ManagedServiceIdentityType

	// READ-ONLY; The principal ID of the resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the resource.
	TenantID *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TaggedResource - The common properties of tracked resources in the service.
type TaggedResource struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// The resource location.
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; The resource type.
	Type *string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// Workspace resource.
type Workspace struct {
	// An etag associated with the resource, used for optimistic concurrency when editing it.
	Etag *string

	// The resource location.
	Location *string

	// Workspaces resource specific properties.
	Properties *WorkspaceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; The resource identifier.
	ID *string

	// READ-ONLY; The resource name.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The resource type.
	Type *string
}

// WorkspaceList - Collection of workspace object with a next link
type WorkspaceList struct {
	// The link used to get the next page.
	NextLink *string

	// Collection of resources.
	Value []*Workspace
}

// WorkspacePatchResource - Workspace patch properties
type WorkspacePatchResource struct {
	// Resource tags.
	Tags map[string]*string
}

// WorkspaceProperties - Workspaces resource specific properties.
type WorkspaceProperties struct {
	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess

	// READ-ONLY; The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The provisioning state.
	ProvisioningState *ProvisioningState
}