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

// ContentSource ContentSource
//
// swagger:model ContentSource
type ContentSource struct {

	// Source custom configuration
	// Example: {"branch":"string","contentType":"string","endpointId":"string","integrationId":"string","path":"string","repository":"string","requestScopeOrg":"boolean"}
	Config interface{} `json:"config,omitempty"`

	// Creation time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created By
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Content Source description
	Description string `json:"description,omitempty"`

	// Content Source id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Update time
	// Read Only: true
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Updated By
	// Read Only: true
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// Content Source name
	// Required: true
	Name *string `json:"name"`

	// Associated org
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Associated projects
	// Required: true
	ProjectID *string `json:"projectId"`

	// Is Sync Enabled
	// Example: false
	SyncEnabled bool `json:"syncEnabled,omitempty"`

	// Content Source type
	// Required: true
	// Enum: [com.github com.github.enterprise com.gitlab org.bitbucket com.vmware.marketplace]
	TypeID *string `json:"typeId"`
}

// Validate validates this content source
func (m *ContentSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentSource) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) validateLastUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedAt", "body", "date-time", m.LastUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

var contentSourceTypeTypeIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["com.github","com.github.enterprise","com.gitlab","org.bitbucket","com.vmware.marketplace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentSourceTypeTypeIDPropEnum = append(contentSourceTypeTypeIDPropEnum, v)
	}
}

const (

	// ContentSourceTypeIDComDotGithub captures enum value "com.github"
	ContentSourceTypeIDComDotGithub string = "com.github"

	// ContentSourceTypeIDComDotGithubDotEnterprise captures enum value "com.github.enterprise"
	ContentSourceTypeIDComDotGithubDotEnterprise string = "com.github.enterprise"

	// ContentSourceTypeIDComDotGitlab captures enum value "com.gitlab"
	ContentSourceTypeIDComDotGitlab string = "com.gitlab"

	// ContentSourceTypeIDOrgDotBitbucket captures enum value "org.bitbucket"
	ContentSourceTypeIDOrgDotBitbucket string = "org.bitbucket"

	// ContentSourceTypeIDComDotVmwareDotMarketplace captures enum value "com.vmware.marketplace"
	ContentSourceTypeIDComDotVmwareDotMarketplace string = "com.vmware.marketplace"
)

// prop value enum
func (m *ContentSource) validateTypeIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contentSourceTypeTypeIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContentSource) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("typeId", "body", m.TypeID); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeIDEnum("typeId", "body", *m.TypeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this content source based on the context it is used
func (m *ContentSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOrgID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentSource) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) contextValidateLastUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedAt", "body", strfmt.DateTime(m.LastUpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) contextValidateLastUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastUpdatedBy", "body", string(m.LastUpdatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *ContentSource) contextValidateOrgID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "orgId", "body", string(m.OrgID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentSource) UnmarshalBinary(b []byte) error {
	var res ContentSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
