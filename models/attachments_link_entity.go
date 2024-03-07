// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AttachmentsLinkEntity Attachments_LinkEntity model
//
// swagger:model Attachments_LinkEntity
type AttachmentsLinkEntity struct {

	// Link can be deleted
	Deletable bool `json:"deletable,omitempty"`

	// display text
	DisplayText string `json:"display_text,omitempty"`

	// Link can be edited
	Editable bool `json:"editable,omitempty"`

	// href url
	HrefURL string `json:"href_url,omitempty"`

	// icon url
	IconURL string `json:"icon_url,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this attachments link entity
func (m *AttachmentsLinkEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this attachments link entity based on context it is used
func (m *AttachmentsLinkEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttachmentsLinkEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachmentsLinkEntity) UnmarshalBinary(b []byte) error {
	var res AttachmentsLinkEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
