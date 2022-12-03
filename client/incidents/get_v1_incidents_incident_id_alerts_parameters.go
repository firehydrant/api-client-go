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

// NewGetV1IncidentsIncidentIDAlertsParams creates a new GetV1IncidentsIncidentIDAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDAlertsParams() *GetV1IncidentsIncidentIDAlertsParams {
	return &GetV1IncidentsIncidentIDAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDAlertsParamsWithTimeout creates a new GetV1IncidentsIncidentIDAlertsParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDAlertsParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDAlertsParams {
	return &GetV1IncidentsIncidentIDAlertsParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDAlertsParamsWithContext creates a new GetV1IncidentsIncidentIDAlertsParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDAlertsParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDAlertsParams {
	return &GetV1IncidentsIncidentIDAlertsParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDAlertsParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDAlertsParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDAlertsParams {
	return &GetV1IncidentsIncidentIDAlertsParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentsIncidentIDAlertsParams contains all the parameters to send to the API endpoint

	for the get v1 incidents incident Id alerts operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDAlertsParams struct {

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDAlertsParams) WithDefaults() *GetV1IncidentsIncidentIDAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDAlertsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id alerts params
func (o *GetV1IncidentsIncidentIDAlertsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
