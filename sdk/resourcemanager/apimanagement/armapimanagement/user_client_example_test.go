//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListUsers.json
func ExampleUserClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUserClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.UserClientListByServiceOptions{Filter: nil,
		Top:          nil,
		Skip:         nil,
		ExpandGroups: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.UserCollection = armapimanagement.UserCollection{
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.UserContract{
		// 		{
		// 			Name: to.Ptr("1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/1"),
		// 			Properties: &armapimanagement.UserContractProperties{
		// 				Identities: []*armapimanagement.UserIdentityContract{
		// 					{
		// 						ID: to.Ptr("admin@live.com"),
		// 						Provider: to.Ptr("Azure"),
		// 				}},
		// 				State: to.Ptr(armapimanagement.UserStateActive),
		// 				Email: to.Ptr("admin@live.com"),
		// 				FirstName: to.Ptr("Administrator"),
		// 				LastName: to.Ptr(""),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-22T01:57:39.677Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("56eaec62baf08b06e46d27fd"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/56eaec62baf08b06e46d27fd"),
		// 			Properties: &armapimanagement.UserContractProperties{
		// 				Identities: []*armapimanagement.UserIdentityContract{
		// 					{
		// 						ID: to.Ptr("foo.bar.83@gmail.com"),
		// 						Provider: to.Ptr("Basic"),
		// 				}},
		// 				State: to.Ptr(armapimanagement.UserStateActive),
		// 				Email: to.Ptr("foo.bar.83@gmail.com"),
		// 				FirstName: to.Ptr("foo"),
		// 				LastName: to.Ptr("bar"),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-03-17T17:41:56.327Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("5931a75ae4bbd512a88c680b"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/users"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
		// 			Properties: &armapimanagement.UserContractProperties{
		// 				Identities: []*armapimanagement.UserIdentityContract{
		// 					{
		// 						ID: to.Ptr("*************"),
		// 						Provider: to.Ptr("Microsoft"),
		// 				}},
		// 				State: to.Ptr(armapimanagement.UserStateActive),
		// 				Email: to.Ptr("foobar@outlook.com"),
		// 				FirstName: to.Ptr("foo"),
		// 				LastName: to.Ptr("bar"),
		// 				RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:58:50.357Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadUser.json
func ExampleUserClient_GetEntityTag() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewUserClient().GetEntityTag(ctx, "rg1", "apimService1", "5931a75ae4bbd512a88c680b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetUser.json
func ExampleUserClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().Get(ctx, "rg1", "apimService1", "5931a75ae4bbd512a88c680b", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserContract = armapimanagement.UserContract{
	// 	Name: to.Ptr("5931a75ae4bbd512a88c680b"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/users"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
	// 	Properties: &armapimanagement.UserContractProperties{
	// 		Identities: []*armapimanagement.UserIdentityContract{
	// 			{
	// 				ID: to.Ptr("*************"),
	// 				Provider: to.Ptr("Microsoft"),
	// 		}},
	// 		State: to.Ptr(armapimanagement.UserStateActive),
	// 		Email: to.Ptr("foobar@outlook.com"),
	// 		FirstName: to.Ptr("foo"),
	// 		LastName: to.Ptr("bar"),
	// 		RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:58:50.357Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateUser.json
func ExampleUserClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().CreateOrUpdate(ctx, "rg1", "apimService1", "5931a75ae4bbd512288c680b", armapimanagement.UserCreateParameters{
		Properties: &armapimanagement.UserCreateParameterProperties{
			Confirmation: to.Ptr(armapimanagement.ConfirmationSignup),
			Email:        to.Ptr("foobar@outlook.com"),
			FirstName:    to.Ptr("foo"),
			LastName:     to.Ptr("bar"),
		},
	}, &armapimanagement.UserClientCreateOrUpdateOptions{Notify: nil,
		IfMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserContract = armapimanagement.UserContract{
	// 	Name: to.Ptr("5931a75ae4bbd512288c680b"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/users"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512288c680b"),
	// 	Properties: &armapimanagement.UserContractProperties{
	// 		Identities: []*armapimanagement.UserIdentityContract{
	// 			{
	// 				ID: to.Ptr("foobar@outlook.com"),
	// 				Provider: to.Ptr("Basic"),
	// 		}},
	// 		State: to.Ptr(armapimanagement.UserStateActive),
	// 		Email: to.Ptr("foobar@outlook.com"),
	// 		FirstName: to.Ptr("foo"),
	// 		Groups: []*armapimanagement.GroupContractProperties{
	// 		},
	// 		LastName: to.Ptr("bar"),
	// 		RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-07T21:21:29.160Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateUser.json
func ExampleUserClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().Update(ctx, "rg1", "apimService1", "5931a75ae4bbd512a88c680b", "*", armapimanagement.UserUpdateParameters{
		Properties: &armapimanagement.UserUpdateParametersProperties{
			Email:     to.Ptr("foobar@outlook.com"),
			FirstName: to.Ptr("foo"),
			LastName:  to.Ptr("bar"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserContract = armapimanagement.UserContract{
	// 	Name: to.Ptr("5931a75ae4bbd512a88c680b"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/users"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
	// 	Properties: &armapimanagement.UserContractProperties{
	// 		Identities: []*armapimanagement.UserIdentityContract{
	// 			{
	// 				ID: to.Ptr("*************"),
	// 				Provider: to.Ptr("Microsoft"),
	// 		}},
	// 		State: to.Ptr(armapimanagement.UserStateActive),
	// 		Email: to.Ptr("foobar@outlook.com"),
	// 		FirstName: to.Ptr("foo"),
	// 		LastName: to.Ptr("bar"),
	// 		RegistrationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:58:50.357Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteUser.json
func ExampleUserClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewUserClient().Delete(ctx, "rg1", "apimService1", "5931a75ae4bbd512288c680b", "*", &armapimanagement.UserClientDeleteOptions{DeleteSubscriptions: nil,
		Notify:  nil,
		AppType: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUserGenerateSsoUrl.json
func ExampleUserClient_GenerateSsoURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().GenerateSsoURL(ctx, "rg1", "apimService1", "57127d485157a511ace86ae7", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GenerateSsoURLResult = armapimanagement.GenerateSsoURLResult{
	// 	Value: to.Ptr("https://apimService1.portal.azure-api.net/signin-sso?token=57127d485157a511ace86ae7%26201706051624%267VY18MlwAom***********2bYr2bDQHg21OzQsNakExQ%3d%3d"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUserToken.json
func ExampleUserClient_GetSharedAccessToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewUserClient().GetSharedAccessToken(ctx, "rg1", "apimService1", "userId1718", armapimanagement.UserTokenParameters{
		Properties: &armapimanagement.UserTokenParameterProperties{
			Expiry:  to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T00:44:24.284Z"); return t }()),
			KeyType: to.Ptr(armapimanagement.KeyTypePrimary),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.UserTokenResult = armapimanagement.UserTokenResult{
	// 	Value: to.Ptr("userId1718&201904210044&9A1GR1f5WIhFvFmzQG+xxxxxxxxxxx/kBeu87DWad3tkasUXuvPL+MgzlwUHyg=="),
	// }
}