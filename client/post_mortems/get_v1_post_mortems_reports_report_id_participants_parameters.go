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
	"github.com/go-openapi/swag"
)

// NewGetV1PostMortemsReportsReportIDParticipantsParams creates a new GetV1PostMortemsReportsReportIDParticipantsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1PostMortemsReportsReportIDParticipantsParams() *GetV1PostMortemsReportsReportIDParticipantsParams {
	return &GetV1PostMortemsReportsReportIDParticipantsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PostMortemsReportsReportIDParticipantsParamsWithTimeout creates a new GetV1PostMortemsReportsReportIDParticipantsParams object
// with the ability to set a timeout on a request.
func NewGetV1PostMortemsReportsReportIDParticipantsParamsWithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDParticipantsParams {
	return &GetV1PostMortemsReportsReportIDParticipantsParams{
		timeout: timeout,
	}
}

// NewGetV1PostMortemsReportsReportIDParticipantsParamsWithContext creates a new GetV1PostMortemsReportsReportIDParticipantsParams object
// with the ability to set a context for a request.
func NewGetV1PostMortemsReportsReportIDParticipantsParamsWithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDParticipantsParams {
	return &GetV1PostMortemsReportsReportIDParticipantsParams{
		Context: ctx,
	}
}

// NewGetV1PostMortemsReportsReportIDParticipantsParamsWithHTTPClient creates a new GetV1PostMortemsReportsReportIDParticipantsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1PostMortemsReportsReportIDParticipantsParamsWithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDParticipantsParams {
	return &GetV1PostMortemsReportsReportIDParticipantsParams{
		HTTPClient: client,
	}
}

/*
GetV1PostMortemsReportsReportIDParticipantsParams contains all the parameters to send to the API endpoint

	for the get v1 post mortems reports report Id participants operation.

	Typically these are written to a http.Request.
*/
type GetV1PostMortemsReportsReportIDParticipantsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	// ReportID.
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 post mortems reports report Id participants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WithDefaults() *GetV1PostMortemsReportsReportIDParticipantsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 post mortems reports report Id participants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDParticipantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDParticipantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDParticipantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WithPage(page *int32) *GetV1PostMortemsReportsReportIDParticipantsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WithPerPage(perPage *int32) *GetV1PostMortemsReportsReportIDParticipantsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithReportID adds the reportID to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WithReportID(reportID string) *GetV1PostMortemsReportsReportIDParticipantsParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the get v1 post mortems reports report Id participants params
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PostMortemsReportsReportIDParticipantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
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
