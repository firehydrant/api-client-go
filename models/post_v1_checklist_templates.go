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

// PostV1ChecklistTemplates Creates a checklist template for the organization
//
// swagger:model postV1ChecklistTemplates
type PostV1ChecklistTemplates struct {

	// An array of checks for the checklist template
	// Required: true
	Checks []*PostV1ChecklistTemplatesChecksItems0 `json:"checks"`

	// Array of service IDs to attach checklist template to
	ConnectedServices []*PostV1ChecklistTemplatesConnectedServicesItems0 `json:"connected_services"`

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// The ID of the Team that owns the checklist template
	TeamID string `json:"team_id,omitempty"`
}

// Validate validates this post v1 checklist templates
func (m *PostV1ChecklistTemplates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectedServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChecklistTemplates) validateChecks(formats strfmt.Registry) error {

	if err := validate.Required("checks", "body", m.Checks); err != nil {
		return err
	}

	for i := 0; i < len(m.Checks); i++ {
		if swag.IsZero(m.Checks[i]) { // not required
			continue
		}

		if m.Checks[i] != nil {
			if err := m.Checks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChecklistTemplates) validateConnectedServices(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectedServices) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectedServices); i++ {
		if swag.IsZero(m.ConnectedServices[i]) { // not required
			continue
		}

		if m.ConnectedServices[i] != nil {
			if err := m.ConnectedServices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connected_services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connected_services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChecklistTemplates) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post v1 checklist templates based on the context it is used
func (m *PostV1ChecklistTemplates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChecklistTemplates) contextValidateChecks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Checks); i++ {

		if m.Checks[i] != nil {
			if err := m.Checks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ChecklistTemplates) contextValidateConnectedServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConnectedServices); i++ {

		if m.ConnectedServices[i] != nil {
			if err := m.ConnectedServices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connected_services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connected_services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ChecklistTemplates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ChecklistTemplates) UnmarshalBinary(b []byte) error {
	var res PostV1ChecklistTemplates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1ChecklistTemplatesChecksItems0 post v1 checklist templates checks items0
//
// swagger:model PostV1ChecklistTemplatesChecksItems0
type PostV1ChecklistTemplatesChecksItems0 struct {

	// The description of the check
	Description string `json:"description,omitempty"`

	// The name of the check
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post v1 checklist templates checks items0
func (m *PostV1ChecklistTemplatesChecksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChecklistTemplatesChecksItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 checklist templates checks items0 based on context it is used
func (m *PostV1ChecklistTemplatesChecksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ChecklistTemplatesChecksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ChecklistTemplatesChecksItems0) UnmarshalBinary(b []byte) error {
	var res PostV1ChecklistTemplatesChecksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1ChecklistTemplatesConnectedServicesItems0 post v1 checklist templates connected services items0
//
// swagger:model PostV1ChecklistTemplatesConnectedServicesItems0
type PostV1ChecklistTemplatesConnectedServicesItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this post v1 checklist templates connected services items0
func (m *PostV1ChecklistTemplatesConnectedServicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ChecklistTemplatesConnectedServicesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 checklist templates connected services items0 based on context it is used
func (m *PostV1ChecklistTemplatesConnectedServicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ChecklistTemplatesConnectedServicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ChecklistTemplatesConnectedServicesItems0) UnmarshalBinary(b []byte) error {
	var res PostV1ChecklistTemplatesConnectedServicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
