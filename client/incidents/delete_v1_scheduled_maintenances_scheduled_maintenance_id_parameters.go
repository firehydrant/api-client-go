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

// NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams creates a new DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams() *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithTimeout creates a new DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithTimeout(timeout time.Duration) *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithContext creates a new DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams object
// with the ability to set a context for a request.
func NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithContext(ctx context.Context) *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		Context: ctx,
	}
}

// NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithHTTPClient creates a new DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ScheduledMaintenancesScheduledMaintenanceIDParamsWithHTTPClient(client *http.Client) *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	return &DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams contains all the parameters to send to the API endpoint

	for the delete v1 scheduled maintenances scheduled maintenance Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams struct {

	// ScheduledMaintenanceID.
	ScheduledMaintenanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 scheduled maintenances scheduled maintenance Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithDefaults() *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 scheduled maintenances scheduled maintenance Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithTimeout(timeout time.Duration) *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithContext(ctx context.Context) *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithHTTPClient(client *http.Client) *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScheduledMaintenanceID adds the scheduledMaintenanceID to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) WithScheduledMaintenanceID(scheduledMaintenanceID string) *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams {
	o.SetScheduledMaintenanceID(scheduledMaintenanceID)
	return o
}

// SetScheduledMaintenanceID adds the scheduledMaintenanceId to the delete v1 scheduled maintenances scheduled maintenance Id params
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) SetScheduledMaintenanceID(scheduledMaintenanceID string) {
	o.ScheduledMaintenanceID = scheduledMaintenanceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scheduled_maintenance_id
	if err := r.SetPathParam("scheduled_maintenance_id", o.ScheduledMaintenanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
