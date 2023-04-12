// Code generated by go-swagger; DO NOT EDIT.

package scheduled_maintenances

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

// NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParams creates a new PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParams() *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithTimeout creates a new PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithTimeout(timeout time.Duration) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		timeout: timeout,
	}
}

// NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithContext creates a new PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams object
// with the ability to set a context for a request.
func NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithContext(ctx context.Context) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		Context: ctx,
	}
}

// NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithHTTPClient creates a new PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithHTTPClient(client *http.Client) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams contains all the parameters to send to the API endpoint

	for the patch v1 scheduled maintenances scheduled maintenance Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams struct {

	// PatchV1ScheduledMaintenancesScheduledMaintenanceID.
	PatchV1ScheduledMaintenancesScheduledMaintenanceID *models.PatchV1ScheduledMaintenancesScheduledMaintenanceID

	// ScheduledMaintenanceID.
	ScheduledMaintenanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 scheduled maintenances scheduled maintenance Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithDefaults() *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 scheduled maintenances scheduled maintenance Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithTimeout(timeout time.Duration) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithContext(ctx context.Context) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithHTTPClient(client *http.Client) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchV1ScheduledMaintenancesScheduledMaintenanceID adds the patchV1ScheduledMaintenancesScheduledMaintenanceID to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithPatchV1ScheduledMaintenancesScheduledMaintenanceID(patchV1ScheduledMaintenancesScheduledMaintenanceID *models.PatchV1ScheduledMaintenancesScheduledMaintenanceID) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetPatchV1ScheduledMaintenancesScheduledMaintenanceID(patchV1ScheduledMaintenancesScheduledMaintenanceID)
	return o
}

// SetPatchV1ScheduledMaintenancesScheduledMaintenanceID adds the patchV1ScheduledMaintenancesScheduledMaintenanceId to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetPatchV1ScheduledMaintenancesScheduledMaintenanceID(patchV1ScheduledMaintenancesScheduledMaintenanceID *models.PatchV1ScheduledMaintenancesScheduledMaintenanceID) {
	o.PatchV1ScheduledMaintenancesScheduledMaintenanceID = patchV1ScheduledMaintenancesScheduledMaintenanceID
}

// WithScheduledMaintenanceID adds the scheduledMaintenanceID to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithScheduledMaintenanceID(scheduledMaintenanceID string) *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetScheduledMaintenanceID(scheduledMaintenanceID)
	return o
}

// SetScheduledMaintenanceID adds the scheduledMaintenanceId to the patch v1 scheduled maintenances scheduled maintenance Id params
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetScheduledMaintenanceID(scheduledMaintenanceID string) {
	o.ScheduledMaintenanceID = scheduledMaintenanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ScheduledMaintenancesScheduledMaintenanceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PatchV1ScheduledMaintenancesScheduledMaintenanceID != nil {
		if err := r.SetBodyParam(o.PatchV1ScheduledMaintenancesScheduledMaintenanceID); err != nil {
			return err
		}
	}

	// path param scheduled_maintenance_id
	if err := r.SetPathParam("scheduled_maintenance_id", o.ScheduledMaintenanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
