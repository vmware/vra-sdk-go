// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CertificateInfo Certificate for a cloud account.
//
// swagger:model CertificateInfo
type CertificateInfo struct {

	// The certificate in string format.
	// Example: PEM for X509Certificate
	// Required: true
	Certificate *string `json:"certificate"`

	// Details about the certificate.
	// Example: UNTRUSTED_CERTIFICATE
	// Enum: [UNTRUSTED_CERTIFICATE EXPIRED_CERTIFICATE NOT_YET_VALID_CERTIFICATE KEYSTORE_TAMPERED_OR_PASSWORD_INCORRECT]
	CertificateErrorDetail string `json:"certificateErrorDetail,omitempty"`

	// Certificate related properties which may provide additional information about the given certificate.
	// Required: true
	Properties map[string]string `json:"properties"`
}

// Validate validates this certificate info
func (m *CertificateInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateErrorDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateInfo) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

var certificateInfoTypeCertificateErrorDetailPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNTRUSTED_CERTIFICATE","EXPIRED_CERTIFICATE","NOT_YET_VALID_CERTIFICATE","KEYSTORE_TAMPERED_OR_PASSWORD_INCORRECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateInfoTypeCertificateErrorDetailPropEnum = append(certificateInfoTypeCertificateErrorDetailPropEnum, v)
	}
}

const (

	// CertificateInfoCertificateErrorDetailUNTRUSTEDCERTIFICATE captures enum value "UNTRUSTED_CERTIFICATE"
	CertificateInfoCertificateErrorDetailUNTRUSTEDCERTIFICATE string = "UNTRUSTED_CERTIFICATE"

	// CertificateInfoCertificateErrorDetailEXPIREDCERTIFICATE captures enum value "EXPIRED_CERTIFICATE"
	CertificateInfoCertificateErrorDetailEXPIREDCERTIFICATE string = "EXPIRED_CERTIFICATE"

	// CertificateInfoCertificateErrorDetailNOTYETVALIDCERTIFICATE captures enum value "NOT_YET_VALID_CERTIFICATE"
	CertificateInfoCertificateErrorDetailNOTYETVALIDCERTIFICATE string = "NOT_YET_VALID_CERTIFICATE"

	// CertificateInfoCertificateErrorDetailKEYSTORETAMPEREDORPASSWORDINCORRECT captures enum value "KEYSTORE_TAMPERED_OR_PASSWORD_INCORRECT"
	CertificateInfoCertificateErrorDetailKEYSTORETAMPEREDORPASSWORDINCORRECT string = "KEYSTORE_TAMPERED_OR_PASSWORD_INCORRECT"
)

// prop value enum
func (m *CertificateInfo) validateCertificateErrorDetailEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateInfoTypeCertificateErrorDetailPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateInfo) validateCertificateErrorDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.CertificateErrorDetail) { // not required
		return nil
	}

	// value enum
	if err := m.validateCertificateErrorDetailEnum("certificateErrorDetail", "body", m.CertificateErrorDetail); err != nil {
		return err
	}

	return nil
}

func (m *CertificateInfo) validateProperties(formats strfmt.Registry) error {

	if err := validate.Required("properties", "body", m.Properties); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate info based on context it is used
func (m *CertificateInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateInfo) UnmarshalBinary(b []byte) error {
	var res CertificateInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
