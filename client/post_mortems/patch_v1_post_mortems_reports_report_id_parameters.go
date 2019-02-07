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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchV1PostMortemsReportsReportIDParams creates a new PatchV1PostMortemsReportsReportIDParams object
// with the default values initialized.
func NewPatchV1PostMortemsReportsReportIDParams() *PatchV1PostMortemsReportsReportIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDParamsWithTimeout creates a new PatchV1PostMortemsReportsReportIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV1PostMortemsReportsReportIDParamsWithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDParams{

		timeout: timeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDParamsWithContext creates a new PatchV1PostMortemsReportsReportIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV1PostMortemsReportsReportIDParamsWithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDParams{

		Context: ctx,
	}
}

// NewPatchV1PostMortemsReportsReportIDParamsWithHTTPClient creates a new PatchV1PostMortemsReportsReportIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV1PostMortemsReportsReportIDParamsWithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDParams{
		HTTPClient: client,
	}
}

/*PatchV1PostMortemsReportsReportIDParams contains all the parameters to send to the API endpoint
for the patch v1 post mortems reports report Id operation typically these are written to a http.Request
*/
type PatchV1PostMortemsReportsReportIDParams struct {

	/*Name*/
	Name *string
	/*ReportID*/
	ReportID string
	/*Summary*/
	Summary *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) WithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) WithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) WithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) WithName(name *string) *PatchV1PostMortemsReportsReportIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) SetName(name *string) {
	o.Name = name
}

// WithReportID adds the reportID to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) WithReportID(reportID string) *PatchV1PostMortemsReportsReportIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WithSummary adds the summary to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) WithSummary(summary *string) *PatchV1PostMortemsReportsReportIDParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the patch v1 post mortems reports report Id params
func (o *PatchV1PostMortemsReportsReportIDParams) SetSummary(summary *string) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PostMortemsReportsReportIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// form param name
		var frName string
		if o.Name != nil {
			frName = *o.Name
		}
		fName := frName
		if fName != "" {
			if err := r.SetFormParam("name", fName); err != nil {
				return err
			}
		}

	}

	// path param report_id
	if err := r.SetPathParam("report_id", o.ReportID); err != nil {
		return err
	}

	if o.Summary != nil {

		// form param summary
		var frSummary string
		if o.Summary != nil {
			frSummary = *o.Summary
		}
		fSummary := frSummary
		if fSummary != "" {
			if err := r.SetFormParam("summary", fSummary); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
