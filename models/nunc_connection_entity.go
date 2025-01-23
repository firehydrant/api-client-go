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

// NuncConnectionEntity NuncConnectionEntity model
//
// swagger:model NuncConnectionEntity
type NuncConnectionEntity struct {

	// button background color
	ButtonBackgroundColor string `json:"button_background_color,omitempty"`

	// button text color
	ButtonTextColor string `json:"button_text_color,omitempty"`

	// cname
	Cname string `json:"cname,omitempty"`

	// company name
	CompanyName string `json:"company_name,omitempty"`

	// company tos url
	CompanyTosURL string `json:"company_tos_url,omitempty"`

	// company website
	CompanyWebsite string `json:"company_website,omitempty"`

	// component groups
	ComponentGroups *NuncComponentGroupEntity `json:"component_groups,omitempty"`

	// components
	Components *NuncComponentEntity `json:"components,omitempty"`

	// conditions
	Conditions *NuncConditionEntity `json:"conditions,omitempty"`

	// cover image
	CoverImage *MediaImageEntity `json:"cover_image,omitempty"`

	// dark logo
	DarkLogo *MediaImageEntity `json:"dark_logo,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// enable histogram
	EnableHistogram bool `json:"enable_histogram,omitempty"`

	// exposed fields
	ExposedFields []string `json:"exposed_fields"`

	// favicon
	Favicon *MediaImageEntity `json:"favicon,omitempty"`

	// greeting body
	GreetingBody string `json:"greeting_body,omitempty"`

	// greeting title
	GreetingTitle string `json:"greeting_title,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// link color
	LinkColor string `json:"link_color,omitempty"`

	// List of links attached to this status page.
	Links []*LinksEntity `json:"links"`

	// logo
	Logo *MediaImageEntity `json:"logo,omitempty"`

	// open graph image
	OpenGraphImage *MediaImageEntity `json:"open_graph_image,omitempty"`

	// operational message
	OperationalMessage string `json:"operational_message,omitempty"`

	// primary color
	PrimaryColor string `json:"primary_color,omitempty"`

	// secondary color
	SecondaryColor string `json:"secondary_color,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// ui version
	UIVersion int32 `json:"ui_version,omitempty"`
}

// Validate validates this nunc connection entity
func (m *NuncConnectionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponentGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoverImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDarkLogo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFavicon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenGraphImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NuncConnectionEntity) validateComponentGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ComponentGroups) { // not required
		return nil
	}

	if m.ComponentGroups != nil {
		if err := m.ComponentGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component_groups")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateComponents(formats strfmt.Registry) error {
	if swag.IsZero(m.Components) { // not required
		return nil
	}

	if m.Components != nil {
		if err := m.Components.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateCoverImage(formats strfmt.Registry) error {
	if swag.IsZero(m.CoverImage) { // not required
		return nil
	}

	if m.CoverImage != nil {
		if err := m.CoverImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cover_image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cover_image")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateDarkLogo(formats strfmt.Registry) error {
	if swag.IsZero(m.DarkLogo) { // not required
		return nil
	}

	if m.DarkLogo != nil {
		if err := m.DarkLogo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dark_logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dark_logo")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateFavicon(formats strfmt.Registry) error {
	if swag.IsZero(m.Favicon) { // not required
		return nil
	}

	if m.Favicon != nil {
		if err := m.Favicon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("favicon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("favicon")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateLinks(formats strfmt.Registry) error {
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

func (m *NuncConnectionEntity) validateLogo(formats strfmt.Registry) error {
	if swag.IsZero(m.Logo) { // not required
		return nil
	}

	if m.Logo != nil {
		if err := m.Logo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) validateOpenGraphImage(formats strfmt.Registry) error {
	if swag.IsZero(m.OpenGraphImage) { // not required
		return nil
	}

	if m.OpenGraphImage != nil {
		if err := m.OpenGraphImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("open_graph_image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("open_graph_image")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nunc connection entity based on the context it is used
func (m *NuncConnectionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComponentGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCoverImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDarkLogo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFavicon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenGraphImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NuncConnectionEntity) contextValidateComponentGroups(ctx context.Context, formats strfmt.Registry) error {

	if m.ComponentGroups != nil {
		if err := m.ComponentGroups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component_groups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component_groups")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateComponents(ctx context.Context, formats strfmt.Registry) error {

	if m.Components != nil {
		if err := m.Components.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("components")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("components")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	if m.Conditions != nil {
		if err := m.Conditions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateCoverImage(ctx context.Context, formats strfmt.Registry) error {

	if m.CoverImage != nil {
		if err := m.CoverImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cover_image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cover_image")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateDarkLogo(ctx context.Context, formats strfmt.Registry) error {

	if m.DarkLogo != nil {
		if err := m.DarkLogo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dark_logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dark_logo")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateFavicon(ctx context.Context, formats strfmt.Registry) error {

	if m.Favicon != nil {
		if err := m.Favicon.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("favicon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("favicon")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NuncConnectionEntity) contextValidateLogo(ctx context.Context, formats strfmt.Registry) error {

	if m.Logo != nil {
		if err := m.Logo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logo")
			}
			return err
		}
	}

	return nil
}

func (m *NuncConnectionEntity) contextValidateOpenGraphImage(ctx context.Context, formats strfmt.Registry) error {

	if m.OpenGraphImage != nil {
		if err := m.OpenGraphImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("open_graph_image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("open_graph_image")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NuncConnectionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NuncConnectionEntity) UnmarshalBinary(b []byte) error {
	var res NuncConnectionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
