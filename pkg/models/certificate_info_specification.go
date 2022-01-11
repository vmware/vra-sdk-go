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

// CertificateInfoSpecification Specification for certificate for a cloud account.
//
// swagger:model CertificateInfoSpecification
type CertificateInfoSpecification struct {

	// The certificate in string format.
	// Example: PEM for X509Certificate
	// Required: true
	Certificate *string `json:"certificate"`
}

// Validate validates this certificate info specification
func (m *CertificateInfoSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateInfoSpecification) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate info specification based on context it is used
func (m *CertificateInfoSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateInfoSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateInfoSpecification) UnmarshalBinary(b []byte) error {
	var res CertificateInfoSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
