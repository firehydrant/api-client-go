// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

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

// NewPutV1PostMortemsQuestionsParams creates a new PutV1PostMortemsQuestionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1PostMortemsQuestionsParams() *PutV1PostMortemsQuestionsParams {
	return &PutV1PostMortemsQuestionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1PostMortemsQuestionsParamsWithTimeout creates a new PutV1PostMortemsQuestionsParams object
// with the ability to set a timeout on a request.
func NewPutV1PostMortemsQuestionsParamsWithTimeout(timeout time.Duration) *PutV1PostMortemsQuestionsParams {
	return &PutV1PostMortemsQuestionsParams{
		timeout: timeout,
	}
}

// NewPutV1PostMortemsQuestionsParamsWithContext creates a new PutV1PostMortemsQuestionsParams object
// with the ability to set a context for a request.
func NewPutV1PostMortemsQuestionsParamsWithContext(ctx context.Context) *PutV1PostMortemsQuestionsParams {
	return &PutV1PostMortemsQuestionsParams{
		Context: ctx,
	}
}

// NewPutV1PostMortemsQuestionsParamsWithHTTPClient creates a new PutV1PostMortemsQuestionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1PostMortemsQuestionsParamsWithHTTPClient(client *http.Client) *PutV1PostMortemsQuestionsParams {
	return &PutV1PostMortemsQuestionsParams{
		HTTPClient: client,
	}
}

/*
PutV1PostMortemsQuestionsParams contains all the parameters to send to the API endpoint

	for the put v1 post mortems questions operation.

	Typically these are written to a http.Request.
*/
type PutV1PostMortemsQuestionsParams struct {

	// PutV1PostMortemsQuestions.
	PutV1PostMortemsQuestions *models.PutV1PostMortemsQuestions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 post mortems questions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1PostMortemsQuestionsParams) WithDefaults() *PutV1PostMortemsQuestionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 post mortems questions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1PostMortemsQuestionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) WithTimeout(timeout time.Duration) *PutV1PostMortemsQuestionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) WithContext(ctx context.Context) *PutV1PostMortemsQuestionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) WithHTTPClient(client *http.Client) *PutV1PostMortemsQuestionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPutV1PostMortemsQuestions adds the putV1PostMortemsQuestions to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) WithPutV1PostMortemsQuestions(putV1PostMortemsQuestions *models.PutV1PostMortemsQuestions) *PutV1PostMortemsQuestionsParams {
	o.SetPutV1PostMortemsQuestions(putV1PostMortemsQuestions)
	return o
}

// SetPutV1PostMortemsQuestions adds the putV1PostMortemsQuestions to the put v1 post mortems questions params
func (o *PutV1PostMortemsQuestionsParams) SetPutV1PostMortemsQuestions(putV1PostMortemsQuestions *models.PutV1PostMortemsQuestions) {
	o.PutV1PostMortemsQuestions = putV1PostMortemsQuestions
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1PostMortemsQuestionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PutV1PostMortemsQuestions != nil {
		if err := r.SetBodyParam(o.PutV1PostMortemsQuestions); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
