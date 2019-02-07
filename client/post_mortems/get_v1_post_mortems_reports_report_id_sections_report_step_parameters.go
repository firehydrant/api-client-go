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

// NewGetV1PostMortemsReportsReportIDSectionsReportStepParams creates a new GetV1PostMortemsReportsReportIDSectionsReportStepParams object
// with the default values initialized.
func NewGetV1PostMortemsReportsReportIDSectionsReportStepParams() *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	var ()
	return &GetV1PostMortemsReportsReportIDSectionsReportStepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1PostMortemsReportsReportIDSectionsReportStepParamsWithTimeout creates a new GetV1PostMortemsReportsReportIDSectionsReportStepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetV1PostMortemsReportsReportIDSectionsReportStepParamsWithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	var ()
	return &GetV1PostMortemsReportsReportIDSectionsReportStepParams{

		timeout: timeout,
	}
}

// NewGetV1PostMortemsReportsReportIDSectionsReportStepParamsWithContext creates a new GetV1PostMortemsReportsReportIDSectionsReportStepParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetV1PostMortemsReportsReportIDSectionsReportStepParamsWithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	var ()
	return &GetV1PostMortemsReportsReportIDSectionsReportStepParams{

		Context: ctx,
	}
}

// NewGetV1PostMortemsReportsReportIDSectionsReportStepParamsWithHTTPClient creates a new GetV1PostMortemsReportsReportIDSectionsReportStepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetV1PostMortemsReportsReportIDSectionsReportStepParamsWithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	var ()
	return &GetV1PostMortemsReportsReportIDSectionsReportStepParams{
		HTTPClient: client,
	}
}

/*GetV1PostMortemsReportsReportIDSectionsReportStepParams contains all the parameters to send to the API endpoint
for the get v1 post mortems reports report Id sections report step operation typically these are written to a http.Request
*/
type GetV1PostMortemsReportsReportIDSectionsReportStepParams struct {

	/*Page*/
	Page *int32
	/*PerPage*/
	PerPage *int32
	/*ReportID*/
	ReportID string
	/*ReportStep
	  The report step the sections belong to

	*/
	ReportStep string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WithTimeout(timeout time.Duration) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WithContext(ctx context.Context) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WithHTTPClient(client *http.Client) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WithPage(page *int32) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WithPerPage(perPage *int32) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithReportID adds the reportID to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WithReportID(reportID string) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WithReportStep adds the reportStep to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WithReportStep(reportStep string) *GetV1PostMortemsReportsReportIDSectionsReportStepParams {
	o.SetReportStep(reportStep)
	return o
}

// SetReportStep adds the reportStep to the get v1 post mortems reports report Id sections report step params
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) SetReportStep(reportStep string) {
	o.ReportStep = reportStep
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1PostMortemsReportsReportIDSectionsReportStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param report_step
	if err := r.SetPathParam("report_step", o.ReportStep); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
