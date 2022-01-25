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

// CloudAccountVsphereSpecification Specification for a vSphere cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.
//
// swagger:model CloudAccountVsphereSpecification
type CloudAccountVsphereSpecification struct {

	// Accept self signed certificate when connecting to vSphere
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// NSX-V or NSX-T account to associate with this vSphere cloud account. vSphere cloud account can be a single NSX-V cloud account or a single NSX-T cloud account.
	// Example: [ \"42f3e0d199d134755684cd935435a\" ]
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	Dcid string `json:"dcid,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Host name for the vSphere endpoint
	// Example: vc.mycompany.com
	// Required: true
	HostName *string `json:"hostName"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Password for the user used to authenticate with the cloud Account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	Password *string `json:"password"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"Datacenter:datacenter-3\",\"externalRegionId\": \"Datacenter:datacenter-3\"}]
	// Required: true
	Regions []*RegionSpecification `json:"regions"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Username to authenticate with the cloud account
	// Example: administrator@mycompany.com
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this cloud account vsphere specification
func (m *CloudAccountVsphereSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
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

func (m *CloudAccountVsphereSpecification) validateCertificateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CertificateInfo) { // not required
		return nil
	}

	if m.CertificateInfo != nil {
		if err := m.CertificateInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountVsphereSpecification) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVsphereSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVsphereSpecification) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountVsphereSpecification) validateRegions(formats strfmt.Registry) error {

	if err := validate.Required("regions", "body", m.Regions); err != nil {
		return err
	}

	for i := 0; i < len(m.Regions); i++ {
		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {
			if err := m.Regions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountVsphereSpecification) validateTags(formats strfmt.Registry) error {
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

func (m *CloudAccountVsphereSpecification) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cloud account vsphere specification based on the context it is used
func (m *CloudAccountVsphereSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegions(ctx, formats); err != nil {
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

func (m *CloudAccountVsphereSpecification) contextValidateCertificateInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CertificateInfo != nil {
		if err := m.CertificateInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountVsphereSpecification) contextValidateRegions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Regions); i++ {

		if m.Regions[i] != nil {
			if err := m.Regions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudAccountVsphereSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudAccountVsphereSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountVsphereSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountVsphereSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
