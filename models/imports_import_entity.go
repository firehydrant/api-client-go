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

// ImportsImportEntity Imports_ImportEntity model
//
// swagger:model Imports_ImportEntity
type ImportsImportEntity struct {

	// state
	// Enum: [preprocessing ready_for_import importing completed]
	State string `json:"state,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this imports import entity
func (m *ImportsImportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
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

var importsImportEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["preprocessing","ready_for_import","importing","completed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		importsImportEntityTypeStatePropEnum = append(importsImportEntityTypeStatePropEnum, v)
	}
}

const (

	// ImportsImportEntityStatePreprocessing captures enum value "preprocessing"
	ImportsImportEntityStatePreprocessing string = "preprocessing"

	// ImportsImportEntityStateReadyForImport captures enum value "ready_for_import"
	ImportsImportEntityStateReadyForImport string = "ready_for_import"

	// ImportsImportEntityStateImporting captures enum value "importing"
	ImportsImportEntityStateImporting string = "importing"

	// ImportsImportEntityStateCompleted captures enum value "completed"
	ImportsImportEntityStateCompleted string = "completed"
)

// prop value enum
func (m *ImportsImportEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, importsImportEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ImportsImportEntity) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *ImportsImportEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this imports import entity based on context it is used
func (m *ImportsImportEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImportsImportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportsImportEntity) UnmarshalBinary(b []byte) error {
	var res ImportsImportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
