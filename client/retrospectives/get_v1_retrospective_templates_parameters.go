// Code generated by go-swagger; DO NOT EDIT.

package retrospectives

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

// NewGetV1RetrospectiveTemplatesParams creates a new GetV1RetrospectiveTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RetrospectiveTemplatesParams() *GetV1RetrospectiveTemplatesParams {
	return &GetV1RetrospectiveTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RetrospectiveTemplatesParamsWithTimeout creates a new GetV1RetrospectiveTemplatesParams object
// with the ability to set a timeout on a request.
func NewGetV1RetrospectiveTemplatesParamsWithTimeout(timeout time.Duration) *GetV1RetrospectiveTemplatesParams {
	return &GetV1RetrospectiveTemplatesParams{
		timeout: timeout,
	}
}

// NewGetV1RetrospectiveTemplatesParamsWithContext creates a new GetV1RetrospectiveTemplatesParams object
// with the ability to set a context for a request.
func NewGetV1RetrospectiveTemplatesParamsWithContext(ctx context.Context) *GetV1RetrospectiveTemplatesParams {
	return &GetV1RetrospectiveTemplatesParams{
		Context: ctx,
	}
}

// NewGetV1RetrospectiveTemplatesParamsWithHTTPClient creates a new GetV1RetrospectiveTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RetrospectiveTemplatesParamsWithHTTPClient(client *http.Client) *GetV1RetrospectiveTemplatesParams {
	return &GetV1RetrospectiveTemplatesParams{
		HTTPClient: client,
	}
}

/*
GetV1RetrospectiveTemplatesParams contains all the parameters to send to the API endpoint

	for the get v1 retrospective templates operation.

	Typically these are written to a http.Request.
*/
type GetV1RetrospectiveTemplatesParams struct {

	// ForIncident.
	ForIncident *string

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 retrospective templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RetrospectiveTemplatesParams) WithDefaults() *GetV1RetrospectiveTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 retrospective templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RetrospectiveTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) WithTimeout(timeout time.Duration) *GetV1RetrospectiveTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) WithContext(ctx context.Context) *GetV1RetrospectiveTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) WithHTTPClient(client *http.Client) *GetV1RetrospectiveTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForIncident adds the forIncident to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) WithForIncident(forIncident *string) *GetV1RetrospectiveTemplatesParams {
	o.SetForIncident(forIncident)
	return o
}

// SetForIncident adds the forIncident to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) SetForIncident(forIncident *string) {
	o.ForIncident = forIncident
}

// WithPage adds the page to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) WithPage(page *int32) *GetV1RetrospectiveTemplatesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) WithPerPage(perPage *int32) *GetV1RetrospectiveTemplatesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 retrospective templates params
func (o *GetV1RetrospectiveTemplatesParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RetrospectiveTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForIncident != nil {

		// query param for_incident
		var qrForIncident string

		if o.ForIncident != nil {
			qrForIncident = *o.ForIncident
		}
		qForIncident := qrForIncident
		if qForIncident != "" {

			if err := r.SetQueryParam("for_incident", qForIncident); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
