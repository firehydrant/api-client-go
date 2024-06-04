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

// PatchV1SignalsEmailTargetsID Update a Signals email target by ID
//
// swagger:model patchV1SignalsEmailTargetsId
type PatchV1SignalsEmailTargetsID struct {

	// A detailed description of the email target.
	Description string `json:"description,omitempty"`

	// The CEL expression that defines the level of an incoming email that is sent to the target.
	LevelCel string `json:"level_cel,omitempty"`

	// The email target's name.
	Name string `json:"name,omitempty"`

	// The email address that will be listening to events.
	Slug string `json:"slug,omitempty"`

	// The CEL expression that defines the status of an incoming email that is sent to the target.
	StatusCel string `json:"status_cel,omitempty"`

	// target
	Target *PatchV1SignalsEmailTargetsIDTarget `json:"target,omitempty"`
}

// Validate validates this patch v1 signals email targets Id
func (m *PatchV1SignalsEmailTargetsID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1SignalsEmailTargetsID) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch v1 signals email targets Id based on the context it is used
func (m *PatchV1SignalsEmailTargetsID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1SignalsEmailTargetsID) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {
		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1SignalsEmailTargetsID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1SignalsEmailTargetsID) UnmarshalBinary(b []byte) error {
	var res PatchV1SignalsEmailTargetsID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1SignalsEmailTargetsIDTarget The target that the email target will notify. This object must contain a `type`
// field that specifies the type of target and an `id` field that specifies the ID of
// the target. The `type` field must be one of "escalation_policy", "on_call_schedule",
// "team", "user", or "slack_channel".
//
// swagger:model PatchV1SignalsEmailTargetsIDTarget
type PatchV1SignalsEmailTargetsIDTarget struct {

	// The ID of the target that the inbound email will notify when matched.
	// Required: true
	ID *string `json:"id"`

	// The type of target that the inbound email will notify when matched.
	// Required: true
	// Enum: [EscalationPolicy OnCallSchedule Team User SlackChannel]
	Type *string `json:"type"`
}

// Validate validates this patch v1 signals email targets ID target
func (m *PatchV1SignalsEmailTargetsIDTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1SignalsEmailTargetsIDTarget) validateID(formats strfmt.Registry) error {

	if err := validate.Required("target"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var patchV1SignalsEmailTargetsIdTargetTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EscalationPolicy","OnCallSchedule","Team","User","SlackChannel"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchV1SignalsEmailTargetsIdTargetTypeTypePropEnum = append(patchV1SignalsEmailTargetsIdTargetTypeTypePropEnum, v)
	}
}

const (

	// PatchV1SignalsEmailTargetsIDTargetTypeEscalationPolicy captures enum value "EscalationPolicy"
	PatchV1SignalsEmailTargetsIDTargetTypeEscalationPolicy string = "EscalationPolicy"

	// PatchV1SignalsEmailTargetsIDTargetTypeOnCallSchedule captures enum value "OnCallSchedule"
	PatchV1SignalsEmailTargetsIDTargetTypeOnCallSchedule string = "OnCallSchedule"

	// PatchV1SignalsEmailTargetsIDTargetTypeTeam captures enum value "Team"
	PatchV1SignalsEmailTargetsIDTargetTypeTeam string = "Team"

	// PatchV1SignalsEmailTargetsIDTargetTypeUser captures enum value "User"
	PatchV1SignalsEmailTargetsIDTargetTypeUser string = "User"

	// PatchV1SignalsEmailTargetsIDTargetTypeSlackChannel captures enum value "SlackChannel"
	PatchV1SignalsEmailTargetsIDTargetTypeSlackChannel string = "SlackChannel"
)

// prop value enum
func (m *PatchV1SignalsEmailTargetsIDTarget) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchV1SignalsEmailTargetsIdTargetTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchV1SignalsEmailTargetsIDTarget) validateType(formats strfmt.Registry) error {

	if err := validate.Required("target"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("target"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 signals email targets ID target based on context it is used
func (m *PatchV1SignalsEmailTargetsIDTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1SignalsEmailTargetsIDTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1SignalsEmailTargetsIDTarget) UnmarshalBinary(b []byte) error {
	var res PatchV1SignalsEmailTargetsIDTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}