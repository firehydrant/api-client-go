// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetV1IntegrationsAwsConnectionsParams creates a new GetV1IntegrationsAwsConnectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IntegrationsAwsConnectionsParams() *GetV1IntegrationsAwsConnectionsParams {
	return &GetV1IntegrationsAwsConnectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsAwsConnectionsParamsWithTimeout creates a new GetV1IntegrationsAwsConnectionsParams object
// with the ability to set a timeout on a request.
func NewGetV1IntegrationsAwsConnectionsParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsAwsConnectionsParams {
	return &GetV1IntegrationsAwsConnectionsParams{
		timeout: timeout,
	}
}

// NewGetV1IntegrationsAwsConnectionsParamsWithContext creates a new GetV1IntegrationsAwsConnectionsParams object
// with the ability to set a context for a request.
func NewGetV1IntegrationsAwsConnectionsParamsWithContext(ctx context.Context) *GetV1IntegrationsAwsConnectionsParams {
	return &GetV1IntegrationsAwsConnectionsParams{
		Context: ctx,
	}
}

// NewGetV1IntegrationsAwsConnectionsParamsWithHTTPClient creates a new GetV1IntegrationsAwsConnectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IntegrationsAwsConnectionsParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsAwsConnectionsParams {
	return &GetV1IntegrationsAwsConnectionsParams{
		HTTPClient: client,
	}
}

/* GetV1IntegrationsAwsConnectionsParams contains all the parameters to send to the API endpoint
   for the get v1 integrations aws connections operation.

   Typically these are written to a http.Request.
*/
type GetV1IntegrationsAwsConnectionsParams struct {

	/* AwsAccountID.

	   AWS account ID containing the role to be assumed
	*/
	AwsAccountID *string

	/* ExternalID.

	   The external ID supplied when assuming the role
	*/
	ExternalID *string

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* TargetArn.

	   ARN of the role to be assumed
	*/
	TargetArn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 integrations aws connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsAwsConnectionsParams) WithDefaults() *GetV1IntegrationsAwsConnectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 integrations aws connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsAwsConnectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsAwsConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithContext(ctx context.Context) *GetV1IntegrationsAwsConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsAwsConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAwsAccountID adds the awsAccountID to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithAwsAccountID(awsAccountID *string) *GetV1IntegrationsAwsConnectionsParams {
	o.SetAwsAccountID(awsAccountID)
	return o
}

// SetAwsAccountID adds the awsAccountId to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetAwsAccountID(awsAccountID *string) {
	o.AwsAccountID = awsAccountID
}

// WithExternalID adds the externalID to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithExternalID(externalID *string) *GetV1IntegrationsAwsConnectionsParams {
	o.SetExternalID(externalID)
	return o
}

// SetExternalID adds the externalId to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetExternalID(externalID *string) {
	o.ExternalID = externalID
}

// WithPage adds the page to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithPage(page *int32) *GetV1IntegrationsAwsConnectionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithPerPage(perPage *int32) *GetV1IntegrationsAwsConnectionsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithTargetArn adds the targetArn to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) WithTargetArn(targetArn *string) *GetV1IntegrationsAwsConnectionsParams {
	o.SetTargetArn(targetArn)
	return o
}

// SetTargetArn adds the targetArn to the get v1 integrations aws connections params
func (o *GetV1IntegrationsAwsConnectionsParams) SetTargetArn(targetArn *string) {
	o.TargetArn = targetArn
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsAwsConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AwsAccountID != nil {

		// query param aws_account_id
		var qrAwsAccountID string

		if o.AwsAccountID != nil {
			qrAwsAccountID = *o.AwsAccountID
		}
		qAwsAccountID := qrAwsAccountID
		if qAwsAccountID != "" {

			if err := r.SetQueryParam("aws_account_id", qAwsAccountID); err != nil {
				return err
			}
		}
	}

	if o.ExternalID != nil {

		// query param external_id
		var qrExternalID string

		if o.ExternalID != nil {
			qrExternalID = *o.ExternalID
		}
		qExternalID := qrExternalID
		if qExternalID != "" {

			if err := r.SetQueryParam("external_id", qExternalID); err != nil {
				return err
			}
		}
	}

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

	if o.TargetArn != nil {

		// query param target_arn
		var qrTargetArn string

		if o.TargetArn != nil {
			qrTargetArn = *o.TargetArn
		}
		qTargetArn := qrTargetArn
		if qTargetArn != "" {

			if err := r.SetQueryParam("target_arn", qTargetArn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}