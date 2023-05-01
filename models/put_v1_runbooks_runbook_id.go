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

// PutV1RunbooksRunbookID Update a runbook and any attachment rules associated with it. This endpoint is used to configure nearly everything
// about a runbook, including but not limited to the steps, environments, attachment rules, and severities.
//
// swagger:model putV1RunbooksRunbookId
type PutV1RunbooksRunbookID struct {

	// attachment rule
	AttachmentRule *PutV1RunbooksRunbookIDAttachmentRule `json:"attachment_rule,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// environments
	Environments []*PutV1RunbooksRunbookIDEnvironmentsItems0 `json:"environments"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner *PutV1RunbooksRunbookIDOwner `json:"owner,omitempty"`

	// services
	Services []*PutV1RunbooksRunbookIDServicesItems0 `json:"services"`

	// severities
	Severities []*PutV1RunbooksRunbookIDSeveritiesItems0 `json:"severities"`

	// steps
	Steps []*PutV1RunbooksRunbookIDStepsItems0 `json:"steps"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this put v1 runbooks runbook Id
func (m *PutV1RunbooksRunbookID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1RunbooksRunbookID) validateAttachmentRule(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachmentRule) { // not required
		return nil
	}

	if m.AttachmentRule != nil {
		if err := m.AttachmentRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment_rule")
			}
			return err
		}
	}

	return nil
}

func (m *PutV1RunbooksRunbookID) validateEnvironments(formats strfmt.Registry) error {
	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1RunbooksRunbookID) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *PutV1RunbooksRunbookID) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1RunbooksRunbookID) validateSeverities(formats strfmt.Registry) error {
	if swag.IsZero(m.Severities) { // not required
		return nil
	}

	for i := 0; i < len(m.Severities); i++ {
		if swag.IsZero(m.Severities[i]) { // not required
			continue
		}

		if m.Severities[i] != nil {
			if err := m.Severities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("severities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("severities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1RunbooksRunbookID) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this put v1 runbooks runbook Id based on the context it is used
func (m *PutV1RunbooksRunbookID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachmentRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1RunbooksRunbookID) contextValidateAttachmentRule(ctx context.Context, formats strfmt.Registry) error {

	if m.AttachmentRule != nil {
		if err := m.AttachmentRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment_rule")
			}
			return err
		}
	}

	return nil
}

func (m *PutV1RunbooksRunbookID) contextValidateEnvironments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Environments); i++ {

		if m.Environments[i] != nil {
			if err := m.Environments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1RunbooksRunbookID) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *PutV1RunbooksRunbookID) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Services); i++ {

		if m.Services[i] != nil {
			if err := m.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1RunbooksRunbookID) contextValidateSeverities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Severities); i++ {

		if m.Severities[i] != nil {
			if err := m.Severities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("severities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("severities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1RunbooksRunbookID) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {
			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookID) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1RunbooksRunbookIDAttachmentRule put v1 runbooks runbook ID attachment rule
//
// swagger:model PutV1RunbooksRunbookIDAttachmentRule
type PutV1RunbooksRunbookIDAttachmentRule struct {

	// The JSON logic for the attaching the runbook
	// Required: true
	Logic *string `json:"logic"`

	// The user data for the rule
	UserData string `json:"user_data,omitempty"`
}

// Validate validates this put v1 runbooks runbook ID attachment rule
func (m *PutV1RunbooksRunbookIDAttachmentRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1RunbooksRunbookIDAttachmentRule) validateLogic(formats strfmt.Registry) error {

	if err := validate.Required("attachment_rule"+"."+"logic", "body", m.Logic); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 runbooks runbook ID attachment rule based on context it is used
func (m *PutV1RunbooksRunbookIDAttachmentRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDAttachmentRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDAttachmentRule) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookIDAttachmentRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1RunbooksRunbookIDEnvironmentsItems0 put v1 runbooks runbook ID environments items0
//
// swagger:model PutV1RunbooksRunbookIDEnvironmentsItems0
type PutV1RunbooksRunbookIDEnvironmentsItems0 struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this put v1 runbooks runbook ID environments items0
func (m *PutV1RunbooksRunbookIDEnvironmentsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put v1 runbooks runbook ID environments items0 based on context it is used
func (m *PutV1RunbooksRunbookIDEnvironmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDEnvironmentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDEnvironmentsItems0) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookIDEnvironmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1RunbooksRunbookIDOwner An object representing a Team that owns the runbook
//
// swagger:model PutV1RunbooksRunbookIDOwner
type PutV1RunbooksRunbookIDOwner struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this put v1 runbooks runbook ID owner
func (m *PutV1RunbooksRunbookIDOwner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put v1 runbooks runbook ID owner based on context it is used
func (m *PutV1RunbooksRunbookIDOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDOwner) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookIDOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1RunbooksRunbookIDServicesItems0 put v1 runbooks runbook ID services items0
//
// swagger:model PutV1RunbooksRunbookIDServicesItems0
type PutV1RunbooksRunbookIDServicesItems0 struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this put v1 runbooks runbook ID services items0
func (m *PutV1RunbooksRunbookIDServicesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put v1 runbooks runbook ID services items0 based on context it is used
func (m *PutV1RunbooksRunbookIDServicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDServicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDServicesItems0) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookIDServicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1RunbooksRunbookIDSeveritiesItems0 put v1 runbooks runbook ID severities items0
//
// swagger:model PutV1RunbooksRunbookIDSeveritiesItems0
type PutV1RunbooksRunbookIDSeveritiesItems0 struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this put v1 runbooks runbook ID severities items0
func (m *PutV1RunbooksRunbookIDSeveritiesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put v1 runbooks runbook ID severities items0 based on context it is used
func (m *PutV1RunbooksRunbookIDSeveritiesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDSeveritiesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDSeveritiesItems0) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookIDSeveritiesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1RunbooksRunbookIDStepsItems0 put v1 runbooks runbook ID steps items0
//
// swagger:model PutV1RunbooksRunbookIDStepsItems0
type PutV1RunbooksRunbookIDStepsItems0 struct {

	// ID of action to use for this step.
	// Required: true
	ActionID *string `json:"action_id"`

	// Name for step
	// Required: true
	Name *string `json:"name"`

	// rule
	Rule *PutV1RunbooksRunbookIDStepsItems0Rule `json:"rule,omitempty"`

	// ID of step to be updated
	StepID string `json:"step_id,omitempty"`
}

// Validate validates this put v1 runbooks runbook ID steps items0
func (m *PutV1RunbooksRunbookIDStepsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1RunbooksRunbookIDStepsItems0) validateActionID(formats strfmt.Registry) error {

	if err := validate.Required("action_id", "body", m.ActionID); err != nil {
		return err
	}

	return nil
}

func (m *PutV1RunbooksRunbookIDStepsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PutV1RunbooksRunbookIDStepsItems0) validateRule(formats strfmt.Registry) error {
	if swag.IsZero(m.Rule) { // not required
		return nil
	}

	if m.Rule != nil {
		if err := m.Rule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put v1 runbooks runbook ID steps items0 based on the context it is used
func (m *PutV1RunbooksRunbookIDStepsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1RunbooksRunbookIDStepsItems0) contextValidateRule(ctx context.Context, formats strfmt.Registry) error {

	if m.Rule != nil {
		if err := m.Rule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDStepsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDStepsItems0) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookIDStepsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1RunbooksRunbookIDStepsItems0Rule put v1 runbooks runbook ID steps items0 rule
//
// swagger:model PutV1RunbooksRunbookIDStepsItems0Rule
type PutV1RunbooksRunbookIDStepsItems0Rule struct {

	// The JSON logic for the rule
	// Required: true
	Logic *string `json:"logic"`

	// The user data for the rule
	UserData string `json:"user_data,omitempty"`
}

// Validate validates this put v1 runbooks runbook ID steps items0 rule
func (m *PutV1RunbooksRunbookIDStepsItems0Rule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1RunbooksRunbookIDStepsItems0Rule) validateLogic(formats strfmt.Registry) error {

	if err := validate.Required("rule"+"."+"logic", "body", m.Logic); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 runbooks runbook ID steps items0 rule based on context it is used
func (m *PutV1RunbooksRunbookIDStepsItems0Rule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDStepsItems0Rule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1RunbooksRunbookIDStepsItems0Rule) UnmarshalBinary(b []byte) error {
	var res PutV1RunbooksRunbookIDStepsItems0Rule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}