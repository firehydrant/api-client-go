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

// IntegrationsAwsCloudtrailBatchEntity Integrations_Aws_CloudtrailBatchEntity model
//
// swagger:model Integrations_Aws_CloudtrailBatchEntity
type IntegrationsAwsCloudtrailBatchEntity struct {

	// connection
	Connection *IntegrationsAwsConnectionEntity `json:"connection,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// ends at
	// Format: date-time
	EndsAt strfmt.DateTime `json:"ends_at,omitempty"`

	// events created
	EventsCreated int32 `json:"events_created,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// starts at
	// Format: date-time
	StartsAt strfmt.DateTime `json:"starts_at,omitempty"`

	// status
	// Enum: [in_progress failed successful pending retried]
	Status string `json:"status,omitempty"`
}

// Validate validates this integrations aws cloudtrail batch entity
func (m *IntegrationsAwsCloudtrailBatchEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
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

func (m *IntegrationsAwsCloudtrailBatchEntity) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *IntegrationsAwsCloudtrailBatchEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsAwsCloudtrailBatchEntity) validateEndsAt(formats strfmt.Registry) error {
	if swag.IsZero(m.EndsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("ends_at", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsAwsCloudtrailBatchEntity) validateStartsAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("starts_at", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var integrationsAwsCloudtrailBatchEntityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in_progress","failed","successful","pending","retried"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationsAwsCloudtrailBatchEntityTypeStatusPropEnum = append(integrationsAwsCloudtrailBatchEntityTypeStatusPropEnum, v)
	}
}

const (

	// IntegrationsAwsCloudtrailBatchEntityStatusInProgress captures enum value "in_progress"
	IntegrationsAwsCloudtrailBatchEntityStatusInProgress string = "in_progress"

	// IntegrationsAwsCloudtrailBatchEntityStatusFailed captures enum value "failed"
	IntegrationsAwsCloudtrailBatchEntityStatusFailed string = "failed"

	// IntegrationsAwsCloudtrailBatchEntityStatusSuccessful captures enum value "successful"
	IntegrationsAwsCloudtrailBatchEntityStatusSuccessful string = "successful"

	// IntegrationsAwsCloudtrailBatchEntityStatusPending captures enum value "pending"
	IntegrationsAwsCloudtrailBatchEntityStatusPending string = "pending"

	// IntegrationsAwsCloudtrailBatchEntityStatusRetried captures enum value "retried"
	IntegrationsAwsCloudtrailBatchEntityStatusRetried string = "retried"
)

// prop value enum
func (m *IntegrationsAwsCloudtrailBatchEntity) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, integrationsAwsCloudtrailBatchEntityTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntegrationsAwsCloudtrailBatchEntity) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this integrations aws cloudtrail batch entity based on the context it is used
func (m *IntegrationsAwsCloudtrailBatchEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsAwsCloudtrailBatchEntity) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {
		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsAwsCloudtrailBatchEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsAwsCloudtrailBatchEntity) UnmarshalBinary(b []byte) error {
	var res IntegrationsAwsCloudtrailBatchEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
