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

// K8SNode K8SNode
//
// swagger:model K8SNode
type K8SNode struct {

	// allocatable pods
	AllocatablePods int32 `json:"allocatablePods,omitempty"`

	// cpu cores
	CPUCores float64 `json:"cpuCores,omitempty"`

	// cpu usage
	CPUUsage float64 `json:"cpuUsage,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// memory usage
	MemoryUsage float64 `json:"memoryUsage,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// total memory
	TotalMemory float64 `json:"totalMemory,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}

// Validate validates this k8 s node
func (m *K8SNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8SNode) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this k8 s node based on context it is used
func (m *K8SNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *K8SNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8SNode) UnmarshalBinary(b []byte) error {
	var res K8SNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
