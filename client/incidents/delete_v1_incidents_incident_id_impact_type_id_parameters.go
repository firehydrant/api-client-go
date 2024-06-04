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

// NewDeleteV1IncidentsIncidentIDImpactTypeIDParams creates a new DeleteV1IncidentsIncidentIDImpactTypeIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1IncidentsIncidentIDImpactTypeIDParams() *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	return &DeleteV1IncidentsIncidentIDImpactTypeIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1IncidentsIncidentIDImpactTypeIDParamsWithTimeout creates a new DeleteV1IncidentsIncidentIDImpactTypeIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1IncidentsIncidentIDImpactTypeIDParamsWithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	return &DeleteV1IncidentsIncidentIDImpactTypeIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1IncidentsIncidentIDImpactTypeIDParamsWithContext creates a new DeleteV1IncidentsIncidentIDImpactTypeIDParams object
// with the ability to set a context for a request.
func NewDeleteV1IncidentsIncidentIDImpactTypeIDParamsWithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	return &DeleteV1IncidentsIncidentIDImpactTypeIDParams{
		Context: ctx,
	}
}

// NewDeleteV1IncidentsIncidentIDImpactTypeIDParamsWithHTTPClient creates a new DeleteV1IncidentsIncidentIDImpactTypeIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1IncidentsIncidentIDImpactTypeIDParamsWithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	return &DeleteV1IncidentsIncidentIDImpactTypeIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1IncidentsIncidentIDImpactTypeIDParams contains all the parameters to send to the API endpoint

	for the delete v1 incidents incident Id impact type Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1IncidentsIncidentIDImpactTypeIDParams struct {

	// ID.
	ID string

	// IncidentID.
	IncidentID string

	// Type.
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 incidents incident Id impact type Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WithDefaults() *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 incidents incident Id impact type Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WithID(id string) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) SetID(id string) {
	o.ID = id
}

// WithIncidentID adds the incidentID to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WithIncidentID(incidentID string) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithType adds the typeVar to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WithType(typeVar string) *DeleteV1IncidentsIncidentIDImpactTypeIDParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the delete v1 incidents incident Id impact type Id params
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1IncidentsIncidentIDImpactTypeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
