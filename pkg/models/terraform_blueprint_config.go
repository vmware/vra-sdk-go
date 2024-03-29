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

// TerraformBlueprintConfig TerraformBlueprintConfig
//
// swagger:model TerraformBlueprintConfig
type TerraformBlueprintConfig struct {

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Blueprint description.
	Description string `json:"description,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Blueprint name.
	Name string `json:"name,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Terraform configuration to blueprint mapping.
	TerraformToBlueprintMapping *TerraformToBlueprintMapping `json:"terraformToBlueprintMapping,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Terraform runtime version.
	Version string `json:"version,omitempty"`
}

// Validate validates this terraform blueprint config
func (m *TerraformBlueprintConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerraformToBlueprintMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformBlueprintConfig) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TerraformBlueprintConfig) validateTerraformToBlueprintMapping(formats strfmt.Registry) error {
	if swag.IsZero(m.TerraformToBlueprintMapping) { // not required
		return nil
	}

	if m.TerraformToBlueprintMapping != nil {
		if err := m.TerraformToBlueprintMapping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraformToBlueprintMapping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraformToBlueprintMapping")
			}
			return err
		}
	}

	return nil
}

func (m *TerraformBlueprintConfig) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this terraform blueprint config based on the context it is used
func (m *TerraformBlueprintConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrgID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTerraformToBlueprintMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraformBlueprintConfig) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *TerraformBlueprintConfig) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *TerraformBlueprintConfig) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *TerraformBlueprintConfig) contextValidateOrgID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "orgId", "body", string(m.OrgID)); err != nil {
		return err
	}

	return nil
}

func (m *TerraformBlueprintConfig) contextValidateProjectName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "projectName", "body", string(m.ProjectName)); err != nil {
		return err
	}

	return nil
}

func (m *TerraformBlueprintConfig) contextValidateTerraformToBlueprintMapping(ctx context.Context, formats strfmt.Registry) error {

	if m.TerraformToBlueprintMapping != nil {
		if err := m.TerraformToBlueprintMapping.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terraformToBlueprintMapping")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("terraformToBlueprintMapping")
			}
			return err
		}
	}

	return nil
}

func (m *TerraformBlueprintConfig) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *TerraformBlueprintConfig) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedBy", "body", string(m.UpdatedBy)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraformBlueprintConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraformBlueprintConfig) UnmarshalBinary(b []byte) error {
	var res TerraformBlueprintConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
