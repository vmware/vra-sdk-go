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

// ApplicationTier ApplicationTier
//
// Application tier.
//
// swagger:model ApplicationTier
type ApplicationTier struct {

	// A description of the tier
	Description string `json:"description,omitempty"`

	// Unique identifier of the tier
	ID string `json:"id,omitempty"`

	// Name of the tier
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this application tier
func (m *ApplicationTier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationTier) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this application tier based on context it is used
func (m *ApplicationTier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationTier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationTier) UnmarshalBinary(b []byte) error {
	var res ApplicationTier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
