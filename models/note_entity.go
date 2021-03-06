// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NoteEntity Update a note
// swagger:model NoteEntity
type NoteEntity struct {

	// body
	Body string `json:"body,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this note entity
func (m *NoteEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NoteEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NoteEntity) UnmarshalBinary(b []byte) error {
	var res NoteEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
