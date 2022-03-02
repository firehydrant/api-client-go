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

// NewGetV1IncidentsIncidentIDRoleAssignmentsParams creates a new GetV1IncidentsIncidentIDRoleAssignmentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDRoleAssignmentsParams() *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	return &GetV1IncidentsIncidentIDRoleAssignmentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDRoleAssignmentsParamsWithTimeout creates a new GetV1IncidentsIncidentIDRoleAssignmentsParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDRoleAssignmentsParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	return &GetV1IncidentsIncidentIDRoleAssignmentsParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDRoleAssignmentsParamsWithContext creates a new GetV1IncidentsIncidentIDRoleAssignmentsParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDRoleAssignmentsParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	return &GetV1IncidentsIncidentIDRoleAssignmentsParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDRoleAssignmentsParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDRoleAssignmentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDRoleAssignmentsParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	return &GetV1IncidentsIncidentIDRoleAssignmentsParams{
		HTTPClient: client,
	}
}

/* GetV1IncidentsIncidentIDRoleAssignmentsParams contains all the parameters to send to the API endpoint
   for the get v1 incidents incident Id role assignments operation.

   Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDRoleAssignmentsParams struct {

	// IncidentID.
	IncidentID string

	/* Status.

	   Filter on status of the role assignment
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id role assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) WithDefaults() *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id role assignments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithStatus adds the status to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) WithStatus(status *string) *GetV1IncidentsIncidentIDRoleAssignmentsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get v1 incidents incident Id role assignments params
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDRoleAssignmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
