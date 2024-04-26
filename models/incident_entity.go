// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IncidentEntity IncidentEntity model
//
// swagger:model IncidentEntity
type IncidentEntity struct {

	// active
	Active bool `json:"active,omitempty"`

	// ai incident summary
	AiIncidentSummary string `json:"ai_incident_summary,omitempty"`

	// channel id
	ChannelID string `json:"channel_id,omitempty"`

	// channel name
	ChannelName string `json:"channel_name,omitempty"`

	// channel reference
	ChannelReference string `json:"channel_reference,omitempty"`

	// inoperative: 0, operational: 1, archived: 2
	ChannelStatus string `json:"channel_status,omitempty"`

	// conference bridges
	ConferenceBridges []*IncidentsConferenceBridgeEntity `json:"conference_bridges"`

	// context object
	ContextObject *IncidentsContextObjectEntity `json:"context_object,omitempty"`

	// conversations
	Conversations []*ConversationsAPIEntitiesReference `json:"conversations"`

	// The time the incident was opened
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy *AuthorEntity `json:"created_by,omitempty"`

	// current milestone
	// Enum: [started detected acknowledged investigating identified mitigated resolved postmortem_started postmortem_completed closed]
	CurrentMilestone string `json:"current_milestone,omitempty"`

	// custom fields
	CustomFields []*CustomFieldsFieldValue `json:"custom_fields"`

	// customer impact summary
	CustomerImpactSummary string `json:"customer_impact_summary,omitempty"`

	// customers impacted
	CustomersImpacted int32 `json:"customers_impacted,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// The time the incident was archived
	// Format: date-time
	DiscardedAt strfmt.DateTime `json:"discarded_at,omitempty"`

	// environments
	Environments []*SuccinctEntity `json:"environments"`

	// functionalities
	Functionalities []*SuccinctEntity `json:"functionalities"`

	// UUID of the Incident
	ID string `json:"id,omitempty"`

	// impacts
	Impacts []*IncidentsImpactEntity `json:"impacts"`

	// incident channels
	IncidentChannels []*IncidentsChannelEntity `json:"incident_channels"`

	// incident tickets
	IncidentTickets []*TicketingTicketEntity `json:"incident_tickets"`

	// incident url
	IncidentURL string `json:"incident_url,omitempty"`

	// A key/value of labels
	Labels interface{} `json:"labels,omitempty"`

	// last note
	LastNote *EventNoteEntity `json:"last_note,omitempty"`

	// last update
	LastUpdate string `json:"last_update,omitempty"`

	// milestones
	Milestones []*IncidentsMilestoneEntity `json:"milestones"`

	// monetary impact
	MonetaryImpact int32 `json:"monetary_impact,omitempty"`

	// monetary impact cents
	MonetaryImpactCents int32 `json:"monetary_impact_cents,omitempty"`

	// Name of the incident
	Name string `json:"name,omitempty"`

	// Incident number
	Number int32 `json:"number,omitempty"`

	// organization
	Organization *OrganizationEntity `json:"organization,omitempty"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// priority
	Priority string `json:"priority,omitempty"`

	// private id
	PrivateID string `json:"private_id,omitempty"`

	// private status page url
	PrivateStatusPageURL string `json:"private_status_page_url,omitempty"`

	// report id
	ReportID string `json:"report_id,omitempty"`

	// A list of objects attached to this item. Can be one of: LinkEntity, CustomerSupportIssueEntity, or GenericAttachmentEntity
	RetroExports []string `json:"retro_exports"`

	// role assignments
	RoleAssignments []*IncidentsRoleAssignmentEntity `json:"role_assignments"`

	// services
	Services []*SuccinctEntity `json:"services"`

	// severity
	Severity string `json:"severity,omitempty"`

	// severity color
	SeverityColor string `json:"severity_color,omitempty"`

	// severity condition
	SeverityCondition string `json:"severity_condition,omitempty"`

	// severity condition object
	SeverityConditionObject *SeverityMatrixConditionEntity `json:"severity_condition_object,omitempty"`

	// severity impact
	SeverityImpact string `json:"severity_impact,omitempty"`

	// severity impact object
	SeverityImpactObject *SeverityMatrixImpactEntity `json:"severity_impact_object,omitempty"`

	// The time the incident started
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// status pages
	StatusPages []*IncidentsStatusPageEntity `json:"status_pages"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tag list
	TagList []string `json:"tag_list"`

	// team assignments
	TeamAssignments []*IncidentsTeamAssignmentEntity `json:"team_assignments"`

	// ticket
	Ticket *TicketingTicketEntity `json:"ticket,omitempty"`
}

// Validate validates this incident entity
func (m *IncidentEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConferenceBridges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentMilestone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscardedAt(formats); err != nil {
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

	if err := m.validateIncidentChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncidentTickets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastNote(formats); err != nil {
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

	if err := m.validateTeamAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicket(formats); err != nil {
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

	for i := 0; i < len(m.ConferenceBridges); i++ {
		if swag.IsZero(m.ConferenceBridges[i]) { // not required
			continue
		}

		if m.ConferenceBridges[i] != nil {
			if err := m.ConferenceBridges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conference_bridges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conference_bridges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) validateContextObject(formats strfmt.Registry) error {
	if swag.IsZero(m.ContextObject) { // not required
		return nil
	}

	if m.ContextObject != nil {
		if err := m.ContextObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context_object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context_object")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) validateConversations(formats strfmt.Registry) error {
	if swag.IsZero(m.Conversations) { // not required
		return nil
	}

	for i := 0; i < len(m.Conversations); i++ {
		if swag.IsZero(m.Conversations[i]) { // not required
			continue
		}

		if m.Conversations[i] != nil {
			if err := m.Conversations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conversations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conversations" + "." + strconv.Itoa(i))
				}
				return err
			}
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

var incidentEntityTypeCurrentMilestonePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["started","detected","acknowledged","investigating","identified","mitigated","resolved","postmortem_started","postmortem_completed","closed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		incidentEntityTypeCurrentMilestonePropEnum = append(incidentEntityTypeCurrentMilestonePropEnum, v)
	}
}

const (

	// IncidentEntityCurrentMilestoneStarted captures enum value "started"
	IncidentEntityCurrentMilestoneStarted string = "started"

	// IncidentEntityCurrentMilestoneDetected captures enum value "detected"
	IncidentEntityCurrentMilestoneDetected string = "detected"

	// IncidentEntityCurrentMilestoneAcknowledged captures enum value "acknowledged"
	IncidentEntityCurrentMilestoneAcknowledged string = "acknowledged"

	// IncidentEntityCurrentMilestoneInvestigating captures enum value "investigating"
	IncidentEntityCurrentMilestoneInvestigating string = "investigating"

	// IncidentEntityCurrentMilestoneIdentified captures enum value "identified"
	IncidentEntityCurrentMilestoneIdentified string = "identified"

	// IncidentEntityCurrentMilestoneMitigated captures enum value "mitigated"
	IncidentEntityCurrentMilestoneMitigated string = "mitigated"

	// IncidentEntityCurrentMilestoneResolved captures enum value "resolved"
	IncidentEntityCurrentMilestoneResolved string = "resolved"

	// IncidentEntityCurrentMilestonePostmortemStarted captures enum value "postmortem_started"
	IncidentEntityCurrentMilestonePostmortemStarted string = "postmortem_started"

	// IncidentEntityCurrentMilestonePostmortemCompleted captures enum value "postmortem_completed"
	IncidentEntityCurrentMilestonePostmortemCompleted string = "postmortem_completed"

	// IncidentEntityCurrentMilestoneClosed captures enum value "closed"
	IncidentEntityCurrentMilestoneClosed string = "closed"
)

// prop value enum
func (m *IncidentEntity) validateCurrentMilestoneEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, incidentEntityTypeCurrentMilestonePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IncidentEntity) validateCurrentMilestone(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentMilestone) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrentMilestoneEnum("current_milestone", "body", m.CurrentMilestone); err != nil {
		return err
	}

	return nil
}

func (m *IncidentEntity) validateCustomFields(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomFields) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomFields); i++ {
		if swag.IsZero(m.CustomFields[i]) { // not required
			continue
		}

		if m.CustomFields[i] != nil {
			if err := m.CustomFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) validateDiscardedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscardedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("discarded_at", "body", "date-time", m.DiscardedAt.String(), formats); err != nil {
		return err
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

func (m *IncidentEntity) validateIncidentChannels(formats strfmt.Registry) error {
	if swag.IsZero(m.IncidentChannels) { // not required
		return nil
	}

	for i := 0; i < len(m.IncidentChannels); i++ {
		if swag.IsZero(m.IncidentChannels[i]) { // not required
			continue
		}

		if m.IncidentChannels[i] != nil {
			if err := m.IncidentChannels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incident_channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incident_channels" + "." + strconv.Itoa(i))
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

	for i := 0; i < len(m.IncidentTickets); i++ {
		if swag.IsZero(m.IncidentTickets[i]) { // not required
			continue
		}

		if m.IncidentTickets[i] != nil {
			if err := m.IncidentTickets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incident_tickets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incident_tickets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) validateLastNote(formats strfmt.Registry) error {
	if swag.IsZero(m.LastNote) { // not required
		return nil
	}

	if m.LastNote != nil {
		if err := m.LastNote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_note")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_note")
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

	for i := 0; i < len(m.RoleAssignments); i++ {
		if swag.IsZero(m.RoleAssignments[i]) { // not required
			continue
		}

		if m.RoleAssignments[i] != nil {
			if err := m.RoleAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("role_assignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("role_assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
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

func (m *IncidentEntity) validateTeamAssignments(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamAssignments) { // not required
		return nil
	}

	for i := 0; i < len(m.TeamAssignments); i++ {
		if swag.IsZero(m.TeamAssignments[i]) { // not required
			continue
		}

		if m.TeamAssignments[i] != nil {
			if err := m.TeamAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("team_assignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("team_assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) validateTicket(formats strfmt.Registry) error {
	if swag.IsZero(m.Ticket) { // not required
		return nil
	}

	if m.Ticket != nil {
		if err := m.Ticket.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ticket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ticket")
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

	if err := m.contextValidateContextObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomFields(ctx, formats); err != nil {
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

	if err := m.contextValidateIncidentChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncidentTickets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastNote(ctx, formats); err != nil {
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

	if err := m.contextValidateTeamAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTicket(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IncidentEntity) contextValidateConferenceBridges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConferenceBridges); i++ {

		if m.ConferenceBridges[i] != nil {
			if err := m.ConferenceBridges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conference_bridges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conference_bridges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) contextValidateContextObject(ctx context.Context, formats strfmt.Registry) error {

	if m.ContextObject != nil {
		if err := m.ContextObject.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context_object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context_object")
			}
			return err
		}
	}

	return nil
}

func (m *IncidentEntity) contextValidateConversations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conversations); i++ {

		if m.Conversations[i] != nil {
			if err := m.Conversations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conversations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conversations" + "." + strconv.Itoa(i))
				}
				return err
			}
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

func (m *IncidentEntity) contextValidateCustomFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomFields); i++ {

		if m.CustomFields[i] != nil {
			if err := m.CustomFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
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

	for i := 0; i < len(m.Impacts); i++ {

		if m.Impacts[i] != nil {
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

func (m *IncidentEntity) contextValidateIncidentChannels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncidentChannels); i++ {

		if m.IncidentChannels[i] != nil {
			if err := m.IncidentChannels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incident_channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incident_channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) contextValidateIncidentTickets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IncidentTickets); i++ {

		if m.IncidentTickets[i] != nil {
			if err := m.IncidentTickets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incident_tickets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("incident_tickets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) contextValidateLastNote(ctx context.Context, formats strfmt.Registry) error {

	if m.LastNote != nil {
		if err := m.LastNote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_note")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_note")
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

	for i := 0; i < len(m.RoleAssignments); i++ {

		if m.RoleAssignments[i] != nil {
			if err := m.RoleAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("role_assignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("role_assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
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

func (m *IncidentEntity) contextValidateTeamAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TeamAssignments); i++ {

		if m.TeamAssignments[i] != nil {
			if err := m.TeamAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("team_assignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("team_assignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IncidentEntity) contextValidateTicket(ctx context.Context, formats strfmt.Registry) error {

	if m.Ticket != nil {
		if err := m.Ticket.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ticket")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ticket")
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
