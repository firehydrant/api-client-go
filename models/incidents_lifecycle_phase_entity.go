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
)

// IncidentsLifecyclePhaseEntity incidents lifecycle phase entity
//
// swagger:model Incidents_LifecyclePhaseEntity
type IncidentsLifecyclePhaseEntity struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// milestones
	Milestones []*IncidentsLifecycleMilestoneEntity `json:"milestones"`

	// name
	Name string `json:"name,omitempty"`

	// position
	Position int32 `json:"position,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this incidents lifecycle phase entity
func (m *IncidentsLifecyclePhaseEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMilestones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentsLifecyclePhaseEntity) validateMilestones(formats strfmt.Registry) error {
	if swag.IsZero(m.Milestones) { // not required
		return nil
	}

	for i := 0; i < len(m.Milestones); i++ {
		if swag.IsZero(m.Milestones[i]) { // not required
			continue
		}

		if m.Milestones[i] != nil {
			if err := m.Milestones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("milestones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("milestones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this incidents lifecycle phase entity based on the context it is used
func (m *IncidentsLifecyclePhaseEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMilestones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentsLifecyclePhaseEntity) contextValidateMilestones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Milestones); i++ {

		if m.Milestones[i] != nil {
			if err := m.Milestones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("milestones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("milestones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncidentsLifecyclePhaseEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentsLifecyclePhaseEntity) UnmarshalBinary(b []byte) error {
	var res IncidentsLifecyclePhaseEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
