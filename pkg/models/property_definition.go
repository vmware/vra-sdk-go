// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PropertyDefinition PropertyDefinition
//
// swagger:model PropertyDefinition
type PropertyDefinition struct {

	// dollar data
	DollarData string `json:"$data,omitempty"`

	// dollar dynamic default
	DollarDynamicDefault string `json:"$dynamicDefault,omitempty"`

	// dollar ref
	DollarRef string `json:"$ref,omitempty"`

	// additional properties
	AdditionalProperties bool `json:"additionalProperties,omitempty"`

	// all of
	AllOf []*PropertyDefinition `json:"allOf"`

	// any of
	AnyOf []*PropertyDefinition `json:"anyOf"`

	// computed
	Computed bool `json:"computed,omitempty"`

	// const
	Const interface{} `json:"const,omitempty"`

	// default
	Default interface{} `json:"default,omitempty"`

	// dependencies
	Dependencies map[string][]string `json:"dependencies,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// encrypted
	Encrypted bool `json:"encrypted,omitempty"`

	// enum
	Enum []interface{} `json:"enum"`

	// format
	Format string `json:"format,omitempty"`

	// ignore case on diff
	IgnoreCaseOnDiff bool `json:"ignoreCaseOnDiff,omitempty"`

	// ignore on update
	IgnoreOnUpdate bool `json:"ignoreOnUpdate,omitempty"`

	// items
	Items *PropertyDefinition `json:"items,omitempty"`

	// max items
	MaxItems int64 `json:"maxItems,omitempty"`

	// max length
	MaxLength int32 `json:"maxLength,omitempty"`

	// max properties
	MaxProperties int32 `json:"maxProperties,omitempty"`

	// maximum
	Maximum int64 `json:"maximum,omitempty"`

	// min items
	MinItems int64 `json:"minItems,omitempty"`

	// min length
	MinLength int32 `json:"minLength,omitempty"`

	// min properties
	MinProperties int32 `json:"minProperties,omitempty"`

	// minimum
	Minimum int64 `json:"minimum,omitempty"`

	// not
	Not *PropertyDefinition `json:"not,omitempty"`

	// one of
	OneOf []*PropertyDefinition `json:"oneOf"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// properties
	Properties map[string]PropertyDefinition `json:"properties,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// recreate on update
	RecreateOnUpdate bool `json:"recreateOnUpdate,omitempty"`

	// required
	Required []string `json:"required"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unique items
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// write only
	WriteOnly bool `json:"writeOnly,omitempty"`
}

// Validate validates this property definition
func (m *PropertyDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnyOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneOf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyDefinition) validateAllOf(formats strfmt.Registry) error {
	if swag.IsZero(m.AllOf) { // not required
		return nil
	}

	for i := 0; i < len(m.AllOf); i++ {
		if swag.IsZero(m.AllOf[i]) { // not required
			continue
		}

		if m.AllOf[i] != nil {
			if err := m.AllOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDefinition) validateAnyOf(formats strfmt.Registry) error {
	if swag.IsZero(m.AnyOf) { // not required
		return nil
	}

	for i := 0; i < len(m.AnyOf); i++ {
		if swag.IsZero(m.AnyOf[i]) { // not required
			continue
		}

		if m.AnyOf[i] != nil {
			if err := m.AnyOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("anyOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("anyOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDefinition) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	if m.Items != nil {
		if err := m.Items.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("items")
			}
			return err
		}
	}

	return nil
}

func (m *PropertyDefinition) validateNot(formats strfmt.Registry) error {
	if swag.IsZero(m.Not) { // not required
		return nil
	}

	if m.Not != nil {
		if err := m.Not.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("not")
			}
			return err
		}
	}

	return nil
}

func (m *PropertyDefinition) validateOneOf(formats strfmt.Registry) error {
	if swag.IsZero(m.OneOf) { // not required
		return nil
	}

	for i := 0; i < len(m.OneOf); i++ {
		if swag.IsZero(m.OneOf[i]) { // not required
			continue
		}

		if m.OneOf[i] != nil {
			if err := m.OneOf[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oneOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("oneOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDefinition) validateProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	for k := range m.Properties {

		if err := validate.Required("properties"+"."+k, "body", m.Properties[k]); err != nil {
			return err
		}
		if val, ok := m.Properties[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("properties" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this property definition based on the context it is used
func (m *PropertyDefinition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAnyOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOneOf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PropertyDefinition) contextValidateAllOf(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllOf); i++ {

		if m.AllOf[i] != nil {
			if err := m.AllOf[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDefinition) contextValidateAnyOf(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AnyOf); i++ {

		if m.AnyOf[i] != nil {
			if err := m.AnyOf[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("anyOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("anyOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDefinition) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	if m.Items != nil {
		if err := m.Items.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("items")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("items")
			}
			return err
		}
	}

	return nil
}

func (m *PropertyDefinition) contextValidateNot(ctx context.Context, formats strfmt.Registry) error {

	if m.Not != nil {
		if err := m.Not.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("not")
			}
			return err
		}
	}

	return nil
}

func (m *PropertyDefinition) contextValidateOneOf(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OneOf); i++ {

		if m.OneOf[i] != nil {
			if err := m.OneOf[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("oneOf" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("oneOf" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PropertyDefinition) contextValidateProperties(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Properties {

		if val, ok := m.Properties[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PropertyDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PropertyDefinition) UnmarshalBinary(b []byte) error {
	var res PropertyDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
