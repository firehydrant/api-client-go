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
)

// NewGetV1IncidentsIncidentIDTasksTaskIDParams creates a new GetV1IncidentsIncidentIDTasksTaskIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDTasksTaskIDParams() *GetV1IncidentsIncidentIDTasksTaskIDParams {
	return &GetV1IncidentsIncidentIDTasksTaskIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDTasksTaskIDParamsWithTimeout creates a new GetV1IncidentsIncidentIDTasksTaskIDParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDTasksTaskIDParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	return &GetV1IncidentsIncidentIDTasksTaskIDParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDTasksTaskIDParamsWithContext creates a new GetV1IncidentsIncidentIDTasksTaskIDParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDTasksTaskIDParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	return &GetV1IncidentsIncidentIDTasksTaskIDParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDTasksTaskIDParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDTasksTaskIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDTasksTaskIDParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	return &GetV1IncidentsIncidentIDTasksTaskIDParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentsIncidentIDTasksTaskIDParams contains all the parameters to send to the API endpoint

	for the get v1 incidents incident Id tasks task Id operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDTasksTaskIDParams struct {

	// IncidentID.
	IncidentID string

	// TaskID.
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id tasks task Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) WithDefaults() *GetV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id tasks task Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithTaskID adds the taskID to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) WithTaskID(taskID string) *GetV1IncidentsIncidentIDTasksTaskIDParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the get v1 incidents incident Id tasks task Id params
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDTasksTaskIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
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
