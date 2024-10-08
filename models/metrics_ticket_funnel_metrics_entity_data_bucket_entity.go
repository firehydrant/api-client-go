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

// MetricsTicketFunnelMetricsEntityDataBucketEntity metrics ticket funnel metrics entity data bucket entity
//
// swagger:model Metrics_TicketFunnelMetricsEntity_DataBucketEntity
type MetricsTicketFunnelMetricsEntityDataBucketEntity struct {

	// filter params
	FilterParams *MetricsTicketFunnelMetricsEntityDataBucketFilterParamsEntity `json:"filter_params,omitempty"`

	// The number of follow ups created
	FollowUpsCreated int32 `json:"follow_ups_created,omitempty"`

	// The number of follow ups completed
	FollowUpsDone int32 `json:"follow_ups_done,omitempty"`

	// The number of tasks created
	TasksCreated int32 `json:"tasks_created,omitempty"`

	// The number of tasks completed
	TasksDone int32 `json:"tasks_done,omitempty"`

	// The start datetime for the period
	// Format: date-time
	TimeBucket strfmt.DateTime `json:"time_bucket,omitempty"`
}

// Validate validates this metrics ticket funnel metrics entity data bucket entity
func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeBucket(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) validateFilterParams(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterParams) { // not required
		return nil
	}

	if m.FilterParams != nil {
		if err := m.FilterParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter_params")
			}
			return err
		}
	}

	return nil
}

func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) validateTimeBucket(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeBucket) { // not required
		return nil
	}

	if err := validate.FormatOf("time_bucket", "body", "date-time", m.TimeBucket.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrics ticket funnel metrics entity data bucket entity based on the context it is used
func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilterParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) contextValidateFilterParams(ctx context.Context, formats strfmt.Registry) error {

	if m.FilterParams != nil {
		if err := m.FilterParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsTicketFunnelMetricsEntityDataBucketEntity) UnmarshalBinary(b []byte) error {
	var res MetricsTicketFunnelMetricsEntityDataBucketEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
