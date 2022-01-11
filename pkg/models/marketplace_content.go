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

// MarketplaceContent MarketplaceContent
//
// swagger:model MarketplaceContent
type MarketplaceContent struct {

	// Description
	Description string `json:"description,omitempty"`

	// Developer
	Developer string `json:"developer,omitempty"`

	// Display Name
	DisplayName string `json:"displayName,omitempty"`

	// Documents
	Documents []*Document `json:"documents"`

	// EULA
	Eula string `json:"eula,omitempty"`

	// Highlights
	Highlights []string `json:"highlights"`

	// Icon
	Icon string `json:"icon,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// Overview
	Overview string `json:"overview,omitempty"`

	// Rating
	Rating string `json:"rating,omitempty"`

	// Screenshots
	Screenshots []string `json:"screenshots"`

	// Short Name
	ShortName string `json:"shortName,omitempty"`

	// support
	Support *Support `json:"support,omitempty"`

	// Tech Specs
	TechSpecs string `json:"techSpecs,omitempty"`

	// Total Reviews
	TotalReviews int32 `json:"totalReviews,omitempty"`

	// Updated Date
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Version
	Version string `json:"version,omitempty"`

	// Video
	Video string `json:"video,omitempty"`
}

// Validate validates this marketplace content
func (m *MarketplaceContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceContent) validateDocuments(formats strfmt.Registry) error {
	if swag.IsZero(m.Documents) { // not required
		return nil
	}

	for i := 0; i < len(m.Documents); i++ {
		if swag.IsZero(m.Documents[i]) { // not required
			continue
		}

		if m.Documents[i] != nil {
			if err := m.Documents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("documents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MarketplaceContent) validateSupport(formats strfmt.Registry) error {
	if swag.IsZero(m.Support) { // not required
		return nil
	}

	if m.Support != nil {
		if err := m.Support.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("support")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("support")
			}
			return err
		}
	}

	return nil
}

func (m *MarketplaceContent) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this marketplace content based on the context it is used
func (m *MarketplaceContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocuments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupport(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceContent) contextValidateDocuments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Documents); i++ {

		if m.Documents[i] != nil {
			if err := m.Documents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("documents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("documents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MarketplaceContent) contextValidateSupport(ctx context.Context, formats strfmt.Registry) error {

	if m.Support != nil {
		if err := m.Support.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("support")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("support")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceContent) UnmarshalBinary(b []byte) error {
	var res MarketplaceContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
