// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntegrationsStatuspageConnectionEntity Integrations_Statuspage_ConnectionEntity model
//
// swagger:model Integrations_Statuspage_ConnectionEntity
type IntegrationsStatuspageConnectionEntity struct {

	// conditions
	Conditions *IntegrationsStatuspageConditionEntity `json:"conditions,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// page id
	PageID string `json:"page_id,omitempty"`

	// page name
	PageName string `json:"page_name,omitempty"`

	// severities
	Severities *IntegrationsStatuspageSeverityEntity `json:"severities,omitempty"`
}

// Validate validates this integrations statuspage connection entity
func (m *IntegrationsStatuspageConnectionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsStatuspageConnectionEntity) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationsStatuspageConnectionEntity) validateSeverities(formats strfmt.Registry) error {
	if swag.IsZero(m.Severities) { // not required
		return nil
	}

	if m.Severities != nil {
		if err := m.Severities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severities")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this integrations statuspage connection entity based on the context it is used
func (m *IntegrationsStatuspageConnectionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsStatuspageConnectionEntity) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	if m.Conditions != nil {

		if swag.IsZero(m.Conditions) { // not required
			return nil
		}

		if err := m.Conditions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationsStatuspageConnectionEntity) contextValidateSeverities(ctx context.Context, formats strfmt.Registry) error {

	if m.Severities != nil {

		if swag.IsZero(m.Severities) { // not required
			return nil
		}

		if err := m.Severities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severities")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsStatuspageConnectionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsStatuspageConnectionEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsStatuspageConnectionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
