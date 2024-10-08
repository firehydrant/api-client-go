// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsMilestonesFunnelEntityGroupingsEntity metrics milestones funnel entity groupings entity
//
// swagger:model Metrics_MilestonesFunnelEntity_GroupingsEntity
type MetricsMilestonesFunnelEntityGroupingsEntity struct {

	// The bucket size for the data
	BucketSize string `json:"bucket_size,omitempty"`
}

// Validate validates this metrics milestones funnel entity groupings entity
func (m *MetricsMilestonesFunnelEntityGroupingsEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics milestones funnel entity groupings entity based on context it is used
func (m *MetricsMilestonesFunnelEntityGroupingsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntityGroupingsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntityGroupingsEntity) UnmarshalBinary(b []byte) error {
	var res MetricsMilestonesFunnelEntityGroupingsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
