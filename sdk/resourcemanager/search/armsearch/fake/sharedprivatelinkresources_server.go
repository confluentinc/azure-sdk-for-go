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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"net/http"
	"net/url"
	"regexp"
)

// SharedPrivateLinkResourcesServer is a fake server for instances of the armsearch.SharedPrivateLinkResourcesClient type.
type SharedPrivateLinkResourcesServer struct {
	// BeginCreateOrUpdate is the fake for method SharedPrivateLinkResourcesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, sharedPrivateLinkResource armsearch.SharedPrivateLinkResource, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsearch.SharedPrivateLinkResourcesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SharedPrivateLinkResourcesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent, http.StatusNotFound
	BeginDelete func(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.SharedPrivateLinkResourcesClientBeginDeleteOptions) (resp azfake.PollerResponder[armsearch.SharedPrivateLinkResourcesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SharedPrivateLinkResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.SharedPrivateLinkResourcesClientGetOptions) (resp azfake.Responder[armsearch.SharedPrivateLinkResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method SharedPrivateLinkResourcesClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, searchServiceName string, searchManagementRequestOptions *armsearch.SearchManagementRequestOptions, options *armsearch.SharedPrivateLinkResourcesClientListByServiceOptions) (resp azfake.PagerResponder[armsearch.SharedPrivateLinkResourcesClientListByServiceResponse])
}

// NewSharedPrivateLinkResourcesServerTransport creates a new instance of SharedPrivateLinkResourcesServerTransport with the provided implementation.
// The returned SharedPrivateLinkResourcesServerTransport instance is connected to an instance of armsearch.SharedPrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSharedPrivateLinkResourcesServerTransport(srv *SharedPrivateLinkResourcesServer) *SharedPrivateLinkResourcesServerTransport {
	return &SharedPrivateLinkResourcesServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armsearch.SharedPrivateLinkResourcesClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armsearch.SharedPrivateLinkResourcesClientDeleteResponse]](),
		newListByServicePager: newTracker[azfake.PagerResponder[armsearch.SharedPrivateLinkResourcesClientListByServiceResponse]](),
	}
}

// SharedPrivateLinkResourcesServerTransport connects instances of armsearch.SharedPrivateLinkResourcesClient to instances of SharedPrivateLinkResourcesServer.
// Don't use this type directly, use NewSharedPrivateLinkResourcesServerTransport instead.
type SharedPrivateLinkResourcesServerTransport struct {
	srv                   *SharedPrivateLinkResourcesServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armsearch.SharedPrivateLinkResourcesClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armsearch.SharedPrivateLinkResourcesClientDeleteResponse]]
	newListByServicePager *tracker[azfake.PagerResponder[armsearch.SharedPrivateLinkResourcesClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for SharedPrivateLinkResourcesServerTransport.
func (s *SharedPrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SharedPrivateLinkResourcesClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SharedPrivateLinkResourcesClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SharedPrivateLinkResourcesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SharedPrivateLinkResourcesClient.NewListByServicePager":
		resp, err = s.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsearch.SharedPrivateLinkResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
		if err != nil {
			return nil, err
		}
		sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
		if err != nil {
			return nil, err
		}
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
		var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
		if clientRequestIDParam != nil {
			searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
				ClientRequestID: clientRequestIDParam,
			}
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, searchServiceNameParam, sharedPrivateLinkResourceNameParam, body, searchManagementRequestOptions, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
		if err != nil {
			return nil, err
		}
		sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
		if err != nil {
			return nil, err
		}
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
		var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
		if clientRequestIDParam != nil {
			searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
				ClientRequestID: clientRequestIDParam,
			}
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, searchServiceNameParam, sharedPrivateLinkResourceNameParam, searchManagementRequestOptions, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent, http.StatusNotFound}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent, http.StatusNotFound", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
	if err != nil {
		return nil, err
	}
	sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
	var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
	if clientRequestIDParam != nil {
		searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, searchServiceNameParam, sharedPrivateLinkResourceNameParam, searchManagementRequestOptions, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SharedPrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SharedPrivateLinkResourcesServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := s.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "x-ms-client-request-id"))
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
		if err != nil {
			return nil, err
		}
		var searchManagementRequestOptions *armsearch.SearchManagementRequestOptions
		if clientRequestIDParam != nil {
			searchManagementRequestOptions = &armsearch.SearchManagementRequestOptions{
				ClientRequestID: clientRequestIDParam,
			}
		}
		resp := s.srv.NewListByServicePager(resourceGroupNameParam, searchServiceNameParam, searchManagementRequestOptions, nil)
		newListByServicePager = &resp
		s.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armsearch.SharedPrivateLinkResourcesClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		s.newListByServicePager.remove(req)
	}
	return resp, nil
}
