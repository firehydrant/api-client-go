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

// StatusPageEntity Add a status page to an incident.
//
// swagger:model StatusPageEntity
type StatusPageEntity struct {

	// external id
	ExternalID string `json:"external_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// integration
	Integration *IntegrationEntity `json:"integration,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this status page entity
func (m *StatusPageEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusPageEntity) validateIntegration(formats strfmt.Registry) error {
	if swag.IsZero(m.Integration) { // not required
		return nil
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this status page entity based on the context it is used
func (m *StatusPageEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusPageEntity) contextValidateIntegration(ctx context.Context, formats strfmt.Registry) error {

	if m.Integration != nil {
		if err := m.Integration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusPageEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusPageEntity) UnmarshalBinary(b []byte) error {
	var res StatusPageEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
