// Code generated by go-swagger; DO NOT EDIT.

package incident_settings

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

// NewGetV1SeverityMatrixImpactsParams creates a new GetV1SeverityMatrixImpactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SeverityMatrixImpactsParams() *GetV1SeverityMatrixImpactsParams {
	return &GetV1SeverityMatrixImpactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SeverityMatrixImpactsParamsWithTimeout creates a new GetV1SeverityMatrixImpactsParams object
// with the ability to set a timeout on a request.
func NewGetV1SeverityMatrixImpactsParamsWithTimeout(timeout time.Duration) *GetV1SeverityMatrixImpactsParams {
	return &GetV1SeverityMatrixImpactsParams{
		timeout: timeout,
	}
}

// NewGetV1SeverityMatrixImpactsParamsWithContext creates a new GetV1SeverityMatrixImpactsParams object
// with the ability to set a context for a request.
func NewGetV1SeverityMatrixImpactsParamsWithContext(ctx context.Context) *GetV1SeverityMatrixImpactsParams {
	return &GetV1SeverityMatrixImpactsParams{
		Context: ctx,
	}
}

// NewGetV1SeverityMatrixImpactsParamsWithHTTPClient creates a new GetV1SeverityMatrixImpactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SeverityMatrixImpactsParamsWithHTTPClient(client *http.Client) *GetV1SeverityMatrixImpactsParams {
	return &GetV1SeverityMatrixImpactsParams{
		HTTPClient: client,
	}
}

/*
GetV1SeverityMatrixImpactsParams contains all the parameters to send to the API endpoint

	for the get v1 severity matrix impacts operation.

	Typically these are written to a http.Request.
*/
type GetV1SeverityMatrixImpactsParams struct {

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

// WithDefaults hydrates default values in the get v1 severity matrix impacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeverityMatrixImpactsParams) WithDefaults() *GetV1SeverityMatrixImpactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 severity matrix impacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeverityMatrixImpactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) WithTimeout(timeout time.Duration) *GetV1SeverityMatrixImpactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) WithContext(ctx context.Context) *GetV1SeverityMatrixImpactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) WithHTTPClient(client *http.Client) *GetV1SeverityMatrixImpactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) WithPage(page *int32) *GetV1SeverityMatrixImpactsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) WithPerPage(perPage *int32) *GetV1SeverityMatrixImpactsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 severity matrix impacts params
func (o *GetV1SeverityMatrixImpactsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SeverityMatrixImpactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
