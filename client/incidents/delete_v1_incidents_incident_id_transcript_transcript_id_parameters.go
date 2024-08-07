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

// NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams creates a new DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams() *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	return &DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParamsWithTimeout creates a new DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParamsWithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	return &DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParamsWithContext creates a new DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams object
// with the ability to set a context for a request.
func NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParamsWithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	return &DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams{
		Context: ctx,
	}
}

// NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParamsWithHTTPClient creates a new DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1IncidentsIncidentIDTranscriptTranscriptIDParamsWithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	return &DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams contains all the parameters to send to the API endpoint

	for the delete v1 incidents incident Id transcript transcript Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams struct {

	// IncidentID.
	IncidentID string

	// TranscriptID.
	TranscriptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 incidents incident Id transcript transcript Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) WithDefaults() *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 incidents incident Id transcript transcript Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) WithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) WithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) WithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) WithIncidentID(incidentID string) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithTranscriptID adds the transcriptID to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) WithTranscriptID(transcriptID string) *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams {
	o.SetTranscriptID(transcriptID)
	return o
}

// SetTranscriptID adds the transcriptId to the delete v1 incidents incident Id transcript transcript Id params
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) SetTranscriptID(transcriptID string) {
	o.TranscriptID = transcriptID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1IncidentsIncidentIDTranscriptTranscriptIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	// path param transcript_id
	if err := r.SetPathParam("transcript_id", o.TranscriptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
