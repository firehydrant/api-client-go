// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImportsImportableResourceEntity imports importable resource entity
//
// swagger:model Imports_ImportableResourceEntity
type ImportsImportableResourceEntity struct {

	// import errors
	ImportErrors []*ImportsImportErrorEntity `json:"import_errors"`

	// imported at
	// Format: date-time
	ImportedAt strfmt.DateTime `json:"imported_at,omitempty"`

	// remote id
	RemoteID string `json:"remote_id,omitempty"`

	// state
	// Enum: [selected skipped imported errored]
	State string `json:"state,omitempty"`
}

// Validate validates this imports importable resource entity
func (m *ImportsImportableResourceEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImportErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportsImportableResourceEntity) validateImportErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportErrors) { // not required
		return nil
	}

	for i := 0; i < len(m.ImportErrors); i++ {
		if swag.IsZero(m.ImportErrors[i]) { // not required
			continue
		}

		if m.ImportErrors[i] != nil {
			if err := m.ImportErrors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("import_errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("import_errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ImportsImportableResourceEntity) validateImportedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("imported_at", "body", "date-time", m.ImportedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var importsImportableResourceEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["selected","skipped","imported","errored"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		importsImportableResourceEntityTypeStatePropEnum = append(importsImportableResourceEntityTypeStatePropEnum, v)
	}
}

const (

	// ImportsImportableResourceEntityStateSelected captures enum value "selected"
	ImportsImportableResourceEntityStateSelected string = "selected"

	// ImportsImportableResourceEntityStateSkipped captures enum value "skipped"
	ImportsImportableResourceEntityStateSkipped string = "skipped"

	// ImportsImportableResourceEntityStateImported captures enum value "imported"
	ImportsImportableResourceEntityStateImported string = "imported"

	// ImportsImportableResourceEntityStateErrored captures enum value "errored"
	ImportsImportableResourceEntityStateErrored string = "errored"
)

// prop value enum
func (m *ImportsImportableResourceEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, importsImportableResourceEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ImportsImportableResourceEntity) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this imports importable resource entity based on the context it is used
func (m *ImportsImportableResourceEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImportErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportsImportableResourceEntity) contextValidateImportErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImportErrors); i++ {

		if m.ImportErrors[i] != nil {
			if err := m.ImportErrors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("import_errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("import_errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImportsImportableResourceEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportsImportableResourceEntity) UnmarshalBinary(b []byte) error {
	var res ImportsImportableResourceEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
