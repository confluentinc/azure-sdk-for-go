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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
	"net/http"
	"net/url"
	"regexp"
)

// Server is a fake server for instances of the armfeatures.Client type.
type Server struct {
	// Get is the fake for method Client.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceProviderNamespace string, featureName string, options *armfeatures.ClientGetOptions) (resp azfake.Responder[armfeatures.ClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method Client.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceProviderNamespace string, options *armfeatures.ClientListOptions) (resp azfake.PagerResponder[armfeatures.ClientListResponse])

	// NewListAllPager is the fake for method Client.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armfeatures.ClientListAllOptions) (resp azfake.PagerResponder[armfeatures.ClientListAllResponse])

	// Register is the fake for method Client.Register
	// HTTP status codes to indicate success: http.StatusOK
	Register func(ctx context.Context, resourceProviderNamespace string, featureName string, options *armfeatures.ClientRegisterOptions) (resp azfake.Responder[armfeatures.ClientRegisterResponse], errResp azfake.ErrorResponder)

	// Unregister is the fake for method Client.Unregister
	// HTTP status codes to indicate success: http.StatusOK
	Unregister func(ctx context.Context, resourceProviderNamespace string, featureName string, options *armfeatures.ClientUnregisterOptions) (resp azfake.Responder[armfeatures.ClientUnregisterResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armfeatures.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{
		srv:             srv,
		newListPager:    newTracker[azfake.PagerResponder[armfeatures.ClientListResponse]](),
		newListAllPager: newTracker[azfake.PagerResponder[armfeatures.ClientListAllResponse]](),
	}
}

// ServerTransport connects instances of armfeatures.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv             *Server
	newListPager    *tracker[azfake.PagerResponder[armfeatures.ClientListResponse]]
	newListAllPager *tracker[azfake.PagerResponder[armfeatures.ClientListAllResponse]]
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.Get":
		resp, err = s.dispatchGet(req)
	case "Client.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "Client.NewListAllPager":
		resp, err = s.dispatchNewListAllPager(req)
	case "Client.Register":
		resp, err = s.dispatchRegister(req)
	case "Client.Unregister":
		resp, err = s.dispatchUnregister(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Features/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features/(?P<featureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	featureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("featureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceProviderNamespaceParam, featureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FeatureResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Features/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceProviderNamespaceParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armfeatures.ClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := s.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Features/features`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListAllPager(nil)
		newListAllPager = &resp
		s.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armfeatures.ClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		s.newListAllPager.remove(req)
	}
	return resp, nil
}

func (s *ServerTransport) dispatchRegister(req *http.Request) (*http.Response, error) {
	if s.srv.Register == nil {
		return nil, &nonRetriableError{errors.New("fake for method Register not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Features/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features/(?P<featureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/register`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	featureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("featureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Register(req.Context(), resourceProviderNamespaceParam, featureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FeatureResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchUnregister(req *http.Request) (*http.Response, error) {
	if s.srv.Unregister == nil {
		return nil, &nonRetriableError{errors.New("fake for method Unregister not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Features/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/features/(?P<featureName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unregister`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	featureNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("featureName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Unregister(req.Context(), resourceProviderNamespaceParam, featureNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FeatureResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}