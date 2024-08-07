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

// NewGetV1IncidentsIncidentIDTranscriptParams creates a new GetV1IncidentsIncidentIDTranscriptParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsIncidentIDTranscriptParams() *GetV1IncidentsIncidentIDTranscriptParams {
	return &GetV1IncidentsIncidentIDTranscriptParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsIncidentIDTranscriptParamsWithTimeout creates a new GetV1IncidentsIncidentIDTranscriptParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsIncidentIDTranscriptParamsWithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTranscriptParams {
	return &GetV1IncidentsIncidentIDTranscriptParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsIncidentIDTranscriptParamsWithContext creates a new GetV1IncidentsIncidentIDTranscriptParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsIncidentIDTranscriptParamsWithContext(ctx context.Context) *GetV1IncidentsIncidentIDTranscriptParams {
	return &GetV1IncidentsIncidentIDTranscriptParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsIncidentIDTranscriptParamsWithHTTPClient creates a new GetV1IncidentsIncidentIDTranscriptParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsIncidentIDTranscriptParamsWithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTranscriptParams {
	return &GetV1IncidentsIncidentIDTranscriptParams{
		HTTPClient: client,
	}
}

/*
GetV1IncidentsIncidentIDTranscriptParams contains all the parameters to send to the API endpoint

	for the get v1 incidents incident Id transcript operation.

	Typically these are written to a http.Request.
*/
type GetV1IncidentsIncidentIDTranscriptParams struct {

	/* After.

	   The ID of the transcript entry to start after.
	*/
	After *string

	/* Before.

	   The ID of the transcript entry to start before.
	*/
	Before *string

	// IncidentID.
	IncidentID string

	/* Sort.

	   The order to sort the transcript entries.

	   Default: "asc"
	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents incident Id transcript params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithDefaults() *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents incident Id transcript params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetDefaults() {
	var (
		sortDefault = string("asc")
	)

	val := GetV1IncidentsIncidentIDTranscriptParams{
		Sort: &sortDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithTimeout(timeout time.Duration) *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithContext(ctx context.Context) *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithHTTPClient(client *http.Client) *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithAfter(after *string) *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithBefore(before *string) *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetBefore(before *string) {
	o.Before = before
}

// WithIncidentID adds the incidentID to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithIncidentID(incidentID string) *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithSort adds the sort to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) WithSort(sort *string) *GetV1IncidentsIncidentIDTranscriptParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get v1 incidents incident Id transcript params
func (o *GetV1IncidentsIncidentIDTranscriptParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsIncidentIDTranscriptParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore string

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
