// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthorEntity AuthorEntity model
//
// swagger:model AuthorEntity
type AuthorEntity struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// source
	Source string `json:"source,omitempty"`
}

// Validate validates this author entity
func (m *AuthorEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this author entity based on context it is used
func (m *AuthorEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthorEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorEntity) UnmarshalBinary(b []byte) error {
	var res AuthorEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
