// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewGetV1IncidentsParams creates a new GetV1IncidentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IncidentsParams() *GetV1IncidentsParams {
	return &GetV1IncidentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IncidentsParamsWithTimeout creates a new GetV1IncidentsParams object
// with the ability to set a timeout on a request.
func NewGetV1IncidentsParamsWithTimeout(timeout time.Duration) *GetV1IncidentsParams {
	return &GetV1IncidentsParams{
		timeout: timeout,
	}
}

// NewGetV1IncidentsParamsWithContext creates a new GetV1IncidentsParams object
// with the ability to set a context for a request.
func NewGetV1IncidentsParamsWithContext(ctx context.Context) *GetV1IncidentsParams {
	return &GetV1IncidentsParams{
		Context: ctx,
	}
}

// NewGetV1IncidentsParamsWithHTTPClient creates a new GetV1IncidentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IncidentsParamsWithHTTPClient(client *http.Client) *GetV1IncidentsParams {
	return &GetV1IncidentsParams{
		HTTPClient: client,
	}
}

/* GetV1IncidentsParams contains all the parameters to send to the API endpoint
   for the get v1 incidents operation.

   Typically these are written to a http.Request.
*/
type GetV1IncidentsParams struct {

	// Conditions.
	Conditions *string

	/* CurrentMilestones.

	   A comma separated list of current milestones
	*/
	CurrentMilestones *string

	/* EndDate.

	   The end date to return incidents from

	   Format: date
	*/
	EndDate *strfmt.Date

	/* Environments.

	   A comma separated list of environment IDs
	*/
	Environments *string

	/* Name.

	   A query to search incidents by their name
	*/
	Name *string

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   A text query for an incident that searches on name, summary, and desciption
	*/
	Query *string

	/* SavedSearchID.

	   The id of a previously saved search.
	*/
	SavedSearchID *string

	/* Services.

	   A comma separated list of service IDs
	*/
	Services *string

	/* Severities.

	   A text value of severity
	*/
	Severities *string

	/* SeverityNotSet.

	   Flag for including incidents where severity has not been set
	*/
	SeverityNotSet *bool

	/* StartDate.

	   The start date to return incidents from

	   Format: date
	*/
	StartDate *strfmt.Date

	/* Status.

	   Incident status
	*/
	Status *string

	/* TagMatchStrategy.

	   A matching strategy for the tags provided
	*/
	TagMatchStrategy *string

	/* Tags.

	   A comma separated list of tags
	*/
	Tags *string

	/* Teams.

	   A comma separated list of team IDs
	*/
	Teams *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 incidents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsParams) WithDefaults() *GetV1IncidentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 incidents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IncidentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 incidents params
func (o *GetV1IncidentsParams) WithTimeout(timeout time.Duration) *GetV1IncidentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 incidents params
func (o *GetV1IncidentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 incidents params
func (o *GetV1IncidentsParams) WithContext(ctx context.Context) *GetV1IncidentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 incidents params
func (o *GetV1IncidentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 incidents params
func (o *GetV1IncidentsParams) WithHTTPClient(client *http.Client) *GetV1IncidentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 incidents params
func (o *GetV1IncidentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConditions adds the conditions to the get v1 incidents params
func (o *GetV1IncidentsParams) WithConditions(conditions *string) *GetV1IncidentsParams {
	o.SetConditions(conditions)
	return o
}

// SetConditions adds the conditions to the get v1 incidents params
func (o *GetV1IncidentsParams) SetConditions(conditions *string) {
	o.Conditions = conditions
}

// WithCurrentMilestones adds the currentMilestones to the get v1 incidents params
func (o *GetV1IncidentsParams) WithCurrentMilestones(currentMilestones *string) *GetV1IncidentsParams {
	o.SetCurrentMilestones(currentMilestones)
	return o
}

// SetCurrentMilestones adds the currentMilestones to the get v1 incidents params
func (o *GetV1IncidentsParams) SetCurrentMilestones(currentMilestones *string) {
	o.CurrentMilestones = currentMilestones
}

// WithEndDate adds the endDate to the get v1 incidents params
func (o *GetV1IncidentsParams) WithEndDate(endDate *strfmt.Date) *GetV1IncidentsParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get v1 incidents params
func (o *GetV1IncidentsParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithEnvironments adds the environments to the get v1 incidents params
func (o *GetV1IncidentsParams) WithEnvironments(environments *string) *GetV1IncidentsParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get v1 incidents params
func (o *GetV1IncidentsParams) SetEnvironments(environments *string) {
	o.Environments = environments
}

// WithName adds the name to the get v1 incidents params
func (o *GetV1IncidentsParams) WithName(name *string) *GetV1IncidentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 incidents params
func (o *GetV1IncidentsParams) SetName(name *string) {
	o.Name = name
}

// WithPage adds the page to the get v1 incidents params
func (o *GetV1IncidentsParams) WithPage(page *int32) *GetV1IncidentsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 incidents params
func (o *GetV1IncidentsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 incidents params
func (o *GetV1IncidentsParams) WithPerPage(perPage *int32) *GetV1IncidentsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 incidents params
func (o *GetV1IncidentsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 incidents params
func (o *GetV1IncidentsParams) WithQuery(query *string) *GetV1IncidentsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 incidents params
func (o *GetV1IncidentsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSavedSearchID adds the savedSearchID to the get v1 incidents params
func (o *GetV1IncidentsParams) WithSavedSearchID(savedSearchID *string) *GetV1IncidentsParams {
	o.SetSavedSearchID(savedSearchID)
	return o
}

// SetSavedSearchID adds the savedSearchId to the get v1 incidents params
func (o *GetV1IncidentsParams) SetSavedSearchID(savedSearchID *string) {
	o.SavedSearchID = savedSearchID
}

// WithServices adds the services to the get v1 incidents params
func (o *GetV1IncidentsParams) WithServices(services *string) *GetV1IncidentsParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the get v1 incidents params
func (o *GetV1IncidentsParams) SetServices(services *string) {
	o.Services = services
}

// WithSeverities adds the severities to the get v1 incidents params
func (o *GetV1IncidentsParams) WithSeverities(severities *string) *GetV1IncidentsParams {
	o.SetSeverities(severities)
	return o
}

// SetSeverities adds the severities to the get v1 incidents params
func (o *GetV1IncidentsParams) SetSeverities(severities *string) {
	o.Severities = severities
}

// WithSeverityNotSet adds the severityNotSet to the get v1 incidents params
func (o *GetV1IncidentsParams) WithSeverityNotSet(severityNotSet *bool) *GetV1IncidentsParams {
	o.SetSeverityNotSet(severityNotSet)
	return o
}

// SetSeverityNotSet adds the severityNotSet to the get v1 incidents params
func (o *GetV1IncidentsParams) SetSeverityNotSet(severityNotSet *bool) {
	o.SeverityNotSet = severityNotSet
}

// WithStartDate adds the startDate to the get v1 incidents params
func (o *GetV1IncidentsParams) WithStartDate(startDate *strfmt.Date) *GetV1IncidentsParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get v1 incidents params
func (o *GetV1IncidentsParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithStatus adds the status to the get v1 incidents params
func (o *GetV1IncidentsParams) WithStatus(status *string) *GetV1IncidentsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get v1 incidents params
func (o *GetV1IncidentsParams) SetStatus(status *string) {
	o.Status = status
}

// WithTagMatchStrategy adds the tagMatchStrategy to the get v1 incidents params
func (o *GetV1IncidentsParams) WithTagMatchStrategy(tagMatchStrategy *string) *GetV1IncidentsParams {
	o.SetTagMatchStrategy(tagMatchStrategy)
	return o
}

// SetTagMatchStrategy adds the tagMatchStrategy to the get v1 incidents params
func (o *GetV1IncidentsParams) SetTagMatchStrategy(tagMatchStrategy *string) {
	o.TagMatchStrategy = tagMatchStrategy
}

// WithTags adds the tags to the get v1 incidents params
func (o *GetV1IncidentsParams) WithTags(tags *string) *GetV1IncidentsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get v1 incidents params
func (o *GetV1IncidentsParams) SetTags(tags *string) {
	o.Tags = tags
}

// WithTeams adds the teams to the get v1 incidents params
func (o *GetV1IncidentsParams) WithTeams(teams *string) *GetV1IncidentsParams {
	o.SetTeams(teams)
	return o
}

// SetTeams adds the teams to the get v1 incidents params
func (o *GetV1IncidentsParams) SetTeams(teams *string) {
	o.Teams = teams
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IncidentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Conditions != nil {

		// query param conditions
		var qrConditions string

		if o.Conditions != nil {
			qrConditions = *o.Conditions
		}
		qConditions := qrConditions
		if qConditions != "" {

			if err := r.SetQueryParam("conditions", qConditions); err != nil {
				return err
			}
		}
	}

	if o.CurrentMilestones != nil {

		// query param current_milestones
		var qrCurrentMilestones string

		if o.CurrentMilestones != nil {
			qrCurrentMilestones = *o.CurrentMilestones
		}
		qCurrentMilestones := qrCurrentMilestones
		if qCurrentMilestones != "" {

			if err := r.SetQueryParam("current_milestones", qCurrentMilestones); err != nil {
				return err
			}
		}
	}

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.Date

		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {

			if err := r.SetQueryParam("end_date", qEndDate); err != nil {
				return err
			}
		}
	}

	if o.Environments != nil {

		// query param environments
		var qrEnvironments string

		if o.Environments != nil {
			qrEnvironments = *o.Environments
		}
		qEnvironments := qrEnvironments
		if qEnvironments != "" {

			if err := r.SetQueryParam("environments", qEnvironments); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

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

	if o.SavedSearchID != nil {

		// query param saved_search_id
		var qrSavedSearchID string

		if o.SavedSearchID != nil {
			qrSavedSearchID = *o.SavedSearchID
		}
		qSavedSearchID := qrSavedSearchID
		if qSavedSearchID != "" {

			if err := r.SetQueryParam("saved_search_id", qSavedSearchID); err != nil {
				return err
			}
		}
	}

	if o.Services != nil {

		// query param services
		var qrServices string

		if o.Services != nil {
			qrServices = *o.Services
		}
		qServices := qrServices
		if qServices != "" {

			if err := r.SetQueryParam("services", qServices); err != nil {
				return err
			}
		}
	}

	if o.Severities != nil {

		// query param severities
		var qrSeverities string

		if o.Severities != nil {
			qrSeverities = *o.Severities
		}
		qSeverities := qrSeverities
		if qSeverities != "" {

			if err := r.SetQueryParam("severities", qSeverities); err != nil {
				return err
			}
		}
	}

	if o.SeverityNotSet != nil {

		// query param severity_not_set
		var qrSeverityNotSet bool

		if o.SeverityNotSet != nil {
			qrSeverityNotSet = *o.SeverityNotSet
		}
		qSeverityNotSet := swag.FormatBool(qrSeverityNotSet)
		if qSeverityNotSet != "" {

			if err := r.SetQueryParam("severity_not_set", qSeverityNotSet); err != nil {
				return err
			}
		}
	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.Date

		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {

			if err := r.SetQueryParam("start_date", qStartDate); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.TagMatchStrategy != nil {

		// query param tag_match_strategy
		var qrTagMatchStrategy string

		if o.TagMatchStrategy != nil {
			qrTagMatchStrategy = *o.TagMatchStrategy
		}
		qTagMatchStrategy := qrTagMatchStrategy
		if qTagMatchStrategy != "" {

			if err := r.SetQueryParam("tag_match_strategy", qTagMatchStrategy); err != nil {
				return err
			}
		}
	}

	if o.Tags != nil {

		// query param tags
		var qrTags string

		if o.Tags != nil {
			qrTags = *o.Tags
		}
		qTags := qrTags
		if qTags != "" {

			if err := r.SetQueryParam("tags", qTags); err != nil {
				return err
			}
		}
	}

	if o.Teams != nil {

		// query param teams
		var qrTeams string

		if o.Teams != nil {
			qrTeams = *o.Teams
		}
		qTeams := qrTeams
		if qTeams != "" {

			if err := r.SetQueryParam("teams", qTeams); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
