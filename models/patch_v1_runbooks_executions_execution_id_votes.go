// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchV1RunbooksExecutionsExecutionIDVotes Upvote or downvote an object
//
// swagger:model patchV1RunbooksExecutionsExecutionIdVotes
type PatchV1RunbooksExecutionsExecutionIDVotes struct {

	// The direction you would like to vote, or if you dig it
	// Required: true
	// Enum: [up down dig]
	Direction *string `json:"direction"`
}

// Validate validates this patch v1 runbooks executions execution Id votes
func (m *PatchV1RunbooksExecutionsExecutionIDVotes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchV1RunbooksExecutionsExecutionIdVotesTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["up","down","dig"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchV1RunbooksExecutionsExecutionIdVotesTypeDirectionPropEnum = append(patchV1RunbooksExecutionsExecutionIdVotesTypeDirectionPropEnum, v)
	}
}

const (

	// PatchV1RunbooksExecutionsExecutionIDVotesDirectionUp captures enum value "up"
	PatchV1RunbooksExecutionsExecutionIDVotesDirectionUp string = "up"

	// PatchV1RunbooksExecutionsExecutionIDVotesDirectionDown captures enum value "down"
	PatchV1RunbooksExecutionsExecutionIDVotesDirectionDown string = "down"

	// PatchV1RunbooksExecutionsExecutionIDVotesDirectionDig captures enum value "dig"
	PatchV1RunbooksExecutionsExecutionIDVotesDirectionDig string = "dig"
)

// prop value enum
func (m *PatchV1RunbooksExecutionsExecutionIDVotes) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchV1RunbooksExecutionsExecutionIdVotesTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchV1RunbooksExecutionsExecutionIDVotes) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("direction", "body", m.Direction); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", *m.Direction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 runbooks executions execution Id votes based on context it is used
func (m *PatchV1RunbooksExecutionsExecutionIDVotes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1RunbooksExecutionsExecutionIDVotes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1RunbooksExecutionsExecutionIDVotes) UnmarshalBinary(b []byte) error {
	var res PatchV1RunbooksExecutionsExecutionIDVotes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
