// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams creates a new DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams() *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	return &DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParamsWithTimeout creates a new DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParamsWithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	return &DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParamsWithContext creates a new DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams object
// with the ability to set a context for a request.
func NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParamsWithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	return &DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams{
		Context: ctx,
	}
}

// NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParamsWithHTTPClient creates a new DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParamsWithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	return &DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams contains all the parameters to send to the API endpoint

	for the delete v1 incidents incident Id alerts incident alert Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams struct {

	// IncidentAlertID.
	IncidentAlertID string

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 incidents incident Id alerts incident alert Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) WithDefaults() *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 incidents incident Id alerts incident alert Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) WithTimeout(timeout time.Duration) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) WithContext(ctx context.Context) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) WithHTTPClient(client *http.Client) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentAlertID adds the incidentAlertID to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) WithIncidentAlertID(incidentAlertID string) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	o.SetIncidentAlertID(incidentAlertID)
	return o
}

// SetIncidentAlertID adds the incidentAlertId to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) SetIncidentAlertID(incidentAlertID string) {
	o.IncidentAlertID = incidentAlertID
}

// WithIncidentID adds the incidentID to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) WithIncidentID(incidentID string) *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the delete v1 incidents incident Id alerts incident alert Id params
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_alert_id
	if err := r.SetPathParam("incident_alert_id", o.IncidentAlertID); err != nil {
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
