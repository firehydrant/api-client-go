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

// NewDeleteV1TeamsTeamIDSignalRulesIDParams creates a new DeleteV1TeamsTeamIDSignalRulesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1TeamsTeamIDSignalRulesIDParams() *DeleteV1TeamsTeamIDSignalRulesIDParams {
	return &DeleteV1TeamsTeamIDSignalRulesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1TeamsTeamIDSignalRulesIDParamsWithTimeout creates a new DeleteV1TeamsTeamIDSignalRulesIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1TeamsTeamIDSignalRulesIDParamsWithTimeout(timeout time.Duration) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	return &DeleteV1TeamsTeamIDSignalRulesIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1TeamsTeamIDSignalRulesIDParamsWithContext creates a new DeleteV1TeamsTeamIDSignalRulesIDParams object
// with the ability to set a context for a request.
func NewDeleteV1TeamsTeamIDSignalRulesIDParamsWithContext(ctx context.Context) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	return &DeleteV1TeamsTeamIDSignalRulesIDParams{
		Context: ctx,
	}
}

// NewDeleteV1TeamsTeamIDSignalRulesIDParamsWithHTTPClient creates a new DeleteV1TeamsTeamIDSignalRulesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1TeamsTeamIDSignalRulesIDParamsWithHTTPClient(client *http.Client) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	return &DeleteV1TeamsTeamIDSignalRulesIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1TeamsTeamIDSignalRulesIDParams contains all the parameters to send to the API endpoint

	for the delete v1 teams team Id signal rules Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1TeamsTeamIDSignalRulesIDParams struct {

	// ID.
	ID string

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 teams team Id signal rules Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) WithDefaults() *DeleteV1TeamsTeamIDSignalRulesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 teams team Id signal rules Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) WithTimeout(timeout time.Duration) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) WithContext(ctx context.Context) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) WithHTTPClient(client *http.Client) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) WithID(id string) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) SetID(id string) {
	o.ID = id
}

// WithTeamID adds the teamID to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) WithTeamID(teamID string) *DeleteV1TeamsTeamIDSignalRulesIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete v1 teams team Id signal rules Id params
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1TeamsTeamIDSignalRulesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
