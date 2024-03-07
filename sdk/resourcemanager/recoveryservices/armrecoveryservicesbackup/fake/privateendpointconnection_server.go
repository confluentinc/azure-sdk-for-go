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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionServer is a fake server for instances of the armrecoveryservicesbackup.PrivateEndpointConnectionClient type.
type PrivateEndpointConnectionServer struct {
	// BeginDelete is the fake for method PrivateEndpointConnectionClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, options *armrecoveryservicesbackup.PrivateEndpointConnectionClientBeginDeleteOptions) (resp azfake.PollerResponder[armrecoveryservicesbackup.PrivateEndpointConnectionClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, options *armrecoveryservicesbackup.PrivateEndpointConnectionClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.PrivateEndpointConnectionClientGetResponse], errResp azfake.ErrorResponder)

	// BeginPut is the fake for method PrivateEndpointConnectionClient.BeginPut
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPut func(ctx context.Context, vaultName string, resourceGroupName string, privateEndpointConnectionName string, parameters armrecoveryservicesbackup.PrivateEndpointConnectionResource, options *armrecoveryservicesbackup.PrivateEndpointConnectionClientBeginPutOptions) (resp azfake.PollerResponder[armrecoveryservicesbackup.PrivateEndpointConnectionClientPutResponse], errResp azfake.ErrorResponder)
}

// NewPrivateEndpointConnectionServerTransport creates a new instance of PrivateEndpointConnectionServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionServerTransport instance is connected to an instance of armrecoveryservicesbackup.PrivateEndpointConnectionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionServerTransport(srv *PrivateEndpointConnectionServer) *PrivateEndpointConnectionServerTransport {
	return &PrivateEndpointConnectionServerTransport{
		srv:         srv,
		beginDelete: newTracker[azfake.PollerResponder[armrecoveryservicesbackup.PrivateEndpointConnectionClientDeleteResponse]](),
		beginPut:    newTracker[azfake.PollerResponder[armrecoveryservicesbackup.PrivateEndpointConnectionClientPutResponse]](),
	}
}

// PrivateEndpointConnectionServerTransport connects instances of armrecoveryservicesbackup.PrivateEndpointConnectionClient to instances of PrivateEndpointConnectionServer.
// Don't use this type directly, use NewPrivateEndpointConnectionServerTransport instead.
type PrivateEndpointConnectionServerTransport struct {
	srv         *PrivateEndpointConnectionServer
	beginDelete *tracker[azfake.PollerResponder[armrecoveryservicesbackup.PrivateEndpointConnectionClientDeleteResponse]]
	beginPut    *tracker[azfake.PollerResponder[armrecoveryservicesbackup.PrivateEndpointConnectionClientPutResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionServerTransport.
func (p *PrivateEndpointConnectionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateEndpointConnectionClient.BeginDelete":
		resp, err = p.dispatchBeginDelete(req)
	case "PrivateEndpointConnectionClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateEndpointConnectionClient.BeginPut":
		resp, err = p.dispatchBeginPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), vaultNameParam, resourceGroupNameParam, privateEndpointConnectionNameParam, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnectionResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionServerTransport) dispatchBeginPut(req *http.Request) (*http.Response, error) {
	if p.srv.BeginPut == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPut not implemented")}
	}
	beginPut := p.beginPut.get(req)
	if beginPut == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.PrivateEndpointConnectionResource](req)
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginPut(req.Context(), vaultNameParam, resourceGroupNameParam, privateEndpointConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPut = &respr
		p.beginPut.add(req, beginPut)
	}

	resp, err := server.PollerResponderNext(beginPut, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginPut.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPut) {
		p.beginPut.remove(req)
	}

	return resp, nil
}