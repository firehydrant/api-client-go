// Code generated by go-swagger; DO NOT EDIT.

package incident_settings

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

// NewPostV1SeverityMatrixConditionsParams creates a new PostV1SeverityMatrixConditionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1SeverityMatrixConditionsParams() *PostV1SeverityMatrixConditionsParams {
	return &PostV1SeverityMatrixConditionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1SeverityMatrixConditionsParamsWithTimeout creates a new PostV1SeverityMatrixConditionsParams object
// with the ability to set a timeout on a request.
func NewPostV1SeverityMatrixConditionsParamsWithTimeout(timeout time.Duration) *PostV1SeverityMatrixConditionsParams {
	return &PostV1SeverityMatrixConditionsParams{
		timeout: timeout,
	}
}

// NewPostV1SeverityMatrixConditionsParamsWithContext creates a new PostV1SeverityMatrixConditionsParams object
// with the ability to set a context for a request.
func NewPostV1SeverityMatrixConditionsParamsWithContext(ctx context.Context) *PostV1SeverityMatrixConditionsParams {
	return &PostV1SeverityMatrixConditionsParams{
		Context: ctx,
	}
}

// NewPostV1SeverityMatrixConditionsParamsWithHTTPClient creates a new PostV1SeverityMatrixConditionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1SeverityMatrixConditionsParamsWithHTTPClient(client *http.Client) *PostV1SeverityMatrixConditionsParams {
	return &PostV1SeverityMatrixConditionsParams{
		HTTPClient: client,
	}
}

/*
PostV1SeverityMatrixConditionsParams contains all the parameters to send to the API endpoint

	for the post v1 severity matrix conditions operation.

	Typically these are written to a http.Request.
*/
type PostV1SeverityMatrixConditionsParams struct {

	// PostV1SeverityMatrixConditions.
	PostV1SeverityMatrixConditions *models.PostV1SeverityMatrixConditions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 severity matrix conditions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SeverityMatrixConditionsParams) WithDefaults() *PostV1SeverityMatrixConditionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 severity matrix conditions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SeverityMatrixConditionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) WithTimeout(timeout time.Duration) *PostV1SeverityMatrixConditionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) WithContext(ctx context.Context) *PostV1SeverityMatrixConditionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) WithHTTPClient(client *http.Client) *PostV1SeverityMatrixConditionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPostV1SeverityMatrixConditions adds the postV1SeverityMatrixConditions to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) WithPostV1SeverityMatrixConditions(postV1SeverityMatrixConditions *models.PostV1SeverityMatrixConditions) *PostV1SeverityMatrixConditionsParams {
	o.SetPostV1SeverityMatrixConditions(postV1SeverityMatrixConditions)
	return o
}

// SetPostV1SeverityMatrixConditions adds the postV1SeverityMatrixConditions to the post v1 severity matrix conditions params
func (o *PostV1SeverityMatrixConditionsParams) SetPostV1SeverityMatrixConditions(postV1SeverityMatrixConditions *models.PostV1SeverityMatrixConditions) {
	o.PostV1SeverityMatrixConditions = postV1SeverityMatrixConditions
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1SeverityMatrixConditionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PostV1SeverityMatrixConditions != nil {
		if err := r.SetBodyParam(o.PostV1SeverityMatrixConditions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
