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

	"github.com/firehydrant/api-client-go/models"
)

// NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParams creates a new PutV1IncidentsIncidentIDMilestonesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParams() *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	return &PutV1IncidentsIncidentIDMilestonesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParamsWithTimeout creates a new PutV1IncidentsIncidentIDMilestonesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParamsWithTimeout(timeout time.Duration) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	return &PutV1IncidentsIncidentIDMilestonesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParamsWithContext creates a new PutV1IncidentsIncidentIDMilestonesBulkUpdateParams object
// with the ability to set a context for a request.
func NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParamsWithContext(ctx context.Context) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	return &PutV1IncidentsIncidentIDMilestonesBulkUpdateParams{
		Context: ctx,
	}
}

// NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParamsWithHTTPClient creates a new PutV1IncidentsIncidentIDMilestonesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1IncidentsIncidentIDMilestonesBulkUpdateParamsWithHTTPClient(client *http.Client) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	return &PutV1IncidentsIncidentIDMilestonesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
PutV1IncidentsIncidentIDMilestonesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the put v1 incidents incident Id milestones bulk update operation.

	Typically these are written to a http.Request.
*/
type PutV1IncidentsIncidentIDMilestonesBulkUpdateParams struct {

	// V1IncidentsIncidentIDMilestonesBulkUpdate.
	V1IncidentsIncidentIDMilestonesBulkUpdate *models.PutV1IncidentsIncidentIDMilestonesBulkUpdate

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 incidents incident Id milestones bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) WithDefaults() *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 incidents incident Id milestones bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) WithTimeout(timeout time.Duration) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) WithContext(ctx context.Context) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) WithHTTPClient(client *http.Client) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IncidentsIncidentIDMilestonesBulkUpdate adds the v1IncidentsIncidentIDMilestonesBulkUpdate to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) WithV1IncidentsIncidentIDMilestonesBulkUpdate(v1IncidentsIncidentIDMilestonesBulkUpdate *models.PutV1IncidentsIncidentIDMilestonesBulkUpdate) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	o.SetV1IncidentsIncidentIDMilestonesBulkUpdate(v1IncidentsIncidentIDMilestonesBulkUpdate)
	return o
}

// SetV1IncidentsIncidentIDMilestonesBulkUpdate adds the v1IncidentsIncidentIdMilestonesBulkUpdate to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) SetV1IncidentsIncidentIDMilestonesBulkUpdate(v1IncidentsIncidentIDMilestonesBulkUpdate *models.PutV1IncidentsIncidentIDMilestonesBulkUpdate) {
	o.V1IncidentsIncidentIDMilestonesBulkUpdate = v1IncidentsIncidentIDMilestonesBulkUpdate
}

// WithIncidentID adds the incidentID to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) WithIncidentID(incidentID string) *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the put v1 incidents incident Id milestones bulk update params
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1IncidentsIncidentIDMilestonesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1IncidentsIncidentIDMilestonesBulkUpdate != nil {
		if err := r.SetBodyParam(o.V1IncidentsIncidentIDMilestonesBulkUpdate); err != nil {
			return err
		}
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
