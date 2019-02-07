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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1IncidentsIncidentIDParams creates a new GetV1IncidentsIncidentIDParams object
// with the default values initialized.
func NewGetV1IncidentsIncidentIDParams() *GetV1IncidentsIncidentIDParams {
	var ()
	return &GetV1IncidentsIncidentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDParamsWithTimeout creates a new GetV1IncidentsIncidentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1IncidentsIncidentIDParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDParams {
	var ()
	return &GetV1IncidentsIncidentIDParams{

		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDParamsWithContext creates a new GetV1IncidentsIncidentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1IncidentsIncidentIDParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDParams {
	var ()
	return &GetV1IncidentsIncidentIDParams{

		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1IncidentsIncidentIDParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDParams {
	var ()
	return &GetV1IncidentsIncidentIDParams{
		HTTPClient: client,
	}
}

/*GetV1IncidentsIncidentIDParams contains all the parameters to send to the API endpoint
for the get v1 incidents incident Id operation typically these are written to a http.Request
*/
type GetV1IncidentsIncidentIDParams struct {

	/*IncidentID*/
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id params
func (o *GetV1IncidentsIncidentIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
