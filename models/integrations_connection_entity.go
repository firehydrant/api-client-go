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

// IntegrationsConnectionEntity Integrations_ConnectionEntity model
//
// swagger:model Integrations_ConnectionEntity
type IntegrationsConnectionEntity struct {

	// authorized by
	AuthorizedBy string `json:"authorized_by,omitempty"`

	// authorized by id
	AuthorizedByID string `json:"authorized_by_id,omitempty"`

	// configuration url
	ConfigurationURL string `json:"configuration_url,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// default authorized actor
	DefaultAuthorizedActor *AuthorEntity `json:"default_authorized_actor,omitempty"`

	// Integration-specific details of this connection. As identified by the integration_slug, this object will be represented by that integration's ConnectionEntity.
	Details interface{} `json:"details,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// integration id
	IntegrationID string `json:"integration_id,omitempty"`

	// integration slug
	IntegrationSlug string `json:"integration_slug,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this integrations connection entity
func (m *IntegrationsConnectionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultAuthorizedActor(formats); err != nil {
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

func (m *IntegrationsConnectionEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsConnectionEntity) validateDefaultAuthorizedActor(formats strfmt.Registry) error {
	if swag.IsZero(m.DefaultAuthorizedActor) { // not required
		return nil
	}

	if m.DefaultAuthorizedActor != nil {
		if err := m.DefaultAuthorizedActor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_authorized_actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_authorized_actor")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationsConnectionEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this integrations connection entity based on the context it is used
func (m *IntegrationsConnectionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultAuthorizedActor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsConnectionEntity) contextValidateDefaultAuthorizedActor(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultAuthorizedActor != nil {
		if err := m.DefaultAuthorizedActor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_authorized_actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_authorized_actor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsConnectionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsConnectionEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsConnectionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
