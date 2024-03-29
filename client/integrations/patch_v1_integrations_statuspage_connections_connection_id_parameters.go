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

// NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParams creates a new PatchV1IntegrationsStatuspageConnectionsConnectionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParams() *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsConnectionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParamsWithTimeout creates a new PatchV1IntegrationsStatuspageConnectionsConnectionIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParamsWithTimeout(timeout time.Duration) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsConnectionIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParamsWithContext creates a new PatchV1IntegrationsStatuspageConnectionsConnectionIDParams object
// with the ability to set a context for a request.
func NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParamsWithContext(ctx context.Context) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsConnectionIDParams{
		Context: ctx,
	}
}

// NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParamsWithHTTPClient creates a new PatchV1IntegrationsStatuspageConnectionsConnectionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IntegrationsStatuspageConnectionsConnectionIDParamsWithHTTPClient(client *http.Client) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	return &PatchV1IntegrationsStatuspageConnectionsConnectionIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1IntegrationsStatuspageConnectionsConnectionIDParams contains all the parameters to send to the API endpoint

	for the patch v1 integrations statuspage connections connection Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1IntegrationsStatuspageConnectionsConnectionIDParams struct {

	/* ConnectionID.

	   Connection UUID
	*/
	ConnectionID string

	// PatchV1IntegrationsStatuspageConnectionsConnectionID.
	PatchV1IntegrationsStatuspageConnectionsConnectionID *models.PatchV1IntegrationsStatuspageConnectionsConnectionID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 integrations statuspage connections connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) WithDefaults() *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 integrations statuspage connections connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) WithTimeout(timeout time.Duration) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) WithContext(ctx context.Context) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) WithHTTPClient(client *http.Client) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) WithConnectionID(connectionID string) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) SetConnectionID(connectionID string) {
	o.ConnectionID = connectionID
}

// WithPatchV1IntegrationsStatuspageConnectionsConnectionID adds the patchV1IntegrationsStatuspageConnectionsConnectionID to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) WithPatchV1IntegrationsStatuspageConnectionsConnectionID(patchV1IntegrationsStatuspageConnectionsConnectionID *models.PatchV1IntegrationsStatuspageConnectionsConnectionID) *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams {
	o.SetPatchV1IntegrationsStatuspageConnectionsConnectionID(patchV1IntegrationsStatuspageConnectionsConnectionID)
	return o
}

// SetPatchV1IntegrationsStatuspageConnectionsConnectionID adds the patchV1IntegrationsStatuspageConnectionsConnectionId to the patch v1 integrations statuspage connections connection Id params
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) SetPatchV1IntegrationsStatuspageConnectionsConnectionID(patchV1IntegrationsStatuspageConnectionsConnectionID *models.PatchV1IntegrationsStatuspageConnectionsConnectionID) {
	o.PatchV1IntegrationsStatuspageConnectionsConnectionID = patchV1IntegrationsStatuspageConnectionsConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IntegrationsStatuspageConnectionsConnectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", o.ConnectionID); err != nil {
		return err
	}
	if o.PatchV1IntegrationsStatuspageConnectionsConnectionID != nil {
		if err := r.SetBodyParam(o.PatchV1IntegrationsStatuspageConnectionsConnectionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
