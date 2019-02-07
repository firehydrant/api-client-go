// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ChangeIdentityEntity Update an identity
// swagger:model ChangeIdentityEntity
type ChangeIdentityEntity struct {

	// change id
	ChangeID string `json:"change_id,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this change identity entity
func (m *ChangeIdentityEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangeIdentityEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeIdentityEntity) UnmarshalBinary(b []byte) error {
	var res ChangeIdentityEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
