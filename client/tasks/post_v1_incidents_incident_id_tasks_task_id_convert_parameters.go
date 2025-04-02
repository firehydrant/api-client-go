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

// NewPostV1IncidentsIncidentIDTasksTaskIDConvertParams creates a new PostV1IncidentsIncidentIDTasksTaskIDConvertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDTasksTaskIDConvertParams() *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	return &PostV1IncidentsIncidentIDTasksTaskIDConvertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDTasksTaskIDConvertParamsWithTimeout creates a new PostV1IncidentsIncidentIDTasksTaskIDConvertParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDTasksTaskIDConvertParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	return &PostV1IncidentsIncidentIDTasksTaskIDConvertParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDTasksTaskIDConvertParamsWithContext creates a new PostV1IncidentsIncidentIDTasksTaskIDConvertParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDTasksTaskIDConvertParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	return &PostV1IncidentsIncidentIDTasksTaskIDConvertParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDTasksTaskIDConvertParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDTasksTaskIDConvertParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDTasksTaskIDConvertParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	return &PostV1IncidentsIncidentIDTasksTaskIDConvertParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDTasksTaskIDConvertParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id tasks task Id convert operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDTasksTaskIDConvertParams struct {

	// IncidentID.
	IncidentID string

	// PostV1IncidentsIncidentIDTasksTaskIDConvert.
	PostV1IncidentsIncidentIDTasksTaskIDConvert *models.PostV1IncidentsIncidentIDTasksTaskIDConvert

	// TaskID.
	TaskID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id tasks task Id convert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WithDefaults() *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id tasks task Id convert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithPostV1IncidentsIncidentIDTasksTaskIDConvert adds the postV1IncidentsIncidentIDTasksTaskIDConvert to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WithPostV1IncidentsIncidentIDTasksTaskIDConvert(postV1IncidentsIncidentIDTasksTaskIDConvert *models.PostV1IncidentsIncidentIDTasksTaskIDConvert) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	o.SetPostV1IncidentsIncidentIDTasksTaskIDConvert(postV1IncidentsIncidentIDTasksTaskIDConvert)
	return o
}

// SetPostV1IncidentsIncidentIDTasksTaskIDConvert adds the postV1IncidentsIncidentIdTasksTaskIdConvert to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) SetPostV1IncidentsIncidentIDTasksTaskIDConvert(postV1IncidentsIncidentIDTasksTaskIDConvert *models.PostV1IncidentsIncidentIDTasksTaskIDConvert) {
	o.PostV1IncidentsIncidentIDTasksTaskIDConvert = postV1IncidentsIncidentIDTasksTaskIDConvert
}

// WithTaskID adds the taskID to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WithTaskID(taskID string) *PostV1IncidentsIncidentIDTasksTaskIDConvertParams {
	o.SetTaskID(taskID)
	return o
}

// SetTaskID adds the taskId to the post v1 incidents incident Id tasks task Id convert params
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) SetTaskID(taskID string) {
	o.TaskID = taskID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}
	if o.PostV1IncidentsIncidentIDTasksTaskIDConvert != nil {
		if err := r.SetBodyParam(o.PostV1IncidentsIncidentIDTasksTaskIDConvert); err != nil {
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
