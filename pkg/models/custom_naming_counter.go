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

// CustomNamingCounter A representation of a Counter.
//
// swagger:model CustomNamingCounter
type CustomNamingCounter struct {

	// active
	Active bool `json:"active,omitempty"`

	// The resource type of custom name
	// Required: true
	// Enum: [COMPUTE NETWORK COMPUTE_STORAGE LOAD_BALANCER RESOURCE_GROUP GATEWAY NAT SECURITY_GROUP]
	CnResourceType *string `json:"cnResourceType"`

	// The current counter of custom name
	// Required: true
	CurrentCounter *int64 `json:"currentCounter"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The project id to which the counter is mapped
	// Required: true
	ProjectID *string `json:"projectId"`
}

// Validate validates this custom naming counter
func (m *CustomNamingCounter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCnResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentCounter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customNamingCounterTypeCnResourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COMPUTE","NETWORK","COMPUTE_STORAGE","LOAD_BALANCER","RESOURCE_GROUP","GATEWAY","NAT","SECURITY_GROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customNamingCounterTypeCnResourceTypePropEnum = append(customNamingCounterTypeCnResourceTypePropEnum, v)
	}
}

const (

	// CustomNamingCounterCnResourceTypeCOMPUTE captures enum value "COMPUTE"
	CustomNamingCounterCnResourceTypeCOMPUTE string = "COMPUTE"

	// CustomNamingCounterCnResourceTypeNETWORK captures enum value "NETWORK"
	CustomNamingCounterCnResourceTypeNETWORK string = "NETWORK"

	// CustomNamingCounterCnResourceTypeCOMPUTESTORAGE captures enum value "COMPUTE_STORAGE"
	CustomNamingCounterCnResourceTypeCOMPUTESTORAGE string = "COMPUTE_STORAGE"

	// CustomNamingCounterCnResourceTypeLOADBALANCER captures enum value "LOAD_BALANCER"
	CustomNamingCounterCnResourceTypeLOADBALANCER string = "LOAD_BALANCER"

	// CustomNamingCounterCnResourceTypeRESOURCEGROUP captures enum value "RESOURCE_GROUP"
	CustomNamingCounterCnResourceTypeRESOURCEGROUP string = "RESOURCE_GROUP"

	// CustomNamingCounterCnResourceTypeGATEWAY captures enum value "GATEWAY"
	CustomNamingCounterCnResourceTypeGATEWAY string = "GATEWAY"

	// CustomNamingCounterCnResourceTypeNAT captures enum value "NAT"
	CustomNamingCounterCnResourceTypeNAT string = "NAT"

	// CustomNamingCounterCnResourceTypeSECURITYGROUP captures enum value "SECURITY_GROUP"
	CustomNamingCounterCnResourceTypeSECURITYGROUP string = "SECURITY_GROUP"
)

// prop value enum
func (m *CustomNamingCounter) validateCnResourceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, customNamingCounterTypeCnResourceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CustomNamingCounter) validateCnResourceType(formats strfmt.Registry) error {

	if err := validate.Required("cnResourceType", "body", m.CnResourceType); err != nil {
		return err
	}

	// value enum
	if err := m.validateCnResourceTypeEnum("cnResourceType", "body", *m.CnResourceType); err != nil {
		return err
	}

	return nil
}

func (m *CustomNamingCounter) validateCurrentCounter(formats strfmt.Registry) error {

	if err := validate.Required("currentCounter", "body", m.CurrentCounter); err != nil {
		return err
	}

	return nil
}

func (m *CustomNamingCounter) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomNamingCounter) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom naming counter based on context it is used
func (m *CustomNamingCounter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomNamingCounter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomNamingCounter) UnmarshalBinary(b []byte) error {
	var res CustomNamingCounter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
