// Code generated by go-swagger; DO NOT EDIT.

package checklist_templates

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

// NewGetV1ChecklistTemplatesIDParams creates a new GetV1ChecklistTemplatesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ChecklistTemplatesIDParams() *GetV1ChecklistTemplatesIDParams {
	return &GetV1ChecklistTemplatesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ChecklistTemplatesIDParamsWithTimeout creates a new GetV1ChecklistTemplatesIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ChecklistTemplatesIDParamsWithTimeout(timeout time.Duration) *GetV1ChecklistTemplatesIDParams {
	return &GetV1ChecklistTemplatesIDParams{
		timeout: timeout,
	}
}

// NewGetV1ChecklistTemplatesIDParamsWithContext creates a new GetV1ChecklistTemplatesIDParams object
// with the ability to set a context for a request.
func NewGetV1ChecklistTemplatesIDParamsWithContext(ctx context.Context) *GetV1ChecklistTemplatesIDParams {
	return &GetV1ChecklistTemplatesIDParams{
		Context: ctx,
	}
}

// NewGetV1ChecklistTemplatesIDParamsWithHTTPClient creates a new GetV1ChecklistTemplatesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ChecklistTemplatesIDParamsWithHTTPClient(client *http.Client) *GetV1ChecklistTemplatesIDParams {
	return &GetV1ChecklistTemplatesIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ChecklistTemplatesIDParams contains all the parameters to send to the API endpoint

	for the get v1 checklist templates Id operation.

	Typically these are written to a http.Request.
*/
type GetV1ChecklistTemplatesIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 checklist templates Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ChecklistTemplatesIDParams) WithDefaults() *GetV1ChecklistTemplatesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 checklist templates Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ChecklistTemplatesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) WithTimeout(timeout time.Duration) *GetV1ChecklistTemplatesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) WithContext(ctx context.Context) *GetV1ChecklistTemplatesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) WithHTTPClient(client *http.Client) *GetV1ChecklistTemplatesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) WithID(id string) *GetV1ChecklistTemplatesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 checklist templates Id params
func (o *GetV1ChecklistTemplatesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ChecklistTemplatesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
