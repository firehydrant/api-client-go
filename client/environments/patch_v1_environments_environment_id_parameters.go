// Code generated by go-swagger; DO NOT EDIT.

package environments

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

// NewPatchV1EnvironmentsEnvironmentIDParams creates a new PatchV1EnvironmentsEnvironmentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1EnvironmentsEnvironmentIDParams() *PatchV1EnvironmentsEnvironmentIDParams {
	return &PatchV1EnvironmentsEnvironmentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1EnvironmentsEnvironmentIDParamsWithTimeout creates a new PatchV1EnvironmentsEnvironmentIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1EnvironmentsEnvironmentIDParamsWithTimeout(timeout time.Duration) *PatchV1EnvironmentsEnvironmentIDParams {
	return &PatchV1EnvironmentsEnvironmentIDParams{
		timeout: timeout,
	}
}

// NewPatchV1EnvironmentsEnvironmentIDParamsWithContext creates a new PatchV1EnvironmentsEnvironmentIDParams object
// with the ability to set a context for a request.
func NewPatchV1EnvironmentsEnvironmentIDParamsWithContext(ctx context.Context) *PatchV1EnvironmentsEnvironmentIDParams {
	return &PatchV1EnvironmentsEnvironmentIDParams{
		Context: ctx,
	}
}

// NewPatchV1EnvironmentsEnvironmentIDParamsWithHTTPClient creates a new PatchV1EnvironmentsEnvironmentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1EnvironmentsEnvironmentIDParamsWithHTTPClient(client *http.Client) *PatchV1EnvironmentsEnvironmentIDParams {
	return &PatchV1EnvironmentsEnvironmentIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1EnvironmentsEnvironmentIDParams contains all the parameters to send to the API endpoint

	for the patch v1 environments environment Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1EnvironmentsEnvironmentIDParams struct {

	/* EnvironmentID.

	   Environment UUID
	*/
	EnvironmentID string

	// PatchV1EnvironmentsEnvironmentID.
	PatchV1EnvironmentsEnvironmentID *models.PatchV1EnvironmentsEnvironmentID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 environments environment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1EnvironmentsEnvironmentIDParams) WithDefaults() *PatchV1EnvironmentsEnvironmentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 environments environment Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1EnvironmentsEnvironmentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) WithTimeout(timeout time.Duration) *PatchV1EnvironmentsEnvironmentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) WithContext(ctx context.Context) *PatchV1EnvironmentsEnvironmentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) WithHTTPClient(client *http.Client) *PatchV1EnvironmentsEnvironmentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironmentID adds the environmentID to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) WithEnvironmentID(environmentID string) *PatchV1EnvironmentsEnvironmentIDParams {
	o.SetEnvironmentID(environmentID)
	return o
}

// SetEnvironmentID adds the environmentId to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) SetEnvironmentID(environmentID string) {
	o.EnvironmentID = environmentID
}

// WithPatchV1EnvironmentsEnvironmentID adds the patchV1EnvironmentsEnvironmentID to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) WithPatchV1EnvironmentsEnvironmentID(patchV1EnvironmentsEnvironmentID *models.PatchV1EnvironmentsEnvironmentID) *PatchV1EnvironmentsEnvironmentIDParams {
	o.SetPatchV1EnvironmentsEnvironmentID(patchV1EnvironmentsEnvironmentID)
	return o
}

// SetPatchV1EnvironmentsEnvironmentID adds the patchV1EnvironmentsEnvironmentId to the patch v1 environments environment Id params
func (o *PatchV1EnvironmentsEnvironmentIDParams) SetPatchV1EnvironmentsEnvironmentID(patchV1EnvironmentsEnvironmentID *models.PatchV1EnvironmentsEnvironmentID) {
	o.PatchV1EnvironmentsEnvironmentID = patchV1EnvironmentsEnvironmentID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1EnvironmentsEnvironmentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param environment_id
	if err := r.SetPathParam("environment_id", o.EnvironmentID); err != nil {
		return err
	}
	if o.PatchV1EnvironmentsEnvironmentID != nil {
		if err := r.SetBodyParam(o.PatchV1EnvironmentsEnvironmentID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
