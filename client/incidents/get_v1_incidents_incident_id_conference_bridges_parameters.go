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

// NewGetV1IncidentsIncidentIDConferenceBridgesParams creates a new GetV1IncidentsIncidentIDConferenceBridgesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDConferenceBridgesParams() *GetV1IncidentsIncidentIDConferenceBridgesParams {
	return &GetV1IncidentsIncidentIDConferenceBridgesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDConferenceBridgesParamsWithTimeout creates a new GetV1IncidentsIncidentIDConferenceBridgesParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDConferenceBridgesParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDConferenceBridgesParams {
	return &GetV1IncidentsIncidentIDConferenceBridgesParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDConferenceBridgesParamsWithContext creates a new GetV1IncidentsIncidentIDConferenceBridgesParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDConferenceBridgesParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDConferenceBridgesParams {
	return &GetV1IncidentsIncidentIDConferenceBridgesParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDConferenceBridgesParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDConferenceBridgesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDConferenceBridgesParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDConferenceBridgesParams {
	return &GetV1IncidentsIncidentIDConferenceBridgesParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentsIncidentIDConferenceBridgesParams contains all the parameters to send to the API endpoint

	for the get v1 incidents incident Id conference bridges operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDConferenceBridgesParams struct {

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id conference bridges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) WithDefaults() *GetV1IncidentsIncidentIDConferenceBridgesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id conference bridges params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDConferenceBridgesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDConferenceBridgesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDConferenceBridgesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDConferenceBridgesParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id conference bridges params
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDConferenceBridgesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
