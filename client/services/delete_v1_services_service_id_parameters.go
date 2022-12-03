// Code generated by go-swagger; DO NOT EDIT.

package services

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
)

// NewDeleteV1ServicesServiceIDParams creates a new DeleteV1ServicesServiceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteV1ServicesServiceIDParams() *DeleteV1ServicesServiceIDParams {
	return &DeleteV1ServicesServiceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteV1ServicesServiceIDParamsWithTimeout creates a new DeleteV1ServicesServiceIDParams object
// with the ability to set a timeout on a request.
func NewDeleteV1ServicesServiceIDParamsWithTimeout(timeout time.Duration) *DeleteV1ServicesServiceIDParams {
	return &DeleteV1ServicesServiceIDParams{
		timeout: timeout,
	}
}

// NewDeleteV1ServicesServiceIDParamsWithContext creates a new DeleteV1ServicesServiceIDParams object
// with the ability to set a context for a request.
func NewDeleteV1ServicesServiceIDParamsWithContext(ctx context.Context) *DeleteV1ServicesServiceIDParams {
	return &DeleteV1ServicesServiceIDParams{
		Context: ctx,
	}
}

// NewDeleteV1ServicesServiceIDParamsWithHTTPClient creates a new DeleteV1ServicesServiceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteV1ServicesServiceIDParamsWithHTTPClient(client *http.Client) *DeleteV1ServicesServiceIDParams {
	return &DeleteV1ServicesServiceIDParams{
		HTTPClient: client,
	}
}

/*
DeleteV1ServicesServiceIDParams contains all the parameters to send to the API endpoint

	for the delete v1 services service Id operation.

	Typically these are written to a http.Request.
*/
type DeleteV1ServicesServiceIDParams struct {

	// ServiceID.
	//
	// Format: int32
	ServiceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete v1 services service Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ServicesServiceIDParams) WithDefaults() *DeleteV1ServicesServiceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete v1 services service Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteV1ServicesServiceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) WithTimeout(timeout time.Duration) *DeleteV1ServicesServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) WithContext(ctx context.Context) *DeleteV1ServicesServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) WithHTTPClient(client *http.Client) *DeleteV1ServicesServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) WithServiceID(serviceID int32) *DeleteV1ServicesServiceIDParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the delete v1 services service Id params
func (o *DeleteV1ServicesServiceIDParams) SetServiceID(serviceID int32) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteV1ServicesServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param service_id
	if err := r.SetPathParam("service_id", swag.FormatInt32(o.ServiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
