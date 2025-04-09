// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntegrationsSlackWorkspaceEntity Integrations_Slack_WorkspaceEntity model
//
// swagger:model Integrations_Slack_WorkspaceEntity
type IntegrationsSlackWorkspaceEntity struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// team id
	TeamID string `json:"team_id,omitempty"`
}

// Validate validates this integrations slack workspace entity
func (m *IntegrationsSlackWorkspaceEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this integrations slack workspace entity based on context it is used
func (m *IntegrationsSlackWorkspaceEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsSlackWorkspaceEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsSlackWorkspaceEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsSlackWorkspaceEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
