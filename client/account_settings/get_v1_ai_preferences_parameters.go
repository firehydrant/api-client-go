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

// NewGetV1AiPreferencesParams creates a new GetV1AiPreferencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AiPreferencesParams() *GetV1AiPreferencesParams {
	return &GetV1AiPreferencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AiPreferencesParamsWithTimeout creates a new GetV1AiPreferencesParams object
// with the ability to set a timeout on a request.
func NewGetV1AiPreferencesParamsWithTimeout(timeout time.Duration) *GetV1AiPreferencesParams {
	return &GetV1AiPreferencesParams{
		timeout: timeout,
	}
}

// NewGetV1AiPreferencesParamsWithContext creates a new GetV1AiPreferencesParams object
// with the ability to set a context for a request.
func NewGetV1AiPreferencesParamsWithContext(ctx context.Context) *GetV1AiPreferencesParams {
	return &GetV1AiPreferencesParams{
		Context: ctx,
	}
}

// NewGetV1AiPreferencesParamsWithHTTPClient creates a new GetV1AiPreferencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AiPreferencesParamsWithHTTPClient(client *http.Client) *GetV1AiPreferencesParams {
	return &GetV1AiPreferencesParams{
		HTTPClient: client,
	}
}

/*
GetV1AiPreferencesParams contains all the parameters to send to the API endpoint

	for the get v1 ai preferences operation.

	Typically these are written to a http.Request.
*/
type GetV1AiPreferencesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 ai preferences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AiPreferencesParams) WithDefaults() *GetV1AiPreferencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 ai preferences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AiPreferencesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 ai preferences params
func (o *GetV1AiPreferencesParams) WithTimeout(timeout time.Duration) *GetV1AiPreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 ai preferences params
func (o *GetV1AiPreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 ai preferences params
func (o *GetV1AiPreferencesParams) WithContext(ctx context.Context) *GetV1AiPreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 ai preferences params
func (o *GetV1AiPreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 ai preferences params
func (o *GetV1AiPreferencesParams) WithHTTPClient(client *http.Client) *GetV1AiPreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 ai preferences params
func (o *GetV1AiPreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AiPreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
