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

// NewGetV1PostMortemsReportsReportIDIncidentEventsParams creates a new GetV1PostMortemsReportsReportIDIncidentEventsParams object
// with the default values initialized.
func NewGetV1PostMortemsReportsReportIDIncidentEventsParams() *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDIncidentEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PostMortemsReportsReportIDIncidentEventsParamsWithTimeout creates a new GetV1PostMortemsReportsReportIDIncidentEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1PostMortemsReportsReportIDIncidentEventsParamsWithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDIncidentEventsParams{

		timeout: timeout,
	}
}

// NewGetV1PostMortemsReportsReportIDIncidentEventsParamsWithContext creates a new GetV1PostMortemsReportsReportIDIncidentEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1PostMortemsReportsReportIDIncidentEventsParamsWithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDIncidentEventsParams{

		Context: ctx,
	}
}

// NewGetV1PostMortemsReportsReportIDIncidentEventsParamsWithHTTPClient creates a new GetV1PostMortemsReportsReportIDIncidentEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1PostMortemsReportsReportIDIncidentEventsParamsWithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	var ()
	return &GetV1PostMortemsReportsReportIDIncidentEventsParams{
		HTTPClient: client,
	}
}

/*GetV1PostMortemsReportsReportIDIncidentEventsParams contains all the parameters to send to the API endpoint
for the get v1 post mortems reports report Id incident events operation typically these are written to a http.Request
*/
type GetV1PostMortemsReportsReportIDIncidentEventsParams struct {

	/*ReportID*/
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) WithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) WithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) WithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReportID adds the reportID to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) WithReportID(reportID string) *GetV1PostMortemsReportsReportIDIncidentEventsParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the get v1 post mortems reports report Id incident events params
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PostMortemsReportsReportIDIncidentEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param report_id
	if err := r.SetPathParam("report_id", o.ReportID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
