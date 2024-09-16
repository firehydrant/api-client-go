// Code generated by go-swagger; DO NOT EDIT.

package lifecycles

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

// NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams creates a new DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams() *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	return &DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParamsWithTimeout creates a new DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParamsWithTimeout(timeout time.Duration) *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	return &DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParamsWithContext creates a new DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams object
// with the ability to set a context for a request.
func NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParamsWithContext(ctx context.Context) *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	return &DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams{
		Context: ctx,
	}
}

// NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParamsWithHTTPClient creates a new DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParamsWithHTTPClient(client *http.Client) *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	return &DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams contains all the parameters to send to the API endpoint

	for the delete v1 lifecycles measurement definitions measurement definition Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams struct {

	// MeasurementDefinitionID.
	MeasurementDefinitionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 lifecycles measurement definitions measurement definition Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) WithDefaults() *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 lifecycles measurement definitions measurement definition Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) WithTimeout(timeout time.Duration) *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) WithContext(ctx context.Context) *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) WithHTTPClient(client *http.Client) *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMeasurementDefinitionID adds the measurementDefinitionID to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) WithMeasurementDefinitionID(measurementDefinitionID string) *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams {
	o.SetMeasurementDefinitionID(measurementDefinitionID)
	return o
}

// SetMeasurementDefinitionID adds the measurementDefinitionId to the delete v1 lifecycles measurement definitions measurement definition Id params
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) SetMeasurementDefinitionID(measurementDefinitionID string) {
	o.MeasurementDefinitionID = measurementDefinitionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param measurement_definition_id
	if err := r.SetPathParam("measurement_definition_id", o.MeasurementDefinitionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
