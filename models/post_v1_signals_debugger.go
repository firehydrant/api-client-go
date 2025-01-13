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
	"github.com/go-openapi/validate"
)

// PostV1SignalsDebugger post v1 signals debugger
//
// swagger:model postV1SignalsDebugger
type PostV1SignalsDebugger struct {

	// CEL expression
	// Required: true
	Expression *string `json:"expression"`

	// List of signals to evaluate rule expression against
	// Required: true
	Signals []*PostV1SignalsDebuggerSignalsItems0 `json:"signals"`
}

// Validate validates this post v1 signals debugger
func (m *PostV1SignalsDebugger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SignalsDebugger) validateExpression(formats strfmt.Registry) error {

	if err := validate.Required("expression", "body", m.Expression); err != nil {
		return err
	}

	return nil
}

func (m *PostV1SignalsDebugger) validateSignals(formats strfmt.Registry) error {

	if err := validate.Required("signals", "body", m.Signals); err != nil {
		return err
	}

	for i := 0; i < len(m.Signals); i++ {
		if swag.IsZero(m.Signals[i]) { // not required
			continue
		}

		if m.Signals[i] != nil {
			if err := m.Signals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post v1 signals debugger based on the context it is used
func (m *PostV1SignalsDebugger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSignals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SignalsDebugger) contextValidateSignals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signals); i++ {

		if m.Signals[i] != nil {
			if err := m.Signals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1SignalsDebugger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SignalsDebugger) UnmarshalBinary(b []byte) error {
	var res PostV1SignalsDebugger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1SignalsDebuggerSignalsItems0 post v1 signals debugger signals items0
//
// swagger:model PostV1SignalsDebuggerSignalsItems0
type PostV1SignalsDebuggerSignalsItems0 struct {

	// annotations
	Annotations interface{} `json:"annotations,omitempty"`

	// body
	Body string `json:"body,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// images
	Images []*PostV1SignalsDebuggerSignalsItems0ImagesItems0 `json:"images"`

	// level
	Level string `json:"level,omitempty"`

	// links
	Links []*PostV1SignalsDebuggerSignalsItems0LinksItems0 `json:"links"`

	// organization id
	OrganizationID string `json:"organization_id,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this post v1 signals debugger signals items0
func (m *PostV1SignalsDebuggerSignalsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SignalsDebuggerSignalsItems0) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1SignalsDebuggerSignalsItems0) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post v1 signals debugger signals items0 based on the context it is used
func (m *PostV1SignalsDebuggerSignalsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1SignalsDebuggerSignalsItems0) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {
			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostV1SignalsDebuggerSignalsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1SignalsDebuggerSignalsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SignalsDebuggerSignalsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1SignalsDebuggerSignalsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1SignalsDebuggerSignalsItems0ImagesItems0 post v1 signals debugger signals items0 images items0
//
// swagger:model PostV1SignalsDebuggerSignalsItems0ImagesItems0
type PostV1SignalsDebuggerSignalsItems0ImagesItems0 struct {

	// alt
	Alt string `json:"alt,omitempty"`

	// src
	Src string `json:"src,omitempty"`
}

// Validate validates this post v1 signals debugger signals items0 images items0
func (m *PostV1SignalsDebuggerSignalsItems0ImagesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post v1 signals debugger signals items0 images items0 based on context it is used
func (m *PostV1SignalsDebuggerSignalsItems0ImagesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1SignalsDebuggerSignalsItems0ImagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SignalsDebuggerSignalsItems0ImagesItems0) UnmarshalBinary(b []byte) error {
	var res PostV1SignalsDebuggerSignalsItems0ImagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1SignalsDebuggerSignalsItems0LinksItems0 post v1 signals debugger signals items0 links items0
//
// swagger:model PostV1SignalsDebuggerSignalsItems0LinksItems0
type PostV1SignalsDebuggerSignalsItems0LinksItems0 struct {

	// href
	Href string `json:"href,omitempty"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this post v1 signals debugger signals items0 links items0
func (m *PostV1SignalsDebuggerSignalsItems0LinksItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post v1 signals debugger signals items0 links items0 based on context it is used
func (m *PostV1SignalsDebuggerSignalsItems0LinksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1SignalsDebuggerSignalsItems0LinksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1SignalsDebuggerSignalsItems0LinksItems0) UnmarshalBinary(b []byte) error {
	var res PostV1SignalsDebuggerSignalsItems0LinksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
