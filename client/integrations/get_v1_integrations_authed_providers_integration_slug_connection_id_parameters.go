// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams creates a new GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams() *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	return &GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParamsWithTimeout creates a new GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams object
// with the ability to set a timeout on a request.
func NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	return &GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams{
		timeout: timeout,
	}
}

// NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParamsWithContext creates a new GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams object
// with the ability to set a context for a request.
func NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParamsWithContext(ctx context.Context) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	return &GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams{
		Context: ctx,
	}
}

// NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParamsWithHTTPClient creates a new GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	return &GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams{
		HTTPClient: client,
	}
}

/*
GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams contains all the parameters to send to the API endpoint

	for the get v1 integrations authed providers integration slug connection Id operation.

	Typically these are written to a http.Request.
*/
type GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams struct {

	/* ConnectionID.

	   Connection ID
	*/
	ConnectionID string

	/* IntegrationSlug.

	   Integration slug
	*/
	IntegrationSlug string

	/* Query.

	   Query for users by name
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 integrations authed providers integration slug connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WithDefaults() *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 integrations authed providers integration slug connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WithContext(ctx context.Context) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WithConnectionID(connectionID string) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) SetConnectionID(connectionID string) {
	o.ConnectionID = connectionID
}

// WithIntegrationSlug adds the integrationSlug to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WithIntegrationSlug(integrationSlug string) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	o.SetIntegrationSlug(integrationSlug)
	return o
}

// SetIntegrationSlug adds the integrationSlug to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) SetIntegrationSlug(integrationSlug string) {
	o.IntegrationSlug = integrationSlug
}

// WithQuery adds the query to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WithQuery(query *string) *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 integrations authed providers integration slug connection Id params
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsAuthedProvidersIntegrationSlugConnectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID); err != nil {
		return err
	}

	// path param integration_slug
	if err := r.SetPathParam("integration_slug", o.IntegrationSlug); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
