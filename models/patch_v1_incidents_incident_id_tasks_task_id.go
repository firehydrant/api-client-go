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

// PatchV1IncidentsIncidentIDTasksTaskID Update a task's attributes
//
// swagger:model patchV1IncidentsIncidentIdTasksTaskId
type PatchV1IncidentsIncidentIDTasksTaskID struct {

	// The ID of the user assigned to the task.
	AssigneeID string `json:"assignee_id,omitempty"`

	// A description of what the task is for.
	Description string `json:"description,omitempty"`

	// The due date of the task.
	// Format: date-time
	DueAt strfmt.DateTime `json:"due_at,omitempty"`

	// The state of the task. One of: open, in_progress, cancelled, done
	State string `json:"state,omitempty"`

	// The title of the task.
	Title string `json:"title,omitempty"`
}

// Validate validates this patch v1 incidents incident Id tasks task Id
func (m *PatchV1IncidentsIncidentIDTasksTaskID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDueAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1IncidentsIncidentIDTasksTaskID) validateDueAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DueAt) { // not required
		return nil
	}

	if err := validate.FormatOf("due_at", "body", "date-time", m.DueAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 incidents incident Id tasks task Id based on context it is used
func (m *PatchV1IncidentsIncidentIDTasksTaskID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDTasksTaskID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDTasksTaskID) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentsIncidentIDTasksTaskID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
