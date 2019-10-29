package marketplacefulfillment

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
    "encoding/json"
    "github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/marketplacefulfillment/mgmt/2018-08-31/marketplacefulfillment"

            // ErrorResponse error response indicates Microsoft.MarketplaceFulfillment service is not able to process
            // the incoming request. The reason is provided in the error message.
            type ErrorResponse struct {
            // Error - The details of the error.
            Error *ErrorResponseError `json:"error,omitempty"`
            }

            // ErrorResponseError the details of the error.
            type ErrorResponseError struct {
            // Code - READ-ONLY; Error code.
            Code *string `json:"code,omitempty"`
            // Message - READ-ONLY; Error message indicating why the operation failed.
            Message *string `json:"message,omitempty"`
            }

            // Subscription subscription Publisher/Offer/Plan tuple with quantity
            type Subscription struct {
            autorest.Response `json:"-"`
            // SubscriptionProperties - Represents the properties of the resource.
            *SubscriptionProperties `json:"properties,omitempty"`
            // ID - GUID.
            ID *string `json:"id,omitempty"`
            // Name - Subscription name.
            Name *string `json:"name,omitempty"`
            // PublisherID - Publisher Id.
            PublisherID *string `json:"publisherId,omitempty"`
            // OfferID - Offer Id.
            OfferID *string `json:"offerId,omitempty"`
            // PlanID - Plan Id.
            PlanID *string `json:"planId,omitempty"`
            // Quantity - Quantity.
            Quantity *string `json:"quantity,omitempty"`
            }

        // MarshalJSON is the custom marshaler for Subscription.
        func (s Subscription)MarshalJSON() ([]byte, error){
        objectMap := make(map[string]interface{})
                if(s.SubscriptionProperties != nil) {
                objectMap["properties"] = s.SubscriptionProperties
                }
                if(s.ID != nil) {
                objectMap["id"] = s.ID
                }
                if(s.Name != nil) {
                objectMap["name"] = s.Name
                }
                if(s.PublisherID != nil) {
                objectMap["publisherId"] = s.PublisherID
                }
                if(s.OfferID != nil) {
                objectMap["offerId"] = s.OfferID
                }
                if(s.PlanID != nil) {
                objectMap["planId"] = s.PlanID
                }
                if(s.Quantity != nil) {
                objectMap["quantity"] = s.Quantity
                }
                return json.Marshal(objectMap)
        }
        // UnmarshalJSON is the custom unmarshaler for Subscription struct.
        func (s *Subscription) UnmarshalJSON(body []byte) error {
        var m map[string]*json.RawMessage
        err := json.Unmarshal(body, &m)
        if err != nil {
        return err
        }
        for k, v := range  m {
        switch k {
                case "properties":
    if v != nil {
        var subscriptionProperties SubscriptionProperties
        err = json.Unmarshal(*v, &subscriptionProperties)
    if err != nil {
    return err
    }
        s.SubscriptionProperties = &subscriptionProperties
    }
                case "id":
    if v != nil {
        var ID string
        err = json.Unmarshal(*v, &ID)
    if err != nil {
    return err
    }
        s.ID = &ID
    }
                case "name":
    if v != nil {
        var name string
        err = json.Unmarshal(*v, &name)
    if err != nil {
    return err
    }
        s.Name = &name
    }
                case "publisherId":
    if v != nil {
        var publisherID string
        err = json.Unmarshal(*v, &publisherID)
    if err != nil {
    return err
    }
        s.PublisherID = &publisherID
    }
                case "offerId":
    if v != nil {
        var offerID string
        err = json.Unmarshal(*v, &offerID)
    if err != nil {
    return err
    }
        s.OfferID = &offerID
    }
                case "planId":
    if v != nil {
        var planID string
        err = json.Unmarshal(*v, &planID)
    if err != nil {
    return err
    }
        s.PlanID = &planID
    }
                case "quantity":
    if v != nil {
        var quantity string
        err = json.Unmarshal(*v, &quantity)
    if err != nil {
    return err
    }
        s.Quantity = &quantity
    }
            }
        }

        return nil
        }

            // SubscriptionProperties subscription definition
            type SubscriptionProperties struct {
            // ID - GUID.
            ID *string `json:"id,omitempty"`
            // Name - Subscription name.
            Name *string `json:"name,omitempty"`
            // PublisherID - Publisher Id.
            PublisherID *string `json:"publisherId,omitempty"`
            // OfferID - Offer Id.
            OfferID *string `json:"offerId,omitempty"`
            // PlanID - Plan Id.
            PlanID *string `json:"planId,omitempty"`
            // Quantity - Quantity.
            Quantity *string `json:"quantity,omitempty"`
            }
