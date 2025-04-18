// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewPatchV1IncidentsIncidentIDTasksTaskIDParams creates a new PatchV1IncidentsIncidentIDTasksTaskIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IncidentsIncidentIDTasksTaskIDParams() *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	return &PatchV1IncidentsIncidentIDTasksTaskIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IncidentsIncidentIDTasksTaskIDParamsWithTimeout creates a new PatchV1IncidentsIncidentIDTasksTaskIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IncidentsIncidentIDTasksTaskIDParamsWithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	return &PatchV1IncidentsIncidentIDTasksTaskIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IncidentsIncidentIDTasksTaskIDParamsWithContext creates a new PatchV1IncidentsIncidentIDTasksTaskIDParams object
// with the ability to set a context for a request.
func NewPatchV1IncidentsIncidentIDTasksTaskIDParamsWithContext(ctx context.Context) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	return &PatchV1IncidentsIncidentIDTasksTaskIDParams{
		Context: ctx,
	}
}

// NewPatchV1IncidentsIncidentIDTasksTaskIDParamsWithHTTPClient creates a new PatchV1IncidentsIncidentIDTasksTaskIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IncidentsIncidentIDTasksTaskIDParamsWithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	return &PatchV1IncidentsIncidentIDTasksTaskIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1IncidentsIncidentIDTasksTaskIDParams contains all the parameters to send to the API endpoint

	for the patch v1 incidents incident Id tasks task Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1IncidentsIncidentIDTasksTaskIDParams struct {

	// IncidentID.
	IncidentID string

	// PatchV1IncidentsIncidentIDTasksTaskID.
	PatchV1IncidentsIncidentIDTasksTaskID *models.PatchV1IncidentsIncidentIDTasksTaskID

	// TaskID.
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 incidents incident Id tasks task Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WithDefaults() *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 incidents incident Id tasks task Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WithContext(ctx context.Context) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WithIncidentID(incidentID string) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithPatchV1IncidentsIncidentIDTasksTaskID adds the patchV1IncidentsIncidentIDTasksTaskID to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WithPatchV1IncidentsIncidentIDTasksTaskID(patchV1IncidentsIncidentIDTasksTaskID *models.PatchV1IncidentsIncidentIDTasksTaskID) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetPatchV1IncidentsIncidentIDTasksTaskID(patchV1IncidentsIncidentIDTasksTaskID)
	return o
}

// SetPatchV1IncidentsIncidentIDTasksTaskID adds the patchV1IncidentsIncidentIdTasksTaskId to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) SetPatchV1IncidentsIncidentIDTasksTaskID(patchV1IncidentsIncidentIDTasksTaskID *models.PatchV1IncidentsIncidentIDTasksTaskID) {
	o.PatchV1IncidentsIncidentIDTasksTaskID = patchV1IncidentsIncidentIDTasksTaskID
}

// WithTaskID adds the taskID to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WithTaskID(taskID string) *PatchV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the patch v1 incidents incident Id tasks task Id params
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IncidentsIncidentIDTasksTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}
	if o.PatchV1IncidentsIncidentIDTasksTaskID != nil {
		if err := r.SetBodyParam(o.PatchV1IncidentsIncidentIDTasksTaskID); err != nil {
			return err
		}
	}

	// path param task_id
	if err := r.SetPathParam("task_id", o.TaskID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
