// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

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

// NewPatchV1ScimV2UsersIDParams creates a new PatchV1ScimV2UsersIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ScimV2UsersIDParams() *PatchV1ScimV2UsersIDParams {
	return &PatchV1ScimV2UsersIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ScimV2UsersIDParamsWithTimeout creates a new PatchV1ScimV2UsersIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1ScimV2UsersIDParamsWithTimeout(timeout time.Duration) *PatchV1ScimV2UsersIDParams {
	return &PatchV1ScimV2UsersIDParams{
		timeout: timeout,
	}
}

// NewPatchV1ScimV2UsersIDParamsWithContext creates a new PatchV1ScimV2UsersIDParams object
// with the ability to set a context for a request.
func NewPatchV1ScimV2UsersIDParamsWithContext(ctx context.Context) *PatchV1ScimV2UsersIDParams {
	return &PatchV1ScimV2UsersIDParams{
		Context: ctx,
	}
}

// NewPatchV1ScimV2UsersIDParamsWithHTTPClient creates a new PatchV1ScimV2UsersIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ScimV2UsersIDParamsWithHTTPClient(client *http.Client) *PatchV1ScimV2UsersIDParams {
	return &PatchV1ScimV2UsersIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1ScimV2UsersIDParams contains all the parameters to send to the API endpoint

	for the patch v1 scim v2 users Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1ScimV2UsersIDParams struct {

	// ID.
	ID string

	// PatchV1ScimV2UsersID.
	PatchV1ScimV2UsersID *models.PatchV1ScimV2UsersID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 scim v2 users Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScimV2UsersIDParams) WithDefaults() *PatchV1ScimV2UsersIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 scim v2 users Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ScimV2UsersIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) WithTimeout(timeout time.Duration) *PatchV1ScimV2UsersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) WithContext(ctx context.Context) *PatchV1ScimV2UsersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) WithHTTPClient(client *http.Client) *PatchV1ScimV2UsersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) WithID(id string) *PatchV1ScimV2UsersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) SetID(id string) {
	o.ID = id
}

// WithPatchV1ScimV2UsersID adds the patchV1ScimV2UsersID to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) WithPatchV1ScimV2UsersID(patchV1ScimV2UsersID *models.PatchV1ScimV2UsersID) *PatchV1ScimV2UsersIDParams {
	o.SetPatchV1ScimV2UsersID(patchV1ScimV2UsersID)
	return o
}

// SetPatchV1ScimV2UsersID adds the patchV1ScimV2UsersId to the patch v1 scim v2 users Id params
func (o *PatchV1ScimV2UsersIDParams) SetPatchV1ScimV2UsersID(patchV1ScimV2UsersID *models.PatchV1ScimV2UsersID) {
	o.PatchV1ScimV2UsersID = patchV1ScimV2UsersID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ScimV2UsersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.PatchV1ScimV2UsersID != nil {
		if err := r.SetBodyParam(o.PatchV1ScimV2UsersID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
