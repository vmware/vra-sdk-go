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

// VsphereStorageProfile Defines a structure that holds storage profile details defined for vSphere for a specific region.**HATEOAS** links:<br>**datastore** - FabricVsphereDatastore - Datastore for this storage profile.<br>**storage-policy** - FabricVsphereStoragePolicy - vSphere storage policy for this storage profile.<br> **region** - Region - Region for the profile.<br>**self** - VsphereStorageProfile - Self link to this vSphere storage profile.
// swagger:model VsphereStorageProfile
type VsphereStorageProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// Indicates if a storage profile contains default storage properties.
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Type of mode for the disk
	DiskMode string `json:"diskMode,omitempty"`

	// The id of the region for which this profile is defined
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Required: true
	ID *string `json:"id"`

	// The upper bound for the I/O operations per second allocated for each disk.
	LimitIops string `json:"limitIops,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrgID string `json:"orgId,omitempty"`

	// This field is deprecated. Use orgId instead. The id of the organization this entity belongs to.
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	Owner string `json:"owner,omitempty"`

	// Type of format for the disk.
	ProvisioningType string `json:"provisioningType,omitempty"`

	// A specific number of shares assigned to each virtual machine.
	Shares string `json:"shares,omitempty"`

	// Shares level are specified as High, Normal, Low or Custom.
	SharesLevel string `json:"sharesLevel,omitempty"`

	// Indicates whether this storage profile should support encryption or not.
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this vsphere storage profile
func (m *VsphereStorageProfile) Validate(formats strfmt.Registry) error {
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

func (m *VsphereStorageProfile) validateLinks(formats strfmt.Registry) error {

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

func (m *VsphereStorageProfile) validateDefaultItem(formats strfmt.Registry) error {

	if err := validate.Required("defaultItem", "body", m.DefaultItem); err != nil {
		return err
	}

	return nil
}

func (m *VsphereStorageProfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VsphereStorageProfile) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		// hello slicevalidator
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
func (m *VsphereStorageProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsphereStorageProfile) UnmarshalBinary(b []byte) error {
	var res VsphereStorageProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
