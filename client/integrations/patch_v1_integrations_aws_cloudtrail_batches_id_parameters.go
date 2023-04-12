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

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1IntegrationsAwsCloudtrailBatchesIDParams creates a new PatchV1IntegrationsAwsCloudtrailBatchesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IntegrationsAwsCloudtrailBatchesIDParams() *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	return &PatchV1IntegrationsAwsCloudtrailBatchesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IntegrationsAwsCloudtrailBatchesIDParamsWithTimeout creates a new PatchV1IntegrationsAwsCloudtrailBatchesIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IntegrationsAwsCloudtrailBatchesIDParamsWithTimeout(timeout time.Duration) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	return &PatchV1IntegrationsAwsCloudtrailBatchesIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IntegrationsAwsCloudtrailBatchesIDParamsWithContext creates a new PatchV1IntegrationsAwsCloudtrailBatchesIDParams object
// with the ability to set a context for a request.
func NewPatchV1IntegrationsAwsCloudtrailBatchesIDParamsWithContext(ctx context.Context) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	return &PatchV1IntegrationsAwsCloudtrailBatchesIDParams{
		Context: ctx,
	}
}

// NewPatchV1IntegrationsAwsCloudtrailBatchesIDParamsWithHTTPClient creates a new PatchV1IntegrationsAwsCloudtrailBatchesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IntegrationsAwsCloudtrailBatchesIDParamsWithHTTPClient(client *http.Client) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	return &PatchV1IntegrationsAwsCloudtrailBatchesIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1IntegrationsAwsCloudtrailBatchesIDParams contains all the parameters to send to the API endpoint

	for the patch v1 integrations aws cloudtrail batches Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1IntegrationsAwsCloudtrailBatchesIDParams struct {

	/* ID.

	   Connection UUID
	*/
	ID string

	// PatchV1IntegrationsAwsCloudtrailBatchesID.
	PatchV1IntegrationsAwsCloudtrailBatchesID *models.PatchV1IntegrationsAwsCloudtrailBatchesID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 integrations aws cloudtrail batches Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) WithDefaults() *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 integrations aws cloudtrail batches Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) WithTimeout(timeout time.Duration) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) WithContext(ctx context.Context) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) WithHTTPClient(client *http.Client) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) WithID(id string) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) SetID(id string) {
	o.ID = id
}

// WithPatchV1IntegrationsAwsCloudtrailBatchesID adds the patchV1IntegrationsAwsCloudtrailBatchesID to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) WithPatchV1IntegrationsAwsCloudtrailBatchesID(patchV1IntegrationsAwsCloudtrailBatchesID *models.PatchV1IntegrationsAwsCloudtrailBatchesID) *PatchV1IntegrationsAwsCloudtrailBatchesIDParams {
	o.SetPatchV1IntegrationsAwsCloudtrailBatchesID(patchV1IntegrationsAwsCloudtrailBatchesID)
	return o
}

// SetPatchV1IntegrationsAwsCloudtrailBatchesID adds the patchV1IntegrationsAwsCloudtrailBatchesId to the patch v1 integrations aws cloudtrail batches Id params
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) SetPatchV1IntegrationsAwsCloudtrailBatchesID(patchV1IntegrationsAwsCloudtrailBatchesID *models.PatchV1IntegrationsAwsCloudtrailBatchesID) {
	o.PatchV1IntegrationsAwsCloudtrailBatchesID = patchV1IntegrationsAwsCloudtrailBatchesID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IntegrationsAwsCloudtrailBatchesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.PatchV1IntegrationsAwsCloudtrailBatchesID != nil {
		if err := r.SetBodyParam(o.PatchV1IntegrationsAwsCloudtrailBatchesID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
