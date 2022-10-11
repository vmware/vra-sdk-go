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

// K8SResource K8SResource
//
// swagger:model K8SResource
type K8SResource struct {

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// index
	Index int32 `json:"index,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// resource
	Resource *HasMetadata `json:"resource,omitempty"`

	// state
	// Enum: [INITIALIZED NOT_INSTALLED INSTALLING INSTALLED READY UNINSTALLING UNREACHABLE FAILED]
	State string `json:"state,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}

// Validate validates this k8 s resource
func (m *K8SResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8SResource) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *K8SResource) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

var k8SResourceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INITIALIZED","NOT_INSTALLED","INSTALLING","INSTALLED","READY","UNINSTALLING","UNREACHABLE","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		k8SResourceTypeStatePropEnum = append(k8SResourceTypeStatePropEnum, v)
	}
}

const (

	// K8SResourceStateINITIALIZED captures enum value "INITIALIZED"
	K8SResourceStateINITIALIZED string = "INITIALIZED"

	// K8SResourceStateNOTINSTALLED captures enum value "NOT_INSTALLED"
	K8SResourceStateNOTINSTALLED string = "NOT_INSTALLED"

	// K8SResourceStateINSTALLING captures enum value "INSTALLING"
	K8SResourceStateINSTALLING string = "INSTALLING"

	// K8SResourceStateINSTALLED captures enum value "INSTALLED"
	K8SResourceStateINSTALLED string = "INSTALLED"

	// K8SResourceStateREADY captures enum value "READY"
	K8SResourceStateREADY string = "READY"

	// K8SResourceStateUNINSTALLING captures enum value "UNINSTALLING"
	K8SResourceStateUNINSTALLING string = "UNINSTALLING"

	// K8SResourceStateUNREACHABLE captures enum value "UNREACHABLE"
	K8SResourceStateUNREACHABLE string = "UNREACHABLE"

	// K8SResourceStateFAILED captures enum value "FAILED"
	K8SResourceStateFAILED string = "FAILED"
)

// prop value enum
func (m *K8SResource) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, k8SResourceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *K8SResource) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this k8 s resource based on the context it is used
func (m *K8SResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8SResource) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {
		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *K8SResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8SResource) UnmarshalBinary(b []byte) error {
	var res K8SResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}