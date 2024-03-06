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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
	"net/http"
	"net/url"
	"regexp"
)

// StorageAccountsServer is a fake server for instances of the armdataboxedge.StorageAccountsClient type.
type StorageAccountsServer struct {
	// BeginCreateOrUpdate is the fake for method StorageAccountsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, storageAccount armdataboxedge.StorageAccount, options *armdataboxedge.StorageAccountsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdataboxedge.StorageAccountsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method StorageAccountsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *armdataboxedge.StorageAccountsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdataboxedge.StorageAccountsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method StorageAccountsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, storageAccountName string, resourceGroupName string, options *armdataboxedge.StorageAccountsClientGetOptions) (resp azfake.Responder[armdataboxedge.StorageAccountsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDataBoxEdgeDevicePager is the fake for method StorageAccountsClient.NewListByDataBoxEdgeDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDataBoxEdgeDevicePager func(deviceName string, resourceGroupName string, options *armdataboxedge.StorageAccountsClientListByDataBoxEdgeDeviceOptions) (resp azfake.PagerResponder[armdataboxedge.StorageAccountsClientListByDataBoxEdgeDeviceResponse])
}

// NewStorageAccountsServerTransport creates a new instance of StorageAccountsServerTransport with the provided implementation.
// The returned StorageAccountsServerTransport instance is connected to an instance of armdataboxedge.StorageAccountsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewStorageAccountsServerTransport(srv *StorageAccountsServer) *StorageAccountsServerTransport {
	return &StorageAccountsServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armdataboxedge.StorageAccountsClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armdataboxedge.StorageAccountsClientDeleteResponse]](),
		newListByDataBoxEdgeDevicePager: newTracker[azfake.PagerResponder[armdataboxedge.StorageAccountsClientListByDataBoxEdgeDeviceResponse]](),
	}
}

// StorageAccountsServerTransport connects instances of armdataboxedge.StorageAccountsClient to instances of StorageAccountsServer.
// Don't use this type directly, use NewStorageAccountsServerTransport instead.
type StorageAccountsServerTransport struct {
	srv                             *StorageAccountsServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armdataboxedge.StorageAccountsClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armdataboxedge.StorageAccountsClientDeleteResponse]]
	newListByDataBoxEdgeDevicePager *tracker[azfake.PagerResponder[armdataboxedge.StorageAccountsClientListByDataBoxEdgeDeviceResponse]]
}

// Do implements the policy.Transporter interface for StorageAccountsServerTransport.
func (s *StorageAccountsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "StorageAccountsClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "StorageAccountsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "StorageAccountsClient.Get":
		resp, err = s.dispatchGet(req)
	case "StorageAccountsClient.NewListByDataBoxEdgeDevicePager":
		resp, err = s.dispatchNewListByDataBoxEdgeDevicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *StorageAccountsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageAccounts/(?P<storageAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.StorageAccount](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		storageAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageAccountName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, storageAccountNameParam, resourceGroupNameParam, body, nil)
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

func (s *StorageAccountsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageAccounts/(?P<storageAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		storageAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageAccountName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), deviceNameParam, storageAccountNameParam, resourceGroupNameParam, nil)
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

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *StorageAccountsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageAccounts/(?P<storageAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	storageAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageAccountName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), deviceNameParam, storageAccountNameParam, resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StorageAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StorageAccountsServerTransport) dispatchNewListByDataBoxEdgeDevicePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByDataBoxEdgeDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDataBoxEdgeDevicePager not implemented")}
	}
	newListByDataBoxEdgeDevicePager := s.newListByDataBoxEdgeDevicePager.get(req)
	if newListByDataBoxEdgeDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storageAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByDataBoxEdgeDevicePager(deviceNameParam, resourceGroupNameParam, nil)
		newListByDataBoxEdgeDevicePager = &resp
		s.newListByDataBoxEdgeDevicePager.add(req, newListByDataBoxEdgeDevicePager)
		server.PagerResponderInjectNextLinks(newListByDataBoxEdgeDevicePager, req, func(page *armdataboxedge.StorageAccountsClientListByDataBoxEdgeDeviceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDataBoxEdgeDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByDataBoxEdgeDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDataBoxEdgeDevicePager) {
		s.newListByDataBoxEdgeDevicePager.remove(req)
	}
	return resp, nil
}
