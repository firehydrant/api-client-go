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

// ServiceDependencyEntity ServiceDependencyEntity model
//
// swagger:model ServiceDependencyEntity
type ServiceDependencyEntity struct {

	// connected service
	ConnectedService *ServiceEntity `json:"connected_service,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// notes
	Notes string `json:"notes,omitempty"`

	// service
	Service *ServiceEntity `json:"service,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this service dependency entity
func (m *ServiceDependencyEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDependencyEntity) validateConnectedService(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedService) { // not required
		return nil
	}

	if m.ConnectedService != nil {
		if err := m.ConnectedService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connected_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connected_service")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDependencyEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDependencyEntity) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDependencyEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service dependency entity based on the context it is used
func (m *ServiceDependencyEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectedService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDependencyEntity) contextValidateConnectedService(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectedService != nil {

		if swag.IsZero(m.ConnectedService) { // not required
			return nil
		}

		if err := m.ConnectedService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connected_service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connected_service")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDependencyEntity) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {

		if swag.IsZero(m.Service) { // not required
			return nil
		}

		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDependencyEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDependencyEntity) UnmarshalBinary(b []byte) error {
	var res ServiceDependencyEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
