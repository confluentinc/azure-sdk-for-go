//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/providerhub/armproviderhub"
	moduleVersion = "v1.2.0"
)

type ExtensionCategory string

const (
	ExtensionCategoryNotSpecified                      ExtensionCategory = "NotSpecified"
	ExtensionCategoryResourceCreationBegin             ExtensionCategory = "ResourceCreationBegin"
	ExtensionCategoryResourceCreationCompleted         ExtensionCategory = "ResourceCreationCompleted"
	ExtensionCategoryResourceCreationValidate          ExtensionCategory = "ResourceCreationValidate"
	ExtensionCategoryResourceDeletionBegin             ExtensionCategory = "ResourceDeletionBegin"
	ExtensionCategoryResourceDeletionCompleted         ExtensionCategory = "ResourceDeletionCompleted"
	ExtensionCategoryResourceDeletionValidate          ExtensionCategory = "ResourceDeletionValidate"
	ExtensionCategoryResourceMoveBegin                 ExtensionCategory = "ResourceMoveBegin"
	ExtensionCategoryResourceMoveCompleted             ExtensionCategory = "ResourceMoveCompleted"
	ExtensionCategoryResourcePatchBegin                ExtensionCategory = "ResourcePatchBegin"
	ExtensionCategoryResourcePatchCompleted            ExtensionCategory = "ResourcePatchCompleted"
	ExtensionCategoryResourcePatchValidate             ExtensionCategory = "ResourcePatchValidate"
	ExtensionCategoryResourcePostAction                ExtensionCategory = "ResourcePostAction"
	ExtensionCategoryResourceReadBegin                 ExtensionCategory = "ResourceReadBegin"
	ExtensionCategoryResourceReadValidate              ExtensionCategory = "ResourceReadValidate"
	ExtensionCategorySubscriptionLifecycleNotification ExtensionCategory = "SubscriptionLifecycleNotification"
)

// PossibleExtensionCategoryValues returns the possible values for the ExtensionCategory const type.
func PossibleExtensionCategoryValues() []ExtensionCategory {
	return []ExtensionCategory{
		ExtensionCategoryNotSpecified,
		ExtensionCategoryResourceCreationBegin,
		ExtensionCategoryResourceCreationCompleted,
		ExtensionCategoryResourceCreationValidate,
		ExtensionCategoryResourceDeletionBegin,
		ExtensionCategoryResourceDeletionCompleted,
		ExtensionCategoryResourceDeletionValidate,
		ExtensionCategoryResourceMoveBegin,
		ExtensionCategoryResourceMoveCompleted,
		ExtensionCategoryResourcePatchBegin,
		ExtensionCategoryResourcePatchCompleted,
		ExtensionCategoryResourcePatchValidate,
		ExtensionCategoryResourcePostAction,
		ExtensionCategoryResourceReadBegin,
		ExtensionCategoryResourceReadValidate,
		ExtensionCategorySubscriptionLifecycleNotification,
	}
}

type ExtensionOptionType string

const (
	ExtensionOptionTypeDoNotMergeExistingReadOnlyAndSecretProperties ExtensionOptionType = "DoNotMergeExistingReadOnlyAndSecretProperties"
	ExtensionOptionTypeIncludeInternalMetadata                       ExtensionOptionType = "IncludeInternalMetadata"
	ExtensionOptionTypeNotSpecified                                  ExtensionOptionType = "NotSpecified"
)

// PossibleExtensionOptionTypeValues returns the possible values for the ExtensionOptionType const type.
func PossibleExtensionOptionTypeValues() []ExtensionOptionType {
	return []ExtensionOptionType{
		ExtensionOptionTypeDoNotMergeExistingReadOnlyAndSecretProperties,
		ExtensionOptionTypeIncludeInternalMetadata,
		ExtensionOptionTypeNotSpecified,
	}
}

type FeaturesPolicy string

const (
	FeaturesPolicyAll FeaturesPolicy = "All"
	FeaturesPolicyAny FeaturesPolicy = "Any"
)

// PossibleFeaturesPolicyValues returns the possible values for the FeaturesPolicy const type.
func PossibleFeaturesPolicyValues() []FeaturesPolicy {
	return []FeaturesPolicy{
		FeaturesPolicyAll,
		FeaturesPolicyAny,
	}
}

type IdentityManagementTypes string

const (
	IdentityManagementTypesActor                     IdentityManagementTypes = "Actor"
	IdentityManagementTypesDelegatedResourceIdentity IdentityManagementTypes = "DelegatedResourceIdentity"
	IdentityManagementTypesNotSpecified              IdentityManagementTypes = "NotSpecified"
	IdentityManagementTypesSystemAssigned            IdentityManagementTypes = "SystemAssigned"
	IdentityManagementTypesUserAssigned              IdentityManagementTypes = "UserAssigned"
)

// PossibleIdentityManagementTypesValues returns the possible values for the IdentityManagementTypes const type.
func PossibleIdentityManagementTypesValues() []IdentityManagementTypes {
	return []IdentityManagementTypes{
		IdentityManagementTypesActor,
		IdentityManagementTypesDelegatedResourceIdentity,
		IdentityManagementTypesNotSpecified,
		IdentityManagementTypesSystemAssigned,
		IdentityManagementTypesUserAssigned,
	}
}

type LinkedAction string

const (
	LinkedActionBlocked      LinkedAction = "Blocked"
	LinkedActionEnabled      LinkedAction = "Enabled"
	LinkedActionNotSpecified LinkedAction = "NotSpecified"
	LinkedActionValidate     LinkedAction = "Validate"
)

// PossibleLinkedActionValues returns the possible values for the LinkedAction const type.
func PossibleLinkedActionValues() []LinkedAction {
	return []LinkedAction{
		LinkedActionBlocked,
		LinkedActionEnabled,
		LinkedActionNotSpecified,
		LinkedActionValidate,
	}
}

type LinkedOperation string

const (
	LinkedOperationCrossResourceGroupResourceMove LinkedOperation = "CrossResourceGroupResourceMove"
	LinkedOperationCrossSubscriptionResourceMove  LinkedOperation = "CrossSubscriptionResourceMove"
	LinkedOperationNone                           LinkedOperation = "None"
)

// PossibleLinkedOperationValues returns the possible values for the LinkedOperation const type.
func PossibleLinkedOperationValues() []LinkedOperation {
	return []LinkedOperation{
		LinkedOperationCrossResourceGroupResourceMove,
		LinkedOperationCrossSubscriptionResourceMove,
		LinkedOperationNone,
	}
}

type LoggingDetails string

const (
	LoggingDetailsBody LoggingDetails = "Body"
	LoggingDetailsNone LoggingDetails = "None"
)

// PossibleLoggingDetailsValues returns the possible values for the LoggingDetails const type.
func PossibleLoggingDetailsValues() []LoggingDetails {
	return []LoggingDetails{
		LoggingDetailsBody,
		LoggingDetailsNone,
	}
}

type LoggingDirections string

const (
	LoggingDirectionsNone     LoggingDirections = "None"
	LoggingDirectionsRequest  LoggingDirections = "Request"
	LoggingDirectionsResponse LoggingDirections = "Response"
)

// PossibleLoggingDirectionsValues returns the possible values for the LoggingDirections const type.
func PossibleLoggingDirectionsValues() []LoggingDirections {
	return []LoggingDirections{
		LoggingDirectionsNone,
		LoggingDirectionsRequest,
		LoggingDirectionsResponse,
	}
}

type ManifestResourceDeletionPolicy string

const (
	ManifestResourceDeletionPolicyCascade      ManifestResourceDeletionPolicy = "Cascade"
	ManifestResourceDeletionPolicyForce        ManifestResourceDeletionPolicy = "Force"
	ManifestResourceDeletionPolicyNotSpecified ManifestResourceDeletionPolicy = "NotSpecified"
)

// PossibleManifestResourceDeletionPolicyValues returns the possible values for the ManifestResourceDeletionPolicy const type.
func PossibleManifestResourceDeletionPolicyValues() []ManifestResourceDeletionPolicy {
	return []ManifestResourceDeletionPolicy{
		ManifestResourceDeletionPolicyCascade,
		ManifestResourceDeletionPolicyForce,
		ManifestResourceDeletionPolicyNotSpecified,
	}
}

type MessageScope string

const (
	MessageScopeNotSpecified            MessageScope = "NotSpecified"
	MessageScopeRegisteredSubscriptions MessageScope = "RegisteredSubscriptions"
)

// PossibleMessageScopeValues returns the possible values for the MessageScope const type.
func PossibleMessageScopeValues() []MessageScope {
	return []MessageScope{
		MessageScopeNotSpecified,
		MessageScopeRegisteredSubscriptions,
	}
}

type NotificationMode string

const (
	NotificationModeEventHub     NotificationMode = "EventHub"
	NotificationModeNotSpecified NotificationMode = "NotSpecified"
	NotificationModeWebHook      NotificationMode = "WebHook"
)

// PossibleNotificationModeValues returns the possible values for the NotificationMode const type.
func PossibleNotificationModeValues() []NotificationMode {
	return []NotificationMode{
		NotificationModeEventHub,
		NotificationModeNotSpecified,
		NotificationModeWebHook,
	}
}

type OperationsDefinitionActionType string

const (
	OperationsDefinitionActionTypeInternal     OperationsDefinitionActionType = "Internal"
	OperationsDefinitionActionTypeNotSpecified OperationsDefinitionActionType = "NotSpecified"
)

// PossibleOperationsDefinitionActionTypeValues returns the possible values for the OperationsDefinitionActionType const type.
func PossibleOperationsDefinitionActionTypeValues() []OperationsDefinitionActionType {
	return []OperationsDefinitionActionType{
		OperationsDefinitionActionTypeInternal,
		OperationsDefinitionActionTypeNotSpecified,
	}
}

type OperationsDefinitionOrigin string

const (
	OperationsDefinitionOriginNotSpecified OperationsDefinitionOrigin = "NotSpecified"
	OperationsDefinitionOriginSystem       OperationsDefinitionOrigin = "System"
	OperationsDefinitionOriginUser         OperationsDefinitionOrigin = "User"
)

// PossibleOperationsDefinitionOriginValues returns the possible values for the OperationsDefinitionOrigin const type.
func PossibleOperationsDefinitionOriginValues() []OperationsDefinitionOrigin {
	return []OperationsDefinitionOrigin{
		OperationsDefinitionOriginNotSpecified,
		OperationsDefinitionOriginSystem,
		OperationsDefinitionOriginUser,
	}
}

type OptInHeaderType string

const (
	OptInHeaderTypeClientGroupMembership          OptInHeaderType = "ClientGroupMembership"
	OptInHeaderTypeNotSpecified                   OptInHeaderType = "NotSpecified"
	OptInHeaderTypeSignedAuxiliaryTokens          OptInHeaderType = "SignedAuxiliaryTokens"
	OptInHeaderTypeSignedUserToken                OptInHeaderType = "SignedUserToken"
	OptInHeaderTypeUnboundedClientGroupMembership OptInHeaderType = "UnboundedClientGroupMembership"
)

// PossibleOptInHeaderTypeValues returns the possible values for the OptInHeaderType const type.
func PossibleOptInHeaderTypeValues() []OptInHeaderType {
	return []OptInHeaderType{
		OptInHeaderTypeClientGroupMembership,
		OptInHeaderTypeNotSpecified,
		OptInHeaderTypeSignedAuxiliaryTokens,
		OptInHeaderTypeSignedUserToken,
		OptInHeaderTypeUnboundedClientGroupMembership,
	}
}

type PreflightOption string

const (
	PreflightOptionContinueDeploymentOnFailure PreflightOption = "ContinueDeploymentOnFailure"
	PreflightOptionDefaultValidationOnly       PreflightOption = "DefaultValidationOnly"
	PreflightOptionNone                        PreflightOption = "None"
)

// PossiblePreflightOptionValues returns the possible values for the PreflightOption const type.
func PossiblePreflightOptionValues() []PreflightOption {
	return []PreflightOption{
		PreflightOptionContinueDeploymentOnFailure,
		PreflightOptionDefaultValidationOnly,
		PreflightOptionNone,
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted          ProvisioningState = "Accepted"
	ProvisioningStateCanceled          ProvisioningState = "Canceled"
	ProvisioningStateCreated           ProvisioningState = "Created"
	ProvisioningStateCreating          ProvisioningState = "Creating"
	ProvisioningStateDeleted           ProvisioningState = "Deleted"
	ProvisioningStateDeleting          ProvisioningState = "Deleting"
	ProvisioningStateFailed            ProvisioningState = "Failed"
	ProvisioningStateMovingResources   ProvisioningState = "MovingResources"
	ProvisioningStateNotSpecified      ProvisioningState = "NotSpecified"
	ProvisioningStateRolloutInProgress ProvisioningState = "RolloutInProgress"
	ProvisioningStateRunning           ProvisioningState = "Running"
	ProvisioningStateSucceeded         ProvisioningState = "Succeeded"
	ProvisioningStateTransientFailure  ProvisioningState = "TransientFailure"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateDeleted,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMovingResources,
		ProvisioningStateNotSpecified,
		ProvisioningStateRolloutInProgress,
		ProvisioningStateRunning,
		ProvisioningStateSucceeded,
		ProvisioningStateTransientFailure,
	}
}

type Regionality string

const (
	RegionalityGlobal       Regionality = "Global"
	RegionalityNotSpecified Regionality = "NotSpecified"
	RegionalityRegional     Regionality = "Regional"
)

// PossibleRegionalityValues returns the possible values for the Regionality const type.
func PossibleRegionalityValues() []Regionality {
	return []Regionality{
		RegionalityGlobal,
		RegionalityNotSpecified,
		RegionalityRegional,
	}
}

type ResourceDeletionPolicy string

const (
	ResourceDeletionPolicyCascadeDeleteAll               ResourceDeletionPolicy = "CascadeDeleteAll"
	ResourceDeletionPolicyCascadeDeleteProxyOnlyChildren ResourceDeletionPolicy = "CascadeDeleteProxyOnlyChildren"
	ResourceDeletionPolicyNotSpecified                   ResourceDeletionPolicy = "NotSpecified"
)

// PossibleResourceDeletionPolicyValues returns the possible values for the ResourceDeletionPolicy const type.
func PossibleResourceDeletionPolicyValues() []ResourceDeletionPolicy {
	return []ResourceDeletionPolicy{
		ResourceDeletionPolicyCascadeDeleteAll,
		ResourceDeletionPolicyCascadeDeleteProxyOnlyChildren,
		ResourceDeletionPolicyNotSpecified,
	}
}

type ResourceProviderCapabilitiesEffect string

const (
	ResourceProviderCapabilitiesEffectAllow        ResourceProviderCapabilitiesEffect = "Allow"
	ResourceProviderCapabilitiesEffectDisallow     ResourceProviderCapabilitiesEffect = "Disallow"
	ResourceProviderCapabilitiesEffectNotSpecified ResourceProviderCapabilitiesEffect = "NotSpecified"
)

// PossibleResourceProviderCapabilitiesEffectValues returns the possible values for the ResourceProviderCapabilitiesEffect const type.
func PossibleResourceProviderCapabilitiesEffectValues() []ResourceProviderCapabilitiesEffect {
	return []ResourceProviderCapabilitiesEffect{
		ResourceProviderCapabilitiesEffectAllow,
		ResourceProviderCapabilitiesEffectDisallow,
		ResourceProviderCapabilitiesEffectNotSpecified,
	}
}

type ResourceProviderManagementResourceAccessPolicy string

const (
	ResourceProviderManagementResourceAccessPolicyAcisActionAllowed ResourceProviderManagementResourceAccessPolicy = "AcisActionAllowed"
	ResourceProviderManagementResourceAccessPolicyAcisReadAllowed   ResourceProviderManagementResourceAccessPolicy = "AcisReadAllowed"
	ResourceProviderManagementResourceAccessPolicyNotSpecified      ResourceProviderManagementResourceAccessPolicy = "NotSpecified"
)

// PossibleResourceProviderManagementResourceAccessPolicyValues returns the possible values for the ResourceProviderManagementResourceAccessPolicy const type.
func PossibleResourceProviderManagementResourceAccessPolicyValues() []ResourceProviderManagementResourceAccessPolicy {
	return []ResourceProviderManagementResourceAccessPolicy{
		ResourceProviderManagementResourceAccessPolicyAcisActionAllowed,
		ResourceProviderManagementResourceAccessPolicyAcisReadAllowed,
		ResourceProviderManagementResourceAccessPolicyNotSpecified,
	}
}

type ResourceProviderType string

const (
	ResourceProviderTypeAuthorizationFree          ResourceProviderType = "AuthorizationFree"
	ResourceProviderTypeExternal                   ResourceProviderType = "External"
	ResourceProviderTypeHidden                     ResourceProviderType = "Hidden"
	ResourceProviderTypeInternal                   ResourceProviderType = "Internal"
	ResourceProviderTypeLegacyRegistrationRequired ResourceProviderType = "LegacyRegistrationRequired"
	ResourceProviderTypeNotSpecified               ResourceProviderType = "NotSpecified"
	ResourceProviderTypeRegistrationFree           ResourceProviderType = "RegistrationFree"
	ResourceProviderTypeTenantOnly                 ResourceProviderType = "TenantOnly"
)

// PossibleResourceProviderTypeValues returns the possible values for the ResourceProviderType const type.
func PossibleResourceProviderTypeValues() []ResourceProviderType {
	return []ResourceProviderType{
		ResourceProviderTypeAuthorizationFree,
		ResourceProviderTypeExternal,
		ResourceProviderTypeHidden,
		ResourceProviderTypeInternal,
		ResourceProviderTypeLegacyRegistrationRequired,
		ResourceProviderTypeNotSpecified,
		ResourceProviderTypeRegistrationFree,
		ResourceProviderTypeTenantOnly,
	}
}

type ResourceTypeMarketplaceType string

const (
	ResourceTypeMarketplaceTypeAddOn        ResourceTypeMarketplaceType = "AddOn"
	ResourceTypeMarketplaceTypeBypass       ResourceTypeMarketplaceType = "Bypass"
	ResourceTypeMarketplaceTypeNotSpecified ResourceTypeMarketplaceType = "NotSpecified"
	ResourceTypeMarketplaceTypeStore        ResourceTypeMarketplaceType = "Store"
)

// PossibleResourceTypeMarketplaceTypeValues returns the possible values for the ResourceTypeMarketplaceType const type.
func PossibleResourceTypeMarketplaceTypeValues() []ResourceTypeMarketplaceType {
	return []ResourceTypeMarketplaceType{
		ResourceTypeMarketplaceTypeAddOn,
		ResourceTypeMarketplaceTypeBypass,
		ResourceTypeMarketplaceTypeNotSpecified,
		ResourceTypeMarketplaceTypeStore,
	}
}

type ResourceTypeRegistrationPropertiesMarketplaceType string

const (
	ResourceTypeRegistrationPropertiesMarketplaceTypeAddOn        ResourceTypeRegistrationPropertiesMarketplaceType = "AddOn"
	ResourceTypeRegistrationPropertiesMarketplaceTypeBypass       ResourceTypeRegistrationPropertiesMarketplaceType = "Bypass"
	ResourceTypeRegistrationPropertiesMarketplaceTypeNotSpecified ResourceTypeRegistrationPropertiesMarketplaceType = "NotSpecified"
	ResourceTypeRegistrationPropertiesMarketplaceTypeStore        ResourceTypeRegistrationPropertiesMarketplaceType = "Store"
)

// PossibleResourceTypeRegistrationPropertiesMarketplaceTypeValues returns the possible values for the ResourceTypeRegistrationPropertiesMarketplaceType const type.
func PossibleResourceTypeRegistrationPropertiesMarketplaceTypeValues() []ResourceTypeRegistrationPropertiesMarketplaceType {
	return []ResourceTypeRegistrationPropertiesMarketplaceType{
		ResourceTypeRegistrationPropertiesMarketplaceTypeAddOn,
		ResourceTypeRegistrationPropertiesMarketplaceTypeBypass,
		ResourceTypeRegistrationPropertiesMarketplaceTypeNotSpecified,
		ResourceTypeRegistrationPropertiesMarketplaceTypeStore,
	}
}

type ResourceValidation string

const (
	ResourceValidationNotSpecified  ResourceValidation = "NotSpecified"
	ResourceValidationProfaneWords  ResourceValidation = "ProfaneWords"
	ResourceValidationReservedWords ResourceValidation = "ReservedWords"
)

// PossibleResourceValidationValues returns the possible values for the ResourceValidation const type.
func PossibleResourceValidationValues() []ResourceValidation {
	return []ResourceValidation{
		ResourceValidationNotSpecified,
		ResourceValidationProfaneWords,
		ResourceValidationReservedWords,
	}
}

type RoutingType string

const (
	RoutingTypeCascadeExtension RoutingType = "CascadeExtension"
	RoutingTypeDefault          RoutingType = "Default"
	RoutingTypeExtension        RoutingType = "Extension"
	RoutingTypeFailover         RoutingType = "Failover"
	RoutingTypeFanout           RoutingType = "Fanout"
	RoutingTypeHostBased        RoutingType = "HostBased"
	RoutingTypeLocationBased    RoutingType = "LocationBased"
	RoutingTypeProxyOnly        RoutingType = "ProxyOnly"
	RoutingTypeTenant           RoutingType = "Tenant"
)

// PossibleRoutingTypeValues returns the possible values for the RoutingType const type.
func PossibleRoutingTypeValues() []RoutingType {
	return []RoutingType{
		RoutingTypeCascadeExtension,
		RoutingTypeDefault,
		RoutingTypeExtension,
		RoutingTypeFailover,
		RoutingTypeFanout,
		RoutingTypeHostBased,
		RoutingTypeLocationBased,
		RoutingTypeProxyOnly,
		RoutingTypeTenant,
	}
}

type SKULocationInfoType string

const (
	SKULocationInfoTypeArcZone      SKULocationInfoType = "ArcZone"
	SKULocationInfoTypeEdgeZone     SKULocationInfoType = "EdgeZone"
	SKULocationInfoTypeNotSpecified SKULocationInfoType = "NotSpecified"
)

// PossibleSKULocationInfoTypeValues returns the possible values for the SKULocationInfoType const type.
func PossibleSKULocationInfoTypeValues() []SKULocationInfoType {
	return []SKULocationInfoType{
		SKULocationInfoTypeArcZone,
		SKULocationInfoTypeEdgeZone,
		SKULocationInfoTypeNotSpecified,
	}
}

type SKUScaleType string

const (
	SKUScaleTypeAutomatic SKUScaleType = "Automatic"
	SKUScaleTypeManual    SKUScaleType = "Manual"
	SKUScaleTypeNone      SKUScaleType = "None"
)

// PossibleSKUScaleTypeValues returns the possible values for the SKUScaleType const type.
func PossibleSKUScaleTypeValues() []SKUScaleType {
	return []SKUScaleType{
		SKUScaleTypeAutomatic,
		SKUScaleTypeManual,
		SKUScaleTypeNone,
	}
}

type SubscriptionNotificationOperation string

const (
	SubscriptionNotificationOperationBillingCancellation    SubscriptionNotificationOperation = "BillingCancellation"
	SubscriptionNotificationOperationDeleteAllResources     SubscriptionNotificationOperation = "DeleteAllResources"
	SubscriptionNotificationOperationNoOp                   SubscriptionNotificationOperation = "NoOp"
	SubscriptionNotificationOperationNotDefined             SubscriptionNotificationOperation = "NotDefined"
	SubscriptionNotificationOperationSoftDeleteAllResources SubscriptionNotificationOperation = "SoftDeleteAllResources"
	SubscriptionNotificationOperationUndoSoftDelete         SubscriptionNotificationOperation = "UndoSoftDelete"
)

// PossibleSubscriptionNotificationOperationValues returns the possible values for the SubscriptionNotificationOperation const type.
func PossibleSubscriptionNotificationOperationValues() []SubscriptionNotificationOperation {
	return []SubscriptionNotificationOperation{
		SubscriptionNotificationOperationBillingCancellation,
		SubscriptionNotificationOperationDeleteAllResources,
		SubscriptionNotificationOperationNoOp,
		SubscriptionNotificationOperationNotDefined,
		SubscriptionNotificationOperationSoftDeleteAllResources,
		SubscriptionNotificationOperationUndoSoftDelete,
	}
}

type SubscriptionReregistrationResult string

const (
	SubscriptionReregistrationResultConditionalUpdate SubscriptionReregistrationResult = "ConditionalUpdate"
	SubscriptionReregistrationResultFailed            SubscriptionReregistrationResult = "Failed"
	SubscriptionReregistrationResultForcedUpdate      SubscriptionReregistrationResult = "ForcedUpdate"
	SubscriptionReregistrationResultNotApplicable     SubscriptionReregistrationResult = "NotApplicable"
)

// PossibleSubscriptionReregistrationResultValues returns the possible values for the SubscriptionReregistrationResult const type.
func PossibleSubscriptionReregistrationResultValues() []SubscriptionReregistrationResult {
	return []SubscriptionReregistrationResult{
		SubscriptionReregistrationResultConditionalUpdate,
		SubscriptionReregistrationResultFailed,
		SubscriptionReregistrationResultForcedUpdate,
		SubscriptionReregistrationResultNotApplicable,
	}
}

type SubscriptionState string

const (
	SubscriptionStateDeleted    SubscriptionState = "Deleted"
	SubscriptionStateDisabled   SubscriptionState = "Disabled"
	SubscriptionStateEnabled    SubscriptionState = "Enabled"
	SubscriptionStateNotDefined SubscriptionState = "NotDefined"
	SubscriptionStatePastDue    SubscriptionState = "PastDue"
	SubscriptionStateWarned     SubscriptionState = "Warned"
)

// PossibleSubscriptionStateValues returns the possible values for the SubscriptionState const type.
func PossibleSubscriptionStateValues() []SubscriptionState {
	return []SubscriptionState{
		SubscriptionStateDeleted,
		SubscriptionStateDisabled,
		SubscriptionStateEnabled,
		SubscriptionStateNotDefined,
		SubscriptionStatePastDue,
		SubscriptionStateWarned,
	}
}

type SubscriptionTransitioningState string

const (
	SubscriptionTransitioningStateDeleted                 SubscriptionTransitioningState = "Deleted"
	SubscriptionTransitioningStateRegistered              SubscriptionTransitioningState = "Registered"
	SubscriptionTransitioningStateSuspended               SubscriptionTransitioningState = "Suspended"
	SubscriptionTransitioningStateSuspendedToDeleted      SubscriptionTransitioningState = "SuspendedToDeleted"
	SubscriptionTransitioningStateSuspendedToRegistered   SubscriptionTransitioningState = "SuspendedToRegistered"
	SubscriptionTransitioningStateSuspendedToUnregistered SubscriptionTransitioningState = "SuspendedToUnregistered"
	SubscriptionTransitioningStateSuspendedToWarned       SubscriptionTransitioningState = "SuspendedToWarned"
	SubscriptionTransitioningStateUnregistered            SubscriptionTransitioningState = "Unregistered"
	SubscriptionTransitioningStateWarned                  SubscriptionTransitioningState = "Warned"
	SubscriptionTransitioningStateWarnedToDeleted         SubscriptionTransitioningState = "WarnedToDeleted"
	SubscriptionTransitioningStateWarnedToRegistered      SubscriptionTransitioningState = "WarnedToRegistered"
	SubscriptionTransitioningStateWarnedToSuspended       SubscriptionTransitioningState = "WarnedToSuspended"
	SubscriptionTransitioningStateWarnedToUnregistered    SubscriptionTransitioningState = "WarnedToUnregistered"
)

// PossibleSubscriptionTransitioningStateValues returns the possible values for the SubscriptionTransitioningState const type.
func PossibleSubscriptionTransitioningStateValues() []SubscriptionTransitioningState {
	return []SubscriptionTransitioningState{
		SubscriptionTransitioningStateDeleted,
		SubscriptionTransitioningStateRegistered,
		SubscriptionTransitioningStateSuspended,
		SubscriptionTransitioningStateSuspendedToDeleted,
		SubscriptionTransitioningStateSuspendedToRegistered,
		SubscriptionTransitioningStateSuspendedToUnregistered,
		SubscriptionTransitioningStateSuspendedToWarned,
		SubscriptionTransitioningStateUnregistered,
		SubscriptionTransitioningStateWarned,
		SubscriptionTransitioningStateWarnedToDeleted,
		SubscriptionTransitioningStateWarnedToRegistered,
		SubscriptionTransitioningStateWarnedToSuspended,
		SubscriptionTransitioningStateWarnedToUnregistered,
	}
}

type TemplateDeploymentCapabilities string

const (
	TemplateDeploymentCapabilitiesDefault   TemplateDeploymentCapabilities = "Default"
	TemplateDeploymentCapabilitiesPreflight TemplateDeploymentCapabilities = "Preflight"
)

// PossibleTemplateDeploymentCapabilitiesValues returns the possible values for the TemplateDeploymentCapabilities const type.
func PossibleTemplateDeploymentCapabilitiesValues() []TemplateDeploymentCapabilities {
	return []TemplateDeploymentCapabilities{
		TemplateDeploymentCapabilitiesDefault,
		TemplateDeploymentCapabilitiesPreflight,
	}
}

type TemplateDeploymentPreflightOptions string

const (
	TemplateDeploymentPreflightOptionsDeploymentRequests TemplateDeploymentPreflightOptions = "DeploymentRequests"
	TemplateDeploymentPreflightOptionsNone               TemplateDeploymentPreflightOptions = "None"
	TemplateDeploymentPreflightOptionsRegisteredOnly     TemplateDeploymentPreflightOptions = "RegisteredOnly"
	TemplateDeploymentPreflightOptionsTestOnly           TemplateDeploymentPreflightOptions = "TestOnly"
	TemplateDeploymentPreflightOptionsValidationRequests TemplateDeploymentPreflightOptions = "ValidationRequests"
)

// PossibleTemplateDeploymentPreflightOptionsValues returns the possible values for the TemplateDeploymentPreflightOptions const type.
func PossibleTemplateDeploymentPreflightOptionsValues() []TemplateDeploymentPreflightOptions {
	return []TemplateDeploymentPreflightOptions{
		TemplateDeploymentPreflightOptionsDeploymentRequests,
		TemplateDeploymentPreflightOptionsNone,
		TemplateDeploymentPreflightOptionsRegisteredOnly,
		TemplateDeploymentPreflightOptionsTestOnly,
		TemplateDeploymentPreflightOptionsValidationRequests,
	}
}

type ThrottlingMetricType string

const (
	ThrottlingMetricTypeNotSpecified      ThrottlingMetricType = "NotSpecified"
	ThrottlingMetricTypeNumberOfRequests  ThrottlingMetricType = "NumberOfRequests"
	ThrottlingMetricTypeNumberOfResources ThrottlingMetricType = "NumberOfResources"
)

// PossibleThrottlingMetricTypeValues returns the possible values for the ThrottlingMetricType const type.
func PossibleThrottlingMetricTypeValues() []ThrottlingMetricType {
	return []ThrottlingMetricType{
		ThrottlingMetricTypeNotSpecified,
		ThrottlingMetricTypeNumberOfRequests,
		ThrottlingMetricTypeNumberOfResources,
	}
}

type TrafficRegionCategory string

const (
	TrafficRegionCategoryCanary                 TrafficRegionCategory = "Canary"
	TrafficRegionCategoryHighTraffic            TrafficRegionCategory = "HighTraffic"
	TrafficRegionCategoryLowTraffic             TrafficRegionCategory = "LowTraffic"
	TrafficRegionCategoryMediumTraffic          TrafficRegionCategory = "MediumTraffic"
	TrafficRegionCategoryNone                   TrafficRegionCategory = "None"
	TrafficRegionCategoryNotSpecified           TrafficRegionCategory = "NotSpecified"
	TrafficRegionCategoryRestOfTheWorldGroupOne TrafficRegionCategory = "RestOfTheWorldGroupOne"
	TrafficRegionCategoryRestOfTheWorldGroupTwo TrafficRegionCategory = "RestOfTheWorldGroupTwo"
)

// PossibleTrafficRegionCategoryValues returns the possible values for the TrafficRegionCategory const type.
func PossibleTrafficRegionCategoryValues() []TrafficRegionCategory {
	return []TrafficRegionCategory{
		TrafficRegionCategoryCanary,
		TrafficRegionCategoryHighTraffic,
		TrafficRegionCategoryLowTraffic,
		TrafficRegionCategoryMediumTraffic,
		TrafficRegionCategoryNone,
		TrafficRegionCategoryNotSpecified,
		TrafficRegionCategoryRestOfTheWorldGroupOne,
		TrafficRegionCategoryRestOfTheWorldGroupTwo,
	}
}