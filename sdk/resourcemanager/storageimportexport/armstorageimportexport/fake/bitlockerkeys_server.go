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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
	"net/http"
	"net/url"
	"regexp"
)

// BitLockerKeysServer is a fake server for instances of the armstorageimportexport.BitLockerKeysClient type.
type BitLockerKeysServer struct {
	// NewListPager is the fake for method BitLockerKeysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(jobName string, resourceGroupName string, options *armstorageimportexport.BitLockerKeysClientListOptions) (resp azfake.PagerResponder[armstorageimportexport.BitLockerKeysClientListResponse])
}

// NewBitLockerKeysServerTransport creates a new instance of BitLockerKeysServerTransport with the provided implementation.
// The returned BitLockerKeysServerTransport instance is connected to an instance of armstorageimportexport.BitLockerKeysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBitLockerKeysServerTransport(srv *BitLockerKeysServer) *BitLockerKeysServerTransport {
	return &BitLockerKeysServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armstorageimportexport.BitLockerKeysClientListResponse]](),
	}
}

// BitLockerKeysServerTransport connects instances of armstorageimportexport.BitLockerKeysClient to instances of BitLockerKeysServer.
// Don't use this type directly, use NewBitLockerKeysServerTransport instead.
type BitLockerKeysServerTransport struct {
	srv          *BitLockerKeysServer
	newListPager *tracker[azfake.PagerResponder[armstorageimportexport.BitLockerKeysClientListResponse]]
}

// Do implements the policy.Transporter interface for BitLockerKeysServerTransport.
func (b *BitLockerKeysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BitLockerKeysClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BitLockerKeysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ImportExport/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listBitLockerKeys`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		acceptLanguageParam := getOptional(getHeaderValue(req.Header, "Accept-Language"))
		var options *armstorageimportexport.BitLockerKeysClientListOptions
		if acceptLanguageParam != nil {
			options = &armstorageimportexport.BitLockerKeysClientListOptions{
				AcceptLanguage: acceptLanguageParam,
			}
		}
		resp := b.srv.NewListPager(jobNameParam, resourceGroupNameParam, options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
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
