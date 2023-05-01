// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TicketingProjectFieldMapBodyEntity ticketing project field map body entity
//
// swagger:model Ticketing_ProjectFieldMapBodyEntity
type TicketingProjectFieldMapBodyEntity struct {

	// cases
	Cases []*TicketingProjectFieldMapCasesEntity `json:"cases"`

	// else
	Else *TicketingProjectFieldMapCasesElseEntity `json:"else,omitempty"`

	// external field
	ExternalField string `json:"external_field,omitempty"`

	// external value
	ExternalValue *TicketingProjectFieldMapExternalValueEntity `json:"external_value,omitempty"`

	// strategy
	// Enum: [basic logic]
	Strategy string `json:"strategy,omitempty"`

	// user data
	UserData interface{} `json:"user_data,omitempty"`
}

// Validate validates this ticketing project field map body entity
func (m *TicketingProjectFieldMapBodyEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketingProjectFieldMapBodyEntity) validateCases(formats strfmt.Registry) error {
	if swag.IsZero(m.Cases) { // not required
		return nil
	}

	for i := 0; i < len(m.Cases); i++ {
		if swag.IsZero(m.Cases[i]) { // not required
			continue
		}

		if m.Cases[i] != nil {
			if err := m.Cases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TicketingProjectFieldMapBodyEntity) validateElse(formats strfmt.Registry) error {
	if swag.IsZero(m.Else) { // not required
		return nil
	}

	if m.Else != nil {
		if err := m.Else.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("else")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("else")
			}
			return err
		}
	}

	return nil
}

func (m *TicketingProjectFieldMapBodyEntity) validateExternalValue(formats strfmt.Registry) error {
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

var ticketingProjectFieldMapBodyEntityTypeStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["basic","logic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ticketingProjectFieldMapBodyEntityTypeStrategyPropEnum = append(ticketingProjectFieldMapBodyEntityTypeStrategyPropEnum, v)
	}
}

const (

	// TicketingProjectFieldMapBodyEntityStrategyBasic captures enum value "basic"
	TicketingProjectFieldMapBodyEntityStrategyBasic string = "basic"

	// TicketingProjectFieldMapBodyEntityStrategyLogic captures enum value "logic"
	TicketingProjectFieldMapBodyEntityStrategyLogic string = "logic"
)

// prop value enum
func (m *TicketingProjectFieldMapBodyEntity) validateStrategyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ticketingProjectFieldMapBodyEntityTypeStrategyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TicketingProjectFieldMapBodyEntity) validateStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.Strategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrategyEnum("strategy", "body", m.Strategy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ticketing project field map body entity based on the context it is used
func (m *TicketingProjectFieldMapBodyEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketingProjectFieldMapBodyEntity) contextValidateCases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cases); i++ {

		if m.Cases[i] != nil {
			if err := m.Cases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TicketingProjectFieldMapBodyEntity) contextValidateElse(ctx context.Context, formats strfmt.Registry) error {

	if m.Else != nil {
		if err := m.Else.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("else")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("else")
			}
			return err
		}
	}

	return nil
}

func (m *TicketingProjectFieldMapBodyEntity) contextValidateExternalValue(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalValue != nil {
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
func (m *TicketingProjectFieldMapBodyEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketingProjectFieldMapBodyEntity) UnmarshalBinary(b []byte) error {
	var res TicketingProjectFieldMapBodyEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}