// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Module module
// swagger:model Module
type Module struct {

	// annotations
	Annotations []Annotation `json:"annotations"`

	// class loader
	ClassLoader *ClassLoader `json:"classLoader,omitempty"`

	// declared annotations
	DeclaredAnnotations []Annotation `json:"declaredAnnotations"`

	// descriptor
	Descriptor *ModuleDescriptor `json:"descriptor,omitempty"`

	// layer
	Layer ModuleLayer `json:"layer,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// named
	Named bool `json:"named,omitempty"`

	// packages
	// Unique: true
	Packages []string `json:"packages"`
}

// Validate validates this module
func (m *Module) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassLoader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Module) validateClassLoader(formats strfmt.Registry) error {

	if swag.IsZero(m.ClassLoader) { // not required
		return nil
	}

	if m.ClassLoader != nil {
		if err := m.ClassLoader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classLoader")
			}
			return err
		}
	}

	return nil
}

func (m *Module) validateDescriptor(formats strfmt.Registry) error {

	if swag.IsZero(m.Descriptor) { // not required
		return nil
	}

	if m.Descriptor != nil {
		if err := m.Descriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descriptor")
			}
			return err
		}
	}

	return nil
}

func (m *Module) validatePackages(formats strfmt.Registry) error {

	if swag.IsZero(m.Packages) { // not required
		return nil
	}

	if err := validate.UniqueItems("packages", "body", m.Packages); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Module) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Module) UnmarshalBinary(b []byte) error {
	var res Module
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}