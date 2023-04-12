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

// PostV1Functionalities Creates a functionality for the organization
//
// swagger:model postV1Functionalities
type PostV1Functionalities struct {

	// alert on add
	AlertOnAdd bool `json:"alert_on_add,omitempty"`

	// auto add responding team
	AutoAddRespondingTeam bool `json:"auto_add_responding_team,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// An array of external resources to attach to this service.
	ExternalResources []*PostV1FunctionalitiesExternalResourcesItems0 `json:"external_resources"`

	// A hash of label keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// An array of links to associate with this service
	Links []*PostV1FunctionalitiesLinksItems0 `json:"links"`

	// name
	// Required: true
	Name *string `json:"name"`

	// owner
	Owner *PostV1FunctionalitiesOwner `json:"owner,omitempty"`

	// services
	Services []*PostV1FunctionalitiesServicesItems0 `json:"services"`

	// An array of teams to attach to this service.
	Teams []*PostV1FunctionalitiesTeamsItems0 `json:"teams"`
}

// Validate validates this post v1 functionalities
func (m *PostV1Functionalities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1Functionalities) validateExternalResources(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalResources) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalResources); i++ {
		if swag.IsZero(m.ExternalResources[i]) { // not required
			continue
		}

		if m.ExternalResources[i] != nil {
			if err := m.ExternalResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1Functionalities) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1Functionalities) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostV1Functionalities) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *PostV1Functionalities) validateServices(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1Functionalities) validateTeams(formats strfmt.Registry) error {
	if swag.IsZero(m.Teams) { // not required
		return nil
	}

	for i := 0; i < len(m.Teams); i++ {
		if swag.IsZero(m.Teams[i]) { // not required
			continue
		}

		if m.Teams[i] != nil {
			if err := m.Teams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post v1 functionalities based on the context it is used
func (m *PostV1Functionalities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1Functionalities) contextValidateExternalResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalResources); i++ {

		if m.ExternalResources[i] != nil {
			if err := m.ExternalResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external_resources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("external_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1Functionalities) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1Functionalities) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *PostV1Functionalities) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Services); i++ {

		if m.Services[i] != nil {
			if err := m.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1Functionalities) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Teams); i++ {

		if m.Teams[i] != nil {
			if err := m.Teams[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("teams" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("teams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1Functionalities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1Functionalities) UnmarshalBinary(b []byte) error {
	var res PostV1Functionalities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1FunctionalitiesExternalResourcesItems0 post v1 functionalities external resources items0
//
// swagger:model PostV1FunctionalitiesExternalResourcesItems0
type PostV1FunctionalitiesExternalResourcesItems0 struct {

	// The integration slug for the external resource. Can be one of: github, opsgenie, pager_duty, statuspage, victorops. Not required if the resource has already been imported.
	ConnectionType string `json:"connection_type,omitempty"`

	// remote id
	// Required: true
	RemoteID *string `json:"remote_id"`
}

// Validate validates this post v1 functionalities external resources items0
func (m *PostV1FunctionalitiesExternalResourcesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1FunctionalitiesExternalResourcesItems0) validateRemoteID(formats strfmt.Registry) error {

	if err := validate.Required("remote_id", "body", m.RemoteID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 functionalities external resources items0 based on context it is used
func (m *PostV1FunctionalitiesExternalResourcesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1FunctionalitiesExternalResourcesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1FunctionalitiesExternalResourcesItems0) UnmarshalBinary(b []byte) error {
	var res PostV1FunctionalitiesExternalResourcesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1FunctionalitiesLinksItems0 post v1 functionalities links items0
//
// swagger:model PostV1FunctionalitiesLinksItems0
type PostV1FunctionalitiesLinksItems0 struct {

	// URL
	// Required: true
	HrefURL *string `json:"href_url"`

	// An optional URL to an icon representing this link
	IconURL string `json:"icon_url,omitempty"`

	// Short name used to display and identify this link
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post v1 functionalities links items0
func (m *PostV1FunctionalitiesLinksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHrefURL(formats); err != nil {
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

func (m *PostV1FunctionalitiesLinksItems0) validateHrefURL(formats strfmt.Registry) error {

	if err := validate.Required("href_url", "body", m.HrefURL); err != nil {
		return err
	}

	return nil
}

func (m *PostV1FunctionalitiesLinksItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 functionalities links items0 based on context it is used
func (m *PostV1FunctionalitiesLinksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1FunctionalitiesLinksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1FunctionalitiesLinksItems0) UnmarshalBinary(b []byte) error {
	var res PostV1FunctionalitiesLinksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1FunctionalitiesOwner An object representing a Team that owns the service
//
// swagger:model PostV1FunctionalitiesOwner
type PostV1FunctionalitiesOwner struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this post v1 functionalities owner
func (m *PostV1FunctionalitiesOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1FunctionalitiesOwner) validateID(formats strfmt.Registry) error {

	if err := validate.Required("owner"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 functionalities owner based on context it is used
func (m *PostV1FunctionalitiesOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1FunctionalitiesOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1FunctionalitiesOwner) UnmarshalBinary(b []byte) error {
	var res PostV1FunctionalitiesOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1FunctionalitiesServicesItems0 post v1 functionalities services items0
//
// swagger:model PostV1FunctionalitiesServicesItems0
type PostV1FunctionalitiesServicesItems0 struct {

	// ID of a service
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this post v1 functionalities services items0
func (m *PostV1FunctionalitiesServicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1FunctionalitiesServicesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 functionalities services items0 based on context it is used
func (m *PostV1FunctionalitiesServicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1FunctionalitiesServicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1FunctionalitiesServicesItems0) UnmarshalBinary(b []byte) error {
	var res PostV1FunctionalitiesServicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1FunctionalitiesTeamsItems0 post v1 functionalities teams items0
//
// swagger:model PostV1FunctionalitiesTeamsItems0
type PostV1FunctionalitiesTeamsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this post v1 functionalities teams items0
func (m *PostV1FunctionalitiesTeamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1FunctionalitiesTeamsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 functionalities teams items0 based on context it is used
func (m *PostV1FunctionalitiesTeamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1FunctionalitiesTeamsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1FunctionalitiesTeamsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1FunctionalitiesTeamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
