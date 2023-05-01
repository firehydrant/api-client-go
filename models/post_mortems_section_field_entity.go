// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostMortemsSectionFieldEntity PostMortems_SectionFieldEntity model
//
// swagger:model PostMortems_SectionFieldEntity
type PostMortemsSectionFieldEntity struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this post mortems section field entity
func (m *PostMortemsSectionFieldEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post mortems section field entity based on context it is used
func (m *PostMortemsSectionFieldEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostMortemsSectionFieldEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostMortemsSectionFieldEntity) UnmarshalBinary(b []byte) error {
	var res PostMortemsSectionFieldEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}