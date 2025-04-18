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

// NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParams creates a new PatchV1PostMortemsReportsReportIDReasonsReasonIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParams() *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	return &PatchV1PostMortemsReportsReportIDReasonsReasonIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParamsWithTimeout creates a new PatchV1PostMortemsReportsReportIDReasonsReasonIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParamsWithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	return &PatchV1PostMortemsReportsReportIDReasonsReasonIDParams{
		timeout: timeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParamsWithContext creates a new PatchV1PostMortemsReportsReportIDReasonsReasonIDParams object
// with the ability to set a context for a request.
func NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParamsWithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	return &PatchV1PostMortemsReportsReportIDReasonsReasonIDParams{
		Context: ctx,
	}
}

// NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParamsWithHTTPClient creates a new PatchV1PostMortemsReportsReportIDReasonsReasonIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1PostMortemsReportsReportIDReasonsReasonIDParamsWithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	return &PatchV1PostMortemsReportsReportIDReasonsReasonIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1PostMortemsReportsReportIDReasonsReasonIDParams contains all the parameters to send to the API endpoint

	for the patch v1 post mortems reports report Id reasons reason Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1PostMortemsReportsReportIDReasonsReasonIDParams struct {

	// PatchV1PostMortemsReportsReportIDReasonsReasonID.
	PatchV1PostMortemsReportsReportIDReasonsReasonID *models.PatchV1PostMortemsReportsReportIDReasonsReasonID

	// ReasonID.
	ReasonID string

	// ReportID.
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 post mortems reports report Id reasons reason Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WithDefaults() *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 post mortems reports report Id reasons reason Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchV1PostMortemsReportsReportIDReasonsReasonID adds the patchV1PostMortemsReportsReportIDReasonsReasonID to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WithPatchV1PostMortemsReportsReportIDReasonsReasonID(patchV1PostMortemsReportsReportIDReasonsReasonID *models.PatchV1PostMortemsReportsReportIDReasonsReasonID) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	o.SetPatchV1PostMortemsReportsReportIDReasonsReasonID(patchV1PostMortemsReportsReportIDReasonsReasonID)
	return o
}

// SetPatchV1PostMortemsReportsReportIDReasonsReasonID adds the patchV1PostMortemsReportsReportIdReasonsReasonId to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) SetPatchV1PostMortemsReportsReportIDReasonsReasonID(patchV1PostMortemsReportsReportIDReasonsReasonID *models.PatchV1PostMortemsReportsReportIDReasonsReasonID) {
	o.PatchV1PostMortemsReportsReportIDReasonsReasonID = patchV1PostMortemsReportsReportIDReasonsReasonID
}

// WithReasonID adds the reasonID to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WithReasonID(reasonID string) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	o.SetReasonID(reasonID)
	return o
}

// SetReasonID adds the reasonId to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) SetReasonID(reasonID string) {
	o.ReasonID = reasonID
}

// WithReportID adds the reportID to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WithReportID(reportID string) *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the patch v1 post mortems reports report Id reasons reason Id params
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PostMortemsReportsReportIDReasonsReasonIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PatchV1PostMortemsReportsReportIDReasonsReasonID != nil {
		if err := r.SetBodyParam(o.PatchV1PostMortemsReportsReportIDReasonsReasonID); err != nil {
			return err
		}
	}

	// path param reason_id
	if err := r.SetPathParam("reason_id", o.ReasonID); err != nil {
		return err
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
