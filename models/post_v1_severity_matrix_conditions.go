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

// PostV1SeverityMatrixConditions Create a new condition
//
// swagger:model postV1SeverityMatrixConditions
type PostV1SeverityMatrixConditions struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// position
	Position int32 `json:"position,omitempty"`
}

// Validate validates this post v1 severity matrix conditions
func (m *PostV1SeverityMatrixConditions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SeverityMatrixConditions) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 severity matrix conditions based on context it is used
func (m *PostV1SeverityMatrixConditions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1SeverityMatrixConditions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SeverityMatrixConditions) UnmarshalBinary(b []byte) error {
	var res PostV1SeverityMatrixConditions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}