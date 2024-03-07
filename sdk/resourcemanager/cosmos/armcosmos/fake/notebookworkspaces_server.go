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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
	"net/http"
	"net/url"
	"regexp"
)

// NotebookWorkspacesServer is a fake server for instances of the armcosmos.NotebookWorkspacesClient type.
type NotebookWorkspacesServer struct {
	// BeginCreateOrUpdate is the fake for method NotebookWorkspacesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName armcosmos.NotebookWorkspaceName, notebookCreateUpdateParameters armcosmos.NotebookWorkspaceCreateUpdateParameters, options *armcosmos.NotebookWorkspacesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcosmos.NotebookWorkspacesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NotebookWorkspacesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName armcosmos.NotebookWorkspaceName, options *armcosmos.NotebookWorkspacesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcosmos.NotebookWorkspacesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NotebookWorkspacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName armcosmos.NotebookWorkspaceName, options *armcosmos.NotebookWorkspacesClientGetOptions) (resp azfake.Responder[armcosmos.NotebookWorkspacesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabaseAccountPager is the fake for method NotebookWorkspacesClient.NewListByDatabaseAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabaseAccountPager func(resourceGroupName string, accountName string, options *armcosmos.NotebookWorkspacesClientListByDatabaseAccountOptions) (resp azfake.PagerResponder[armcosmos.NotebookWorkspacesClientListByDatabaseAccountResponse])

	// ListConnectionInfo is the fake for method NotebookWorkspacesClient.ListConnectionInfo
	// HTTP status codes to indicate success: http.StatusOK
	ListConnectionInfo func(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName armcosmos.NotebookWorkspaceName, options *armcosmos.NotebookWorkspacesClientListConnectionInfoOptions) (resp azfake.Responder[armcosmos.NotebookWorkspacesClientListConnectionInfoResponse], errResp azfake.ErrorResponder)

	// BeginRegenerateAuthToken is the fake for method NotebookWorkspacesClient.BeginRegenerateAuthToken
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRegenerateAuthToken func(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName armcosmos.NotebookWorkspaceName, options *armcosmos.NotebookWorkspacesClientBeginRegenerateAuthTokenOptions) (resp azfake.PollerResponder[armcosmos.NotebookWorkspacesClientRegenerateAuthTokenResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method NotebookWorkspacesClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, accountName string, notebookWorkspaceName armcosmos.NotebookWorkspaceName, options *armcosmos.NotebookWorkspacesClientBeginStartOptions) (resp azfake.PollerResponder[armcosmos.NotebookWorkspacesClientStartResponse], errResp azfake.ErrorResponder)
}

// NewNotebookWorkspacesServerTransport creates a new instance of NotebookWorkspacesServerTransport with the provided implementation.
// The returned NotebookWorkspacesServerTransport instance is connected to an instance of armcosmos.NotebookWorkspacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNotebookWorkspacesServerTransport(srv *NotebookWorkspacesServer) *NotebookWorkspacesServerTransport {
	return &NotebookWorkspacesServerTransport{
		srv:                           srv,
		beginCreateOrUpdate:           newTracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientCreateOrUpdateResponse]](),
		beginDelete:                   newTracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientDeleteResponse]](),
		newListByDatabaseAccountPager: newTracker[azfake.PagerResponder[armcosmos.NotebookWorkspacesClientListByDatabaseAccountResponse]](),
		beginRegenerateAuthToken:      newTracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientRegenerateAuthTokenResponse]](),
		beginStart:                    newTracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientStartResponse]](),
	}
}

// NotebookWorkspacesServerTransport connects instances of armcosmos.NotebookWorkspacesClient to instances of NotebookWorkspacesServer.
// Don't use this type directly, use NewNotebookWorkspacesServerTransport instead.
type NotebookWorkspacesServerTransport struct {
	srv                           *NotebookWorkspacesServer
	beginCreateOrUpdate           *tracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientCreateOrUpdateResponse]]
	beginDelete                   *tracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientDeleteResponse]]
	newListByDatabaseAccountPager *tracker[azfake.PagerResponder[armcosmos.NotebookWorkspacesClientListByDatabaseAccountResponse]]
	beginRegenerateAuthToken      *tracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientRegenerateAuthTokenResponse]]
	beginStart                    *tracker[azfake.PollerResponder[armcosmos.NotebookWorkspacesClientStartResponse]]
}

// Do implements the policy.Transporter interface for NotebookWorkspacesServerTransport.
func (n *NotebookWorkspacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NotebookWorkspacesClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NotebookWorkspacesClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NotebookWorkspacesClient.Get":
		resp, err = n.dispatchGet(req)
	case "NotebookWorkspacesClient.NewListByDatabaseAccountPager":
		resp, err = n.dispatchNewListByDatabaseAccountPager(req)
	case "NotebookWorkspacesClient.ListConnectionInfo":
		resp, err = n.dispatchListConnectionInfo(req)
	case "NotebookWorkspacesClient.BeginRegenerateAuthToken":
		resp, err = n.dispatchBeginRegenerateAuthToken(req)
	case "NotebookWorkspacesClient.BeginStart":
		resp, err = n.dispatchBeginStart(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NotebookWorkspacesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebookWorkspaces/(?P<notebookWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcosmos.NotebookWorkspaceCreateUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		notebookWorkspaceNameParam, err := parseWithCast(matches[regex.SubexpIndex("notebookWorkspaceName")], func(v string) (armcosmos.NotebookWorkspaceName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armcosmos.NotebookWorkspaceName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, accountNameParam, notebookWorkspaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NotebookWorkspacesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebookWorkspaces/(?P<notebookWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		notebookWorkspaceNameParam, err := parseWithCast(matches[regex.SubexpIndex("notebookWorkspaceName")], func(v string) (armcosmos.NotebookWorkspaceName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armcosmos.NotebookWorkspaceName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, notebookWorkspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NotebookWorkspacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebookWorkspaces/(?P<notebookWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	notebookWorkspaceNameParam, err := parseWithCast(matches[regex.SubexpIndex("notebookWorkspaceName")], func(v string) (armcosmos.NotebookWorkspaceName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armcosmos.NotebookWorkspaceName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, notebookWorkspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotebookWorkspace, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotebookWorkspacesServerTransport) dispatchNewListByDatabaseAccountPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByDatabaseAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabaseAccountPager not implemented")}
	}
	newListByDatabaseAccountPager := n.newListByDatabaseAccountPager.get(req)
	if newListByDatabaseAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebookWorkspaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByDatabaseAccountPager(resourceGroupNameParam, accountNameParam, nil)
		newListByDatabaseAccountPager = &resp
		n.newListByDatabaseAccountPager.add(req, newListByDatabaseAccountPager)
	}
	resp, err := server.PagerResponderNext(newListByDatabaseAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByDatabaseAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabaseAccountPager) {
		n.newListByDatabaseAccountPager.remove(req)
	}
	return resp, nil
}

func (n *NotebookWorkspacesServerTransport) dispatchListConnectionInfo(req *http.Request) (*http.Response, error) {
	if n.srv.ListConnectionInfo == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListConnectionInfo not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebookWorkspaces/(?P<notebookWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listConnectionInfo`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	notebookWorkspaceNameParam, err := parseWithCast(matches[regex.SubexpIndex("notebookWorkspaceName")], func(v string) (armcosmos.NotebookWorkspaceName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armcosmos.NotebookWorkspaceName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.ListConnectionInfo(req.Context(), resourceGroupNameParam, accountNameParam, notebookWorkspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NotebookWorkspaceConnectionInfoResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NotebookWorkspacesServerTransport) dispatchBeginRegenerateAuthToken(req *http.Request) (*http.Response, error) {
	if n.srv.BeginRegenerateAuthToken == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRegenerateAuthToken not implemented")}
	}
	beginRegenerateAuthToken := n.beginRegenerateAuthToken.get(req)
	if beginRegenerateAuthToken == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebookWorkspaces/(?P<notebookWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateAuthToken`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		notebookWorkspaceNameParam, err := parseWithCast(matches[regex.SubexpIndex("notebookWorkspaceName")], func(v string) (armcosmos.NotebookWorkspaceName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armcosmos.NotebookWorkspaceName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginRegenerateAuthToken(req.Context(), resourceGroupNameParam, accountNameParam, notebookWorkspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRegenerateAuthToken = &respr
		n.beginRegenerateAuthToken.add(req, beginRegenerateAuthToken)
	}

	resp, err := server.PollerResponderNext(beginRegenerateAuthToken, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginRegenerateAuthToken.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRegenerateAuthToken) {
		n.beginRegenerateAuthToken.remove(req)
	}

	return resp, nil
}

func (n *NotebookWorkspacesServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if n.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := n.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/notebookWorkspaces/(?P<notebookWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		notebookWorkspaceNameParam, err := parseWithCast(matches[regex.SubexpIndex("notebookWorkspaceName")], func(v string) (armcosmos.NotebookWorkspaceName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armcosmos.NotebookWorkspaceName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginStart(req.Context(), resourceGroupNameParam, accountNameParam, notebookWorkspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		n.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		n.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		n.beginStart.remove(req)
	}

	return resp, nil
}