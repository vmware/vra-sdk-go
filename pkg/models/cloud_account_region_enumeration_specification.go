// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudAccountRegionEnumerationSpecification Specification for a region enumeration of generic cloud account.
//
// swagger:model CloudAccountRegionEnumerationSpecification
type CloudAccountRegionEnumerationSpecification struct {

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Existing cloud account id. Either provide existing cloud account Id, or privateKeyId/privateKey credentials pair.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Cloud Account specific properties supplied in as name value pairs. In case of AAP, provide environment property here. Example: "cloudAccountProperties": {
	//         "environment": "aap"
	//     }
	// Example: {\"supportPublicImages\": \"true\", \"acceptSelfSignedCertificate\": \"true\" }
	// Required: true
	CloudAccountProperties map[string]string `json:"cloudAccountProperties"`

	// Cloud account type
	// Example: vsphere, aws, azure, nsxv, nsxt
	CloudAccountType string `json:"cloudAccountType,omitempty"`

	// Secret access key or password to be used to authenticate with the cloud account. Either provide privateKey or provide a cloudAccountId of an existing account.
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	PrivateKey string `json:"privateKey,omitempty"`

	// Access key id or username to be used to authenticate with the cloud account. Either provide privateKeyId or provide a cloudAccountId of an existing account.
	// Example: ACDC55DB4MFH6ADG75KK
	PrivateKeyID string `json:"privateKeyId,omitempty"`
}

// Validate validates this cloud account region enumeration specification
func (m *CloudAccountRegionEnumerationSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountRegionEnumerationSpecification) validateCertificateInfo(formats strfmt.Registry) error {
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

func (m *CloudAccountRegionEnumerationSpecification) validateCloudAccountProperties(formats strfmt.Registry) error {

	if err := validate.Required("cloudAccountProperties", "body", m.CloudAccountProperties); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cloud account region enumeration specification based on the context it is used
func (m *CloudAccountRegionEnumerationSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountRegionEnumerationSpecification) contextValidateCertificateInfo(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudAccountRegionEnumerationSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountRegionEnumerationSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountRegionEnumerationSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}