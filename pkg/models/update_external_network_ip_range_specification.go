// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateExternalNetworkIPRangeSpecification Specification for updating an ExternalNetworkIPRange
//
// swagger:model UpdateExternalNetworkIPRangeSpecification
type UpdateExternalNetworkIPRangeSpecification struct {

	// Deprecated. Use 'fabricNetworkIds'.
	FabricNetworkID string `json:"fabricNetworkId,omitempty"`

	// A list of fabric network Ids that this IP range should be associated with.
	FabricNetworkIds []string `json:"fabricNetworkIds"`
}

// Validate validates this update external network IP range specification
func (m *UpdateExternalNetworkIPRangeSpecification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update external network IP range specification based on context it is used
func (m *UpdateExternalNetworkIPRangeSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateExternalNetworkIPRangeSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateExternalNetworkIPRangeSpecification) UnmarshalBinary(b []byte) error {
	var res UpdateExternalNetworkIPRangeSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
