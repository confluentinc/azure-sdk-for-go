// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package appplatform

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/appplatform/mgmt/2019-05-01-preview/appplatform"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AppResourceProvisioningState = original.AppResourceProvisioningState

const (
	Failed    AppResourceProvisioningState = original.Failed
	Succeeded AppResourceProvisioningState = original.Succeeded
)

type ConfigServerState = original.ConfigServerState

const (
	ConfigServerStateDeleted      ConfigServerState = original.ConfigServerStateDeleted
	ConfigServerStateFailed       ConfigServerState = original.ConfigServerStateFailed
	ConfigServerStateNotAvailable ConfigServerState = original.ConfigServerStateNotAvailable
	ConfigServerStateSucceeded    ConfigServerState = original.ConfigServerStateSucceeded
	ConfigServerStateUpdating     ConfigServerState = original.ConfigServerStateUpdating
)

type DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningState

const (
	DeploymentResourceProvisioningStateCreating   DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateCreating
	DeploymentResourceProvisioningStateFailed     DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateFailed
	DeploymentResourceProvisioningStateProcessing DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateProcessing
	DeploymentResourceProvisioningStateSucceeded  DeploymentResourceProvisioningState = original.DeploymentResourceProvisioningStateSucceeded
)

type DeploymentResourceStatus = original.DeploymentResourceStatus

const (
	DeploymentResourceStatusAllocating DeploymentResourceStatus = original.DeploymentResourceStatusAllocating
	DeploymentResourceStatusCompiling  DeploymentResourceStatus = original.DeploymentResourceStatusCompiling
	DeploymentResourceStatusFailed     DeploymentResourceStatus = original.DeploymentResourceStatusFailed
	DeploymentResourceStatusProcessing DeploymentResourceStatus = original.DeploymentResourceStatusProcessing
	DeploymentResourceStatusRunning    DeploymentResourceStatus = original.DeploymentResourceStatusRunning
	DeploymentResourceStatusStopped    DeploymentResourceStatus = original.DeploymentResourceStatusStopped
	DeploymentResourceStatusUnknown    DeploymentResourceStatus = original.DeploymentResourceStatusUnknown
	DeploymentResourceStatusUpgrading  DeploymentResourceStatus = original.DeploymentResourceStatusUpgrading
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCreating   ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleted    ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting   ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed     ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoved      ProvisioningState = original.ProvisioningStateMoved
	ProvisioningStateMoveFailed ProvisioningState = original.ProvisioningStateMoveFailed
	ProvisioningStateMoving     ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateSucceeded  ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateUpdating   ProvisioningState = original.ProvisioningStateUpdating
)

type RuntimeVersion = original.RuntimeVersion

const (
	Java11 RuntimeVersion = original.Java11
	Java8  RuntimeVersion = original.Java8
)

type TestKeyType = original.TestKeyType

const (
	Primary   TestKeyType = original.Primary
	Secondary TestKeyType = original.Secondary
)

type TraceProxyState = original.TraceProxyState

const (
	TraceProxyStateFailed       TraceProxyState = original.TraceProxyStateFailed
	TraceProxyStateNotAvailable TraceProxyState = original.TraceProxyStateNotAvailable
	TraceProxyStateSucceeded    TraceProxyState = original.TraceProxyStateSucceeded
	TraceProxyStateUpdating     TraceProxyState = original.TraceProxyStateUpdating
)

type UserSourceType = original.UserSourceType

const (
	Jar    UserSourceType = original.Jar
	Source UserSourceType = original.Source
)

type AppResource = original.AppResource
type AppResourceCollection = original.AppResourceCollection
type AppResourceCollectionIterator = original.AppResourceCollectionIterator
type AppResourceCollectionPage = original.AppResourceCollectionPage
type AppResourceProperties = original.AppResourceProperties
type AppsClient = original.AppsClient
type AvailableOperations = original.AvailableOperations
type AvailableOperationsIterator = original.AvailableOperationsIterator
type AvailableOperationsPage = original.AvailableOperationsPage
type BaseClient = original.BaseClient
type BindingResource = original.BindingResource
type BindingResourceCollection = original.BindingResourceCollection
type BindingResourceCollectionIterator = original.BindingResourceCollectionIterator
type BindingResourceCollectionPage = original.BindingResourceCollectionPage
type BindingResourceProperties = original.BindingResourceProperties
type BindingsClient = original.BindingsClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ClusterResourceProperties = original.ClusterResourceProperties
type ConfigServerGitProperty = original.ConfigServerGitProperty
type ConfigServerProperties = original.ConfigServerProperties
type ConfigServerSettings = original.ConfigServerSettings
type DeploymentInstance = original.DeploymentInstance
type DeploymentResource = original.DeploymentResource
type DeploymentResourceCollection = original.DeploymentResourceCollection
type DeploymentResourceCollectionIterator = original.DeploymentResourceCollectionIterator
type DeploymentResourceCollectionPage = original.DeploymentResourceCollectionPage
type DeploymentResourceProperties = original.DeploymentResourceProperties
type DeploymentSettings = original.DeploymentSettings
type DeploymentsClient = original.DeploymentsClient
type DeploymentsCreateOrUpdateFuture = original.DeploymentsCreateOrUpdateFuture
type DeploymentsRestartFuture = original.DeploymentsRestartFuture
type DeploymentsStartFuture = original.DeploymentsStartFuture
type DeploymentsStopFuture = original.DeploymentsStopFuture
type DeploymentsUpdateFuture = original.DeploymentsUpdateFuture
type Error = original.Error
type GitPatternRepository = original.GitPatternRepository
type LogFileURLResponse = original.LogFileURLResponse
type LogSpecification = original.LogSpecification
type MetricDimension = original.MetricDimension
type MetricSpecification = original.MetricSpecification
type NameAvailability = original.NameAvailability
type NameAvailabilityParameters = original.NameAvailabilityParameters
type OperationDetail = original.OperationDetail
type OperationDisplay = original.OperationDisplay
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type PersistentDisk = original.PersistentDisk
type ProxyResource = original.ProxyResource
type RegenerateTestKeyRequestPayload = original.RegenerateTestKeyRequestPayload
type Resource = original.Resource
type ResourceUploadDefinition = original.ResourceUploadDefinition
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceListIterator = original.ServiceResourceListIterator
type ServiceResourceListPage = original.ServiceResourceListPage
type ServiceSpecification = original.ServiceSpecification
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesUpdateFuture = original.ServicesUpdateFuture
type TemporaryDisk = original.TemporaryDisk
type TestKeys = original.TestKeys
type TraceProperties = original.TraceProperties
type TrackedResource = original.TrackedResource
type UserSourceInfo = original.UserSourceInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAppResourceCollectionIterator(page AppResourceCollectionPage) AppResourceCollectionIterator {
	return original.NewAppResourceCollectionIterator(page)
}
func NewAppResourceCollectionPage(getNextPage func(context.Context, AppResourceCollection) (AppResourceCollection, error)) AppResourceCollectionPage {
	return original.NewAppResourceCollectionPage(getNextPage)
}
func NewAppsClient(subscriptionID string) AppsClient {
	return original.NewAppsClient(subscriptionID)
}
func NewAppsClientWithBaseURI(baseURI string, subscriptionID string) AppsClient {
	return original.NewAppsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAvailableOperationsIterator(page AvailableOperationsPage) AvailableOperationsIterator {
	return original.NewAvailableOperationsIterator(page)
}
func NewAvailableOperationsPage(getNextPage func(context.Context, AvailableOperations) (AvailableOperations, error)) AvailableOperationsPage {
	return original.NewAvailableOperationsPage(getNextPage)
}
func NewBindingResourceCollectionIterator(page BindingResourceCollectionPage) BindingResourceCollectionIterator {
	return original.NewBindingResourceCollectionIterator(page)
}
func NewBindingResourceCollectionPage(getNextPage func(context.Context, BindingResourceCollection) (BindingResourceCollection, error)) BindingResourceCollectionPage {
	return original.NewBindingResourceCollectionPage(getNextPage)
}
func NewBindingsClient(subscriptionID string) BindingsClient {
	return original.NewBindingsClient(subscriptionID)
}
func NewBindingsClientWithBaseURI(baseURI string, subscriptionID string) BindingsClient {
	return original.NewBindingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeploymentResourceCollectionIterator(page DeploymentResourceCollectionPage) DeploymentResourceCollectionIterator {
	return original.NewDeploymentResourceCollectionIterator(page)
}
func NewDeploymentResourceCollectionPage(getNextPage func(context.Context, DeploymentResourceCollection) (DeploymentResourceCollection, error)) DeploymentResourceCollectionPage {
	return original.NewDeploymentResourceCollectionPage(getNextPage)
}
func NewDeploymentsClient(subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClient(subscriptionID)
}
func NewDeploymentsClientWithBaseURI(baseURI string, subscriptionID string) DeploymentsClient {
	return original.NewDeploymentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceResourceListIterator(page ServiceResourceListPage) ServiceResourceListIterator {
	return original.NewServiceResourceListIterator(page)
}
func NewServiceResourceListPage(getNextPage func(context.Context, ServiceResourceList) (ServiceResourceList, error)) ServiceResourceListPage {
	return original.NewServiceResourceListPage(getNextPage)
}
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAppResourceProvisioningStateValues() []AppResourceProvisioningState {
	return original.PossibleAppResourceProvisioningStateValues()
}
func PossibleConfigServerStateValues() []ConfigServerState {
	return original.PossibleConfigServerStateValues()
}
func PossibleDeploymentResourceProvisioningStateValues() []DeploymentResourceProvisioningState {
	return original.PossibleDeploymentResourceProvisioningStateValues()
}
func PossibleDeploymentResourceStatusValues() []DeploymentResourceStatus {
	return original.PossibleDeploymentResourceStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleRuntimeVersionValues() []RuntimeVersion {
	return original.PossibleRuntimeVersionValues()
}
func PossibleTestKeyTypeValues() []TestKeyType {
	return original.PossibleTestKeyTypeValues()
}
func PossibleTraceProxyStateValues() []TraceProxyState {
	return original.PossibleTraceProxyStateValues()
}
func PossibleUserSourceTypeValues() []UserSourceType {
	return original.PossibleUserSourceTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
