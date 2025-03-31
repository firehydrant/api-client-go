// Code generated by go-swagger; DO NOT EDIT.

package status_pages

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

	// CompanyName.
	CompanyName *string

	// CompanyTosURL.
	CompanyTosURL *string

	// CompanyWebsite.
	CompanyWebsite *string

	// ComponentsInfrastructureID.
	ComponentsInfrastructureID []string

	// ComponentsInfrastructureType.
	ComponentsInfrastructureType []string

	/* ConditionsConditionID.

	   Severity matrix condition id
	*/
	ConditionsConditionID []string

	/* ConditionsNuncCondition.

	   Status page condition to map your severity matrix condition to
	*/
	ConditionsNuncCondition []string

	// EnableHistogram.
	EnableHistogram *bool

	// ExposedFields.
	ExposedFields []string

	// GreetingBody.
	GreetingBody *string

	// GreetingTitle.
	GreetingTitle *string

	// NuncConnectionID.
	NuncConnectionID string

	// OperationalMessage.
	OperationalMessage *string

	// PrimaryColor.
	PrimaryColor *string

	// SecondaryColor.
	SecondaryColor *string

	// Title.
	Title *string

	// UIVersion.
	//
	// Format: int32
	UIVersion *int32

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

// WithComponentsInfrastructureID adds the componentsInfrastructureID to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithComponentsInfrastructureID(componentsInfrastructureID []string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetComponentsInfrastructureID(componentsInfrastructureID)
	return o
}

// SetComponentsInfrastructureID adds the componentsInfrastructureId to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetComponentsInfrastructureID(componentsInfrastructureID []string) {
	o.ComponentsInfrastructureID = componentsInfrastructureID
}

// WithComponentsInfrastructureType adds the componentsInfrastructureType to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithComponentsInfrastructureType(componentsInfrastructureType []string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetComponentsInfrastructureType(componentsInfrastructureType)
	return o
}

// SetComponentsInfrastructureType adds the componentsInfrastructureType to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetComponentsInfrastructureType(componentsInfrastructureType []string) {
	o.ComponentsInfrastructureType = componentsInfrastructureType
}

// WithConditionsConditionID adds the conditionsConditionID to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithConditionsConditionID(conditionsConditionID []string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetConditionsConditionID(conditionsConditionID)
	return o
}

// SetConditionsConditionID adds the conditionsConditionId to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetConditionsConditionID(conditionsConditionID []string) {
	o.ConditionsConditionID = conditionsConditionID
}

// WithConditionsNuncCondition adds the conditionsNuncCondition to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithConditionsNuncCondition(conditionsNuncCondition []string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetConditionsNuncCondition(conditionsNuncCondition)
	return o
}

// SetConditionsNuncCondition adds the conditionsNuncCondition to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetConditionsNuncCondition(conditionsNuncCondition []string) {
	o.ConditionsNuncCondition = conditionsNuncCondition
}

// WithEnableHistogram adds the enableHistogram to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithEnableHistogram(enableHistogram *bool) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetEnableHistogram(enableHistogram)
	return o
}

// SetEnableHistogram adds the enableHistogram to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetEnableHistogram(enableHistogram *bool) {
	o.EnableHistogram = enableHistogram
}

// WithExposedFields adds the exposedFields to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithExposedFields(exposedFields []string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetExposedFields(exposedFields)
	return o
}

// SetExposedFields adds the exposedFields to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetExposedFields(exposedFields []string) {
	o.ExposedFields = exposedFields
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

// WithNuncConnectionID adds the nuncConnectionID to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithNuncConnectionID(nuncConnectionID string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetNuncConnectionID(nuncConnectionID)
	return o
}

// SetNuncConnectionID adds the nuncConnectionId to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetNuncConnectionID(nuncConnectionID string) {
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

// WithTitle adds the title to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithTitle(title *string) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetTitle(title)
	return o
}

// SetTitle adds the title to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetTitle(title *string) {
	o.Title = title
}

// WithUIVersion adds the uIVersion to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) WithUIVersion(uIVersion *int32) *PutV1NuncConnectionsNuncConnectionIDParams {
	o.SetUIVersion(uIVersion)
	return o
}

// SetUIVersion adds the uiVersion to the put v1 nunc connections nunc connection Id params
func (o *PutV1NuncConnectionsNuncConnectionIDParams) SetUIVersion(uIVersion *int32) {
	o.UIVersion = uIVersion
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

	if o.ComponentsInfrastructureID != nil {

		// binding items for components[infrastructure_id]
		joinedComponentsInfrastructureID := o.bindParamComponentsInfrastructureID(reg)

		// form array param components[infrastructure_id]
		if err := r.SetFormParam("components[infrastructure_id]", joinedComponentsInfrastructureID...); err != nil {
			return err
		}
	}

	if o.ComponentsInfrastructureType != nil {

		// binding items for components[infrastructure_type]
		joinedComponentsInfrastructureType := o.bindParamComponentsInfrastructureType(reg)

		// form array param components[infrastructure_type]
		if err := r.SetFormParam("components[infrastructure_type]", joinedComponentsInfrastructureType...); err != nil {
			return err
		}
	}

	if o.ConditionsConditionID != nil {

		// binding items for conditions[condition_id]
		joinedConditionsConditionID := o.bindParamConditionsConditionID(reg)

		// form array param conditions[condition_id]
		if err := r.SetFormParam("conditions[condition_id]", joinedConditionsConditionID...); err != nil {
			return err
		}
	}

	if o.ConditionsNuncCondition != nil {

		// binding items for conditions[nunc_condition]
		joinedConditionsNuncCondition := o.bindParamConditionsNuncCondition(reg)

		// form array param conditions[nunc_condition]
		if err := r.SetFormParam("conditions[nunc_condition]", joinedConditionsNuncCondition...); err != nil {
			return err
		}
	}

	if o.EnableHistogram != nil {

		// form param enable_histogram
		var frEnableHistogram bool
		if o.EnableHistogram != nil {
			frEnableHistogram = *o.EnableHistogram
		}
		fEnableHistogram := swag.FormatBool(frEnableHistogram)
		if fEnableHistogram != "" {
			if err := r.SetFormParam("enable_histogram", fEnableHistogram); err != nil {
				return err
			}
		}
	}

	if o.ExposedFields != nil {

		// binding items for exposed_fields
		joinedExposedFields := o.bindParamExposedFields(reg)

		// form array param exposed_fields
		if err := r.SetFormParam("exposed_fields", joinedExposedFields...); err != nil {
			return err
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

	// path param nunc_connection_id
	if err := r.SetPathParam("nunc_connection_id", o.NuncConnectionID); err != nil {
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

	if o.Title != nil {

		// form param title
		var frTitle string
		if o.Title != nil {
			frTitle = *o.Title
		}
		fTitle := frTitle
		if fTitle != "" {
			if err := r.SetFormParam("title", fTitle); err != nil {
				return err
			}
		}
	}

	if o.UIVersion != nil {

		// form param ui_version
		var frUIVersion int32
		if o.UIVersion != nil {
			frUIVersion = *o.UIVersion
		}
		fUIVersion := swag.FormatInt32(frUIVersion)
		if fUIVersion != "" {
			if err := r.SetFormParam("ui_version", fUIVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPutV1NuncConnectionsNuncConnectionID binds the parameter components[infrastructure_id]
func (o *PutV1NuncConnectionsNuncConnectionIDParams) bindParamComponentsInfrastructureID(formats strfmt.Registry) []string {
	componentsInfrastructureIDIR := o.ComponentsInfrastructureID

	var componentsInfrastructureIDIC []string
	for _, componentsInfrastructureIDIIR := range componentsInfrastructureIDIR { // explode []string

		componentsInfrastructureIDIIV := componentsInfrastructureIDIIR // string as string
		componentsInfrastructureIDIC = append(componentsInfrastructureIDIC, componentsInfrastructureIDIIV)
	}

	// items.CollectionFormat: ""
	componentsInfrastructureIDIS := swag.JoinByFormat(componentsInfrastructureIDIC, "")

	return componentsInfrastructureIDIS
}

// bindParamPutV1NuncConnectionsNuncConnectionID binds the parameter components[infrastructure_type]
func (o *PutV1NuncConnectionsNuncConnectionIDParams) bindParamComponentsInfrastructureType(formats strfmt.Registry) []string {
	componentsInfrastructureTypeIR := o.ComponentsInfrastructureType

	var componentsInfrastructureTypeIC []string
	for _, componentsInfrastructureTypeIIR := range componentsInfrastructureTypeIR { // explode []string

		componentsInfrastructureTypeIIV := componentsInfrastructureTypeIIR // string as string
		componentsInfrastructureTypeIC = append(componentsInfrastructureTypeIC, componentsInfrastructureTypeIIV)
	}

	// items.CollectionFormat: ""
	componentsInfrastructureTypeIS := swag.JoinByFormat(componentsInfrastructureTypeIC, "")

	return componentsInfrastructureTypeIS
}

// bindParamPutV1NuncConnectionsNuncConnectionID binds the parameter conditions[condition_id]
func (o *PutV1NuncConnectionsNuncConnectionIDParams) bindParamConditionsConditionID(formats strfmt.Registry) []string {
	conditionsConditionIDIR := o.ConditionsConditionID

	var conditionsConditionIDIC []string
	for _, conditionsConditionIDIIR := range conditionsConditionIDIR { // explode []string

		conditionsConditionIDIIV := conditionsConditionIDIIR // string as string
		conditionsConditionIDIC = append(conditionsConditionIDIC, conditionsConditionIDIIV)
	}

	// items.CollectionFormat: ""
	conditionsConditionIDIS := swag.JoinByFormat(conditionsConditionIDIC, "")

	return conditionsConditionIDIS
}

// bindParamPutV1NuncConnectionsNuncConnectionID binds the parameter conditions[nunc_condition]
func (o *PutV1NuncConnectionsNuncConnectionIDParams) bindParamConditionsNuncCondition(formats strfmt.Registry) []string {
	conditionsNuncConditionIR := o.ConditionsNuncCondition

	var conditionsNuncConditionIC []string
	for _, conditionsNuncConditionIIR := range conditionsNuncConditionIR { // explode []string

		conditionsNuncConditionIIV := conditionsNuncConditionIIR // string as string
		conditionsNuncConditionIC = append(conditionsNuncConditionIC, conditionsNuncConditionIIV)
	}

	// items.CollectionFormat: ""
	conditionsNuncConditionIS := swag.JoinByFormat(conditionsNuncConditionIC, "")

	return conditionsNuncConditionIS
}

// bindParamPutV1NuncConnectionsNuncConnectionID binds the parameter exposed_fields
func (o *PutV1NuncConnectionsNuncConnectionIDParams) bindParamExposedFields(formats strfmt.Registry) []string {
	exposedFieldsIR := o.ExposedFields

	var exposedFieldsIC []string
	for _, exposedFieldsIIR := range exposedFieldsIR { // explode []string

		exposedFieldsIIV := exposedFieldsIIR // string as string
		exposedFieldsIC = append(exposedFieldsIC, exposedFieldsIIV)
	}

	// items.CollectionFormat: ""
	exposedFieldsIS := swag.JoinByFormat(exposedFieldsIC, "")

	return exposedFieldsIS
}
