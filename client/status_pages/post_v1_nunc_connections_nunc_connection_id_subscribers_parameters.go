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

// NewPostV1NuncConnectionsNuncConnectionIDSubscribersParams creates a new PostV1NuncConnectionsNuncConnectionIDSubscribersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1NuncConnectionsNuncConnectionIDSubscribersParams() *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	return &PostV1NuncConnectionsNuncConnectionIDSubscribersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDSubscribersParamsWithTimeout creates a new PostV1NuncConnectionsNuncConnectionIDSubscribersParams object
// with the ability to set a timeout on a request.
func NewPostV1NuncConnectionsNuncConnectionIDSubscribersParamsWithTimeout(timeout time.Duration) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	return &PostV1NuncConnectionsNuncConnectionIDSubscribersParams{
		timeout: timeout,
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDSubscribersParamsWithContext creates a new PostV1NuncConnectionsNuncConnectionIDSubscribersParams object
// with the ability to set a context for a request.
func NewPostV1NuncConnectionsNuncConnectionIDSubscribersParamsWithContext(ctx context.Context) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	return &PostV1NuncConnectionsNuncConnectionIDSubscribersParams{
		Context: ctx,
	}
}

// NewPostV1NuncConnectionsNuncConnectionIDSubscribersParamsWithHTTPClient creates a new PostV1NuncConnectionsNuncConnectionIDSubscribersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1NuncConnectionsNuncConnectionIDSubscribersParamsWithHTTPClient(client *http.Client) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	return &PostV1NuncConnectionsNuncConnectionIDSubscribersParams{
		HTTPClient: client,
	}
}

/*
PostV1NuncConnectionsNuncConnectionIDSubscribersParams contains all the parameters to send to the API endpoint

	for the post v1 nunc connections nunc connection Id subscribers operation.

	Typically these are written to a http.Request.
*/
type PostV1NuncConnectionsNuncConnectionIDSubscribersParams struct {

	/* Emails.

	   A comma-separated list of emails to subscribe.
	*/
	Emails string

	// NuncConnectionID.
	NuncConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 nunc connections nunc connection Id subscribers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) WithDefaults() *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 nunc connections nunc connection Id subscribers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) WithTimeout(timeout time.Duration) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) WithContext(ctx context.Context) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) WithHTTPClient(client *http.Client) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmails adds the emails to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) WithEmails(emails string) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	o.SetEmails(emails)
	return o
}

// SetEmails adds the emails to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) SetEmails(emails string) {
	o.Emails = emails
}

// WithNuncConnectionID adds the nuncConnectionID to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) WithNuncConnectionID(nuncConnectionID string) *PostV1NuncConnectionsNuncConnectionIDSubscribersParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the post v1 nunc connections nunc connection Id subscribers params
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) SetNuncConnectionID(nuncConnectionID string) {
	o.NuncConnectionID = nuncConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1NuncConnectionsNuncConnectionIDSubscribersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param emails
	frEmails := o.Emails
	fEmails := frEmails
	if fEmails != "" {
		if err := r.SetFormParam("emails", fEmails); err != nil {
			return err
		}
	}

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", o.NuncConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
