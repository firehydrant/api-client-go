// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IncidentTeamAssignmentsEntity Assign a team for this incident
//
// swagger:model IncidentTeamAssignmentsEntity
type IncidentTeamAssignmentsEntity struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this incident team assignments entity
func (m *IncidentTeamAssignmentsEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this incident team assignments entity based on context it is used
func (m *IncidentTeamAssignmentsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IncidentTeamAssignmentsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentTeamAssignmentsEntity) UnmarshalBinary(b []byte) error {
	var res IncidentTeamAssignmentsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
