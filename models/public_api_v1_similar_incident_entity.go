// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicAPIV1SimilarIncidentEntity public API v1 similar incident entity
//
// swagger:model PublicAPI_V1_SimilarIncidentEntity
type PublicAPIV1SimilarIncidentEntity struct {

	// distance
	Distance float32 `json:"distance,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this public API v1 similar incident entity
func (m *PublicAPIV1SimilarIncidentEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this public API v1 similar incident entity based on context it is used
func (m *PublicAPIV1SimilarIncidentEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicAPIV1SimilarIncidentEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicAPIV1SimilarIncidentEntity) UnmarshalBinary(b []byte) error {
	var res PublicAPIV1SimilarIncidentEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
