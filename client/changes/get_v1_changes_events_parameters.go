// Code generated by go-swagger; DO NOT EDIT.

package changes

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

// NewGetV1ChangesEventsParams creates a new GetV1ChangesEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ChangesEventsParams() *GetV1ChangesEventsParams {
	return &GetV1ChangesEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ChangesEventsParamsWithTimeout creates a new GetV1ChangesEventsParams object
// with the ability to set a timeout on a request.
func NewGetV1ChangesEventsParamsWithTimeout(timeout time.Duration) *GetV1ChangesEventsParams {
	return &GetV1ChangesEventsParams{
		timeout: timeout,
	}
}

// NewGetV1ChangesEventsParamsWithContext creates a new GetV1ChangesEventsParams object
// with the ability to set a context for a request.
func NewGetV1ChangesEventsParamsWithContext(ctx context.Context) *GetV1ChangesEventsParams {
	return &GetV1ChangesEventsParams{
		Context: ctx,
	}
}

// NewGetV1ChangesEventsParamsWithHTTPClient creates a new GetV1ChangesEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ChangesEventsParamsWithHTTPClient(client *http.Client) *GetV1ChangesEventsParams {
	return &GetV1ChangesEventsParams{
		HTTPClient: client,
	}
}

/*
GetV1ChangesEventsParams contains all the parameters to send to the API endpoint

	for the get v1 changes events operation.

	Typically these are written to a http.Request.
*/
type GetV1ChangesEventsParams struct {

	/* EndsAt.

	   The end time to return change events up to

	   Format: date-time
	*/
	EndsAt *strfmt.DateTime

	/* Environments.

	   A comma separated list of environment IDs
	*/
	Environments *string

	/* Labels.

	   A comma separated list of label key / values in the format of "key=value,key2=value2". To filter change events that have a key (with no specific value), omit the value
	*/
	Labels *string

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   A text query for change events
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

	/* StartsAt.

	   The start time to start returning change events from
	*/
	StartsAt *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 changes events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ChangesEventsParams) WithDefaults() *GetV1ChangesEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 changes events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ChangesEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithTimeout(timeout time.Duration) *GetV1ChangesEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithContext(ctx context.Context) *GetV1ChangesEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithHTTPClient(client *http.Client) *GetV1ChangesEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndsAt adds the endsAt to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithEndsAt(endsAt *strfmt.DateTime) *GetV1ChangesEventsParams {
	o.SetEndsAt(endsAt)
	return o
}

// SetEndsAt adds the endsAt to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetEndsAt(endsAt *strfmt.DateTime) {
	o.EndsAt = endsAt
}

// WithEnvironments adds the environments to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithEnvironments(environments *string) *GetV1ChangesEventsParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetEnvironments(environments *string) {
	o.Environments = environments
}

// WithLabels adds the labels to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithLabels(labels *string) *GetV1ChangesEventsParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetLabels(labels *string) {
	o.Labels = labels
}

// WithPage adds the page to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithPage(page *int32) *GetV1ChangesEventsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithPerPage(perPage *int32) *GetV1ChangesEventsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithQuery(query *string) *GetV1ChangesEventsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSavedSearchID adds the savedSearchID to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithSavedSearchID(savedSearchID *string) *GetV1ChangesEventsParams {
	o.SetSavedSearchID(savedSearchID)
	return o
}

// SetSavedSearchID adds the savedSearchId to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetSavedSearchID(savedSearchID *string) {
	o.SavedSearchID = savedSearchID
}

// WithServices adds the services to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithServices(services *string) *GetV1ChangesEventsParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetServices(services *string) {
	o.Services = services
}

// WithStartsAt adds the startsAt to the get v1 changes events params
func (o *GetV1ChangesEventsParams) WithStartsAt(startsAt *string) *GetV1ChangesEventsParams {
	o.SetStartsAt(startsAt)
	return o
}

// SetStartsAt adds the startsAt to the get v1 changes events params
func (o *GetV1ChangesEventsParams) SetStartsAt(startsAt *string) {
	o.StartsAt = startsAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ChangesEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndsAt != nil {

		// query param ends_at
		var qrEndsAt strfmt.DateTime

		if o.EndsAt != nil {
			qrEndsAt = *o.EndsAt
		}
		qEndsAt := qrEndsAt.String()
		if qEndsAt != "" {

			if err := r.SetQueryParam("ends_at", qEndsAt); err != nil {
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

	if o.Labels != nil {

		// query param labels
		var qrLabels string

		if o.Labels != nil {
			qrLabels = *o.Labels
		}
		qLabels := qrLabels
		if qLabels != "" {

			if err := r.SetQueryParam("labels", qLabels); err != nil {
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

	if o.StartsAt != nil {

		// query param starts_at
		var qrStartsAt string

		if o.StartsAt != nil {
			qrStartsAt = *o.StartsAt
		}
		qStartsAt := qrStartsAt
		if qStartsAt != "" {

			if err := r.SetQueryParam("starts_at", qStartsAt); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
