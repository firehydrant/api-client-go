// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationsCustomFieldDefinitionEntity Organizations_CustomFieldDefinitionEntity model
//
// swagger:model Organizations_CustomFieldDefinitionEntity
type OrganizationsCustomFieldDefinitionEntity struct {

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"display_name,omitempty"`

	// field id
	FieldID string `json:"field_id,omitempty"`

	// field type
	FieldType string `json:"field_type,omitempty"`

	// permissible values
	PermissibleValues string `json:"permissible_values,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this organizations custom field definition entity
func (m *OrganizationsCustomFieldDefinitionEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organizations custom field definition entity based on context it is used
func (m *OrganizationsCustomFieldDefinitionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationsCustomFieldDefinitionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationsCustomFieldDefinitionEntity) UnmarshalBinary(b []byte) error {
	var res OrganizationsCustomFieldDefinitionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}