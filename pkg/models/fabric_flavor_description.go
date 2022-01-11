// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FabricFlavorDescription Represents fabric flavor instance type description. Used when creating flavor profiles.
//
// swagger:model FabricFlavorDescription
type FabricFlavorDescription struct {

	// Number of CPU cores. Mandatory for private clouds such as vSphere. Not populated when inapplicable.
	// Example: 4
	CPUCount int32 `json:"cpuCount,omitempty"`

	// The id of the instance type in the corresponding cloud.
	// Example: 901004, 901008
	ID string `json:"id,omitempty"`

	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Not populated when inapplicable.
	// Example: 4096
	MemoryInMB int64 `json:"memoryInMB,omitempty"`

	// The value of the instance type in the corresponding cloud. Valid and mandatory for public clouds
	// Example: t2.small, t2.medium
	Name string `json:"name,omitempty"`
}

// Validate validates this fabric flavor description
func (m *FabricFlavorDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this fabric flavor description based on context it is used
func (m *FabricFlavorDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FabricFlavorDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricFlavorDescription) UnmarshalBinary(b []byte) error {
	var res FabricFlavorDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
