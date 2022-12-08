// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchV1IncidentsIncidentIDGenericChatMessages Update an existing generic chat message on an incident.
//
// swagger:model patchV1IncidentsIncidentIdGenericChatMessages
type PatchV1IncidentsIncidentIDGenericChatMessages struct {

	// body
	// Required: true
	Body *string `json:"body"`
}

// Validate validates this patch v1 incidents incident Id generic chat messages
func (m *PatchV1IncidentsIncidentIDGenericChatMessages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1IncidentsIncidentIDGenericChatMessages) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 incidents incident Id generic chat messages based on context it is used
func (m *PatchV1IncidentsIncidentIDGenericChatMessages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDGenericChatMessages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDGenericChatMessages) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentsIncidentIDGenericChatMessages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}