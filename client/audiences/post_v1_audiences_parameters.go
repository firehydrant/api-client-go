// Code generated by go-swagger; DO NOT EDIT.

package audiences

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

// NewPostV1AudiencesParams creates a new PostV1AudiencesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1AudiencesParams() *PostV1AudiencesParams {
	return &PostV1AudiencesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1AudiencesParamsWithTimeout creates a new PostV1AudiencesParams object
// with the ability to set a timeout on a request.
func NewPostV1AudiencesParamsWithTimeout(timeout time.Duration) *PostV1AudiencesParams {
	return &PostV1AudiencesParams{
		timeout: timeout,
	}
}

// NewPostV1AudiencesParamsWithContext creates a new PostV1AudiencesParams object
// with the ability to set a context for a request.
func NewPostV1AudiencesParamsWithContext(ctx context.Context) *PostV1AudiencesParams {
	return &PostV1AudiencesParams{
		Context: ctx,
	}
}

// NewPostV1AudiencesParamsWithHTTPClient creates a new PostV1AudiencesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1AudiencesParamsWithHTTPClient(client *http.Client) *PostV1AudiencesParams {
	return &PostV1AudiencesParams{
		HTTPClient: client,
	}
}

/*
PostV1AudiencesParams contains all the parameters to send to the API endpoint

	for the post v1 audiences operation.

	Typically these are written to a http.Request.
*/
type PostV1AudiencesParams struct {

	/* Default.

	   Whether this is the default audience
	*/
	Default *bool

	/* Description.

	   Description of the audience (max 4000 characters)
	*/
	Description string

	/* DetailsPrompt.

	   The prompt to display when collecting this detail
	*/
	DetailsPrompt []string

	/* DetailsQuestion.

	   The incident detail question (max 255 characters)
	*/
	DetailsQuestion []string

	/* DetailsSlug.

	   Optional unique identifier for this detail
	*/
	DetailsSlug []string

	/* Name.

	   Name of the audience (max 255 characters)
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 audiences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AudiencesParams) WithDefaults() *PostV1AudiencesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 audiences params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1AudiencesParams) SetDefaults() {
	var (
		defaultVarDefault = bool(false)
	)

	val := PostV1AudiencesParams{
		Default: &defaultVarDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the post v1 audiences params
func (o *PostV1AudiencesParams) WithTimeout(timeout time.Duration) *PostV1AudiencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 audiences params
func (o *PostV1AudiencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 audiences params
func (o *PostV1AudiencesParams) WithContext(ctx context.Context) *PostV1AudiencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 audiences params
func (o *PostV1AudiencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 audiences params
func (o *PostV1AudiencesParams) WithHTTPClient(client *http.Client) *PostV1AudiencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 audiences params
func (o *PostV1AudiencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDefault adds the defaultVar to the post v1 audiences params
func (o *PostV1AudiencesParams) WithDefault(defaultVar *bool) *PostV1AudiencesParams {
	o.SetDefault(defaultVar)
	return o
}

// SetDefault adds the default to the post v1 audiences params
func (o *PostV1AudiencesParams) SetDefault(defaultVar *bool) {
	o.Default = defaultVar
}

// WithDescription adds the description to the post v1 audiences params
func (o *PostV1AudiencesParams) WithDescription(description string) *PostV1AudiencesParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the post v1 audiences params
func (o *PostV1AudiencesParams) SetDescription(description string) {
	o.Description = description
}

// WithDetailsPrompt adds the detailsPrompt to the post v1 audiences params
func (o *PostV1AudiencesParams) WithDetailsPrompt(detailsPrompt []string) *PostV1AudiencesParams {
	o.SetDetailsPrompt(detailsPrompt)
	return o
}

// SetDetailsPrompt adds the detailsPrompt to the post v1 audiences params
func (o *PostV1AudiencesParams) SetDetailsPrompt(detailsPrompt []string) {
	o.DetailsPrompt = detailsPrompt
}

// WithDetailsQuestion adds the detailsQuestion to the post v1 audiences params
func (o *PostV1AudiencesParams) WithDetailsQuestion(detailsQuestion []string) *PostV1AudiencesParams {
	o.SetDetailsQuestion(detailsQuestion)
	return o
}

// SetDetailsQuestion adds the detailsQuestion to the post v1 audiences params
func (o *PostV1AudiencesParams) SetDetailsQuestion(detailsQuestion []string) {
	o.DetailsQuestion = detailsQuestion
}

// WithDetailsSlug adds the detailsSlug to the post v1 audiences params
func (o *PostV1AudiencesParams) WithDetailsSlug(detailsSlug []string) *PostV1AudiencesParams {
	o.SetDetailsSlug(detailsSlug)
	return o
}

// SetDetailsSlug adds the detailsSlug to the post v1 audiences params
func (o *PostV1AudiencesParams) SetDetailsSlug(detailsSlug []string) {
	o.DetailsSlug = detailsSlug
}

// WithName adds the name to the post v1 audiences params
func (o *PostV1AudiencesParams) WithName(name string) *PostV1AudiencesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post v1 audiences params
func (o *PostV1AudiencesParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1AudiencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Default != nil {

		// form param default
		var frDefault bool
		if o.Default != nil {
			frDefault = *o.Default
		}
		fDefault := swag.FormatBool(frDefault)
		if fDefault != "" {
			if err := r.SetFormParam("default", fDefault); err != nil {
				return err
			}
		}
	}

	// form param description
	frDescription := o.Description
	fDescription := frDescription
	if fDescription != "" {
		if err := r.SetFormParam("description", fDescription); err != nil {
			return err
		}
	}

	if o.DetailsPrompt != nil {

		// binding items for details[prompt]
		joinedDetailsPrompt := o.bindParamDetailsPrompt(reg)

		// form array param details[prompt]
		if err := r.SetFormParam("details[prompt]", joinedDetailsPrompt...); err != nil {
			return err
		}
	}

	if o.DetailsQuestion != nil {

		// binding items for details[question]
		joinedDetailsQuestion := o.bindParamDetailsQuestion(reg)

		// form array param details[question]
		if err := r.SetFormParam("details[question]", joinedDetailsQuestion...); err != nil {
			return err
		}
	}

	if o.DetailsSlug != nil {

		// binding items for details[slug]
		joinedDetailsSlug := o.bindParamDetailsSlug(reg)

		// form array param details[slug]
		if err := r.SetFormParam("details[slug]", joinedDetailsSlug...); err != nil {
			return err
		}
	}

	// form param name
	frName := o.Name
	fName := frName
	if fName != "" {
		if err := r.SetFormParam("name", fName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPostV1Audiences binds the parameter details[prompt]
func (o *PostV1AudiencesParams) bindParamDetailsPrompt(formats strfmt.Registry) []string {
	detailsPromptIR := o.DetailsPrompt

	var detailsPromptIC []string
	for _, detailsPromptIIR := range detailsPromptIR { // explode []string

		detailsPromptIIV := detailsPromptIIR // string as string
		detailsPromptIC = append(detailsPromptIC, detailsPromptIIV)
	}

	// items.CollectionFormat: ""
	detailsPromptIS := swag.JoinByFormat(detailsPromptIC, "")

	return detailsPromptIS
}

// bindParamPostV1Audiences binds the parameter details[question]
func (o *PostV1AudiencesParams) bindParamDetailsQuestion(formats strfmt.Registry) []string {
	detailsQuestionIR := o.DetailsQuestion

	var detailsQuestionIC []string
	for _, detailsQuestionIIR := range detailsQuestionIR { // explode []string

		detailsQuestionIIV := detailsQuestionIIR // string as string
		detailsQuestionIC = append(detailsQuestionIC, detailsQuestionIIV)
	}

	// items.CollectionFormat: ""
	detailsQuestionIS := swag.JoinByFormat(detailsQuestionIC, "")

	return detailsQuestionIS
}

// bindParamPostV1Audiences binds the parameter details[slug]
func (o *PostV1AudiencesParams) bindParamDetailsSlug(formats strfmt.Registry) []string {
	detailsSlugIR := o.DetailsSlug

	var detailsSlugIC []string
	for _, detailsSlugIIR := range detailsSlugIR { // explode []string

		detailsSlugIIV := detailsSlugIIR // string as string
		detailsSlugIC = append(detailsSlugIC, detailsSlugIIV)
	}

	// items.CollectionFormat: ""
	detailsSlugIS := swag.JoinByFormat(detailsSlugIC, "")

	return detailsSlugIS
}
