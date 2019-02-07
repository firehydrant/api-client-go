// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1ComponentsComponentIDEnvironments Add an environment to a component
// swagger:model postV1ComponentsComponentIdEnvironments
type PostV1ComponentsComponentIDEnvironments struct {

	// environment id
	// Required: true
	EnvironmentID *string `json:"environment_id"`
}

// Validate validates this post v1 components component Id environments
func (m *PostV1ComponentsComponentIDEnvironments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ComponentsComponentIDEnvironments) validateEnvironmentID(formats strfmt.Registry) error {

	if err := validate.Required("environment_id", "body", m.EnvironmentID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ComponentsComponentIDEnvironments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ComponentsComponentIDEnvironments) UnmarshalBinary(b []byte) error {
	var res PostV1ComponentsComponentIDEnvironments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
