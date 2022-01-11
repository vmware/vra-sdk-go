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

// CloudAccountNsxVResult State object representing a query result of Nsx-V cloud accounts.
//
// swagger:model CloudAccountNsxVResult
type CloudAccountNsxVResult struct {

	// List of content items
	// Read Only: true
	Content []*CloudAccountNsxV `json:"content"`

	// Number of elements in the current page
	// Example: 1
	// Read Only: true
	NumberOfElements int64 `json:"numberOfElements,omitempty"`

	// Total number of elements. In some cases the field may not be populated
	// Example: 1
	// Read Only: true
	TotalElements int64 `json:"totalElements,omitempty"`
}

// Validate validates this cloud account nsx v result
func (m *CloudAccountNsxVResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountNsxVResult) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cloud account nsx v result based on the context it is used
func (m *CloudAccountNsxVResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumberOfElements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalElements(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountNsxVResult) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "content", "body", []*CloudAccountNsxV(m.Content)); err != nil {
		return err
	}

	for i := 0; i < len(m.Content); i++ {

		if m.Content[i] != nil {
			if err := m.Content[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountNsxVResult) contextValidateNumberOfElements(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "numberOfElements", "body", int64(m.NumberOfElements)); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountNsxVResult) contextValidateTotalElements(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "totalElements", "body", int64(m.TotalElements)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountNsxVResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountNsxVResult) UnmarshalBinary(b []byte) error {
	var res CloudAccountNsxVResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
