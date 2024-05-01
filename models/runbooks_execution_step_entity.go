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

// RunbooksExecutionStepEntity runbooks execution step entity
//
// swagger:model Runbooks_ExecutionStepEntity
type RunbooksExecutionStepEntity struct {

	// action slug
	ActionSlug string `json:"action_slug,omitempty"`

	// action type
	ActionType string `json:"action_type,omitempty"`

	// automatic
	Automatic bool `json:"automatic,omitempty"`

	// config
	Config interface{} `json:"config,omitempty"`

	// executable
	Executable bool `json:"executable,omitempty"`

	// execution
	Execution *RunbooksExecutionStepExecutionEntity `json:"execution,omitempty"`

	// has been rerun
	HasBeenRerun bool `json:"has_been_rerun,omitempty"`

	// has been retried
	HasBeenRetried bool `json:"has_been_retried,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// integration name
	IntegrationName string `json:"integration_name,omitempty"`

	// integration slug
	IntegrationSlug string `json:"integration_slug,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// repeatable
	Repeatable bool `json:"repeatable,omitempty"`

	// repeats
	Repeats bool `json:"repeats,omitempty"`

	// repeats at
	// Format: date-time
	RepeatsAt strfmt.DateTime `json:"repeats_at,omitempty"`

	// ISO8601 formatted duration string
	RepeatsDuration string `json:"repeats_duration,omitempty"`

	// rule
	Rule *RulesRuleEntity `json:"rule,omitempty"`

	// step elements
	StepElements []interface{} `json:"step_elements"`
}

// Validate validates this runbooks execution step entity
func (m *RunbooksExecutionStepEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepeatsAt(formats); err != nil {
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

func (m *RunbooksExecutionStepEntity) validateExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.Execution) { // not required
		return nil
	}

	if m.Execution != nil {
		if err := m.Execution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("execution")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksExecutionStepEntity) validateRepeatsAt(formats strfmt.Registry) error {
	if swag.IsZero(m.RepeatsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("repeats_at", "body", "date-time", m.RepeatsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunbooksExecutionStepEntity) validateRule(formats strfmt.Registry) error {
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

// ContextValidate validate this runbooks execution step entity based on the context it is used
func (m *RunbooksExecutionStepEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExecution(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbooksExecutionStepEntity) contextValidateExecution(ctx context.Context, formats strfmt.Registry) error {

	if m.Execution != nil {
		if err := m.Execution.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("execution")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("execution")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksExecutionStepEntity) contextValidateRule(ctx context.Context, formats strfmt.Registry) error {

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
func (m *RunbooksExecutionStepEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunbooksExecutionStepEntity) UnmarshalBinary(b []byte) error {
	var res RunbooksExecutionStepEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
