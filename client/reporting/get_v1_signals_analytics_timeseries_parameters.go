// Code generated by go-swagger; DO NOT EDIT.

package reporting

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

// NewGetV1SignalsAnalyticsTimeseriesParams creates a new GetV1SignalsAnalyticsTimeseriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SignalsAnalyticsTimeseriesParams() *GetV1SignalsAnalyticsTimeseriesParams {
	return &GetV1SignalsAnalyticsTimeseriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SignalsAnalyticsTimeseriesParamsWithTimeout creates a new GetV1SignalsAnalyticsTimeseriesParams object
// with the ability to set a timeout on a request.
func NewGetV1SignalsAnalyticsTimeseriesParamsWithTimeout(timeout time.Duration) *GetV1SignalsAnalyticsTimeseriesParams {
	return &GetV1SignalsAnalyticsTimeseriesParams{
		timeout: timeout,
	}
}

// NewGetV1SignalsAnalyticsTimeseriesParamsWithContext creates a new GetV1SignalsAnalyticsTimeseriesParams object
// with the ability to set a context for a request.
func NewGetV1SignalsAnalyticsTimeseriesParamsWithContext(ctx context.Context) *GetV1SignalsAnalyticsTimeseriesParams {
	return &GetV1SignalsAnalyticsTimeseriesParams{
		Context: ctx,
	}
}

// NewGetV1SignalsAnalyticsTimeseriesParamsWithHTTPClient creates a new GetV1SignalsAnalyticsTimeseriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SignalsAnalyticsTimeseriesParamsWithHTTPClient(client *http.Client) *GetV1SignalsAnalyticsTimeseriesParams {
	return &GetV1SignalsAnalyticsTimeseriesParams{
		HTTPClient: client,
	}
}

/*
GetV1SignalsAnalyticsTimeseriesParams contains all the parameters to send to the API endpoint

	for the get v1 signals analytics timeseries operation.

	Typically these are written to a http.Request.
*/
type GetV1SignalsAnalyticsTimeseriesParams struct {

	/* Bucket.

	   String that determines how records are grouped
	*/
	Bucket *string

	/* EndDate.

	   The end date to return metrics from

	   Format: date-time
	*/
	EndDate *strfmt.DateTime

	/* Environments.

	   A comma separated list of environment IDs
	*/
	Environments *string

	/* GroupBy.

	   String that determines how records are grouped
	*/
	GroupBy *string

	/* Services.

	   A comma separated list of service IDs
	*/
	Services *string

	/* SignalRules.

	   A comma separated list of signal rule IDs
	*/
	SignalRules *string

	/* SortBy.

	   String that determines how records are sorted
	*/
	SortBy *string

	/* SortDirection.

	   String that determines how records are sorted
	*/
	SortDirection *string

	/* StartDate.

	   The start date to return metrics from

	   Format: date-time
	*/
	StartDate *strfmt.DateTime

	/* Tags.

	   A comma separated list of tags
	*/
	Tags *string

	/* Teams.

	   A comma separated list of team IDs
	*/
	Teams *string

	/* Users.

	   A comma separated list of user IDs
	*/
	Users *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 signals analytics timeseries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithDefaults() *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 signals analytics timeseries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithTimeout(timeout time.Duration) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithContext(ctx context.Context) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithHTTPClient(client *http.Client) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucket adds the bucket to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithBucket(bucket *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetBucket(bucket)
	return o
}

// SetBucket adds the bucket to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetBucket(bucket *string) {
	o.Bucket = bucket
}

// WithEndDate adds the endDate to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithEndDate(endDate *strfmt.DateTime) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithEnvironments adds the environments to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithEnvironments(environments *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetEnvironments(environments *string) {
	o.Environments = environments
}

// WithGroupBy adds the groupBy to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithGroupBy(groupBy *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetGroupBy(groupBy *string) {
	o.GroupBy = groupBy
}

// WithServices adds the services to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithServices(services *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetServices(services *string) {
	o.Services = services
}

// WithSignalRules adds the signalRules to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithSignalRules(signalRules *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetSignalRules(signalRules)
	return o
}

// SetSignalRules adds the signalRules to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetSignalRules(signalRules *string) {
	o.SignalRules = signalRules
}

// WithSortBy adds the sortBy to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithSortBy(sortBy *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithSortDirection(sortDirection *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithStartDate adds the startDate to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithStartDate(startDate *strfmt.DateTime) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithTags adds the tags to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithTags(tags *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetTags(tags *string) {
	o.Tags = tags
}

// WithTeams adds the teams to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithTeams(teams *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetTeams(teams)
	return o
}

// SetTeams adds the teams to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetTeams(teams *string) {
	o.Teams = teams
}

// WithUsers adds the users to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) WithUsers(users *string) *GetV1SignalsAnalyticsTimeseriesParams {
	o.SetUsers(users)
	return o
}

// SetUsers adds the users to the get v1 signals analytics timeseries params
func (o *GetV1SignalsAnalyticsTimeseriesParams) SetUsers(users *string) {
	o.Users = users
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SignalsAnalyticsTimeseriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Bucket != nil {

		// query param bucket
		var qrBucket string

		if o.Bucket != nil {
			qrBucket = *o.Bucket
		}
		qBucket := qrBucket
		if qBucket != "" {

			if err := r.SetQueryParam("bucket", qBucket); err != nil {
				return err
			}
		}
	}

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.DateTime

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

	if o.GroupBy != nil {

		// query param group_by
		var qrGroupBy string

		if o.GroupBy != nil {
			qrGroupBy = *o.GroupBy
		}
		qGroupBy := qrGroupBy
		if qGroupBy != "" {

			if err := r.SetQueryParam("group_by", qGroupBy); err != nil {
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

	if o.SignalRules != nil {

		// query param signal_rules
		var qrSignalRules string

		if o.SignalRules != nil {
			qrSignalRules = *o.SignalRules
		}
		qSignalRules := qrSignalRules
		if qSignalRules != "" {

			if err := r.SetQueryParam("signal_rules", qSignalRules); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sort_by
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sort_by", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortDirection != nil {

		// query param sort_direction
		var qrSortDirection string

		if o.SortDirection != nil {
			qrSortDirection = *o.SortDirection
		}
		qSortDirection := qrSortDirection
		if qSortDirection != "" {

			if err := r.SetQueryParam("sort_direction", qSortDirection); err != nil {
				return err
			}
		}
	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.DateTime

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

	if o.Users != nil {

		// query param users
		var qrUsers string

		if o.Users != nil {
			qrUsers = *o.Users
		}
		qUsers := qrUsers
		if qUsers != "" {

			if err := r.SetQueryParam("users", qUsers); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
