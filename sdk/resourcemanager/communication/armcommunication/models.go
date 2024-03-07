//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommunication

import "time"

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string

	// Indicates if the resource name is available.
	NameAvailable *bool

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason
}

// DNSRecord - A class that represents a VerificationStatus record.
type DNSRecord struct {
	// READ-ONLY; Name of the DNS record.
	Name *string

	// READ-ONLY; Represents an expiry time in seconds to represent how long this entry can be cached by the resolver, default
	// = 3600sec.
	TTL *int32

	// READ-ONLY; Type of the DNS record. Example: TXT
	Type *string

	// READ-ONLY; Value of the DNS record.
	Value *string
}

// DomainProperties - A class that describes the properties of a Domains resource.
type DomainProperties struct {
	// REQUIRED; Describes how a Domains resource is being managed.
	DomainManagement *DomainManagement

	// Describes whether user engagement tracking is enabled or disabled.
	UserEngagementTracking *UserEngagementTracking

	// READ-ONLY; The location where the Domains resource data is stored at rest.
	DataLocation *string

	// READ-ONLY; P2 sender domain that is displayed to the email recipients [RFC 5322].
	FromSenderDomain *string

	// READ-ONLY; P1 sender domain that is present on the email envelope [RFC 5321].
	MailFromSenderDomain *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *DomainsProvisioningState

	// READ-ONLY; List of DnsRecord
	VerificationRecords *DomainPropertiesVerificationRecords

	// READ-ONLY; List of VerificationStatusRecord
	VerificationStates *DomainPropertiesVerificationStates
}

// DomainPropertiesVerificationRecords - List of DnsRecord
type DomainPropertiesVerificationRecords struct {
	// A class that represents a VerificationStatus record.
	DKIM *DNSRecord

	// A class that represents a VerificationStatus record.
	DKIM2 *DNSRecord

	// A class that represents a VerificationStatus record.
	DMARC *DNSRecord

	// A class that represents a VerificationStatus record.
	Domain *DNSRecord

	// A class that represents a VerificationStatus record.
	SPF *DNSRecord
}

// DomainPropertiesVerificationStates - List of VerificationStatusRecord
type DomainPropertiesVerificationStates struct {
	// A class that represents a VerificationStatus record.
	DKIM *VerificationStatusRecord

	// A class that represents a VerificationStatus record.
	DKIM2 *VerificationStatusRecord

	// A class that represents a VerificationStatus record.
	DMARC *VerificationStatusRecord

	// A class that represents a VerificationStatus record.
	Domain *VerificationStatusRecord

	// A class that represents a VerificationStatus record.
	SPF *VerificationStatusRecord
}

// DomainResource - A class representing a Domains resource.
type DomainResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The properties of a Domains resource.
	Properties *DomainProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DomainResourceList - Object that includes an array of Domains resource and a possible link for next set.
type DomainResourceList struct {
	// The URL the client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// List of Domains resource
	Value []*DomainResource
}

// EmailServiceProperties - A class that describes the properties of the EmailService.
type EmailServiceProperties struct {
	// REQUIRED; The location where the email service stores its data at rest.
	DataLocation *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *EmailServicesProvisioningState
}

// EmailServiceResource - A class representing an EmailService resource.
type EmailServiceResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The properties of the service.
	Properties *EmailServiceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EmailServiceResourceList - Object that includes an array of EmailServices and a possible link for next set.
type EmailServiceResourceList struct {
	// The URL the client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// List of EmailService
	Value []*EmailServiceResource
}

// EmailServiceResourceUpdate - A class representing update parameters for EmailService resource.
type EmailServiceResourceUpdate struct {
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]*string
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

// LinkNotificationHubParameters - Description of an Azure Notification Hub to link to the communication service
type LinkNotificationHubParameters struct {
	// REQUIRED; Connection string for the notification hub
	ConnectionString *string

	// REQUIRED; The resource ID of the notification hub
	ResourceID *string
}

// LinkedNotificationHub - A notification hub that has been linked to the communication service
type LinkedNotificationHub struct {
	// The resource ID of the notification hub
	ResourceID *string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

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

// NameAvailabilityParameters - Data POST-ed to the nameAvailability action
type NameAvailabilityParameters struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
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

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RegenerateKeyParameters - Parameters describes the request to regenerate access keys
type RegenerateKeyParameters struct {
	// The keyType to regenerate. Must be either 'primary' or 'secondary'(case-insensitive).
	KeyType *KeyType
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SenderUsernameProperties - A class that describes the properties of a SenderUsername resource.
type SenderUsernameProperties struct {
	// REQUIRED; A sender senderUsername to be used when sending emails.
	Username *string

	// The display name for the senderUsername.
	DisplayName *string

	// READ-ONLY; The location where the SenderUsername resource data is stored at rest.
	DataLocation *string

	// READ-ONLY; Provisioning state of the resource. Unknown is the default state for Communication Services.
	ProvisioningState *ProvisioningState
}

// SenderUsernameResource - A class representing a SenderUsername resource.
type SenderUsernameResource struct {
	// The properties of a SenderUsername resource.
	Properties *SenderUsernameProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SenderUsernameResourceCollection - A class representing a Domains SenderUsernames collection.
type SenderUsernameResourceCollection struct {
	// The URL the client should use to fetch the next page (per server side paging).
	NextLink *string

	// List of SenderUsernames
	Value []*SenderUsernameResource
}

// ServiceKeys - A class representing the access keys of a CommunicationService.
type ServiceKeys struct {
	// CommunicationService connection string constructed via the primaryKey
	PrimaryConnectionString *string

	// The primary access key.
	PrimaryKey *string

	// CommunicationService connection string constructed via the secondaryKey
	SecondaryConnectionString *string

	// The secondary access key.
	SecondaryKey *string
}

// ServiceProperties - A class that describes the properties of the CommunicationService.
type ServiceProperties struct {
	// REQUIRED; The location where the communication service stores its data at rest.
	DataLocation *string

	// List of email Domain resource Ids.
	LinkedDomains []*string

	// READ-ONLY; FQDN of the CommunicationService instance.
	HostName *string

	// READ-ONLY; The immutable resource Id of the communication service.
	ImmutableResourceID *string

	// READ-ONLY; Resource ID of an Azure Notification Hub linked to this resource.
	NotificationHubID *string

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *CommunicationServicesProvisioningState

	// READ-ONLY; Version of the CommunicationService resource. Probably you need the same or higher version of client SDKs.
	Version *string
}

// ServiceResource - A class representing a CommunicationService resource.
type ServiceResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// The properties of the service.
	Properties *ServiceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServiceResourceList - Object that includes an array of CommunicationServices and a possible link for next set.
type ServiceResourceList struct {
	// The URL the client should use to fetch the next page (per server side paging). It's null for now, added for future use.
	NextLink *string

	// List of CommunicationService
	Value []*ServiceResource
}

// ServiceResourceUpdate - A class representing update parameters for CommunicationService resource.
type ServiceResourceUpdate struct {
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// The properties of the service.
	Properties *ServiceUpdateProperties

	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]*string
}

// ServiceUpdateProperties - A class that describes the properties that can be updated for CommunicationService resource.
type ServiceUpdateProperties struct {
	// List of email Domain resource Ids.
	LinkedDomains []*string
}

// SuppressionListAddressProperties - A class that describes the properties of a SuppressionListAddress resource.
type SuppressionListAddressProperties struct {
	// REQUIRED; Email address of the recipient.
	Email *string

	// The first name of the email recipient.
	FirstName *string

	// The last name of the email recipient.
	LastName *string

	// An optional property to provide contextual notes or a description for an address.
	Notes *string

	// READ-ONLY; The location where the SuppressionListAddress data is stored at rest. This value is inherited from the parent
	// Domains resource.
	DataLocation *string

	// READ-ONLY; The date the address was last updated in a suppression list.
	LastModified *time.Time
}

// SuppressionListAddressResource - A object that represents a SuppressionList record.
type SuppressionListAddressResource struct {
	// The properties of a SuppressionListAddress resource.
	Properties *SuppressionListAddressProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SuppressionListAddressResourceCollection - Collection of addresses in a suppression list. Response will include a nextLink
// if response contains more pages.
type SuppressionListAddressResourceCollection struct {
	// The URL the client should use to fetch the next page (per server side paging).
	NextLink *string

	// List of suppressed email addresses.
	Value []*SuppressionListAddressResource
}

// SuppressionListProperties - A class that describes the properties of a SuppressionList resource.
type SuppressionListProperties struct {
	// The the name of the suppression list. This value must match one of the valid sender usernames of the sending domain.
	ListName *string

	// READ-ONLY; The date the resource was created.
	CreatedTimeStamp *string

	// READ-ONLY; The location where the SuppressionListAddress data is stored at rest. This value is inherited from the parent
	// Domains resource.
	DataLocation *string

	// READ-ONLY; The date the resource was last updated.
	LastUpdatedTimeStamp *string
}

// SuppressionListResource - A class representing a SuppressionList resource.
type SuppressionListResource struct {
	// The properties of a SuppressionList resource.
	Properties *SuppressionListProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SuppressionListResourceCollection - A class representing a Domains SuppressionListResource collection.
type SuppressionListResourceCollection struct {
	// The URL the client should use to fetch the next page (per server side paging).
	NextLink *string

	// List of SuppressionListResource
	Value []*SuppressionListResource
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

// TaggedResource - An ARM resource with that can accept tags
type TaggedResource struct {
	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]*string
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// UpdateDomainProperties - A class that describes the updatable properties of a Domains resource.
type UpdateDomainProperties struct {
	// Describes whether user engagement tracking is enabled or disabled.
	UserEngagementTracking *UserEngagementTracking
}

// UpdateDomainRequestParameters - A class that describes the PATCH request parameters of a Domains resource.
type UpdateDomainRequestParameters struct {
	// A class that describes the updatable properties of a Domains resource.
	Properties *UpdateDomainProperties

	// Tags of the service which is a list of key value pairs that describe the resource.
	Tags map[string]*string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// VerificationParameter - Input parameter for verification APIs
type VerificationParameter struct {
	// REQUIRED; Type of verification.
	VerificationType *VerificationType
}

// VerificationStatusRecord - A class that represents a VerificationStatus record.
type VerificationStatusRecord struct {
	// READ-ONLY; Error code. This property will only be present if the status is UnableToVerify.
	ErrorCode *string

	// READ-ONLY; Status of the verification operation.
	Status *VerificationStatus
}