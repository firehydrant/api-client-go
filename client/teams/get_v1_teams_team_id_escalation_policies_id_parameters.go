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

// NewGetV1TeamsTeamIDEscalationPoliciesIDParams creates a new GetV1TeamsTeamIDEscalationPoliciesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TeamsTeamIDEscalationPoliciesIDParams() *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	return &GetV1TeamsTeamIDEscalationPoliciesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TeamsTeamIDEscalationPoliciesIDParamsWithTimeout creates a new GetV1TeamsTeamIDEscalationPoliciesIDParams object
// with the ability to set a timeout on a request.
func NewGetV1TeamsTeamIDEscalationPoliciesIDParamsWithTimeout(timeout time.Duration) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	return &GetV1TeamsTeamIDEscalationPoliciesIDParams{
		timeout: timeout,
	}
}

// NewGetV1TeamsTeamIDEscalationPoliciesIDParamsWithContext creates a new GetV1TeamsTeamIDEscalationPoliciesIDParams object
// with the ability to set a context for a request.
func NewGetV1TeamsTeamIDEscalationPoliciesIDParamsWithContext(ctx context.Context) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	return &GetV1TeamsTeamIDEscalationPoliciesIDParams{
		Context: ctx,
	}
}

// NewGetV1TeamsTeamIDEscalationPoliciesIDParamsWithHTTPClient creates a new GetV1TeamsTeamIDEscalationPoliciesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TeamsTeamIDEscalationPoliciesIDParamsWithHTTPClient(client *http.Client) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	return &GetV1TeamsTeamIDEscalationPoliciesIDParams{
		HTTPClient: client,
	}
}

/*
GetV1TeamsTeamIDEscalationPoliciesIDParams contains all the parameters to send to the API endpoint

	for the get v1 teams team Id escalation policies Id operation.

	Typically these are written to a http.Request.
*/
type GetV1TeamsTeamIDEscalationPoliciesIDParams struct {

	/* ID.

	   The ID of the escalation policy you want to manage.
	*/
	ID string

	/* TeamID.

	   The ID of the team for which you want to manage Signals escalation policies.
	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 teams team Id escalation policies Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) WithDefaults() *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 teams team Id escalation policies Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) WithTimeout(timeout time.Duration) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) WithContext(ctx context.Context) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) WithHTTPClient(client *http.Client) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) WithID(id string) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) SetID(id string) {
	o.ID = id
}

// WithTeamID adds the teamID to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) WithTeamID(teamID string) *GetV1TeamsTeamIDEscalationPoliciesIDParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get v1 teams team Id escalation policies Id params
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TeamsTeamIDEscalationPoliciesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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