// Code generated by go-swagger; DO NOT EDIT.

package account_settings

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

// NewGetV1BootstrapParams creates a new GetV1BootstrapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1BootstrapParams() *GetV1BootstrapParams {
	return &GetV1BootstrapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1BootstrapParamsWithTimeout creates a new GetV1BootstrapParams object
// with the ability to set a timeout on a request.
func NewGetV1BootstrapParamsWithTimeout(timeout time.Duration) *GetV1BootstrapParams {
	return &GetV1BootstrapParams{
		timeout: timeout,
	}
}

// NewGetV1BootstrapParamsWithContext creates a new GetV1BootstrapParams object
// with the ability to set a context for a request.
func NewGetV1BootstrapParamsWithContext(ctx context.Context) *GetV1BootstrapParams {
	return &GetV1BootstrapParams{
		Context: ctx,
	}
}

// NewGetV1BootstrapParamsWithHTTPClient creates a new GetV1BootstrapParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1BootstrapParamsWithHTTPClient(client *http.Client) *GetV1BootstrapParams {
	return &GetV1BootstrapParams{
		HTTPClient: client,
	}
}

/*
GetV1BootstrapParams contains all the parameters to send to the API endpoint

	for the get v1 bootstrap operation.

	Typically these are written to a http.Request.
*/
type GetV1BootstrapParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 bootstrap params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1BootstrapParams) WithDefaults() *GetV1BootstrapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 bootstrap params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1BootstrapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 bootstrap params
func (o *GetV1BootstrapParams) WithTimeout(timeout time.Duration) *GetV1BootstrapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 bootstrap params
func (o *GetV1BootstrapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 bootstrap params
func (o *GetV1BootstrapParams) WithContext(ctx context.Context) *GetV1BootstrapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 bootstrap params
func (o *GetV1BootstrapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 bootstrap params
func (o *GetV1BootstrapParams) WithHTTPClient(client *http.Client) *GetV1BootstrapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 bootstrap params
func (o *GetV1BootstrapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1BootstrapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
