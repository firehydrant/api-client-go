// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1IncidentsIncidentIDGenericChatMessages Create a new generic chat message on an incident timeline. These are independent of any specific chat provider.
//
// swagger:model postV1IncidentsIncidentIdGenericChatMessages
type PostV1IncidentsIncidentIDGenericChatMessages struct {

	// body
	// Required: true
	Body *string `json:"body"`

	// ISO8601 timestamp for when the chat message occurred
	// Format: date-time
	OccurredAt strfmt.DateTime `json:"occurred_at,omitempty"`

	// vote direction
	// Enum: [up down]
	VoteDirection string `json:"vote_direction,omitempty"`
}

// Validate validates this post v1 incidents incident Id generic chat messages
func (m *PostV1IncidentsIncidentIDGenericChatMessages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOccurredAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoteDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentsIncidentIDGenericChatMessages) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentsIncidentIDGenericChatMessages) validateOccurredAt(formats strfmt.Registry) error {
	if swag.IsZero(m.OccurredAt) { // not required
		return nil
	}

	if err := validate.FormatOf("occurred_at", "body", "date-time", m.OccurredAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var postV1IncidentsIncidentIdGenericChatMessagesTypeVoteDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1IncidentsIncidentIdGenericChatMessagesTypeVoteDirectionPropEnum = append(postV1IncidentsIncidentIdGenericChatMessagesTypeVoteDirectionPropEnum, v)
	}
}

const (

	// PostV1IncidentsIncidentIDGenericChatMessagesVoteDirectionUp captures enum value "up"
	PostV1IncidentsIncidentIDGenericChatMessagesVoteDirectionUp string = "up"

	// PostV1IncidentsIncidentIDGenericChatMessagesVoteDirectionDown captures enum value "down"
	PostV1IncidentsIncidentIDGenericChatMessagesVoteDirectionDown string = "down"
)

// prop value enum
func (m *PostV1IncidentsIncidentIDGenericChatMessages) validateVoteDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1IncidentsIncidentIdGenericChatMessagesTypeVoteDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1IncidentsIncidentIDGenericChatMessages) validateVoteDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.VoteDirection) { // not required
		return nil
	}

	// value enum
	if err := m.validateVoteDirectionEnum("vote_direction", "body", m.VoteDirection); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 incidents incident Id generic chat messages based on context it is used
func (m *PostV1IncidentsIncidentIDGenericChatMessages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDGenericChatMessages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentsIncidentIDGenericChatMessages) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentsIncidentIDGenericChatMessages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
