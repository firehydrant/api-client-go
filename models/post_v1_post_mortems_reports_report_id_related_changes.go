// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1PostMortemsReportsReportIDRelatedChanges Add a related change to the post mortem report
// swagger:model postV1PostMortemsReportsReportIdRelatedChanges
type PostV1PostMortemsReportsReportIDRelatedChanges struct {

	// An ID of a change in the FireHydrant organization
	ChangeID string `json:"change_id,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// significance type
	// Enum: [caused fixed]
	SignificanceType string `json:"significance_type,omitempty"`
}

// Validate validates this post v1 post mortems reports report Id related changes
func (m *PostV1PostMortemsReportsReportIDRelatedChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSignificanceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postV1PostMortemsReportsReportIdRelatedChangesTypeSignificanceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["caused","fixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1PostMortemsReportsReportIdRelatedChangesTypeSignificanceTypePropEnum = append(postV1PostMortemsReportsReportIdRelatedChangesTypeSignificanceTypePropEnum, v)
	}
}

const (

	// PostV1PostMortemsReportsReportIDRelatedChangesSignificanceTypeCaused captures enum value "caused"
	PostV1PostMortemsReportsReportIDRelatedChangesSignificanceTypeCaused string = "caused"

	// PostV1PostMortemsReportsReportIDRelatedChangesSignificanceTypeFixed captures enum value "fixed"
	PostV1PostMortemsReportsReportIDRelatedChangesSignificanceTypeFixed string = "fixed"
)

// prop value enum
func (m *PostV1PostMortemsReportsReportIDRelatedChanges) validateSignificanceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postV1PostMortemsReportsReportIdRelatedChangesTypeSignificanceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostV1PostMortemsReportsReportIDRelatedChanges) validateSignificanceType(formats strfmt.Registry) error {

	if swag.IsZero(m.SignificanceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSignificanceTypeEnum("significance_type", "body", m.SignificanceType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1PostMortemsReportsReportIDRelatedChanges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1PostMortemsReportsReportIDRelatedChanges) UnmarshalBinary(b []byte) error {
	var res PostV1PostMortemsReportsReportIDRelatedChanges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
