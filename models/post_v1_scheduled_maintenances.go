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

// PostV1ScheduledMaintenances Create a new scheduled maintenance event
//
// swagger:model postV1ScheduledMaintenances
type PostV1ScheduledMaintenances struct {

	// description
	Description string `json:"description,omitempty"`

	// ISO8601 timestamp for the end time of the scheduled maintenance
	// Required: true
	// Format: date-time
	EndsAt *strfmt.DateTime `json:"ends_at"`

	// An array of impact/condition combinations
	Impacts []*PostV1ScheduledMaintenancesImpactsItems0 `json:"impacts"`

	// A json object of label keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// ISO8601 timestamp for the start time of the scheduled maintenance
	// Required: true
	// Format: date-time
	StartsAt *strfmt.DateTime `json:"starts_at"`

	// An array of status pages to display this maintenance on
	StatusPages []*PostV1ScheduledMaintenancesStatusPagesItems0 `json:"status_pages"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this post v1 scheduled maintenances
func (m *PostV1ScheduledMaintenances) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusPages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ScheduledMaintenances) validateEndsAt(formats strfmt.Registry) error {

	if err := validate.Required("ends_at", "body", m.EndsAt); err != nil {
		return err
	}

	if err := validate.FormatOf("ends_at", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ScheduledMaintenances) validateImpacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Impacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Impacts); i++ {
		if swag.IsZero(m.Impacts[i]) { // not required
			continue
		}

		if m.Impacts[i] != nil {
			if err := m.Impacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("impacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("impacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ScheduledMaintenances) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ScheduledMaintenances) validateStartsAt(formats strfmt.Registry) error {

	if err := validate.Required("starts_at", "body", m.StartsAt); err != nil {
		return err
	}

	if err := validate.FormatOf("starts_at", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ScheduledMaintenances) validateStatusPages(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusPages) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusPages); i++ {
		if swag.IsZero(m.StatusPages[i]) { // not required
			continue
		}

		if m.StatusPages[i] != nil {
			if err := m.StatusPages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_pages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status_pages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post v1 scheduled maintenances based on the context it is used
func (m *PostV1ScheduledMaintenances) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImpacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusPages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ScheduledMaintenances) contextValidateImpacts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Impacts); i++ {

		if m.Impacts[i] != nil {

			if swag.IsZero(m.Impacts[i]) { // not required
				return nil
			}

			if err := m.Impacts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("impacts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("impacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1ScheduledMaintenances) contextValidateStatusPages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusPages); i++ {

		if m.StatusPages[i] != nil {

			if swag.IsZero(m.StatusPages[i]) { // not required
				return nil
			}

			if err := m.StatusPages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("status_pages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("status_pages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ScheduledMaintenances) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ScheduledMaintenances) UnmarshalBinary(b []byte) error {
	var res PostV1ScheduledMaintenances
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1ScheduledMaintenancesImpactsItems0 post v1 scheduled maintenances impacts items0
//
// swagger:model PostV1ScheduledMaintenancesImpactsItems0
type PostV1ScheduledMaintenancesImpactsItems0 struct {

	// The id of the condition
	// Required: true
	ConditionID *string `json:"condition_id"`

	// The id of impact
	// Required: true
	ID *string `json:"id"`

	// The type of impact
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this post v1 scheduled maintenances impacts items0
func (m *PostV1ScheduledMaintenancesImpactsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ScheduledMaintenancesImpactsItems0) validateConditionID(formats strfmt.Registry) error {

	if err := validate.Required("condition_id", "body", m.ConditionID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ScheduledMaintenancesImpactsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ScheduledMaintenancesImpactsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 scheduled maintenances impacts items0 based on context it is used
func (m *PostV1ScheduledMaintenancesImpactsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ScheduledMaintenancesImpactsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ScheduledMaintenancesImpactsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1ScheduledMaintenancesImpactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1ScheduledMaintenancesStatusPagesItems0 post v1 scheduled maintenances status pages items0
//
// swagger:model PostV1ScheduledMaintenancesStatusPagesItems0
type PostV1ScheduledMaintenancesStatusPagesItems0 struct {

	// The UUID of the status page to display this maintenance on
	// Required: true
	ConnectionID *string `json:"connection_id"`

	// The slug identifying the type of status page
	IntegrationSlug string `json:"integration_slug,omitempty"`
}

// Validate validates this post v1 scheduled maintenances status pages items0
func (m *PostV1ScheduledMaintenancesStatusPagesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ScheduledMaintenancesStatusPagesItems0) validateConnectionID(formats strfmt.Registry) error {

	if err := validate.Required("connection_id", "body", m.ConnectionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 scheduled maintenances status pages items0 based on context it is used
func (m *PostV1ScheduledMaintenancesStatusPagesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ScheduledMaintenancesStatusPagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ScheduledMaintenancesStatusPagesItems0) UnmarshalBinary(b []byte) error {
	var res PostV1ScheduledMaintenancesStatusPagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
