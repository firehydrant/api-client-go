// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1IncidentTypes Create a new incident type
//
// swagger:model postV1IncidentTypes
type PostV1IncidentTypes struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// template
	// Required: true
	Template *PostV1IncidentTypesTemplate `json:"template"`
}

// Validate validates this post v1 incident types
func (m *PostV1IncidentTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentTypes) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentTypes) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post v1 incident types based on the context it is used
func (m *PostV1IncidentTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentTypes) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentTypes) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1IncidentTypesTemplate post v1 incident types template
//
// swagger:model PostV1IncidentTypesTemplate
type PostV1IncidentTypesTemplate struct {

	// customer impact summary
	CustomerImpactSummary string `json:"customer_impact_summary,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// An array of impact/condition combinations
	Impacts []*PostV1IncidentTypesTemplateImpactsItems0 `json:"impacts"`

	// A labels hash of keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// priority
	Priority string `json:"priority,omitempty"`

	// private incident
	PrivateIncident bool `json:"private_incident,omitempty"`

	// List of ids of Runbooks to attach to incidents created from this type
	RunbookIds []string `json:"runbook_ids"`

	// severity
	Severity string `json:"severity,omitempty"`

	// List of tags for the incident
	TagList []string `json:"tag_list"`

	// List of ids of teams to be assigned to incidents
	TeamIds []string `json:"team_ids"`
}

// Validate validates this post v1 incident types template
func (m *PostV1IncidentTypesTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImpacts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentTypesTemplate) validateImpacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Impacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Impacts); i++ {
		if swag.IsZero(m.Impacts[i]) { // not required
			continue
		}

		if m.Impacts[i] != nil {
			if err := m.Impacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("template" + "." + "impacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("template" + "." + "impacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post v1 incident types template based on the context it is used
func (m *PostV1IncidentTypesTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImpacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentTypesTemplate) contextValidateImpacts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Impacts); i++ {

		if m.Impacts[i] != nil {
			if err := m.Impacts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("template" + "." + "impacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("template" + "." + "impacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentTypesTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentTypesTemplate) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentTypesTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1IncidentTypesTemplateImpactsItems0 post v1 incident types template impacts items0
//
// swagger:model PostV1IncidentTypesTemplateImpactsItems0
type PostV1IncidentTypesTemplateImpactsItems0 struct {

	// The id of the condition
	// Required: true
	ConditionID *string `json:"condition_id"`

	// The id of impact
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this post v1 incident types template impacts items0
func (m *PostV1IncidentTypesTemplateImpactsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentTypesTemplateImpactsItems0) validateConditionID(formats strfmt.Registry) error {

	if err := validate.Required("condition_id", "body", m.ConditionID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentTypesTemplateImpactsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 incident types template impacts items0 based on context it is used
func (m *PostV1IncidentTypesTemplateImpactsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentTypesTemplateImpactsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentTypesTemplateImpactsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentTypesTemplateImpactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
