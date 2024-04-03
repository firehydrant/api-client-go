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

// PatchV1SeverityMatrix Update available severities and impacts in your organization's severity matrix.
//
// swagger:model patchV1SeverityMatrix
type PatchV1SeverityMatrix struct {

	// data
	// Required: true
	Data []*PatchV1SeverityMatrixDataItems0 `json:"data"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this patch v1 severity matrix
func (m *PatchV1SeverityMatrix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1SeverityMatrix) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this patch v1 severity matrix based on the context it is used
func (m *PatchV1SeverityMatrix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1SeverityMatrix) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {

			if swag.IsZero(m.Data[i]) { // not required
				return nil
			}

			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1SeverityMatrix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1SeverityMatrix) UnmarshalBinary(b []byte) error {
	var res PatchV1SeverityMatrix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1SeverityMatrixDataItems0 patch v1 severity matrix data items0
//
// swagger:model PatchV1SeverityMatrixDataItems0
type PatchV1SeverityMatrixDataItems0 struct {

	// Condition id
	// Required: true
	ConditionID *string `json:"condition_id"`

	// Impact id
	// Required: true
	ImpactID *string `json:"impact_id"`

	// Slug of a severity
	// Required: true
	Severity *string `json:"severity"`
}

// Validate validates this patch v1 severity matrix data items0
func (m *PatchV1SeverityMatrixDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpactID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1SeverityMatrixDataItems0) validateConditionID(formats strfmt.Registry) error {

	if err := validate.Required("condition_id", "body", m.ConditionID); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1SeverityMatrixDataItems0) validateImpactID(formats strfmt.Registry) error {

	if err := validate.Required("impact_id", "body", m.ImpactID); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1SeverityMatrixDataItems0) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 severity matrix data items0 based on context it is used
func (m *PatchV1SeverityMatrixDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1SeverityMatrixDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1SeverityMatrixDataItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1SeverityMatrixDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
