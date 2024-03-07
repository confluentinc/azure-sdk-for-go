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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CatalogDevBoxDefinitionsServer is a fake server for instances of the armdevcenter.CatalogDevBoxDefinitionsClient type.
type CatalogDevBoxDefinitionsServer struct {
	// Get is the fake for method CatalogDevBoxDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, devBoxDefinitionName string, options *armdevcenter.CatalogDevBoxDefinitionsClientGetOptions) (resp azfake.Responder[armdevcenter.CatalogDevBoxDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetErrorDetails is the fake for method CatalogDevBoxDefinitionsClient.GetErrorDetails
	// HTTP status codes to indicate success: http.StatusOK
	GetErrorDetails func(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, devBoxDefinitionName string, options *armdevcenter.CatalogDevBoxDefinitionsClientGetErrorDetailsOptions) (resp azfake.Responder[armdevcenter.CatalogDevBoxDefinitionsClientGetErrorDetailsResponse], errResp azfake.ErrorResponder)

	// NewListByCatalogPager is the fake for method CatalogDevBoxDefinitionsClient.NewListByCatalogPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCatalogPager func(resourceGroupName string, devCenterName string, catalogName string, options *armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogOptions) (resp azfake.PagerResponder[armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogResponse])
}

// NewCatalogDevBoxDefinitionsServerTransport creates a new instance of CatalogDevBoxDefinitionsServerTransport with the provided implementation.
// The returned CatalogDevBoxDefinitionsServerTransport instance is connected to an instance of armdevcenter.CatalogDevBoxDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCatalogDevBoxDefinitionsServerTransport(srv *CatalogDevBoxDefinitionsServer) *CatalogDevBoxDefinitionsServerTransport {
	return &CatalogDevBoxDefinitionsServerTransport{
		srv:                   srv,
		newListByCatalogPager: newTracker[azfake.PagerResponder[armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogResponse]](),
	}
}

// CatalogDevBoxDefinitionsServerTransport connects instances of armdevcenter.CatalogDevBoxDefinitionsClient to instances of CatalogDevBoxDefinitionsServer.
// Don't use this type directly, use NewCatalogDevBoxDefinitionsServerTransport instead.
type CatalogDevBoxDefinitionsServerTransport struct {
	srv                   *CatalogDevBoxDefinitionsServer
	newListByCatalogPager *tracker[azfake.PagerResponder[armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogResponse]]
}

// Do implements the policy.Transporter interface for CatalogDevBoxDefinitionsServerTransport.
func (c *CatalogDevBoxDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CatalogDevBoxDefinitionsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CatalogDevBoxDefinitionsClient.GetErrorDetails":
		resp, err = c.dispatchGetErrorDetails(req)
	case "CatalogDevBoxDefinitionsClient.NewListByCatalogPager":
		resp, err = c.dispatchNewListByCatalogPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CatalogDevBoxDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devboxdefinitions/(?P<devBoxDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	devBoxDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devBoxDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, devBoxDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DevBoxDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CatalogDevBoxDefinitionsServerTransport) dispatchGetErrorDetails(req *http.Request) (*http.Response, error) {
	if c.srv.GetErrorDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetErrorDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devboxdefinitions/(?P<devBoxDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getErrorDetails`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
	if err != nil {
		return nil, err
	}
	devBoxDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devBoxDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetErrorDetails(req.Context(), resourceGroupNameParam, devCenterNameParam, catalogNameParam, devBoxDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CatalogResourceValidationErrorDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CatalogDevBoxDefinitionsServerTransport) dispatchNewListByCatalogPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByCatalogPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCatalogPager not implemented")}
	}
	newListByCatalogPager := c.newListByCatalogPager.get(req)
	if newListByCatalogPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/catalogs/(?P<catalogName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devboxdefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		catalogNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("catalogName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogOptions
		if topParam != nil {
			options = &armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogOptions{
				Top: topParam,
			}
		}
		resp := c.srv.NewListByCatalogPager(resourceGroupNameParam, devCenterNameParam, catalogNameParam, options)
		newListByCatalogPager = &resp
		c.newListByCatalogPager.add(req, newListByCatalogPager)
		server.PagerResponderInjectNextLinks(newListByCatalogPager, req, func(page *armdevcenter.CatalogDevBoxDefinitionsClientListByCatalogResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCatalogPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByCatalogPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCatalogPager) {
		c.newListByCatalogPager.remove(req)
	}
	return resp, nil
}