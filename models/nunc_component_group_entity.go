// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NuncComponentGroupEntity nunc component group entity
//
// swagger:model NuncComponentGroupEntity
type NuncComponentGroupEntity struct {

	// component group id
	ComponentGroupID string `json:"component_group_id,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`
}

// Validate validates this nunc component group entity
func (m *NuncComponentGroupEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this nunc component group entity based on context it is used
func (m *NuncComponentGroupEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NuncComponentGroupEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NuncComponentGroupEntity) UnmarshalBinary(b []byte) error {
	var res NuncComponentGroupEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
