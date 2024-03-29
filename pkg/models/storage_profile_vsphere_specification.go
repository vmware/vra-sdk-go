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

// StorageProfileVsphereSpecification Represents a specification of vSphere storage profile.
//
// swagger:model StorageProfileVsphereSpecification
type StorageProfileVsphereSpecification struct {

	// Id of the vSphere Datastore for placing disk and VM.
	// Example: 08d28
	DatastoreID string `json:"datastoreId,omitempty"`

	// Indicates if a storage profile acts as a default storage profile for a disk.
	// Example: true
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Type of mode for the disk
	// Example: undefined / independent-persistent / independent-nonpersistent
	DiskMode string `json:"diskMode,omitempty"`

	// Disk types are specified as
	//
	//  	Standard - Simple vSphere virtual disks which cannot be managed independently without an attached VM.
	// 	First Class - Improved version of standard virtual disks, designed to be fully mananged independent storage objects.
	//
	// Empty value is considered as Standard
	// Example: standard / firstClass
	DiskType string `json:"diskType,omitempty"`

	// The upper bound for the I/O operations per second allocated for each virtual disk.
	// Example: 1000
	LimitIops string `json:"limitIops,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Type of provisioning policy for the disk.
	// Example: thin / thick / eagerZeroedThick
	ProvisioningType string `json:"provisioningType,omitempty"`

	// The Id of the region that is associated with the storage profile.
	// Example: 31186
	// Required: true
	RegionID *string `json:"regionId"`

	// A specific number of shares assigned to each virtual machine.
	// Example: 2000
	Shares string `json:"shares,omitempty"`

	// Shares are specified as High, Normal, Low or Custom and these values specify share values with a 4:2:1 ratio, respectively.
	// Example: low / normal / high / custom
	SharesLevel string `json:"sharesLevel,omitempty"`

	// Id of the vSphere Storage Policy to be applied.
	// Example: 6b59743af31d
	StoragePolicyID string `json:"storagePolicyId,omitempty"`

	// Indicates whether this storage profile supports encryption or not.
	// Example: false
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile.
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this storage profile vsphere specification
func (m *StorageProfileVsphereSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
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

func (m *StorageProfileVsphereSpecification) validateDefaultItem(formats strfmt.Registry) error {

	if err := validate.Required("defaultItem", "body", m.DefaultItem); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileVsphereSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileVsphereSpecification) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileVsphereSpecification) validateTags(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this storage profile vsphere specification based on the context it is used
func (m *StorageProfileVsphereSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageProfileVsphereSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageProfileVsphereSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageProfileVsphereSpecification) UnmarshalBinary(b []byte) error {
	var res StorageProfileVsphereSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
