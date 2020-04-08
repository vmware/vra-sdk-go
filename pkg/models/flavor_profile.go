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

// FlavorProfile Represents a structure that holds flavor mappings defined for the corresponding cloud end-point region.<br>**HATEOAS** links:<br>**region** - Region - Region for the profile.<br>**self** - FlavorProfile - Self link to this flavor profile
// swagger:model FlavorProfile
type FlavorProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// The id of the region for which this profile is defined
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// A list of the flavor mappings defined for the corresponding cloud end-point region
	// Required: true
	FlavorMappings *FlavorMapping `json:"flavorMappings"`

	// The id of this resource instance
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrgID string `json:"orgId,omitempty"`

	// This field is deprecated. Use orgId instead. The id of the organization this entity belongs to.
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	Owner string `json:"owner,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this flavor profile
func (m *FlavorProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlavorMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlavorProfile) validateLinks(formats strfmt.Registry) error {

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

func (m *FlavorProfile) validateFlavorMappings(formats strfmt.Registry) error {

	if err := validate.Required("flavorMappings", "body", m.FlavorMappings); err != nil {
		return err
	}

	if m.FlavorMappings != nil {
		if err := m.FlavorMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flavorMappings")
			}
			return err
		}
	}

	return nil
}

func (m *FlavorProfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlavorProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlavorProfile) UnmarshalBinary(b []byte) error {
	var res FlavorProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
