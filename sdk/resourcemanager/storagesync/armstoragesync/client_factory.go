//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragesync

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
//   - subscriptionID - The ID of the target subscription.
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

// NewCloudEndpointsClient creates a new instance of CloudEndpointsClient.
func (c *ClientFactory) NewCloudEndpointsClient() *CloudEndpointsClient {
	subClient, _ := NewCloudEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMicrosoftStorageSyncClient creates a new instance of MicrosoftStorageSyncClient.
func (c *ClientFactory) NewMicrosoftStorageSyncClient() *MicrosoftStorageSyncClient {
	subClient, _ := NewMicrosoftStorageSyncClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationStatusClient creates a new instance of OperationStatusClient.
func (c *ClientFactory) NewOperationStatusClient() *OperationStatusClient {
	subClient, _ := NewOperationStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
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

// NewRegisteredServersClient creates a new instance of RegisteredServersClient.
func (c *ClientFactory) NewRegisteredServersClient() *RegisteredServersClient {
	subClient, _ := NewRegisteredServersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServerEndpointsClient creates a new instance of ServerEndpointsClient.
func (c *ClientFactory) NewServerEndpointsClient() *ServerEndpointsClient {
	subClient, _ := NewServerEndpointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServicesClient creates a new instance of ServicesClient.
func (c *ClientFactory) NewServicesClient() *ServicesClient {
	subClient, _ := NewServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewSyncGroupsClient creates a new instance of SyncGroupsClient.
func (c *ClientFactory) NewSyncGroupsClient() *SyncGroupsClient {
	subClient, _ := NewSyncGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewWorkflowsClient creates a new instance of WorkflowsClient.
func (c *ClientFactory) NewWorkflowsClient() *WorkflowsClient {
	subClient, _ := NewWorkflowsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}