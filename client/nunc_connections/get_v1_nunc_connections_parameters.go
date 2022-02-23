// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

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

// NewGetV1NuncConnectionsParams creates a new GetV1NuncConnectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1NuncConnectionsParams() *GetV1NuncConnectionsParams {
	return &GetV1NuncConnectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1NuncConnectionsParamsWithTimeout creates a new GetV1NuncConnectionsParams object
// with the ability to set a timeout on a request.
func NewGetV1NuncConnectionsParamsWithTimeout(timeout time.Duration) *GetV1NuncConnectionsParams {
	return &GetV1NuncConnectionsParams{
		timeout: timeout,
	}
}

// NewGetV1NuncConnectionsParamsWithContext creates a new GetV1NuncConnectionsParams object
// with the ability to set a context for a request.
func NewGetV1NuncConnectionsParamsWithContext(ctx context.Context) *GetV1NuncConnectionsParams {
	return &GetV1NuncConnectionsParams{
		Context: ctx,
	}
}

// NewGetV1NuncConnectionsParamsWithHTTPClient creates a new GetV1NuncConnectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1NuncConnectionsParamsWithHTTPClient(client *http.Client) *GetV1NuncConnectionsParams {
	return &GetV1NuncConnectionsParams{
		HTTPClient: client,
	}
}

/* GetV1NuncConnectionsParams contains all the parameters to send to the API endpoint
   for the get v1 nunc connections operation.

   Typically these are written to a http.Request.
*/
type GetV1NuncConnectionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 nunc connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1NuncConnectionsParams) WithDefaults() *GetV1NuncConnectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 nunc connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1NuncConnectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 nunc connections params
func (o *GetV1NuncConnectionsParams) WithTimeout(timeout time.Duration) *GetV1NuncConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 nunc connections params
func (o *GetV1NuncConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 nunc connections params
func (o *GetV1NuncConnectionsParams) WithContext(ctx context.Context) *GetV1NuncConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 nunc connections params
func (o *GetV1NuncConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 nunc connections params
func (o *GetV1NuncConnectionsParams) WithHTTPClient(client *http.Client) *GetV1NuncConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 nunc connections params
func (o *GetV1NuncConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1NuncConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
