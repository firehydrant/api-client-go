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

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams() *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithTimeout creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithTimeout(timeout time.Duration) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		timeout: timeout,
	}
}

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithContext creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a context for a request.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithContext(ctx context.Context) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		Context: ctx,
	}
}

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithHTTPClient creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParamsWithHTTPClient(client *http.Client) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams contains all the parameters to send to the API endpoint

	for the patch v1 teams team Id on call schedules schedule Id shifts Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams struct {

	// ID.
	ID string

	// PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID.
	PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID *models.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID

	// ScheduleID.
	ScheduleID string

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 teams team Id on call schedules schedule Id shifts Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithDefaults() *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 teams team Id on call schedules schedule Id shifts Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithTimeout(timeout time.Duration) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithContext(ctx context.Context) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithHTTPClient(client *http.Client) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithID(id string) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetID(id string) {
	o.ID = id
}

// WithPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID adds the patchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID(patchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID *models.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID(patchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID)
	return o
}

// SetPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID adds the patchV1TeamsTeamIdOnCallSchedulesScheduleIdShiftsId to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetPatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID(patchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID *models.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID) {
	o.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID = patchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID
}

// WithScheduleID adds the scheduleID to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithScheduleID(scheduleID string) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetScheduleID(scheduleID string) {
	o.ScheduleID = scheduleID
}

// WithTeamID adds the teamID to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WithTeamID(teamID string) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the patch v1 teams team Id on call schedules schedule Id shifts Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID != nil {
		if err := r.SetBodyParam(o.PatchV1TeamsTeamIDOnCallSchedulesScheduleIDShiftsID); err != nil {
			return err
		}
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
