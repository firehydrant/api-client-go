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

// RunbooksActionsEntity runbooks actions entity
//
// swagger:model Runbooks_ActionsEntity
type RunbooksActionsEntity struct {

	// automatable
	Automatable bool `json:"automatable,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// config
	Config *RunbooksActionConfigEntity `json:"config,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// default logic
	DefaultLogic interface{} `json:"default_logic,omitempty"`

	// default rule data
	DefaultRuleData interface{} `json:"default_rule_data,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// integration
	Integration *IntegrationsIntegrationEntity `json:"integration,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// prerequisites
	Prerequisites interface{} `json:"prerequisites,omitempty"`

	// repeatable
	Repeatable bool `json:"repeatable,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// supported runbook types
	SupportedRunbookTypes []string `json:"supported_runbook_types"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this runbooks actions entity
func (m *RunbooksActionsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbooksActionsEntity) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksActionsEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunbooksActionsEntity) validateIntegration(formats strfmt.Registry) error {
	if swag.IsZero(m.Integration) { // not required
		return nil
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksActionsEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this runbooks actions entity based on the context it is used
func (m *RunbooksActionsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntegration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbooksActionsEntity) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {

		if swag.IsZero(m.Config) { // not required
			return nil
		}

		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RunbooksActionsEntity) contextValidateIntegration(ctx context.Context, formats strfmt.Registry) error {

	if m.Integration != nil {

		if swag.IsZero(m.Integration) { // not required
			return nil
		}

		if err := m.Integration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunbooksActionsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunbooksActionsEntity) UnmarshalBinary(b []byte) error {
	var res RunbooksActionsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
