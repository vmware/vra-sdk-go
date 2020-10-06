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

// AwsStorageProfile Defines a structure that holds list of storage policies defined for AWS for a specific region.**HATEOAS** links:<br>**region** - Region - Region for the profile.<br>**self** - AwsStorageProfile - Self link to this aws Storage Profile
// swagger:model AwsStorageProfile
type AwsStorageProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Id of the cloud account this storage profile belongs to.
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// Indicates whether this storage profile is default or not..
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Indicates the type of storage device.
	DeviceType string `json:"deviceType,omitempty"`

	// The id of the region for which this profile is defined
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Required: true
	ID *string `json:"id"`

	// Indicates maximum I/O operations per second in range(1-20,000).
	Iops string `json:"iops,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrgID string `json:"orgId,omitempty"`

	// This field is deprecated. Use orgId instead. The id of the organization this entity belongs to.
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	Owner string `json:"owner,omitempty"`

	// Indicates whether this storage profile supports encryption or not.
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`

	// Indicates the type of volume associated with type of storage device.
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this aws storage profile
func (m *AwsStorageProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsStorageProfile) validateLinks(formats strfmt.Registry) error {

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

func (m *AwsStorageProfile) validateDefaultItem(formats strfmt.Registry) error {

	if err := validate.Required("defaultItem", "body", m.DefaultItem); err != nil {
		return err
	}

	return nil
}

func (m *AwsStorageProfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AwsStorageProfile) validateTags(formats strfmt.Registry) error {

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
func (m *AwsStorageProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsStorageProfile) UnmarshalBinary(b []byte) error {
	var res AwsStorageProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
