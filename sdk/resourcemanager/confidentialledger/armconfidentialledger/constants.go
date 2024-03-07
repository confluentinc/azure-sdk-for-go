//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfidentialledger

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confidentialledger/armconfidentialledger"
	moduleVersion = "v1.3.0-beta.1"
)

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
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

// LanguageRuntime - Object representing LanguageRuntime for Manged CCF.
type LanguageRuntime string

const (
	LanguageRuntimeCPP LanguageRuntime = "CPP"
	LanguageRuntimeJS  LanguageRuntime = "JS"
)

// PossibleLanguageRuntimeValues returns the possible values for the LanguageRuntime const type.
func PossibleLanguageRuntimeValues() []LanguageRuntime {
	return []LanguageRuntime{
		LanguageRuntimeCPP,
		LanguageRuntimeJS,
	}
}

// LedgerRoleName - LedgerRole associated with the Security Principal of Ledger
type LedgerRoleName string

const (
	LedgerRoleNameAdministrator LedgerRoleName = "Administrator"
	LedgerRoleNameContributor   LedgerRoleName = "Contributor"
	LedgerRoleNameReader        LedgerRoleName = "Reader"
)

// PossibleLedgerRoleNameValues returns the possible values for the LedgerRoleName const type.
func PossibleLedgerRoleNameValues() []LedgerRoleName {
	return []LedgerRoleName{
		LedgerRoleNameAdministrator,
		LedgerRoleNameContributor,
		LedgerRoleNameReader,
	}
}

// LedgerType - Type of the ledger. Private means transaction data is encrypted.
type LedgerType string

const (
	LedgerTypePrivate LedgerType = "Private"
	LedgerTypePublic  LedgerType = "Public"
	LedgerTypeUnknown LedgerType = "Unknown"
)

// PossibleLedgerTypeValues returns the possible values for the LedgerType const type.
func PossibleLedgerTypeValues() []LedgerType {
	return []LedgerType{
		LedgerTypePrivate,
		LedgerTypePublic,
		LedgerTypeUnknown,
	}
}

// ProvisioningState - Object representing ProvisioningState for Confidential Ledger.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
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
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
		ProvisioningStateUpdating,
	}
}

// RunningState - Object representing RunningState for Confidential Ledger.
type RunningState string

const (
	RunningStateActive   RunningState = "Active"
	RunningStatePaused   RunningState = "Paused"
	RunningStatePausing  RunningState = "Pausing"
	RunningStateResuming RunningState = "Resuming"
	RunningStateUnknown  RunningState = "Unknown"
)

// PossibleRunningStateValues returns the possible values for the RunningState const type.
func PossibleRunningStateValues() []RunningState {
	return []RunningState{
		RunningStateActive,
		RunningStatePaused,
		RunningStatePausing,
		RunningStateResuming,
		RunningStateUnknown,
	}
}