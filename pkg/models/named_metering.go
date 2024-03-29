// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NamedMetering NamedMetering
//
// swagger:model NamedMetering
type NamedMetering struct {

	// metering
	Metering *Metering `json:"metering,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this named metering
func (m *NamedMetering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetering(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedMetering) validateMetering(formats strfmt.Registry) error {
	if swag.IsZero(m.Metering) { // not required
		return nil
	}

	if m.Metering != nil {
		if err := m.Metering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metering")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metering")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this named metering based on the context it is used
func (m *NamedMetering) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedMetering) contextValidateMetering(ctx context.Context, formats strfmt.Registry) error {

	if m.Metering != nil {
		if err := m.Metering.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metering")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metering")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamedMetering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedMetering) UnmarshalBinary(b []byte) error {
	var res NamedMetering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
