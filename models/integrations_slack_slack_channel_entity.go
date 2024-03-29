// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntegrationsSlackSlackChannelEntity integrations slack slack channel entity
//
// swagger:model Integrations_Slack_SlackChannelEntity
type IntegrationsSlackSlackChannelEntity struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// slack channel id
	SlackChannelID string `json:"slack_channel_id,omitempty"`
}

// Validate validates this integrations slack slack channel entity
func (m *IntegrationsSlackSlackChannelEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this integrations slack slack channel entity based on context it is used
func (m *IntegrationsSlackSlackChannelEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsSlackSlackChannelEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsSlackSlackChannelEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsSlackSlackChannelEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
