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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"net/http"
	"net/url"
	"regexp"
)

// ServerParametersServer is a fake server for instances of the armmysql.ServerParametersClient type.
type ServerParametersServer struct {
	// BeginListUpdateConfigurations is the fake for method ServerParametersClient.BeginListUpdateConfigurations
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListUpdateConfigurations func(ctx context.Context, resourceGroupName string, serverName string, value armmysql.ConfigurationListResult, options *armmysql.ServerParametersClientBeginListUpdateConfigurationsOptions) (resp azfake.PollerResponder[armmysql.ServerParametersClientListUpdateConfigurationsResponse], errResp azfake.ErrorResponder)
}

// NewServerParametersServerTransport creates a new instance of ServerParametersServerTransport with the provided implementation.
// The returned ServerParametersServerTransport instance is connected to an instance of armmysql.ServerParametersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerParametersServerTransport(srv *ServerParametersServer) *ServerParametersServerTransport {
	return &ServerParametersServerTransport{
		srv:                           srv,
		beginListUpdateConfigurations: newTracker[azfake.PollerResponder[armmysql.ServerParametersClientListUpdateConfigurationsResponse]](),
	}
}

// ServerParametersServerTransport connects instances of armmysql.ServerParametersClient to instances of ServerParametersServer.
// Don't use this type directly, use NewServerParametersServerTransport instead.
type ServerParametersServerTransport struct {
	srv                           *ServerParametersServer
	beginListUpdateConfigurations *tracker[azfake.PollerResponder[armmysql.ServerParametersClientListUpdateConfigurationsResponse]]
}

// Do implements the policy.Transporter interface for ServerParametersServerTransport.
func (s *ServerParametersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServerParametersClient.BeginListUpdateConfigurations":
		resp, err = s.dispatchBeginListUpdateConfigurations(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerParametersServerTransport) dispatchBeginListUpdateConfigurations(req *http.Request) (*http.Response, error) {
	if s.srv.BeginListUpdateConfigurations == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListUpdateConfigurations not implemented")}
	}
	beginListUpdateConfigurations := s.beginListUpdateConfigurations.get(req)
	if beginListUpdateConfigurations == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforMySQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmysql.ConfigurationListResult](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginListUpdateConfigurations(req.Context(), resourceGroupNameParam, serverNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginListUpdateConfigurations = &respr
		s.beginListUpdateConfigurations.add(req, beginListUpdateConfigurations)
	}

	resp, err := server.PollerResponderNext(beginListUpdateConfigurations, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginListUpdateConfigurations.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginListUpdateConfigurations) {
		s.beginListUpdateConfigurations.remove(req)
	}

	return resp, nil
}
