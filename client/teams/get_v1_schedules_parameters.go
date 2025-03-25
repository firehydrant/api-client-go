// Code generated by go-swagger; DO NOT EDIT.

package teams

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
	"github.com/go-openapi/swag"
)

// NewGetV1SchedulesParams creates a new GetV1SchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SchedulesParams() *GetV1SchedulesParams {
	return &GetV1SchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SchedulesParamsWithTimeout creates a new GetV1SchedulesParams object
// with the ability to set a timeout on a request.
func NewGetV1SchedulesParamsWithTimeout(timeout time.Duration) *GetV1SchedulesParams {
	return &GetV1SchedulesParams{
		timeout: timeout,
	}
}

// NewGetV1SchedulesParamsWithContext creates a new GetV1SchedulesParams object
// with the ability to set a context for a request.
func NewGetV1SchedulesParamsWithContext(ctx context.Context) *GetV1SchedulesParams {
	return &GetV1SchedulesParams{
		Context: ctx,
	}
}

// NewGetV1SchedulesParamsWithHTTPClient creates a new GetV1SchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SchedulesParamsWithHTTPClient(client *http.Client) *GetV1SchedulesParams {
	return &GetV1SchedulesParams{
		HTTPClient: client,
	}
}

/*
GetV1SchedulesParams contains all the parameters to send to the API endpoint

	for the get v1 schedules operation.

	Typically these are written to a http.Request.
*/
type GetV1SchedulesParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   Filter schedules with a query on their name
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SchedulesParams) WithDefaults() *GetV1SchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 schedules params
func (o *GetV1SchedulesParams) WithTimeout(timeout time.Duration) *GetV1SchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 schedules params
func (o *GetV1SchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 schedules params
func (o *GetV1SchedulesParams) WithContext(ctx context.Context) *GetV1SchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 schedules params
func (o *GetV1SchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 schedules params
func (o *GetV1SchedulesParams) WithHTTPClient(client *http.Client) *GetV1SchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 schedules params
func (o *GetV1SchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 schedules params
func (o *GetV1SchedulesParams) WithPage(page *int32) *GetV1SchedulesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 schedules params
func (o *GetV1SchedulesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 schedules params
func (o *GetV1SchedulesParams) WithPerPage(perPage *int32) *GetV1SchedulesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 schedules params
func (o *GetV1SchedulesParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 schedules params
func (o *GetV1SchedulesParams) WithQuery(query *string) *GetV1SchedulesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 schedules params
func (o *GetV1SchedulesParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

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
