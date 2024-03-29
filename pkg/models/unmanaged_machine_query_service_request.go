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

// UnmanagedMachineQueryServiceRequest unmanaged machine query service request
//
// swagger:model UnmanagedMachineQueryServiceRequest
type UnmanagedMachineQueryServiceRequest struct {

	// expand fields
	ExpandFields []string `json:"expandFields"`

	// filters
	Filters []*ResourceFieldFilter `json:"filters"`

	// next page link
	NextPageLink string `json:"nextPageLink,omitempty"`

	// option exclude plan machines
	OptionExcludePlanMachines bool `json:"optionExcludePlanMachines,omitempty"`

	// option include only plan machines
	OptionIncludeOnlyPlanMachines bool `json:"optionIncludeOnlyPlanMachines,omitempty"`

	// page result limit
	PageResultLimit int64 `json:"pageResultLimit,omitempty"`

	// plan link
	PlanLink string `json:"planLink,omitempty"`
}

// Validate validates this unmanaged machine query service request
func (m *UnmanagedMachineQueryServiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnmanagedMachineQueryServiceRequest) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unmanaged machine query service request based on the context it is used
func (m *UnmanagedMachineQueryServiceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UnmanagedMachineQueryServiceRequest) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UnmanagedMachineQueryServiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UnmanagedMachineQueryServiceRequest) UnmarshalBinary(b []byte) error {
	var res UnmanagedMachineQueryServiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
