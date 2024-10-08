// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsMilestonesFunnelEntity Metrics_MilestonesFunnelEntity model
//
// swagger:model Metrics_MilestonesFunnelEntity
type MetricsMilestonesFunnelEntity struct {

	// columns
	Columns []*MetricsMilestonesFunnelEntityColumnEntity `json:"columns"`

	// data
	Data []*MetricsMilestonesFunnelEntityDataBucketEntity `json:"data"`

	// groupings
	Groupings *MetricsMilestonesFunnelEntityGroupingsEntity `json:"groupings,omitempty"`

	// meta
	Meta *MetricsMilestonesFunnelEntityMetaEntity `json:"meta,omitempty"`
}

// Validate validates this metrics milestones funnel entity
func (m *MetricsMilestonesFunnelEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsMilestonesFunnelEntity) validateColumns(formats strfmt.Registry) error {
	if swag.IsZero(m.Columns) { // not required
		return nil
	}

	for i := 0; i < len(m.Columns); i++ {
		if swag.IsZero(m.Columns[i]) { // not required
			continue
		}

		if m.Columns[i] != nil {
			if err := m.Columns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetricsMilestonesFunnelEntity) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetricsMilestonesFunnelEntity) validateGroupings(formats strfmt.Registry) error {
	if swag.IsZero(m.Groupings) { // not required
		return nil
	}

	if m.Groupings != nil {
		if err := m.Groupings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupings")
			}
			return err
		}
	}

	return nil
}

func (m *MetricsMilestonesFunnelEntity) validateMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrics milestones funnel entity based on the context it is used
func (m *MetricsMilestonesFunnelEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumns(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsMilestonesFunnelEntity) contextValidateColumns(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Columns); i++ {

		if m.Columns[i] != nil {
			if err := m.Columns[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("columns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("columns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetricsMilestonesFunnelEntity) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetricsMilestonesFunnelEntity) contextValidateGroupings(ctx context.Context, formats strfmt.Registry) error {

	if m.Groupings != nil {
		if err := m.Groupings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("groupings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("groupings")
			}
			return err
		}
	}

	return nil
}

func (m *MetricsMilestonesFunnelEntity) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if m.Meta != nil {
		if err := m.Meta.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsMilestonesFunnelEntity) UnmarshalBinary(b []byte) error {
	var res MetricsMilestonesFunnelEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
