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

// PostV1Incidents Create a new incident
//
// swagger:model postV1Incidents
type PostV1Incidents struct {

	// List of alert IDs that this incident should be associated to
	AlertIds []string `json:"alert_ids"`

	// context object
	ContextObject *PostV1IncidentsContextObject `json:"context_object,omitempty"`

	// customer impact summary
	CustomerImpactSummary string `json:"customer_impact_summary,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// An array of impacted infrastructure
	Impacts []*PostV1IncidentsImpactsItems0 `json:"impacts"`

	// Key:value pairs to track custom data for the incident
	Labels map[string]string `json:"labels,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// priority
	Priority string `json:"priority,omitempty"`

	// restricted
	Restricted bool `json:"restricted,omitempty"`

	// List of ids of Runbooks to attach to this incident. Foregoes any conditions these Runbooks may have guarding automatic attachment.
	RunbookIds []string `json:"runbook_ids"`

	// severity
	Severity string `json:"severity,omitempty"`

	// severity condition id
	SeverityConditionID string `json:"severity_condition_id,omitempty"`

	// severity impact id
	SeverityImpactID string `json:"severity_impact_id,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// List of tags for the incident
	TagList []string `json:"tag_list"`
}

// Validate validates this post v1 incidents
func (m *PostV1Incidents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContextObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpacts(formats); err != nil {
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

func (m *PostV1Incidents) validateContextObject(formats strfmt.Registry) error {
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

func (m *PostV1Incidents) validateImpacts(formats strfmt.Registry) error {
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

func (m *PostV1Incidents) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post v1 incidents based on the context it is used
func (m *PostV1Incidents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContextObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImpacts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1Incidents) contextValidateContextObject(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PostV1Incidents) contextValidateImpacts(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PostV1Incidents) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1Incidents) UnmarshalBinary(b []byte) error {
	var res PostV1Incidents
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1IncidentsContextObject post v1 incidents context object
//
// swagger:model PostV1IncidentsContextObject
type PostV1IncidentsContextObject struct {

	// context description
	ContextDescription string `json:"context_description,omitempty"`

	// context tag
	// Required: true
	// Enum: [runbook_testing runbook_isolation_testing]
	ContextTag *string `json:"context_tag"`

	// object id
	// Required: true
	ObjectID *string `json:"object_id"`

	// object type
	// Required: true
	// Enum: [Runbooks::Runbook]
	ObjectType *string `json:"object_type"`
}

// Validate validates this post v1 incidents context object
func (m *PostV1IncidentsContextObject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContextTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postV1IncidentsContextObjectTypeContextTagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["runbook_testing","runbook_isolation_testing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1IncidentsContextObjectTypeContextTagPropEnum = append(postV1IncidentsContextObjectTypeContextTagPropEnum, v)
	}
}

const (

	// PostV1IncidentsContextObjectContextTagRunbookTesting captures enum value "runbook_testing"
	PostV1IncidentsContextObjectContextTagRunbookTesting string = "runbook_testing"

	// PostV1IncidentsContextObjectContextTagRunbookIsolationTesting captures enum value "runbook_isolation_testing"
	PostV1IncidentsContextObjectContextTagRunbookIsolationTesting string = "runbook_isolation_testing"
)

// prop value enum
func (m *PostV1IncidentsContextObject) validateContextTagEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1IncidentsContextObjectTypeContextTagPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1IncidentsContextObject) validateContextTag(formats strfmt.Registry) error {

	if err := validate.Required("context_object"+"."+"context_tag", "body", m.ContextTag); err != nil {
		return err
	}

	// value enum
	if err := m.validateContextTagEnum("context_object"+"."+"context_tag", "body", *m.ContextTag); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentsContextObject) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("context_object"+"."+"object_id", "body", m.ObjectID); err != nil {
		return err
	}

	return nil
}

var postV1IncidentsContextObjectTypeObjectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Runbooks::Runbook"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1IncidentsContextObjectTypeObjectTypePropEnum = append(postV1IncidentsContextObjectTypeObjectTypePropEnum, v)
	}
}

const (

	// PostV1IncidentsContextObjectObjectTypeRunbooksRunbook captures enum value "Runbooks::Runbook"
	PostV1IncidentsContextObjectObjectTypeRunbooksRunbook string = "Runbooks::Runbook"
)

// prop value enum
func (m *PostV1IncidentsContextObject) validateObjectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1IncidentsContextObjectTypeObjectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1IncidentsContextObject) validateObjectType(formats strfmt.Registry) error {

	if err := validate.Required("context_object"+"."+"object_type", "body", m.ObjectType); err != nil {
		return err
	}

	// value enum
	if err := m.validateObjectTypeEnum("context_object"+"."+"object_type", "body", *m.ObjectType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 incidents context object based on context it is used
func (m *PostV1IncidentsContextObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentsContextObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentsContextObject) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentsContextObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1IncidentsImpactsItems0 post v1 incidents impacts items0
//
// swagger:model PostV1IncidentsImpactsItems0
type PostV1IncidentsImpactsItems0 struct {

	// The ID of the impact condition. Find these at /v1/severity_matrix/conditions
	// Required: true
	ConditionID *string `json:"condition_id"`

	// The ID of the impacted infrastructure
	// Required: true
	ID *string `json:"id"`

	// The type of impacted infrastructure. One of: environment, functionality, or service
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this post v1 incidents impacts items0
func (m *PostV1IncidentsImpactsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentsImpactsItems0) validateConditionID(formats strfmt.Registry) error {

	if err := validate.Required("condition_id", "body", m.ConditionID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentsImpactsItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentsImpactsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 incidents impacts items0 based on context it is used
func (m *PostV1IncidentsImpactsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentsImpactsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentsImpactsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentsImpactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
