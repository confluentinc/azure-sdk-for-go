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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionsServer is a fake server for instances of the armcosmosforpostgresql.PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method PrivateEndpointConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, parameters armcosmosforpostgresql.PrivateEndpointConnection, options *armcosmosforpostgresql.PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateEndpointConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, options *armcosmosforpostgresql.PrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, privateEndpointConnectionName string, options *armcosmosforpostgresql.PrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armcosmosforpostgresql.PrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByClusterPager is the fake for method PrivateEndpointConnectionsClient.NewListByClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByClusterPager func(resourceGroupName string, clusterName string, options *armcosmosforpostgresql.PrivateEndpointConnectionsClientListByClusterOptions) (resp azfake.PagerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientListByClusterResponse])
}

// NewPrivateEndpointConnectionsServerTransport creates a new instance of PrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionsServerTransport instance is connected to an instance of armcosmosforpostgresql.PrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionsServerTransport(srv *PrivateEndpointConnectionsServer) *PrivateEndpointConnectionsServerTransport {
	return &PrivateEndpointConnectionsServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientDeleteResponse]](),
		newListByClusterPager: newTracker[azfake.PagerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientListByClusterResponse]](),
	}
}

// PrivateEndpointConnectionsServerTransport connects instances of armcosmosforpostgresql.PrivateEndpointConnectionsClient to instances of PrivateEndpointConnectionsServer.
// Don't use this type directly, use NewPrivateEndpointConnectionsServerTransport instead.
type PrivateEndpointConnectionsServerTransport struct {
	srv                   *PrivateEndpointConnectionsServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientDeleteResponse]]
	newListByClusterPager *tracker[azfake.PagerResponder[armcosmosforpostgresql.PrivateEndpointConnectionsClientListByClusterResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionsServerTransport.
func (p *PrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateEndpointConnectionsClient.BeginCreateOrUpdate":
		resp, err = p.dispatchBeginCreateOrUpdate(req)
	case "PrivateEndpointConnectionsClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateEndpointConnectionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateEndpointConnectionsClient.NewListByClusterPager":
		resp, err = p.dispatchNewListByClusterPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/serverGroupsv2/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcosmosforpostgresql.PrivateEndpointConnection](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, clusterNameParam, privateEndpointConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/serverGroupsv2/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterNameParam, privateEndpointConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/serverGroupsv2/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchNewListByClusterPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByClusterPager not implemented")}
	}
	newListByClusterPager := p.newListByClusterPager.get(req)
	if newListByClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/serverGroupsv2/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByClusterPager(resourceGroupNameParam, clusterNameParam, nil)
		newListByClusterPager = &resp
		p.newListByClusterPager.add(req, newListByClusterPager)
	}
	resp, err := server.PagerResponderNext(newListByClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByClusterPager) {
		p.newListByClusterPager.remove(req)
	}
	return resp, nil
}