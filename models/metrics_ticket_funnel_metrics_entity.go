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

// MetricsTicketFunnelMetricsEntity Metrics_TicketFunnelMetricsEntity model
//
// swagger:model Metrics_TicketFunnelMetricsEntity
type MetricsTicketFunnelMetricsEntity struct {

	// data
	Data []*MetricsTicketFunnelMetricsEntityDataBucketEntity `json:"data"`

	// groupings
	Groupings *MetricsTicketFunnelMetricsEntityGroupingsEntity `json:"groupings,omitempty"`
}

// Validate validates this metrics ticket funnel metrics entity
func (m *MetricsTicketFunnelMetricsEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsTicketFunnelMetricsEntity) validateData(formats strfmt.Registry) error {
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

func (m *MetricsTicketFunnelMetricsEntity) validateGroupings(formats strfmt.Registry) error {
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

// ContextValidate validate this metrics ticket funnel metrics entity based on the context it is used
func (m *MetricsTicketFunnelMetricsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricsTicketFunnelMetricsEntity) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MetricsTicketFunnelMetricsEntity) contextValidateGroupings(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *MetricsTicketFunnelMetricsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsTicketFunnelMetricsEntity) UnmarshalBinary(b []byte) error {
	var res MetricsTicketFunnelMetricsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}