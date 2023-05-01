// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsRetrospectiveEntityDataEntity metrics retrospective entity data entity
//
// swagger:model Metrics_RetrospectiveEntity_DataEntity
type MetricsRetrospectiveEntityDataEntity struct {

	// x
	X string `json:"x,omitempty"`

	// y
	Y float32 `json:"y,omitempty"`
}

// Validate validates this metrics retrospective entity data entity
func (m *MetricsRetrospectiveEntityDataEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics retrospective entity data entity based on context it is used
func (m *MetricsRetrospectiveEntityDataEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsRetrospectiveEntityDataEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsRetrospectiveEntityDataEntity) UnmarshalBinary(b []byte) error {
	var res MetricsRetrospectiveEntityDataEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}