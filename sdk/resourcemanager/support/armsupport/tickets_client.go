//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// TicketsClient contains the methods for the SupportTickets group.
// Don't use this type directly, use NewTicketsClient() instead.
type TicketsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTicketsClient creates a new instance of TicketsClient with the specified values.
//   - subscriptionID - Azure subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTicketsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TicketsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TicketsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - Check the availability of a resource name. This API should be used to check the uniqueness of the
// name for support ticket creation for the selected subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - checkNameAvailabilityInput - Input to check.
//   - options - TicketsClientCheckNameAvailabilityOptions contains the optional parameters for the TicketsClient.CheckNameAvailability
//     method.
func (client *TicketsClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, options *TicketsClientCheckNameAvailabilityOptions) (TicketsClientCheckNameAvailabilityResponse, error) {
	var err error
	const operationName = "TicketsClient.CheckNameAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkNameAvailabilityCreateRequest(ctx, checkNameAvailabilityInput, options)
	if err != nil {
		return TicketsClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TicketsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TicketsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *TicketsClient) checkNameAvailabilityCreateRequest(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput, options *TicketsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkNameAvailabilityInput); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *TicketsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (TicketsClientCheckNameAvailabilityResponse, error) {
	result := TicketsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityOutput); err != nil {
		return TicketsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreate - Creates a new support ticket for Subscription and Service limits (Quota), Technical, Billing, and Subscription
// Management issues for the specified subscription. Learn the prerequisites
// [https://aka.ms/supportAPI] required to create a support ticket.
// Always call the Services and ProblemClassifications API to get the most recent set of services and problem categories required
// for support ticket creation.
// Adding attachments is not currently supported via the API. To add a file to an existing support ticket, visit the Manage
// support ticket [https://portal.azure.com/#blade/MicrosoftAzure
// Support/HelpAndSupportBlade/managesupportrequest] page in the Azure portal, select the support ticket, and use the file
// upload control to add a new file.
// Providing consent to share diagnostic information with Azure support is currently not supported via the API. The Azure
// support engineer working on your ticket will reach out to you for consent if your
// issue requires gathering diagnostic information from your Azure resources.
// Creating a support ticket for on-behalf-of: Include x-ms-authorization-auxiliary header to provide an auxiliary token as
// per documentation
// [https://docs.microsoft.com/azure/azure-resource-manager/management/authenticate-multi-tenant]. The primary token will
// be from the tenant for whom a support ticket is being raised against the
// subscription, i.e. Cloud solution provider (CSP) customer tenant. The auxiliary token will be from the Cloud solution provider
// (CSP) partner tenant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - supportTicketName - Support ticket name.
//   - createSupportTicketParameters - Support ticket request payload.
//   - options - TicketsClientBeginCreateOptions contains the optional parameters for the TicketsClient.BeginCreate method.
func (client *TicketsClient) BeginCreate(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails, options *TicketsClientBeginCreateOptions) (*runtime.Poller[TicketsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, supportTicketName, createSupportTicketParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TicketsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[TicketsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a new support ticket for Subscription and Service limits (Quota), Technical, Billing, and Subscription
// Management issues for the specified subscription. Learn the prerequisites
// [https://aka.ms/supportAPI] required to create a support ticket.
// Always call the Services and ProblemClassifications API to get the most recent set of services and problem categories required
// for support ticket creation.
// Adding attachments is not currently supported via the API. To add a file to an existing support ticket, visit the Manage
// support ticket [https://portal.azure.com/#blade/MicrosoftAzure
// Support/HelpAndSupportBlade/managesupportrequest] page in the Azure portal, select the support ticket, and use the file
// upload control to add a new file.
// Providing consent to share diagnostic information with Azure support is currently not supported via the API. The Azure
// support engineer working on your ticket will reach out to you for consent if your
// issue requires gathering diagnostic information from your Azure resources.
// Creating a support ticket for on-behalf-of: Include x-ms-authorization-auxiliary header to provide an auxiliary token as
// per documentation
// [https://docs.microsoft.com/azure/azure-resource-manager/management/authenticate-multi-tenant]. The primary token will
// be from the tenant for whom a support ticket is being raised against the
// subscription, i.e. Cloud solution provider (CSP) customer tenant. The auxiliary token will be from the Cloud solution provider
// (CSP) partner tenant.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *TicketsClient) create(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails, options *TicketsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "TicketsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, supportTicketName, createSupportTicketParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *TicketsClient) createCreateRequest(ctx context.Context, supportTicketName string, createSupportTicketParameters TicketDetails, options *TicketsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createSupportTicketParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get ticket details for an Azure subscription. Support ticket data is available for 18 months after ticket creation.
// If a ticket was created more than 18 months ago, a request for data might cause an
// error.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - supportTicketName - Support ticket name.
//   - options - TicketsClientGetOptions contains the optional parameters for the TicketsClient.Get method.
func (client *TicketsClient) Get(ctx context.Context, supportTicketName string, options *TicketsClientGetOptions) (TicketsClientGetResponse, error) {
	var err error
	const operationName = "TicketsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, supportTicketName, options)
	if err != nil {
		return TicketsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TicketsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TicketsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TicketsClient) getCreateRequest(ctx context.Context, supportTicketName string, options *TicketsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TicketsClient) getHandleResponse(resp *http.Response) (TicketsClientGetResponse, error) {
	result := TicketsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TicketDetails); err != nil {
		return TicketsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the support tickets for an Azure subscription. You can also filter the support tickets by Status,
// CreatedDate, ServiceId, and ProblemClassificationId using the $filter parameter. Output will
// be a paged result with nextLink, using which you can retrieve the next set of support tickets.
// Support ticket data is available for 18 months after ticket creation. If a ticket was created more than 18 months ago,
// a request for data might cause an error.
//
// Generated from API version 2022-09-01-preview
//   - options - TicketsClientListOptions contains the optional parameters for the TicketsClient.NewListPager method.
func (client *TicketsClient) NewListPager(options *TicketsClientListOptions) *runtime.Pager[TicketsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TicketsClientListResponse]{
		More: func(page TicketsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TicketsClientListResponse) (TicketsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TicketsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return TicketsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TicketsClient) listCreateRequest(ctx context.Context, options *TicketsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TicketsClient) listHandleResponse(resp *http.Response) (TicketsClientListResponse, error) {
	result := TicketsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TicketsListResult); err != nil {
		return TicketsClientListResponse{}, err
	}
	return result, nil
}

// Update - This API allows you to update the severity level, ticket status, advanced diagnostic consent and your contact
// information in the support ticket.
// Note: The severity levels cannot be changed if a support ticket is actively being worked upon by an Azure support engineer.
// In such a case, contact your support engineer to request severity update by
// adding a new communication using the Communications API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - supportTicketName - Support ticket name.
//   - updateSupportTicket - UpdateSupportTicket object.
//   - options - TicketsClientUpdateOptions contains the optional parameters for the TicketsClient.Update method.
func (client *TicketsClient) Update(ctx context.Context, supportTicketName string, updateSupportTicket UpdateSupportTicket, options *TicketsClientUpdateOptions) (TicketsClientUpdateResponse, error) {
	var err error
	const operationName = "TicketsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, supportTicketName, updateSupportTicket, options)
	if err != nil {
		return TicketsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TicketsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TicketsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TicketsClient) updateCreateRequest(ctx context.Context, supportTicketName string, updateSupportTicket UpdateSupportTicket, options *TicketsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Support/supportTickets/{supportTicketName}"
	if supportTicketName == "" {
		return nil, errors.New("parameter supportTicketName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{supportTicketName}", url.PathEscape(supportTicketName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateSupportTicket); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TicketsClient) updateHandleResponse(resp *http.Response) (TicketsClientUpdateResponse, error) {
	result := TicketsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TicketDetails); err != nil {
		return TicketsClientUpdateResponse{}, err
	}
	return result, nil
}