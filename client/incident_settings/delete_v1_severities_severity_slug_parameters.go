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
)

// NewDeleteV1SeveritiesSeveritySlugParams creates a new DeleteV1SeveritiesSeveritySlugParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1SeveritiesSeveritySlugParams() *DeleteV1SeveritiesSeveritySlugParams {
	return &DeleteV1SeveritiesSeveritySlugParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1SeveritiesSeveritySlugParamsWithTimeout creates a new DeleteV1SeveritiesSeveritySlugParams object
// with the ability to set a timeout on a request.
func NewDeleteV1SeveritiesSeveritySlugParamsWithTimeout(timeout time.Duration) *DeleteV1SeveritiesSeveritySlugParams {
	return &DeleteV1SeveritiesSeveritySlugParams{
		timeout: timeout,
	}
}

// NewDeleteV1SeveritiesSeveritySlugParamsWithContext creates a new DeleteV1SeveritiesSeveritySlugParams object
// with the ability to set a context for a request.
func NewDeleteV1SeveritiesSeveritySlugParamsWithContext(ctx context.Context) *DeleteV1SeveritiesSeveritySlugParams {
	return &DeleteV1SeveritiesSeveritySlugParams{
		Context: ctx,
	}
}

// NewDeleteV1SeveritiesSeveritySlugParamsWithHTTPClient creates a new DeleteV1SeveritiesSeveritySlugParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1SeveritiesSeveritySlugParamsWithHTTPClient(client *http.Client) *DeleteV1SeveritiesSeveritySlugParams {
	return &DeleteV1SeveritiesSeveritySlugParams{
		HTTPClient: client,
	}
}

/*
DeleteV1SeveritiesSeveritySlugParams contains all the parameters to send to the API endpoint

	for the delete v1 severities severity slug operation.

	Typically these are written to a http.Request.
*/
type DeleteV1SeveritiesSeveritySlugParams struct {

	// SeveritySlug.
	SeveritySlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 severities severity slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SeveritiesSeveritySlugParams) WithDefaults() *DeleteV1SeveritiesSeveritySlugParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 severities severity slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SeveritiesSeveritySlugParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) WithTimeout(timeout time.Duration) *DeleteV1SeveritiesSeveritySlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) WithContext(ctx context.Context) *DeleteV1SeveritiesSeveritySlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) WithHTTPClient(client *http.Client) *DeleteV1SeveritiesSeveritySlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSeveritySlug adds the severitySlug to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) WithSeveritySlug(severitySlug string) *DeleteV1SeveritiesSeveritySlugParams {
	o.SetSeveritySlug(severitySlug)
	return o
}

// SetSeveritySlug adds the severitySlug to the delete v1 severities severity slug params
func (o *DeleteV1SeveritiesSeveritySlugParams) SetSeveritySlug(severitySlug string) {
	o.SeveritySlug = severitySlug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1SeveritiesSeveritySlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param severity_slug
	if err := r.SetPathParam("severity_slug", o.SeveritySlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
