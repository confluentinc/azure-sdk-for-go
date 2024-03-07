//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
	moduleVersion = "v3.1.0"
)

// AuthenticationType - How to authenticate users who access local diagnostics APIs.
type AuthenticationType string

const (
	// AuthenticationTypeAAD - Use AAD SSO to authenticate the user (this requires internet access).
	AuthenticationTypeAAD AuthenticationType = "AAD"
	// AuthenticationTypePassword - Use locally stored passwords to authenticate the user.
	AuthenticationTypePassword AuthenticationType = "Password"
)

// PossibleAuthenticationTypeValues returns the possible values for the AuthenticationType const type.
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return []AuthenticationType{
		AuthenticationTypeAAD,
		AuthenticationTypePassword,
	}
}

// BillingSKU - The SKU of the packet core control plane resource. The SKU list may change over time when a new SKU gets added
// or an exiting SKU gets removed.
type BillingSKU string

const (
	// BillingSKUG0 - 100 Mbps, 20 active SIMs plan, 2 RANs
	BillingSKUG0 BillingSKU = "G0"
	// BillingSKUG1 - 1 Gbps, 100 active SIMs plan, 5 RANs
	BillingSKUG1 BillingSKU = "G1"
	// BillingSKUG10 - 10 Gbps, 1000 active SIMs plan
	BillingSKUG10 BillingSKU = "G10"
	// BillingSKUG2 - 2 Gbps, 200 active SIMs plan, 10 RANs
	BillingSKUG2 BillingSKU = "G2"
	// BillingSKUG5 - 5 Gbps, 500 active SIMs plan
	BillingSKUG5 BillingSKU = "G5"
)

// PossibleBillingSKUValues returns the possible values for the BillingSKU const type.
func PossibleBillingSKUValues() []BillingSKU {
	return []BillingSKU{
		BillingSKUG0,
		BillingSKUG1,
		BillingSKUG10,
		BillingSKUG2,
		BillingSKUG5,
	}
}

// CertificateProvisioningState - The certificate's provisioning state
type CertificateProvisioningState string

const (
	// CertificateProvisioningStateFailed - The certificate failed to be provisioned. The "reason" property explains why.
	CertificateProvisioningStateFailed CertificateProvisioningState = "Failed"
	// CertificateProvisioningStateNotProvisioned - The certificate has not been provisioned.
	CertificateProvisioningStateNotProvisioned CertificateProvisioningState = "NotProvisioned"
	// CertificateProvisioningStateProvisioned - The certificate has been provisioned.
	CertificateProvisioningStateProvisioned CertificateProvisioningState = "Provisioned"
)

// PossibleCertificateProvisioningStateValues returns the possible values for the CertificateProvisioningState const type.
func PossibleCertificateProvisioningStateValues() []CertificateProvisioningState {
	return []CertificateProvisioningState{
		CertificateProvisioningStateFailed,
		CertificateProvisioningStateNotProvisioned,
		CertificateProvisioningStateProvisioned,
	}
}

// CoreNetworkType - The core network technology generation (5G core, EPC / 4G core or EPC / 4G + 5G core).
type CoreNetworkType string

const (
	// CoreNetworkTypeEPC - EPC / 4G core
	CoreNetworkTypeEPC CoreNetworkType = "EPC"
	// CoreNetworkTypeEPC5GC - Combined EPC / 4G and 5G core
	CoreNetworkTypeEPC5GC CoreNetworkType = "EPC + 5GC"
	// CoreNetworkTypeFiveGC - 5G core
	CoreNetworkTypeFiveGC CoreNetworkType = "5GC"
)

// PossibleCoreNetworkTypeValues returns the possible values for the CoreNetworkType const type.
func PossibleCoreNetworkTypeValues() []CoreNetworkType {
	return []CoreNetworkType{
		CoreNetworkTypeEPC,
		CoreNetworkTypeEPC5GC,
		CoreNetworkTypeFiveGC,
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

// DesiredInstallationState - The desired installation state of the packet core.
type DesiredInstallationState string

const (
	// DesiredInstallationStateInstalled - Install the packet core.
	DesiredInstallationStateInstalled DesiredInstallationState = "Installed"
	// DesiredInstallationStateUninstalled - Don't install the packet core.
	DesiredInstallationStateUninstalled DesiredInstallationState = "Uninstalled"
)

// PossibleDesiredInstallationStateValues returns the possible values for the DesiredInstallationState const type.
func PossibleDesiredInstallationStateValues() []DesiredInstallationState {
	return []DesiredInstallationState{
		DesiredInstallationStateInstalled,
		DesiredInstallationStateUninstalled,
	}
}

// DiagnosticsPackageStatus - The status of the diagnostics package collection.
type DiagnosticsPackageStatus string

const (
	DiagnosticsPackageStatusCollected  DiagnosticsPackageStatus = "Collected"
	DiagnosticsPackageStatusCollecting DiagnosticsPackageStatus = "Collecting"
	DiagnosticsPackageStatusError      DiagnosticsPackageStatus = "Error"
	DiagnosticsPackageStatusNotStarted DiagnosticsPackageStatus = "NotStarted"
)

// PossibleDiagnosticsPackageStatusValues returns the possible values for the DiagnosticsPackageStatus const type.
func PossibleDiagnosticsPackageStatusValues() []DiagnosticsPackageStatus {
	return []DiagnosticsPackageStatus{
		DiagnosticsPackageStatusCollected,
		DiagnosticsPackageStatusCollecting,
		DiagnosticsPackageStatusError,
		DiagnosticsPackageStatusNotStarted,
	}
}

// InstallationReason - The reason or list of reasons why a packet core has not been installed or requires a reinstall.
type InstallationReason string

const (
	// InstallationReasonControlPlaneAccessInterfaceHasChanged - A reinstall is required as the packet core is running with out-of-date
	// control plane access interface information.
	InstallationReasonControlPlaneAccessInterfaceHasChanged InstallationReason = "ControlPlaneAccessInterfaceHasChanged"
	// InstallationReasonControlPlaneAccessVirtualIPv4AddressesHasChanged - A reinstall is required as the packet core is running
	// with out-of-date control plane access network virtual IP address.
	InstallationReasonControlPlaneAccessVirtualIPv4AddressesHasChanged InstallationReason = "ControlPlaneAccessVirtualIpv4AddressesHasChanged"
	// InstallationReasonNoAttachedDataNetworks - The packet core has not been installed as the packet core has no attached data
	// networks.
	InstallationReasonNoAttachedDataNetworks InstallationReason = "NoAttachedDataNetworks"
	// InstallationReasonNoPacketCoreDataPlane - The packet core has not been installed as there is no configured data plane for
	// this packet core.
	InstallationReasonNoPacketCoreDataPlane InstallationReason = "NoPacketCoreDataPlane"
	// InstallationReasonNoSlices - The packet core has not been installed as the mobile network does not have any applicable
	// configured slices.
	InstallationReasonNoSlices InstallationReason = "NoSlices"
	// InstallationReasonPublicLandMobileNetworkIdentifierHasChanged - A reinstall is required as the packet core is running with
	// out-of-date PLMN ID.
	InstallationReasonPublicLandMobileNetworkIdentifierHasChanged InstallationReason = "PublicLandMobileNetworkIdentifierHasChanged"
	// InstallationReasonUserPlaneAccessInterfaceHasChanged - A reinstall is required as the packet core is running with out-of-date
	// user plane core interface.
	InstallationReasonUserPlaneAccessInterfaceHasChanged InstallationReason = "UserPlaneAccessInterfaceHasChanged"
	// InstallationReasonUserPlaneAccessVirtualIPv4AddressesHasChanged - A reinstall is required as the packet core is running
	// with out-of-date user plane access network virtual IP address.
	InstallationReasonUserPlaneAccessVirtualIPv4AddressesHasChanged InstallationReason = "UserPlaneAccessVirtualIpv4AddressesHasChanged"
	// InstallationReasonUserPlaneDataInterfaceHasChanged - A reinstall is required as the packet core is running with out-of-date
	// user plane access interface.
	InstallationReasonUserPlaneDataInterfaceHasChanged InstallationReason = "UserPlaneDataInterfaceHasChanged"
)

// PossibleInstallationReasonValues returns the possible values for the InstallationReason const type.
func PossibleInstallationReasonValues() []InstallationReason {
	return []InstallationReason{
		InstallationReasonControlPlaneAccessInterfaceHasChanged,
		InstallationReasonControlPlaneAccessVirtualIPv4AddressesHasChanged,
		InstallationReasonNoAttachedDataNetworks,
		InstallationReasonNoPacketCoreDataPlane,
		InstallationReasonNoSlices,
		InstallationReasonPublicLandMobileNetworkIdentifierHasChanged,
		InstallationReasonUserPlaneAccessInterfaceHasChanged,
		InstallationReasonUserPlaneAccessVirtualIPv4AddressesHasChanged,
		InstallationReasonUserPlaneDataInterfaceHasChanged,
	}
}

// InstallationState - The installation state of the packet core.
type InstallationState string

const (
	// InstallationStateFailed - The packet core is in failed state.
	InstallationStateFailed InstallationState = "Failed"
	// InstallationStateInstalled - The packet core is installed.
	InstallationStateInstalled InstallationState = "Installed"
	// InstallationStateInstalling - The packet core is installing.
	InstallationStateInstalling InstallationState = "Installing"
	// InstallationStateReinstalling - The packet core is reinstalling.
	InstallationStateReinstalling InstallationState = "Reinstalling"
	// InstallationStateRollingBack - The packet core is rolling back to its previous version.
	InstallationStateRollingBack InstallationState = "RollingBack"
	// InstallationStateUninstalled - The packet core is uninstalled.
	InstallationStateUninstalled InstallationState = "Uninstalled"
	// InstallationStateUninstalling - The packet core is uninstalling.
	InstallationStateUninstalling InstallationState = "Uninstalling"
	// InstallationStateUpdating - The packet core is updating its configuration.
	InstallationStateUpdating InstallationState = "Updating"
	// InstallationStateUpgrading - The packet core is upgrading to a different software version.
	InstallationStateUpgrading InstallationState = "Upgrading"
)

// PossibleInstallationStateValues returns the possible values for the InstallationState const type.
func PossibleInstallationStateValues() []InstallationState {
	return []InstallationState{
		InstallationStateFailed,
		InstallationStateInstalled,
		InstallationStateInstalling,
		InstallationStateReinstalling,
		InstallationStateRollingBack,
		InstallationStateUninstalled,
		InstallationStateUninstalling,
		InstallationStateUpdating,
		InstallationStateUpgrading,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (currently only UserAssigned allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone         ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeUserAssigned,
	}
}

// NaptEnabled - Whether network address and port translation is enabled.
type NaptEnabled string

const (
	// NaptEnabledDisabled - NAPT is disabled
	NaptEnabledDisabled NaptEnabled = "Disabled"
	// NaptEnabledEnabled - NAPT is enabled
	NaptEnabledEnabled NaptEnabled = "Enabled"
)

// PossibleNaptEnabledValues returns the possible values for the NaptEnabled const type.
func PossibleNaptEnabledValues() []NaptEnabled {
	return []NaptEnabled{
		NaptEnabledDisabled,
		NaptEnabledEnabled,
	}
}

// ObsoleteVersion - Indicates whether this version is obsolete.
type ObsoleteVersion string

const (
	// ObsoleteVersionNotObsolete - This version is not obsolete for use in new packet core control plane deployments.
	ObsoleteVersionNotObsolete ObsoleteVersion = "NotObsolete"
	// ObsoleteVersionObsolete - This version is obsolete for use in new packet core control plane deployments.
	ObsoleteVersionObsolete ObsoleteVersion = "Obsolete"
)

// PossibleObsoleteVersionValues returns the possible values for the ObsoleteVersion const type.
func PossibleObsoleteVersionValues() []ObsoleteVersion {
	return []ObsoleteVersion{
		ObsoleteVersionNotObsolete,
		ObsoleteVersionObsolete,
	}
}

// PacketCaptureStatus - The status of the packet capture session.
type PacketCaptureStatus string

const (
	PacketCaptureStatusError      PacketCaptureStatus = "Error"
	PacketCaptureStatusNotStarted PacketCaptureStatus = "NotStarted"
	PacketCaptureStatusRunning    PacketCaptureStatus = "Running"
	PacketCaptureStatusStopped    PacketCaptureStatus = "Stopped"
)

// PossiblePacketCaptureStatusValues returns the possible values for the PacketCaptureStatus const type.
func PossiblePacketCaptureStatusValues() []PacketCaptureStatus {
	return []PacketCaptureStatus{
		PacketCaptureStatusError,
		PacketCaptureStatusNotStarted,
		PacketCaptureStatusRunning,
		PacketCaptureStatusStopped,
	}
}

// PduSessionType - PDU session type (IPv4/IPv6).
type PduSessionType string

const (
	PduSessionTypeIPv4 PduSessionType = "IPv4"
	PduSessionTypeIPv6 PduSessionType = "IPv6"
)

// PossiblePduSessionTypeValues returns the possible values for the PduSessionType const type.
func PossiblePduSessionTypeValues() []PduSessionType {
	return []PduSessionType{
		PduSessionTypeIPv4,
		PduSessionTypeIPv6,
	}
}

// PlatformType - The platform type where packet core is deployed. The contents of this enum can change.
type PlatformType string

const (
	// PlatformTypeAKSHCI - If this option is chosen, you must set one of "azureStackEdgeDevice", "connectedCluster" or "customLocation".
	// If multiple are set, they must be consistent with each other.
	PlatformTypeAKSHCI PlatformType = "AKS-HCI"
	// PlatformTypeThreePAZURESTACKHCI - If this option is chosen, you must set one of "azureStackHciCluster", "connectedCluster"
	// or "customLocation". If multiple are set, they must be consistent with each other.
	PlatformTypeThreePAZURESTACKHCI PlatformType = "3P-AZURE-STACK-HCI"
)

// PossiblePlatformTypeValues returns the possible values for the PlatformType const type.
func PossiblePlatformTypeValues() []PlatformType {
	return []PlatformType{
		PlatformTypeAKSHCI,
		PlatformTypeThreePAZURESTACKHCI,
	}
}

// PreemptionCapability - Preemption capability.
type PreemptionCapability string

const (
	// PreemptionCapabilityMayPreempt - May preempt
	PreemptionCapabilityMayPreempt PreemptionCapability = "MayPreempt"
	// PreemptionCapabilityNotPreempt - Cannot preempt
	PreemptionCapabilityNotPreempt PreemptionCapability = "NotPreempt"
)

// PossiblePreemptionCapabilityValues returns the possible values for the PreemptionCapability const type.
func PossiblePreemptionCapabilityValues() []PreemptionCapability {
	return []PreemptionCapability{
		PreemptionCapabilityMayPreempt,
		PreemptionCapabilityNotPreempt,
	}
}

// PreemptionVulnerability - Preemption vulnerability.
type PreemptionVulnerability string

const (
	// PreemptionVulnerabilityNotPreemptable - Cannot be preempted
	PreemptionVulnerabilityNotPreemptable PreemptionVulnerability = "NotPreemptable"
	// PreemptionVulnerabilityPreemptable - May be preempted
	PreemptionVulnerabilityPreemptable PreemptionVulnerability = "Preemptable"
)

// PossiblePreemptionVulnerabilityValues returns the possible values for the PreemptionVulnerability const type.
func PossiblePreemptionVulnerabilityValues() []PreemptionVulnerability {
	return []PreemptionVulnerability{
		PreemptionVulnerabilityNotPreemptable,
		PreemptionVulnerabilityPreemptable,
	}
}

// ProvisioningState - The current provisioning state.
type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUnknown,
	}
}

// RecommendedVersion - Indicates whether this is the recommended version to use for new packet core control plane deployments.
type RecommendedVersion string

const (
	// RecommendedVersionNotRecommended - This is not the recommended version to use for new packet core control plane deployments.
	RecommendedVersionNotRecommended RecommendedVersion = "NotRecommended"
	// RecommendedVersionRecommended - This is the recommended version to use for new packet core control plane deployments.
	RecommendedVersionRecommended RecommendedVersion = "Recommended"
)

// PossibleRecommendedVersionValues returns the possible values for the RecommendedVersion const type.
func PossibleRecommendedVersionValues() []RecommendedVersion {
	return []RecommendedVersion{
		RecommendedVersionNotRecommended,
		RecommendedVersionRecommended,
	}
}

// ReinstallRequired - Whether a reinstall of the packet core is required to pick up the latest configuration changes.
type ReinstallRequired string

const (
	// ReinstallRequiredNotRequired - A reinstall of the packet core is not required.
	ReinstallRequiredNotRequired ReinstallRequired = "NotRequired"
	// ReinstallRequiredRequired - A reinstall of the packet core is required.
	ReinstallRequiredRequired ReinstallRequired = "Required"
)

// PossibleReinstallRequiredValues returns the possible values for the ReinstallRequired const type.
func PossibleReinstallRequiredValues() []ReinstallRequired {
	return []ReinstallRequired{
		ReinstallRequiredNotRequired,
		ReinstallRequiredRequired,
	}
}

// SdfDirection - Service data flow direction.
type SdfDirection string

const (
	// SdfDirectionBidirectional - Traffic flowing both to and from the UE.
	SdfDirectionBidirectional SdfDirection = "Bidirectional"
	// SdfDirectionDownlink - Traffic flowing from the data network to the UE.
	SdfDirectionDownlink SdfDirection = "Downlink"
	// SdfDirectionUplink - Traffic flowing from the UE to the data network.
	SdfDirectionUplink SdfDirection = "Uplink"
)

// PossibleSdfDirectionValues returns the possible values for the SdfDirection const type.
func PossibleSdfDirectionValues() []SdfDirection {
	return []SdfDirection{
		SdfDirectionBidirectional,
		SdfDirectionDownlink,
		SdfDirectionUplink,
	}
}

// SimState - The state of the SIM resource.
type SimState string

const (
	// SimStateDisabled - The SIM is disabled because not all configuration required for enabling is present.
	SimStateDisabled SimState = "Disabled"
	// SimStateEnabled - The SIM is enabled.
	SimStateEnabled SimState = "Enabled"
	// SimStateInvalid - The SIM cannot be enabled because some of the associated configuration is invalid.
	SimStateInvalid SimState = "Invalid"
)

// PossibleSimStateValues returns the possible values for the SimState const type.
func PossibleSimStateValues() []SimState {
	return []SimState{
		SimStateDisabled,
		SimStateEnabled,
		SimStateInvalid,
	}
}

// SiteProvisioningState - The provisioning state of a resource e.g. SIM/SIM policy on a site.
type SiteProvisioningState string

const (
	// SiteProvisioningStateAdding - The resource is being added to this site.
	SiteProvisioningStateAdding SiteProvisioningState = "Adding"
	// SiteProvisioningStateDeleting - The resource is being deleted from this site.
	SiteProvisioningStateDeleting SiteProvisioningState = "Deleting"
	// SiteProvisioningStateFailed - The resource failed to be provisioned on this site.
	SiteProvisioningStateFailed SiteProvisioningState = "Failed"
	// SiteProvisioningStateNotApplicable - The resource should not be provisioned on this site.
	SiteProvisioningStateNotApplicable SiteProvisioningState = "NotApplicable"
	// SiteProvisioningStateProvisioned - The resource is provisioned on this site.
	SiteProvisioningStateProvisioned SiteProvisioningState = "Provisioned"
	// SiteProvisioningStateUpdating - The resource is being updated on this site.
	SiteProvisioningStateUpdating SiteProvisioningState = "Updating"
)

// PossibleSiteProvisioningStateValues returns the possible values for the SiteProvisioningState const type.
func PossibleSiteProvisioningStateValues() []SiteProvisioningState {
	return []SiteProvisioningState{
		SiteProvisioningStateAdding,
		SiteProvisioningStateDeleting,
		SiteProvisioningStateFailed,
		SiteProvisioningStateNotApplicable,
		SiteProvisioningStateProvisioned,
		SiteProvisioningStateUpdating,
	}
}

// TrafficControlPermission - Traffic control permission.
type TrafficControlPermission string

const (
	// TrafficControlPermissionBlocked - Traffic matching this rule is not allowed to flow.
	TrafficControlPermissionBlocked TrafficControlPermission = "Blocked"
	// TrafficControlPermissionEnabled - Traffic matching this rule is allowed to flow.
	TrafficControlPermissionEnabled TrafficControlPermission = "Enabled"
)

// PossibleTrafficControlPermissionValues returns the possible values for the TrafficControlPermission const type.
func PossibleTrafficControlPermissionValues() []TrafficControlPermission {
	return []TrafficControlPermission{
		TrafficControlPermissionBlocked,
		TrafficControlPermissionEnabled,
	}
}

// VersionState - The state of this packet core control plane version.
type VersionState string

const (
	// VersionStateActive - This version is active and suitable for production use.
	VersionStateActive VersionState = "Active"
	// VersionStateDeprecated - This version is deprecated and is no longer supported.
	VersionStateDeprecated VersionState = "Deprecated"
	// VersionStatePreview - This version is a preview and is not suitable for production use.
	VersionStatePreview VersionState = "Preview"
	// VersionStateUnknown - The state of this version is unknown.
	VersionStateUnknown VersionState = "Unknown"
	// VersionStateValidating - This version is currently being validated.
	VersionStateValidating VersionState = "Validating"
	// VersionStateValidationFailed - This version failed validation.
	VersionStateValidationFailed VersionState = "ValidationFailed"
)

// PossibleVersionStateValues returns the possible values for the VersionState const type.
func PossibleVersionStateValues() []VersionState {
	return []VersionState{
		VersionStateActive,
		VersionStateDeprecated,
		VersionStatePreview,
		VersionStateUnknown,
		VersionStateValidating,
		VersionStateValidationFailed,
	}
}