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

// PostV1SignalsEmailTargets Create a Signals email target for a team.
//
// swagger:model postV1SignalsEmailTargets
type PostV1SignalsEmailTargets struct {

	// A list of email addresses that are allowed to send events to the target. Must be exact match.
	AllowedSenders []string `json:"allowed_senders"`

	// A detailed description of the email target.
	Description string `json:"description,omitempty"`

	// The CEL expression that defines the level of an incoming email that is sent to the target.
	LevelCel string `json:"level_cel,omitempty"`

	// The email target's name.
	// Required: true
	Name *string `json:"name"`

	// Whether or not all rules must match, or if only one rule must match.
	// Enum: [all any]
	RuleMatchingStrategy string `json:"rule_matching_strategy,omitempty"`

	// A list of CEL expressions that should be evaluated and matched to determine if the target should be notified.
	Rules []string `json:"rules"`

	// The email address that will be listening to events.
	Slug string `json:"slug,omitempty"`

	// The CEL expression that defines the status of an incoming email that is sent to the target.
	StatusCel string `json:"status_cel,omitempty"`

	// target
	Target *PostV1SignalsEmailTargetsTarget `json:"target,omitempty"`
}

// Validate validates this post v1 signals email targets
func (m *PostV1SignalsEmailTargets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleMatchingStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SignalsEmailTargets) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var postV1SignalsEmailTargetsTypeRuleMatchingStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1SignalsEmailTargetsTypeRuleMatchingStrategyPropEnum = append(postV1SignalsEmailTargetsTypeRuleMatchingStrategyPropEnum, v)
	}
}

const (

	// PostV1SignalsEmailTargetsRuleMatchingStrategyAll captures enum value "all"
	PostV1SignalsEmailTargetsRuleMatchingStrategyAll string = "all"

	// PostV1SignalsEmailTargetsRuleMatchingStrategyAny captures enum value "any"
	PostV1SignalsEmailTargetsRuleMatchingStrategyAny string = "any"
)

// prop value enum
func (m *PostV1SignalsEmailTargets) validateRuleMatchingStrategyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1SignalsEmailTargetsTypeRuleMatchingStrategyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1SignalsEmailTargets) validateRuleMatchingStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.RuleMatchingStrategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateRuleMatchingStrategyEnum("rule_matching_strategy", "body", m.RuleMatchingStrategy); err != nil {
		return err
	}

	return nil
}

func (m *PostV1SignalsEmailTargets) validateTarget(formats strfmt.Registry) error {
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

// ContextValidate validate this post v1 signals email targets based on the context it is used
func (m *PostV1SignalsEmailTargets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SignalsEmailTargets) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PostV1SignalsEmailTargets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SignalsEmailTargets) UnmarshalBinary(b []byte) error {
	var res PostV1SignalsEmailTargets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1SignalsEmailTargetsTarget The target that the email target will notify. This object must contain a `type`
// field that specifies the type of target and an `id` field that specifies the ID of
// the target. The `type` field must be one of "escalation_policy", "on_call_schedule",
// "team", "user", or "slack_channel".
//
// swagger:model PostV1SignalsEmailTargetsTarget
type PostV1SignalsEmailTargetsTarget struct {

	// The ID of the target that the inbound email will notify when matched.
	// Required: true
	ID *string `json:"id"`

	// The type of target that the inbound email will notify when matched.
	// Required: true
	// Enum: [Team EntireTeam EscalationPolicy OnCallSchedule User SlackChannel Webhook]
	Type *string `json:"type"`
}

// Validate validates this post v1 signals email targets target
func (m *PostV1SignalsEmailTargetsTarget) Validate(formats strfmt.Registry) error {
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

func (m *PostV1SignalsEmailTargetsTarget) validateID(formats strfmt.Registry) error {

	if err := validate.Required("target"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var postV1SignalsEmailTargetsTargetTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Team","EntireTeam","EscalationPolicy","OnCallSchedule","User","SlackChannel","Webhook"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1SignalsEmailTargetsTargetTypeTypePropEnum = append(postV1SignalsEmailTargetsTargetTypeTypePropEnum, v)
	}
}

const (

	// PostV1SignalsEmailTargetsTargetTypeTeam captures enum value "Team"
	PostV1SignalsEmailTargetsTargetTypeTeam string = "Team"

	// PostV1SignalsEmailTargetsTargetTypeEntireTeam captures enum value "EntireTeam"
	PostV1SignalsEmailTargetsTargetTypeEntireTeam string = "EntireTeam"

	// PostV1SignalsEmailTargetsTargetTypeEscalationPolicy captures enum value "EscalationPolicy"
	PostV1SignalsEmailTargetsTargetTypeEscalationPolicy string = "EscalationPolicy"

	// PostV1SignalsEmailTargetsTargetTypeOnCallSchedule captures enum value "OnCallSchedule"
	PostV1SignalsEmailTargetsTargetTypeOnCallSchedule string = "OnCallSchedule"

	// PostV1SignalsEmailTargetsTargetTypeUser captures enum value "User"
	PostV1SignalsEmailTargetsTargetTypeUser string = "User"

	// PostV1SignalsEmailTargetsTargetTypeSlackChannel captures enum value "SlackChannel"
	PostV1SignalsEmailTargetsTargetTypeSlackChannel string = "SlackChannel"

	// PostV1SignalsEmailTargetsTargetTypeWebhook captures enum value "Webhook"
	PostV1SignalsEmailTargetsTargetTypeWebhook string = "Webhook"
)

// prop value enum
func (m *PostV1SignalsEmailTargetsTarget) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1SignalsEmailTargetsTargetTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1SignalsEmailTargetsTarget) validateType(formats strfmt.Registry) error {

	if err := validate.Required("target"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("target"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 signals email targets target based on context it is used
func (m *PostV1SignalsEmailTargetsTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1SignalsEmailTargetsTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SignalsEmailTargetsTarget) UnmarshalBinary(b []byte) error {
	var res PostV1SignalsEmailTargetsTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
