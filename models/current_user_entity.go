// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CurrentUserEntity CurrentUserEntity model
//
// swagger:model CurrentUserEntity
type CurrentUserEntity struct {

	// account id
	AccountID int32 `json:"account_id,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// organization name
	OrganizationName string `json:"organization_name,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// source
	Source string `json:"source,omitempty"`
}

// Validate validates this current user entity
func (m *CurrentUserEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this current user entity based on context it is used
func (m *CurrentUserEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CurrentUserEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrentUserEntity) UnmarshalBinary(b []byte) error {
	var res CurrentUserEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
