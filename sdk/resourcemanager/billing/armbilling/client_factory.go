//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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
//   - subscriptionID - The ID that uniquely identifies an Azure subscription.
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

// NewAccountsClient creates a new instance of AccountsClient.
func (c *ClientFactory) NewAccountsClient() *AccountsClient {
	subClient, _ := NewAccountsClient(c.credential, c.options)
	return subClient
}

// NewAddressClient creates a new instance of AddressClient.
func (c *ClientFactory) NewAddressClient() *AddressClient {
	subClient, _ := NewAddressClient(c.credential, c.options)
	return subClient
}

// NewAgreementsClient creates a new instance of AgreementsClient.
func (c *ClientFactory) NewAgreementsClient() *AgreementsClient {
	subClient, _ := NewAgreementsClient(c.credential, c.options)
	return subClient
}

// NewAvailableBalancesClient creates a new instance of AvailableBalancesClient.
func (c *ClientFactory) NewAvailableBalancesClient() *AvailableBalancesClient {
	subClient, _ := NewAvailableBalancesClient(c.credential, c.options)
	return subClient
}

// NewCustomersClient creates a new instance of CustomersClient.
func (c *ClientFactory) NewCustomersClient() *CustomersClient {
	subClient, _ := NewCustomersClient(c.credential, c.options)
	return subClient
}

// NewEnrollmentAccountsClient creates a new instance of EnrollmentAccountsClient.
func (c *ClientFactory) NewEnrollmentAccountsClient() *EnrollmentAccountsClient {
	subClient, _ := NewEnrollmentAccountsClient(c.credential, c.options)
	return subClient
}

// NewInstructionsClient creates a new instance of InstructionsClient.
func (c *ClientFactory) NewInstructionsClient() *InstructionsClient {
	subClient, _ := NewInstructionsClient(c.credential, c.options)
	return subClient
}

// NewInvoiceSectionsClient creates a new instance of InvoiceSectionsClient.
func (c *ClientFactory) NewInvoiceSectionsClient() *InvoiceSectionsClient {
	subClient, _ := NewInvoiceSectionsClient(c.credential, c.options)
	return subClient
}

// NewInvoicesClient creates a new instance of InvoicesClient.
func (c *ClientFactory) NewInvoicesClient() *InvoicesClient {
	subClient, _ := NewInvoicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewPeriodsClient creates a new instance of PeriodsClient.
func (c *ClientFactory) NewPeriodsClient() *PeriodsClient {
	subClient, _ := NewPeriodsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewPermissionsClient creates a new instance of PermissionsClient.
func (c *ClientFactory) NewPermissionsClient() *PermissionsClient {
	subClient, _ := NewPermissionsClient(c.credential, c.options)
	return subClient
}

// NewPoliciesClient creates a new instance of PoliciesClient.
func (c *ClientFactory) NewPoliciesClient() *PoliciesClient {
	subClient, _ := NewPoliciesClient(c.credential, c.options)
	return subClient
}

// NewProductsClient creates a new instance of ProductsClient.
func (c *ClientFactory) NewProductsClient() *ProductsClient {
	subClient, _ := NewProductsClient(c.credential, c.options)
	return subClient
}

// NewProfilesClient creates a new instance of ProfilesClient.
func (c *ClientFactory) NewProfilesClient() *ProfilesClient {
	subClient, _ := NewProfilesClient(c.credential, c.options)
	return subClient
}

// NewPropertyClient creates a new instance of PropertyClient.
func (c *ClientFactory) NewPropertyClient() *PropertyClient {
	subClient, _ := NewPropertyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewReservationsClient creates a new instance of ReservationsClient.
func (c *ClientFactory) NewReservationsClient() *ReservationsClient {
	subClient, _ := NewReservationsClient(c.credential, c.options)
	return subClient
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient.
func (c *ClientFactory) NewRoleAssignmentsClient() *RoleAssignmentsClient {
	subClient, _ := NewRoleAssignmentsClient(c.credential, c.options)
	return subClient
}

// NewRoleDefinitionsClient creates a new instance of RoleDefinitionsClient.
func (c *ClientFactory) NewRoleDefinitionsClient() *RoleDefinitionsClient {
	subClient, _ := NewRoleDefinitionsClient(c.credential, c.options)
	return subClient
}

// NewSubscriptionsClient creates a new instance of SubscriptionsClient.
func (c *ClientFactory) NewSubscriptionsClient() *SubscriptionsClient {
	subClient, _ := NewSubscriptionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewTransactionsClient creates a new instance of TransactionsClient.
func (c *ClientFactory) NewTransactionsClient() *TransactionsClient {
	subClient, _ := NewTransactionsClient(c.credential, c.options)
	return subClient
}