package marketplacefulfillment

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

// Client is the REST API for MarketplaceFulfillment.
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

// Resolve resolve marketplace subscription.
    // Parameters:
        // xMsMarketplaceToken - the token query parameter in the URL when the user is redirected to the SaaS partner’s
        // website from Azure (for example: https://contoso.com/signup?token=..). Note: The URL decodes the token value
        // from the browser before using it.
        // xMsRequestid - a unique string value for tracking the request from the client, preferably a GUID. If this
        // value isn't provided, one will be generated and provided in the response headers.
        // xMsCorrelationid - a unique string value for operation on the client. This parameter correlates all events
        // from client operation with events on the server side. If this value isn't provided, one will be generated
        // and provided in the response headers.
func (client Client) Resolve(ctx context.Context, xMsMarketplaceToken string, xMsRequestid string, xMsCorrelationid string) (result Subscription, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/Client.Resolve")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ResolvePreparer(ctx, xMsMarketplaceToken, xMsRequestid, xMsCorrelationid)
    if err != nil {
    err = autorest.NewErrorWithError(err, "marketplacefulfillment.Client", "Resolve", nil , "Failure preparing request")
    return
    }

            resp, err := client.ResolveSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "marketplacefulfillment.Client", "Resolve", resp, "Failure sending request")
            return
            }

            result, err = client.ResolveResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "marketplacefulfillment.Client", "Resolve", resp, "Failure responding to request")
            }

    return
    }

    // ResolvePreparer prepares the Resolve request.
    func (client Client) ResolvePreparer(ctx context.Context, xMsMarketplaceToken string, xMsRequestid string, xMsCorrelationid string) (*http.Request, error) {
                    const APIVersion = "2018-08-31"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPath("/resolve"),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("x-ms-marketplace-token", autorest.String(xMsMarketplaceToken)))
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

    // ResolveSender sends the Resolve request. The method will close the
    // http.Response Body if it receives an error.
    func (client Client) ResolveSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ResolveResponder handles the response to the Resolve request. The method always
// closes the http.Response Body.
func (client Client) ResolveResponder(resp *http.Response) (result Subscription, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }
