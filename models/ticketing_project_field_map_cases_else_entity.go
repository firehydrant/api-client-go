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

// TicketingProjectFieldMapCasesElseEntity ticketing project field map cases else entity
//
// swagger:model Ticketing_ProjectFieldMapCasesElseEntity
type TicketingProjectFieldMapCasesElseEntity struct {

	// external value
	ExternalValue *TicketingProjectFieldMapExternalValueEntity `json:"external_value,omitempty"`
}

// Validate validates this ticketing project field map cases else entity
func (m *TicketingProjectFieldMapCasesElseEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketingProjectFieldMapCasesElseEntity) validateExternalValue(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalValue) { // not required
		return nil
	}

	if m.ExternalValue != nil {
		if err := m.ExternalValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ticketing project field map cases else entity based on the context it is used
func (m *TicketingProjectFieldMapCasesElseEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketingProjectFieldMapCasesElseEntity) contextValidateExternalValue(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalValue != nil {

		if swag.IsZero(m.ExternalValue) { // not required
			return nil
		}

		if err := m.ExternalValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("external_value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("external_value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TicketingProjectFieldMapCasesElseEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketingProjectFieldMapCasesElseEntity) UnmarshalBinary(b []byte) error {
	var res TicketingProjectFieldMapCasesElseEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
