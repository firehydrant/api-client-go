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

// AlertsProcessingLogEntryEntity alerts processing log entry entity
//
// swagger:model Alerts_ProcessingLogEntryEntity
type AlertsProcessingLogEntryEntity struct {

	// An unstructured representation of this log entry's context.
	Context interface{} `json:"context,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// level
	// Enum: [unknown debug info warn error fatal]
	Level string `json:"level,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// message type
	// Enum: [i18n_key custom]
	MessageType string `json:"message_type,omitempty"`
}

// Validate validates this alerts processing log entry entity
func (m *AlertsProcessingLogEntryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertsProcessingLogEntryEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var alertsProcessingLogEntryEntityTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","debug","info","warn","error","fatal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertsProcessingLogEntryEntityTypeLevelPropEnum = append(alertsProcessingLogEntryEntityTypeLevelPropEnum, v)
	}
}

const (

	// AlertsProcessingLogEntryEntityLevelUnknown captures enum value "unknown"
	AlertsProcessingLogEntryEntityLevelUnknown string = "unknown"

	// AlertsProcessingLogEntryEntityLevelDebug captures enum value "debug"
	AlertsProcessingLogEntryEntityLevelDebug string = "debug"

	// AlertsProcessingLogEntryEntityLevelInfo captures enum value "info"
	AlertsProcessingLogEntryEntityLevelInfo string = "info"

	// AlertsProcessingLogEntryEntityLevelWarn captures enum value "warn"
	AlertsProcessingLogEntryEntityLevelWarn string = "warn"

	// AlertsProcessingLogEntryEntityLevelError captures enum value "error"
	AlertsProcessingLogEntryEntityLevelError string = "error"

	// AlertsProcessingLogEntryEntityLevelFatal captures enum value "fatal"
	AlertsProcessingLogEntryEntityLevelFatal string = "fatal"
)

// prop value enum
func (m *AlertsProcessingLogEntryEntity) validateLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertsProcessingLogEntryEntityTypeLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AlertsProcessingLogEntryEntity) validateLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", m.Level); err != nil {
		return err
	}

	return nil
}

var alertsProcessingLogEntryEntityTypeMessageTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["i18n_key","custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertsProcessingLogEntryEntityTypeMessageTypePropEnum = append(alertsProcessingLogEntryEntityTypeMessageTypePropEnum, v)
	}
}

const (

	// AlertsProcessingLogEntryEntityMessageTypeI18nKey captures enum value "i18n_key"
	AlertsProcessingLogEntryEntityMessageTypeI18nKey string = "i18n_key"

	// AlertsProcessingLogEntryEntityMessageTypeCustom captures enum value "custom"
	AlertsProcessingLogEntryEntityMessageTypeCustom string = "custom"
)

// prop value enum
func (m *AlertsProcessingLogEntryEntity) validateMessageTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertsProcessingLogEntryEntityTypeMessageTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AlertsProcessingLogEntryEntity) validateMessageType(formats strfmt.Registry) error {
	if swag.IsZero(m.MessageType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMessageTypeEnum("message_type", "body", m.MessageType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this alerts processing log entry entity based on context it is used
func (m *AlertsProcessingLogEntryEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertsProcessingLogEntryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertsProcessingLogEntryEntity) UnmarshalBinary(b []byte) error {
	var res AlertsProcessingLogEntryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}