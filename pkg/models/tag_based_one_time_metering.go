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

// TagBasedOneTimeMetering TagBasedOneTimeMetering
//
// swagger:model TagBasedOneTimeMetering
type TagBasedOneTimeMetering struct {

	// key
	Key string `json:"key,omitempty"`

	// one time metering
	OneTimeMetering *OneTimeMetering `json:"oneTimeMetering,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this tag based one time metering
func (m *TagBasedOneTimeMetering) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOneTimeMetering(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagBasedOneTimeMetering) validateOneTimeMetering(formats strfmt.Registry) error {
	if swag.IsZero(m.OneTimeMetering) { // not required
		return nil
	}

	if m.OneTimeMetering != nil {
		if err := m.OneTimeMetering.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oneTimeMetering")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oneTimeMetering")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tag based one time metering based on the context it is used
func (m *TagBasedOneTimeMetering) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOneTimeMetering(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TagBasedOneTimeMetering) contextValidateOneTimeMetering(ctx context.Context, formats strfmt.Registry) error {

	if m.OneTimeMetering != nil {
		if err := m.OneTimeMetering.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oneTimeMetering")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("oneTimeMetering")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TagBasedOneTimeMetering) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagBasedOneTimeMetering) UnmarshalBinary(b []byte) error {
	var res TagBasedOneTimeMetering
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
