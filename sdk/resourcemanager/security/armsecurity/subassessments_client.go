//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SubAssessmentsClient contains the methods for the SubAssessments group.
// Don't use this type directly, use NewSubAssessmentsClient() instead.
type SubAssessmentsClient struct {
	internal *arm.Client
}

// NewSubAssessmentsClient creates a new instance of SubAssessmentsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSubAssessmentsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SubAssessmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SubAssessmentsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get a security sub-assessment on your scanned resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-01-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - assessmentName - The Assessment Key - Unique key for the assessment type
//   - subAssessmentName - The Sub-Assessment Key - Unique key for the sub-assessment type
//   - options - SubAssessmentsClientGetOptions contains the optional parameters for the SubAssessmentsClient.Get method.
func (client *SubAssessmentsClient) Get(ctx context.Context, scope string, assessmentName string, subAssessmentName string, options *SubAssessmentsClientGetOptions) (SubAssessmentsClientGetResponse, error) {
	var err error
	const operationName = "SubAssessmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, assessmentName, subAssessmentName, options)
	if err != nil {
		return SubAssessmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SubAssessmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SubAssessmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SubAssessmentsClient) getCreateRequest(ctx context.Context, scope string, assessmentName string, subAssessmentName string, options *SubAssessmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments/{subAssessmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	if subAssessmentName == "" {
		return nil, errors.New("parameter subAssessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subAssessmentName}", url.PathEscape(subAssessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubAssessmentsClient) getHandleResponse(resp *http.Response) (SubAssessmentsClientGetResponse, error) {
	result := SubAssessmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubAssessment); err != nil {
		return SubAssessmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get security sub-assessments on all your scanned resources inside a scope
//
// Generated from API version 2019-01-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - assessmentName - The Assessment Key - Unique key for the assessment type
//   - options - SubAssessmentsClientListOptions contains the optional parameters for the SubAssessmentsClient.NewListPager method.
func (client *SubAssessmentsClient) NewListPager(scope string, assessmentName string, options *SubAssessmentsClientListOptions) *runtime.Pager[SubAssessmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubAssessmentsClientListResponse]{
		More: func(page SubAssessmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubAssessmentsClientListResponse) (SubAssessmentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SubAssessmentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, assessmentName, options)
			}, nil)
			if err != nil {
				return SubAssessmentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SubAssessmentsClient) listCreateRequest(ctx context.Context, scope string, assessmentName string, options *SubAssessmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/assessments/{assessmentName}/subAssessments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if assessmentName == "" {
		return nil, errors.New("parameter assessmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assessmentName}", url.PathEscape(assessmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubAssessmentsClient) listHandleResponse(resp *http.Response) (SubAssessmentsClientListResponse, error) {
	result := SubAssessmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubAssessmentList); err != nil {
		return SubAssessmentsClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Get security sub-assessments on all your scanned resources inside a subscription scope
//
// Generated from API version 2019-01-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - options - SubAssessmentsClientListAllOptions contains the optional parameters for the SubAssessmentsClient.NewListAllPager
//     method.
func (client *SubAssessmentsClient) NewListAllPager(scope string, options *SubAssessmentsClientListAllOptions) *runtime.Pager[SubAssessmentsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[SubAssessmentsClientListAllResponse]{
		More: func(page SubAssessmentsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SubAssessmentsClientListAllResponse) (SubAssessmentsClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SubAssessmentsClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return SubAssessmentsClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *SubAssessmentsClient) listAllCreateRequest(ctx context.Context, scope string, options *SubAssessmentsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/subAssessments"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *SubAssessmentsClient) listAllHandleResponse(resp *http.Response) (SubAssessmentsClientListAllResponse, error) {
	result := SubAssessmentsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubAssessmentList); err != nil {
		return SubAssessmentsClientListAllResponse{}, err
	}
	return result, nil
}