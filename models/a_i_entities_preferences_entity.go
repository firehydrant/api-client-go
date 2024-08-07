// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AIEntitiesPreferencesEntity AI_Entities_PreferencesEntity model
//
// swagger:model AI_Entities_PreferencesEntity
type AIEntitiesPreferencesEntity struct {

	// ai
	Ai bool `json:"ai,omitempty"`

	// description
	Description bool `json:"description,omitempty"`

	// followups
	Followups bool `json:"followups,omitempty"`

	// impact
	Impact bool `json:"impact,omitempty"`

	// retros
	Retros bool `json:"retros,omitempty"`

	// similar incidents
	SimilarIncidents bool `json:"similar_incidents,omitempty"`

	// summaries
	Summaries bool `json:"summaries,omitempty"`

	// updates
	Updates bool `json:"updates,omitempty"`
}

// Validate validates this a i entities preferences entity
func (m *AIEntitiesPreferencesEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a i entities preferences entity based on context it is used
func (m *AIEntitiesPreferencesEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AIEntitiesPreferencesEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AIEntitiesPreferencesEntity) UnmarshalBinary(b []byte) error {
	var res AIEntitiesPreferencesEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
