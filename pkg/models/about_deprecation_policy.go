// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AboutDeprecationPolicy About$DeprecationPolicy
// swagger:model AboutDeprecationPolicy
type AboutDeprecationPolicy struct {

	// The date the api was deprecated in yyyy-MM-dd format (UTC). Could be empty if the api is not deprecated.
	DeprecatedAt string `json:"deprecatedAt,omitempty"`

	// A free text description that contains information about why this api is deprecated and how to migrate to a newer version.
	Description string `json:"description,omitempty"`

	// The date the api support will be dropped in yyyy-MM-dd format (UTC). The api may still be available for use after that date but this is not guaranteed.
	ExpiresAt string `json:"expiresAt,omitempty"`
}

// Validate validates this about deprecation policy
func (m *AboutDeprecationPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AboutDeprecationPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AboutDeprecationPolicy) UnmarshalBinary(b []byte) error {
	var res AboutDeprecationPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
