//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"net/http"
	"net/url"
	"regexp"
)

// BackupWorkloadItemsServer is a fake server for instances of the armrecoveryservicesbackup.BackupWorkloadItemsClient type.
type BackupWorkloadItemsServer struct {
	// NewListPager is the fake for method BackupWorkloadItemsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(vaultName string, resourceGroupName string, fabricName string, containerName string, options *armrecoveryservicesbackup.BackupWorkloadItemsClientListOptions) (resp azfake.PagerResponder[armrecoveryservicesbackup.BackupWorkloadItemsClientListResponse])
}

// NewBackupWorkloadItemsServerTransport creates a new instance of BackupWorkloadItemsServerTransport with the provided implementation.
// The returned BackupWorkloadItemsServerTransport instance is connected to an instance of armrecoveryservicesbackup.BackupWorkloadItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupWorkloadItemsServerTransport(srv *BackupWorkloadItemsServer) *BackupWorkloadItemsServerTransport {
	return &BackupWorkloadItemsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupWorkloadItemsClientListResponse]](),
	}
}

// BackupWorkloadItemsServerTransport connects instances of armrecoveryservicesbackup.BackupWorkloadItemsClient to instances of BackupWorkloadItemsServer.
// Don't use this type directly, use NewBackupWorkloadItemsServerTransport instead.
type BackupWorkloadItemsServerTransport struct {
	srv          *BackupWorkloadItemsServer
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicesbackup.BackupWorkloadItemsClientListResponse]]
}

// Do implements the policy.Transporter interface for BackupWorkloadItemsServerTransport.
func (b *BackupWorkloadItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BackupWorkloadItemsClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BackupWorkloadItemsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/items`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armrecoveryservicesbackup.BackupWorkloadItemsClientListOptions
		if filterParam != nil || skipTokenParam != nil {
			options = &armrecoveryservicesbackup.BackupWorkloadItemsClientListOptions{
				Filter:    filterParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := b.srv.NewListPager(vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicesbackup.BackupWorkloadItemsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}