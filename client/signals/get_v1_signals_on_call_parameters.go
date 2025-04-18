// Code generated by go-swagger; DO NOT EDIT.

package signals

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

// NewGetV1SignalsOnCallParams creates a new GetV1SignalsOnCallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SignalsOnCallParams() *GetV1SignalsOnCallParams {
	return &GetV1SignalsOnCallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SignalsOnCallParamsWithTimeout creates a new GetV1SignalsOnCallParams object
// with the ability to set a timeout on a request.
func NewGetV1SignalsOnCallParamsWithTimeout(timeout time.Duration) *GetV1SignalsOnCallParams {
	return &GetV1SignalsOnCallParams{
		timeout: timeout,
	}
}

// NewGetV1SignalsOnCallParamsWithContext creates a new GetV1SignalsOnCallParams object
// with the ability to set a context for a request.
func NewGetV1SignalsOnCallParamsWithContext(ctx context.Context) *GetV1SignalsOnCallParams {
	return &GetV1SignalsOnCallParams{
		Context: ctx,
	}
}

// NewGetV1SignalsOnCallParamsWithHTTPClient creates a new GetV1SignalsOnCallParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SignalsOnCallParamsWithHTTPClient(client *http.Client) *GetV1SignalsOnCallParams {
	return &GetV1SignalsOnCallParams{
		HTTPClient: client,
	}
}

/*
GetV1SignalsOnCallParams contains all the parameters to send to the API endpoint

	for the get v1 signals on call operation.

	Typically these are written to a http.Request.
*/
type GetV1SignalsOnCallParams struct {

	/* TeamID.

	   An optional comma separated list of team IDs to filter currently on-call users by
	*/
	TeamID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 signals on call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsOnCallParams) WithDefaults() *GetV1SignalsOnCallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 signals on call params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsOnCallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) WithTimeout(timeout time.Duration) *GetV1SignalsOnCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) WithContext(ctx context.Context) *GetV1SignalsOnCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) WithHTTPClient(client *http.Client) *GetV1SignalsOnCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) WithTeamID(teamID *string) *GetV1SignalsOnCallParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get v1 signals on call params
func (o *GetV1SignalsOnCallParams) SetTeamID(teamID *string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SignalsOnCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.TeamID != nil {

		// query param team_id
		var qrTeamID string

		if o.TeamID != nil {
			qrTeamID = *o.TeamID
		}
		qTeamID := qrTeamID
		if qTeamID != "" {

			if err := r.SetQueryParam("team_id", qTeamID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
