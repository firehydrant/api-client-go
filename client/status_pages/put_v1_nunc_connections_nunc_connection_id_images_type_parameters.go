// Code generated by go-swagger; DO NOT EDIT.

package status_pages

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

// NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParams creates a new PutV1NuncConnectionsNuncConnectionIDImagesTypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParams() *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	return &PutV1NuncConnectionsNuncConnectionIDImagesTypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParamsWithTimeout creates a new PutV1NuncConnectionsNuncConnectionIDImagesTypeParams object
// with the ability to set a timeout on a request.
func NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParamsWithTimeout(timeout time.Duration) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	return &PutV1NuncConnectionsNuncConnectionIDImagesTypeParams{
		timeout: timeout,
	}
}

// NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParamsWithContext creates a new PutV1NuncConnectionsNuncConnectionIDImagesTypeParams object
// with the ability to set a context for a request.
func NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParamsWithContext(ctx context.Context) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	return &PutV1NuncConnectionsNuncConnectionIDImagesTypeParams{
		Context: ctx,
	}
}

// NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParamsWithHTTPClient creates a new PutV1NuncConnectionsNuncConnectionIDImagesTypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1NuncConnectionsNuncConnectionIDImagesTypeParamsWithHTTPClient(client *http.Client) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	return &PutV1NuncConnectionsNuncConnectionIDImagesTypeParams{
		HTTPClient: client,
	}
}

/*
PutV1NuncConnectionsNuncConnectionIDImagesTypeParams contains all the parameters to send to the API endpoint

	for the put v1 nunc connections nunc connection Id images type operation.

	Typically these are written to a http.Request.
*/
type PutV1NuncConnectionsNuncConnectionIDImagesTypeParams struct {

	// File.
	File runtime.NamedReadCloser

	// NuncConnectionID.
	NuncConnectionID string

	// Type.
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 nunc connections nunc connection Id images type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WithDefaults() *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 nunc connections nunc connection Id images type params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WithTimeout(timeout time.Duration) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WithContext(ctx context.Context) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WithHTTPClient(client *http.Client) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WithFile(file runtime.NamedReadCloser) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithNuncConnectionID adds the nuncConnectionID to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WithNuncConnectionID(nuncConnectionID string) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) SetNuncConnectionID(nuncConnectionID string) {
	o.NuncConnectionID = nuncConnectionID
}

// WithType adds the typeVar to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WithType(typeVar string) *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the put v1 nunc connections nunc connection Id images type params
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1NuncConnectionsNuncConnectionIDImagesTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.File != nil {

		if o.File != nil {
			// form file param file
			if err := r.SetFileParam("file", o.File); err != nil {
				return err
			}
		}
	}

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", o.NuncConnectionID); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
