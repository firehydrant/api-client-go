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

// PostV1ConversationsConversationIDCommentsCommentIDReactions ALPHA - Create a reaction on a comment
//
// swagger:model postV1ConversationsConversationIdCommentsCommentIdReactions
type PostV1ConversationsConversationIDCommentsCommentIDReactions struct {

	// CLDR short name of Unicode emojis
	// Required: true
	Reaction *string `json:"reaction"`
}

// Validate validates this post v1 conversations conversation Id comments comment Id reactions
func (m *PostV1ConversationsConversationIDCommentsCommentIDReactions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ConversationsConversationIDCommentsCommentIDReactions) validateReaction(formats strfmt.Registry) error {

	if err := validate.Required("reaction", "body", m.Reaction); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 conversations conversation Id comments comment Id reactions based on context it is used
func (m *PostV1ConversationsConversationIDCommentsCommentIDReactions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ConversationsConversationIDCommentsCommentIDReactions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ConversationsConversationIDCommentsCommentIDReactions) UnmarshalBinary(b []byte) error {
	var res PostV1ConversationsConversationIDCommentsCommentIDReactions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
