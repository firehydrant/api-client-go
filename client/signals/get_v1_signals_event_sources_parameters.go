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

// NewGetV1SignalsEventSourcesParams creates a new GetV1SignalsEventSourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SignalsEventSourcesParams() *GetV1SignalsEventSourcesParams {
	return &GetV1SignalsEventSourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SignalsEventSourcesParamsWithTimeout creates a new GetV1SignalsEventSourcesParams object
// with the ability to set a timeout on a request.
func NewGetV1SignalsEventSourcesParamsWithTimeout(timeout time.Duration) *GetV1SignalsEventSourcesParams {
	return &GetV1SignalsEventSourcesParams{
		timeout: timeout,
	}
}

// NewGetV1SignalsEventSourcesParamsWithContext creates a new GetV1SignalsEventSourcesParams object
// with the ability to set a context for a request.
func NewGetV1SignalsEventSourcesParamsWithContext(ctx context.Context) *GetV1SignalsEventSourcesParams {
	return &GetV1SignalsEventSourcesParams{
		Context: ctx,
	}
}

// NewGetV1SignalsEventSourcesParamsWithHTTPClient creates a new GetV1SignalsEventSourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SignalsEventSourcesParamsWithHTTPClient(client *http.Client) *GetV1SignalsEventSourcesParams {
	return &GetV1SignalsEventSourcesParams{
		HTTPClient: client,
	}
}

/*
GetV1SignalsEventSourcesParams contains all the parameters to send to the API endpoint

	for the get v1 signals event sources operation.

	Typically these are written to a http.Request.
*/
type GetV1SignalsEventSourcesParams struct {

	/* EscalationPolicyID.

	   Escalation policy ID to send signals to directly. `team_id` is required if this is provided.
	*/
	EscalationPolicyID *string

	/* OnCallScheduleID.

	   On-call schedule ID to send signals to directly. `team_id` is required if this is provided.
	*/
	OnCallScheduleID *string

	/* TeamID.

	   Team ID to send signals to directly
	*/
	TeamID *string

	/* UserID.

	   User ID to send signals to directly
	*/
	UserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 signals event sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsEventSourcesParams) WithDefaults() *GetV1SignalsEventSourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 signals event sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsEventSourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) WithTimeout(timeout time.Duration) *GetV1SignalsEventSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) WithContext(ctx context.Context) *GetV1SignalsEventSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) WithHTTPClient(client *http.Client) *GetV1SignalsEventSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEscalationPolicyID adds the escalationPolicyID to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) WithEscalationPolicyID(escalationPolicyID *string) *GetV1SignalsEventSourcesParams {
	o.SetEscalationPolicyID(escalationPolicyID)
	return o
}

// SetEscalationPolicyID adds the escalationPolicyId to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) SetEscalationPolicyID(escalationPolicyID *string) {
	o.EscalationPolicyID = escalationPolicyID
}

// WithOnCallScheduleID adds the onCallScheduleID to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) WithOnCallScheduleID(onCallScheduleID *string) *GetV1SignalsEventSourcesParams {
	o.SetOnCallScheduleID(onCallScheduleID)
	return o
}

// SetOnCallScheduleID adds the onCallScheduleId to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) SetOnCallScheduleID(onCallScheduleID *string) {
	o.OnCallScheduleID = onCallScheduleID
}

// WithTeamID adds the teamID to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) WithTeamID(teamID *string) *GetV1SignalsEventSourcesParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) SetTeamID(teamID *string) {
	o.TeamID = teamID
}

// WithUserID adds the userID to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) WithUserID(userID *string) *GetV1SignalsEventSourcesParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get v1 signals event sources params
func (o *GetV1SignalsEventSourcesParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SignalsEventSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EscalationPolicyID != nil {

		// query param escalation_policy_id
		var qrEscalationPolicyID string

		if o.EscalationPolicyID != nil {
			qrEscalationPolicyID = *o.EscalationPolicyID
		}
		qEscalationPolicyID := qrEscalationPolicyID
		if qEscalationPolicyID != "" {

			if err := r.SetQueryParam("escalation_policy_id", qEscalationPolicyID); err != nil {
				return err
			}
		}
	}

	if o.OnCallScheduleID != nil {

		// query param on_call_schedule_id
		var qrOnCallScheduleID string

		if o.OnCallScheduleID != nil {
			qrOnCallScheduleID = *o.OnCallScheduleID
		}
		qOnCallScheduleID := qrOnCallScheduleID
		if qOnCallScheduleID != "" {

			if err := r.SetQueryParam("on_call_schedule_id", qOnCallScheduleID); err != nil {
				return err
			}
		}
	}

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

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
