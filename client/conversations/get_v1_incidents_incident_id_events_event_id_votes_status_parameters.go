// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParams creates a new GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParams() *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	return &GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParamsWithTimeout creates a new GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	return &GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParamsWithContext creates a new GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	return &GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDEventsEventIDVotesStatusParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	return &GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams contains all the parameters to send to the API endpoint

	for the get v1 incidents incident Id events event Id votes status operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams struct {

	// EventID.
	EventID string

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id events event Id votes status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) WithDefaults() *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id events event Id votes status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEventID adds the eventID to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) WithEventID(eventID string) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) SetEventID(eventID string) {
	o.EventID = eventID
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id events event Id votes status params
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param event_id
	if err := r.SetPathParam("event_id", o.EventID); err != nil {
		return err
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
