package marketplacemeteredbilling

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
)

// Client is the REST API for MarketplaceMeteredBilling.
type Client struct {
    BaseClient
}
// NewClient creates an instance of the Client client.
func NewClient() Client {
    return NewClientWithBaseURI(DefaultBaseURI, )
}

// NewClientWithBaseURI creates an instance of the Client client.
    func NewClientWithBaseURI(baseURI string, ) Client {
        return Client{ NewWithBaseURI(baseURI, )}
    }

// BatchUsageEventMethod report batch usage events.
    // Parameters:
        // parameters - parameters supplied to report batch usage events.
func (client Client) BatchUsageEventMethod(ctx context.Context, parameters BatchUsageEvent) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/Client.BatchUsageEventMethod")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.BatchUsageEventMethodPreparer(ctx, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "marketplacemeteredbilling.Client", "BatchUsageEventMethod", nil , "Failure preparing request")
    return
    }

            resp, err := client.BatchUsageEventMethodSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "marketplacemeteredbilling.Client", "BatchUsageEventMethod", resp, "Failure sending request")
            return
            }

            result, err = client.BatchUsageEventMethodResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "marketplacemeteredbilling.Client", "BatchUsageEventMethod", resp, "Failure responding to request")
            }

    return
    }

    // BatchUsageEventMethodPreparer prepares the BatchUsageEventMethod request.
    func (client Client) BatchUsageEventMethodPreparer(ctx context.Context, parameters BatchUsageEvent) (*http.Request, error) {
                    const APIVersion = "2018-08-31"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPath("/batchUsageEvent"),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // BatchUsageEventMethodSender sends the BatchUsageEventMethod request. The method will close the
    // http.Response Body if it receives an error.
    func (client Client) BatchUsageEventMethodSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// BatchUsageEventMethodResponder handles the response to the BatchUsageEventMethod request. The method always
// closes the http.Response Body.
func (client Client) BatchUsageEventMethodResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusBadRequest,http.StatusForbidden),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// UsageEventMethod report single usage event.
    // Parameters:
        // parameters - parameters supplied to report a single usage event.
        // xMsRequestid - a unique string value for tracking the request from the client, preferably a GUID. If this
        // value isn't provided, one will be generated and provided in the response headers.
        // xMsCorrelationid - a unique string value for operation on the client. This parameter correlates all events
        // from client operation with events on the server side. If this value isn't provided, one will be generated
        // and provided in the response headers.
func (client Client) UsageEventMethod(ctx context.Context, parameters UsageEvent, xMsRequestid string, xMsCorrelationid string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/Client.UsageEventMethod")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UsageEventMethodPreparer(ctx, parameters, xMsRequestid, xMsCorrelationid)
    if err != nil {
    err = autorest.NewErrorWithError(err, "marketplacemeteredbilling.Client", "UsageEventMethod", nil , "Failure preparing request")
    return
    }

            resp, err := client.UsageEventMethodSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "marketplacemeteredbilling.Client", "UsageEventMethod", resp, "Failure sending request")
            return
            }

            result, err = client.UsageEventMethodResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "marketplacemeteredbilling.Client", "UsageEventMethod", resp, "Failure responding to request")
            }

    return
    }

    // UsageEventMethodPreparer prepares the UsageEventMethod request.
    func (client Client) UsageEventMethodPreparer(ctx context.Context, parameters UsageEvent, xMsRequestid string, xMsCorrelationid string) (*http.Request, error) {
                    const APIVersion = "2018-08-31"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPath("/usageEvent"),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
            if len(xMsRequestid) > 0 {
            preparer = autorest.DecoratePreparer(preparer,
            autorest.WithHeader("x-ms-requestid",autorest.String(xMsRequestid)))
            }
            if len(xMsCorrelationid) > 0 {
            preparer = autorest.DecoratePreparer(preparer,
            autorest.WithHeader("x-ms-correlationid",autorest.String(xMsCorrelationid)))
            }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UsageEventMethodSender sends the UsageEventMethod request. The method will close the
    // http.Response Body if it receives an error.
    func (client Client) UsageEventMethodSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// UsageEventMethodResponder handles the response to the UsageEventMethod request. The method always
// closes the http.Response Body.
func (client Client) UsageEventMethodResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusBadRequest,http.StatusForbidden,http.StatusConflict),
    autorest.ByClosing())
    result.Response = resp
        return
    }

