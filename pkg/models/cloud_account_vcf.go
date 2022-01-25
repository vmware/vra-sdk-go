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

// CloudAccountVcf State object representing a VCF cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.<br>**HATEOAS** links:<br>**regions** - array[Region] - List of regions that are enabled for this cloud account.<br>**self** - CloudAccountVcf - Self link to this cloud account
//
// swagger:model CloudAccountVcf
type CloudAccountVcf struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"isExternal\" : \"false\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// A list of regions that are enabled for provisioning on this cloud account
	EnabledRegions []*Region `json:"enabledRegions"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// NSX Hostname in a workload domain
	// Example: nsx.mycompany.com
	// Required: true
	NsxHostName *string `json:"nsxHostName"`

	// Username to authenticate to NSX manager in workload domain
	// Example: administrator@mycompany.com
	// Required: true
	NsxUsername *string `json:"nsxUsername"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// SDDC Manager ID
	// Example: 2e5bb71d-0c14-4066-a999-2cb6c693654a
	// Required: true
	SddcManagerID *string `json:"sddcManagerId"`

	// A set of tag keys and optional values to set on the Cloud Account.Cloud account capability tags may enable different features.
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`

	// vCenter Hostname in a workload domain
	// Example: vcenter.mycompany.com
	// Required: true
	VcenterHostName *string `json:"vcenterHostName"`

	// Username to authenticate to vCenter Server in workload domain
	// Example: administrator@mycompany.com
	// Required: true
	VcenterUsername *string `json:"vcenterUsername"`

	// Id of the VCF worload domain.
	// Example: 587db412-6037-43e4-8e1e-49ebbaf6cd35
	// Required: true
	VcfDomainID *string `json:"vcfDomainId"`

	// Name of the VCF worload domain.
	// Example: Workload Domain - 1
	// Required: true
	VcfDomainName *string `json:"vcfDomainName"`
}

// Validate validates this cloud account vcf
func (m *CloudAccountVcf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabledRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNsxUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSddcManagerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcenterHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcenterUsername(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcfDomainID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcfDomainName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVcf) validateLinks(formats strfmt.Registry) error {

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

func (m *CloudAccountVcf) validateEnabledRegions(formats strfmt.Registry) error {
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

func (m *CloudAccountVcf) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcf) validateNsxHostName(formats strfmt.Registry) error {

	if err := validate.Required("nsxHostName", "body", m.NsxHostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcf) validateNsxUsername(formats strfmt.Registry) error {

	if err := validate.Required("nsxUsername", "body", m.NsxUsername); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcf) validateSddcManagerID(formats strfmt.Registry) error {

	if err := validate.Required("sddcManagerId", "body", m.SddcManagerID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcf) validateTags(formats strfmt.Registry) error {
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

func (m *CloudAccountVcf) validateVcenterHostName(formats strfmt.Registry) error {

	if err := validate.Required("vcenterHostName", "body", m.VcenterHostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcf) validateVcenterUsername(formats strfmt.Registry) error {

	if err := validate.Required("vcenterUsername", "body", m.VcenterUsername); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcf) validateVcfDomainID(formats strfmt.Registry) error {

	if err := validate.Required("vcfDomainId", "body", m.VcfDomainID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVcf) validateVcfDomainName(formats strfmt.Registry) error {

	if err := validate.Required("vcfDomainName", "body", m.VcfDomainName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cloud account vcf based on the context it is used
func (m *CloudAccountVcf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *CloudAccountVcf) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CloudAccountVcf) contextValidateEnabledRegions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CloudAccountVcf) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudAccountVcf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountVcf) UnmarshalBinary(b []byte) error {
	var res CloudAccountVcf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
