// Code generated by go-swagger; DO NOT EDIT.

package incident_types

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

// NewGetV1IncidentTypesParams creates a new GetV1IncidentTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentTypesParams() *GetV1IncidentTypesParams {
	return &GetV1IncidentTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentTypesParamsWithTimeout creates a new GetV1IncidentTypesParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentTypesParamsWithTimeout(timeout time.Duration) *GetV1IncidentTypesParams {
	return &GetV1IncidentTypesParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentTypesParamsWithContext creates a new GetV1IncidentTypesParams object
// with the ability to set a context for a request.
func NewGetV1IncidentTypesParamsWithContext(ctx context.Context) *GetV1IncidentTypesParams {
	return &GetV1IncidentTypesParams{
		Context: ctx,
	}
}

// NewGetV1IncidentTypesParamsWithHTTPClient creates a new GetV1IncidentTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentTypesParamsWithHTTPClient(client *http.Client) *GetV1IncidentTypesParams {
	return &GetV1IncidentTypesParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentTypesParams contains all the parameters to send to the API endpoint

	for the get v1 incident types operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentTypesParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   A query to search incident types by their name
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incident types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentTypesParams) WithDefaults() *GetV1IncidentTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incident types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incident types params
func (o *GetV1IncidentTypesParams) WithTimeout(timeout time.Duration) *GetV1IncidentTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incident types params
func (o *GetV1IncidentTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incident types params
func (o *GetV1IncidentTypesParams) WithContext(ctx context.Context) *GetV1IncidentTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incident types params
func (o *GetV1IncidentTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incident types params
func (o *GetV1IncidentTypesParams) WithHTTPClient(client *http.Client) *GetV1IncidentTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incident types params
func (o *GetV1IncidentTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 incident types params
func (o *GetV1IncidentTypesParams) WithPage(page *int32) *GetV1IncidentTypesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 incident types params
func (o *GetV1IncidentTypesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 incident types params
func (o *GetV1IncidentTypesParams) WithPerPage(perPage *int32) *GetV1IncidentTypesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 incident types params
func (o *GetV1IncidentTypesParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 incident types params
func (o *GetV1IncidentTypesParams) WithQuery(query *string) *GetV1IncidentTypesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 incident types params
func (o *GetV1IncidentTypesParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
