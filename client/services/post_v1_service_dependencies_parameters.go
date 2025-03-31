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

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1ServiceDependenciesParams creates a new PostV1ServiceDependenciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ServiceDependenciesParams() *PostV1ServiceDependenciesParams {
	return &PostV1ServiceDependenciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ServiceDependenciesParamsWithTimeout creates a new PostV1ServiceDependenciesParams object
// with the ability to set a timeout on a request.
func NewPostV1ServiceDependenciesParamsWithTimeout(timeout time.Duration) *PostV1ServiceDependenciesParams {
	return &PostV1ServiceDependenciesParams{
		timeout: timeout,
	}
}

// NewPostV1ServiceDependenciesParamsWithContext creates a new PostV1ServiceDependenciesParams object
// with the ability to set a context for a request.
func NewPostV1ServiceDependenciesParamsWithContext(ctx context.Context) *PostV1ServiceDependenciesParams {
	return &PostV1ServiceDependenciesParams{
		Context: ctx,
	}
}

// NewPostV1ServiceDependenciesParamsWithHTTPClient creates a new PostV1ServiceDependenciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ServiceDependenciesParamsWithHTTPClient(client *http.Client) *PostV1ServiceDependenciesParams {
	return &PostV1ServiceDependenciesParams{
		HTTPClient: client,
	}
}

/*
PostV1ServiceDependenciesParams contains all the parameters to send to the API endpoint

	for the post v1 service dependencies operation.

	Typically these are written to a http.Request.
*/
type PostV1ServiceDependenciesParams struct {

	// PostV1ServiceDependencies.
	PostV1ServiceDependencies *models.PostV1ServiceDependencies

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 service dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ServiceDependenciesParams) WithDefaults() *PostV1ServiceDependenciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 service dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ServiceDependenciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) WithTimeout(timeout time.Duration) *PostV1ServiceDependenciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) WithContext(ctx context.Context) *PostV1ServiceDependenciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) WithHTTPClient(client *http.Client) *PostV1ServiceDependenciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostV1ServiceDependencies adds the postV1ServiceDependencies to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) WithPostV1ServiceDependencies(postV1ServiceDependencies *models.PostV1ServiceDependencies) *PostV1ServiceDependenciesParams {
	o.SetPostV1ServiceDependencies(postV1ServiceDependencies)
	return o
}

// SetPostV1ServiceDependencies adds the postV1ServiceDependencies to the post v1 service dependencies params
func (o *PostV1ServiceDependenciesParams) SetPostV1ServiceDependencies(postV1ServiceDependencies *models.PostV1ServiceDependencies) {
	o.PostV1ServiceDependencies = postV1ServiceDependencies
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ServiceDependenciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostV1ServiceDependencies != nil {
		if err := r.SetBodyParam(o.PostV1ServiceDependencies); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
