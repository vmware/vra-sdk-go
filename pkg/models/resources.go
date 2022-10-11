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

// Resources Resources
//
// swagger:model Resources
type Resources struct {

	// limits
	Limits *Limits `json:"limits,omitempty"`

	// requests
	Requests *Requests `json:"requests,omitempty"`
}

// Validate validates this resources
func (m *Resources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resources) validateLimits(formats strfmt.Registry) error {
	if swag.IsZero(m.Limits) { // not required
		return nil
	}

	if m.Limits != nil {
		if err := m.Limits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limits")
			}
			return err
		}
	}

	return nil
}

func (m *Resources) validateRequests(formats strfmt.Registry) error {
	if swag.IsZero(m.Requests) { // not required
		return nil
	}

	if m.Requests != nil {
		if err := m.Requests.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requests")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requests")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this resources based on the context it is used
func (m *Resources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resources) contextValidateLimits(ctx context.Context, formats strfmt.Registry) error {

	if m.Limits != nil {
		if err := m.Limits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limits")
			}
			return err
		}
	}

	return nil
}

func (m *Resources) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

	if m.Requests != nil {
		if err := m.Requests.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requests")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requests")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resources) UnmarshalBinary(b []byte) error {
	var res Resources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
