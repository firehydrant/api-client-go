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

// PostV1TicketingTickets Creates a ticket for a project
//
// swagger:model postV1TicketingTickets
type PostV1TicketingTickets struct {

	// description
	Description string `json:"description,omitempty"`

	// priority id
	PriorityID string `json:"priority_id,omitempty"`

	// project id
	ProjectID string `json:"project_id,omitempty"`

	// The remote URL for an existing ticket in a supported and configured ticketing integration
	RemoteURL string `json:"remote_url,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// summary
	// Required: true
	Summary *string `json:"summary"`

	// List of tags for the ticket
	TagList []string `json:"tag_list"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this post v1 ticketing tickets
func (m *PostV1TicketingTickets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1TicketingTickets) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 ticketing tickets based on context it is used
func (m *PostV1TicketingTickets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1TicketingTickets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1TicketingTickets) UnmarshalBinary(b []byte) error {
	var res PostV1TicketingTickets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
