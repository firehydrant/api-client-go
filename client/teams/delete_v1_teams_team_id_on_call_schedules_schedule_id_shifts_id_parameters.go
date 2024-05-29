// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams() *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithTimeout creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithTimeout(timeout time.Duration) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithContext creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a context for a request.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithContext(ctx context.Context) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		Context: ctx,
	}
}

// NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithHTTPClient creates a new DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithHTTPClient(client *http.Client) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams contains all the parameters to send to the API endpoint

	for the delete v1 teams team Id on call schedules schedule Id shifts Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams struct {

	// ID.
	ID string

	// ScheduleID.
	ScheduleID string

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 teams team Id on call schedules schedule Id shifts Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithDefaults() *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 teams team Id on call schedules schedule Id shifts Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithTimeout(timeout time.Duration) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithContext(ctx context.Context) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithHTTPClient(client *http.Client) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithID(id string) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetID(id string) {
	o.ID = id
}

// WithScheduleID adds the scheduleID to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithScheduleID(scheduleID string) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WithTeamID adds the teamID to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithTeamID(teamID string) *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete v1 teams team Id on call schedules schedule Id shifts Id params
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

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
