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

// RunbookEntity RunbookEntity model
//
// swagger:model RunbookEntity
type RunbookEntity struct {

	// attachment rule
	AttachmentRule *RulesRuleEntity `json:"attachment_rule,omitempty"`

	// auto attach to restricted incidents
	AutoAttachToRestrictedIncidents bool `json:"auto_attach_to_restricted_incidents,omitempty"`

	// categories the runbook applies to
	Categories []string `json:"categories"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// created by
	CreatedBy *AuthorEntity `json:"created_by,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is editable
	IsEditable bool `json:"is_editable,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Team that owns the runbook
	Owner *TeamEntityLite `json:"owner,omitempty"`

	// runbook template id
	RunbookTemplateID string `json:"runbook_template_id,omitempty"`

	// steps
	Steps *RunbookStepEntity `json:"steps,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tutorial
	Tutorial bool `json:"tutorial,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// updated by
	UpdatedBy *AuthorEntity `json:"updated_by,omitempty"`

	// votes
	Votes *VotesEntity `json:"votes,omitempty"`
}

// Validate validates this runbook entity
func (m *RunbookEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbookEntity) validateAttachmentRule(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachmentRule) { // not required
		return nil
	}

	if m.AttachmentRule != nil {
		if err := m.AttachmentRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment_rule")
			}
			return err
		}
	}

	return nil
}

func (m *RunbookEntity) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunbookEntity) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *RunbookEntity) validateOwner(formats strfmt.Registry) error {
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

func (m *RunbookEntity) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	if m.Steps != nil {
		if err := m.Steps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("steps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("steps")
			}
			return err
		}
	}

	return nil
}

func (m *RunbookEntity) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RunbookEntity) validateUpdatedBy(formats strfmt.Registry) error {
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

func (m *RunbookEntity) validateVotes(formats strfmt.Registry) error {
	if swag.IsZero(m.Votes) { // not required
		return nil
	}

	if m.Votes != nil {
		if err := m.Votes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("votes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("votes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this runbook entity based on the context it is used
func (m *RunbookEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachmentRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RunbookEntity) contextValidateAttachmentRule(ctx context.Context, formats strfmt.Registry) error {

	if m.AttachmentRule != nil {
		if err := m.AttachmentRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment_rule")
			}
			return err
		}
	}

	return nil
}

func (m *RunbookEntity) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RunbookEntity) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RunbookEntity) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	if m.Steps != nil {
		if err := m.Steps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("steps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("steps")
			}
			return err
		}
	}

	return nil
}

func (m *RunbookEntity) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *RunbookEntity) contextValidateVotes(ctx context.Context, formats strfmt.Registry) error {

	if m.Votes != nil {
		if err := m.Votes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("votes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("votes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RunbookEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RunbookEntity) UnmarshalBinary(b []byte) error {
	var res RunbookEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
