// Code generated by go-swagger; DO NOT EDIT.

package retrospectives

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

// NewPutV1PostMortemsReportsReportIDReasonsOrderParams creates a new PutV1PostMortemsReportsReportIDReasonsOrderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1PostMortemsReportsReportIDReasonsOrderParams() *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	return &PutV1PostMortemsReportsReportIDReasonsOrderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1PostMortemsReportsReportIDReasonsOrderParamsWithTimeout creates a new PutV1PostMortemsReportsReportIDReasonsOrderParams object
// with the ability to set a timeout on a request.
func NewPutV1PostMortemsReportsReportIDReasonsOrderParamsWithTimeout(timeout time.Duration) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	return &PutV1PostMortemsReportsReportIDReasonsOrderParams{
		timeout: timeout,
	}
}

// NewPutV1PostMortemsReportsReportIDReasonsOrderParamsWithContext creates a new PutV1PostMortemsReportsReportIDReasonsOrderParams object
// with the ability to set a context for a request.
func NewPutV1PostMortemsReportsReportIDReasonsOrderParamsWithContext(ctx context.Context) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	return &PutV1PostMortemsReportsReportIDReasonsOrderParams{
		Context: ctx,
	}
}

// NewPutV1PostMortemsReportsReportIDReasonsOrderParamsWithHTTPClient creates a new PutV1PostMortemsReportsReportIDReasonsOrderParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1PostMortemsReportsReportIDReasonsOrderParamsWithHTTPClient(client *http.Client) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	return &PutV1PostMortemsReportsReportIDReasonsOrderParams{
		HTTPClient: client,
	}
}

/*
PutV1PostMortemsReportsReportIDReasonsOrderParams contains all the parameters to send to the API endpoint

	for the put v1 post mortems reports report Id reasons order operation.

	Typically these are written to a http.Request.
*/
type PutV1PostMortemsReportsReportIDReasonsOrderParams struct {

	// PutV1PostMortemsReportsReportIDReasonsOrder.
	PutV1PostMortemsReportsReportIDReasonsOrder *models.PutV1PostMortemsReportsReportIDReasonsOrder

	// ReportID.
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 post mortems reports report Id reasons order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) WithDefaults() *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 post mortems reports report Id reasons order params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) WithTimeout(timeout time.Duration) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) WithContext(ctx context.Context) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) WithHTTPClient(client *http.Client) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPutV1PostMortemsReportsReportIDReasonsOrder adds the putV1PostMortemsReportsReportIDReasonsOrder to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) WithPutV1PostMortemsReportsReportIDReasonsOrder(putV1PostMortemsReportsReportIDReasonsOrder *models.PutV1PostMortemsReportsReportIDReasonsOrder) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	o.SetPutV1PostMortemsReportsReportIDReasonsOrder(putV1PostMortemsReportsReportIDReasonsOrder)
	return o
}

// SetPutV1PostMortemsReportsReportIDReasonsOrder adds the putV1PostMortemsReportsReportIdReasonsOrder to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) SetPutV1PostMortemsReportsReportIDReasonsOrder(putV1PostMortemsReportsReportIDReasonsOrder *models.PutV1PostMortemsReportsReportIDReasonsOrder) {
	o.PutV1PostMortemsReportsReportIDReasonsOrder = putV1PostMortemsReportsReportIDReasonsOrder
}

// WithReportID adds the reportID to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) WithReportID(reportID string) *PutV1PostMortemsReportsReportIDReasonsOrderParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the put v1 post mortems reports report Id reasons order params
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1PostMortemsReportsReportIDReasonsOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PutV1PostMortemsReportsReportIDReasonsOrder != nil {
		if err := r.SetBodyParam(o.PutV1PostMortemsReportsReportIDReasonsOrder); err != nil {
			return err
		}
	}

	// path param report_id
	if err := r.SetPathParam("report_id", o.ReportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
