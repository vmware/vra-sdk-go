// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FilterSchema FilterSchema
// swagger:model FilterSchema
type FilterSchema struct {

	// filters
	Filters []*Filter `json:"filters"`
}

// Validate validates this filter schema
func (m *FilterSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilterSchema) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		// hello slicevalidator
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilterSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterSchema) UnmarshalBinary(b []byte) error {
	var res FilterSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
