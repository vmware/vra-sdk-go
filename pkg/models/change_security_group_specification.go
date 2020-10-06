// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeSecurityGroupSpecification Specification for a second day change security groups operation for a vsphere machine
// swagger:model ChangeSecurityGroupSpecification
type ChangeSecurityGroupSpecification struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// The id of this resource instance
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// A set of network interface controller specifications for this machine. If not specified, then no reconfiguration will be performed.
	NetworkInterfaceSpecifications []*NetworkInterfaceSpecification `json:"networkInterfaceSpecifications"`

	// The id of the organization this entity belongs to.
	OrgID string `json:"orgId,omitempty"`

	// This field is deprecated. Use orgId instead. The id of the organization this entity belongs to.
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	Owner string `json:"owner,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this change security group specification
func (m *ChangeSecurityGroupSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkInterfaceSpecifications(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeSecurityGroupSpecification) validateLinks(formats strfmt.Registry) error {

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ChangeSecurityGroupSpecification) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ChangeSecurityGroupSpecification) validateNetworkInterfaceSpecifications(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkInterfaceSpecifications) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkInterfaceSpecifications); i++ {
		if swag.IsZero(m.NetworkInterfaceSpecifications[i]) { // not required
			continue
		}

		if m.NetworkInterfaceSpecifications[i] != nil {
			if err := m.NetworkInterfaceSpecifications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkInterfaceSpecifications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeSecurityGroupSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeSecurityGroupSpecification) UnmarshalBinary(b []byte) error {
	var res ChangeSecurityGroupSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
