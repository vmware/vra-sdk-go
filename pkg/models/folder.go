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

// Folder Represents a structure that holds details of vSphere folders**HATEOAS** links:<br>**self** - Folder - Self link to this folder
//
// swagger:model Folder
type Folder struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// A set of cloud account IDs that are associated with the particular Folder
	// Example: [ \"cloud-account-id-1\", \"cloud-account-id-2\" } ]
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"folderType\" : \"1\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user or display name of the group that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Type of a owner(user/ad_group) that owns the entity.
	// Example: ad_group
	OwnerType string `json:"ownerType,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this folder
func (m *Folder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Folder) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_links" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_links" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *Folder) validateCloudAccountIds(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudAccountIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("cloudAccountIds", "body", m.CloudAccountIds); err != nil {
		return err
	}

	return nil
}

func (m *Folder) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this folder based on the context it is used
func (m *Folder) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Folder) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if val, ok := m.Links[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Folder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Folder) UnmarshalBinary(b []byte) error {
	var res Folder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
