// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaginationEntity pagination entity
//
// swagger:model PaginationEntity
type PaginationEntity struct {

	// count
	Count int32 `json:"count,omitempty"`

	// items
	Items int32 `json:"items,omitempty"`

	// last
	Last int32 `json:"last,omitempty"`

	// next
	Next *int32 `json:"next,omitempty"`

	// page
	Page int32 `json:"page,omitempty"`

	// pages
	Pages int32 `json:"pages,omitempty"`

	// prev
	Prev *int32 `json:"prev,omitempty"`
}

// Validate validates this pagination entity
func (m *PaginationEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pagination entity based on context it is used
func (m *PaginationEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PaginationEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginationEntity) UnmarshalBinary(b []byte) error {
	var res PaginationEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
