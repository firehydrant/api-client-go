// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1CustomFieldsDefinitions Create a new custom field definition
//
// swagger:model postV1CustomFieldsDefinitions
type PostV1CustomFieldsDefinitions struct {

	// description
	Description string `json:"description,omitempty"`

	// display name
	// Required: true
	DisplayName *string `json:"display_name"`

	// field type
	// Required: true
	FieldType *string `json:"field_type"`

	// permissible values
	PermissibleValues []string `json:"permissible_values"`

	// required
	// Required: true
	Required *bool `json:"required"`
}

// Validate validates this post v1 custom fields definitions
func (m *PostV1CustomFieldsDefinitions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1CustomFieldsDefinitions) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *PostV1CustomFieldsDefinitions) validateFieldType(formats strfmt.Registry) error {

	if err := validate.Required("field_type", "body", m.FieldType); err != nil {
		return err
	}

	return nil
}

func (m *PostV1CustomFieldsDefinitions) validateRequired(formats strfmt.Registry) error {

	if err := validate.Required("required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 custom fields definitions based on context it is used
func (m *PostV1CustomFieldsDefinitions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1CustomFieldsDefinitions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1CustomFieldsDefinitions) UnmarshalBinary(b []byte) error {
	var res PostV1CustomFieldsDefinitions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}