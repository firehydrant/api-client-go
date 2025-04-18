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

// NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams creates a new GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams() *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	return &GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParamsWithTimeout creates a new GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams object
// with the ability to set a timeout on a request.
func NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParamsWithTimeout(timeout time.Duration) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	return &GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams{
		timeout: timeout,
	}
}

// NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParamsWithContext creates a new GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams object
// with the ability to set a context for a request.
func NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParamsWithContext(ctx context.Context) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	return &GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams{
		Context: ctx,
	}
}

// NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParamsWithHTTPClient creates a new GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParamsWithHTTPClient(client *http.Client) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	return &GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams{
		HTTPClient: client,
	}
}

/*
GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams contains all the parameters to send to the API endpoint

	for the get v1 ai summarize incident incident Id generated summary Id voted operation.

	Typically these are written to a http.Request.
*/
type GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams struct {

	// GeneratedSummaryID.
	GeneratedSummaryID string

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 ai summarize incident incident Id generated summary Id voted params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) WithDefaults() *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 ai summarize incident incident Id generated summary Id voted params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) WithTimeout(timeout time.Duration) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) WithContext(ctx context.Context) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) WithHTTPClient(client *http.Client) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneratedSummaryID adds the generatedSummaryID to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) WithGeneratedSummaryID(generatedSummaryID string) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	o.SetGeneratedSummaryID(generatedSummaryID)
	return o
}

// SetGeneratedSummaryID adds the generatedSummaryId to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) SetGeneratedSummaryID(generatedSummaryID string) {
	o.GeneratedSummaryID = generatedSummaryID
}

// WithIncidentID adds the incidentID to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) WithIncidentID(incidentID string) *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the get v1 ai summarize incident incident Id generated summary Id voted params
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVotedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param generated_summary_id
	if err := r.SetPathParam("generated_summary_id", o.GeneratedSummaryID); err != nil {
		return err
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
