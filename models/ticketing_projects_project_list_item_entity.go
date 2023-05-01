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

// TicketingProjectsProjectListItemEntity Ticketing_Projects_ProjectListItemEntity model
//
// swagger:model Ticketing_Projects_ProjectListItemEntity
type TicketingProjectsProjectListItemEntity struct {

	// config
	Config *TicketingProjectConfigEntity `json:"config,omitempty"`

	// connection slug
	ConnectionSlug string `json:"connection_slug,omitempty"`

	// field map
	FieldMap *TicketingProjectFieldMapEntity `json:"field_map,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this ticketing projects project list item entity
func (m *TicketingProjectsProjectListItemEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldMap(formats); err != nil {
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

func (m *TicketingProjectsProjectListItemEntity) validateConfig(formats strfmt.Registry) error {
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

func (m *TicketingProjectsProjectListItemEntity) validateFieldMap(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldMap) { // not required
		return nil
	}

	if m.FieldMap != nil {
		if err := m.FieldMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field_map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("field_map")
			}
			return err
		}
	}

	return nil
}

func (m *TicketingProjectsProjectListItemEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ticketing projects project list item entity based on the context it is used
func (m *TicketingProjectsProjectListItemEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFieldMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketingProjectsProjectListItemEntity) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
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

func (m *TicketingProjectsProjectListItemEntity) contextValidateFieldMap(ctx context.Context, formats strfmt.Registry) error {

	if m.FieldMap != nil {
		if err := m.FieldMap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field_map")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("field_map")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TicketingProjectsProjectListItemEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketingProjectsProjectListItemEntity) UnmarshalBinary(b []byte) error {
	var res TicketingProjectsProjectListItemEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}