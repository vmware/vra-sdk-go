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

// UpdateCloudAccountVsphereSpecification update cloud account vsphere specification
// swagger:model UpdateCloudAccountVsphereSpecification
type UpdateCloudAccountVsphereSpecification struct {

	// NSX-V or NSX-T account to associate with this vSphere cloud account. vSphere cloud account can be a single NSX-V cloud account or a single NSX-T cloud account.
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Create default cloud zones for the enabled regions.
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// Password to authenticate with the cloud account
	Password string `json:"password,omitempty"`

	// A set of Region names to enable provisioning on.
	RegionIds []string `json:"regionIds"`

	// A set of tag keys and optional values to set on the Cloud Account
	Tags []*Tag `json:"tags"`

	// Username to authenticate with the cloud account
	Username string `json:"username,omitempty"`
}

// Validate validates this update cloud account vsphere specification
func (m *UpdateCloudAccountVsphereSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCloudAccountVsphereSpecification) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCloudAccountVsphereSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCloudAccountVsphereSpecification) UnmarshalBinary(b []byte) error {
	var res UpdateCloudAccountVsphereSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
