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

// IntegrationsConnectionStatusEntity Integrations_ConnectionStatusEntity model
//
// swagger:model Integrations_ConnectionStatusEntity
type IntegrationsConnectionStatusEntity struct {

	// check type
	CheckType string `json:"check_type,omitempty"`

	// checked at
	// Format: date-time
	CheckedAt strfmt.DateTime `json:"checked_at,omitempty"`

	// connection id
	ConnectionID string `json:"connection_id,omitempty"`

	// Additional unstructured data about the status check.
	Data interface{} `json:"data,omitempty"`

	// error type
	ErrorType string `json:"error_type,omitempty"`

	// integration slug
	IntegrationSlug string `json:"integration_slug,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// performed by
	PerformedBy *AuthorEntity `json:"performed_by,omitempty"`

	// status
	// Enum: [ok info warning error]
	Status string `json:"status,omitempty"`
}

// Validate validates this integrations connection status entity
func (m *IntegrationsConnectionStatusEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsConnectionStatusEntity) validateCheckedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CheckedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("checked_at", "body", "date-time", m.CheckedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsConnectionStatusEntity) validatePerformedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformedBy) { // not required
		return nil
	}

	if m.PerformedBy != nil {
		if err := m.PerformedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performed_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("performed_by")
			}
			return err
		}
	}

	return nil
}

var integrationsConnectionStatusEntityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","info","warning","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationsConnectionStatusEntityTypeStatusPropEnum = append(integrationsConnectionStatusEntityTypeStatusPropEnum, v)
	}
}

const (

	// IntegrationsConnectionStatusEntityStatusOk captures enum value "ok"
	IntegrationsConnectionStatusEntityStatusOk string = "ok"

	// IntegrationsConnectionStatusEntityStatusInfo captures enum value "info"
	IntegrationsConnectionStatusEntityStatusInfo string = "info"

	// IntegrationsConnectionStatusEntityStatusWarning captures enum value "warning"
	IntegrationsConnectionStatusEntityStatusWarning string = "warning"

	// IntegrationsConnectionStatusEntityStatusError captures enum value "error"
	IntegrationsConnectionStatusEntityStatusError string = "error"
)

// prop value enum
func (m *IntegrationsConnectionStatusEntity) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, integrationsConnectionStatusEntityTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationsConnectionStatusEntity) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this integrations connection status entity based on the context it is used
func (m *IntegrationsConnectionStatusEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePerformedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsConnectionStatusEntity) contextValidatePerformedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.PerformedBy != nil {
		if err := m.PerformedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performed_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("performed_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsConnectionStatusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsConnectionStatusEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsConnectionStatusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
