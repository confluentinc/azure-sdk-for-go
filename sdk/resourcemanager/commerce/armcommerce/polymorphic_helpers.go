//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcommerce

import "encoding/json"

func unmarshalOfferTermInfoAutoGeneratedClassification(rawMsg json.RawMessage) (OfferTermInfoAutoGeneratedClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OfferTermInfoAutoGeneratedClassification
	switch m["Name"] {
	case string(OfferTermInfoMonetaryCommitment):
		b = &MonetaryCommitment{}
	case string(OfferTermInfoMonetaryCredit):
		b = &MonetaryCredit{}
	case string(OfferTermInfoRecurringCharge):
		b = &RecurringCharge{}
	default:
		b = &OfferTermInfoAutoGenerated{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOfferTermInfoAutoGeneratedClassificationArray(rawMsg json.RawMessage) ([]OfferTermInfoAutoGeneratedClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]OfferTermInfoAutoGeneratedClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalOfferTermInfoAutoGeneratedClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}