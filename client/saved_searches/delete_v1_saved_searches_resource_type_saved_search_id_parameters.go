// Code generated by go-swagger; DO NOT EDIT.

package saved_searches

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

// NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParams creates a new DeleteV1SavedSearchesResourceTypeSavedSearchIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParams() *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &DeleteV1SavedSearchesResourceTypeSavedSearchIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParamsWithTimeout creates a new DeleteV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParamsWithTimeout(timeout time.Duration) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &DeleteV1SavedSearchesResourceTypeSavedSearchIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParamsWithContext creates a new DeleteV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the ability to set a context for a request.
func NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParamsWithContext(ctx context.Context) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &DeleteV1SavedSearchesResourceTypeSavedSearchIDParams{
		Context: ctx,
	}
}

// NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParamsWithHTTPClient creates a new DeleteV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1SavedSearchesResourceTypeSavedSearchIDParamsWithHTTPClient(client *http.Client) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &DeleteV1SavedSearchesResourceTypeSavedSearchIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1SavedSearchesResourceTypeSavedSearchIDParams contains all the parameters to send to the API endpoint

	for the delete v1 saved searches resource type saved search Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1SavedSearchesResourceTypeSavedSearchIDParams struct {

	// ResourceType.
	ResourceType string

	/* SavedSearchID.

	   ID of a saved search
	*/
	SavedSearchID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 saved searches resource type saved search Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) WithDefaults() *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 saved searches resource type saved search Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) WithTimeout(timeout time.Duration) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) WithContext(ctx context.Context) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) WithHTTPClient(client *http.Client) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceType adds the resourceType to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) WithResourceType(resourceType string) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithSavedSearchID adds the savedSearchID to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) WithSavedSearchID(savedSearchID string) *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetSavedSearchID(savedSearchID)
	return o
}

// SetSavedSearchID adds the savedSearchId to the delete v1 saved searches resource type saved search Id params
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) SetSavedSearchID(savedSearchID string) {
	o.SavedSearchID = savedSearchID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1SavedSearchesResourceTypeSavedSearchIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_type
	if err := r.SetPathParam("resource_type", o.ResourceType); err != nil {
		return err
	}

	// path param saved_search_id
	if err := r.SetPathParam("saved_search_id", o.SavedSearchID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
