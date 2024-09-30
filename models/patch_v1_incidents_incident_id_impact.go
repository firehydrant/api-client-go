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

// PatchV1IncidentsIncidentIDImpact Allows updating an incident's impacted infrastructure, with the option to
// move the incident into a different milestone and provide a note to update
// the incident timeline and any attached status pages. If this method is
// requested with the PUT verb, impacts will be completely replaced with the
// information in the request body, even if not provided (effectively clearing
// all impacts). If this method is requested with the PATCH verb, the provided
// impacts will be added or updated, but no impacts will be removed.
//
// swagger:model patchV1IncidentsIncidentIdImpact
type PatchV1IncidentsIncidentIDImpact struct {

	// impact
	Impact []*PatchV1IncidentsIncidentIDImpactImpactItems0 `json:"impact"`

	// milestone
	Milestone string `json:"milestone,omitempty"`

	// note
	Note string `json:"note,omitempty"`

	// status pages
	StatusPages []*PatchV1IncidentsIncidentIDImpactStatusPagesItems0 `json:"status_pages"`
}

// Validate validates this patch v1 incidents incident Id impact
func (m *PatchV1IncidentsIncidentIDImpact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImpact(formats); err != nil {
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

func (m *PatchV1IncidentsIncidentIDImpact) validateImpact(formats strfmt.Registry) error {
	if swag.IsZero(m.Impact) { // not required
		return nil
	}

	for i := 0; i < len(m.Impact); i++ {
		if swag.IsZero(m.Impact[i]) { // not required
			continue
		}

		if m.Impact[i] != nil {
			if err := m.Impact[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("impact" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("impact" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1IncidentsIncidentIDImpact) validateStatusPages(formats strfmt.Registry) error {
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

// ContextValidate validate this patch v1 incidents incident Id impact based on the context it is used
func (m *PatchV1IncidentsIncidentIDImpact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImpact(ctx, formats); err != nil {
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

func (m *PatchV1IncidentsIncidentIDImpact) contextValidateImpact(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Impact); i++ {

		if m.Impact[i] != nil {
			if err := m.Impact[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("impact" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("impact" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchV1IncidentsIncidentIDImpact) contextValidateStatusPages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusPages); i++ {

		if m.StatusPages[i] != nil {
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
func (m *PatchV1IncidentsIncidentIDImpact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDImpact) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentsIncidentIDImpact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1IncidentsIncidentIDImpactImpactItems0 patch v1 incidents incident ID impact impact items0
//
// swagger:model PatchV1IncidentsIncidentIDImpactImpactItems0
type PatchV1IncidentsIncidentIDImpactImpactItems0 struct {

	// condition id
	// Required: true
	ConditionID *string `json:"condition_id"`

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this patch v1 incidents incident ID impact impact items0
func (m *PatchV1IncidentsIncidentIDImpactImpactItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1IncidentsIncidentIDImpactImpactItems0) validateConditionID(formats strfmt.Registry) error {

	if err := validate.Required("condition_id", "body", m.ConditionID); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1IncidentsIncidentIDImpactImpactItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 incidents incident ID impact impact items0 based on context it is used
func (m *PatchV1IncidentsIncidentIDImpactImpactItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDImpactImpactItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDImpactImpactItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentsIncidentIDImpactImpactItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1IncidentsIncidentIDImpactStatusPagesItems0 patch v1 incidents incident ID impact status pages items0
//
// swagger:model PatchV1IncidentsIncidentIDImpactStatusPagesItems0
type PatchV1IncidentsIncidentIDImpactStatusPagesItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// integration slug
	// Required: true
	IntegrationSlug *string `json:"integration_slug"`
}

// Validate validates this patch v1 incidents incident ID impact status pages items0
func (m *PatchV1IncidentsIncidentIDImpactStatusPagesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1IncidentsIncidentIDImpactStatusPagesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1IncidentsIncidentIDImpactStatusPagesItems0) validateIntegrationSlug(formats strfmt.Registry) error {

	if err := validate.Required("integration_slug", "body", m.IntegrationSlug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 incidents incident ID impact status pages items0 based on context it is used
func (m *PatchV1IncidentsIncidentIDImpactStatusPagesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDImpactStatusPagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDImpactStatusPagesItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentsIncidentIDImpactStatusPagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
