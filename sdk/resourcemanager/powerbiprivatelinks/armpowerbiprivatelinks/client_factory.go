//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPowerBIResourcesClient creates a new instance of PowerBIResourcesClient.
func (c *ClientFactory) NewPowerBIResourcesClient() *PowerBIResourcesClient {
	subClient, _ := NewPowerBIResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient.
func (c *ClientFactory) NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient {
	subClient, _ := NewPrivateEndpointConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient.
func (c *ClientFactory) NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient {
	subClient, _ := NewPrivateLinkResourcesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateLinkServiceResourceOperationResultsClient creates a new instance of PrivateLinkServiceResourceOperationResultsClient.
func (c *ClientFactory) NewPrivateLinkServiceResourceOperationResultsClient() *PrivateLinkServiceResourceOperationResultsClient {
	subClient, _ := NewPrivateLinkServiceResourceOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateLinkServicesClient creates a new instance of PrivateLinkServicesClient.
func (c *ClientFactory) NewPrivateLinkServicesClient() *PrivateLinkServicesClient {
	subClient, _ := NewPrivateLinkServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPrivateLinkServicesForPowerBIClient creates a new instance of PrivateLinkServicesForPowerBIClient.
func (c *ClientFactory) NewPrivateLinkServicesForPowerBIClient() *PrivateLinkServicesForPowerBIClient {
	subClient, _ := NewPrivateLinkServicesForPowerBIClient(c.subscriptionID, c.credential, c.options)
	return subClient
}