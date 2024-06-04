// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

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

// NewPostV1NuncConnectionsNuncConnectionIDLinksParams creates a new PostV1NuncConnectionsNuncConnectionIDLinksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1NuncConnectionsNuncConnectionIDLinksParams() *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	return &PostV1NuncConnectionsNuncConnectionIDLinksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDLinksParamsWithTimeout creates a new PostV1NuncConnectionsNuncConnectionIDLinksParams object
// with the ability to set a timeout on a request.
func NewPostV1NuncConnectionsNuncConnectionIDLinksParamsWithTimeout(timeout time.Duration) *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	return &PostV1NuncConnectionsNuncConnectionIDLinksParams{
		timeout: timeout,
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDLinksParamsWithContext creates a new PostV1NuncConnectionsNuncConnectionIDLinksParams object
// with the ability to set a context for a request.
func NewPostV1NuncConnectionsNuncConnectionIDLinksParamsWithContext(ctx context.Context) *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	return &PostV1NuncConnectionsNuncConnectionIDLinksParams{
		Context: ctx,
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDLinksParamsWithHTTPClient creates a new PostV1NuncConnectionsNuncConnectionIDLinksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1NuncConnectionsNuncConnectionIDLinksParamsWithHTTPClient(client *http.Client) *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	return &PostV1NuncConnectionsNuncConnectionIDLinksParams{
		HTTPClient: client,
	}
}

/*
PostV1NuncConnectionsNuncConnectionIDLinksParams contains all the parameters to send to the API endpoint

	for the post v1 nunc connections nunc connection Id links operation.

	Typically these are written to a http.Request.
*/
type PostV1NuncConnectionsNuncConnectionIDLinksParams struct {

	// NuncConnectionID.
	NuncConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 nunc connections nunc connection Id links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) WithDefaults() *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 nunc connections nunc connection Id links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) WithTimeout(timeout time.Duration) *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) WithContext(ctx context.Context) *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) WithHTTPClient(client *http.Client) *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNuncConnectionID adds the nuncConnectionID to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) WithNuncConnectionID(nuncConnectionID string) *PostV1NuncConnectionsNuncConnectionIDLinksParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the post v1 nunc connections nunc connection Id links params
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) SetNuncConnectionID(nuncConnectionID string) {
	o.NuncConnectionID = nuncConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1NuncConnectionsNuncConnectionIDLinksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", o.NuncConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
