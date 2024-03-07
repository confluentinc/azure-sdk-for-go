//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybriddatamanager

import "time"

// AvailableProviderOperation - Class represents provider operation
type AvailableProviderOperation struct {
	// REQUIRED; Gets or Sets Name of the operations
	Name *string

	// Gets or sets Display information Contains the localized display information for this particular operation/action
	Display *AvailableProviderOperationDisplay

	// Gets or sets Origin The intended executor of the operation; governs the display of the operation in the RBAC UX and the
	// audit logs UX. Default value is “user,system”
	Origin *string

	// Gets or sets Properties Reserved for future use
	Properties any
}

// AvailableProviderOperationDisplay - Contains the localized display information for this particular operation / action.
// These value will be used by several clients for (1) custom role definitions for RBAC; (2) complex query filters for
// the event service; and (3) audit history / records for management operations.
type AvailableProviderOperationDisplay struct {
	// Gets or sets Description The localized friendly description for the operation, as it should be shown to the user. It should
	// be thorough, yet concise – it will be used in tool tips and detailed views.
	Description *string

	// Gets or sets Operation The localized friendly name for the operation, as it should be shown to the user. It should be concise
	// (to fit in drop downs) but clear (i.e. self-documenting). It should use
	// Title Casing and include the entity/resource to which it applies.
	Operation *string

	// Gets or sets Provider The localized friendly form of the resource provider name – it is expected to also include the publisher/company
	// responsible. It should use Title Casing and begin with
	// “Microsoft” for 1st party services.
	Provider *string

	// Gets or sets Resource The localized friendly form of the resource type related to this action/operation – it should match
	// the public documentation for the resource provider. It should use Title Casing
	// – for examples, please refer to the “name” section.
	Resource *string
}

// AvailableProviderOperations - Class for set of operations used for discovery of available provider operations.
type AvailableProviderOperations struct {
	// Link for the next set of operations.
	NextLink *string

	// List of operations.
	Value []*AvailableProviderOperation
}

// CustomerSecret - The pair of customer secret.
type CustomerSecret struct {
	// REQUIRED; The encryption algorithm used to encrypt data.
	Algorithm *SupportedAlgorithm

	// REQUIRED; The identifier to the data service input object which this secret corresponds to.
	KeyIdentifier *string

	// REQUIRED; It contains the encrypted customer secret.
	KeyValue *string
}

// DataManager - The DataManager resource.
type DataManager struct {
	// REQUIRED; The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US,
	// East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it
	// is created, but if an identical geo region is specified on update the request will succeed.
	Location *string

	// Etag of the Resource.
	Etag *string

	// The sku type.
	SKU *SKU

	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across
	// resource groups).
	Tags map[string]*string

	// READ-ONLY; The Resource Id.
	ID *string

	// READ-ONLY; The Resource Name.
	Name *string

	// READ-ONLY; The Resource type.
	Type *string
}

// DataManagerList - DataManager resources Collection.
type DataManagerList struct {
	// Link for the next set of data stores.
	NextLink *string

	// List of data manager resources.
	Value []*DataManager
}

// DataManagerUpdateParameter - The DataManagerUpdateParameter.
type DataManagerUpdateParameter struct {
	// The sku type.
	SKU *SKU

	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across
	// resource groups).
	Tags map[string]*string
}

// DataService - Data Service.
type DataService struct {
	// REQUIRED; DataService properties.
	Properties *DataServiceProperties

	// READ-ONLY; Id of the object.
	ID *string

	// READ-ONLY; Name of the object.
	Name *string

	// READ-ONLY; Type of the object.
	Type *string
}

// DataServiceList - Data Service Collection.
type DataServiceList struct {
	// Link for the next set of data services.
	NextLink *string

	// List of data services.
	Value []*DataService
}

// DataServiceProperties - Data Service properties.
type DataServiceProperties struct {
	// REQUIRED; State of the data service.
	State *State

	// Supported data store types which can be used as a sink.
	SupportedDataSinkTypes []*string

	// Supported data store types which can be used as a source.
	SupportedDataSourceTypes []*string
}

// DataStore - Data store.
type DataStore struct {
	// REQUIRED; DataStore properties.
	Properties *DataStoreProperties

	// READ-ONLY; Id of the object.
	ID *string

	// READ-ONLY; Name of the object.
	Name *string

	// READ-ONLY; Type of the object.
	Type *string
}

// DataStoreFilter - Contains the information about the filters for the DataStore.
type DataStoreFilter struct {
	// The data store type id.
	DataStoreTypeID *string
}

// DataStoreList - Data Store Collection.
type DataStoreList struct {
	// Link for the next set of data stores.
	NextLink *string

	// List of data stores.
	Value []*DataStore
}

// DataStoreProperties - Data Store for sources and sinks
type DataStoreProperties struct {
	// REQUIRED; The arm id of the data store type.
	DataStoreTypeID *string

	// REQUIRED; State of the data source.
	State *State

	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source
	// to understand the key. Value contains customer secret encrypted by the
	// encryptionKeys.
	CustomerSecrets []*CustomerSecret

	// A generic json used differently by each data source type.
	ExtendedProperties any

	// Arm Id for the manager resource to which the data source is associated. This is optional.
	RepositoryID *string
}

// DataStoreType - Data Store Type.
type DataStoreType struct {
	// REQUIRED; DataStoreType properties.
	Properties *DataStoreTypeProperties

	// READ-ONLY; Id of the object.
	ID *string

	// READ-ONLY; Name of the object.
	Name *string

	// READ-ONLY; Type of the object.
	Type *string
}

// DataStoreTypeList - Data Store Type Collection.
type DataStoreTypeList struct {
	// Link for the next set of data store types.
	NextLink *string

	// List of DataStoreType.
	Value []*DataStoreType
}

// DataStoreTypeProperties - Data Store Type properties.
type DataStoreTypeProperties struct {
	// REQUIRED; State of the data store type.
	State *State

	// Arm type for the manager resource to which the data source type is associated. This is optional.
	RepositoryType *string

	// Supported data services where it can be used as a sink.
	SupportedDataServicesAsSink []*string

	// Supported data services where it can be used as a source.
	SupportedDataServicesAsSource []*string
}

// DmsBaseObject - Base class for all objects under DataManager Service
type DmsBaseObject struct {
	// READ-ONLY; Id of the object.
	ID *string

	// READ-ONLY; Name of the object.
	Name *string

	// READ-ONLY; Type of the object.
	Type *string
}

// Error - Top level error for the job.
type Error struct {
	// REQUIRED; Error code that can be used to programmatically identify the error.
	Code *string

	// Describes the error in detail and provides debugging information.
	Message *string
}

// ErrorDetails - Error Details
type ErrorDetails struct {
	// Error code.
	ErrorCode *int32

	// Error message.
	ErrorMessage *string

	// Contains the non localized exception message
	ExceptionMessage *string

	// Recommended action for the error.
	RecommendedAction *string
}

// Job - Data service job.
type Job struct {
	// REQUIRED; Job properties.
	Properties *JobProperties

	// REQUIRED; Time at which the job was started in UTC ISO 8601 format.
	StartTime *time.Time

	// REQUIRED; Status of the job.
	Status *JobStatus

	// Time at which the job ended in UTC ISO 8601 format.
	EndTime *time.Time

	// Top level error for the job.
	Error *Error

	// READ-ONLY; Id of the object.
	ID *string

	// READ-ONLY; Name of the object.
	Name *string

	// READ-ONLY; Type of the object.
	Type *string
}

// JobDefinition - Job Definition.
type JobDefinition struct {
	// REQUIRED; JobDefinition properties.
	Properties *JobDefinitionProperties

	// READ-ONLY; Id of the object.
	ID *string

	// READ-ONLY; Name of the object.
	Name *string

	// READ-ONLY; Type of the object.
	Type *string
}

// JobDefinitionFilter - Contains the supported job definition filters.
type JobDefinitionFilter struct {
	// REQUIRED; The state of the job definition.
	State *State

	// The data source associated with the job definition
	DataSource *string

	// The last modified date time of the data source.
	LastModified *time.Time
}

// JobDefinitionList - Job Definition Collection.
type JobDefinitionList struct {
	// Link for the next set of job definitions.
	NextLink *string

	// List of job definitions.
	Value []*JobDefinition
}

// JobDefinitionProperties - Job Definition
type JobDefinitionProperties struct {
	// REQUIRED; Data Sink Id associated to the job definition.
	DataSinkID *string

	// REQUIRED; Data Source Id associated to the job definition.
	DataSourceID *string

	// REQUIRED; State of the job definition.
	State *State

	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source
	// to understand the key. Value contains customer secret encrypted by the
	// encryptionKeys.
	CustomerSecrets []*CustomerSecret

	// A generic json used differently by each data service type.
	DataServiceInput any

	// Last modified time of the job definition.
	LastModifiedTime *time.Time

	// This is the preferred geo location for the job to run.
	RunLocation *RunLocation

	// Schedule for running the job definition
	Schedules []*Schedule

	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation *UserConfirmation
}

// JobDetails - Job details.
type JobDetails struct {
	// Error details for failure. This is optional.
	ErrorDetails []*ErrorDetails

	// Item Details Link to download files or see details
	ItemDetailsLink *string

	// JobDefinition at the time of the run
	JobDefinition *JobDefinition

	// List of stages that ran in the job
	JobStages []*JobStages
}

// JobFilter - Contains the information about the filters for the job.
type JobFilter struct {
	// REQUIRED; The status of the job.
	Status *JobStatus

	// The start time of the job.
	StartTime *time.Time
}

// JobList - Job Collection.
type JobList struct {
	// Link for the next set of jobs.
	NextLink *string

	// List of jobs.
	Value []*Job
}

// JobProperties - Job Properties
type JobProperties struct {
	// REQUIRED; Describes whether the job is cancellable.
	IsCancellable *IsJobCancellable

	// Number of bytes processed by the job as of now.
	BytesProcessed *int64

	// Name of the data sink on which the job was triggered.
	DataSinkName *string

	// Name of the data source on which the job was triggered.
	DataSourceName *string

	// Details of a job run. This field will only be sent for expand details filter.
	Details *JobDetails

	// Number of items processed by the job as of now
	ItemsProcessed *int64

	// Number of bytes to be processed by the job in total.
	TotalBytesToProcess *int64

	// Number of items to be processed by the job in total
	TotalItemsToProcess *int64
}

// JobStages - Job stages.
type JobStages struct {
	// REQUIRED; Status of the job stage.
	StageStatus *JobStatus

	// Error details for the stage. This is optional
	ErrorDetails []*ErrorDetails

	// Job Stage Details
	JobStageDetails any

	// Name of the job stage.
	StageName *string
}

// Key - Encryption Key.
type Key struct {
	// REQUIRED; The maximum byte size that can be encrypted by the key. For a key size larger than the size, break into chunks
	// and encrypt each chunk, append each encrypted chunk with : to mark the end of the chunk.
	EncryptionChunkSizeInBytes *int32

	// REQUIRED; Exponent of the encryption key.
	KeyExponent *string

	// REQUIRED; Modulus of the encryption key.
	KeyModulus *string
}

// PublicKey - Public key
type PublicKey struct {
	// REQUIRED; Public key property.
	Properties *PublicKeyProperties

	// READ-ONLY; Id of the object.
	ID *string

	// READ-ONLY; Name of the object.
	Name *string

	// READ-ONLY; Type of the object.
	Type *string
}

// PublicKeyList - PublicKey Collection
type PublicKeyList struct {
	// Link for the next set of public keys.
	NextLink *string

	// List of public keys.
	Value []*PublicKey
}

// PublicKeyProperties - PublicKey Properties
type PublicKeyProperties struct {
	// REQUIRED; Level one public key for encryption
	DataServiceLevel1Key *Key

	// REQUIRED; Level two public key for encryption
	DataServiceLevel2Key *Key
}

// Resource - Model of the Resource.
type Resource struct {
	// REQUIRED; The location of the resource. This will be one of the supported and registered Azure Geo Regions (e.g. West US,
	// East US, Southeast Asia, etc.). The geo region of a resource cannot be changed once it
	// is created, but if an identical geo region is specified on update the request will succeed.
	Location *string

	// The sku type.
	SKU *SKU

	// The list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across
	// resource groups).
	Tags map[string]*string

	// READ-ONLY; The Resource Id.
	ID *string

	// READ-ONLY; The Resource Name.
	Name *string

	// READ-ONLY; The Resource type.
	Type *string
}

// RunParameters - Run parameters for a job.
type RunParameters struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source
	// to understand the key. Value contains customer secret encrypted by the
	// encryptionKeys.
	CustomerSecrets []*CustomerSecret

	// A generic json used differently by each data service type.
	DataServiceInput any

	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation *UserConfirmation
}

// SKU - The sku type.
type SKU struct {
	// The sku name. Required for data manager creation, optional for update.
	Name *string

	// The sku tier. This is based on the SKU name.
	Tier *string
}

// Schedule for the job run.
type Schedule struct {
	// Name of the schedule.
	Name *string

	// A list of repetition intervals in ISO 8601 format.
	PolicyList []*string
}