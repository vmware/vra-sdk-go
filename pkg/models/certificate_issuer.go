// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CertificateIssuer CertificateIssuer
// swagger:model CertificateIssuer
type CertificateIssuer struct {

	// A human-friendly name used as an identifier for the holding body.
	CommonName string `json:"commonName,omitempty"`

	// Name of the organisation.
	Organization string `json:"organization,omitempty"`
}

// Validate validates this certificate issuer
func (m *CertificateIssuer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateIssuer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateIssuer) UnmarshalBinary(b []byte) error {
	var res CertificateIssuer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
