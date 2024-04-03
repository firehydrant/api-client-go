// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsSingleMetricsEntity Metrics_SingleMetricsEntity model
//
// swagger:model Metrics_SingleMetricsEntity
type MetricsSingleMetricsEntity struct {

	// Number of incidents in this time period for this component
	Count int32 `json:"count,omitempty"`

	// The UUID of the component
	ID string `json:"id,omitempty"`

	// Mean Time To Acknowledgement (seconds) for all incidents for this component in this time period
	Mtta int32 `json:"mtta,omitempty"`

	// Mean Time To Detection (seconds) for all incidents for this component in this time period
	Mttd int32 `json:"mttd,omitempty"`

	// Mean Time To Mitigation (seconds) for all incidents for this component in this time period
	Mttm int32 `json:"mttm,omitempty"`

	// Mean Time To Resolution (seconds) for all incidents for this component in this time period
	Mttr int32 `json:"mttr,omitempty"`

	// The name of the component
	Name string `json:"name,omitempty"`

	// Total time (seconds) the component was impacted (MTTR x Incident Count)
	TotalTime int32 `json:"total_time,omitempty"`
}

// Validate validates this metrics single metrics entity
func (m *MetricsSingleMetricsEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics single metrics entity based on context it is used
func (m *MetricsSingleMetricsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsSingleMetricsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsSingleMetricsEntity) UnmarshalBinary(b []byte) error {
	var res MetricsSingleMetricsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
