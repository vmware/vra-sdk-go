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

// SupervisorCluster SupervisorCluster
//
// swagger:model SupervisorCluster
type SupervisorCluster struct {

	// address
	Address string `json:"address,omitempty"`

	// cpu capacity
	CPUCapacity int64 `json:"cpuCapacity,omitempty"`

	// cpu used
	CPUUsed int64 `json:"cpuUsed,omitempty"`

	// document self link
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// endpoint link
	EndpointLink string `json:"endpointLink,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// memory capacity
	MemoryCapacity int64 `json:"memoryCapacity,omitempty"`

	// memory used
	MemoryUsed int64 `json:"memoryUsed,omitempty"`

	// moref
	Moref string `json:"moref,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// registered
	Registered bool `json:"registered,omitempty"`

	// software version
	SoftwareVersion string `json:"softwareVersion,omitempty"`

	// status
	// Enum: [ON CONFIGURING REMOVING ERROR]
	Status string `json:"status,omitempty"`

	// status message
	StatusMessage string `json:"statusMessage,omitempty"`

	// storage capacity
	StorageCapacity int64 `json:"storageCapacity,omitempty"`

	// storage used
	StorageUsed int64 `json:"storageUsed,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}

// Validate validates this supervisor cluster
func (m *SupervisorCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var supervisorClusterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ON","CONFIGURING","REMOVING","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supervisorClusterTypeStatusPropEnum = append(supervisorClusterTypeStatusPropEnum, v)
	}
}

const (

	// SupervisorClusterStatusON captures enum value "ON"
	SupervisorClusterStatusON string = "ON"

	// SupervisorClusterStatusCONFIGURING captures enum value "CONFIGURING"
	SupervisorClusterStatusCONFIGURING string = "CONFIGURING"

	// SupervisorClusterStatusREMOVING captures enum value "REMOVING"
	SupervisorClusterStatusREMOVING string = "REMOVING"

	// SupervisorClusterStatusERROR captures enum value "ERROR"
	SupervisorClusterStatusERROR string = "ERROR"
)

// prop value enum
func (m *SupervisorCluster) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, supervisorClusterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SupervisorCluster) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this supervisor cluster based on context it is used
func (m *SupervisorCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupervisorCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupervisorCluster) UnmarshalBinary(b []byte) error {
	var res SupervisorCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
