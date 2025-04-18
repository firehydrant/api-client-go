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

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1SavedSearchesResourceTypeSavedSearchIDParams creates a new PatchV1SavedSearchesResourceTypeSavedSearchIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1SavedSearchesResourceTypeSavedSearchIDParams() *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &PatchV1SavedSearchesResourceTypeSavedSearchIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1SavedSearchesResourceTypeSavedSearchIDParamsWithTimeout creates a new PatchV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1SavedSearchesResourceTypeSavedSearchIDParamsWithTimeout(timeout time.Duration) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &PatchV1SavedSearchesResourceTypeSavedSearchIDParams{
		timeout: timeout,
	}
}

// NewPatchV1SavedSearchesResourceTypeSavedSearchIDParamsWithContext creates a new PatchV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the ability to set a context for a request.
func NewPatchV1SavedSearchesResourceTypeSavedSearchIDParamsWithContext(ctx context.Context) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &PatchV1SavedSearchesResourceTypeSavedSearchIDParams{
		Context: ctx,
	}
}

// NewPatchV1SavedSearchesResourceTypeSavedSearchIDParamsWithHTTPClient creates a new PatchV1SavedSearchesResourceTypeSavedSearchIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1SavedSearchesResourceTypeSavedSearchIDParamsWithHTTPClient(client *http.Client) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	return &PatchV1SavedSearchesResourceTypeSavedSearchIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1SavedSearchesResourceTypeSavedSearchIDParams contains all the parameters to send to the API endpoint

	for the patch v1 saved searches resource type saved search Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1SavedSearchesResourceTypeSavedSearchIDParams struct {

	// PatchV1SavedSearchesResourceTypeSavedSearchID.
	PatchV1SavedSearchesResourceTypeSavedSearchID *models.PatchV1SavedSearchesResourceTypeSavedSearchID

	// ResourceType.
	ResourceType string

	// SavedSearchID.
	SavedSearchID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 saved searches resource type saved search Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WithDefaults() *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 saved searches resource type saved search Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WithTimeout(timeout time.Duration) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WithContext(ctx context.Context) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WithHTTPClient(client *http.Client) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchV1SavedSearchesResourceTypeSavedSearchID adds the patchV1SavedSearchesResourceTypeSavedSearchID to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WithPatchV1SavedSearchesResourceTypeSavedSearchID(patchV1SavedSearchesResourceTypeSavedSearchID *models.PatchV1SavedSearchesResourceTypeSavedSearchID) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetPatchV1SavedSearchesResourceTypeSavedSearchID(patchV1SavedSearchesResourceTypeSavedSearchID)
	return o
}

// SetPatchV1SavedSearchesResourceTypeSavedSearchID adds the patchV1SavedSearchesResourceTypeSavedSearchId to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) SetPatchV1SavedSearchesResourceTypeSavedSearchID(patchV1SavedSearchesResourceTypeSavedSearchID *models.PatchV1SavedSearchesResourceTypeSavedSearchID) {
	o.PatchV1SavedSearchesResourceTypeSavedSearchID = patchV1SavedSearchesResourceTypeSavedSearchID
}

// WithResourceType adds the resourceType to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WithResourceType(resourceType string) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithSavedSearchID adds the savedSearchID to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WithSavedSearchID(savedSearchID string) *PatchV1SavedSearchesResourceTypeSavedSearchIDParams {
	o.SetSavedSearchID(savedSearchID)
	return o
}

// SetSavedSearchID adds the savedSearchId to the patch v1 saved searches resource type saved search Id params
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) SetSavedSearchID(savedSearchID string) {
	o.SavedSearchID = savedSearchID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1SavedSearchesResourceTypeSavedSearchIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PatchV1SavedSearchesResourceTypeSavedSearchID != nil {
		if err := r.SetBodyParam(o.PatchV1SavedSearchesResourceTypeSavedSearchID); err != nil {
			return err
		}
	}

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
