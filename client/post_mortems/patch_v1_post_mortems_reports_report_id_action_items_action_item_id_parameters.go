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

	models "github.com/firehydrant/api-client-go/models"
)

// NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams creates a new PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams object
// with the default values initialized.
func NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams() *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParamsWithTimeout creates a new PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParamsWithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams{

		timeout: timeout,
	}
}

// NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParamsWithContext creates a new PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParamsWithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams{

		Context: ctx,
	}
}

// NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParamsWithHTTPClient creates a new PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchV1PostMortemsReportsReportIDActionItemsActionItemIDParamsWithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	var ()
	return &PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams{
		HTTPClient: client,
	}
}

/*PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams contains all the parameters to send to the API endpoint
for the patch v1 post mortems reports report Id action items action item Id operation typically these are written to a http.Request
*/
type PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams struct {

	/*V1PostMortemsReportsReportIDActionItems*/
	V1PostMortemsReportsReportIDActionItems *models.PatchV1PostMortemsReportsReportIDActionItems
	/*ActionItemID*/
	ActionItemID string
	/*ReportID*/
	ReportID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) WithTimeout(timeout time.Duration) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) WithContext(ctx context.Context) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) WithHTTPClient(client *http.Client) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1PostMortemsReportsReportIDActionItems adds the v1PostMortemsReportsReportIDActionItems to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) WithV1PostMortemsReportsReportIDActionItems(v1PostMortemsReportsReportIDActionItems *models.PatchV1PostMortemsReportsReportIDActionItems) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	o.SetV1PostMortemsReportsReportIDActionItems(v1PostMortemsReportsReportIDActionItems)
	return o
}

// SetV1PostMortemsReportsReportIDActionItems adds the v1PostMortemsReportsReportIdActionItems to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) SetV1PostMortemsReportsReportIDActionItems(v1PostMortemsReportsReportIDActionItems *models.PatchV1PostMortemsReportsReportIDActionItems) {
	o.V1PostMortemsReportsReportIDActionItems = v1PostMortemsReportsReportIDActionItems
}

// WithActionItemID adds the actionItemID to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) WithActionItemID(actionItemID string) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	o.SetActionItemID(actionItemID)
	return o
}

// SetActionItemID adds the actionItemId to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) SetActionItemID(actionItemID string) {
	o.ActionItemID = actionItemID
}

// WithReportID adds the reportID to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) WithReportID(reportID string) *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams {
	o.SetReportID(reportID)
	return o
}

// SetReportID adds the reportId to the patch v1 post mortems reports report Id action items action item Id params
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) SetReportID(reportID string) {
	o.ReportID = reportID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1PostMortemsReportsReportIDActionItemsActionItemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.V1PostMortemsReportsReportIDActionItems != nil {
		if err := r.SetBodyParam(o.V1PostMortemsReportsReportIDActionItems); err != nil {
			return err
		}
	}

	// path param action_item_id
	if err := r.SetPathParam("action_item_id", o.ActionItemID); err != nil {
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
