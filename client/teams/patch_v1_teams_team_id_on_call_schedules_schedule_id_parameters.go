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

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams() *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithTimeout creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithTimeout(timeout time.Duration) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		timeout: timeout,
	}
}

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithContext creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a context for a request.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithContext(ctx context.Context) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		Context: ctx,
	}
}

// NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithHTTPClient creates a new PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1TeamsTeamIDOnCallSchedulesScheduleIDParamsWithHTTPClient(client *http.Client) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	return &PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams contains all the parameters to send to the API endpoint

	for the patch v1 teams team Id on call schedules schedule Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams struct {

	// PatchV1TeamsTeamIDOnCallSchedulesScheduleID.
	PatchV1TeamsTeamIDOnCallSchedulesScheduleID *models.PatchV1TeamsTeamIDOnCallSchedulesScheduleID

	// ScheduleID.
	//
	// Format: int32
	ScheduleID int32

	/* TeamID.

	   The ID of the team for which you want to manage Signals on-call schedules.
	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 teams team Id on call schedules schedule Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithDefaults() *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 teams team Id on call schedules schedule Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithTimeout(timeout time.Duration) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithContext(ctx context.Context) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithHTTPClient(client *http.Client) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchV1TeamsTeamIDOnCallSchedulesScheduleID adds the patchV1TeamsTeamIDOnCallSchedulesScheduleID to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithPatchV1TeamsTeamIDOnCallSchedulesScheduleID(patchV1TeamsTeamIDOnCallSchedulesScheduleID *models.PatchV1TeamsTeamIDOnCallSchedulesScheduleID) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetPatchV1TeamsTeamIDOnCallSchedulesScheduleID(patchV1TeamsTeamIDOnCallSchedulesScheduleID)
	return o
}

// SetPatchV1TeamsTeamIDOnCallSchedulesScheduleID adds the patchV1TeamsTeamIdOnCallSchedulesScheduleId to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetPatchV1TeamsTeamIDOnCallSchedulesScheduleID(patchV1TeamsTeamIDOnCallSchedulesScheduleID *models.PatchV1TeamsTeamIDOnCallSchedulesScheduleID) {
	o.PatchV1TeamsTeamIDOnCallSchedulesScheduleID = patchV1TeamsTeamIDOnCallSchedulesScheduleID
}

// WithScheduleID adds the scheduleID to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithScheduleID(scheduleID int32) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetScheduleID(scheduleID)
	return o
}

// SetScheduleID adds the scheduleId to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetScheduleID(scheduleID int32) {
	o.ScheduleID = scheduleID
}

// WithTeamID adds the teamID to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WithTeamID(teamID string) *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the patch v1 teams team Id on call schedules schedule Id params
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1TeamsTeamIDOnCallSchedulesScheduleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PatchV1TeamsTeamIDOnCallSchedulesScheduleID != nil {
		if err := r.SetBodyParam(o.PatchV1TeamsTeamIDOnCallSchedulesScheduleID); err != nil {
			return err
		}
	}

	// path param schedule_id
	if err := r.SetPathParam("schedule_id", swag.FormatInt32(o.ScheduleID)); err != nil {
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
