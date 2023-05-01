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

// IncidentsRoleAssignmentEntity Incidents_RoleAssignmentEntity model
//
// swagger:model Incidents_RoleAssignmentEntity
type IncidentsRoleAssignmentEntity struct {

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// incident role
	IncidentRole *IncidentRoleEntity `json:"incident_role,omitempty"`

	// status
	// Enum: [active inactive]
	Status string `json:"status,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// user
	User *UserEntity `json:"user,omitempty"`
}

// Validate validates this incidents role assignment entity
func (m *IncidentsRoleAssignmentEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidentRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentsRoleAssignmentEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentsRoleAssignmentEntity) validateIncidentRole(formats strfmt.Registry) error {
	if swag.IsZero(m.IncidentRole) { // not required
		return nil
	}

	if m.IncidentRole != nil {
		if err := m.IncidentRole.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident_role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incident_role")
			}
			return err
		}
	}

	return nil
}

var incidentsRoleAssignmentEntityTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		incidentsRoleAssignmentEntityTypeStatusPropEnum = append(incidentsRoleAssignmentEntityTypeStatusPropEnum, v)
	}
}

const (

	// IncidentsRoleAssignmentEntityStatusActive captures enum value "active"
	IncidentsRoleAssignmentEntityStatusActive string = "active"

	// IncidentsRoleAssignmentEntityStatusInactive captures enum value "inactive"
	IncidentsRoleAssignmentEntityStatusInactive string = "inactive"
)

// prop value enum
func (m *IncidentsRoleAssignmentEntity) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, incidentsRoleAssignmentEntityTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IncidentsRoleAssignmentEntity) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IncidentsRoleAssignmentEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentsRoleAssignmentEntity) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this incidents role assignment entity based on the context it is used
func (m *IncidentsRoleAssignmentEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIncidentRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentsRoleAssignmentEntity) contextValidateIncidentRole(ctx context.Context, formats strfmt.Registry) error {

	if m.IncidentRole != nil {
		if err := m.IncidentRole.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident_role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incident_role")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentsRoleAssignmentEntity) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncidentsRoleAssignmentEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentsRoleAssignmentEntity) UnmarshalBinary(b []byte) error {
	var res IncidentsRoleAssignmentEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}