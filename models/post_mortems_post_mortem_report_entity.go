// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostMortemsPostMortemReportEntity PostMortems_PostMortemReportEntity model
//
// swagger:model PostMortems_PostMortemReportEntity
type PostMortemsPostMortemReportEntity struct {

	// additional details
	AdditionalDetails []string `json:"additional_details"`

	// calendar events
	CalendarEvents *CalendarsEventEntity `json:"calendar_events,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// incident
	Incident *IncidentEntity `json:"incident,omitempty"`

	// incident id
	IncidentID string `json:"incident_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// questions
	Questions *PostMortemsQuestionEntity `json:"questions,omitempty"`

	// retrospective id
	RetrospectiveID string `json:"retrospective_id,omitempty"`

	// retrospective note
	RetrospectiveNote string `json:"retrospective_note,omitempty"`

	// retrospective shim
	RetrospectiveShim bool `json:"retrospective_shim,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tag list
	TagList []string `json:"tag_list"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this post mortems post mortem report entity
func (m *PostMortemsPostMortemReportEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalendarEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncident(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuestions(formats); err != nil {
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

func (m *PostMortemsPostMortemReportEntity) validateCalendarEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.CalendarEvents) { // not required
		return nil
	}

	if m.CalendarEvents != nil {
		if err := m.CalendarEvents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calendar_events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("calendar_events")
			}
			return err
		}
	}

	return nil
}

func (m *PostMortemsPostMortemReportEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostMortemsPostMortemReportEntity) validateIncident(formats strfmt.Registry) error {
	if swag.IsZero(m.Incident) { // not required
		return nil
	}

	if m.Incident != nil {
		if err := m.Incident.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incident")
			}
			return err
		}
	}

	return nil
}

func (m *PostMortemsPostMortemReportEntity) validateQuestions(formats strfmt.Registry) error {
	if swag.IsZero(m.Questions) { // not required
		return nil
	}

	if m.Questions != nil {
		if err := m.Questions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("questions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("questions")
			}
			return err
		}
	}

	return nil
}

func (m *PostMortemsPostMortemReportEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post mortems post mortem report entity based on the context it is used
func (m *PostMortemsPostMortemReportEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCalendarEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncident(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuestions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostMortemsPostMortemReportEntity) contextValidateCalendarEvents(ctx context.Context, formats strfmt.Registry) error {

	if m.CalendarEvents != nil {
		if err := m.CalendarEvents.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calendar_events")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("calendar_events")
			}
			return err
		}
	}

	return nil
}

func (m *PostMortemsPostMortemReportEntity) contextValidateIncident(ctx context.Context, formats strfmt.Registry) error {

	if m.Incident != nil {
		if err := m.Incident.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incident")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incident")
			}
			return err
		}
	}

	return nil
}

func (m *PostMortemsPostMortemReportEntity) contextValidateQuestions(ctx context.Context, formats strfmt.Registry) error {

	if m.Questions != nil {
		if err := m.Questions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("questions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("questions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostMortemsPostMortemReportEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostMortemsPostMortemReportEntity) UnmarshalBinary(b []byte) error {
	var res PostMortemsPostMortemReportEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
