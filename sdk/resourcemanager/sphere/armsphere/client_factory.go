//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsphere

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

// NewCatalogsClient creates a new instance of CatalogsClient.
func (c *ClientFactory) NewCatalogsClient() *CatalogsClient {
	subClient, _ := NewCatalogsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCertificatesClient creates a new instance of CertificatesClient.
func (c *ClientFactory) NewCertificatesClient() *CertificatesClient {
	subClient, _ := NewCertificatesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDeploymentsClient creates a new instance of DeploymentsClient.
func (c *ClientFactory) NewDeploymentsClient() *DeploymentsClient {
	subClient, _ := NewDeploymentsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDeviceGroupsClient creates a new instance of DeviceGroupsClient.
func (c *ClientFactory) NewDeviceGroupsClient() *DeviceGroupsClient {
	subClient, _ := NewDeviceGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDevicesClient creates a new instance of DevicesClient.
func (c *ClientFactory) NewDevicesClient() *DevicesClient {
	subClient, _ := NewDevicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewImagesClient creates a new instance of ImagesClient.
func (c *ClientFactory) NewImagesClient() *ImagesClient {
	subClient, _ := NewImagesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewProductsClient creates a new instance of ProductsClient.
func (c *ClientFactory) NewProductsClient() *ProductsClient {
	subClient, _ := NewProductsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}