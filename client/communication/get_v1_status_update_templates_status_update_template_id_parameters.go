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
)

// NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams creates a new GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams() *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	return &GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParamsWithTimeout creates a new GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams object
// with the ability to set a timeout on a request.
func NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParamsWithTimeout(timeout time.Duration) *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	return &GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams{
		timeout: timeout,
	}
}

// NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParamsWithContext creates a new GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams object
// with the ability to set a context for a request.
func NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParamsWithContext(ctx context.Context) *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	return &GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams{
		Context: ctx,
	}
}

// NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParamsWithHTTPClient creates a new GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1StatusUpdateTemplatesStatusUpdateTemplateIDParamsWithHTTPClient(client *http.Client) *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	return &GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams{
		HTTPClient: client,
	}
}

/*
GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams contains all the parameters to send to the API endpoint

	for the get v1 status update templates status update template Id operation.

	Typically these are written to a http.Request.
*/
type GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams struct {

	// StatusUpdateTemplateID.
	StatusUpdateTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 status update templates status update template Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) WithDefaults() *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 status update templates status update template Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) WithTimeout(timeout time.Duration) *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) WithContext(ctx context.Context) *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) WithHTTPClient(client *http.Client) *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStatusUpdateTemplateID adds the statusUpdateTemplateID to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) WithStatusUpdateTemplateID(statusUpdateTemplateID string) *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams {
	o.SetStatusUpdateTemplateID(statusUpdateTemplateID)
	return o
}

// SetStatusUpdateTemplateID adds the statusUpdateTemplateId to the get v1 status update templates status update template Id params
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) SetStatusUpdateTemplateID(statusUpdateTemplateID string) {
	o.StatusUpdateTemplateID = statusUpdateTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1StatusUpdateTemplatesStatusUpdateTemplateIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param status_update_template_id
	if err := r.SetPathParam("status_update_template_id", o.StatusUpdateTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
