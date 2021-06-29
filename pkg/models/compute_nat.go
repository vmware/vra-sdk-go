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

// ComputeNat The compute nat object is used to configure NAT rules on the Edge Gateway or Tier-1 logical router in NSX to enable port forwarding.<br>**HATEOAS** links:<br>**self** - Nat - Self link to this nat
//
// swagger:model ComputeNat
type ComputeNat struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Set of ids of the cloud accounts this resource belongs to.
	// Example: [9e49]
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base resource.
	// Example: { \"property\" : \"value\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// Deployment id that is associated with this resource.
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The external regionId of the resource.
	// Example: us-east-1
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// The external zoneId of the resource.
	// Example: us-east-1a
	// Required: true
	ExternalZoneID *string `json:"externalZoneId"`

	// The compute gateway to which the compute nat is attached
	// Required: true
	Gateway *string `json:"gateway"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// A list of NAT rule(s) to be created on the NSX network
	// Required: true
	NatRules []*NatRule `json:"natRules"`

	// The id of the organization this entity belongs to.
	// Example: 9e49
	OrgID string `json:"orgId,omitempty"`

	// This field is deprecated. Use orgId instead. The id of the organization this entity belongs to.
	// Example: deprecated
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// The id of the project this resource belongs to.
	// Example: 9e49
	ProjectID string `json:"projectId,omitempty"`

	// A set of tag keys and optional values that were set on this resource.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this compute nat
func (m *ComputeNat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeNat) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ComputeNat) validateCloudAccountIds(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudAccountIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("cloudAccountIds", "body", m.CloudAccountIds); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNat) validateExternalRegionID(formats strfmt.Registry) error {

	if err := validate.Required("externalRegionId", "body", m.ExternalRegionID); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNat) validateExternalZoneID(formats strfmt.Registry) error {

	if err := validate.Required("externalZoneId", "body", m.ExternalZoneID); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNat) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNat) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNat) validateNatRules(formats strfmt.Registry) error {

	if err := validate.Required("natRules", "body", m.NatRules); err != nil {
		return err
	}

	for i := 0; i < len(m.NatRules); i++ {
		if swag.IsZero(m.NatRules[i]) { // not required
			continue
		}

		if m.NatRules[i] != nil {
			if err := m.NatRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("natRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeNat) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this compute nat based on the context it is used
func (m *ComputeNat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNatRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeNat) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ComputeNat) contextValidateNatRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NatRules); i++ {

		if m.NatRules[i] != nil {
			if err := m.NatRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("natRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeNat) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputeNat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeNat) UnmarshalBinary(b []byte) error {
	var res ComputeNat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}