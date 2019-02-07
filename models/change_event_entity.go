// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeEventEntity Retrieve a change event
// swagger:model ChangeEventEntity
type ChangeEventEntity struct {

	// attachments
	Attachments []interface{} `json:"attachments"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// duration iso8601
	DurationIso8601 *string `json:"duration_iso8601,omitempty"`

	// duration ms
	DurationMs *int32 `json:"duration_ms,omitempty"`

	// ends at
	// Format: date-time
	EndsAt *strfmt.DateTime `json:"ends_at,omitempty"`

	// environments
	Environments []*EnvironmentEntity `json:"environments"`

	// id
	ID string `json:"id,omitempty"`

	// identities
	Identities []*ChangeIdentityEntity `json:"identities"`

	// labels
	Labels interface{} `json:"labels,omitempty"`

	// related changes
	RelatedChanges []*ChangeEntity `json:"related_changes"`

	// services
	Services []*ServiceEntity `json:"services"`

	// starts at
	// Format: date-time
	StartsAt strfmt.DateTime `json:"starts_at,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this change event entity
func (m *ChangeEventEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndsAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeEventEntity) validateEndsAt(formats strfmt.Registry) error {

	if swag.IsZero(m.EndsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("ends_at", "body", "date-time", m.EndsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChangeEventEntity) validateEnvironments(formats strfmt.Registry) error {

	if swag.IsZero(m.Environments) { // not required
		return nil
	}

	for i := 0; i < len(m.Environments); i++ {
		if swag.IsZero(m.Environments[i]) { // not required
			continue
		}

		if m.Environments[i] != nil {
			if err := m.Environments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChangeEventEntity) validateIdentities(formats strfmt.Registry) error {

	if swag.IsZero(m.Identities) { // not required
		return nil
	}

	for i := 0; i < len(m.Identities); i++ {
		if swag.IsZero(m.Identities[i]) { // not required
			continue
		}

		if m.Identities[i] != nil {
			if err := m.Identities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChangeEventEntity) validateRelatedChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.RelatedChanges) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedChanges); i++ {
		if swag.IsZero(m.RelatedChanges[i]) { // not required
			continue
		}

		if m.RelatedChanges[i] != nil {
			if err := m.RelatedChanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("related_changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChangeEventEntity) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChangeEventEntity) validateStartsAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartsAt) { // not required
		return nil
	}

	if err := validate.FormatOf("starts_at", "body", "date-time", m.StartsAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeEventEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeEventEntity) UnmarshalBinary(b []byte) error {
	var res ChangeEventEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
