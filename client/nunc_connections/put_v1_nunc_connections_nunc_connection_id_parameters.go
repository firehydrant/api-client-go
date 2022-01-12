// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

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

// NewPutV1NuncConnectionsNuncConnectionIDParams creates a new PutV1NuncConnectionsNuncConnectionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutV1NuncConnectionsNuncConnectionIDParams() *PutV1NuncConnectionsNuncConnectionIDParams {
	return &PutV1NuncConnectionsNuncConnectionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutV1NuncConnectionsNuncConnectionIDParamsWithTimeout creates a new PutV1NuncConnectionsNuncConnectionIDParams object
// with the ability to set a timeout on a request.
func NewPutV1NuncConnectionsNuncConnectionIDParamsWithTimeout(timeout time.Duration) *PutV1NuncConnectionsNuncConnectionIDParams {
	return &PutV1NuncConnectionsNuncConnectionIDParams{
		timeout: timeout,
	}
}

// NewPutV1NuncConnectionsNuncConnectionIDParamsWithContext creates a new PutV1NuncConnectionsNuncConnectionIDParams object
// with the ability to set a context for a request.
func NewPutV1NuncConnectionsNuncConnectionIDParamsWithContext(ctx context.Context) *PutV1NuncConnectionsNuncConnectionIDParams {
	return &PutV1NuncConnectionsNuncConnectionIDParams{
		Context: ctx,
	}
}

// NewPutV1NuncConnectionsNuncConnectionIDParamsWithHTTPClient creates a new PutV1NuncConnectionsNuncConnectionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutV1NuncConnectionsNuncConnectionIDParamsWithHTTPClient(client *http.Client) *PutV1NuncConnectionsNuncConnectionIDParams {
	return &PutV1NuncConnectionsNuncConnectionIDParams{
		HTTPClient: client,
	}
}

/* PutV1NuncConnectionsNuncConnectionIDParams contains all the parameters to send to the API endpoint
   for the put v1 nunc connections nunc connection Id operation.

   Typically these are written to a http.Request.
*/
type PutV1NuncConnectionsNuncConnectionIDParams struct {

	// CompanyName.
	CompanyName *string

	// CompanyTosURL.
	CompanyTosURL *string

	// CompanyWebsite.
	CompanyWebsite *string

	// GreetingBody.
	GreetingBody *string

	// GreetingTitle.
	GreetingTitle *string

	// Logo.
	Logo runtime.NamedReadCloser

	// NuncConnectionID.
	//
	// Format: int32
	NuncConnectionID int32

	// OperationalMessage.
	OperationalMessage *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put v1 nunc connections nunc connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithDefaults() *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put v1 nunc connections nunc connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithTimeout(timeout time.Duration) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithContext(ctx context.Context) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithHTTPClient(client *http.Client) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompanyName adds the companyName to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithCompanyName(companyName *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetCompanyName(companyName)
	return o
}

// SetCompanyName adds the companyName to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetCompanyName(companyName *string) {
	o.CompanyName = companyName
}

// WithCompanyTosURL adds the companyTosURL to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithCompanyTosURL(companyTosURL *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetCompanyTosURL(companyTosURL)
	return o
}

// SetCompanyTosURL adds the companyTosUrl to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetCompanyTosURL(companyTosURL *string) {
	o.CompanyTosURL = companyTosURL
}

// WithCompanyWebsite adds the companyWebsite to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithCompanyWebsite(companyWebsite *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetCompanyWebsite(companyWebsite)
	return o
}

// SetCompanyWebsite adds the companyWebsite to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetCompanyWebsite(companyWebsite *string) {
	o.CompanyWebsite = companyWebsite
}

// WithGreetingBody adds the greetingBody to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithGreetingBody(greetingBody *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetGreetingBody(greetingBody)
	return o
}

// SetGreetingBody adds the greetingBody to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetGreetingBody(greetingBody *string) {
	o.GreetingBody = greetingBody
}

// WithGreetingTitle adds the greetingTitle to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithGreetingTitle(greetingTitle *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetGreetingTitle(greetingTitle)
	return o
}

// SetGreetingTitle adds the greetingTitle to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetGreetingTitle(greetingTitle *string) {
	o.GreetingTitle = greetingTitle
}

// WithLogo adds the logo to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithLogo(logo runtime.NamedReadCloser) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetLogo(logo)
	return o
}

// SetLogo adds the logo to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetLogo(logo runtime.NamedReadCloser) {
	o.Logo = logo
}

// WithNuncConnectionID adds the nuncConnectionID to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithNuncConnectionID(nuncConnectionID int32) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetNuncConnectionID(nuncConnectionID int32) {
	o.NuncConnectionID = nuncConnectionID
}

// WithOperationalMessage adds the operationalMessage to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithOperationalMessage(operationalMessage *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetOperationalMessage(operationalMessage)
	return o
}

// SetOperationalMessage adds the operationalMessage to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetOperationalMessage(operationalMessage *string) {
	o.OperationalMessage = operationalMessage
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CompanyName != nil {

		// form param company_name
		var frCompanyName string
		if o.CompanyName != nil {
			frCompanyName = *o.CompanyName
		}
		fCompanyName := frCompanyName
		if fCompanyName != "" {
			if err := r.SetFormParam("company_name", fCompanyName); err != nil {
				return err
			}
		}
	}

	if o.CompanyTosURL != nil {

		// form param company_tos_url
		var frCompanyTosURL string
		if o.CompanyTosURL != nil {
			frCompanyTosURL = *o.CompanyTosURL
		}
		fCompanyTosURL := frCompanyTosURL
		if fCompanyTosURL != "" {
			if err := r.SetFormParam("company_tos_url", fCompanyTosURL); err != nil {
				return err
			}
		}
	}

	if o.CompanyWebsite != nil {

		// form param company_website
		var frCompanyWebsite string
		if o.CompanyWebsite != nil {
			frCompanyWebsite = *o.CompanyWebsite
		}
		fCompanyWebsite := frCompanyWebsite
		if fCompanyWebsite != "" {
			if err := r.SetFormParam("company_website", fCompanyWebsite); err != nil {
				return err
			}
		}
	}

	if o.GreetingBody != nil {

		// form param greeting_body
		var frGreetingBody string
		if o.GreetingBody != nil {
			frGreetingBody = *o.GreetingBody
		}
		fGreetingBody := frGreetingBody
		if fGreetingBody != "" {
			if err := r.SetFormParam("greeting_body", fGreetingBody); err != nil {
				return err
			}
		}
	}

	if o.GreetingTitle != nil {

		// form param greeting_title
		var frGreetingTitle string
		if o.GreetingTitle != nil {
			frGreetingTitle = *o.GreetingTitle
		}
		fGreetingTitle := frGreetingTitle
		if fGreetingTitle != "" {
			if err := r.SetFormParam("greeting_title", fGreetingTitle); err != nil {
				return err
			}
		}
	}

	if o.Logo != nil {

		if o.Logo != nil {
			// form file param logo
			if err := r.SetFileParam("logo", o.Logo); err != nil {
				return err
			}
		}
	}

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", swag.FormatInt32(o.NuncConnectionID)); err != nil {
		return err
	}

	if o.OperationalMessage != nil {

		// form param operational_message
		var frOperationalMessage string
		if o.OperationalMessage != nil {
			frOperationalMessage = *o.OperationalMessage
		}
		fOperationalMessage := frOperationalMessage
		if fOperationalMessage != "" {
			if err := r.SetFormParam("operational_message", fOperationalMessage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}