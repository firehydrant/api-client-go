// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewPostV1IncidentsIncidentIDRetrospectivesParams creates a new PostV1IncidentsIncidentIDRetrospectivesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDRetrospectivesParams() *PostV1IncidentsIncidentIDRetrospectivesParams {
	return &PostV1IncidentsIncidentIDRetrospectivesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDRetrospectivesParamsWithTimeout creates a new PostV1IncidentsIncidentIDRetrospectivesParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDRetrospectivesParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDRetrospectivesParams {
	return &PostV1IncidentsIncidentIDRetrospectivesParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDRetrospectivesParamsWithContext creates a new PostV1IncidentsIncidentIDRetrospectivesParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDRetrospectivesParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDRetrospectivesParams {
	return &PostV1IncidentsIncidentIDRetrospectivesParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDRetrospectivesParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDRetrospectivesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDRetrospectivesParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDRetrospectivesParams {
	return &PostV1IncidentsIncidentIDRetrospectivesParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDRetrospectivesParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id retrospectives operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDRetrospectivesParams struct {

	// IncidentID.
	IncidentID string

	/* RetrospectiveTemplateID.

	   The id of the retrospective template to apply.
	*/
	RetrospectiveTemplateID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id retrospectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) WithDefaults() *PostV1IncidentsIncidentIDRetrospectivesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id retrospectives params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDRetrospectivesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDRetrospectivesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDRetrospectivesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDRetrospectivesParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithRetrospectiveTemplateID adds the retrospectiveTemplateID to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) WithRetrospectiveTemplateID(retrospectiveTemplateID string) *PostV1IncidentsIncidentIDRetrospectivesParams {
	o.SetRetrospectiveTemplateID(retrospectiveTemplateID)
	return o
}

// SetRetrospectiveTemplateID adds the retrospectiveTemplateId to the post v1 incidents incident Id retrospectives params
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) SetRetrospectiveTemplateID(retrospectiveTemplateID string) {
	o.RetrospectiveTemplateID = retrospectiveTemplateID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDRetrospectivesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	// form param retrospective_template_id
	frRetrospectiveTemplateID := o.RetrospectiveTemplateID
	fRetrospectiveTemplateID := frRetrospectiveTemplateID
	if fRetrospectiveTemplateID != "" {
		if err := r.SetFormParam("retrospective_template_id", fRetrospectiveTemplateID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
