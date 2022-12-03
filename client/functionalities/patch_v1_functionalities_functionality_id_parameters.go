// Code generated by go-swagger; DO NOT EDIT.

package functionalities

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

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1FunctionalitiesFunctionalityIDParams creates a new PatchV1FunctionalitiesFunctionalityIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1FunctionalitiesFunctionalityIDParams() *PatchV1FunctionalitiesFunctionalityIDParams {
	return &PatchV1FunctionalitiesFunctionalityIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1FunctionalitiesFunctionalityIDParamsWithTimeout creates a new PatchV1FunctionalitiesFunctionalityIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1FunctionalitiesFunctionalityIDParamsWithTimeout(timeout time.Duration) *PatchV1FunctionalitiesFunctionalityIDParams {
	return &PatchV1FunctionalitiesFunctionalityIDParams{
		timeout: timeout,
	}
}

// NewPatchV1FunctionalitiesFunctionalityIDParamsWithContext creates a new PatchV1FunctionalitiesFunctionalityIDParams object
// with the ability to set a context for a request.
func NewPatchV1FunctionalitiesFunctionalityIDParamsWithContext(ctx context.Context) *PatchV1FunctionalitiesFunctionalityIDParams {
	return &PatchV1FunctionalitiesFunctionalityIDParams{
		Context: ctx,
	}
}

// NewPatchV1FunctionalitiesFunctionalityIDParamsWithHTTPClient creates a new PatchV1FunctionalitiesFunctionalityIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1FunctionalitiesFunctionalityIDParamsWithHTTPClient(client *http.Client) *PatchV1FunctionalitiesFunctionalityIDParams {
	return &PatchV1FunctionalitiesFunctionalityIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1FunctionalitiesFunctionalityIDParams contains all the parameters to send to the API endpoint

	for the patch v1 functionalities functionality Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1FunctionalitiesFunctionalityIDParams struct {

	// V1Functionalities.
	V1Functionalities *models.PatchV1Functionalities

	// FunctionalityID.
	//
	// Format: int32
	FunctionalityID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 functionalities functionality Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1FunctionalitiesFunctionalityIDParams) WithDefaults() *PatchV1FunctionalitiesFunctionalityIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 functionalities functionality Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1FunctionalitiesFunctionalityIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) WithTimeout(timeout time.Duration) *PatchV1FunctionalitiesFunctionalityIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) WithContext(ctx context.Context) *PatchV1FunctionalitiesFunctionalityIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) WithHTTPClient(client *http.Client) *PatchV1FunctionalitiesFunctionalityIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Functionalities adds the v1Functionalities to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) WithV1Functionalities(v1Functionalities *models.PatchV1Functionalities) *PatchV1FunctionalitiesFunctionalityIDParams {
	o.SetV1Functionalities(v1Functionalities)
	return o
}

// SetV1Functionalities adds the v1Functionalities to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) SetV1Functionalities(v1Functionalities *models.PatchV1Functionalities) {
	o.V1Functionalities = v1Functionalities
}

// WithFunctionalityID adds the functionalityID to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) WithFunctionalityID(functionalityID int32) *PatchV1FunctionalitiesFunctionalityIDParams {
	o.SetFunctionalityID(functionalityID)
	return o
}

// SetFunctionalityID adds the functionalityId to the patch v1 functionalities functionality Id params
func (o *PatchV1FunctionalitiesFunctionalityIDParams) SetFunctionalityID(functionalityID int32) {
	o.FunctionalityID = functionalityID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1FunctionalitiesFunctionalityIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Functionalities != nil {
		if err := r.SetBodyParam(o.V1Functionalities); err != nil {
			return err
		}
	}

	// path param functionality_id
	if err := r.SetPathParam("functionality_id", swag.FormatInt32(o.FunctionalityID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
