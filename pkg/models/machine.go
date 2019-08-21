// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Machine Represents a cloud agnostic machine.<br>**HATEOAS** links:<br>**operations** - array[String] - Supported operations for the machine.<br>**network-interfaces** - array[NetworkInterface] - Network interfaces for the machine.<br>**disks** - array[MachineDisk] - disks for the machine.<br>**deployment** - Deployment - Deployment that this machine is part of.<br>**cloud-accounts** - array[CloudAccount] - Cloud accounts where this machine is provisioned.<br>**self** - Machine - Self link to this machine
// swagger:model Machine
type Machine struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
	Address string `json:"address,omitempty"`

	// Set of ids of the cloud accounts this entity belongs to.
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// External entity Id on the provider side.
	ExternalID string `json:"externalId,omitempty"`

	// The external regionId of the resource
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// The external zoneId of the resource.
	// Required: true
	ExternalZoneID *string `json:"externalZoneId"`

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

	// Power state of machine.
	// Required: true
	// Enum: [ON OFF GUEST_OFF UNKNOWN SUSPEND]
	PowerState *string `json:"powerState"`

	// The id of the project this entity belongs to.
	ProjectID string `json:"projectId,omitempty"`

	// A set of tag keys and optional values that were set on this resource instance.
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this machine
func (m *Machine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerState(formats); err != nil {
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

func (m *Machine) validateLinks(formats strfmt.Registry) error {

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

func (m *Machine) validateCloudAccountIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudAccountIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("cloudAccountIds", "body", m.CloudAccountIds); err != nil {
		return err
	}

	return nil
}

func (m *Machine) validateExternalRegionID(formats strfmt.Registry) error {

	if err := validate.Required("externalRegionId", "body", m.ExternalRegionID); err != nil {
		return err
	}

	return nil
}

func (m *Machine) validateExternalZoneID(formats strfmt.Registry) error {

	if err := validate.Required("externalZoneId", "body", m.ExternalZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Machine) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var machineTypePowerStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ON","OFF","GUEST_OFF","UNKNOWN","SUSPEND"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		machineTypePowerStatePropEnum = append(machineTypePowerStatePropEnum, v)
	}
}

const (

	// MachinePowerStateON captures enum value "ON"
	MachinePowerStateON string = "ON"

	// MachinePowerStateOFF captures enum value "OFF"
	MachinePowerStateOFF string = "OFF"

	// MachinePowerStateGUESTOFF captures enum value "GUEST_OFF"
	MachinePowerStateGUESTOFF string = "GUEST_OFF"

	// MachinePowerStateUNKNOWN captures enum value "UNKNOWN"
	MachinePowerStateUNKNOWN string = "UNKNOWN"

	// MachinePowerStateSUSPEND captures enum value "SUSPEND"
	MachinePowerStateSUSPEND string = "SUSPEND"
)

// prop value enum
func (m *Machine) validatePowerStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, machineTypePowerStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Machine) validatePowerState(formats strfmt.Registry) error {

	if err := validate.Required("powerState", "body", m.PowerState); err != nil {
		return err
	}

	// value enum
	if err := m.validatePowerStateEnum("powerState", "body", *m.PowerState); err != nil {
		return err
	}

	return nil
}

func (m *Machine) validateTags(formats strfmt.Registry) error {

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
func (m *Machine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Machine) UnmarshalBinary(b []byte) error {
	var res Machine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
