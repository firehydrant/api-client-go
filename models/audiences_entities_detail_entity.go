// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AudiencesEntitiesDetailEntity audiences entities detail entity
//
// swagger:model Audiences_Entities_DetailEntity
type AudiencesEntitiesDetailEntity struct {

	// Unique identifier for the detail item
	ID string `json:"id,omitempty"`

	// Order position of this item in the list
	Position int32 `json:"position,omitempty"`

	// AI prompt used to gather this information
	Prompt string `json:"prompt,omitempty"`

	// The need-to-know question (maximum 255 characters)
	Question string `json:"question,omitempty"`

	// Slug of the detail, unique and autogenerated
	Slug string `json:"slug,omitempty"`
}

// Validate validates this audiences entities detail entity
func (m *AudiencesEntitiesDetailEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this audiences entities detail entity based on context it is used
func (m *AudiencesEntitiesDetailEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AudiencesEntitiesDetailEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AudiencesEntitiesDetailEntity) UnmarshalBinary(b []byte) error {
	var res AudiencesEntitiesDetailEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
