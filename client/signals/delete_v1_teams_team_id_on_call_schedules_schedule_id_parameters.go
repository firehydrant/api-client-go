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

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams() *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithTimeout creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithTimeout(timeout time.Duration) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithContext creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a context for a request.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithContext(ctx context.Context) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		Context: ctx,
	}
}

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithHTTPClient creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithHTTPClient(client *http.Client) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams contains all the parameters to send to the API endpoint

	for the delete v1 teams team Id on call schedules schedule Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams struct {

	// ScheduleID.
	ScheduleID string

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 teams team Id on call schedules schedule Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithDefaults() *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 teams team Id on call schedules schedule Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithTimeout(timeout time.Duration) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithContext(ctx context.Context) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithHTTPClient(client *http.Client) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScheduleID adds the scheduleID to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithScheduleID(scheduleID string) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WithTeamID adds the teamID to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithTeamID(teamID string) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete v1 teams team Id on call schedules schedule Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param schedule_id
	if err := r.SetPathParam("schedule_id", o.ScheduleID); err != nil {
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
