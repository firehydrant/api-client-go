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
)

// RunbooksElementDynamicSelectEntity runbooks element dynamic select entity
//
// swagger:model Runbooks_ElementDynamicSelectEntity
type RunbooksElementDynamicSelectEntity struct {

	// async url
	AsyncURL string `json:"async_url,omitempty"`

	// clearable
	Clearable bool `json:"clearable,omitempty"`

	// default value
	DefaultValue *RunbooksElementDynamicSelectEntitySelectOptionEntity `json:"default_value,omitempty"`

	// is multi
	IsMulti bool `json:"is_multi,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// options
	Options []*RunbooksElementDynamicSelectEntitySelectOptionEntity `json:"options"`

	// placeholder
	Placeholder string `json:"placeholder,omitempty"`

	// required
	Required bool `json:"required,omitempty"`
}

// Validate validates this runbooks element dynamic select entity
func (m *RunbooksElementDynamicSelectEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbooksElementDynamicSelectEntity) validateDefaultValue(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultValue) { // not required
		return nil
	}

	if m.DefaultValue != nil {
		if err := m.DefaultValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_value")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementDynamicSelectEntity) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {
		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {
			if err := m.Options[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this runbooks element dynamic select entity based on the context it is used
func (m *RunbooksElementDynamicSelectEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbooksElementDynamicSelectEntity) contextValidateDefaultValue(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultValue != nil {
		if err := m.DefaultValue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_value")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksElementDynamicSelectEntity) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Options); i++ {

		if m.Options[i] != nil {
			if err := m.Options[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunbooksElementDynamicSelectEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunbooksElementDynamicSelectEntity) UnmarshalBinary(b []byte) error {
	var res RunbooksElementDynamicSelectEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
