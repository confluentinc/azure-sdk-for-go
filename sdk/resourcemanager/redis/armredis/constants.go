//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
	moduleVersion = "v3.1.0"
)

// AccessPolicyAssignmentProvisioningState - Provisioning state of an access policy assignment set
type AccessPolicyAssignmentProvisioningState string

const (
	AccessPolicyAssignmentProvisioningStateCanceled  AccessPolicyAssignmentProvisioningState = "Canceled"
	AccessPolicyAssignmentProvisioningStateDeleted   AccessPolicyAssignmentProvisioningState = "Deleted"
	AccessPolicyAssignmentProvisioningStateDeleting  AccessPolicyAssignmentProvisioningState = "Deleting"
	AccessPolicyAssignmentProvisioningStateFailed    AccessPolicyAssignmentProvisioningState = "Failed"
	AccessPolicyAssignmentProvisioningStateSucceeded AccessPolicyAssignmentProvisioningState = "Succeeded"
	AccessPolicyAssignmentProvisioningStateUpdating  AccessPolicyAssignmentProvisioningState = "Updating"
)

// PossibleAccessPolicyAssignmentProvisioningStateValues returns the possible values for the AccessPolicyAssignmentProvisioningState const type.
func PossibleAccessPolicyAssignmentProvisioningStateValues() []AccessPolicyAssignmentProvisioningState {
	return []AccessPolicyAssignmentProvisioningState{
		AccessPolicyAssignmentProvisioningStateCanceled,
		AccessPolicyAssignmentProvisioningStateDeleted,
		AccessPolicyAssignmentProvisioningStateDeleting,
		AccessPolicyAssignmentProvisioningStateFailed,
		AccessPolicyAssignmentProvisioningStateSucceeded,
		AccessPolicyAssignmentProvisioningStateUpdating,
	}
}

// AccessPolicyProvisioningState - Provisioning state of access policy
type AccessPolicyProvisioningState string

const (
	AccessPolicyProvisioningStateCanceled  AccessPolicyProvisioningState = "Canceled"
	AccessPolicyProvisioningStateDeleted   AccessPolicyProvisioningState = "Deleted"
	AccessPolicyProvisioningStateDeleting  AccessPolicyProvisioningState = "Deleting"
	AccessPolicyProvisioningStateFailed    AccessPolicyProvisioningState = "Failed"
	AccessPolicyProvisioningStateSucceeded AccessPolicyProvisioningState = "Succeeded"
	AccessPolicyProvisioningStateUpdating  AccessPolicyProvisioningState = "Updating"
)

// PossibleAccessPolicyProvisioningStateValues returns the possible values for the AccessPolicyProvisioningState const type.
func PossibleAccessPolicyProvisioningStateValues() []AccessPolicyProvisioningState {
	return []AccessPolicyProvisioningState{
		AccessPolicyProvisioningStateCanceled,
		AccessPolicyProvisioningStateDeleted,
		AccessPolicyProvisioningStateDeleting,
		AccessPolicyProvisioningStateFailed,
		AccessPolicyProvisioningStateSucceeded,
		AccessPolicyProvisioningStateUpdating,
	}
}

// AccessPolicyType - Built-In or Custom access policy
type AccessPolicyType string

const (
	AccessPolicyTypeBuiltIn AccessPolicyType = "BuiltIn"
	AccessPolicyTypeCustom  AccessPolicyType = "Custom"
)

// PossibleAccessPolicyTypeValues returns the possible values for the AccessPolicyType const type.
func PossibleAccessPolicyTypeValues() []AccessPolicyType {
	return []AccessPolicyType{
		AccessPolicyTypeBuiltIn,
		AccessPolicyTypeCustom,
	}
}

// DayOfWeek - Day of the week when a cache can be patched.
type DayOfWeek string

const (
	DayOfWeekEveryday  DayOfWeek = "Everyday"
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
	DayOfWeekWeekend   DayOfWeek = "Weekend"
)

// PossibleDayOfWeekValues returns the possible values for the DayOfWeek const type.
func PossibleDayOfWeekValues() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekEveryday,
		DayOfWeekFriday,
		DayOfWeekMonday,
		DayOfWeekSaturday,
		DayOfWeekSunday,
		DayOfWeekThursday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
		DayOfWeekWeekend,
	}
}

type DefaultName string

const (
	DefaultNameDefault DefaultName = "default"
)

// PossibleDefaultNameValues returns the possible values for the DefaultName const type.
func PossibleDefaultNameValues() []DefaultName {
	return []DefaultName{
		DefaultNameDefault,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningState - Redis instance provisioning status.
type ProvisioningState string

const (
	ProvisioningStateConfiguringAAD         ProvisioningState = "ConfiguringAAD"
	ProvisioningStateCreating               ProvisioningState = "Creating"
	ProvisioningStateDeleting               ProvisioningState = "Deleting"
	ProvisioningStateDisabled               ProvisioningState = "Disabled"
	ProvisioningStateFailed                 ProvisioningState = "Failed"
	ProvisioningStateLinking                ProvisioningState = "Linking"
	ProvisioningStateProvisioning           ProvisioningState = "Provisioning"
	ProvisioningStateRecoveringScaleFailure ProvisioningState = "RecoveringScaleFailure"
	ProvisioningStateScaling                ProvisioningState = "Scaling"
	ProvisioningStateSucceeded              ProvisioningState = "Succeeded"
	ProvisioningStateUnlinking              ProvisioningState = "Unlinking"
	ProvisioningStateUnprovisioning         ProvisioningState = "Unprovisioning"
	ProvisioningStateUpdating               ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateConfiguringAAD,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateDisabled,
		ProvisioningStateFailed,
		ProvisioningStateLinking,
		ProvisioningStateProvisioning,
		ProvisioningStateRecoveringScaleFailure,
		ProvisioningStateScaling,
		ProvisioningStateSucceeded,
		ProvisioningStateUnlinking,
		ProvisioningStateUnprovisioning,
		ProvisioningStateUpdating,
	}
}

// PublicNetworkAccess - Whether or not public endpoint access is allowed for this cache. Value is optional but if passed
// in, must be 'Enabled' or 'Disabled'. If 'Disabled', private endpoints are the exclusive access method.
// Default value is 'Enabled'
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// RebootType - Which Redis node(s) to reboot. Depending on this value data loss is possible.
type RebootType string

const (
	RebootTypeAllNodes      RebootType = "AllNodes"
	RebootTypePrimaryNode   RebootType = "PrimaryNode"
	RebootTypeSecondaryNode RebootType = "SecondaryNode"
)

// PossibleRebootTypeValues returns the possible values for the RebootType const type.
func PossibleRebootTypeValues() []RebootType {
	return []RebootType{
		RebootTypeAllNodes,
		RebootTypePrimaryNode,
		RebootTypeSecondaryNode,
	}
}

// RedisKeyType - The Redis access key to regenerate.
type RedisKeyType string

const (
	RedisKeyTypePrimary   RedisKeyType = "Primary"
	RedisKeyTypeSecondary RedisKeyType = "Secondary"
)

// PossibleRedisKeyTypeValues returns the possible values for the RedisKeyType const type.
func PossibleRedisKeyTypeValues() []RedisKeyType {
	return []RedisKeyType{
		RedisKeyTypePrimary,
		RedisKeyTypeSecondary,
	}
}

// ReplicationRole - Role of the linked server.
type ReplicationRole string

const (
	ReplicationRolePrimary   ReplicationRole = "Primary"
	ReplicationRoleSecondary ReplicationRole = "Secondary"
)

// PossibleReplicationRoleValues returns the possible values for the ReplicationRole const type.
func PossibleReplicationRoleValues() []ReplicationRole {
	return []ReplicationRole{
		ReplicationRolePrimary,
		ReplicationRoleSecondary,
	}
}

// SKUFamily - The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).
type SKUFamily string

const (
	SKUFamilyC SKUFamily = "C"
	SKUFamilyP SKUFamily = "P"
)

// PossibleSKUFamilyValues returns the possible values for the SKUFamily const type.
func PossibleSKUFamilyValues() []SKUFamily {
	return []SKUFamily{
		SKUFamilyC,
		SKUFamilyP,
	}
}

// SKUName - The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)
type SKUName string

const (
	SKUNameBasic    SKUName = "Basic"
	SKUNamePremium  SKUName = "Premium"
	SKUNameStandard SKUName = "Standard"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNamePremium,
		SKUNameStandard,
	}
}

// TLSVersion - Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, '1.0', '1.1', '1.2')
type TLSVersion string

const (
	TLSVersionOne0 TLSVersion = "1.0"
	TLSVersionOne1 TLSVersion = "1.1"
	TLSVersionOne2 TLSVersion = "1.2"
)

// PossibleTLSVersionValues returns the possible values for the TLSVersion const type.
func PossibleTLSVersionValues() []TLSVersion {
	return []TLSVersion{
		TLSVersionOne0,
		TLSVersionOne1,
		TLSVersionOne2,
	}
}

// UpdateChannel - Optional: Specifies the update channel for the monthly Redis updates your Redis Cache will receive. Caches
// using 'Preview' update channel get latest Redis updates at least 4 weeks ahead of 'Stable'
// channel caches. Default value is 'Stable'.
type UpdateChannel string

const (
	UpdateChannelPreview UpdateChannel = "Preview"
	UpdateChannelStable  UpdateChannel = "Stable"
)

// PossibleUpdateChannelValues returns the possible values for the UpdateChannel const type.
func PossibleUpdateChannelValues() []UpdateChannel {
	return []UpdateChannel{
		UpdateChannelPreview,
		UpdateChannelStable,
	}
}