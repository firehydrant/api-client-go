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
	"github.com/go-openapi/swag"
)

// NewGetV1IncidentsIncidentIDTasksParams creates a new GetV1IncidentsIncidentIDTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDTasksParams() *GetV1IncidentsIncidentIDTasksParams {
	return &GetV1IncidentsIncidentIDTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDTasksParamsWithTimeout creates a new GetV1IncidentsIncidentIDTasksParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDTasksParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTasksParams {
	return &GetV1IncidentsIncidentIDTasksParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDTasksParamsWithContext creates a new GetV1IncidentsIncidentIDTasksParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDTasksParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDTasksParams {
	return &GetV1IncidentsIncidentIDTasksParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDTasksParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDTasksParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTasksParams {
	return &GetV1IncidentsIncidentIDTasksParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentsIncidentIDTasksParams contains all the parameters to send to the API endpoint

	for the get v1 incidents incident Id tasks operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDTasksParams struct {

	// IncidentID.
	IncidentID string

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDTasksParams) WithDefaults() *GetV1IncidentsIncidentIDTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDTasksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDTasksParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithPage adds the page to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithPage(page *int32) *GetV1IncidentsIncidentIDTasksParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) WithPerPage(perPage *int32) *GetV1IncidentsIncidentIDTasksParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 incidents incident Id tasks params
func (o *GetV1IncidentsIncidentIDTasksParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
