// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewGetV1MetricsMilestoneFunnelParams creates a new GetV1MetricsMilestoneFunnelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1MetricsMilestoneFunnelParams() *GetV1MetricsMilestoneFunnelParams {
	return &GetV1MetricsMilestoneFunnelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1MetricsMilestoneFunnelParamsWithTimeout creates a new GetV1MetricsMilestoneFunnelParams object
// with the ability to set a timeout on a request.
func NewGetV1MetricsMilestoneFunnelParamsWithTimeout(timeout time.Duration) *GetV1MetricsMilestoneFunnelParams {
	return &GetV1MetricsMilestoneFunnelParams{
		timeout: timeout,
	}
}

// NewGetV1MetricsMilestoneFunnelParamsWithContext creates a new GetV1MetricsMilestoneFunnelParams object
// with the ability to set a context for a request.
func NewGetV1MetricsMilestoneFunnelParamsWithContext(ctx context.Context) *GetV1MetricsMilestoneFunnelParams {
	return &GetV1MetricsMilestoneFunnelParams{
		Context: ctx,
	}
}

// NewGetV1MetricsMilestoneFunnelParamsWithHTTPClient creates a new GetV1MetricsMilestoneFunnelParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1MetricsMilestoneFunnelParamsWithHTTPClient(client *http.Client) *GetV1MetricsMilestoneFunnelParams {
	return &GetV1MetricsMilestoneFunnelParams{
		HTTPClient: client,
	}
}

/*
GetV1MetricsMilestoneFunnelParams contains all the parameters to send to the API endpoint

	for the get v1 metrics milestone funnel operation.

	Typically these are written to a http.Request.
*/
type GetV1MetricsMilestoneFunnelParams struct {

	/* Archived.

	   Return archived incidents
	*/
	Archived *bool

	/* AssignedTeams.

	   A comma separated list of IDs for assigned teams or 'is_empty' to filter for incidents with no active team assignments
	*/
	AssignedTeams *string

	/* ClosedAtOrAfter.

	   Filters for incidents that were closed at or after this time

	   Format: date-time
	*/
	ClosedAtOrAfter *strfmt.DateTime

	/* ClosedAtOrBefore.

	   Filters for incidents that were closed at or before this time

	   Format: date-time
	*/
	ClosedAtOrBefore *strfmt.DateTime

	/* Conditions.

	   A JSON string that defines 'logic' and 'user_data'
	*/
	Conditions *string

	/* CreatedAtOrAfter.

	   Filters for incidents that were created at or after this time

	   Format: date-time
	*/
	CreatedAtOrAfter *strfmt.DateTime

	/* CreatedAtOrBefore.

	   Filters for incidents that were created at or before this time

	   Format: date-time
	*/
	CreatedAtOrBefore *strfmt.DateTime

	/* CurrentMilestones.

	   A comma separated list of current milestones
	*/
	CurrentMilestones *string

	/* EndDate.

	   Filters for incidents that started on or before this date

	   Format: date
	*/
	EndDate *strfmt.Date

	/* Environments.

	   A comma separated list of environment IDs or 'is_empty' to filter for incidents with no impacted environments
	*/
	Environments *string

	/* ExcludedInfrastructureIds.

	   A comma separated list of infrastructure IDs. Returns incidents that do not have the following infrastructure ids associated with them.
	*/
	ExcludedInfrastructureIds *string

	/* Functionalities.

	   A comma separated list of functionality IDs or 'is_empty' to filter for incidents with no impacted functionalities
	*/
	Functionalities *string

	// GroupBy.
	GroupBy []string

	/* IncidentTypeID.

	   A comma separated list of incident type IDs
	*/
	IncidentTypeID *string

	/* Name.

	   A query to search incidents by their name
	*/
	Name *string

	/* Priorities.

	   A text value of priority
	*/
	Priorities *string

	/* PriorityNotSet.

	   Flag for including incidents where priority has not been set
	*/
	PriorityNotSet *bool

	/* Query.

	   A text query for an incident that searches on name, summary, and desciption
	*/
	Query *string

	/* ResolvedAtOrAfter.

	   Filters for incidents that were resolved at or after this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.

	   Format: date-time
	*/
	ResolvedAtOrAfter *strfmt.DateTime

	/* ResolvedAtOrBefore.

	   Filters for incidents that were resolved at or before this time. Combine this with the `current_milestones` parameter if you wish to omit incidents that were re-opened and are still active.

	   Format: date-time
	*/
	ResolvedAtOrBefore *strfmt.DateTime

	/* SavedSearchID.

	   The id of a previously saved search.
	*/
	SavedSearchID *string

	/* Services.

	   A comma separated list of service IDs or 'is_empty' to filter for incidents with no impacted services
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

	   Filters for incidents that started on or after this date

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

	/* UpdatedAfter.

	   Filters for incidents that were updated after this date

	   Format: date-time
	*/
	UpdatedAfter *strfmt.DateTime

	/* UpdatedBefore.

	   Filters for incidents that were updated before this date

	   Format: date-time
	*/
	UpdatedBefore *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 metrics milestone funnel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MetricsMilestoneFunnelParams) WithDefaults() *GetV1MetricsMilestoneFunnelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 metrics milestone funnel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MetricsMilestoneFunnelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithTimeout(timeout time.Duration) *GetV1MetricsMilestoneFunnelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithContext(ctx context.Context) *GetV1MetricsMilestoneFunnelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithHTTPClient(client *http.Client) *GetV1MetricsMilestoneFunnelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchived adds the archived to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithArchived(archived *bool) *GetV1MetricsMilestoneFunnelParams {
	o.SetArchived(archived)
	return o
}

// SetArchived adds the archived to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetArchived(archived *bool) {
	o.Archived = archived
}

// WithAssignedTeams adds the assignedTeams to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithAssignedTeams(assignedTeams *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetAssignedTeams(assignedTeams)
	return o
}

// SetAssignedTeams adds the assignedTeams to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetAssignedTeams(assignedTeams *string) {
	o.AssignedTeams = assignedTeams
}

// WithClosedAtOrAfter adds the closedAtOrAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithClosedAtOrAfter(closedAtOrAfter *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetClosedAtOrAfter(closedAtOrAfter)
	return o
}

// SetClosedAtOrAfter adds the closedAtOrAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetClosedAtOrAfter(closedAtOrAfter *strfmt.DateTime) {
	o.ClosedAtOrAfter = closedAtOrAfter
}

// WithClosedAtOrBefore adds the closedAtOrBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithClosedAtOrBefore(closedAtOrBefore *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetClosedAtOrBefore(closedAtOrBefore)
	return o
}

// SetClosedAtOrBefore adds the closedAtOrBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetClosedAtOrBefore(closedAtOrBefore *strfmt.DateTime) {
	o.ClosedAtOrBefore = closedAtOrBefore
}

// WithConditions adds the conditions to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithConditions(conditions *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetConditions(conditions)
	return o
}

// SetConditions adds the conditions to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetConditions(conditions *string) {
	o.Conditions = conditions
}

// WithCreatedAtOrAfter adds the createdAtOrAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithCreatedAtOrAfter(createdAtOrAfter *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetCreatedAtOrAfter(createdAtOrAfter)
	return o
}

// SetCreatedAtOrAfter adds the createdAtOrAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetCreatedAtOrAfter(createdAtOrAfter *strfmt.DateTime) {
	o.CreatedAtOrAfter = createdAtOrAfter
}

// WithCreatedAtOrBefore adds the createdAtOrBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithCreatedAtOrBefore(createdAtOrBefore *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetCreatedAtOrBefore(createdAtOrBefore)
	return o
}

// SetCreatedAtOrBefore adds the createdAtOrBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetCreatedAtOrBefore(createdAtOrBefore *strfmt.DateTime) {
	o.CreatedAtOrBefore = createdAtOrBefore
}

// WithCurrentMilestones adds the currentMilestones to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithCurrentMilestones(currentMilestones *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetCurrentMilestones(currentMilestones)
	return o
}

// SetCurrentMilestones adds the currentMilestones to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetCurrentMilestones(currentMilestones *string) {
	o.CurrentMilestones = currentMilestones
}

// WithEndDate adds the endDate to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithEndDate(endDate *strfmt.Date) *GetV1MetricsMilestoneFunnelParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithEnvironments adds the environments to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithEnvironments(environments *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetEnvironments(environments *string) {
	o.Environments = environments
}

// WithExcludedInfrastructureIds adds the excludedInfrastructureIds to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithExcludedInfrastructureIds(excludedInfrastructureIds *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetExcludedInfrastructureIds(excludedInfrastructureIds)
	return o
}

// SetExcludedInfrastructureIds adds the excludedInfrastructureIds to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetExcludedInfrastructureIds(excludedInfrastructureIds *string) {
	o.ExcludedInfrastructureIds = excludedInfrastructureIds
}

// WithFunctionalities adds the functionalities to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithFunctionalities(functionalities *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetFunctionalities(functionalities)
	return o
}

// SetFunctionalities adds the functionalities to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetFunctionalities(functionalities *string) {
	o.Functionalities = functionalities
}

// WithGroupBy adds the groupBy to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithGroupBy(groupBy []string) *GetV1MetricsMilestoneFunnelParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetGroupBy(groupBy []string) {
	o.GroupBy = groupBy
}

// WithIncidentTypeID adds the incidentTypeID to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithIncidentTypeID(incidentTypeID *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetIncidentTypeID(incidentTypeID)
	return o
}

// SetIncidentTypeID adds the incidentTypeId to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetIncidentTypeID(incidentTypeID *string) {
	o.IncidentTypeID = incidentTypeID
}

// WithName adds the name to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithName(name *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetName(name *string) {
	o.Name = name
}

// WithPriorities adds the priorities to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithPriorities(priorities *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetPriorities(priorities)
	return o
}

// SetPriorities adds the priorities to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetPriorities(priorities *string) {
	o.Priorities = priorities
}

// WithPriorityNotSet adds the priorityNotSet to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithPriorityNotSet(priorityNotSet *bool) *GetV1MetricsMilestoneFunnelParams {
	o.SetPriorityNotSet(priorityNotSet)
	return o
}

// SetPriorityNotSet adds the priorityNotSet to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetPriorityNotSet(priorityNotSet *bool) {
	o.PriorityNotSet = priorityNotSet
}

// WithQuery adds the query to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithQuery(query *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetQuery(query *string) {
	o.Query = query
}

// WithResolvedAtOrAfter adds the resolvedAtOrAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithResolvedAtOrAfter(resolvedAtOrAfter *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetResolvedAtOrAfter(resolvedAtOrAfter)
	return o
}

// SetResolvedAtOrAfter adds the resolvedAtOrAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetResolvedAtOrAfter(resolvedAtOrAfter *strfmt.DateTime) {
	o.ResolvedAtOrAfter = resolvedAtOrAfter
}

// WithResolvedAtOrBefore adds the resolvedAtOrBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithResolvedAtOrBefore(resolvedAtOrBefore *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetResolvedAtOrBefore(resolvedAtOrBefore)
	return o
}

// SetResolvedAtOrBefore adds the resolvedAtOrBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetResolvedAtOrBefore(resolvedAtOrBefore *strfmt.DateTime) {
	o.ResolvedAtOrBefore = resolvedAtOrBefore
}

// WithSavedSearchID adds the savedSearchID to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithSavedSearchID(savedSearchID *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetSavedSearchID(savedSearchID)
	return o
}

// SetSavedSearchID adds the savedSearchId to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetSavedSearchID(savedSearchID *string) {
	o.SavedSearchID = savedSearchID
}

// WithServices adds the services to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithServices(services *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetServices(services *string) {
	o.Services = services
}

// WithSeverities adds the severities to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithSeverities(severities *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetSeverities(severities)
	return o
}

// SetSeverities adds the severities to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetSeverities(severities *string) {
	o.Severities = severities
}

// WithSeverityNotSet adds the severityNotSet to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithSeverityNotSet(severityNotSet *bool) *GetV1MetricsMilestoneFunnelParams {
	o.SetSeverityNotSet(severityNotSet)
	return o
}

// SetSeverityNotSet adds the severityNotSet to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetSeverityNotSet(severityNotSet *bool) {
	o.SeverityNotSet = severityNotSet
}

// WithStartDate adds the startDate to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithStartDate(startDate *strfmt.Date) *GetV1MetricsMilestoneFunnelParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithStatus adds the status to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithStatus(status *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetStatus(status *string) {
	o.Status = status
}

// WithTagMatchStrategy adds the tagMatchStrategy to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithTagMatchStrategy(tagMatchStrategy *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetTagMatchStrategy(tagMatchStrategy)
	return o
}

// SetTagMatchStrategy adds the tagMatchStrategy to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetTagMatchStrategy(tagMatchStrategy *string) {
	o.TagMatchStrategy = tagMatchStrategy
}

// WithTags adds the tags to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithTags(tags *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetTags(tags *string) {
	o.Tags = tags
}

// WithTeams adds the teams to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithTeams(teams *string) *GetV1MetricsMilestoneFunnelParams {
	o.SetTeams(teams)
	return o
}

// SetTeams adds the teams to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetTeams(teams *string) {
	o.Teams = teams
}

// WithUpdatedAfter adds the updatedAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithUpdatedAfter(updatedAfter *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetUpdatedAfter(updatedAfter)
	return o
}

// SetUpdatedAfter adds the updatedAfter to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetUpdatedAfter(updatedAfter *strfmt.DateTime) {
	o.UpdatedAfter = updatedAfter
}

// WithUpdatedBefore adds the updatedBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) WithUpdatedBefore(updatedBefore *strfmt.DateTime) *GetV1MetricsMilestoneFunnelParams {
	o.SetUpdatedBefore(updatedBefore)
	return o
}

// SetUpdatedBefore adds the updatedBefore to the get v1 metrics milestone funnel params
func (o *GetV1MetricsMilestoneFunnelParams) SetUpdatedBefore(updatedBefore *strfmt.DateTime) {
	o.UpdatedBefore = updatedBefore
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1MetricsMilestoneFunnelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Archived != nil {

		// query param archived
		var qrArchived bool

		if o.Archived != nil {
			qrArchived = *o.Archived
		}
		qArchived := swag.FormatBool(qrArchived)
		if qArchived != "" {

			if err := r.SetQueryParam("archived", qArchived); err != nil {
				return err
			}
		}
	}

	if o.AssignedTeams != nil {

		// query param assigned_teams
		var qrAssignedTeams string

		if o.AssignedTeams != nil {
			qrAssignedTeams = *o.AssignedTeams
		}
		qAssignedTeams := qrAssignedTeams
		if qAssignedTeams != "" {

			if err := r.SetQueryParam("assigned_teams", qAssignedTeams); err != nil {
				return err
			}
		}
	}

	if o.ClosedAtOrAfter != nil {

		// query param closed_at_or_after
		var qrClosedAtOrAfter strfmt.DateTime

		if o.ClosedAtOrAfter != nil {
			qrClosedAtOrAfter = *o.ClosedAtOrAfter
		}
		qClosedAtOrAfter := qrClosedAtOrAfter.String()
		if qClosedAtOrAfter != "" {

			if err := r.SetQueryParam("closed_at_or_after", qClosedAtOrAfter); err != nil {
				return err
			}
		}
	}

	if o.ClosedAtOrBefore != nil {

		// query param closed_at_or_before
		var qrClosedAtOrBefore strfmt.DateTime

		if o.ClosedAtOrBefore != nil {
			qrClosedAtOrBefore = *o.ClosedAtOrBefore
		}
		qClosedAtOrBefore := qrClosedAtOrBefore.String()
		if qClosedAtOrBefore != "" {

			if err := r.SetQueryParam("closed_at_or_before", qClosedAtOrBefore); err != nil {
				return err
			}
		}
	}

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

	if o.CreatedAtOrAfter != nil {

		// query param created_at_or_after
		var qrCreatedAtOrAfter strfmt.DateTime

		if o.CreatedAtOrAfter != nil {
			qrCreatedAtOrAfter = *o.CreatedAtOrAfter
		}
		qCreatedAtOrAfter := qrCreatedAtOrAfter.String()
		if qCreatedAtOrAfter != "" {

			if err := r.SetQueryParam("created_at_or_after", qCreatedAtOrAfter); err != nil {
				return err
			}
		}
	}

	if o.CreatedAtOrBefore != nil {

		// query param created_at_or_before
		var qrCreatedAtOrBefore strfmt.DateTime

		if o.CreatedAtOrBefore != nil {
			qrCreatedAtOrBefore = *o.CreatedAtOrBefore
		}
		qCreatedAtOrBefore := qrCreatedAtOrBefore.String()
		if qCreatedAtOrBefore != "" {

			if err := r.SetQueryParam("created_at_or_before", qCreatedAtOrBefore); err != nil {
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

	if o.ExcludedInfrastructureIds != nil {

		// query param excluded_infrastructure_ids
		var qrExcludedInfrastructureIds string

		if o.ExcludedInfrastructureIds != nil {
			qrExcludedInfrastructureIds = *o.ExcludedInfrastructureIds
		}
		qExcludedInfrastructureIds := qrExcludedInfrastructureIds
		if qExcludedInfrastructureIds != "" {

			if err := r.SetQueryParam("excluded_infrastructure_ids", qExcludedInfrastructureIds); err != nil {
				return err
			}
		}
	}

	if o.Functionalities != nil {

		// query param functionalities
		var qrFunctionalities string

		if o.Functionalities != nil {
			qrFunctionalities = *o.Functionalities
		}
		qFunctionalities := qrFunctionalities
		if qFunctionalities != "" {

			if err := r.SetQueryParam("functionalities", qFunctionalities); err != nil {
				return err
			}
		}
	}

	if o.GroupBy != nil {

		// binding items for group_by
		joinedGroupBy := o.bindParamGroupBy(reg)

		// form array param group_by
		if err := r.SetFormParam("group_by", joinedGroupBy...); err != nil {
			return err
		}
	}

	if o.IncidentTypeID != nil {

		// query param incident_type_id
		var qrIncidentTypeID string

		if o.IncidentTypeID != nil {
			qrIncidentTypeID = *o.IncidentTypeID
		}
		qIncidentTypeID := qrIncidentTypeID
		if qIncidentTypeID != "" {

			if err := r.SetQueryParam("incident_type_id", qIncidentTypeID); err != nil {
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

	if o.Priorities != nil {

		// query param priorities
		var qrPriorities string

		if o.Priorities != nil {
			qrPriorities = *o.Priorities
		}
		qPriorities := qrPriorities
		if qPriorities != "" {

			if err := r.SetQueryParam("priorities", qPriorities); err != nil {
				return err
			}
		}
	}

	if o.PriorityNotSet != nil {

		// query param priority_not_set
		var qrPriorityNotSet bool

		if o.PriorityNotSet != nil {
			qrPriorityNotSet = *o.PriorityNotSet
		}
		qPriorityNotSet := swag.FormatBool(qrPriorityNotSet)
		if qPriorityNotSet != "" {

			if err := r.SetQueryParam("priority_not_set", qPriorityNotSet); err != nil {
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

	if o.ResolvedAtOrAfter != nil {

		// query param resolved_at_or_after
		var qrResolvedAtOrAfter strfmt.DateTime

		if o.ResolvedAtOrAfter != nil {
			qrResolvedAtOrAfter = *o.ResolvedAtOrAfter
		}
		qResolvedAtOrAfter := qrResolvedAtOrAfter.String()
		if qResolvedAtOrAfter != "" {

			if err := r.SetQueryParam("resolved_at_or_after", qResolvedAtOrAfter); err != nil {
				return err
			}
		}
	}

	if o.ResolvedAtOrBefore != nil {

		// query param resolved_at_or_before
		var qrResolvedAtOrBefore strfmt.DateTime

		if o.ResolvedAtOrBefore != nil {
			qrResolvedAtOrBefore = *o.ResolvedAtOrBefore
		}
		qResolvedAtOrBefore := qrResolvedAtOrBefore.String()
		if qResolvedAtOrBefore != "" {

			if err := r.SetQueryParam("resolved_at_or_before", qResolvedAtOrBefore); err != nil {
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

	if o.UpdatedAfter != nil {

		// query param updated_after
		var qrUpdatedAfter strfmt.DateTime

		if o.UpdatedAfter != nil {
			qrUpdatedAfter = *o.UpdatedAfter
		}
		qUpdatedAfter := qrUpdatedAfter.String()
		if qUpdatedAfter != "" {

			if err := r.SetQueryParam("updated_after", qUpdatedAfter); err != nil {
				return err
			}
		}
	}

	if o.UpdatedBefore != nil {

		// query param updated_before
		var qrUpdatedBefore strfmt.DateTime

		if o.UpdatedBefore != nil {
			qrUpdatedBefore = *o.UpdatedBefore
		}
		qUpdatedBefore := qrUpdatedBefore.String()
		if qUpdatedBefore != "" {

			if err := r.SetQueryParam("updated_before", qUpdatedBefore); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetV1MetricsMilestoneFunnel binds the parameter group_by
func (o *GetV1MetricsMilestoneFunnelParams) bindParamGroupBy(formats strfmt.Registry) []string {
	groupByIR := o.GroupBy

	var groupByIC []string
	for _, groupByIIR := range groupByIR { // explode []string

		groupByIIV := groupByIIR // string as string
		groupByIC = append(groupByIC, groupByIIV)
	}

	// items.CollectionFormat: ""
	groupByIS := swag.JoinByFormat(groupByIC, "")

	return groupByIS
}
