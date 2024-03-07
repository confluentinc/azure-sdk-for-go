//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
	moduleVersion = "v1.2.0"
)

// CalculationWindowTypes - The calculation window.
type CalculationWindowTypes string

const (
	CalculationWindowTypesDay      CalculationWindowTypes = "Day"
	CalculationWindowTypesHour     CalculationWindowTypes = "Hour"
	CalculationWindowTypesLifetime CalculationWindowTypes = "Lifetime"
	CalculationWindowTypesMonth    CalculationWindowTypes = "Month"
	CalculationWindowTypesWeek     CalculationWindowTypes = "Week"
)

// PossibleCalculationWindowTypesValues returns the possible values for the CalculationWindowTypes const type.
func PossibleCalculationWindowTypesValues() []CalculationWindowTypes {
	return []CalculationWindowTypes{
		CalculationWindowTypesDay,
		CalculationWindowTypesHour,
		CalculationWindowTypesLifetime,
		CalculationWindowTypesMonth,
		CalculationWindowTypesWeek,
	}
}

// CanonicalPropertyValueType - Type of canonical property value.
type CanonicalPropertyValueType string

const (
	CanonicalPropertyValueTypeCategorical        CanonicalPropertyValueType = "Categorical"
	CanonicalPropertyValueTypeDerivedCategorical CanonicalPropertyValueType = "DerivedCategorical"
	CanonicalPropertyValueTypeDerivedNumeric     CanonicalPropertyValueType = "DerivedNumeric"
	CanonicalPropertyValueTypeNumeric            CanonicalPropertyValueType = "Numeric"
)

// PossibleCanonicalPropertyValueTypeValues returns the possible values for the CanonicalPropertyValueType const type.
func PossibleCanonicalPropertyValueTypeValues() []CanonicalPropertyValueType {
	return []CanonicalPropertyValueType{
		CanonicalPropertyValueTypeCategorical,
		CanonicalPropertyValueTypeDerivedCategorical,
		CanonicalPropertyValueTypeDerivedNumeric,
		CanonicalPropertyValueTypeNumeric,
	}
}

// CardinalityTypes - The Relationship Cardinality.
type CardinalityTypes string

const (
	CardinalityTypesManyToMany CardinalityTypes = "ManyToMany"
	CardinalityTypesOneToMany  CardinalityTypes = "OneToMany"
	CardinalityTypesOneToOne   CardinalityTypes = "OneToOne"
)

// PossibleCardinalityTypesValues returns the possible values for the CardinalityTypes const type.
func PossibleCardinalityTypesValues() []CardinalityTypes {
	return []CardinalityTypes{
		CardinalityTypesManyToMany,
		CardinalityTypesOneToMany,
		CardinalityTypesOneToOne,
	}
}

// CompletionOperationTypes - The type of completion operation.
type CompletionOperationTypes string

const (
	CompletionOperationTypesDeleteFile CompletionOperationTypes = "DeleteFile"
	CompletionOperationTypesDoNothing  CompletionOperationTypes = "DoNothing"
	CompletionOperationTypesMoveFile   CompletionOperationTypes = "MoveFile"
)

// PossibleCompletionOperationTypesValues returns the possible values for the CompletionOperationTypes const type.
func PossibleCompletionOperationTypesValues() []CompletionOperationTypes {
	return []CompletionOperationTypes{
		CompletionOperationTypesDeleteFile,
		CompletionOperationTypesDoNothing,
		CompletionOperationTypesMoveFile,
	}
}

// ConnectorMappingStates - State of connector mapping.
type ConnectorMappingStates string

const (
	ConnectorMappingStatesCreated  ConnectorMappingStates = "Created"
	ConnectorMappingStatesCreating ConnectorMappingStates = "Creating"
	ConnectorMappingStatesExpiring ConnectorMappingStates = "Expiring"
	ConnectorMappingStatesFailed   ConnectorMappingStates = "Failed"
	ConnectorMappingStatesReady    ConnectorMappingStates = "Ready"
	ConnectorMappingStatesRunning  ConnectorMappingStates = "Running"
	ConnectorMappingStatesStopped  ConnectorMappingStates = "Stopped"
)

// PossibleConnectorMappingStatesValues returns the possible values for the ConnectorMappingStates const type.
func PossibleConnectorMappingStatesValues() []ConnectorMappingStates {
	return []ConnectorMappingStates{
		ConnectorMappingStatesCreated,
		ConnectorMappingStatesCreating,
		ConnectorMappingStatesExpiring,
		ConnectorMappingStatesFailed,
		ConnectorMappingStatesReady,
		ConnectorMappingStatesRunning,
		ConnectorMappingStatesStopped,
	}
}

// ConnectorStates - State of connector.
type ConnectorStates string

const (
	ConnectorStatesCreated  ConnectorStates = "Created"
	ConnectorStatesCreating ConnectorStates = "Creating"
	ConnectorStatesDeleting ConnectorStates = "Deleting"
	ConnectorStatesExpiring ConnectorStates = "Expiring"
	ConnectorStatesFailed   ConnectorStates = "Failed"
	ConnectorStatesReady    ConnectorStates = "Ready"
)

// PossibleConnectorStatesValues returns the possible values for the ConnectorStates const type.
func PossibleConnectorStatesValues() []ConnectorStates {
	return []ConnectorStates{
		ConnectorStatesCreated,
		ConnectorStatesCreating,
		ConnectorStatesDeleting,
		ConnectorStatesExpiring,
		ConnectorStatesFailed,
		ConnectorStatesReady,
	}
}

// ConnectorTypes - Type of connector.
type ConnectorTypes string

const (
	ConnectorTypesAzureBlob      ConnectorTypes = "AzureBlob"
	ConnectorTypesCRM            ConnectorTypes = "CRM"
	ConnectorTypesExchangeOnline ConnectorTypes = "ExchangeOnline"
	ConnectorTypesNone           ConnectorTypes = "None"
	ConnectorTypesOutbound       ConnectorTypes = "Outbound"
	ConnectorTypesSalesforce     ConnectorTypes = "Salesforce"
)

// PossibleConnectorTypesValues returns the possible values for the ConnectorTypes const type.
func PossibleConnectorTypesValues() []ConnectorTypes {
	return []ConnectorTypes{
		ConnectorTypesAzureBlob,
		ConnectorTypesCRM,
		ConnectorTypesExchangeOnline,
		ConnectorTypesNone,
		ConnectorTypesOutbound,
		ConnectorTypesSalesforce,
	}
}

// DataSourceType - The data source type.
type DataSourceType string

const (
	DataSourceTypeConnector       DataSourceType = "Connector"
	DataSourceTypeLinkInteraction DataSourceType = "LinkInteraction"
	DataSourceTypeSystemDefault   DataSourceType = "SystemDefault"
)

// PossibleDataSourceTypeValues returns the possible values for the DataSourceType const type.
func PossibleDataSourceTypeValues() []DataSourceType {
	return []DataSourceType{
		DataSourceTypeConnector,
		DataSourceTypeLinkInteraction,
		DataSourceTypeSystemDefault,
	}
}

// EntityType - Type of source entity.
type EntityType string

const (
	EntityTypeInteraction  EntityType = "Interaction"
	EntityTypeNone         EntityType = "None"
	EntityTypeProfile      EntityType = "Profile"
	EntityTypeRelationship EntityType = "Relationship"
)

// PossibleEntityTypeValues returns the possible values for the EntityType const type.
func PossibleEntityTypeValues() []EntityType {
	return []EntityType{
		EntityTypeInteraction,
		EntityTypeNone,
		EntityTypeProfile,
		EntityTypeRelationship,
	}
}

// EntityTypes - Type of entity.
type EntityTypes string

const (
	EntityTypesInteraction  EntityTypes = "Interaction"
	EntityTypesNone         EntityTypes = "None"
	EntityTypesProfile      EntityTypes = "Profile"
	EntityTypesRelationship EntityTypes = "Relationship"
)

// PossibleEntityTypesValues returns the possible values for the EntityTypes const type.
func PossibleEntityTypesValues() []EntityTypes {
	return []EntityTypes{
		EntityTypesInteraction,
		EntityTypesNone,
		EntityTypesProfile,
		EntityTypesRelationship,
	}
}

// ErrorManagementTypes - The type of error management to use for the mapping.
type ErrorManagementTypes string

const (
	ErrorManagementTypesRejectAndContinue ErrorManagementTypes = "RejectAndContinue"
	ErrorManagementTypesRejectUntilLimit  ErrorManagementTypes = "RejectUntilLimit"
	ErrorManagementTypesStopImport        ErrorManagementTypes = "StopImport"
)

// PossibleErrorManagementTypesValues returns the possible values for the ErrorManagementTypes const type.
func PossibleErrorManagementTypesValues() []ErrorManagementTypes {
	return []ErrorManagementTypes{
		ErrorManagementTypesRejectAndContinue,
		ErrorManagementTypesRejectUntilLimit,
		ErrorManagementTypesStopImport,
	}
}

// FrequencyTypes - The frequency to update.
type FrequencyTypes string

const (
	FrequencyTypesDay    FrequencyTypes = "Day"
	FrequencyTypesHour   FrequencyTypes = "Hour"
	FrequencyTypesMinute FrequencyTypes = "Minute"
	FrequencyTypesMonth  FrequencyTypes = "Month"
	FrequencyTypesWeek   FrequencyTypes = "Week"
)

// PossibleFrequencyTypesValues returns the possible values for the FrequencyTypes const type.
func PossibleFrequencyTypesValues() []FrequencyTypes {
	return []FrequencyTypes{
		FrequencyTypesDay,
		FrequencyTypesHour,
		FrequencyTypesMinute,
		FrequencyTypesMonth,
		FrequencyTypesWeek,
	}
}

// InstanceOperationType - Determines whether this link is supposed to create or delete instances if Link is NOT Reference
// Only.
type InstanceOperationType string

const (
	InstanceOperationTypeDelete InstanceOperationType = "Delete"
	InstanceOperationTypeUpsert InstanceOperationType = "Upsert"
)

// PossibleInstanceOperationTypeValues returns the possible values for the InstanceOperationType const type.
func PossibleInstanceOperationTypeValues() []InstanceOperationType {
	return []InstanceOperationType{
		InstanceOperationTypeDelete,
		InstanceOperationTypeUpsert,
	}
}

// KpiFunctions - The computation function for the KPI.
type KpiFunctions string

const (
	KpiFunctionsAvg           KpiFunctions = "Avg"
	KpiFunctionsCount         KpiFunctions = "Count"
	KpiFunctionsCountDistinct KpiFunctions = "CountDistinct"
	KpiFunctionsLast          KpiFunctions = "Last"
	KpiFunctionsMax           KpiFunctions = "Max"
	KpiFunctionsMin           KpiFunctions = "Min"
	KpiFunctionsNone          KpiFunctions = "None"
	KpiFunctionsSum           KpiFunctions = "Sum"
)

// PossibleKpiFunctionsValues returns the possible values for the KpiFunctions const type.
func PossibleKpiFunctionsValues() []KpiFunctions {
	return []KpiFunctions{
		KpiFunctionsAvg,
		KpiFunctionsCount,
		KpiFunctionsCountDistinct,
		KpiFunctionsLast,
		KpiFunctionsMax,
		KpiFunctionsMin,
		KpiFunctionsNone,
		KpiFunctionsSum,
	}
}

// LinkTypes - Link type.
type LinkTypes string

const (
	LinkTypesCopyIfNull   LinkTypes = "CopyIfNull"
	LinkTypesUpdateAlways LinkTypes = "UpdateAlways"
)

// PossibleLinkTypesValues returns the possible values for the LinkTypes const type.
func PossibleLinkTypesValues() []LinkTypes {
	return []LinkTypes{
		LinkTypesCopyIfNull,
		LinkTypesUpdateAlways,
	}
}

// PermissionTypes - Supported permission types.
type PermissionTypes string

const (
	PermissionTypesManage PermissionTypes = "Manage"
	PermissionTypesRead   PermissionTypes = "Read"
	PermissionTypesWrite  PermissionTypes = "Write"
)

// PossiblePermissionTypesValues returns the possible values for the PermissionTypes const type.
func PossiblePermissionTypesValues() []PermissionTypes {
	return []PermissionTypes{
		PermissionTypesManage,
		PermissionTypesRead,
		PermissionTypesWrite,
	}
}

// PredictionModelLifeCycle - Prediction model life cycle. When prediction is in PendingModelConfirmation status, it is allowed
// to update the status to PendingFeaturing or Active through API.
type PredictionModelLifeCycle string

const (
	PredictionModelLifeCycleActive                   PredictionModelLifeCycle = "Active"
	PredictionModelLifeCycleDeleted                  PredictionModelLifeCycle = "Deleted"
	PredictionModelLifeCycleDiscovering              PredictionModelLifeCycle = "Discovering"
	PredictionModelLifeCycleEvaluating               PredictionModelLifeCycle = "Evaluating"
	PredictionModelLifeCycleEvaluatingFailed         PredictionModelLifeCycle = "EvaluatingFailed"
	PredictionModelLifeCycleFailed                   PredictionModelLifeCycle = "Failed"
	PredictionModelLifeCycleFeaturing                PredictionModelLifeCycle = "Featuring"
	PredictionModelLifeCycleFeaturingFailed          PredictionModelLifeCycle = "FeaturingFailed"
	PredictionModelLifeCycleHumanIntervention        PredictionModelLifeCycle = "HumanIntervention"
	PredictionModelLifeCycleNew                      PredictionModelLifeCycle = "New"
	PredictionModelLifeCyclePendingDiscovering       PredictionModelLifeCycle = "PendingDiscovering"
	PredictionModelLifeCyclePendingFeaturing         PredictionModelLifeCycle = "PendingFeaturing"
	PredictionModelLifeCyclePendingModelConfirmation PredictionModelLifeCycle = "PendingModelConfirmation"
	PredictionModelLifeCyclePendingTraining          PredictionModelLifeCycle = "PendingTraining"
	PredictionModelLifeCycleProvisioning             PredictionModelLifeCycle = "Provisioning"
	PredictionModelLifeCycleProvisioningFailed       PredictionModelLifeCycle = "ProvisioningFailed"
	PredictionModelLifeCycleTraining                 PredictionModelLifeCycle = "Training"
	PredictionModelLifeCycleTrainingFailed           PredictionModelLifeCycle = "TrainingFailed"
)

// PossiblePredictionModelLifeCycleValues returns the possible values for the PredictionModelLifeCycle const type.
func PossiblePredictionModelLifeCycleValues() []PredictionModelLifeCycle {
	return []PredictionModelLifeCycle{
		PredictionModelLifeCycleActive,
		PredictionModelLifeCycleDeleted,
		PredictionModelLifeCycleDiscovering,
		PredictionModelLifeCycleEvaluating,
		PredictionModelLifeCycleEvaluatingFailed,
		PredictionModelLifeCycleFailed,
		PredictionModelLifeCycleFeaturing,
		PredictionModelLifeCycleFeaturingFailed,
		PredictionModelLifeCycleHumanIntervention,
		PredictionModelLifeCycleNew,
		PredictionModelLifeCyclePendingDiscovering,
		PredictionModelLifeCyclePendingFeaturing,
		PredictionModelLifeCyclePendingModelConfirmation,
		PredictionModelLifeCyclePendingTraining,
		PredictionModelLifeCycleProvisioning,
		PredictionModelLifeCycleProvisioningFailed,
		PredictionModelLifeCycleTraining,
		PredictionModelLifeCycleTrainingFailed,
	}
}

// ProvisioningStates - Provisioning state.
type ProvisioningStates string

const (
	ProvisioningStatesDeleting          ProvisioningStates = "Deleting"
	ProvisioningStatesExpiring          ProvisioningStates = "Expiring"
	ProvisioningStatesFailed            ProvisioningStates = "Failed"
	ProvisioningStatesHumanIntervention ProvisioningStates = "HumanIntervention"
	ProvisioningStatesProvisioning      ProvisioningStates = "Provisioning"
	ProvisioningStatesSucceeded         ProvisioningStates = "Succeeded"
)

// PossibleProvisioningStatesValues returns the possible values for the ProvisioningStates const type.
func PossibleProvisioningStatesValues() []ProvisioningStates {
	return []ProvisioningStates{
		ProvisioningStatesDeleting,
		ProvisioningStatesExpiring,
		ProvisioningStatesFailed,
		ProvisioningStatesHumanIntervention,
		ProvisioningStatesProvisioning,
		ProvisioningStatesSucceeded,
	}
}

// RoleTypes - Type of roles.
type RoleTypes string

const (
	RoleTypesAdmin        RoleTypes = "Admin"
	RoleTypesDataAdmin    RoleTypes = "DataAdmin"
	RoleTypesDataReader   RoleTypes = "DataReader"
	RoleTypesManageAdmin  RoleTypes = "ManageAdmin"
	RoleTypesManageReader RoleTypes = "ManageReader"
	RoleTypesReader       RoleTypes = "Reader"
)

// PossibleRoleTypesValues returns the possible values for the RoleTypes const type.
func PossibleRoleTypesValues() []RoleTypes {
	return []RoleTypes{
		RoleTypesAdmin,
		RoleTypesDataAdmin,
		RoleTypesDataReader,
		RoleTypesManageAdmin,
		RoleTypesManageReader,
		RoleTypesReader,
	}
}

// Status - The data source status.
type Status string

const (
	StatusActive  Status = "Active"
	StatusDeleted Status = "Deleted"
	StatusNone    Status = "None"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusActive,
		StatusDeleted,
		StatusNone,
	}
}