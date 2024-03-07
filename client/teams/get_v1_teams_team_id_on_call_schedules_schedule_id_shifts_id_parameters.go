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
	"github.com/go-openapi/swag"
)

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams() *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithTimeout creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a timeout on a request.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithTimeout(timeout time.Duration) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		timeout: timeout,
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithContext creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a context for a request.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithContext(ctx context.Context) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		Context: ctx,
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithHTTPClient creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithHTTPClient(client *http.Client) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		HTTPClient: client,
	}
}

/*
GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams contains all the parameters to send to the API endpoint

	for the get v1 teams team Id on call schedules schedule Id shifts Id operation.

	Typically these are written to a http.Request.
*/
type GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams struct {

	/* ID.

	   The ID of the on-call shift you want to manage.
	*/
	ID string

	// ScheduleID.
	//
	// Format: int32
	ScheduleID int32

	// TeamID.
	//
	// Format: int32
	TeamID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 teams team Id on call schedules schedule Id shifts Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithDefaults() *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 teams team Id on call schedules schedule Id shifts Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithTimeout(timeout time.Duration) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithContext(ctx context.Context) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithHTTPClient(client *http.Client) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithID(id string) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetID(id string) {
	o.ID = id
}

// WithScheduleID adds the scheduleID to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithScheduleID(scheduleID int32) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetScheduleID(scheduleID int32) {
	o.ScheduleID = scheduleID
}

// WithTeamID adds the teamID to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithTeamID(teamID int32) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get v1 teams team Id on call schedules schedule Id shifts Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetTeamID(teamID int32) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param schedule_id
	if err := r.SetPathParam("schedule_id", swag.FormatInt32(o.ScheduleID)); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt32(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
