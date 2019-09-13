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

package vmwarecloudsimple

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/vmwarecloudsimple/mgmt/2019-04-01/vmwarecloudsimple"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationType = original.AggregationType

const (
	Average AggregationType = original.Average
	Total   AggregationType = original.Total
)

type DiskIndependenceMode = original.DiskIndependenceMode

const (
	IndependentNonpersistent DiskIndependenceMode = original.IndependentNonpersistent
	IndependentPersistent    DiskIndependenceMode = original.IndependentPersistent
	Persistent               DiskIndependenceMode = original.Persistent
)

type GuestOSType = original.GuestOSType

const (
	Linux   GuestOSType = original.Linux
	Other   GuestOSType = original.Other
	Windows GuestOSType = original.Windows
)

type NICType = original.NICType

const (
	E1000   NICType = original.E1000
	E1000E  NICType = original.E1000E
	PCNET32 NICType = original.PCNET32
	VMXNET  NICType = original.VMXNET
	VMXNET2 NICType = original.VMXNET2
	VMXNET3 NICType = original.VMXNET3
)

type NodeStatus = original.NodeStatus

const (
	Unused NodeStatus = original.Unused
	Used   NodeStatus = original.Used
)

type OnboardingStatus = original.OnboardingStatus

const (
	NotOnBoarded     OnboardingStatus = original.NotOnBoarded
	OnBoarded        OnboardingStatus = original.OnBoarded
	OnBoarding       OnboardingStatus = original.OnBoarding
	OnBoardingFailed OnboardingStatus = original.OnBoardingFailed
)

type OperationOrigin = original.OperationOrigin

const (
	System     OperationOrigin = original.System
	User       OperationOrigin = original.User
	Usersystem OperationOrigin = original.Usersystem
)

type PrivateCloudResourceType = original.PrivateCloudResourceType

const (
	MicrosoftVMwareCloudSimpleprivateClouds PrivateCloudResourceType = original.MicrosoftVMwareCloudSimpleprivateClouds
)

type StopMode = original.StopMode

const (
	Poweroff StopMode = original.Poweroff
	Reboot   StopMode = original.Reboot
	Shutdown StopMode = original.Shutdown
	Suspend  StopMode = original.Suspend
)

type UsageCount = original.UsageCount

const (
	Bytes          UsageCount = original.Bytes
	BytesPerSecond UsageCount = original.BytesPerSecond
	Count          UsageCount = original.Count
	CountPerSecond UsageCount = original.CountPerSecond
	Percent        UsageCount = original.Percent
	Seconds        UsageCount = original.Seconds
)

type VirtualMachineStatus = original.VirtualMachineStatus

const (
	Deallocating VirtualMachineStatus = original.Deallocating
	Deleting     VirtualMachineStatus = original.Deleting
	Poweredoff   VirtualMachineStatus = original.Poweredoff
	Running      VirtualMachineStatus = original.Running
	Suspended    VirtualMachineStatus = original.Suspended
	Updating     VirtualMachineStatus = original.Updating
)

type AvailableOperation = original.AvailableOperation
type AvailableOperationDisplay = original.AvailableOperationDisplay
type AvailableOperationDisplayPropertyServiceSpecification = original.AvailableOperationDisplayPropertyServiceSpecification
type AvailableOperationDisplayPropertyServiceSpecificationMetricsItem = original.AvailableOperationDisplayPropertyServiceSpecificationMetricsItem
type AvailableOperationDisplayPropertyServiceSpecificationMetricsList = original.AvailableOperationDisplayPropertyServiceSpecificationMetricsList
type AvailableOperationsListResponse = original.AvailableOperationsListResponse
type AvailableOperationsListResponseIterator = original.AvailableOperationsListResponseIterator
type AvailableOperationsListResponsePage = original.AvailableOperationsListResponsePage
type BaseClient = original.BaseClient
type CSRPError = original.CSRPError
type CSRPErrorBody = original.CSRPErrorBody
type CreateDedicatedCloudNodeProperties = original.CreateDedicatedCloudNodeProperties
type CreateDedicatedCloudNodeRequest = original.CreateDedicatedCloudNodeRequest
type DedicatedCloudNode = original.DedicatedCloudNode
type DedicatedCloudNodeListResponse = original.DedicatedCloudNodeListResponse
type DedicatedCloudNodeListResponseIterator = original.DedicatedCloudNodeListResponseIterator
type DedicatedCloudNodeListResponsePage = original.DedicatedCloudNodeListResponsePage
type DedicatedCloudNodeProperties = original.DedicatedCloudNodeProperties
type DedicatedCloudNodesClient = original.DedicatedCloudNodesClient
type DedicatedCloudNodesCreateOrUpdateFuture = original.DedicatedCloudNodesCreateOrUpdateFuture
type DedicatedCloudService = original.DedicatedCloudService
type DedicatedCloudServiceListResponse = original.DedicatedCloudServiceListResponse
type DedicatedCloudServiceListResponseIterator = original.DedicatedCloudServiceListResponseIterator
type DedicatedCloudServiceListResponsePage = original.DedicatedCloudServiceListResponsePage
type DedicatedCloudServiceProperties = original.DedicatedCloudServiceProperties
type DedicatedCloudServicesClient = original.DedicatedCloudServicesClient
type DedicatedCloudServicesDeleteFuture = original.DedicatedCloudServicesDeleteFuture
type OperationError = original.OperationError
type OperationResource = original.OperationResource
type OperationsClient = original.OperationsClient
type PatchPayload = original.PatchPayload
type PrivateCloud = original.PrivateCloud
type PrivateCloudList = original.PrivateCloudList
type PrivateCloudListIterator = original.PrivateCloudListIterator
type PrivateCloudListPage = original.PrivateCloudListPage
type PrivateCloudProperties = original.PrivateCloudProperties
type PrivateCloudsClient = original.PrivateCloudsClient
type ResourcePool = original.ResourcePool
type ResourcePoolProperties = original.ResourcePoolProperties
type ResourcePoolsClient = original.ResourcePoolsClient
type ResourcePoolsListResponse = original.ResourcePoolsListResponse
type ResourcePoolsListResponseIterator = original.ResourcePoolsListResponseIterator
type ResourcePoolsListResponsePage = original.ResourcePoolsListResponsePage
type Sku = original.Sku
type SkuAvailability = original.SkuAvailability
type SkuAvailabilityListResponse = original.SkuAvailabilityListResponse
type SkuAvailabilityListResponseIterator = original.SkuAvailabilityListResponseIterator
type SkuAvailabilityListResponsePage = original.SkuAvailabilityListResponsePage
type SkuDescription = original.SkuDescription
type SkusAvailabilityClient = original.SkusAvailabilityClient
type Usage = original.Usage
type UsageListResponse = original.UsageListResponse
type UsageListResponseIterator = original.UsageListResponseIterator
type UsageListResponsePage = original.UsageListResponsePage
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient
type VirtualDisk = original.VirtualDisk
type VirtualDiskController = original.VirtualDiskController
type VirtualMachine = original.VirtualMachine
type VirtualMachineListResponse = original.VirtualMachineListResponse
type VirtualMachineListResponseIterator = original.VirtualMachineListResponseIterator
type VirtualMachineListResponsePage = original.VirtualMachineListResponsePage
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineStopMode = original.VirtualMachineStopMode
type VirtualMachineTemplate = original.VirtualMachineTemplate
type VirtualMachineTemplateListResponse = original.VirtualMachineTemplateListResponse
type VirtualMachineTemplateListResponseIterator = original.VirtualMachineTemplateListResponseIterator
type VirtualMachineTemplateListResponsePage = original.VirtualMachineTemplateListResponsePage
type VirtualMachineTemplateProperties = original.VirtualMachineTemplateProperties
type VirtualMachineTemplatesClient = original.VirtualMachineTemplatesClient
type VirtualMachinesClient = original.VirtualMachinesClient
type VirtualMachinesCreateOrUpdateFuture = original.VirtualMachinesCreateOrUpdateFuture
type VirtualMachinesDeleteFuture = original.VirtualMachinesDeleteFuture
type VirtualMachinesStartFuture = original.VirtualMachinesStartFuture
type VirtualMachinesStopFuture = original.VirtualMachinesStopFuture
type VirtualMachinesUpdateFuture = original.VirtualMachinesUpdateFuture
type VirtualNetwork = original.VirtualNetwork
type VirtualNetworkListResponse = original.VirtualNetworkListResponse
type VirtualNetworkListResponseIterator = original.VirtualNetworkListResponseIterator
type VirtualNetworkListResponsePage = original.VirtualNetworkListResponsePage
type VirtualNetworkProperties = original.VirtualNetworkProperties
type VirtualNetworksClient = original.VirtualNetworksClient
type VirtualNic = original.VirtualNic

func New(referer string, subscriptionID string) BaseClient {
	return original.New(referer, subscriptionID)
}
func NewAvailableOperationsListResponseIterator(page AvailableOperationsListResponsePage) AvailableOperationsListResponseIterator {
	return original.NewAvailableOperationsListResponseIterator(page)
}
func NewAvailableOperationsListResponsePage(getNextPage func(context.Context, AvailableOperationsListResponse) (AvailableOperationsListResponse, error)) AvailableOperationsListResponsePage {
	return original.NewAvailableOperationsListResponsePage(getNextPage)
}
func NewDedicatedCloudNodeListResponseIterator(page DedicatedCloudNodeListResponsePage) DedicatedCloudNodeListResponseIterator {
	return original.NewDedicatedCloudNodeListResponseIterator(page)
}
func NewDedicatedCloudNodeListResponsePage(getNextPage func(context.Context, DedicatedCloudNodeListResponse) (DedicatedCloudNodeListResponse, error)) DedicatedCloudNodeListResponsePage {
	return original.NewDedicatedCloudNodeListResponsePage(getNextPage)
}
func NewDedicatedCloudNodesClient(referer string, subscriptionID string) DedicatedCloudNodesClient {
	return original.NewDedicatedCloudNodesClient(referer, subscriptionID)
}
func NewDedicatedCloudNodesClientWithBaseURI(baseURI string, referer string, subscriptionID string) DedicatedCloudNodesClient {
	return original.NewDedicatedCloudNodesClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewDedicatedCloudServiceListResponseIterator(page DedicatedCloudServiceListResponsePage) DedicatedCloudServiceListResponseIterator {
	return original.NewDedicatedCloudServiceListResponseIterator(page)
}
func NewDedicatedCloudServiceListResponsePage(getNextPage func(context.Context, DedicatedCloudServiceListResponse) (DedicatedCloudServiceListResponse, error)) DedicatedCloudServiceListResponsePage {
	return original.NewDedicatedCloudServiceListResponsePage(getNextPage)
}
func NewDedicatedCloudServicesClient(referer string, subscriptionID string) DedicatedCloudServicesClient {
	return original.NewDedicatedCloudServicesClient(referer, subscriptionID)
}
func NewDedicatedCloudServicesClientWithBaseURI(baseURI string, referer string, subscriptionID string) DedicatedCloudServicesClient {
	return original.NewDedicatedCloudServicesClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewOperationsClient(referer string, subscriptionID string) OperationsClient {
	return original.NewOperationsClient(referer, subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, referer string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewPrivateCloudListIterator(page PrivateCloudListPage) PrivateCloudListIterator {
	return original.NewPrivateCloudListIterator(page)
}
func NewPrivateCloudListPage(getNextPage func(context.Context, PrivateCloudList) (PrivateCloudList, error)) PrivateCloudListPage {
	return original.NewPrivateCloudListPage(getNextPage)
}
func NewPrivateCloudsClient(referer string, subscriptionID string) PrivateCloudsClient {
	return original.NewPrivateCloudsClient(referer, subscriptionID)
}
func NewPrivateCloudsClientWithBaseURI(baseURI string, referer string, subscriptionID string) PrivateCloudsClient {
	return original.NewPrivateCloudsClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewResourcePoolsClient(referer string, subscriptionID string) ResourcePoolsClient {
	return original.NewResourcePoolsClient(referer, subscriptionID)
}
func NewResourcePoolsClientWithBaseURI(baseURI string, referer string, subscriptionID string) ResourcePoolsClient {
	return original.NewResourcePoolsClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewResourcePoolsListResponseIterator(page ResourcePoolsListResponsePage) ResourcePoolsListResponseIterator {
	return original.NewResourcePoolsListResponseIterator(page)
}
func NewResourcePoolsListResponsePage(getNextPage func(context.Context, ResourcePoolsListResponse) (ResourcePoolsListResponse, error)) ResourcePoolsListResponsePage {
	return original.NewResourcePoolsListResponsePage(getNextPage)
}
func NewSkuAvailabilityListResponseIterator(page SkuAvailabilityListResponsePage) SkuAvailabilityListResponseIterator {
	return original.NewSkuAvailabilityListResponseIterator(page)
}
func NewSkuAvailabilityListResponsePage(getNextPage func(context.Context, SkuAvailabilityListResponse) (SkuAvailabilityListResponse, error)) SkuAvailabilityListResponsePage {
	return original.NewSkuAvailabilityListResponsePage(getNextPage)
}
func NewSkusAvailabilityClient(referer string, subscriptionID string) SkusAvailabilityClient {
	return original.NewSkusAvailabilityClient(referer, subscriptionID)
}
func NewSkusAvailabilityClientWithBaseURI(baseURI string, referer string, subscriptionID string) SkusAvailabilityClient {
	return original.NewSkusAvailabilityClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewUsageListResponseIterator(page UsageListResponsePage) UsageListResponseIterator {
	return original.NewUsageListResponseIterator(page)
}
func NewUsageListResponsePage(getNextPage func(context.Context, UsageListResponse) (UsageListResponse, error)) UsageListResponsePage {
	return original.NewUsageListResponsePage(getNextPage)
}
func NewUsagesClient(referer string, subscriptionID string) UsagesClient {
	return original.NewUsagesClient(referer, subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, referer string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewVirtualMachineListResponseIterator(page VirtualMachineListResponsePage) VirtualMachineListResponseIterator {
	return original.NewVirtualMachineListResponseIterator(page)
}
func NewVirtualMachineListResponsePage(getNextPage func(context.Context, VirtualMachineListResponse) (VirtualMachineListResponse, error)) VirtualMachineListResponsePage {
	return original.NewVirtualMachineListResponsePage(getNextPage)
}
func NewVirtualMachineTemplateListResponseIterator(page VirtualMachineTemplateListResponsePage) VirtualMachineTemplateListResponseIterator {
	return original.NewVirtualMachineTemplateListResponseIterator(page)
}
func NewVirtualMachineTemplateListResponsePage(getNextPage func(context.Context, VirtualMachineTemplateListResponse) (VirtualMachineTemplateListResponse, error)) VirtualMachineTemplateListResponsePage {
	return original.NewVirtualMachineTemplateListResponsePage(getNextPage)
}
func NewVirtualMachineTemplatesClient(referer string, subscriptionID string) VirtualMachineTemplatesClient {
	return original.NewVirtualMachineTemplatesClient(referer, subscriptionID)
}
func NewVirtualMachineTemplatesClientWithBaseURI(baseURI string, referer string, subscriptionID string) VirtualMachineTemplatesClient {
	return original.NewVirtualMachineTemplatesClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewVirtualMachinesClient(referer string, subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClient(referer, subscriptionID)
}
func NewVirtualMachinesClientWithBaseURI(baseURI string, referer string, subscriptionID string) VirtualMachinesClient {
	return original.NewVirtualMachinesClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewVirtualNetworkListResponseIterator(page VirtualNetworkListResponsePage) VirtualNetworkListResponseIterator {
	return original.NewVirtualNetworkListResponseIterator(page)
}
func NewVirtualNetworkListResponsePage(getNextPage func(context.Context, VirtualNetworkListResponse) (VirtualNetworkListResponse, error)) VirtualNetworkListResponsePage {
	return original.NewVirtualNetworkListResponsePage(getNextPage)
}
func NewVirtualNetworksClient(referer string, subscriptionID string) VirtualNetworksClient {
	return original.NewVirtualNetworksClient(referer, subscriptionID)
}
func NewVirtualNetworksClientWithBaseURI(baseURI string, referer string, subscriptionID string) VirtualNetworksClient {
	return original.NewVirtualNetworksClientWithBaseURI(baseURI, referer, subscriptionID)
}
func NewWithBaseURI(baseURI string, referer string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, referer, subscriptionID)
}
func PossibleAggregationTypeValues() []AggregationType {
	return original.PossibleAggregationTypeValues()
}
func PossibleDiskIndependenceModeValues() []DiskIndependenceMode {
	return original.PossibleDiskIndependenceModeValues()
}
func PossibleGuestOSTypeValues() []GuestOSType {
	return original.PossibleGuestOSTypeValues()
}
func PossibleNICTypeValues() []NICType {
	return original.PossibleNICTypeValues()
}
func PossibleNodeStatusValues() []NodeStatus {
	return original.PossibleNodeStatusValues()
}
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return original.PossibleOnboardingStatusValues()
}
func PossibleOperationOriginValues() []OperationOrigin {
	return original.PossibleOperationOriginValues()
}
func PossiblePrivateCloudResourceTypeValues() []PrivateCloudResourceType {
	return original.PossiblePrivateCloudResourceTypeValues()
}
func PossibleStopModeValues() []StopMode {
	return original.PossibleStopModeValues()
}
func PossibleUsageCountValues() []UsageCount {
	return original.PossibleUsageCountValues()
}
func PossibleVirtualMachineStatusValues() []VirtualMachineStatus {
	return original.PossibleVirtualMachineStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
