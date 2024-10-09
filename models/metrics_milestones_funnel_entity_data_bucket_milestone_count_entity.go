// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity metrics milestones funnel entity data bucket milestone count entity
//
// swagger:model Metrics_MilestonesFunnelEntity_DataBucketMilestoneCountEntity
type MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity struct {

	// The frequency count of that milestone for the period
	Count int32 `json:"count,omitempty"`

	// The UUID of the milestone
	MilestoneID string `json:"milestone_id,omitempty"`
}

// Validate validates this metrics milestones funnel entity data bucket milestone count entity
func (m *MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics milestones funnel entity data bucket milestone count entity based on context it is used
func (m *MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity) UnmarshalBinary(b []byte) error {
	var res MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}