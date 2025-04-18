// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewDeleteV1ChecklistTemplatesIDParams creates a new DeleteV1ChecklistTemplatesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ChecklistTemplatesIDParams() *DeleteV1ChecklistTemplatesIDParams {
	return &DeleteV1ChecklistTemplatesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ChecklistTemplatesIDParamsWithTimeout creates a new DeleteV1ChecklistTemplatesIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ChecklistTemplatesIDParamsWithTimeout(timeout time.Duration) *DeleteV1ChecklistTemplatesIDParams {
	return &DeleteV1ChecklistTemplatesIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1ChecklistTemplatesIDParamsWithContext creates a new DeleteV1ChecklistTemplatesIDParams object
// with the ability to set a context for a request.
func NewDeleteV1ChecklistTemplatesIDParamsWithContext(ctx context.Context) *DeleteV1ChecklistTemplatesIDParams {
	return &DeleteV1ChecklistTemplatesIDParams{
		Context: ctx,
	}
}

// NewDeleteV1ChecklistTemplatesIDParamsWithHTTPClient creates a new DeleteV1ChecklistTemplatesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ChecklistTemplatesIDParamsWithHTTPClient(client *http.Client) *DeleteV1ChecklistTemplatesIDParams {
	return &DeleteV1ChecklistTemplatesIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ChecklistTemplatesIDParams contains all the parameters to send to the API endpoint

	for the delete v1 checklist templates Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ChecklistTemplatesIDParams struct {

	/* ID.

	   Checklist Template UUID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 checklist templates Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ChecklistTemplatesIDParams) WithDefaults() *DeleteV1ChecklistTemplatesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 checklist templates Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ChecklistTemplatesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) WithTimeout(timeout time.Duration) *DeleteV1ChecklistTemplatesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) WithContext(ctx context.Context) *DeleteV1ChecklistTemplatesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) WithHTTPClient(client *http.Client) *DeleteV1ChecklistTemplatesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) WithID(id string) *DeleteV1ChecklistTemplatesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 checklist templates Id params
func (o *DeleteV1ChecklistTemplatesIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ChecklistTemplatesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
