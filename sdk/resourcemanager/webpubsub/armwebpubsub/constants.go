//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
	moduleVersion = "v1.3.0-beta.1"
)

// ACLAction - Azure Networking ACL Action.
type ACLAction string

const (
	ACLActionAllow ACLAction = "Allow"
	ACLActionDeny  ACLAction = "Deny"
)

// PossibleACLActionValues returns the possible values for the ACLAction const type.
func PossibleACLActionValues() []ACLAction {
	return []ACLAction{
		ACLActionAllow,
		ACLActionDeny,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

type EventListenerEndpointDiscriminator string

const (
	EventListenerEndpointDiscriminatorEventHub EventListenerEndpointDiscriminator = "EventHub"
)

// PossibleEventListenerEndpointDiscriminatorValues returns the possible values for the EventListenerEndpointDiscriminator const type.
func PossibleEventListenerEndpointDiscriminatorValues() []EventListenerEndpointDiscriminator {
	return []EventListenerEndpointDiscriminator{
		EventListenerEndpointDiscriminatorEventHub,
	}
}

type EventListenerFilterDiscriminator string

const (
	EventListenerFilterDiscriminatorEventName EventListenerFilterDiscriminator = "EventName"
)

// PossibleEventListenerFilterDiscriminatorValues returns the possible values for the EventListenerFilterDiscriminator const type.
func PossibleEventListenerFilterDiscriminatorValues() []EventListenerFilterDiscriminator {
	return []EventListenerFilterDiscriminator{
		EventListenerFilterDiscriminatorEventName,
	}
}

// KeyType - The type of access key.
type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSalt      KeyType = "Salt"
	KeyTypeSecondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimary,
		KeyTypeSalt,
		KeyTypeSecondary,
	}
}

// ManagedIdentityType - Represents the identity type: systemAssigned, userAssigned, None
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone           ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAssigned ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeUserAssigned   ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeUserAssigned,
	}
}

// PrivateLinkServiceConnectionStatus - Indicates whether the connection has been Approved/Rejected/Removed by the owner of
// the service.
type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

// PossiblePrivateLinkServiceConnectionStatusValues returns the possible values for the PrivateLinkServiceConnectionStatus const type.
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return []PrivateLinkServiceConnectionStatus{
		PrivateLinkServiceConnectionStatusApproved,
		PrivateLinkServiceConnectionStatusDisconnected,
		PrivateLinkServiceConnectionStatusPending,
		PrivateLinkServiceConnectionStatusRejected,
	}
}

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMoving,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
		ProvisioningStateUpdating,
	}
}

// ScaleType - The scale type applicable to the sku.
type ScaleType string

const (
	ScaleTypeAutomatic ScaleType = "Automatic"
	ScaleTypeManual    ScaleType = "Manual"
	ScaleTypeNone      ScaleType = "None"
)

// PossibleScaleTypeValues returns the possible values for the ScaleType const type.
func PossibleScaleTypeValues() []ScaleType {
	return []ScaleType{
		ScaleTypeAutomatic,
		ScaleTypeManual,
		ScaleTypeNone,
	}
}

// ServiceKind - The kind of the service
type ServiceKind string

const (
	ServiceKindSocketIO  ServiceKind = "SocketIO"
	ServiceKindWebPubSub ServiceKind = "WebPubSub"
)

// PossibleServiceKindValues returns the possible values for the ServiceKind const type.
func PossibleServiceKindValues() []ServiceKind {
	return []ServiceKind{
		ServiceKindSocketIO,
		ServiceKindWebPubSub,
	}
}

// SharedPrivateLinkResourceStatus - Status of the shared private link resource
type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
	SharedPrivateLinkResourceStatusTimeout      SharedPrivateLinkResourceStatus = "Timeout"
)

// PossibleSharedPrivateLinkResourceStatusValues returns the possible values for the SharedPrivateLinkResourceStatus const type.
func PossibleSharedPrivateLinkResourceStatusValues() []SharedPrivateLinkResourceStatus {
	return []SharedPrivateLinkResourceStatus{
		SharedPrivateLinkResourceStatusApproved,
		SharedPrivateLinkResourceStatusDisconnected,
		SharedPrivateLinkResourceStatusPending,
		SharedPrivateLinkResourceStatusRejected,
		SharedPrivateLinkResourceStatusTimeout,
	}
}

// UpstreamAuthType - Upstream auth type enum.
type UpstreamAuthType string

const (
	UpstreamAuthTypeManagedIdentity UpstreamAuthType = "ManagedIdentity"
	UpstreamAuthTypeNone            UpstreamAuthType = "None"
)

// PossibleUpstreamAuthTypeValues returns the possible values for the UpstreamAuthType const type.
func PossibleUpstreamAuthTypeValues() []UpstreamAuthType {
	return []UpstreamAuthType{
		UpstreamAuthTypeManagedIdentity,
		UpstreamAuthTypeNone,
	}
}

// WebPubSubRequestType - The incoming request type to the service
type WebPubSubRequestType string

const (
	WebPubSubRequestTypeClientConnection WebPubSubRequestType = "ClientConnection"
	WebPubSubRequestTypeRESTAPI          WebPubSubRequestType = "RESTAPI"
	WebPubSubRequestTypeServerConnection WebPubSubRequestType = "ServerConnection"
	WebPubSubRequestTypeTrace            WebPubSubRequestType = "Trace"
)

// PossibleWebPubSubRequestTypeValues returns the possible values for the WebPubSubRequestType const type.
func PossibleWebPubSubRequestTypeValues() []WebPubSubRequestType {
	return []WebPubSubRequestType{
		WebPubSubRequestTypeClientConnection,
		WebPubSubRequestTypeRESTAPI,
		WebPubSubRequestTypeServerConnection,
		WebPubSubRequestTypeTrace,
	}
}

// WebPubSubSKUTier - Optional tier of this particular SKU. 'Standard' or 'Free'.
// Basic is deprecated, use Standard instead.
type WebPubSubSKUTier string

const (
	WebPubSubSKUTierBasic    WebPubSubSKUTier = "Basic"
	WebPubSubSKUTierFree     WebPubSubSKUTier = "Free"
	WebPubSubSKUTierPremium  WebPubSubSKUTier = "Premium"
	WebPubSubSKUTierStandard WebPubSubSKUTier = "Standard"
)

// PossibleWebPubSubSKUTierValues returns the possible values for the WebPubSubSKUTier const type.
func PossibleWebPubSubSKUTierValues() []WebPubSubSKUTier {
	return []WebPubSubSKUTier{
		WebPubSubSKUTierBasic,
		WebPubSubSKUTierFree,
		WebPubSubSKUTierPremium,
		WebPubSubSKUTierStandard,
	}
}