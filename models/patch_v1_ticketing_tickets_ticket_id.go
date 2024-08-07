// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchV1TicketingTicketsTicketID Update a ticket's attributes
//
// swagger:model patchV1TicketingTicketsTicketId
type PatchV1TicketingTicketsTicketID struct {

	// description
	Description string `json:"description,omitempty"`

	// priority id
	PriorityID string `json:"priority_id,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// List of tags for the ticket
	TagList []string `json:"tag_list"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this patch v1 ticketing tickets ticket Id
func (m *PatchV1TicketingTicketsTicketID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 ticketing tickets ticket Id based on context it is used
func (m *PatchV1TicketingTicketsTicketID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TicketingTicketsTicketID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TicketingTicketsTicketID) UnmarshalBinary(b []byte) error {
	var res PatchV1TicketingTicketsTicketID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
