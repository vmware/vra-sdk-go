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

// ResourceAction ResourceAction
//
// swagger:model ResourceAction
type ResourceAction struct {

	// Resource action type
	// Enum: [RESOURCE_ACTION RESOURCE_EXTENSION]
	ActionType string `json:"actionType,omitempty"`

	// Dependent resources
	Dependents []string `json:"dependents"`

	// Resource action description
	Description string `json:"description,omitempty"`

	// Resource action display name
	DisplayName string `json:"displayName,omitempty"`

	// Resource action custom UI definition. Optional
	FormDefinition *FormDefinition `json:"formDefinition,omitempty"`

	// Resource action id
	ID string `json:"id,omitempty"`

	// Resource action name
	Name string `json:"name,omitempty"`

	// Resource action org ID
	OrgID string `json:"orgId,omitempty"`

	// Resource action project ID
	ProjectID string `json:"projectId,omitempty"`

	// Resource action input schema
	Schema interface{} `json:"schema,omitempty"`

	// Resource action is valid for current state
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this resource action
func (m *ResourceAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormDefinition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resourceActionTypeActionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RESOURCE_ACTION","RESOURCE_EXTENSION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceActionTypeActionTypePropEnum = append(resourceActionTypeActionTypePropEnum, v)
	}
}

const (

	// ResourceActionActionTypeRESOURCEACTION captures enum value "RESOURCE_ACTION"
	ResourceActionActionTypeRESOURCEACTION string = "RESOURCE_ACTION"

	// ResourceActionActionTypeRESOURCEEXTENSION captures enum value "RESOURCE_EXTENSION"
	ResourceActionActionTypeRESOURCEEXTENSION string = "RESOURCE_EXTENSION"
)

// prop value enum
func (m *ResourceAction) validateActionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resourceActionTypeActionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResourceAction) validateActionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionTypeEnum("actionType", "body", m.ActionType); err != nil {
		return err
	}

	return nil
}

func (m *ResourceAction) validateFormDefinition(formats strfmt.Registry) error {
	if swag.IsZero(m.FormDefinition) { // not required
		return nil
	}

	if m.FormDefinition != nil {
		if err := m.FormDefinition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formDefinition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("formDefinition")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this resource action based on the context it is used
func (m *ResourceAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFormDefinition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceAction) contextValidateFormDefinition(ctx context.Context, formats strfmt.Registry) error {

	if m.FormDefinition != nil {
		if err := m.FormDefinition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formDefinition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("formDefinition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceAction) UnmarshalBinary(b []byte) error {
	var res ResourceAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
