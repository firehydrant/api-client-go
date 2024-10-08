// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsMilestonesFunnelEntityMetaEntity metrics milestones funnel entity meta entity
//
// swagger:model Metrics_MilestonesFunnelEntity_MetaEntity
type MetricsMilestonesFunnelEntityMetaEntity struct {

	// added milestones
	AddedMilestones []string `json:"added_milestones"`

	// deleted milestones
	DeletedMilestones []string `json:"deleted_milestones"`
}

// Validate validates this metrics milestones funnel entity meta entity
func (m *MetricsMilestonesFunnelEntityMetaEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics milestones funnel entity meta entity based on context it is used
func (m *MetricsMilestonesFunnelEntityMetaEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntityMetaEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntityMetaEntity) UnmarshalBinary(b []byte) error {
	var res MetricsMilestonesFunnelEntityMetaEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
