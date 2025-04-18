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
)

// NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams creates a new GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams() *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	return &GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParamsWithTimeout creates a new GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams object
// with the ability to set a timeout on a request.
func NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParamsWithTimeout(timeout time.Duration) *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	return &GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams{
		timeout: timeout,
	}
}

// NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParamsWithContext creates a new GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams object
// with the ability to set a context for a request.
func NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParamsWithContext(ctx context.Context) *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	return &GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams{
		Context: ctx,
	}
}

// NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParamsWithHTTPClient creates a new GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RetrospectiveTemplatesRetrospectiveTemplateIDParamsWithHTTPClient(client *http.Client) *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	return &GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams{
		HTTPClient: client,
	}
}

/*
GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams contains all the parameters to send to the API endpoint

	for the get v1 retrospective templates retrospective template Id operation.

	Typically these are written to a http.Request.
*/
type GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams struct {

	// RetrospectiveTemplateID.
	RetrospectiveTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 retrospective templates retrospective template Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) WithDefaults() *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 retrospective templates retrospective template Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) WithTimeout(timeout time.Duration) *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) WithContext(ctx context.Context) *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) WithHTTPClient(client *http.Client) *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRetrospectiveTemplateID adds the retrospectiveTemplateID to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) WithRetrospectiveTemplateID(retrospectiveTemplateID string) *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams {
	o.SetRetrospectiveTemplateID(retrospectiveTemplateID)
	return o
}

// SetRetrospectiveTemplateID adds the retrospectiveTemplateId to the get v1 retrospective templates retrospective template Id params
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) SetRetrospectiveTemplateID(retrospectiveTemplateID string) {
	o.RetrospectiveTemplateID = retrospectiveTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RetrospectiveTemplatesRetrospectiveTemplateIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param retrospective_template_id
	if err := r.SetPathParam("retrospective_template_id", o.RetrospectiveTemplateID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
