// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SignalsAPITargetEntity signals API target entity
//
// swagger:model Signals_API_TargetEntity
type SignalsAPITargetEntity struct {

	// id
	ID string `json:"id,omitempty"`

	// is pageable
	IsPageable bool `json:"is_pageable,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this signals API target entity
func (m *SignalsAPITargetEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this signals API target entity based on context it is used
func (m *SignalsAPITargetEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignalsAPITargetEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignalsAPITargetEntity) UnmarshalBinary(b []byte) error {
	var res SignalsAPITargetEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}