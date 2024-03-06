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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/integrationspaces/armintegrationspaces"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ApplicationsServer is a fake server for instances of the armintegrationspaces.ApplicationsClient type.
type ApplicationsServer struct {
	// CreateOrUpdate is the fake for method ApplicationsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, resource armintegrationspaces.Application, options *armintegrationspaces.ApplicationsClientCreateOrUpdateOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ApplicationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, options *armintegrationspaces.ApplicationsClientDeleteOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteBusinessProcessDevelopmentArtifact is the fake for method ApplicationsClient.DeleteBusinessProcessDevelopmentArtifact
	// HTTP status codes to indicate success: http.StatusOK
	DeleteBusinessProcessDevelopmentArtifact func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, body armintegrationspaces.GetOrDeleteBusinessProcessDevelopmentArtifactRequest, options *armintegrationspaces.ApplicationsClientDeleteBusinessProcessDevelopmentArtifactOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientDeleteBusinessProcessDevelopmentArtifactResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ApplicationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, options *armintegrationspaces.ApplicationsClientGetOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientGetResponse], errResp azfake.ErrorResponder)

	// GetBusinessProcessDevelopmentArtifact is the fake for method ApplicationsClient.GetBusinessProcessDevelopmentArtifact
	// HTTP status codes to indicate success: http.StatusOK
	GetBusinessProcessDevelopmentArtifact func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, body armintegrationspaces.GetOrDeleteBusinessProcessDevelopmentArtifactRequest, options *armintegrationspaces.ApplicationsClientGetBusinessProcessDevelopmentArtifactOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientGetBusinessProcessDevelopmentArtifactResponse], errResp azfake.ErrorResponder)

	// ListBusinessProcessDevelopmentArtifacts is the fake for method ApplicationsClient.ListBusinessProcessDevelopmentArtifacts
	// HTTP status codes to indicate success: http.StatusOK
	ListBusinessProcessDevelopmentArtifacts func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, options *armintegrationspaces.ApplicationsClientListBusinessProcessDevelopmentArtifactsOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientListBusinessProcessDevelopmentArtifactsResponse], errResp azfake.ErrorResponder)

	// NewListBySpacePager is the fake for method ApplicationsClient.NewListBySpacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySpacePager func(resourceGroupName string, spaceName string, options *armintegrationspaces.ApplicationsClientListBySpaceOptions) (resp azfake.PagerResponder[armintegrationspaces.ApplicationsClientListBySpaceResponse])

	// Patch is the fake for method ApplicationsClient.Patch
	// HTTP status codes to indicate success: http.StatusOK
	Patch func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, properties armintegrationspaces.ApplicationUpdate, options *armintegrationspaces.ApplicationsClientPatchOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientPatchResponse], errResp azfake.ErrorResponder)

	// SaveBusinessProcessDevelopmentArtifact is the fake for method ApplicationsClient.SaveBusinessProcessDevelopmentArtifact
	// HTTP status codes to indicate success: http.StatusOK
	SaveBusinessProcessDevelopmentArtifact func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, body armintegrationspaces.SaveOrValidateBusinessProcessDevelopmentArtifactRequest, options *armintegrationspaces.ApplicationsClientSaveBusinessProcessDevelopmentArtifactOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientSaveBusinessProcessDevelopmentArtifactResponse], errResp azfake.ErrorResponder)

	// ValidateBusinessProcessDevelopmentArtifact is the fake for method ApplicationsClient.ValidateBusinessProcessDevelopmentArtifact
	// HTTP status codes to indicate success: http.StatusOK
	ValidateBusinessProcessDevelopmentArtifact func(ctx context.Context, resourceGroupName string, spaceName string, applicationName string, body armintegrationspaces.SaveOrValidateBusinessProcessDevelopmentArtifactRequest, options *armintegrationspaces.ApplicationsClientValidateBusinessProcessDevelopmentArtifactOptions) (resp azfake.Responder[armintegrationspaces.ApplicationsClientValidateBusinessProcessDevelopmentArtifactResponse], errResp azfake.ErrorResponder)
}

// NewApplicationsServerTransport creates a new instance of ApplicationsServerTransport with the provided implementation.
// The returned ApplicationsServerTransport instance is connected to an instance of armintegrationspaces.ApplicationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationsServerTransport(srv *ApplicationsServer) *ApplicationsServerTransport {
	return &ApplicationsServerTransport{
		srv:                 srv,
		newListBySpacePager: newTracker[azfake.PagerResponder[armintegrationspaces.ApplicationsClientListBySpaceResponse]](),
	}
}

// ApplicationsServerTransport connects instances of armintegrationspaces.ApplicationsClient to instances of ApplicationsServer.
// Don't use this type directly, use NewApplicationsServerTransport instead.
type ApplicationsServerTransport struct {
	srv                 *ApplicationsServer
	newListBySpacePager *tracker[azfake.PagerResponder[armintegrationspaces.ApplicationsClientListBySpaceResponse]]
}

// Do implements the policy.Transporter interface for ApplicationsServerTransport.
func (a *ApplicationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationsClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "ApplicationsClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "ApplicationsClient.DeleteBusinessProcessDevelopmentArtifact":
		resp, err = a.dispatchDeleteBusinessProcessDevelopmentArtifact(req)
	case "ApplicationsClient.Get":
		resp, err = a.dispatchGet(req)
	case "ApplicationsClient.GetBusinessProcessDevelopmentArtifact":
		resp, err = a.dispatchGetBusinessProcessDevelopmentArtifact(req)
	case "ApplicationsClient.ListBusinessProcessDevelopmentArtifacts":
		resp, err = a.dispatchListBusinessProcessDevelopmentArtifacts(req)
	case "ApplicationsClient.NewListBySpacePager":
		resp, err = a.dispatchNewListBySpacePager(req)
	case "ApplicationsClient.Patch":
		resp, err = a.dispatchPatch(req)
	case "ApplicationsClient.SaveBusinessProcessDevelopmentArtifact":
		resp, err = a.dispatchSaveBusinessProcessDevelopmentArtifact(req)
	case "ApplicationsClient.ValidateBusinessProcessDevelopmentArtifact":
		resp, err = a.dispatchValidateBusinessProcessDevelopmentArtifact(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.Application](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Application, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, nil)
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

func (a *ApplicationsServerTransport) dispatchDeleteBusinessProcessDevelopmentArtifact(req *http.Request) (*http.Response, error) {
	if a.srv.DeleteBusinessProcessDevelopmentArtifact == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteBusinessProcessDevelopmentArtifact not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deleteBusinessProcessDevelopmentArtifact`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.GetOrDeleteBusinessProcessDevelopmentArtifactRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.DeleteBusinessProcessDevelopmentArtifact(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, body, nil)
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

func (a *ApplicationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Application, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchGetBusinessProcessDevelopmentArtifact(req *http.Request) (*http.Response, error) {
	if a.srv.GetBusinessProcessDevelopmentArtifact == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBusinessProcessDevelopmentArtifact not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getBusinessProcessDevelopmentArtifact`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.GetOrDeleteBusinessProcessDevelopmentArtifactRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetBusinessProcessDevelopmentArtifact(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SaveOrGetBusinessProcessDevelopmentArtifactResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchListBusinessProcessDevelopmentArtifacts(req *http.Request) (*http.Response, error) {
	if a.srv.ListBusinessProcessDevelopmentArtifacts == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListBusinessProcessDevelopmentArtifacts not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listBusinessProcessDevelopmentArtifacts`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ListBusinessProcessDevelopmentArtifacts(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ListBusinessProcessDevelopmentArtifactsResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchNewListBySpacePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySpacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySpacePager not implemented")}
	}
	newListBySpacePager := a.newListBySpacePager.get(req)
	if newListBySpacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications`
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
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		selectEscaped := qp["select"]
		selectParam := make([]string, len(selectEscaped))
		for i, v := range selectEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			selectParam[i] = u
		}
		expandEscaped := qp["expand"]
		expandParam := make([]string, len(expandEscaped))
		for i, v := range expandEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			expandParam[i] = u
		}
		orderbyEscaped := qp["orderby"]
		orderbyParam := make([]string, len(orderbyEscaped))
		for i, v := range orderbyEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			orderbyParam[i] = u
		}
		spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
		if err != nil {
			return nil, err
		}
		var options *armintegrationspaces.ApplicationsClientListBySpaceOptions
		if topParam != nil || skipParam != nil || maxpagesizeParam != nil || filterParam != nil || len(selectParam) > 0 || len(expandParam) > 0 || len(orderbyParam) > 0 {
			options = &armintegrationspaces.ApplicationsClientListBySpaceOptions{
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
				Filter:      filterParam,
				Select:      selectParam,
				Expand:      expandParam,
				Orderby:     orderbyParam,
			}
		}
		resp := a.srv.NewListBySpacePager(resourceGroupNameParam, spaceNameParam, options)
		newListBySpacePager = &resp
		a.newListBySpacePager.add(req, newListBySpacePager)
		server.PagerResponderInjectNextLinks(newListBySpacePager, req, func(page *armintegrationspaces.ApplicationsClientListBySpaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySpacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySpacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySpacePager) {
		a.newListBySpacePager.remove(req)
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchPatch(req *http.Request) (*http.Response, error) {
	if a.srv.Patch == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.ApplicationUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Patch(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Application, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchSaveBusinessProcessDevelopmentArtifact(req *http.Request) (*http.Response, error) {
	if a.srv.SaveBusinessProcessDevelopmentArtifact == nil {
		return nil, &nonRetriableError{errors.New("fake for method SaveBusinessProcessDevelopmentArtifact not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/saveBusinessProcessDevelopmentArtifact`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.SaveOrValidateBusinessProcessDevelopmentArtifactRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.SaveBusinessProcessDevelopmentArtifact(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SaveOrGetBusinessProcessDevelopmentArtifactResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *ApplicationsServerTransport) dispatchValidateBusinessProcessDevelopmentArtifact(req *http.Request) (*http.Response, error) {
	if a.srv.ValidateBusinessProcessDevelopmentArtifact == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidateBusinessProcessDevelopmentArtifact not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validateBusinessProcessDevelopmentArtifact`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.SaveOrValidateBusinessProcessDevelopmentArtifactRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.ValidateBusinessProcessDevelopmentArtifact(req.Context(), resourceGroupNameParam, spaceNameParam, applicationNameParam, body, nil)
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
