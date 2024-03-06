//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
	"net/http"
	"net/url"
	"regexp"
)

// ComponentCurrentBillingFeaturesServer is a fake server for instances of the armapplicationinsights.ComponentCurrentBillingFeaturesClient type.
type ComponentCurrentBillingFeaturesServer struct {
	// Get is the fake for method ComponentCurrentBillingFeaturesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armapplicationinsights.ComponentCurrentBillingFeaturesClientGetOptions) (resp azfake.Responder[armapplicationinsights.ComponentCurrentBillingFeaturesClientGetResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method ComponentCurrentBillingFeaturesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, resourceName string, billingFeaturesProperties armapplicationinsights.ComponentBillingFeatures, options *armapplicationinsights.ComponentCurrentBillingFeaturesClientUpdateOptions) (resp azfake.Responder[armapplicationinsights.ComponentCurrentBillingFeaturesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewComponentCurrentBillingFeaturesServerTransport creates a new instance of ComponentCurrentBillingFeaturesServerTransport with the provided implementation.
// The returned ComponentCurrentBillingFeaturesServerTransport instance is connected to an instance of armapplicationinsights.ComponentCurrentBillingFeaturesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewComponentCurrentBillingFeaturesServerTransport(srv *ComponentCurrentBillingFeaturesServer) *ComponentCurrentBillingFeaturesServerTransport {
	return &ComponentCurrentBillingFeaturesServerTransport{srv: srv}
}

// ComponentCurrentBillingFeaturesServerTransport connects instances of armapplicationinsights.ComponentCurrentBillingFeaturesClient to instances of ComponentCurrentBillingFeaturesServer.
// Don't use this type directly, use NewComponentCurrentBillingFeaturesServerTransport instead.
type ComponentCurrentBillingFeaturesServerTransport struct {
	srv *ComponentCurrentBillingFeaturesServer
}

// Do implements the policy.Transporter interface for ComponentCurrentBillingFeaturesServerTransport.
func (c *ComponentCurrentBillingFeaturesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ComponentCurrentBillingFeaturesClient.Get":
		resp, err = c.dispatchGet(req)
	case "ComponentCurrentBillingFeaturesClient.Update":
		resp, err = c.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ComponentCurrentBillingFeaturesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/currentbillingfeatures`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentBillingFeatures, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ComponentCurrentBillingFeaturesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/components/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/currentbillingfeatures`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapplicationinsights.ComponentBillingFeatures](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentBillingFeatures, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
