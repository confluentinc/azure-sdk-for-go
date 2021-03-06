package marketplacemeteredbilling

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
    "github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/marketplacemeteredbilling/mgmt/2018-08-31/marketplacemeteredbilling"

            // AdditionalInfoProperties additionalInfoProperties definition
            type AdditionalInfoProperties struct {
            // AcceptedMessage - Details about the accepted message.
            AcceptedMessage *ResultProperties `json:"acceptedMessage,omitempty"`
            }

            // BasicErrorProperties basicErrorProperties definition
            type BasicErrorProperties struct {
            // Message - Error Message.
            Message *string `json:"message,omitempty"`
            // Target - API request target.
            Target *string `json:"target,omitempty"`
            // Details - More details about the error.
            Details *[]DetailsProperties `json:"details,omitempty"`
            // Code - Error code enum.
            Code *string `json:"code,omitempty"`
            }

            // BatchUsageEventRequest the batchUsageEvent endpoint allows you to emit usage events for a specific
            // purchased entity
            type BatchUsageEventRequest struct {
            Request *[]UsageEventRequestProperties `json:"request,omitempty"`
            }

            // BatchUsageEventRequestProperties batchUsageEvent definition
            type BatchUsageEventRequestProperties struct {
            Request *[]UsageEventRequestProperties `json:"request,omitempty"`
            }

            // BatchUsageEventResponse the batchUsageEventResponse definition
            type BatchUsageEventResponse struct {
            autorest.Response `json:"-"`
            // Result - List of result objects
            Result *[]ResultProperties `json:"result,omitempty"`
            // Count - Count of result objects
            Count *int32 `json:"count,omitempty"`
            }

            // BatchUsageEventResponseProperties batchUsageEventResponseProperties definition
            type BatchUsageEventResponseProperties struct {
            // Result - List of result objects
            Result *[]ResultProperties `json:"result,omitempty"`
            // Count - Count of result objects
            Count *int32 `json:"count,omitempty"`
            }

            // DetailsProperties detailsProperties definition
            type DetailsProperties struct {
            // Message - Error Message.
            Message *string `json:"message,omitempty"`
            // Target - API request target.
            Target *string `json:"target,omitempty"`
            // Code - Error code enum.
            Code *string `json:"code,omitempty"`
            }

            // ErrorProperties errorProperties definition
            type ErrorProperties struct {
            // Message - Error Message.
            Message *string `json:"message,omitempty"`
            // Target - API request target.
            Target *string `json:"target,omitempty"`
            // Details - More details about the error.
            Details *[]DetailsProperties `json:"details,omitempty"`
            // Code - Error code enum.
            Code *string `json:"code,omitempty"`
            // AdditionalInfo - Additional information about the error.
            AdditionalInfo *AdditionalInfoProperties `json:"additionalInfo,omitempty"`
            }

            // ResultProperties resultProperties definition
            type ResultProperties struct {
            // UsageEventID - Unique ID for the reported usage event.
            UsageEventID *string `json:"usageEventId,omitempty"`
            // Status - Status of the result.
            Status *string `json:"status,omitempty"`
            // MessageTime - Time when the message was sent.
            MessageTime *string `json:"messageTime,omitempty"`
            // ResourceID - Resource ID against which the metric was reported.
            ResourceID *string `json:"resourceId,omitempty"`
            // Quantity - Quantity of the metric that was reported.
            Quantity *float64 `json:"quantity,omitempty"`
            // Dimension - The metric id itself.
            Dimension *string `json:"dimension,omitempty"`
            // EffectiveStartTime - The start time of the metric that was reported.
            EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
            // PlanID - Plan id of the subscriber.
            PlanID *string `json:"planId,omitempty"`
            // Error - Tenant that purchased the SaaS subscription.
            Error *ErrorProperties `json:"error,omitempty"`
            }

            // SetObject ...
            type SetObject struct {
            autorest.Response `json:"-"`
            Value interface{} `json:"value,omitempty"`
            }

            // UsageEventRequest the UsageEvent endpoint allows you to emit usage events for a specific purchased
            // entity
            type UsageEventRequest struct {
            // ResourceID - Identifier of the resource against which usage is emitted.
            ResourceID *string `json:"resourceId,omitempty"`
            // Quantity - Quantity of the dimension being reported.
            Quantity *float64 `json:"quantity,omitempty"`
            // Dimension - Dimension identifier.
            Dimension *string `json:"dimension,omitempty"`
            // EffectiveStartTime - Time in UTC when the usage event occurred.
            EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
            // PlanID - Plan associated with the purchased offer.
            PlanID *string `json:"planId,omitempty"`
            }

            // UsageEventRequestProperties usageEvent definition
            type UsageEventRequestProperties struct {
            // ResourceID - Identifier of the resource against which usage is emitted.
            ResourceID *string `json:"resourceId,omitempty"`
            // Quantity - Quantity of the dimension being reported.
            Quantity *float64 `json:"quantity,omitempty"`
            // Dimension - Dimension identifier.
            Dimension *string `json:"dimension,omitempty"`
            // EffectiveStartTime - Time in UTC when the usage event occurred.
            EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
            // PlanID - Plan associated with the purchased offer.
            PlanID *string `json:"planId,omitempty"`
            }

            // UsageEventResponse the batchUsageEventResponse definition
            type UsageEventResponse struct {
            // UsageEventID - Unique ID for the reported usage event.
            UsageEventID *string `json:"usageEventId,omitempty"`
            // Status - Status of the result.
            Status *string `json:"status,omitempty"`
            // MessageTime - Time when the message was sent.
            MessageTime *string `json:"messageTime,omitempty"`
            // ResourceID - Resource ID against which the metric was reported.
            ResourceID *string `json:"resourceId,omitempty"`
            // Quantity - Quantity of the metric that was reported.
            Quantity *float64 `json:"quantity,omitempty"`
            // Dimension - The metric id itself.
            Dimension *string `json:"dimension,omitempty"`
            // EffectiveStartTime - The start time of the metric that was reported.
            EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
            // PlanID - Plan id of the subscriber.
            PlanID *string `json:"planId,omitempty"`
            // Error - Tenant that purchased the SaaS subscription.
            Error *ErrorProperties `json:"error,omitempty"`
            }

