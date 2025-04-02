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

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1SeveritiesParams creates a new PostV1SeveritiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1SeveritiesParams() *PostV1SeveritiesParams {
	return &PostV1SeveritiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1SeveritiesParamsWithTimeout creates a new PostV1SeveritiesParams object
// with the ability to set a timeout on a request.
func NewPostV1SeveritiesParamsWithTimeout(timeout time.Duration) *PostV1SeveritiesParams {
	return &PostV1SeveritiesParams{
		timeout: timeout,
	}
}

// NewPostV1SeveritiesParamsWithContext creates a new PostV1SeveritiesParams object
// with the ability to set a context for a request.
func NewPostV1SeveritiesParamsWithContext(ctx context.Context) *PostV1SeveritiesParams {
	return &PostV1SeveritiesParams{
		Context: ctx,
	}
}

// NewPostV1SeveritiesParamsWithHTTPClient creates a new PostV1SeveritiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1SeveritiesParamsWithHTTPClient(client *http.Client) *PostV1SeveritiesParams {
	return &PostV1SeveritiesParams{
		HTTPClient: client,
	}
}

/*
PostV1SeveritiesParams contains all the parameters to send to the API endpoint

	for the post v1 severities operation.

	Typically these are written to a http.Request.
*/
type PostV1SeveritiesParams struct {

	// PostV1Severities.
	PostV1Severities *models.PostV1Severities

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 severities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SeveritiesParams) WithDefaults() *PostV1SeveritiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 severities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SeveritiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 severities params
func (o *PostV1SeveritiesParams) WithTimeout(timeout time.Duration) *PostV1SeveritiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 severities params
func (o *PostV1SeveritiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 severities params
func (o *PostV1SeveritiesParams) WithContext(ctx context.Context) *PostV1SeveritiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 severities params
func (o *PostV1SeveritiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 severities params
func (o *PostV1SeveritiesParams) WithHTTPClient(client *http.Client) *PostV1SeveritiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 severities params
func (o *PostV1SeveritiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostV1Severities adds the postV1Severities to the post v1 severities params
func (o *PostV1SeveritiesParams) WithPostV1Severities(postV1Severities *models.PostV1Severities) *PostV1SeveritiesParams {
	o.SetPostV1Severities(postV1Severities)
	return o
}

// SetPostV1Severities adds the postV1Severities to the post v1 severities params
func (o *PostV1SeveritiesParams) SetPostV1Severities(postV1Severities *models.PostV1Severities) {
	o.PostV1Severities = postV1Severities
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1SeveritiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostV1Severities != nil {
		if err := r.SetBodyParam(o.PostV1Severities); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
