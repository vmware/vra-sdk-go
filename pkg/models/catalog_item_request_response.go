// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogItemRequestResponse CatalogItemRequestResponse
//
// # The blueprint service's response to a deployment request
//
// swagger:model CatalogItemRequestResponse
type CatalogItemRequestResponse struct {

	// The created deployment's ID
	DeploymentID string `json:"deploymentId,omitempty"`

	// The created deployment's name
	DeploymentName string `json:"deploymentName,omitempty"`
}

// Validate validates this catalog item request response
func (m *CatalogItemRequestResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this catalog item request response based on context it is used
func (m *CatalogItemRequestResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogItemRequestResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemRequestResponse) UnmarshalBinary(b []byte) error {
	var res CatalogItemRequestResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
