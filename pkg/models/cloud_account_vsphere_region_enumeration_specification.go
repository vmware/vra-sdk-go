// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudAccountVsphereRegionEnumerationSpecification Specification for a region enumeration of vshpere cloud account.
//
// swagger:model CloudAccountVsphereRegionEnumerationSpecification
type CloudAccountVsphereRegionEnumerationSpecification struct {

	// Accept self signed certificate when connecting to vSphere
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Existing cloud account id. Either provide existing cloud account Id, or hostName, username, password.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	Dcid string `json:"dcid,omitempty"`

	// The environment where data collectors are deployed. When the data collectors are deployed on a cloud gateway appliance, use "aap".
	// Example: aap
	Environment string `json:"environment,omitempty"`

	// Host name for the vSphere endpoint. Either provide hostName or provide a cloudAccountId of an existing account.
	// Example: vc.mycompany.com
	HostName string `json:"hostName,omitempty"`

	// Password for the user used to authenticate with the cloud Account. Either provide password or provide a cloudAccountId of an existing account.
	// Example: cndhjslacd90ascdbasyoucbdh
	Password string `json:"password,omitempty"`

	// Username to authenticate with the cloud account. Either provide username or provide a cloudAccountId of an existing account.
	// Example: administrator@mycompany.com
	Username string `json:"username,omitempty"`
}

// Validate validates this cloud account vsphere region enumeration specification
func (m *CloudAccountVsphereRegionEnumerationSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVsphereRegionEnumerationSpecification) validateCertificateInfo(formats strfmt.Registry) error {
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

// ContextValidate validate this cloud account vsphere region enumeration specification based on the context it is used
func (m *CloudAccountVsphereRegionEnumerationSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVsphereRegionEnumerationSpecification) contextValidateCertificateInfo(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CloudAccountVsphereRegionEnumerationSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountVsphereRegionEnumerationSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountVsphereRegionEnumerationSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
