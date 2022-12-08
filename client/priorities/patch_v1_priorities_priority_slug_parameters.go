// Code generated by go-swagger; DO NOT EDIT.

package priorities

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

// NewPatchV1PrioritiesPrioritySlugParams creates a new PatchV1PrioritiesPrioritySlugParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1PrioritiesPrioritySlugParams() *PatchV1PrioritiesPrioritySlugParams {
	return &PatchV1PrioritiesPrioritySlugParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PrioritiesPrioritySlugParamsWithTimeout creates a new PatchV1PrioritiesPrioritySlugParams object
// with the ability to set a timeout on a request.
func NewPatchV1PrioritiesPrioritySlugParamsWithTimeout(timeout time.Duration) *PatchV1PrioritiesPrioritySlugParams {
	return &PatchV1PrioritiesPrioritySlugParams{
		timeout: timeout,
	}
}

// NewPatchV1PrioritiesPrioritySlugParamsWithContext creates a new PatchV1PrioritiesPrioritySlugParams object
// with the ability to set a context for a request.
func NewPatchV1PrioritiesPrioritySlugParamsWithContext(ctx context.Context) *PatchV1PrioritiesPrioritySlugParams {
	return &PatchV1PrioritiesPrioritySlugParams{
		Context: ctx,
	}
}

// NewPatchV1PrioritiesPrioritySlugParamsWithHTTPClient creates a new PatchV1PrioritiesPrioritySlugParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1PrioritiesPrioritySlugParamsWithHTTPClient(client *http.Client) *PatchV1PrioritiesPrioritySlugParams {
	return &PatchV1PrioritiesPrioritySlugParams{
		HTTPClient: client,
	}
}

/*
PatchV1PrioritiesPrioritySlugParams contains all the parameters to send to the API endpoint

	for the patch v1 priorities priority slug operation.

	Typically these are written to a http.Request.
*/
type PatchV1PrioritiesPrioritySlugParams struct {

	// V1Priorities.
	V1Priorities *models.PatchV1Priorities

	// PrioritySlug.
	PrioritySlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 priorities priority slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PrioritiesPrioritySlugParams) WithDefaults() *PatchV1PrioritiesPrioritySlugParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 priorities priority slug params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PrioritiesPrioritySlugParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) WithTimeout(timeout time.Duration) *PatchV1PrioritiesPrioritySlugParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) WithContext(ctx context.Context) *PatchV1PrioritiesPrioritySlugParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) WithHTTPClient(client *http.Client) *PatchV1PrioritiesPrioritySlugParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Priorities adds the v1Priorities to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) WithV1Priorities(v1Priorities *models.PatchV1Priorities) *PatchV1PrioritiesPrioritySlugParams {
	o.SetV1Priorities(v1Priorities)
	return o
}

// SetV1Priorities adds the v1Priorities to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) SetV1Priorities(v1Priorities *models.PatchV1Priorities) {
	o.V1Priorities = v1Priorities
}

// WithPrioritySlug adds the prioritySlug to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) WithPrioritySlug(prioritySlug string) *PatchV1PrioritiesPrioritySlugParams {
	o.SetPrioritySlug(prioritySlug)
	return o
}

// SetPrioritySlug adds the prioritySlug to the patch v1 priorities priority slug params
func (o *PatchV1PrioritiesPrioritySlugParams) SetPrioritySlug(prioritySlug string) {
	o.PrioritySlug = prioritySlug
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PrioritiesPrioritySlugParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Priorities != nil {
		if err := r.SetBodyParam(o.V1Priorities); err != nil {
			return err
		}
	}

	// path param priority_slug
	if err := r.SetPathParam("priority_slug", o.PrioritySlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}