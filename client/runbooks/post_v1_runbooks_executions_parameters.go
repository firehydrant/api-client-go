// Code generated by go-swagger; DO NOT EDIT.

package runbooks

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

// NewPostV1RunbooksExecutionsParams creates a new PostV1RunbooksExecutionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1RunbooksExecutionsParams() *PostV1RunbooksExecutionsParams {
	return &PostV1RunbooksExecutionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1RunbooksExecutionsParamsWithTimeout creates a new PostV1RunbooksExecutionsParams object
// with the ability to set a timeout on a request.
func NewPostV1RunbooksExecutionsParamsWithTimeout(timeout time.Duration) *PostV1RunbooksExecutionsParams {
	return &PostV1RunbooksExecutionsParams{
		timeout: timeout,
	}
}

// NewPostV1RunbooksExecutionsParamsWithContext creates a new PostV1RunbooksExecutionsParams object
// with the ability to set a context for a request.
func NewPostV1RunbooksExecutionsParamsWithContext(ctx context.Context) *PostV1RunbooksExecutionsParams {
	return &PostV1RunbooksExecutionsParams{
		Context: ctx,
	}
}

// NewPostV1RunbooksExecutionsParamsWithHTTPClient creates a new PostV1RunbooksExecutionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1RunbooksExecutionsParamsWithHTTPClient(client *http.Client) *PostV1RunbooksExecutionsParams {
	return &PostV1RunbooksExecutionsParams{
		HTTPClient: client,
	}
}

/*
PostV1RunbooksExecutionsParams contains all the parameters to send to the API endpoint

	for the post v1 runbooks executions operation.

	Typically these are written to a http.Request.
*/
type PostV1RunbooksExecutionsParams struct {

	// PostV1RunbooksExecutions.
	PostV1RunbooksExecutions *models.PostV1RunbooksExecutions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 runbooks executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1RunbooksExecutionsParams) WithDefaults() *PostV1RunbooksExecutionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 runbooks executions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1RunbooksExecutionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) WithTimeout(timeout time.Duration) *PostV1RunbooksExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) WithContext(ctx context.Context) *PostV1RunbooksExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) WithHTTPClient(client *http.Client) *PostV1RunbooksExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostV1RunbooksExecutions adds the postV1RunbooksExecutions to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) WithPostV1RunbooksExecutions(postV1RunbooksExecutions *models.PostV1RunbooksExecutions) *PostV1RunbooksExecutionsParams {
	o.SetPostV1RunbooksExecutions(postV1RunbooksExecutions)
	return o
}

// SetPostV1RunbooksExecutions adds the postV1RunbooksExecutions to the post v1 runbooks executions params
func (o *PostV1RunbooksExecutionsParams) SetPostV1RunbooksExecutions(postV1RunbooksExecutions *models.PostV1RunbooksExecutions) {
	o.PostV1RunbooksExecutions = postV1RunbooksExecutions
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1RunbooksExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostV1RunbooksExecutions != nil {
		if err := r.SetBodyParam(o.PostV1RunbooksExecutions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
