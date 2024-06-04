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
)

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams() *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParamsWithTimeout creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams object
// with the ability to set a timeout on a request.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParamsWithTimeout(timeout time.Duration) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams{
		timeout: timeout,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParamsWithContext creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams object
// with the ability to set a context for a request.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParamsWithContext(ctx context.Context) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams{
		Context: ctx,
	}
}

// NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParamsWithHTTPClient creates a new GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParamsWithHTTPClient(client *http.Client) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	return &GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams{
		HTTPClient: client,
	}
}

/*
GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams contains all the parameters to send to the API endpoint

	for the get v1 runbooks executions execution Id steps step Id votes status operation.

	Typically these are written to a http.Request.
*/
type GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams struct {

	// ExecutionID.
	ExecutionID string

	// StepID.
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 runbooks executions execution Id steps step Id votes status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) WithDefaults() *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 runbooks executions execution Id steps step Id votes status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) WithTimeout(timeout time.Duration) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) WithContext(ctx context.Context) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) WithHTTPClient(client *http.Client) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecutionID adds the executionID to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) WithExecutionID(executionID string) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	o.SetExecutionID(executionID)
	return o
}

// SetExecutionID adds the executionId to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) SetExecutionID(executionID string) {
	o.ExecutionID = executionID
}

// WithStepID adds the stepID to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) WithStepID(stepID string) *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the get v1 runbooks executions execution Id steps step Id votes status params
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1RunbooksExecutionsExecutionIDStepsStepIDVotesStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param execution_id
	if err := r.SetPathParam("execution_id", o.ExecutionID); err != nil {
		return err
	}

	// path param step_id
	if err := r.SetPathParam("step_id", o.StepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
