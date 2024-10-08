// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchV1SignalsWebhookTargetsID Update a Signals webhook target by ID
//
// swagger:model patchV1SignalsWebhookTargetsId
type PatchV1SignalsWebhookTargetsID struct {

	// An optional detailed description of the webhook target.
	Description string `json:"description,omitempty"`

	// The webhook target's name.
	Name string `json:"name,omitempty"`

	// An optional secret we will provide in the `FH-Signature` header
	// when sending a payload to the webhook target. This key will not be
	// shown in any response once configured.
	//
	SigningKey string `json:"signing_key,omitempty"`

	// The URL that the webhook target will notify.
	URL string `json:"url,omitempty"`
}

// Validate validates this patch v1 signals webhook targets Id
func (m *PatchV1SignalsWebhookTargetsID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 signals webhook targets Id based on context it is used
func (m *PatchV1SignalsWebhookTargetsID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1SignalsWebhookTargetsID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1SignalsWebhookTargetsID) UnmarshalBinary(b []byte) error {
	var res PatchV1SignalsWebhookTargetsID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
