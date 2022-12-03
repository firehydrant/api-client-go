// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetV1UsersParams creates a new GetV1UsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1UsersParams() *GetV1UsersParams {
	return &GetV1UsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1UsersParamsWithTimeout creates a new GetV1UsersParams object
// with the ability to set a timeout on a request.
func NewGetV1UsersParamsWithTimeout(timeout time.Duration) *GetV1UsersParams {
	return &GetV1UsersParams{
		timeout: timeout,
	}
}

// NewGetV1UsersParamsWithContext creates a new GetV1UsersParams object
// with the ability to set a context for a request.
func NewGetV1UsersParamsWithContext(ctx context.Context) *GetV1UsersParams {
	return &GetV1UsersParams{
		Context: ctx,
	}
}

// NewGetV1UsersParamsWithHTTPClient creates a new GetV1UsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1UsersParamsWithHTTPClient(client *http.Client) *GetV1UsersParams {
	return &GetV1UsersParams{
		HTTPClient: client,
	}
}

/*
GetV1UsersParams contains all the parameters to send to the API endpoint

	for the get v1 users operation.

	Typically these are written to a http.Request.
*/
type GetV1UsersParams struct {

	/* Name.

	   Text string of a query to filter users by name
	*/
	Name *string

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   Text string of a query to filter users by name or email
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UsersParams) WithDefaults() *GetV1UsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1UsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 users params
func (o *GetV1UsersParams) WithTimeout(timeout time.Duration) *GetV1UsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 users params
func (o *GetV1UsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 users params
func (o *GetV1UsersParams) WithContext(ctx context.Context) *GetV1UsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 users params
func (o *GetV1UsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 users params
func (o *GetV1UsersParams) WithHTTPClient(client *http.Client) *GetV1UsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 users params
func (o *GetV1UsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get v1 users params
func (o *GetV1UsersParams) WithName(name *string) *GetV1UsersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 users params
func (o *GetV1UsersParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the get v1 users params
func (o *GetV1UsersParams) WithPage(page *int32) *GetV1UsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 users params
func (o *GetV1UsersParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 users params
func (o *GetV1UsersParams) WithPerPage(perPage *int32) *GetV1UsersParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 users params
func (o *GetV1UsersParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 users params
func (o *GetV1UsersParams) WithQuery(query *string) *GetV1UsersParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 users params
func (o *GetV1UsersParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1UsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

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
