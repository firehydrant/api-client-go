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

// PatchV1FunctionalitiesFunctionalityID Update a functionalities attributes
//
// swagger:model patchV1FunctionalitiesFunctionalityId
type PatchV1FunctionalitiesFunctionalityID struct {

	// alert on add
	AlertOnAdd bool `json:"alert_on_add,omitempty"`

	// auto add responding team
	AutoAddRespondingTeam bool `json:"auto_add_responding_team,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// An array of external resources to attach to this service.
	ExternalResources []*PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0 `json:"external_resources"`

	// A hash of label keys and values
	Labels map[string]string `json:"labels,omitempty"`

	// An array of links to associate with this functionality. This will remove all links not present in the patch. Only acts if 'links' key is included in the payload.
	Links []*PatchV1FunctionalitiesFunctionalityIDLinksItems0 `json:"links"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner *PatchV1FunctionalitiesFunctionalityIDOwner `json:"owner,omitempty"`

	// If you are trying to remove a team as an owner from a functionality, set this to 'true'
	RemoveOwner bool `json:"remove_owner,omitempty"`

	// If set to true, any external_resources tagged on the service that are not included in the given array will be removed. Set this to true if you want to do a replacement operation for the external_resources
	RemoveRemainingExternalResources bool `json:"remove_remaining_external_resources,omitempty"`

	// Set this to true if you want to remove all of the services that are not included in the services array from the functionality
	RemoveRemainingServices *bool `json:"remove_remaining_services,omitempty"`

	// If set to true, any teams tagged on the service that are not included in the given array will be removed. Set this to true if you want to do a replacement operation for the teams
	RemoveRemainingTeams bool `json:"remove_remaining_teams,omitempty"`

	// services
	Services []*PatchV1FunctionalitiesFunctionalityIDServicesItems0 `json:"services"`

	// An array of teams to attach to this functionality.
	Teams []*PatchV1FunctionalitiesFunctionalityIDTeamsItems0 `json:"teams"`
}

// Validate validates this patch v1 functionalities functionality Id
func (m *PatchV1FunctionalitiesFunctionalityID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
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

func (m *PatchV1FunctionalitiesFunctionalityID) validateExternalResources(formats strfmt.Registry) error {
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

func (m *PatchV1FunctionalitiesFunctionalityID) validateLinks(formats strfmt.Registry) error {
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

func (m *PatchV1FunctionalitiesFunctionalityID) validateOwner(formats strfmt.Registry) error {
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

func (m *PatchV1FunctionalitiesFunctionalityID) validateServices(formats strfmt.Registry) error {
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

func (m *PatchV1FunctionalitiesFunctionalityID) validateTeams(formats strfmt.Registry) error {
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

// ContextValidate validate this patch v1 functionalities functionality Id based on the context it is used
func (m *PatchV1FunctionalitiesFunctionalityID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PatchV1FunctionalitiesFunctionalityID) contextValidateExternalResources(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PatchV1FunctionalitiesFunctionalityID) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PatchV1FunctionalitiesFunctionalityID) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PatchV1FunctionalitiesFunctionalityID) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PatchV1FunctionalitiesFunctionalityID) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PatchV1FunctionalitiesFunctionalityID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityID) UnmarshalBinary(b []byte) error {
	var res PatchV1FunctionalitiesFunctionalityID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0 patch v1 functionalities functionality ID external resources items0
//
// swagger:model PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0
type PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0 struct {

	// The integration slug for the external resource. Can be one of: github, opsgenie, pager_duty, statuspage, victorops. Not required if the resource has already been imported.
	ConnectionType string `json:"connection_type,omitempty"`

	// remote id
	// Required: true
	RemoteID *string `json:"remote_id"`

	// If you are trying to remove an external resource from a service, set this to 'true'.
	Remove bool `json:"remove,omitempty"`
}

// Validate validates this patch v1 functionalities functionality ID external resources items0
func (m *PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemoteID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0) validateRemoteID(formats strfmt.Registry) error {

	if err := validate.Required("remote_id", "body", m.RemoteID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 functionalities functionality ID external resources items0 based on context it is used
func (m *PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1FunctionalitiesFunctionalityIDExternalResourcesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1FunctionalitiesFunctionalityIDLinksItems0 patch v1 functionalities functionality ID links items0
//
// swagger:model PatchV1FunctionalitiesFunctionalityIDLinksItems0
type PatchV1FunctionalitiesFunctionalityIDLinksItems0 struct {

	// URL
	// Required: true
	HrefURL *string `json:"href_url"`

	// An optional URL to an icon representing this link
	IconURL string `json:"icon_url,omitempty"`

	// If updating an existing link, specify it's id.
	ID string `json:"id,omitempty"`

	// Short name used to display and identify this link
	// Required: true
	Name *string `json:"name"`

	// If you are trying to remove a link, set this to 'true'
	Remove bool `json:"remove,omitempty"`
}

// Validate validates this patch v1 functionalities functionality ID links items0
func (m *PatchV1FunctionalitiesFunctionalityIDLinksItems0) Validate(formats strfmt.Registry) error {
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

func (m *PatchV1FunctionalitiesFunctionalityIDLinksItems0) validateHrefURL(formats strfmt.Registry) error {

	if err := validate.Required("href_url", "body", m.HrefURL); err != nil {
		return err
	}

	return nil
}

func (m *PatchV1FunctionalitiesFunctionalityIDLinksItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 functionalities functionality ID links items0 based on context it is used
func (m *PatchV1FunctionalitiesFunctionalityIDLinksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDLinksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDLinksItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1FunctionalitiesFunctionalityIDLinksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1FunctionalitiesFunctionalityIDOwner An object representing a Team that owns the functionality
//
// swagger:model PatchV1FunctionalitiesFunctionalityIDOwner
type PatchV1FunctionalitiesFunctionalityIDOwner struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this patch v1 functionalities functionality ID owner
func (m *PatchV1FunctionalitiesFunctionalityIDOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1FunctionalitiesFunctionalityIDOwner) validateID(formats strfmt.Registry) error {

	if err := validate.Required("owner"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 functionalities functionality ID owner based on context it is used
func (m *PatchV1FunctionalitiesFunctionalityIDOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDOwner) UnmarshalBinary(b []byte) error {
	var res PatchV1FunctionalitiesFunctionalityIDOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1FunctionalitiesFunctionalityIDServicesItems0 patch v1 functionalities functionality ID services items0
//
// swagger:model PatchV1FunctionalitiesFunctionalityIDServicesItems0
type PatchV1FunctionalitiesFunctionalityIDServicesItems0 struct {

	// ID of a service
	// Required: true
	ID *string `json:"id"`

	// Set to true if you want to remove the given service from the functionality
	Remove bool `json:"remove,omitempty"`
}

// Validate validates this patch v1 functionalities functionality ID services items0
func (m *PatchV1FunctionalitiesFunctionalityIDServicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1FunctionalitiesFunctionalityIDServicesItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 functionalities functionality ID services items0 based on context it is used
func (m *PatchV1FunctionalitiesFunctionalityIDServicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDServicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDServicesItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1FunctionalitiesFunctionalityIDServicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1FunctionalitiesFunctionalityIDTeamsItems0 patch v1 functionalities functionality ID teams items0
//
// swagger:model PatchV1FunctionalitiesFunctionalityIDTeamsItems0
type PatchV1FunctionalitiesFunctionalityIDTeamsItems0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// If you are trying to remove a team from a functionality, set this to 'true'
	Remove bool `json:"remove,omitempty"`
}

// Validate validates this patch v1 functionalities functionality ID teams items0
func (m *PatchV1FunctionalitiesFunctionalityIDTeamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1FunctionalitiesFunctionalityIDTeamsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 functionalities functionality ID teams items0 based on context it is used
func (m *PatchV1FunctionalitiesFunctionalityIDTeamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDTeamsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1FunctionalitiesFunctionalityIDTeamsItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1FunctionalitiesFunctionalityIDTeamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
