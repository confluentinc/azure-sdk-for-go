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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PipelineRunsServer is a fake server for instances of the armdatafactory.PipelineRunsClient type.
type PipelineRunsServer struct {
	// Cancel is the fake for method PipelineRunsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, resourceGroupName string, factoryName string, runID string, options *armdatafactory.PipelineRunsClientCancelOptions) (resp azfake.Responder[armdatafactory.PipelineRunsClientCancelResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PipelineRunsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, factoryName string, runID string, options *armdatafactory.PipelineRunsClientGetOptions) (resp azfake.Responder[armdatafactory.PipelineRunsClientGetResponse], errResp azfake.ErrorResponder)

	// QueryByFactory is the fake for method PipelineRunsClient.QueryByFactory
	// HTTP status codes to indicate success: http.StatusOK
	QueryByFactory func(ctx context.Context, resourceGroupName string, factoryName string, filterParameters armdatafactory.RunFilterParameters, options *armdatafactory.PipelineRunsClientQueryByFactoryOptions) (resp azfake.Responder[armdatafactory.PipelineRunsClientQueryByFactoryResponse], errResp azfake.ErrorResponder)
}

// NewPipelineRunsServerTransport creates a new instance of PipelineRunsServerTransport with the provided implementation.
// The returned PipelineRunsServerTransport instance is connected to an instance of armdatafactory.PipelineRunsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPipelineRunsServerTransport(srv *PipelineRunsServer) *PipelineRunsServerTransport {
	return &PipelineRunsServerTransport{srv: srv}
}

// PipelineRunsServerTransport connects instances of armdatafactory.PipelineRunsClient to instances of PipelineRunsServer.
// Don't use this type directly, use NewPipelineRunsServerTransport instead.
type PipelineRunsServerTransport struct {
	srv *PipelineRunsServer
}

// Do implements the policy.Transporter interface for PipelineRunsServerTransport.
func (p *PipelineRunsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PipelineRunsClient.Cancel":
		resp, err = p.dispatchCancel(req)
	case "PipelineRunsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PipelineRunsClient.QueryByFactory":
		resp, err = p.dispatchQueryByFactory(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PipelineRunsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if p.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelineruns/(?P<runId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
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
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	runIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("runId")])
	if err != nil {
		return nil, err
	}
	isRecursiveUnescaped, err := url.QueryUnescape(qp.Get("isRecursive"))
	if err != nil {
		return nil, err
	}
	isRecursiveParam, err := parseOptional(isRecursiveUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armdatafactory.PipelineRunsClientCancelOptions
	if isRecursiveParam != nil {
		options = &armdatafactory.PipelineRunsClientCancelOptions{
			IsRecursive: isRecursiveParam,
		}
	}
	respr, errRespr := p.srv.Cancel(req.Context(), resourceGroupNameParam, factoryNameParam, runIDParam, options)
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

func (p *PipelineRunsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelineruns/(?P<runId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	runIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("runId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, factoryNameParam, runIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PipelineRun, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PipelineRunsServerTransport) dispatchQueryByFactory(req *http.Request) (*http.Response, error) {
	if p.srv.QueryByFactory == nil {
		return nil, &nonRetriableError{errors.New("fake for method QueryByFactory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryPipelineRuns`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.RunFilterParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.QueryByFactory(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PipelineRunsQueryResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}