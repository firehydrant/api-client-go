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

// PatchV1ScimV2UsersID PATCH SCIM endpoint to update a User. This endpoint is used to update a resource's attributes.
//
// swagger:model patchV1ScimV2UsersId
type PatchV1ScimV2UsersID struct {

	// An array of operations to perform on the user
	// Required: true
	Operations []*PatchV1ScimV2UsersIDOperationsItems0 `json:"Operations"`

	// An optional trail to log the request
	Trail string `json:"trail,omitempty"`
}

// Validate validates this patch v1 scim v2 users Id
func (m *PatchV1ScimV2UsersID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1ScimV2UsersID) validateOperations(formats strfmt.Registry) error {

	if err := validate.Required("Operations", "body", m.Operations); err != nil {
		return err
	}

	for i := 0; i < len(m.Operations); i++ {
		if swag.IsZero(m.Operations[i]) { // not required
			continue
		}

		if m.Operations[i] != nil {
			if err := m.Operations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Operations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Operations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this patch v1 scim v2 users Id based on the context it is used
func (m *PatchV1ScimV2UsersID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1ScimV2UsersID) contextValidateOperations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Operations); i++ {

		if m.Operations[i] != nil {
			if err := m.Operations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Operations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Operations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1ScimV2UsersID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1ScimV2UsersID) UnmarshalBinary(b []byte) error {
	var res PatchV1ScimV2UsersID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1ScimV2UsersIDOperationsItems0 patch v1 scim v2 users ID operations items0
//
// swagger:model PatchV1ScimV2UsersIDOperationsItems0
type PatchV1ScimV2UsersIDOperationsItems0 struct {

	// The operation to perform on the user. Options are add, remove, replace
	// Required: true
	Op *string `json:"op"`

	// The path to the attribute to be modified
	// Required: true
	Path *string `json:"path"`
}

// Validate validates this patch v1 scim v2 users ID operations items0
func (m *PatchV1ScimV2UsersIDOperationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1ScimV2UsersIDOperationsItems0) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1ScimV2UsersIDOperationsItems0) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 scim v2 users ID operations items0 based on context it is used
func (m *PatchV1ScimV2UsersIDOperationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1ScimV2UsersIDOperationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1ScimV2UsersIDOperationsItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1ScimV2UsersIDOperationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
