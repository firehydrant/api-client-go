// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1CatalogsCatalogIDIngest Accepts catalog data in the configured format and asyncronously processes the data to incorporate changes into service catalog.
//
// swagger:model postV1CatalogsCatalogIdIngest
type PostV1CatalogsCatalogIDIngest struct {

	// Service data
	// Required: true
	Data *string `json:"data"`

	// Encoding for submitted data
	// Required: true
	// Enum: [text/yaml application/x-yaml application/json]
	Encoding *string `json:"encoding"`
}

// Validate validates this post v1 catalogs catalog Id ingest
func (m *PostV1CatalogsCatalogIDIngest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncoding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1CatalogsCatalogIDIngest) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

var postV1CatalogsCatalogIdIngestTypeEncodingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text/yaml","application/x-yaml","application/json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1CatalogsCatalogIdIngestTypeEncodingPropEnum = append(postV1CatalogsCatalogIdIngestTypeEncodingPropEnum, v)
	}
}

const (

	// PostV1CatalogsCatalogIDIngestEncodingTextYaml captures enum value "text/yaml"
	PostV1CatalogsCatalogIDIngestEncodingTextYaml string = "text/yaml"

	// PostV1CatalogsCatalogIDIngestEncodingApplicationxDashYaml captures enum value "application/x-yaml"
	PostV1CatalogsCatalogIDIngestEncodingApplicationxDashYaml string = "application/x-yaml"

	// PostV1CatalogsCatalogIDIngestEncodingApplicationJSON captures enum value "application/json"
	PostV1CatalogsCatalogIDIngestEncodingApplicationJSON string = "application/json"
)

// prop value enum
func (m *PostV1CatalogsCatalogIDIngest) validateEncodingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1CatalogsCatalogIdIngestTypeEncodingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1CatalogsCatalogIDIngest) validateEncoding(formats strfmt.Registry) error {

	if err := validate.Required("encoding", "body", m.Encoding); err != nil {
		return err
	}

	// value enum
	if err := m.validateEncodingEnum("encoding", "body", *m.Encoding); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 catalogs catalog Id ingest based on context it is used
func (m *PostV1CatalogsCatalogIDIngest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1CatalogsCatalogIDIngest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1CatalogsCatalogIDIngest) UnmarshalBinary(b []byte) error {
	var res PostV1CatalogsCatalogIDIngest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}