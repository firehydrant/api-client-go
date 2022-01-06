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

// IncidentEntity Retrieve an incident
//
// swagger:model IncidentEntity
type IncidentEntity struct {

	// active
	Active string `json:"active,omitempty"`

	// channel id
	ChannelID string `json:"channel_id,omitempty"`

	// channel name
	ChannelName string `json:"channel_name,omitempty"`

	// channel reference
	ChannelReference string `json:"channel_reference,omitempty"`

	// channel status
	ChannelStatus string `json:"channel_status,omitempty"`

	// conference bridges
	ConferenceBridges *ConferenceBridgeEntity `json:"conference_bridges,omitempty"`

	// context object
	ContextObject string `json:"context_object,omitempty"`

	// The time the incident was opened
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy *AuthorEntity `json:"created_by,omitempty"`

	// current milestone
	CurrentMilestone string `json:"current_milestone,omitempty"`

	// customer impact summary
	CustomerImpactSummary string `json:"customer_impact_summary,omitempty"`

	// customers impacted
	CustomersImpacted string `json:"customers_impacted,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// environments
	Environments []*SuccinctEntity `json:"environments"`

	// functionalities
	Functionalities []*SuccinctEntity `json:"functionalities"`

	// UUID of the Incident
	ID string `json:"id,omitempty"`

	// impacts
	Impacts *ImpactEntity `json:"impacts,omitempty"`

	// incident roles
	IncidentRoles []*IncidentRoleEntity `json:"incident_roles"`

	// incident tickets
	IncidentTickets *TicketEntity `json:"incident_tickets,omitempty"`

	// incident url
	IncidentURL string `json:"incident_url,omitempty"`

	// A key/value of labels
	Labels interface{} `json:"labels,omitempty"`

	// last update
	LastUpdate string `json:"last_update,omitempty"`

	// milestones
	Milestones []*MilestoneEntity `json:"milestones"`

	// monetary impact
	MonetaryImpact string `json:"monetary_impact,omitempty"`

	// monetary impact cents
	MonetaryImpactCents string `json:"monetary_impact_cents,omitempty"`

	// Name of the incident
	Name string `json:"name,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// organization
	Organization *OrganizationEntity `json:"organization,omitempty"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// private id
	PrivateID string `json:"private_id,omitempty"`

	// private status page url
	PrivateStatusPageURL string `json:"private_status_page_url,omitempty"`

	// report id
	ReportID string `json:"report_id,omitempty"`

	// role assignments
	RoleAssignments *RoleAssignmentEntity `json:"role_assignments,omitempty"`

	// services
	Services []*SuccinctEntity `json:"services"`

	// severity
	Severity string `json:"severity,omitempty"`

	// severity condition
	SeverityCondition string `json:"severity_condition,omitempty"`

	// severity condition object
	SeverityConditionObject *ConditionEntity `json:"severity_condition_object,omitempty"`

	// severity impact
	SeverityImpact string `json:"severity_impact,omitempty"`

	// severity impact object
	SeverityImpactObject *ImpactEntity `json:"severity_impact_object,omitempty"`

	// The time the incident started
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// status pages
	StatusPages *StatusPageEntity `json:"status_pages,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tag list
	TagList string `json:"tag_list,omitempty"`
}

// Validate validates this incident entity
func (m *IncidentEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConferenceBridges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionalities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidentRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidentTickets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMilestones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverityConditionObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverityImpactObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
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

func (m *IncidentEntity) validateConferenceBridges(formats strfmt.Registry) error {
	if swag.IsZero(m.ConferenceBridges) { // not required
		return nil
	}

	if m.ConferenceBridges != nil {
		if err := m.ConferenceBridges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conference_bridges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conference_bridges")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentEntity) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateEnvironments(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) validateFunctionalities(formats strfmt.Registry) error {
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

func (m *IncidentEntity) validateImpacts(formats strfmt.Registry) error {
	if swag.IsZero(m.Impacts) { // not required
		return nil
	}

	if m.Impacts != nil {
		if err := m.Impacts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impacts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impacts")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateIncidentRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.IncidentRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.IncidentRoles); i++ {
		if swag.IsZero(m.IncidentRoles[i]) { // not required
			continue
		}

		if m.IncidentRoles[i] != nil {
			if err := m.IncidentRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incident_roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incident_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) validateIncidentTickets(formats strfmt.Registry) error {
	if swag.IsZero(m.IncidentTickets) { // not required
		return nil
	}

	if m.IncidentTickets != nil {
		if err := m.IncidentTickets.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident_tickets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incident_tickets")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateMilestones(formats strfmt.Registry) error {
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

func (m *IncidentEntity) validateOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateRoleAssignments(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleAssignments) { // not required
		return nil
	}

	if m.RoleAssignments != nil {
		if err := m.RoleAssignments.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_assignments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_assignments")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateServices(formats strfmt.Registry) error {
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

func (m *IncidentEntity) validateSeverityConditionObject(formats strfmt.Registry) error {
	if swag.IsZero(m.SeverityConditionObject) { // not required
		return nil
	}

	if m.SeverityConditionObject != nil {
		if err := m.SeverityConditionObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severity_condition_object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severity_condition_object")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateSeverityImpactObject(formats strfmt.Registry) error {
	if swag.IsZero(m.SeverityImpactObject) { // not required
		return nil
	}

	if m.SeverityImpactObject != nil {
		if err := m.SeverityImpactObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severity_impact_object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severity_impact_object")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IncidentEntity) validateStatusPages(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusPages) { // not required
		return nil
	}

	if m.StatusPages != nil {
		if err := m.StatusPages.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_pages")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_pages")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this incident entity based on the context it is used
func (m *IncidentEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConferenceBridges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFunctionalities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImpacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncidentRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncidentTickets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMilestones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverityConditionObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSeverityImpactObject(ctx, formats); err != nil {
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

func (m *IncidentEntity) contextValidateConferenceBridges(ctx context.Context, formats strfmt.Registry) error {

	if m.ConferenceBridges != nil {
		if err := m.ConferenceBridges.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conference_bridges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conference_bridges")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {
		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("created_by")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("created_by")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateEnvironments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Environments); i++ {

		if m.Environments[i] != nil {
			if err := m.Environments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("environments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("environments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) contextValidateFunctionalities(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IncidentEntity) contextValidateImpacts(ctx context.Context, formats strfmt.Registry) error {

	if m.Impacts != nil {
		if err := m.Impacts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("impacts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("impacts")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateIncidentRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncidentRoles); i++ {

		if m.IncidentRoles[i] != nil {
			if err := m.IncidentRoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incident_roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incident_roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) contextValidateIncidentTickets(ctx context.Context, formats strfmt.Registry) error {

	if m.IncidentTickets != nil {
		if err := m.IncidentTickets.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident_tickets")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incident_tickets")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateMilestones(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IncidentEntity) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.Organization != nil {
		if err := m.Organization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateRoleAssignments(ctx context.Context, formats strfmt.Registry) error {

	if m.RoleAssignments != nil {
		if err := m.RoleAssignments.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_assignments")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role_assignments")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IncidentEntity) contextValidateSeverityConditionObject(ctx context.Context, formats strfmt.Registry) error {

	if m.SeverityConditionObject != nil {
		if err := m.SeverityConditionObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severity_condition_object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severity_condition_object")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateSeverityImpactObject(ctx context.Context, formats strfmt.Registry) error {

	if m.SeverityImpactObject != nil {
		if err := m.SeverityImpactObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("severity_impact_object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("severity_impact_object")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateStatusPages(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusPages != nil {
		if err := m.StatusPages.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_pages")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_pages")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IncidentEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IncidentEntity) UnmarshalBinary(b []byte) error {
	var res IncidentEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
