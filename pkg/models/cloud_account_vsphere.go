// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudAccountVsphere State object representing a vSphere cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.<br>**HATEOAS** links:<br>**regions** - array[Region] - List of regions that are enabled for this cloud account.<br>**self** - CloudAccountVsphere - Self link to this cloud account.
//
// swagger:model CloudAccountVsphere
type CloudAccountVsphere struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Workload mobility associations to other vSphere cloud accounts. ID refers to an associated cloud account, and directionality can be unidirectional or bidirectional.
	// Example: { \"42f3e0d199d134755684cd935435a\": \"BIDIRECTIONAL\" }
	AssociatedMobilityCloudAccountIds map[string]string `json:"associatedMobilityCloudAccountIds,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"isExternal\" : \"false\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	Dcid string `json:"dcid,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// A list of regions that are enabled for provisioning on this cloud account
	EnabledRegions []*Region `json:"enabledRegions"`

	// Host name for the vSphere cloud account
	// Example: vc1.vmware.com
	// Required: true
	HostName *string `json:"hostName"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// A set of tag keys and optional values that were set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`

	// Username to authenticate with the cloud account
	// Example: administrator@mycompany.com
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this cloud account vsphere
func (m *CloudAccountVsphere) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedMobilityCloudAccountIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabledRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVsphere) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_links" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_links" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// additional properties value enum
var cloudAccountVsphereAssociatedMobilityCloudAccountIdsValueEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNIDIRECTIONAL","BIDIRECTIONAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudAccountVsphereAssociatedMobilityCloudAccountIdsValueEnum = append(cloudAccountVsphereAssociatedMobilityCloudAccountIdsValueEnum, v)
	}
}

func (m *CloudAccountVsphere) validateAssociatedMobilityCloudAccountIdsValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloudAccountVsphereAssociatedMobilityCloudAccountIdsValueEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloudAccountVsphere) validateAssociatedMobilityCloudAccountIds(formats strfmt.Registry) error {
	if swag.IsZero(m.AssociatedMobilityCloudAccountIds) { // not required
		return nil
	}

	for k := range m.AssociatedMobilityCloudAccountIds {

		// value enum
		if err := m.validateAssociatedMobilityCloudAccountIdsValueEnum("associatedMobilityCloudAccountIds"+"."+k, "body", m.AssociatedMobilityCloudAccountIds[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *CloudAccountVsphere) validateEnabledRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.EnabledRegions) { // not required
		return nil
	}

	for i := 0; i < len(m.EnabledRegions); i++ {
		if swag.IsZero(m.EnabledRegions[i]) { // not required
			continue
		}

		if m.EnabledRegions[i] != nil {
			if err := m.EnabledRegions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enabledRegions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("enabledRegions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountVsphere) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVsphere) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVsphere) validateTags(formats strfmt.Registry) error {
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

func (m *CloudAccountVsphere) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cloud account vsphere based on the context it is used
func (m *CloudAccountVsphere) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnabledRegions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVsphere) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if val, ok := m.Links[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountVsphere) contextValidateEnabledRegions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EnabledRegions); i++ {

		if m.EnabledRegions[i] != nil {
			if err := m.EnabledRegions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("enabledRegions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("enabledRegions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountVsphere) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudAccountVsphere) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountVsphere) UnmarshalBinary(b []byte) error {
	var res CloudAccountVsphere
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
