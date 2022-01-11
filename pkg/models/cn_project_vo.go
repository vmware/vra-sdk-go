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

// CnProjectVo cn project vo
//
// swagger:model CnProjectVo
type CnProjectVo struct {

	// active
	Active bool `json:"active,omitempty"`

	// default org
	DefaultOrg bool `json:"defaultOrg,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this cn project vo
func (m *CnProjectVo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CnProjectVo) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cn project vo based on context it is used
func (m *CnProjectVo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CnProjectVo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CnProjectVo) UnmarshalBinary(b []byte) error {
	var res CnProjectVo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
