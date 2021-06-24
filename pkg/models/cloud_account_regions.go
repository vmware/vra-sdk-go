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

// CloudAccountRegions State object representing cloud account region.<br><br>**externalRegions** - array[RegionSpecification] - Set of regions that can be enabled for this cloud account.<br>**externalRegionIds** - array[String] - Set of ids of regions that can be enabled for this cloud account.<br>
//
// swagger:model CloudAccountRegions
type CloudAccountRegions struct {

	// A set of ids of regions that can be enabled for this cloud account. Deprecated - use externalRegions to obtain enumerated regions.
	// Example: [ \"us-east-1\", \"ap-northeast-1\" ]
	ExternalRegionIds []string `json:"externalRegionIds"`

	// A set of regions that can be enabled for this cloud account.
	ExternalRegions []*RegionSpecification `json:"externalRegions"`
}

// Validate validates this cloud account regions
func (m *CloudAccountRegions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalRegions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountRegions) validateExternalRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalRegions) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalRegions); i++ {
		if swag.IsZero(m.ExternalRegions[i]) { // not required
			continue
		}

		if m.ExternalRegions[i] != nil {
			if err := m.ExternalRegions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalRegions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cloud account regions based on the context it is used
func (m *CloudAccountRegions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalRegions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountRegions) contextValidateExternalRegions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalRegions); i++ {

		if m.ExternalRegions[i] != nil {
			if err := m.ExternalRegions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalRegions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountRegions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountRegions) UnmarshalBinary(b []byte) error {
	var res CloudAccountRegions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
