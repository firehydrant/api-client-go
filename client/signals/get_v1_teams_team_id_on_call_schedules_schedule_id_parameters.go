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

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParams creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParams() *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithTimeout creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a timeout on a request.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithTimeout(timeout time.Duration) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		timeout: timeout,
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithContext creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a context for a request.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithContext(ctx context.Context) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		Context: ctx,
	}
}

// NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithHTTPClient creates a new GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithHTTPClient(client *http.Client) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		HTTPClient: client,
	}
}

/*
GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams contains all the parameters to send to the API endpoint

	for the get v1 teams team Id on call schedules schedule Id operation.

	Typically these are written to a http.Request.
*/
type GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams struct {

	// ScheduleID.
	ScheduleID string

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 teams team Id on call schedules schedule Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithDefaults() *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 teams team Id on call schedules schedule Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithTimeout(timeout time.Duration) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithContext(ctx context.Context) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithHTTPClient(client *http.Client) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScheduleID adds the scheduleID to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithScheduleID(scheduleID string) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WithTeamID adds the teamID to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithTeamID(teamID string) *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get v1 teams team Id on call schedules schedule Id params
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
