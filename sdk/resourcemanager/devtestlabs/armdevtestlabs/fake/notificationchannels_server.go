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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// NotificationChannelsServer is a fake server for instances of the armdevtestlabs.NotificationChannelsClient type.
type NotificationChannelsServer struct {
	// CreateOrUpdate is the fake for method NotificationChannelsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel armdevtestlabs.NotificationChannel, options *armdevtestlabs.NotificationChannelsClientCreateOrUpdateOptions) (resp azfake.Responder[armdevtestlabs.NotificationChannelsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method NotificationChannelsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, labName string, name string, options *armdevtestlabs.NotificationChannelsClientDeleteOptions) (resp azfake.Responder[armdevtestlabs.NotificationChannelsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NotificationChannelsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, labName string, name string, options *armdevtestlabs.NotificationChannelsClientGetOptions) (resp azfake.Responder[armdevtestlabs.NotificationChannelsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method NotificationChannelsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, labName string, options *armdevtestlabs.NotificationChannelsClientListOptions) (resp azfake.PagerResponder[armdevtestlabs.NotificationChannelsClientListResponse])

	// Notify is the fake for method NotificationChannelsClient.Notify
	// HTTP status codes to indicate success: http.StatusOK
	Notify func(ctx context.Context, resourceGroupName string, labName string, name string, notifyParameters armdevtestlabs.NotifyParameters, options *armdevtestlabs.NotificationChannelsClientNotifyOptions) (resp azfake.Responder[armdevtestlabs.NotificationChannelsClientNotifyResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method NotificationChannelsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel armdevtestlabs.NotificationChannelFragment, options *armdevtestlabs.NotificationChannelsClientUpdateOptions) (resp azfake.Responder[armdevtestlabs.NotificationChannelsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewNotificationChannelsServerTransport creates a new instance of NotificationChannelsServerTransport with the provided implementation.
// The returned NotificationChannelsServerTransport instance is connected to an instance of armdevtestlabs.NotificationChannelsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNotificationChannelsServerTransport(srv *NotificationChannelsServer) *NotificationChannelsServerTransport {
	return &NotificationChannelsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdevtestlabs.NotificationChannelsClientListResponse]](),
	}
}

// NotificationChannelsServerTransport connects instances of armdevtestlabs.NotificationChannelsClient to instances of NotificationChannelsServer.
// Don't use this type directly, use NewNotificationChannelsServerTransport instead.
type NotificationChannelsServerTransport struct {
	srv          *NotificationChannelsServer
	newListPager *tracker[azfake.PagerResponder[armdevtestlabs.NotificationChannelsClientListResponse]]
}

// Do implements the policy.Transporter interface for NotificationChannelsServerTransport.
func (n *NotificationChannelsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NotificationChannelsClient.CreateOrUpdate":
		resp, err = n.dispatchCreateOrUpdate(req)
	case "NotificationChannelsClient.Delete":
		resp, err = n.dispatchDelete(req)
	case "NotificationChannelsClient.Get":
		resp, err = n.dispatchGet(req)
	case "NotificationChannelsClient.NewListPager":
		resp, err = n.dispatchNewListPager(req)
	case "NotificationChannelsClient.Notify":
		resp, err = n.dispatchNotify(req)
	case "NotificationChannelsClient.Update":
		resp, err = n.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NotificationChannelsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationchannels/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.NotificationChannel](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, labNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationChannel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationChannelsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if n.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationchannels/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Delete(req.Context(), resourceGroupNameParam, labNameParam, nameParam, nil)
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

func (n *NotificationChannelsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationchannels/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armdevtestlabs.NotificationChannelsClientGetOptions
	if expandParam != nil {
		options = &armdevtestlabs.NotificationChannelsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, labNameParam, nameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationChannel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationChannelsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := n.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationchannels`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armdevtestlabs.NotificationChannelsClientListOptions
		if expandParam != nil || filterParam != nil || topParam != nil || orderbyParam != nil {
			options = &armdevtestlabs.NotificationChannelsClientListOptions{
				Expand:  expandParam,
				Filter:  filterParam,
				Top:     topParam,
				Orderby: orderbyParam,
			}
		}
		resp := n.srv.NewListPager(resourceGroupNameParam, labNameParam, options)
		newListPager = &resp
		n.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdevtestlabs.NotificationChannelsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		n.newListPager.remove(req)
	}
	return resp, nil
}

func (n *NotificationChannelsServerTransport) dispatchNotify(req *http.Request) (*http.Response, error) {
	if n.srv.Notify == nil {
		return nil, &nonRetriableError{errors.New("fake for method Notify not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationchannels/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notify`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.NotifyParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Notify(req.Context(), resourceGroupNameParam, labNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotificationChannelsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notificationchannels/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.NotificationChannelFragment](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Update(req.Context(), resourceGroupNameParam, labNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotificationChannel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
