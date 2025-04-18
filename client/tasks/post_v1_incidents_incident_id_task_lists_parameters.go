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

// NewPostV1IncidentsIncidentIDTaskListsParams creates a new PostV1IncidentsIncidentIDTaskListsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDTaskListsParams() *PostV1IncidentsIncidentIDTaskListsParams {
	return &PostV1IncidentsIncidentIDTaskListsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDTaskListsParamsWithTimeout creates a new PostV1IncidentsIncidentIDTaskListsParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDTaskListsParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTaskListsParams {
	return &PostV1IncidentsIncidentIDTaskListsParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDTaskListsParamsWithContext creates a new PostV1IncidentsIncidentIDTaskListsParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDTaskListsParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDTaskListsParams {
	return &PostV1IncidentsIncidentIDTaskListsParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDTaskListsParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDTaskListsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDTaskListsParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTaskListsParams {
	return &PostV1IncidentsIncidentIDTaskListsParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDTaskListsParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id task lists operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDTaskListsParams struct {

	// IncidentID.
	IncidentID string

	// PostV1IncidentsIncidentIDTaskLists.
	PostV1IncidentsIncidentIDTaskLists *models.PostV1IncidentsIncidentIDTaskLists

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id task lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTaskListsParams) WithDefaults() *PostV1IncidentsIncidentIDTaskListsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id task lists params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDTaskListsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDTaskListsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDTaskListsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDTaskListsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDTaskListsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithPostV1IncidentsIncidentIDTaskLists adds the postV1IncidentsIncidentIDTaskLists to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) WithPostV1IncidentsIncidentIDTaskLists(postV1IncidentsIncidentIDTaskLists *models.PostV1IncidentsIncidentIDTaskLists) *PostV1IncidentsIncidentIDTaskListsParams {
	o.SetPostV1IncidentsIncidentIDTaskLists(postV1IncidentsIncidentIDTaskLists)
	return o
}

// SetPostV1IncidentsIncidentIDTaskLists adds the postV1IncidentsIncidentIdTaskLists to the post v1 incidents incident Id task lists params
func (o *PostV1IncidentsIncidentIDTaskListsParams) SetPostV1IncidentsIncidentIDTaskLists(postV1IncidentsIncidentIDTaskLists *models.PostV1IncidentsIncidentIDTaskLists) {
	o.PostV1IncidentsIncidentIDTaskLists = postV1IncidentsIncidentIDTaskLists
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDTaskListsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}
	if o.PostV1IncidentsIncidentIDTaskLists != nil {
		if err := r.SetBodyParam(o.PostV1IncidentsIncidentIDTaskLists); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
