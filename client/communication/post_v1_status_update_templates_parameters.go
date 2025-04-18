// Code generated by go-swagger; DO NOT EDIT.

package communication

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

// NewPostV1StatusUpdateTemplatesParams creates a new PostV1StatusUpdateTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1StatusUpdateTemplatesParams() *PostV1StatusUpdateTemplatesParams {
	return &PostV1StatusUpdateTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1StatusUpdateTemplatesParamsWithTimeout creates a new PostV1StatusUpdateTemplatesParams object
// with the ability to set a timeout on a request.
func NewPostV1StatusUpdateTemplatesParamsWithTimeout(timeout time.Duration) *PostV1StatusUpdateTemplatesParams {
	return &PostV1StatusUpdateTemplatesParams{
		timeout: timeout,
	}
}

// NewPostV1StatusUpdateTemplatesParamsWithContext creates a new PostV1StatusUpdateTemplatesParams object
// with the ability to set a context for a request.
func NewPostV1StatusUpdateTemplatesParamsWithContext(ctx context.Context) *PostV1StatusUpdateTemplatesParams {
	return &PostV1StatusUpdateTemplatesParams{
		Context: ctx,
	}
}

// NewPostV1StatusUpdateTemplatesParamsWithHTTPClient creates a new PostV1StatusUpdateTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1StatusUpdateTemplatesParamsWithHTTPClient(client *http.Client) *PostV1StatusUpdateTemplatesParams {
	return &PostV1StatusUpdateTemplatesParams{
		HTTPClient: client,
	}
}

/*
PostV1StatusUpdateTemplatesParams contains all the parameters to send to the API endpoint

	for the post v1 status update templates operation.

	Typically these are written to a http.Request.
*/
type PostV1StatusUpdateTemplatesParams struct {

	// PostV1StatusUpdateTemplates.
	PostV1StatusUpdateTemplates *models.PostV1StatusUpdateTemplates

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 status update templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1StatusUpdateTemplatesParams) WithDefaults() *PostV1StatusUpdateTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 status update templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1StatusUpdateTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) WithTimeout(timeout time.Duration) *PostV1StatusUpdateTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) WithContext(ctx context.Context) *PostV1StatusUpdateTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) WithHTTPClient(client *http.Client) *PostV1StatusUpdateTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostV1StatusUpdateTemplates adds the postV1StatusUpdateTemplates to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) WithPostV1StatusUpdateTemplates(postV1StatusUpdateTemplates *models.PostV1StatusUpdateTemplates) *PostV1StatusUpdateTemplatesParams {
	o.SetPostV1StatusUpdateTemplates(postV1StatusUpdateTemplates)
	return o
}

// SetPostV1StatusUpdateTemplates adds the postV1StatusUpdateTemplates to the post v1 status update templates params
func (o *PostV1StatusUpdateTemplatesParams) SetPostV1StatusUpdateTemplates(postV1StatusUpdateTemplates *models.PostV1StatusUpdateTemplates) {
	o.PostV1StatusUpdateTemplates = postV1StatusUpdateTemplates
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1StatusUpdateTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostV1StatusUpdateTemplates != nil {
		if err := r.SetBodyParam(o.PostV1StatusUpdateTemplates); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
