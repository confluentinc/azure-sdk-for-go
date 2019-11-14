package marketplacefulfillmentapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "context"
    "github.com/Azure/azure-sdk-for-go/services/marketplacefulfillment/mgmt/2018-08-31/marketplacefulfillment"
)

        // ClientAPI contains the set of methods on the Client type.
        type ClientAPI interface {
            ActivateSubscription(ctx context.Context, activatePayload marketplacefulfillment.Subscription, subscriptionID string, authorization string, xMsRequestid string, xMsCorrelationid string) (result marketplacefulfillment.Subscription, err error)
            GetOperation(ctx context.Context, subscriptionID string, operationID string, authorization string, xMsRequestid string, xMsCorrelationid string) (result marketplacefulfillment.Operation, err error)
            GetSubscription(ctx context.Context, subscriptionID string, authorization string, xMsRequestid string, xMsCorrelationid string) (result marketplacefulfillment.Subscription, err error)
            PatchOperation(ctx context.Context, patchPayload marketplacefulfillment.Operation, subscriptionID string, operationID string, authorization string, xMsRequestid string, xMsCorrelationid string) (result marketplacefulfillment.Operation, err error)
            Resolve(ctx context.Context, authorization string, xMsMarketplaceToken string, xMsRequestid string, xMsCorrelationid string) (result marketplacefulfillment.Subscription, err error)
        }

        var _ ClientAPI = (*marketplacefulfillment.Client)(nil)
