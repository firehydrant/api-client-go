// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchV1TeamsTeamID Update a single team from its ID
//
// swagger:model patchV1TeamsTeamId
type PatchV1TeamsTeamID struct {

	// description
	Description string `json:"description,omitempty"`

	// memberships
	Memberships []*PatchV1TeamsTeamIDMembershipsItems0 `json:"memberships"`

	// ms teams channel
	MsTeamsChannel *PatchV1TeamsTeamIDMsTeamsChannel `json:"ms_teams_channel,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The Slack channel ID associated with this team. This may be the reference in FireHydrant's system (i.e. UUID) or the ID value from Slack (e.g. C1234567890).
	//
	SlackChannelID string `json:"slack_channel_id,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this patch v1 teams team Id
func (m *PatchV1TeamsTeamID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemberships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsTeamsChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1TeamsTeamID) validateMemberships(formats strfmt.Registry) error {
	if swag.IsZero(m.Memberships) { // not required
		return nil
	}

	for i := 0; i < len(m.Memberships); i++ {
		if swag.IsZero(m.Memberships[i]) { // not required
			continue
		}

		if m.Memberships[i] != nil {
			if err := m.Memberships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memberships" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memberships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1TeamsTeamID) validateMsTeamsChannel(formats strfmt.Registry) error {
	if swag.IsZero(m.MsTeamsChannel) { // not required
		return nil
	}

	if m.MsTeamsChannel != nil {
		if err := m.MsTeamsChannel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ms_teams_channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ms_teams_channel")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch v1 teams team Id based on the context it is used
func (m *PatchV1TeamsTeamID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMemberships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMsTeamsChannel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1TeamsTeamID) contextValidateMemberships(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Memberships); i++ {

		if m.Memberships[i] != nil {
			if err := m.Memberships[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("memberships" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("memberships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1TeamsTeamID) contextValidateMsTeamsChannel(ctx context.Context, formats strfmt.Registry) error {

	if m.MsTeamsChannel != nil {
		if err := m.MsTeamsChannel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ms_teams_channel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ms_teams_channel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TeamsTeamID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TeamsTeamID) UnmarshalBinary(b []byte) error {
	var res PatchV1TeamsTeamID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1TeamsTeamIDMembershipsItems0 patch v1 teams team ID memberships items0
//
// swagger:model PatchV1TeamsTeamIDMembershipsItems0
type PatchV1TeamsTeamIDMembershipsItems0 struct {

	// An incident role ID that this user will automatically assigned if this team is assigned to an incident
	IncidentRoleID string `json:"incident_role_id,omitempty"`

	// schedule id
	ScheduleID string `json:"schedule_id,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this patch v1 teams team ID memberships items0
func (m *PatchV1TeamsTeamIDMembershipsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 teams team ID memberships items0 based on context it is used
func (m *PatchV1TeamsTeamIDMembershipsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TeamsTeamIDMembershipsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TeamsTeamIDMembershipsItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1TeamsTeamIDMembershipsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1TeamsTeamIDMsTeamsChannel MS Teams channel identity for channel associated with this team
//
// swagger:model PatchV1TeamsTeamIDMsTeamsChannel
type PatchV1TeamsTeamIDMsTeamsChannel struct {

	// channel id
	// Required: true
	ChannelID *string `json:"channel_id"`

	// ms team id
	// Required: true
	MsTeamID *string `json:"ms_team_id"`
}

// Validate validates this patch v1 teams team ID ms teams channel
func (m *PatchV1TeamsTeamIDMsTeamsChannel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMsTeamID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1TeamsTeamIDMsTeamsChannel) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("ms_teams_channel"+"."+"channel_id", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1TeamsTeamIDMsTeamsChannel) validateMsTeamID(formats strfmt.Registry) error {

	if err := validate.Required("ms_teams_channel"+"."+"ms_team_id", "body", m.MsTeamID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 teams team ID ms teams channel based on context it is used
func (m *PatchV1TeamsTeamIDMsTeamsChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TeamsTeamIDMsTeamsChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TeamsTeamIDMsTeamsChannel) UnmarshalBinary(b []byte) error {
	var res PatchV1TeamsTeamIDMsTeamsChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
