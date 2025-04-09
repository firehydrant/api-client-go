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
	"github.com/go-openapi/swag"
)

// NewGetV1TeamsTeamIDSignalRulesParams creates a new GetV1TeamsTeamIDSignalRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TeamsTeamIDSignalRulesParams() *GetV1TeamsTeamIDSignalRulesParams {
	return &GetV1TeamsTeamIDSignalRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TeamsTeamIDSignalRulesParamsWithTimeout creates a new GetV1TeamsTeamIDSignalRulesParams object
// with the ability to set a timeout on a request.
func NewGetV1TeamsTeamIDSignalRulesParamsWithTimeout(timeout time.Duration) *GetV1TeamsTeamIDSignalRulesParams {
	return &GetV1TeamsTeamIDSignalRulesParams{
		timeout: timeout,
	}
}

// NewGetV1TeamsTeamIDSignalRulesParamsWithContext creates a new GetV1TeamsTeamIDSignalRulesParams object
// with the ability to set a context for a request.
func NewGetV1TeamsTeamIDSignalRulesParamsWithContext(ctx context.Context) *GetV1TeamsTeamIDSignalRulesParams {
	return &GetV1TeamsTeamIDSignalRulesParams{
		Context: ctx,
	}
}

// NewGetV1TeamsTeamIDSignalRulesParamsWithHTTPClient creates a new GetV1TeamsTeamIDSignalRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TeamsTeamIDSignalRulesParamsWithHTTPClient(client *http.Client) *GetV1TeamsTeamIDSignalRulesParams {
	return &GetV1TeamsTeamIDSignalRulesParams{
		HTTPClient: client,
	}
}

/*
GetV1TeamsTeamIDSignalRulesParams contains all the parameters to send to the API endpoint

	for the get v1 teams team Id signal rules operation.

	Typically these are written to a http.Request.
*/
type GetV1TeamsTeamIDSignalRulesParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   A query string for searching through the list of alerting rules.
	*/
	Query *string

	// TeamID.
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 teams team Id signal rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDSignalRulesParams) WithDefaults() *GetV1TeamsTeamIDSignalRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 teams team Id signal rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TeamsTeamIDSignalRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) WithTimeout(timeout time.Duration) *GetV1TeamsTeamIDSignalRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) WithContext(ctx context.Context) *GetV1TeamsTeamIDSignalRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) WithHTTPClient(client *http.Client) *GetV1TeamsTeamIDSignalRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) WithPage(page *int32) *GetV1TeamsTeamIDSignalRulesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) WithPerPage(perPage *int32) *GetV1TeamsTeamIDSignalRulesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) WithQuery(query *string) *GetV1TeamsTeamIDSignalRulesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) SetQuery(query *string) {
	o.Query = query
}

// WithTeamID adds the teamID to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) WithTeamID(teamID string) *GetV1TeamsTeamIDSignalRulesParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get v1 teams team Id signal rules params
func (o *GetV1TeamsTeamIDSignalRulesParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TeamsTeamIDSignalRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
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
