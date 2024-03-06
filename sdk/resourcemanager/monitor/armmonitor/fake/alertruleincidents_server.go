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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"regexp"
)

// AlertRuleIncidentsServer is a fake server for instances of the armmonitor.AlertRuleIncidentsClient type.
type AlertRuleIncidentsServer struct {
	// Get is the fake for method AlertRuleIncidentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, ruleName string, incidentName string, options *armmonitor.AlertRuleIncidentsClientGetOptions) (resp azfake.Responder[armmonitor.AlertRuleIncidentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAlertRulePager is the fake for method AlertRuleIncidentsClient.NewListByAlertRulePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAlertRulePager func(resourceGroupName string, ruleName string, options *armmonitor.AlertRuleIncidentsClientListByAlertRuleOptions) (resp azfake.PagerResponder[armmonitor.AlertRuleIncidentsClientListByAlertRuleResponse])
}

// NewAlertRuleIncidentsServerTransport creates a new instance of AlertRuleIncidentsServerTransport with the provided implementation.
// The returned AlertRuleIncidentsServerTransport instance is connected to an instance of armmonitor.AlertRuleIncidentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAlertRuleIncidentsServerTransport(srv *AlertRuleIncidentsServer) *AlertRuleIncidentsServerTransport {
	return &AlertRuleIncidentsServerTransport{
		srv:                     srv,
		newListByAlertRulePager: newTracker[azfake.PagerResponder[armmonitor.AlertRuleIncidentsClientListByAlertRuleResponse]](),
	}
}

// AlertRuleIncidentsServerTransport connects instances of armmonitor.AlertRuleIncidentsClient to instances of AlertRuleIncidentsServer.
// Don't use this type directly, use NewAlertRuleIncidentsServerTransport instead.
type AlertRuleIncidentsServerTransport struct {
	srv                     *AlertRuleIncidentsServer
	newListByAlertRulePager *tracker[azfake.PagerResponder[armmonitor.AlertRuleIncidentsClientListByAlertRuleResponse]]
}

// Do implements the policy.Transporter interface for AlertRuleIncidentsServerTransport.
func (a *AlertRuleIncidentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AlertRuleIncidentsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AlertRuleIncidentsClient.NewListByAlertRulePager":
		resp, err = a.dispatchNewListByAlertRulePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AlertRuleIncidentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft\.insights/alertrules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/incidents/(?P<incidentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
	if err != nil {
		return nil, err
	}
	incidentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, ruleNameParam, incidentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Incident, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AlertRuleIncidentsServerTransport) dispatchNewListByAlertRulePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByAlertRulePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAlertRulePager not implemented")}
	}
	newListByAlertRulePager := a.newListByAlertRulePager.get(req)
	if newListByAlertRulePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/microsoft\.insights/alertrules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/incidents`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByAlertRulePager(resourceGroupNameParam, ruleNameParam, nil)
		newListByAlertRulePager = &resp
		a.newListByAlertRulePager.add(req, newListByAlertRulePager)
	}
	resp, err := server.PagerResponderNext(newListByAlertRulePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByAlertRulePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAlertRulePager) {
		a.newListByAlertRulePager.remove(req)
	}
	return resp, nil
}
