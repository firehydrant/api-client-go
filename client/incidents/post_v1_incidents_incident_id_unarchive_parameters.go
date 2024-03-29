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

// NewPostV1IncidentsIncidentIDUnarchiveParams creates a new PostV1IncidentsIncidentIDUnarchiveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDUnarchiveParams() *PostV1IncidentsIncidentIDUnarchiveParams {
	return &PostV1IncidentsIncidentIDUnarchiveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDUnarchiveParamsWithTimeout creates a new PostV1IncidentsIncidentIDUnarchiveParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDUnarchiveParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDUnarchiveParams {
	return &PostV1IncidentsIncidentIDUnarchiveParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDUnarchiveParamsWithContext creates a new PostV1IncidentsIncidentIDUnarchiveParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDUnarchiveParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDUnarchiveParams {
	return &PostV1IncidentsIncidentIDUnarchiveParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDUnarchiveParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDUnarchiveParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDUnarchiveParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDUnarchiveParams {
	return &PostV1IncidentsIncidentIDUnarchiveParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDUnarchiveParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id unarchive operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDUnarchiveParams struct {

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id unarchive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDUnarchiveParams) WithDefaults() *PostV1IncidentsIncidentIDUnarchiveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id unarchive params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDUnarchiveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDUnarchiveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDUnarchiveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDUnarchiveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDUnarchiveParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id unarchive params
func (o *PostV1IncidentsIncidentIDUnarchiveParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDUnarchiveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
