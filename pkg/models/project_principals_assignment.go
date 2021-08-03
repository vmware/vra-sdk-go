// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectPrincipalsAssignment ProjectPrincipalsAssignment
//
// swagger:model ProjectPrincipalsAssignment
type ProjectPrincipalsAssignment struct {
	modifyField []PrincipalRole

	removeField []PrincipalRole
}

// Modify gets the modify of this base type
func (m *ProjectPrincipalsAssignment) Modify() []PrincipalRole {
	return m.modifyField
}

// SetModify sets the modify of this base type
func (m *ProjectPrincipalsAssignment) SetModify(val []PrincipalRole) {
	m.modifyField = val
}

// Remove gets the remove of this base type
func (m *ProjectPrincipalsAssignment) Remove() []PrincipalRole {
	return m.removeField
}

// SetRemove sets the remove of this base type
func (m *ProjectPrincipalsAssignment) SetRemove(val []PrincipalRole) {
	m.removeField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ProjectPrincipalsAssignment) UnmarshalJSON(raw []byte) error {
	var data struct {
		Modify json.RawMessage `json:"modify"`

		Remove json.RawMessage `json:"remove"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propModify []PrincipalRole
	if string(data.Modify) != "null" {
		modify, err := UnmarshalPrincipalRoleSlice(bytes.NewBuffer(data.Modify), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propModify = modify
	}
	var propRemove []PrincipalRole
	if string(data.Remove) != "null" {
		remove, err := UnmarshalPrincipalRoleSlice(bytes.NewBuffer(data.Remove), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propRemove = remove
	}

	var result ProjectPrincipalsAssignment

	// modify
	result.modifyField = propModify

	// remove
	result.removeField = propRemove

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ProjectPrincipalsAssignment) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Modify []PrincipalRole `json:"modify"`

		Remove []PrincipalRole `json:"remove"`
	}{

		Modify: m.modifyField,

		Remove: m.removeField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this project principals assignment
func (m *ProjectPrincipalsAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModify(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemove(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectPrincipalsAssignment) validateModify(formats strfmt.Registry) error {
	if swag.IsZero(m.Modify()) { // not required
		return nil
	}

	for i := 0; i < len(m.Modify()); i++ {

		if err := m.modifyField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modify" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProjectPrincipalsAssignment) validateRemove(formats strfmt.Registry) error {
	if swag.IsZero(m.Remove()) { // not required
		return nil
	}

	for i := 0; i < len(m.Remove()); i++ {

		if err := m.removeField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remove" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this project principals assignment based on the context it is used
func (m *ProjectPrincipalsAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModify(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemove(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectPrincipalsAssignment) contextValidateModify(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Modify()); i++ {

		if err := m.modifyField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modify" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ProjectPrincipalsAssignment) contextValidateRemove(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Remove()); i++ {

		if err := m.removeField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remove" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectPrincipalsAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectPrincipalsAssignment) UnmarshalBinary(b []byte) error {
	var res ProjectPrincipalsAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
