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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
	"net/http"
	"net/url"
	"regexp"
)

// SubscriptionsServer is a fake server for instances of the armstreamanalytics.SubscriptionsClient type.
type SubscriptionsServer struct {
	// CompileQuery is the fake for method SubscriptionsClient.CompileQuery
	// HTTP status codes to indicate success: http.StatusOK
	CompileQuery func(ctx context.Context, location string, compileQuery armstreamanalytics.CompileQuery, options *armstreamanalytics.SubscriptionsClientCompileQueryOptions) (resp azfake.Responder[armstreamanalytics.SubscriptionsClientCompileQueryResponse], errResp azfake.ErrorResponder)

	// ListQuotas is the fake for method SubscriptionsClient.ListQuotas
	// HTTP status codes to indicate success: http.StatusOK
	ListQuotas func(ctx context.Context, location string, options *armstreamanalytics.SubscriptionsClientListQuotasOptions) (resp azfake.Responder[armstreamanalytics.SubscriptionsClientListQuotasResponse], errResp azfake.ErrorResponder)

	// BeginSampleInput is the fake for method SubscriptionsClient.BeginSampleInput
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginSampleInput func(ctx context.Context, location string, sampleInput armstreamanalytics.SampleInput, options *armstreamanalytics.SubscriptionsClientBeginSampleInputOptions) (resp azfake.PollerResponder[armstreamanalytics.SubscriptionsClientSampleInputResponse], errResp azfake.ErrorResponder)

	// BeginTestInput is the fake for method SubscriptionsClient.BeginTestInput
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginTestInput func(ctx context.Context, location string, testInput armstreamanalytics.TestInput, options *armstreamanalytics.SubscriptionsClientBeginTestInputOptions) (resp azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestInputResponse], errResp azfake.ErrorResponder)

	// BeginTestOutput is the fake for method SubscriptionsClient.BeginTestOutput
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginTestOutput func(ctx context.Context, location string, testOutput armstreamanalytics.TestOutput, options *armstreamanalytics.SubscriptionsClientBeginTestOutputOptions) (resp azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestOutputResponse], errResp azfake.ErrorResponder)

	// BeginTestQuery is the fake for method SubscriptionsClient.BeginTestQuery
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginTestQuery func(ctx context.Context, location string, testQuery armstreamanalytics.TestQuery, options *armstreamanalytics.SubscriptionsClientBeginTestQueryOptions) (resp azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestQueryResponse], errResp azfake.ErrorResponder)
}

// NewSubscriptionsServerTransport creates a new instance of SubscriptionsServerTransport with the provided implementation.
// The returned SubscriptionsServerTransport instance is connected to an instance of armstreamanalytics.SubscriptionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSubscriptionsServerTransport(srv *SubscriptionsServer) *SubscriptionsServerTransport {
	return &SubscriptionsServerTransport{
		srv:              srv,
		beginSampleInput: newTracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientSampleInputResponse]](),
		beginTestInput:   newTracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestInputResponse]](),
		beginTestOutput:  newTracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestOutputResponse]](),
		beginTestQuery:   newTracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestQueryResponse]](),
	}
}

// SubscriptionsServerTransport connects instances of armstreamanalytics.SubscriptionsClient to instances of SubscriptionsServer.
// Don't use this type directly, use NewSubscriptionsServerTransport instead.
type SubscriptionsServerTransport struct {
	srv              *SubscriptionsServer
	beginSampleInput *tracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientSampleInputResponse]]
	beginTestInput   *tracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestInputResponse]]
	beginTestOutput  *tracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestOutputResponse]]
	beginTestQuery   *tracker[azfake.PollerResponder[armstreamanalytics.SubscriptionsClientTestQueryResponse]]
}

// Do implements the policy.Transporter interface for SubscriptionsServerTransport.
func (s *SubscriptionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SubscriptionsClient.CompileQuery":
		resp, err = s.dispatchCompileQuery(req)
	case "SubscriptionsClient.ListQuotas":
		resp, err = s.dispatchListQuotas(req)
	case "SubscriptionsClient.BeginSampleInput":
		resp, err = s.dispatchBeginSampleInput(req)
	case "SubscriptionsClient.BeginTestInput":
		resp, err = s.dispatchBeginTestInput(req)
	case "SubscriptionsClient.BeginTestOutput":
		resp, err = s.dispatchBeginTestOutput(req)
	case "SubscriptionsClient.BeginTestQuery":
		resp, err = s.dispatchBeginTestQuery(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchCompileQuery(req *http.Request) (*http.Response, error) {
	if s.srv.CompileQuery == nil {
		return nil, &nonRetriableError{errors.New("fake for method CompileQuery not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/compileQuery`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstreamanalytics.CompileQuery](req)
	if err != nil {
		return nil, err
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CompileQuery(req.Context(), locationParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueryCompilationResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchListQuotas(req *http.Request) (*http.Response, error) {
	if s.srv.ListQuotas == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListQuotas not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/quotas`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ListQuotas(req.Context(), locationParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionQuotasListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchBeginSampleInput(req *http.Request) (*http.Response, error) {
	if s.srv.BeginSampleInput == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSampleInput not implemented")}
	}
	beginSampleInput := s.beginSampleInput.get(req)
	if beginSampleInput == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sampleInput`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstreamanalytics.SampleInput](req)
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginSampleInput(req.Context(), locationParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSampleInput = &respr
		s.beginSampleInput.add(req, beginSampleInput)
	}

	resp, err := server.PollerResponderNext(beginSampleInput, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		s.beginSampleInput.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSampleInput) {
		s.beginSampleInput.remove(req)
	}

	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchBeginTestInput(req *http.Request) (*http.Response, error) {
	if s.srv.BeginTestInput == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTestInput not implemented")}
	}
	beginTestInput := s.beginTestInput.get(req)
	if beginTestInput == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/testInput`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstreamanalytics.TestInput](req)
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginTestInput(req.Context(), locationParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTestInput = &respr
		s.beginTestInput.add(req, beginTestInput)
	}

	resp, err := server.PollerResponderNext(beginTestInput, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		s.beginTestInput.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTestInput) {
		s.beginTestInput.remove(req)
	}

	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchBeginTestOutput(req *http.Request) (*http.Response, error) {
	if s.srv.BeginTestOutput == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTestOutput not implemented")}
	}
	beginTestOutput := s.beginTestOutput.get(req)
	if beginTestOutput == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/testOutput`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstreamanalytics.TestOutput](req)
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginTestOutput(req.Context(), locationParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTestOutput = &respr
		s.beginTestOutput.add(req, beginTestOutput)
	}

	resp, err := server.PollerResponderNext(beginTestOutput, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		s.beginTestOutput.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTestOutput) {
		s.beginTestOutput.remove(req)
	}

	return resp, nil
}

func (s *SubscriptionsServerTransport) dispatchBeginTestQuery(req *http.Request) (*http.Response, error) {
	if s.srv.BeginTestQuery == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTestQuery not implemented")}
	}
	beginTestQuery := s.beginTestQuery.get(req)
	if beginTestQuery == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StreamAnalytics/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/testQuery`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstreamanalytics.TestQuery](req)
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginTestQuery(req.Context(), locationParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTestQuery = &respr
		s.beginTestQuery.add(req, beginTestQuery)
	}

	resp, err := server.PollerResponderNext(beginTestQuery, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginTestQuery.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTestQuery) {
		s.beginTestQuery.remove(req)
	}

	return resp, nil
}
