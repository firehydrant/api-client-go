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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetV1PostMortemsReportsReportIDEventsParams creates a new GetV1PostMortemsReportsReportIDEventsParams object
// with the default values initialized.
func NewGetV1PostMortemsReportsReportIDEventsParams() *GetV1PostMortemsReportsReportIDEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PostMortemsReportsReportIDEventsParamsWithTimeout creates a new GetV1PostMortemsReportsReportIDEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1PostMortemsReportsReportIDEventsParamsWithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDEventsParams{

		timeout: timeout,
	}
}

// NewGetV1PostMortemsReportsReportIDEventsParamsWithContext creates a new GetV1PostMortemsReportsReportIDEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1PostMortemsReportsReportIDEventsParamsWithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDEventsParams{

		Context: ctx,
	}
}

// NewGetV1PostMortemsReportsReportIDEventsParamsWithHTTPClient creates a new GetV1PostMortemsReportsReportIDEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1PostMortemsReportsReportIDEventsParamsWithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDEventsParams{
		HTTPClient: client,
	}
}

/*GetV1PostMortemsReportsReportIDEventsParams contains all the parameters to send to the API endpoint
for the get v1 post mortems reports report Id events operation typically these are written to a http.Request
*/
type GetV1PostMortemsReportsReportIDEventsParams struct {

	/*ReportID*/
	ReportID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) WithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) WithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) WithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportID adds the reportID to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) WithReportID(reportID int32) *GetV1PostMortemsReportsReportIDEventsParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the get v1 post mortems reports report Id events params
func (o *GetV1PostMortemsReportsReportIDEventsParams) SetReportID(reportID int32) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PostMortemsReportsReportIDEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param report_id
	if err := r.SetPathParam("report_id", swag.FormatInt32(o.ReportID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
