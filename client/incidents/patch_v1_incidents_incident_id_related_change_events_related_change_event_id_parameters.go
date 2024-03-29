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

// NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams creates a new PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams() *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	return &PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParamsWithTimeout creates a new PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParamsWithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	return &PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParamsWithContext creates a new PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams object
// with the ability to set a context for a request.
func NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParamsWithContext(ctx context.Context) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	return &PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams{
		Context: ctx,
	}
}

// NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParamsWithHTTPClient creates a new PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParamsWithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	return &PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams contains all the parameters to send to the API endpoint

	for the patch v1 incidents incident Id related change events related change event Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams struct {

	// IncidentID.
	IncidentID string

	// PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID.
	PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID *models.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID

	// RelatedChangeEventID.
	RelatedChangeEventID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 incidents incident Id related change events related change event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WithDefaults() *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 incidents incident Id related change events related change event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WithContext(ctx context.Context) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WithIncidentID(incidentID string) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID adds the patchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WithPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID(patchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID *models.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	o.SetPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID(patchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID)
	return o
}

// SetPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID adds the patchV1IncidentsIncidentIdRelatedChangeEventsRelatedChangeEventId to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) SetPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID(patchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID *models.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID) {
	o.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID = patchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID
}

// WithRelatedChangeEventID adds the relatedChangeEventID to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WithRelatedChangeEventID(relatedChangeEventID string) *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams {
	o.SetRelatedChangeEventID(relatedChangeEventID)
	return o
}

// SetRelatedChangeEventID adds the relatedChangeEventId to the patch v1 incidents incident Id related change events related change event Id params
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) SetRelatedChangeEventID(relatedChangeEventID string) {
	o.RelatedChangeEventID = relatedChangeEventID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}
	if o.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID != nil {
		if err := r.SetBodyParam(o.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID); err != nil {
			return err
		}
	}

	// path param related_change_event_id
	if err := r.SetPathParam("related_change_event_id", o.RelatedChangeEventID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
