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

/*
PutV1NuncConnectionsNuncConnectionIDParams contains all the parameters to send to the API endpoint

	for the put v1 nunc connections nunc connection Id operation.

	Typically these are written to a http.Request.
*/
type PutV1NuncConnectionsNuncConnectionIDParams struct {

	// ButtonBackgroundColor.
	ButtonBackgroundColor *string

	// ButtonTextColor.
	ButtonTextColor *string

	// CompanyName.
	CompanyName *string

	// CompanyTosURL.
	CompanyTosURL *string

	// CompanyWebsite.
	CompanyWebsite *string

	// CoverImage.
	CoverImage runtime.NamedReadCloser

	// Favicon.
	Favicon runtime.NamedReadCloser

	// GreetingBody.
	GreetingBody *string

	// GreetingTitle.
	GreetingTitle *string

	// LinkColor.
	LinkColor *string

	// Logo.
	Logo runtime.NamedReadCloser

	// NuncConnectionID.
	//
	// Format: int32
	NuncConnectionID int32

	// OpenGraphImage.
	OpenGraphImage runtime.NamedReadCloser

	// OperationalMessage.
	OperationalMessage *string

	// PrimaryColor.
	PrimaryColor *string

	// SecondaryColor.
	SecondaryColor *string

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

// WithButtonBackgroundColor adds the buttonBackgroundColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithButtonBackgroundColor(buttonBackgroundColor *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetButtonBackgroundColor(buttonBackgroundColor)
	return o
}

// SetButtonBackgroundColor adds the buttonBackgroundColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetButtonBackgroundColor(buttonBackgroundColor *string) {
	o.ButtonBackgroundColor = buttonBackgroundColor
}

// WithButtonTextColor adds the buttonTextColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithButtonTextColor(buttonTextColor *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetButtonTextColor(buttonTextColor)
	return o
}

// SetButtonTextColor adds the buttonTextColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetButtonTextColor(buttonTextColor *string) {
	o.ButtonTextColor = buttonTextColor
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

// WithCoverImage adds the coverImage to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithCoverImage(coverImage runtime.NamedReadCloser) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetCoverImage(coverImage)
	return o
}

// SetCoverImage adds the coverImage to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetCoverImage(coverImage runtime.NamedReadCloser) {
	o.CoverImage = coverImage
}

// WithFavicon adds the favicon to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithFavicon(favicon runtime.NamedReadCloser) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetFavicon(favicon)
	return o
}

// SetFavicon adds the favicon to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetFavicon(favicon runtime.NamedReadCloser) {
	o.Favicon = favicon
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

// WithLinkColor adds the linkColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithLinkColor(linkColor *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetLinkColor(linkColor)
	return o
}

// SetLinkColor adds the linkColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetLinkColor(linkColor *string) {
	o.LinkColor = linkColor
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

// WithOpenGraphImage adds the openGraphImage to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithOpenGraphImage(openGraphImage runtime.NamedReadCloser) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetOpenGraphImage(openGraphImage)
	return o
}

// SetOpenGraphImage adds the openGraphImage to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetOpenGraphImage(openGraphImage runtime.NamedReadCloser) {
	o.OpenGraphImage = openGraphImage
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

// WithPrimaryColor adds the primaryColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithPrimaryColor(primaryColor *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetPrimaryColor(primaryColor)
	return o
}

// SetPrimaryColor adds the primaryColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetPrimaryColor(primaryColor *string) {
	o.PrimaryColor = primaryColor
}

// WithSecondaryColor adds the secondaryColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithSecondaryColor(secondaryColor *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetSecondaryColor(secondaryColor)
	return o
}

// SetSecondaryColor adds the secondaryColor to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetSecondaryColor(secondaryColor *string) {
	o.SecondaryColor = secondaryColor
}

// WriteToRequest writes these params to a swagger request
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ButtonBackgroundColor != nil {

		// form param button_background_color
		var frButtonBackgroundColor string
		if o.ButtonBackgroundColor != nil {
			frButtonBackgroundColor = *o.ButtonBackgroundColor
		}
		fButtonBackgroundColor := frButtonBackgroundColor
		if fButtonBackgroundColor != "" {
			if err := r.SetFormParam("button_background_color", fButtonBackgroundColor); err != nil {
				return err
			}
		}
	}

	if o.ButtonTextColor != nil {

		// form param button_text_color
		var frButtonTextColor string
		if o.ButtonTextColor != nil {
			frButtonTextColor = *o.ButtonTextColor
		}
		fButtonTextColor := frButtonTextColor
		if fButtonTextColor != "" {
			if err := r.SetFormParam("button_text_color", fButtonTextColor); err != nil {
				return err
			}
		}
	}

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

	if o.CoverImage != nil {

		if o.CoverImage != nil {
			// form file param cover_image
			if err := r.SetFileParam("cover_image", o.CoverImage); err != nil {
				return err
			}
		}
	}

	if o.Favicon != nil {

		if o.Favicon != nil {
			// form file param favicon
			if err := r.SetFileParam("favicon", o.Favicon); err != nil {
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

	if o.LinkColor != nil {

		// form param link_color
		var frLinkColor string
		if o.LinkColor != nil {
			frLinkColor = *o.LinkColor
		}
		fLinkColor := frLinkColor
		if fLinkColor != "" {
			if err := r.SetFormParam("link_color", fLinkColor); err != nil {
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

	if o.OpenGraphImage != nil {

		if o.OpenGraphImage != nil {
			// form file param open_graph_image
			if err := r.SetFileParam("open_graph_image", o.OpenGraphImage); err != nil {
				return err
			}
		}
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

	if o.PrimaryColor != nil {

		// form param primary_color
		var frPrimaryColor string
		if o.PrimaryColor != nil {
			frPrimaryColor = *o.PrimaryColor
		}
		fPrimaryColor := frPrimaryColor
		if fPrimaryColor != "" {
			if err := r.SetFormParam("primary_color", fPrimaryColor); err != nil {
				return err
			}
		}
	}

	if o.SecondaryColor != nil {

		// form param secondary_color
		var frSecondaryColor string
		if o.SecondaryColor != nil {
			frSecondaryColor = *o.SecondaryColor
		}
		fSecondaryColor := frSecondaryColor
		if fSecondaryColor != "" {
			if err := r.SetFormParam("secondary_color", fSecondaryColor); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
