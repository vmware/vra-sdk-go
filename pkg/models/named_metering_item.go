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

// NamedMeteringItem NamedMeteringItem
//
// swagger:model NamedMeteringItem
type NamedMeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// named meterings
	NamedMeterings []*NamedMetering `json:"namedMeterings"`
}

// Validate validates this named metering item
func (m *NamedMeteringItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamedMeterings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedMeteringItem) validateNamedMeterings(formats strfmt.Registry) error {
	if swag.IsZero(m.NamedMeterings) { // not required
		return nil
	}

	for i := 0; i < len(m.NamedMeterings); i++ {
		if swag.IsZero(m.NamedMeterings[i]) { // not required
			continue
		}

		if m.NamedMeterings[i] != nil {
			if err := m.NamedMeterings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namedMeterings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("namedMeterings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this named metering item based on the context it is used
func (m *NamedMeteringItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNamedMeterings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedMeteringItem) contextValidateNamedMeterings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NamedMeterings); i++ {

		if m.NamedMeterings[i] != nil {
			if err := m.NamedMeterings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namedMeterings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("namedMeterings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamedMeteringItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedMeteringItem) UnmarshalBinary(b []byte) error {
	var res NamedMeteringItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
