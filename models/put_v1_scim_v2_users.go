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

// PutV1ScimV2Users PUT SCIM endpoint to update a User. This endpoint is used to replace a resource's attributes.
//
// swagger:model putV1ScimV2Users
type PutV1ScimV2Users struct {

	// Boolean that represents whether user is active
	Active bool `json:"active,omitempty"`

	// Email addresses for the User
	Emails []*PutV1ScimV2UsersEmailsItems0 `json:"emails"`

	// name
	Name *PutV1ScimV2UsersName `json:"name,omitempty"`

	// Roles for the User
	Roles []string `json:"roles"`

	// A service provider's unique identifier for the user
	UserName string `json:"userName,omitempty"`
}

// Validate validates this put v1 scim v2 users
func (m *PutV1ScimV2Users) Validate(formats strfmt.Registry) error {
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

func (m *PutV1ScimV2Users) validateEmails(formats strfmt.Registry) error {
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

func (m *PutV1ScimV2Users) validateName(formats strfmt.Registry) error {
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

// ContextValidate validate this put v1 scim v2 users based on the context it is used
func (m *PutV1ScimV2Users) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PutV1ScimV2Users) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PutV1ScimV2Users) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PutV1ScimV2Users) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1ScimV2Users) UnmarshalBinary(b []byte) error {
	var res PutV1ScimV2Users
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1ScimV2UsersEmailsItems0 put v1 scim v2 users emails items0
//
// swagger:model PutV1ScimV2UsersEmailsItems0
type PutV1ScimV2UsersEmailsItems0 struct {

	// Boolean which signifies if an email is intended as the primary email for the User
	Primary bool `json:"primary,omitempty"`

	// String that represents an email address for the User
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this put v1 scim v2 users emails items0
func (m *PutV1ScimV2UsersEmailsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutV1ScimV2UsersEmailsItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 scim v2 users emails items0 based on context it is used
func (m *PutV1ScimV2UsersEmailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1ScimV2UsersEmailsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1ScimV2UsersEmailsItems0) UnmarshalBinary(b []byte) error {
	var res PutV1ScimV2UsersEmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutV1ScimV2UsersName The components of the user's name
//
// swagger:model PutV1ScimV2UsersName
type PutV1ScimV2UsersName struct {

	// The given name of the User, or first name in most Western languages
	// Required: true
	FamilyName *string `json:"familyName"`

	// The family name of the User, or last name in most Western languages
	// Required: true
	GivenName *string `json:"givenName"`
}

// Validate validates this put v1 scim v2 users name
func (m *PutV1ScimV2UsersName) Validate(formats strfmt.Registry) error {
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

func (m *PutV1ScimV2UsersName) validateFamilyName(formats strfmt.Registry) error {

	if err := validate.Required("name"+"."+"familyName", "body", m.FamilyName); err != nil {
		return err
	}

	return nil
}

func (m *PutV1ScimV2UsersName) validateGivenName(formats strfmt.Registry) error {

	if err := validate.Required("name"+"."+"givenName", "body", m.GivenName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put v1 scim v2 users name based on context it is used
func (m *PutV1ScimV2UsersName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1ScimV2UsersName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1ScimV2UsersName) UnmarshalBinary(b []byte) error {
	var res PutV1ScimV2UsersName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
