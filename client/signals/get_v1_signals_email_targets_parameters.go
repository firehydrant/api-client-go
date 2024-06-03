// Code generated by go-swagger; DO NOT EDIT.

package signals

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

// NewGetV1SignalsEmailTargetsParams creates a new GetV1SignalsEmailTargetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SignalsEmailTargetsParams() *GetV1SignalsEmailTargetsParams {
	return &GetV1SignalsEmailTargetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SignalsEmailTargetsParamsWithTimeout creates a new GetV1SignalsEmailTargetsParams object
// with the ability to set a timeout on a request.
func NewGetV1SignalsEmailTargetsParamsWithTimeout(timeout time.Duration) *GetV1SignalsEmailTargetsParams {
	return &GetV1SignalsEmailTargetsParams{
		timeout: timeout,
	}
}

// NewGetV1SignalsEmailTargetsParamsWithContext creates a new GetV1SignalsEmailTargetsParams object
// with the ability to set a context for a request.
func NewGetV1SignalsEmailTargetsParamsWithContext(ctx context.Context) *GetV1SignalsEmailTargetsParams {
	return &GetV1SignalsEmailTargetsParams{
		Context: ctx,
	}
}

// NewGetV1SignalsEmailTargetsParamsWithHTTPClient creates a new GetV1SignalsEmailTargetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SignalsEmailTargetsParamsWithHTTPClient(client *http.Client) *GetV1SignalsEmailTargetsParams {
	return &GetV1SignalsEmailTargetsParams{
		HTTPClient: client,
	}
}

/*
GetV1SignalsEmailTargetsParams contains all the parameters to send to the API endpoint

	for the get v1 signals email targets operation.

	Typically these are written to a http.Request.
*/
type GetV1SignalsEmailTargetsParams struct {

	/* Query.

	   A query string to search the list of targets by.
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 signals email targets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsEmailTargetsParams) WithDefaults() *GetV1SignalsEmailTargetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 signals email targets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsEmailTargetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) WithTimeout(timeout time.Duration) *GetV1SignalsEmailTargetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) WithContext(ctx context.Context) *GetV1SignalsEmailTargetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) WithHTTPClient(client *http.Client) *GetV1SignalsEmailTargetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) WithQuery(query *string) *GetV1SignalsEmailTargetsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 signals email targets params
func (o *GetV1SignalsEmailTargetsParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SignalsEmailTargetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
