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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
	"net/http"
	"net/url"
	"regexp"
)

// ServiceConfigurationsServer is a fake server for instances of the armhybridconnectivity.ServiceConfigurationsClient type.
type ServiceConfigurationsServer struct {
	// CreateOrupdate is the fake for method ServiceConfigurationsClient.CreateOrupdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrupdate func(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, serviceConfigurationResource armhybridconnectivity.ServiceConfigurationResource, options *armhybridconnectivity.ServiceConfigurationsClientCreateOrupdateOptions) (resp azfake.Responder[armhybridconnectivity.ServiceConfigurationsClientCreateOrupdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ServiceConfigurationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, options *armhybridconnectivity.ServiceConfigurationsClientDeleteOptions) (resp azfake.Responder[armhybridconnectivity.ServiceConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServiceConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, options *armhybridconnectivity.ServiceConfigurationsClientGetOptions) (resp azfake.Responder[armhybridconnectivity.ServiceConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByEndpointResourcePager is the fake for method ServiceConfigurationsClient.NewListByEndpointResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByEndpointResourcePager func(resourceURI string, endpointName string, options *armhybridconnectivity.ServiceConfigurationsClientListByEndpointResourceOptions) (resp azfake.PagerResponder[armhybridconnectivity.ServiceConfigurationsClientListByEndpointResourceResponse])

	// Update is the fake for method ServiceConfigurationsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, serviceConfigurationResource armhybridconnectivity.ServiceConfigurationResourcePatch, options *armhybridconnectivity.ServiceConfigurationsClientUpdateOptions) (resp azfake.Responder[armhybridconnectivity.ServiceConfigurationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewServiceConfigurationsServerTransport creates a new instance of ServiceConfigurationsServerTransport with the provided implementation.
// The returned ServiceConfigurationsServerTransport instance is connected to an instance of armhybridconnectivity.ServiceConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceConfigurationsServerTransport(srv *ServiceConfigurationsServer) *ServiceConfigurationsServerTransport {
	return &ServiceConfigurationsServerTransport{
		srv:                            srv,
		newListByEndpointResourcePager: newTracker[azfake.PagerResponder[armhybridconnectivity.ServiceConfigurationsClientListByEndpointResourceResponse]](),
	}
}

// ServiceConfigurationsServerTransport connects instances of armhybridconnectivity.ServiceConfigurationsClient to instances of ServiceConfigurationsServer.
// Don't use this type directly, use NewServiceConfigurationsServerTransport instead.
type ServiceConfigurationsServerTransport struct {
	srv                            *ServiceConfigurationsServer
	newListByEndpointResourcePager *tracker[azfake.PagerResponder[armhybridconnectivity.ServiceConfigurationsClientListByEndpointResourceResponse]]
}

// Do implements the policy.Transporter interface for ServiceConfigurationsServerTransport.
func (s *ServiceConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServiceConfigurationsClient.CreateOrupdate":
		resp, err = s.dispatchCreateOrupdate(req)
	case "ServiceConfigurationsClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "ServiceConfigurationsClient.Get":
		resp, err = s.dispatchGet(req)
	case "ServiceConfigurationsClient.NewListByEndpointResourcePager":
		resp, err = s.dispatchNewListByEndpointResourcePager(req)
	case "ServiceConfigurationsClient.Update":
		resp, err = s.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceConfigurationsServerTransport) dispatchCreateOrupdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrupdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrupdate not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceConfigurations/(?P<serviceConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridconnectivity.ServiceConfigurationResource](req)
	if err != nil {
		return nil, err
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	serviceConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrupdate(req.Context(), resourceURIParam, endpointNameParam, serviceConfigurationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceConfigurationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceConfigurationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceConfigurations/(?P<serviceConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	serviceConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceURIParam, endpointNameParam, serviceConfigurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceConfigurations/(?P<serviceConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	serviceConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceURIParam, endpointNameParam, serviceConfigurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceConfigurationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceConfigurationsServerTransport) dispatchNewListByEndpointResourcePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByEndpointResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByEndpointResourcePager not implemented")}
	}
	newListByEndpointResourcePager := s.newListByEndpointResourcePager.get(req)
	if newListByEndpointResourcePager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByEndpointResourcePager(resourceURIParam, endpointNameParam, nil)
		newListByEndpointResourcePager = &resp
		s.newListByEndpointResourcePager.add(req, newListByEndpointResourcePager)
		server.PagerResponderInjectNextLinks(newListByEndpointResourcePager, req, func(page *armhybridconnectivity.ServiceConfigurationsClientListByEndpointResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByEndpointResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByEndpointResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByEndpointResourcePager) {
		s.newListByEndpointResourcePager.remove(req)
	}
	return resp, nil
}

func (s *ServiceConfigurationsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridConnectivity/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceConfigurations/(?P<serviceConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridconnectivity.ServiceConfigurationResourcePatch](req)
	if err != nil {
		return nil, err
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	serviceConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Update(req.Context(), resourceURIParam, endpointNameParam, serviceConfigurationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServiceConfigurationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
