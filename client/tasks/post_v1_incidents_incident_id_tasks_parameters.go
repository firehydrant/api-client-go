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

// NewPostV1IncidentsIncidentIDTasksParams creates a new PostV1IncidentsIncidentIDTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDTasksParams() *PostV1IncidentsIncidentIDTasksParams {
	return &PostV1IncidentsIncidentIDTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDTasksParamsWithTimeout creates a new PostV1IncidentsIncidentIDTasksParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDTasksParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTasksParams {
	return &PostV1IncidentsIncidentIDTasksParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDTasksParamsWithContext creates a new PostV1IncidentsIncidentIDTasksParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDTasksParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDTasksParams {
	return &PostV1IncidentsIncidentIDTasksParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDTasksParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDTasksParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTasksParams {
	return &PostV1IncidentsIncidentIDTasksParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDTasksParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id tasks operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDTasksParams struct {

	// IncidentID.
	IncidentID string

	// PostV1IncidentsIncidentIDTasks.
	PostV1IncidentsIncidentIDTasks *models.PostV1IncidentsIncidentIDTasks

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTasksParams) WithDefaults() *PostV1IncidentsIncidentIDTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDTasksParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithPostV1IncidentsIncidentIDTasks adds the postV1IncidentsIncidentIDTasks to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) WithPostV1IncidentsIncidentIDTasks(postV1IncidentsIncidentIDTasks *models.PostV1IncidentsIncidentIDTasks) *PostV1IncidentsIncidentIDTasksParams {
	o.SetPostV1IncidentsIncidentIDTasks(postV1IncidentsIncidentIDTasks)
	return o
}

// SetPostV1IncidentsIncidentIDTasks adds the postV1IncidentsIncidentIdTasks to the post v1 incidents incident Id tasks params
func (o *PostV1IncidentsIncidentIDTasksParams) SetPostV1IncidentsIncidentIDTasks(postV1IncidentsIncidentIDTasks *models.PostV1IncidentsIncidentIDTasks) {
	o.PostV1IncidentsIncidentIDTasks = postV1IncidentsIncidentIDTasks
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}
	if o.PostV1IncidentsIncidentIDTasks != nil {
		if err := r.SetBodyParam(o.PostV1IncidentsIncidentIDTasks); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
