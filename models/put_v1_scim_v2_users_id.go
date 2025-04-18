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

// PutV1ScimV2UsersID PUT SCIM endpoint to update a User. This endpoint is used to replace a resource's attributes.
//
// swagger:model putV1ScimV2UsersId
type PutV1ScimV2UsersID struct {

	// Boolean that represents whether user is active
	Active bool `json:"active,omitempty"`

	// Email addresses for the User
	Emails []*PutV1ScimV2UsersIDEmailsItems0 `json:"emails"`

	// name
	Name *PutV1ScimV2UsersIDName `json:"name,omitempty"`

	// Roles for the User. Options are owner, member, collaborator, or viewer. Roles may be specified as strings or SCIM role objects.
	Roles interface{} `json:"roles,omitempty"`

	// A service provider's unique identifier for the user
	UserName string `json:"userName,omitempty"`
}

// Validate validates this put v1 scim v2 users Id
func (m *PutV1ScimV2UsersID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1ScimV2UsersID) validateEmails(formats strfmt.Registry) error {
	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1ScimV2UsersID) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put v1 scim v2 users Id based on the context it is used
func (m *PutV1ScimV2UsersID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1ScimV2UsersID) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Emails); i++ {

		if m.Emails[i] != nil {
			if err := m.Emails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutV1ScimV2UsersID) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {
		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutV1ScimV2UsersID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1ScimV2UsersID) UnmarshalBinary(b []byte) error {
	var res PutV1ScimV2UsersID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1ScimV2UsersIDEmailsItems0 put v1 scim v2 users ID emails items0
//
// swagger:model PutV1ScimV2UsersIDEmailsItems0
type PutV1ScimV2UsersIDEmailsItems0 struct {

	// Boolean which signifies if an email is intended as the primary email for the User
	Primary bool `json:"primary,omitempty"`

	// String that represents an email address for the User
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this put v1 scim v2 users ID emails items0
func (m *PutV1ScimV2UsersIDEmailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1ScimV2UsersIDEmailsItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 scim v2 users ID emails items0 based on context it is used
func (m *PutV1ScimV2UsersIDEmailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1ScimV2UsersIDEmailsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1ScimV2UsersIDEmailsItems0) UnmarshalBinary(b []byte) error {
	var res PutV1ScimV2UsersIDEmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1ScimV2UsersIDName The components of the user's name
//
// swagger:model PutV1ScimV2UsersIDName
type PutV1ScimV2UsersIDName struct {

	// The given name of the User, or first name in most Western languages
	// Required: true
	FamilyName *string `json:"familyName"`

	// The family name of the User, or last name in most Western languages
	// Required: true
	GivenName *string `json:"givenName"`
}

// Validate validates this put v1 scim v2 users ID name
func (m *PutV1ScimV2UsersIDName) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFamilyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGivenName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1ScimV2UsersIDName) validateFamilyName(formats strfmt.Registry) error {

	if err := validate.Required("name"+"."+"familyName", "body", m.FamilyName); err != nil {
		return err
	}

	return nil
}

func (m *PutV1ScimV2UsersIDName) validateGivenName(formats strfmt.Registry) error {

	if err := validate.Required("name"+"."+"givenName", "body", m.GivenName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 scim v2 users ID name based on context it is used
func (m *PutV1ScimV2UsersIDName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1ScimV2UsersIDName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1ScimV2UsersIDName) UnmarshalBinary(b []byte) error {
	var res PutV1ScimV2UsersIDName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
