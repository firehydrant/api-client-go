// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ParticipantEntity participant entity
//
// swagger:model ParticipantEntity
type ParticipantEntity struct {

	// actor
	Actor string `json:"actor,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this participant entity
func (m *ParticipantEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this participant entity based on context it is used
func (m *ParticipantEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParticipantEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParticipantEntity) UnmarshalBinary(b []byte) error {
	var res ParticipantEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
