// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Type Type
//
// swagger:model Type
type Type struct {

	// kind
	Kind string `json:"kind,omitempty"`

	// package
	Package string `json:"package,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this type
func (m *Type) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this type based on context it is used
func (m *Type) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Type) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Type) UnmarshalBinary(b []byte) error {
	var res Type
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
