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

// ServiceEntity ServiceEntity model
//
// swagger:model ServiceEntity
type ServiceEntity struct {

	// List of active incident guids
	ActiveIncidents []*IncidentEntity `json:"active_incidents"`

	// alert on add
	AlertOnAdd bool `json:"alert_on_add,omitempty"`

	// allowed params
	AllowedParams []string `json:"allowed_params"`

	// auto add responding team
	AutoAddRespondingTeam bool `json:"auto_add_responding_team,omitempty"`

	// List of checklists associated with a service
	Checklists []*ChecklistTemplateEntity `json:"checklists"`

	// completed checks
	CompletedChecks int32 `json:"completed_checks,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Information about known linkages to representations of services outside of FireHydrant.
	ExternalResources []*ExternalResourceEntity `json:"external_resources"`

	// List of functionalities attached to the service
	Functionalities []*FunctionalityEntity `json:"functionalities"`

	// id
	ID string `json:"id,omitempty"`

	// An object of label key and values
	Labels interface{} `json:"labels,omitempty"`

	// last import
	LastImport *ImportsImportableResourceEntity `json:"last_import,omitempty"`

	// List of links attached to this service.
	Links []*LinksEntity `json:"links"`

	// If set, this field indicates that the service is managed by an integration and thus cannot be set manually
	ManagedBy string `json:"managed_by,omitempty"`

	// Indicates the settings of the catalog that manages this service
	ManagedBySettings interface{} `json:"managed_by_settings,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Team that owns the service
	Owner *TeamEntity `json:"owner,omitempty"`

	// service checklist updated at
	// Format: date-time
	ServiceChecklistUpdatedAt strfmt.DateTime `json:"service_checklist_updated_at,omitempty"`

	// service tier
	ServiceTier int32 `json:"service_tier,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// List of teams attached to the service
	Teams []*TeamEntity `json:"teams"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// The most recent user to update the current service
	UpdatedBy *AuthorEntity `json:"updated_by,omitempty"`
}

// Validate validates this service entity
func (m *ServiceEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveIncidents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChecklists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionalities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastImport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceChecklistUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTeams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEntity) validateActiveIncidents(formats strfmt.Registry) error {
	if swag.IsZero(m.ActiveIncidents) { // not required
		return nil
	}

	for i := 0; i < len(m.ActiveIncidents); i++ {
		if swag.IsZero(m.ActiveIncidents[i]) { // not required
			continue
		}

		if m.ActiveIncidents[i] != nil {
			if err := m.ActiveIncidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active_incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("active_incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceEntity) validateChecklists(formats strfmt.Registry) error {
	if swag.IsZero(m.Checklists) { // not required
		return nil
	}

	for i := 0; i < len(m.Checklists); i++ {
		if swag.IsZero(m.Checklists[i]) { // not required
			continue
		}

		if m.Checklists[i] != nil {
			if err := m.Checklists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checklists" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checklists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceEntity) validateExternalResources(formats strfmt.Registry) error {
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

func (m *ServiceEntity) validateFunctionalities(formats strfmt.Registry) error {
	if swag.IsZero(m.Functionalities) { // not required
		return nil
	}

	for i := 0; i < len(m.Functionalities); i++ {
		if swag.IsZero(m.Functionalities[i]) { // not required
			continue
		}

		if m.Functionalities[i] != nil {
			if err := m.Functionalities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("functionalities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("functionalities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceEntity) validateLastImport(formats strfmt.Registry) error {
	if swag.IsZero(m.LastImport) { // not required
		return nil
	}

	if m.LastImport != nil {
		if err := m.LastImport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_import")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_import")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEntity) validateLinks(formats strfmt.Registry) error {
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

func (m *ServiceEntity) validateOwner(formats strfmt.Registry) error {
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

func (m *ServiceEntity) validateServiceChecklistUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceChecklistUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("service_checklist_updated_at", "body", "date-time", m.ServiceChecklistUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceEntity) validateTeams(formats strfmt.Registry) error {
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

func (m *ServiceEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceEntity) validateUpdatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service entity based on the context it is used
func (m *ServiceEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActiveIncidents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChecklists(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFunctionalities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastImport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTeams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceEntity) contextValidateActiveIncidents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ActiveIncidents); i++ {

		if m.ActiveIncidents[i] != nil {
			if err := m.ActiveIncidents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active_incidents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("active_incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceEntity) contextValidateChecklists(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Checklists); i++ {

		if m.Checklists[i] != nil {
			if err := m.Checklists[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("checklists" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("checklists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceEntity) contextValidateExternalResources(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceEntity) contextValidateFunctionalities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Functionalities); i++ {

		if m.Functionalities[i] != nil {
			if err := m.Functionalities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("functionalities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("functionalities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceEntity) contextValidateLastImport(ctx context.Context, formats strfmt.Registry) error {

	if m.LastImport != nil {
		if err := m.LastImport.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_import")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_import")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceEntity) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceEntity) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceEntity) contextValidateTeams(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ServiceEntity) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updated_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updated_by")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceEntity) UnmarshalBinary(b []byte) error {
	var res ServiceEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
