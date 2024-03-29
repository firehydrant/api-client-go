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

// PostV1StatusUpdateTemplates Create a status update template for your organization
//
// swagger:model postV1StatusUpdateTemplates
type PostV1StatusUpdateTemplates struct {

	// body
	// Required: true
	Body *string `json:"body"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post v1 status update templates
func (m *PostV1StatusUpdateTemplates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1StatusUpdateTemplates) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *PostV1StatusUpdateTemplates) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 status update templates based on context it is used
func (m *PostV1StatusUpdateTemplates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1StatusUpdateTemplates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1StatusUpdateTemplates) UnmarshalBinary(b []byte) error {
	var res PostV1StatusUpdateTemplates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
