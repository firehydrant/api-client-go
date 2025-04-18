// Code generated by go-swagger; DO NOT EDIT.

package reporting

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

// NewGetV1MetricsInfraTypeInfraIDParams creates a new GetV1MetricsInfraTypeInfraIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1MetricsInfraTypeInfraIDParams() *GetV1MetricsInfraTypeInfraIDParams {
	return &GetV1MetricsInfraTypeInfraIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1MetricsInfraTypeInfraIDParamsWithTimeout creates a new GetV1MetricsInfraTypeInfraIDParams object
// with the ability to set a timeout on a request.
func NewGetV1MetricsInfraTypeInfraIDParamsWithTimeout(timeout time.Duration) *GetV1MetricsInfraTypeInfraIDParams {
	return &GetV1MetricsInfraTypeInfraIDParams{
		timeout: timeout,
	}
}

// NewGetV1MetricsInfraTypeInfraIDParamsWithContext creates a new GetV1MetricsInfraTypeInfraIDParams object
// with the ability to set a context for a request.
func NewGetV1MetricsInfraTypeInfraIDParamsWithContext(ctx context.Context) *GetV1MetricsInfraTypeInfraIDParams {
	return &GetV1MetricsInfraTypeInfraIDParams{
		Context: ctx,
	}
}

// NewGetV1MetricsInfraTypeInfraIDParamsWithHTTPClient creates a new GetV1MetricsInfraTypeInfraIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1MetricsInfraTypeInfraIDParamsWithHTTPClient(client *http.Client) *GetV1MetricsInfraTypeInfraIDParams {
	return &GetV1MetricsInfraTypeInfraIDParams{
		HTTPClient: client,
	}
}

/*
GetV1MetricsInfraTypeInfraIDParams contains all the parameters to send to the API endpoint

	for the get v1 metrics infra type infra Id operation.

	Typically these are written to a http.Request.
*/
type GetV1MetricsInfraTypeInfraIDParams struct {

	/* EndDate.

	   The end date to return metrics from, defaults to today

	   Format: date
	*/
	EndDate *strfmt.Date

	/* InfraID.

	   Component UUID
	*/
	InfraID string

	// InfraType.
	InfraType string

	/* StartDate.

	   The start date to return metrics from; defaults to 30 days ago

	   Format: date
	*/
	StartDate *strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 metrics infra type infra Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MetricsInfraTypeInfraIDParams) WithDefaults() *GetV1MetricsInfraTypeInfraIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 metrics infra type infra Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MetricsInfraTypeInfraIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) WithTimeout(timeout time.Duration) *GetV1MetricsInfraTypeInfraIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) WithContext(ctx context.Context) *GetV1MetricsInfraTypeInfraIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) WithHTTPClient(client *http.Client) *GetV1MetricsInfraTypeInfraIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) WithEndDate(endDate *strfmt.Date) *GetV1MetricsInfraTypeInfraIDParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithInfraID adds the infraID to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) WithInfraID(infraID string) *GetV1MetricsInfraTypeInfraIDParams {
	o.SetInfraID(infraID)
	return o
}

// SetInfraID adds the infraId to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) SetInfraID(infraID string) {
	o.InfraID = infraID
}

// WithInfraType adds the infraType to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) WithInfraType(infraType string) *GetV1MetricsInfraTypeInfraIDParams {
	o.SetInfraType(infraType)
	return o
}

// SetInfraType adds the infraType to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) SetInfraType(infraType string) {
	o.InfraType = infraType
}

// WithStartDate adds the startDate to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) WithStartDate(startDate *strfmt.Date) *GetV1MetricsInfraTypeInfraIDParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get v1 metrics infra type infra Id params
func (o *GetV1MetricsInfraTypeInfraIDParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1MetricsInfraTypeInfraIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.Date

		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {

			if err := r.SetQueryParam("end_date", qEndDate); err != nil {
				return err
			}
		}
	}

	// path param infra_id
	if err := r.SetPathParam("infra_id", o.InfraID); err != nil {
		return err
	}

	// path param infra_type
	if err := r.SetPathParam("infra_type", o.InfraType); err != nil {
		return err
	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.Date

		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {

			if err := r.SetQueryParam("start_date", qStartDate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
