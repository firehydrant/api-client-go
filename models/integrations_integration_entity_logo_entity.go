// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntegrationsIntegrationEntityLogoEntity integrations integration entity logo entity
//
// swagger:model Integrations_IntegrationEntity_LogoEntity
type IntegrationsIntegrationEntityLogoEntity struct {

	// logo url
	LogoURL string `json:"logo_url,omitempty"`
}

// Validate validates this integrations integration entity logo entity
func (m *IntegrationsIntegrationEntityLogoEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this integrations integration entity logo entity based on context it is used
func (m *IntegrationsIntegrationEntityLogoEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsIntegrationEntityLogoEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsIntegrationEntityLogoEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsIntegrationEntityLogoEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}