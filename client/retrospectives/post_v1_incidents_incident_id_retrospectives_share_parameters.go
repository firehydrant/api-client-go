// Code generated by go-swagger; DO NOT EDIT.

package retrospectives

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

// NewPostV1IncidentsIncidentIDRetrospectivesShareParams creates a new PostV1IncidentsIncidentIDRetrospectivesShareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDRetrospectivesShareParams() *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	return &PostV1IncidentsIncidentIDRetrospectivesShareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDRetrospectivesShareParamsWithTimeout creates a new PostV1IncidentsIncidentIDRetrospectivesShareParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDRetrospectivesShareParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	return &PostV1IncidentsIncidentIDRetrospectivesShareParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDRetrospectivesShareParamsWithContext creates a new PostV1IncidentsIncidentIDRetrospectivesShareParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDRetrospectivesShareParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	return &PostV1IncidentsIncidentIDRetrospectivesShareParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDRetrospectivesShareParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDRetrospectivesShareParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDRetrospectivesShareParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	return &PostV1IncidentsIncidentIDRetrospectivesShareParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDRetrospectivesShareParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id retrospectives share operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDRetrospectivesShareParams struct {

	// IncidentID.
	IncidentID string

	/* RetrospectiveIds.

	   An array of retrospective IDs to share
	*/
	RetrospectiveIds []string

	/* TeamIds.

	   An array of team IDs with whom to share the report
	*/
	TeamIds []string

	/* UserIds.

	   An array of user IDs with whom to share the report
	*/
	UserIds []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id retrospectives share params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithDefaults() *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id retrospectives share params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithRetrospectiveIds adds the retrospectiveIds to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithRetrospectiveIds(retrospectiveIds []string) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetRetrospectiveIds(retrospectiveIds)
	return o
}

// SetRetrospectiveIds adds the retrospectiveIds to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetRetrospectiveIds(retrospectiveIds []string) {
	o.RetrospectiveIds = retrospectiveIds
}

// WithTeamIds adds the teamIds to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithTeamIds(teamIds []string) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetTeamIds(teamIds)
	return o
}

// SetTeamIds adds the teamIds to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetTeamIds(teamIds []string) {
	o.TeamIds = teamIds
}

// WithUserIds adds the userIds to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WithUserIds(userIds []string) *PostV1IncidentsIncidentIDRetrospectivesShareParams {
	o.SetUserIds(userIds)
	return o
}

// SetUserIds adds the userIds to the post v1 incidents incident Id retrospectives share params
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) SetUserIds(userIds []string) {
	o.UserIds = userIds
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if o.RetrospectiveIds != nil {

		// binding items for retrospective_ids
		joinedRetrospectiveIds := o.bindParamRetrospectiveIds(reg)

		// form array param retrospective_ids
		if err := r.SetFormParam("retrospective_ids", joinedRetrospectiveIds...); err != nil {
			return err
		}
	}

	if o.TeamIds != nil {

		// binding items for team_ids
		joinedTeamIds := o.bindParamTeamIds(reg)

		// form array param team_ids
		if err := r.SetFormParam("team_ids", joinedTeamIds...); err != nil {
			return err
		}
	}

	if o.UserIds != nil {

		// binding items for user_ids
		joinedUserIds := o.bindParamUserIds(reg)

		// form array param user_ids
		if err := r.SetFormParam("user_ids", joinedUserIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPostV1IncidentsIncidentIDRetrospectivesShare binds the parameter retrospective_ids
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) bindParamRetrospectiveIds(formats strfmt.Registry) []string {
	retrospectiveIdsIR := o.RetrospectiveIds

	var retrospectiveIdsIC []string
	for _, retrospectiveIdsIIR := range retrospectiveIdsIR { // explode []string

		retrospectiveIdsIIV := retrospectiveIdsIIR // string as string
		retrospectiveIdsIC = append(retrospectiveIdsIC, retrospectiveIdsIIV)
	}

	// items.CollectionFormat: ""
	retrospectiveIdsIS := swag.JoinByFormat(retrospectiveIdsIC, "")

	return retrospectiveIdsIS
}

// bindParamPostV1IncidentsIncidentIDRetrospectivesShare binds the parameter team_ids
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) bindParamTeamIds(formats strfmt.Registry) []string {
	teamIdsIR := o.TeamIds

	var teamIdsIC []string
	for _, teamIdsIIR := range teamIdsIR { // explode []string

		teamIdsIIV := teamIdsIIR // string as string
		teamIdsIC = append(teamIdsIC, teamIdsIIV)
	}

	// items.CollectionFormat: ""
	teamIdsIS := swag.JoinByFormat(teamIdsIC, "")

	return teamIdsIS
}

// bindParamPostV1IncidentsIncidentIDRetrospectivesShare binds the parameter user_ids
func (o *PostV1IncidentsIncidentIDRetrospectivesShareParams) bindParamUserIds(formats strfmt.Registry) []string {
	userIdsIR := o.UserIds

	var userIdsIC []string
	for _, userIdsIIR := range userIdsIR { // explode []string

		userIdsIIV := userIdsIIR // string as string
		userIdsIC = append(userIdsIC, userIdsIIV)
	}

	// items.CollectionFormat: ""
	userIdsIS := swag.JoinByFormat(userIdsIC, "")

	return userIdsIS
}
