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

// TicketingTicketEntity Ticketing_TicketEntity model
//
// swagger:model Ticketing_TicketEntity
type TicketingTicketEntity struct {

	// assignees
	Assignees []*AuthorEntity `json:"assignees"`

	// A list of objects attached to this item. Can be one of: LinkEntity, CustomerSupportIssueEntity, or GenericAttachmentEntity
	Attachments []interface{} `json:"attachments"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy *AuthorEntity `json:"created_by,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// due at
	// Format: date-time
	DueAt strfmt.DateTime `json:"due_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Milestone of incident that this ticket is related to
	IncidentCurrentMilestone string `json:"incident_current_milestone,omitempty"`

	// ID of incident that this ticket is related to
	IncidentID string `json:"incident_id,omitempty"`

	// Name of incident that this ticket is related to
	IncidentName string `json:"incident_name,omitempty"`

	// link
	Link *AttachmentsLinkEntity `json:"link,omitempty"`

	// priority
	Priority *TicketingPriorityEntity `json:"priority,omitempty"`

	// state
	// Enum: [open in_progress cancelled done]
	State string `json:"state,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tag list
	TagList []string `json:"tag_list"`

	// ID of task that this ticket is related to
	TaskID string `json:"task_id,omitempty"`

	// type
	// Enum: [incident task follow_up]
	Type string `json:"type,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this ticketing ticket entity
func (m *TicketingTicketEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDueAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *TicketingTicketEntity) validateAssignees(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignees) { // not required
		return nil
	}

	for i := 0; i < len(m.Assignees); i++ {
		if swag.IsZero(m.Assignees[i]) { // not required
			continue
		}

		if m.Assignees[i] != nil {
			if err := m.Assignees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TicketingTicketEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TicketingTicketEntity) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *TicketingTicketEntity) validateDueAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DueAt) { // not required
		return nil
	}

	if err := validate.FormatOf("due_at", "body", "date-time", m.DueAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TicketingTicketEntity) validateLink(formats strfmt.Registry) error {
	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

func (m *TicketingTicketEntity) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if m.Priority != nil {
		if err := m.Priority.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priority")
			}
			return err
		}
	}

	return nil
}

var ticketingTicketEntityTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["open","in_progress","cancelled","done"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ticketingTicketEntityTypeStatePropEnum = append(ticketingTicketEntityTypeStatePropEnum, v)
	}
}

const (

	// TicketingTicketEntityStateOpen captures enum value "open"
	TicketingTicketEntityStateOpen string = "open"

	// TicketingTicketEntityStateInProgress captures enum value "in_progress"
	TicketingTicketEntityStateInProgress string = "in_progress"

	// TicketingTicketEntityStateCancelled captures enum value "cancelled"
	TicketingTicketEntityStateCancelled string = "cancelled"

	// TicketingTicketEntityStateDone captures enum value "done"
	TicketingTicketEntityStateDone string = "done"
)

// prop value enum
func (m *TicketingTicketEntity) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ticketingTicketEntityTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TicketingTicketEntity) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var ticketingTicketEntityTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["incident","task","follow_up"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ticketingTicketEntityTypeTypePropEnum = append(ticketingTicketEntityTypeTypePropEnum, v)
	}
}

const (

	// TicketingTicketEntityTypeIncident captures enum value "incident"
	TicketingTicketEntityTypeIncident string = "incident"

	// TicketingTicketEntityTypeTask captures enum value "task"
	TicketingTicketEntityTypeTask string = "task"

	// TicketingTicketEntityTypeFollowUp captures enum value "follow_up"
	TicketingTicketEntityTypeFollowUp string = "follow_up"
)

// prop value enum
func (m *TicketingTicketEntity) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ticketingTicketEntityTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TicketingTicketEntity) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *TicketingTicketEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ticketing ticket entity based on the context it is used
func (m *TicketingTicketEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLink(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TicketingTicketEntity) contextValidateAssignees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Assignees); i++ {

		if m.Assignees[i] != nil {
			if err := m.Assignees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TicketingTicketEntity) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TicketingTicketEntity) contextValidateLink(ctx context.Context, formats strfmt.Registry) error {

	if m.Link != nil {
		if err := m.Link.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("link")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("link")
			}
			return err
		}
	}

	return nil
}

func (m *TicketingTicketEntity) contextValidatePriority(ctx context.Context, formats strfmt.Registry) error {

	if m.Priority != nil {
		if err := m.Priority.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priority")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priority")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TicketingTicketEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TicketingTicketEntity) UnmarshalBinary(b []byte) error {
	var res TicketingTicketEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
