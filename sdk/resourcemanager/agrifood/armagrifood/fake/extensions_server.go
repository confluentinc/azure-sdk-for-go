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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ExtensionsServer is a fake server for instances of the armagrifood.ExtensionsClient type.
type ExtensionsServer struct {
	// Create is the fake for method ExtensionsClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, extensionID string, options *armagrifood.ExtensionsClientCreateOptions) (resp azfake.Responder[armagrifood.ExtensionsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ExtensionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, extensionID string, options *armagrifood.ExtensionsClientDeleteOptions) (resp azfake.Responder[armagrifood.ExtensionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExtensionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, extensionID string, options *armagrifood.ExtensionsClientGetOptions) (resp azfake.Responder[armagrifood.ExtensionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFarmBeatsPager is the fake for method ExtensionsClient.NewListByFarmBeatsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFarmBeatsPager func(resourceGroupName string, farmBeatsResourceName string, options *armagrifood.ExtensionsClientListByFarmBeatsOptions) (resp azfake.PagerResponder[armagrifood.ExtensionsClientListByFarmBeatsResponse])

	// Update is the fake for method ExtensionsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, extensionID string, options *armagrifood.ExtensionsClientUpdateOptions) (resp azfake.Responder[armagrifood.ExtensionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewExtensionsServerTransport creates a new instance of ExtensionsServerTransport with the provided implementation.
// The returned ExtensionsServerTransport instance is connected to an instance of armagrifood.ExtensionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExtensionsServerTransport(srv *ExtensionsServer) *ExtensionsServerTransport {
	return &ExtensionsServerTransport{
		srv:                     srv,
		newListByFarmBeatsPager: newTracker[azfake.PagerResponder[armagrifood.ExtensionsClientListByFarmBeatsResponse]](),
	}
}

// ExtensionsServerTransport connects instances of armagrifood.ExtensionsClient to instances of ExtensionsServer.
// Don't use this type directly, use NewExtensionsServerTransport instead.
type ExtensionsServerTransport struct {
	srv                     *ExtensionsServer
	newListByFarmBeatsPager *tracker[azfake.PagerResponder[armagrifood.ExtensionsClientListByFarmBeatsResponse]]
}

// Do implements the policy.Transporter interface for ExtensionsServerTransport.
func (e *ExtensionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExtensionsClient.Create":
		resp, err = e.dispatchCreate(req)
	case "ExtensionsClient.Delete":
		resp, err = e.dispatchDelete(req)
	case "ExtensionsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExtensionsClient.NewListByFarmBeatsPager":
		resp, err = e.dispatchNewListByFarmBeatsPager(req)
	case "ExtensionsClient.Update":
		resp, err = e.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if e.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<farmBeatsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	farmBeatsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("farmBeatsResourceName")])
	if err != nil {
		return nil, err
	}
	extensionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Create(req.Context(), resourceGroupNameParam, farmBeatsResourceNameParam, extensionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Extension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if e.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<farmBeatsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	farmBeatsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("farmBeatsResourceName")])
	if err != nil {
		return nil, err
	}
	extensionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Delete(req.Context(), resourceGroupNameParam, farmBeatsResourceNameParam, extensionIDParam, nil)
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

func (e *ExtensionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<farmBeatsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	farmBeatsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("farmBeatsResourceName")])
	if err != nil {
		return nil, err
	}
	extensionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, farmBeatsResourceNameParam, extensionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Extension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchNewListByFarmBeatsPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByFarmBeatsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFarmBeatsPager not implemented")}
	}
	newListByFarmBeatsPager := e.newListByFarmBeatsPager.get(req)
	if newListByFarmBeatsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<farmBeatsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions`
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
		farmBeatsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("farmBeatsResourceName")])
		if err != nil {
			return nil, err
		}
		extensionIDsEscaped := qp["extensionIds"]
		extensionIDsParam := make([]string, len(extensionIDsEscaped))
		for i, v := range extensionIDsEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			extensionIDsParam[i] = u
		}
		extensionCategoriesEscaped := qp["extensionCategories"]
		extensionCategoriesParam := make([]string, len(extensionCategoriesEscaped))
		for i, v := range extensionCategoriesEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			extensionCategoriesParam[i] = u
		}
		maxPageSizeUnescaped, err := url.QueryUnescape(qp.Get("$maxPageSize"))
		if err != nil {
			return nil, err
		}
		maxPageSizeParam, err := parseOptional(maxPageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armagrifood.ExtensionsClientListByFarmBeatsOptions
		if len(extensionIDsParam) > 0 || len(extensionCategoriesParam) > 0 || maxPageSizeParam != nil || skipTokenParam != nil {
			options = &armagrifood.ExtensionsClientListByFarmBeatsOptions{
				ExtensionIDs:        extensionIDsParam,
				ExtensionCategories: extensionCategoriesParam,
				MaxPageSize:         maxPageSizeParam,
				SkipToken:           skipTokenParam,
			}
		}
		resp := e.srv.NewListByFarmBeatsPager(resourceGroupNameParam, farmBeatsResourceNameParam, options)
		newListByFarmBeatsPager = &resp
		e.newListByFarmBeatsPager.add(req, newListByFarmBeatsPager)
		server.PagerResponderInjectNextLinks(newListByFarmBeatsPager, req, func(page *armagrifood.ExtensionsClientListByFarmBeatsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFarmBeatsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByFarmBeatsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFarmBeatsPager) {
		e.newListByFarmBeatsPager.remove(req)
	}
	return resp, nil
}

func (e *ExtensionsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<farmBeatsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/extensions/(?P<extensionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	farmBeatsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("farmBeatsResourceName")])
	if err != nil {
		return nil, err
	}
	extensionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("extensionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Update(req.Context(), resourceGroupNameParam, farmBeatsResourceNameParam, extensionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Extension, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}