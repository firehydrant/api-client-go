// Code generated by go-swagger; DO NOT EDIT.

package status_pages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteV1NuncSubscriptionsUnsubscribeTokenParams creates a new DeleteV1NuncSubscriptionsUnsubscribeTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1NuncSubscriptionsUnsubscribeTokenParams() *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	return &DeleteV1NuncSubscriptionsUnsubscribeTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1NuncSubscriptionsUnsubscribeTokenParamsWithTimeout creates a new DeleteV1NuncSubscriptionsUnsubscribeTokenParams object
// with the ability to set a timeout on a request.
func NewDeleteV1NuncSubscriptionsUnsubscribeTokenParamsWithTimeout(timeout time.Duration) *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	return &DeleteV1NuncSubscriptionsUnsubscribeTokenParams{
		timeout: timeout,
	}
}

// NewDeleteV1NuncSubscriptionsUnsubscribeTokenParamsWithContext creates a new DeleteV1NuncSubscriptionsUnsubscribeTokenParams object
// with the ability to set a context for a request.
func NewDeleteV1NuncSubscriptionsUnsubscribeTokenParamsWithContext(ctx context.Context) *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	return &DeleteV1NuncSubscriptionsUnsubscribeTokenParams{
		Context: ctx,
	}
}

// NewDeleteV1NuncSubscriptionsUnsubscribeTokenParamsWithHTTPClient creates a new DeleteV1NuncSubscriptionsUnsubscribeTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1NuncSubscriptionsUnsubscribeTokenParamsWithHTTPClient(client *http.Client) *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	return &DeleteV1NuncSubscriptionsUnsubscribeTokenParams{
		HTTPClient: client,
	}
}

/*
DeleteV1NuncSubscriptionsUnsubscribeTokenParams contains all the parameters to send to the API endpoint

	for the delete v1 nunc subscriptions unsubscribe token operation.

	Typically these are written to a http.Request.
*/
type DeleteV1NuncSubscriptionsUnsubscribeTokenParams struct {

	// UnsubscribeToken.
	UnsubscribeToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 nunc subscriptions unsubscribe token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) WithDefaults() *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 nunc subscriptions unsubscribe token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) WithTimeout(timeout time.Duration) *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) WithContext(ctx context.Context) *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) WithHTTPClient(client *http.Client) *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUnsubscribeToken adds the unsubscribeToken to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) WithUnsubscribeToken(unsubscribeToken string) *DeleteV1NuncSubscriptionsUnsubscribeTokenParams {
	o.SetUnsubscribeToken(unsubscribeToken)
	return o
}

// SetUnsubscribeToken adds the unsubscribeToken to the delete v1 nunc subscriptions unsubscribe token params
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) SetUnsubscribeToken(unsubscribeToken string) {
	o.UnsubscribeToken = unsubscribeToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1NuncSubscriptionsUnsubscribeTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param unsubscribe_token
	if err := r.SetPathParam("unsubscribe_token", o.UnsubscribeToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
