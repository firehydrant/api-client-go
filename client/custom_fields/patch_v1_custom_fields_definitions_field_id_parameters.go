// Code generated by go-swagger; DO NOT EDIT.

package custom_fields

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

// NewPatchV1CustomFieldsDefinitionsFieldIDParams creates a new PatchV1CustomFieldsDefinitionsFieldIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1CustomFieldsDefinitionsFieldIDParams() *PatchV1CustomFieldsDefinitionsFieldIDParams {
	return &PatchV1CustomFieldsDefinitionsFieldIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1CustomFieldsDefinitionsFieldIDParamsWithTimeout creates a new PatchV1CustomFieldsDefinitionsFieldIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1CustomFieldsDefinitionsFieldIDParamsWithTimeout(timeout time.Duration) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	return &PatchV1CustomFieldsDefinitionsFieldIDParams{
		timeout: timeout,
	}
}

// NewPatchV1CustomFieldsDefinitionsFieldIDParamsWithContext creates a new PatchV1CustomFieldsDefinitionsFieldIDParams object
// with the ability to set a context for a request.
func NewPatchV1CustomFieldsDefinitionsFieldIDParamsWithContext(ctx context.Context) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	return &PatchV1CustomFieldsDefinitionsFieldIDParams{
		Context: ctx,
	}
}

// NewPatchV1CustomFieldsDefinitionsFieldIDParamsWithHTTPClient creates a new PatchV1CustomFieldsDefinitionsFieldIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1CustomFieldsDefinitionsFieldIDParamsWithHTTPClient(client *http.Client) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	return &PatchV1CustomFieldsDefinitionsFieldIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1CustomFieldsDefinitionsFieldIDParams contains all the parameters to send to the API endpoint

	for the patch v1 custom fields definitions field Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1CustomFieldsDefinitionsFieldIDParams struct {

	// FieldID.
	FieldID string

	// PatchV1CustomFieldsDefinitionsFieldID.
	PatchV1CustomFieldsDefinitionsFieldID *models.PatchV1CustomFieldsDefinitionsFieldID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 custom fields definitions field Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) WithDefaults() *PatchV1CustomFieldsDefinitionsFieldIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 custom fields definitions field Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) WithTimeout(timeout time.Duration) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) WithContext(ctx context.Context) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) WithHTTPClient(client *http.Client) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldID adds the fieldID to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) WithFieldID(fieldID string) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	o.SetFieldID(fieldID)
	return o
}

// SetFieldID adds the fieldId to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) SetFieldID(fieldID string) {
	o.FieldID = fieldID
}

// WithPatchV1CustomFieldsDefinitionsFieldID adds the patchV1CustomFieldsDefinitionsFieldID to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) WithPatchV1CustomFieldsDefinitionsFieldID(patchV1CustomFieldsDefinitionsFieldID *models.PatchV1CustomFieldsDefinitionsFieldID) *PatchV1CustomFieldsDefinitionsFieldIDParams {
	o.SetPatchV1CustomFieldsDefinitionsFieldID(patchV1CustomFieldsDefinitionsFieldID)
	return o
}

// SetPatchV1CustomFieldsDefinitionsFieldID adds the patchV1CustomFieldsDefinitionsFieldId to the patch v1 custom fields definitions field Id params
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) SetPatchV1CustomFieldsDefinitionsFieldID(patchV1CustomFieldsDefinitionsFieldID *models.PatchV1CustomFieldsDefinitionsFieldID) {
	o.PatchV1CustomFieldsDefinitionsFieldID = patchV1CustomFieldsDefinitionsFieldID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1CustomFieldsDefinitionsFieldIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field_id
	if err := r.SetPathParam("field_id", o.FieldID); err != nil {
		return err
	}
	if o.PatchV1CustomFieldsDefinitionsFieldID != nil {
		if err := r.SetBodyParam(o.PatchV1CustomFieldsDefinitionsFieldID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}