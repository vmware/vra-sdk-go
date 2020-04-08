// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MeteringPolicyAssignment MeteringPolicyAssignment
//
// Pricing card assignment for project/cloud zone
// swagger:model MeteringPolicyAssignment
type MeteringPolicyAssignment struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Pricing card assigned entity id
	EntityID string `json:"entityId,omitempty"`

	// Pricing card assigned entity name
	EntityName string `json:"entityName,omitempty"`

	// Pricing card assigned entity type
	// Enum: [ALL PROJECT CLOUDZONE]
	EntityType string `json:"entityType,omitempty"`

	// Id of the pricingCardAssignment
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Updated time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// OrgId of the pricingCardAssignment
	OrgID string `json:"orgId,omitempty"`

	// Pricing card id
	// Format: uuid
	PricingCardID strfmt.UUID `json:"pricingCardId,omitempty"`

	// Pricing card name
	PricingCardName string `json:"pricingCardName,omitempty"`
}

// Validate validates this metering policy assignment
func (m *MeteringPolicyAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePricingCardID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MeteringPolicyAssignment) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var meteringPolicyAssignmentTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","PROJECT","CLOUDZONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		meteringPolicyAssignmentTypeEntityTypePropEnum = append(meteringPolicyAssignmentTypeEntityTypePropEnum, v)
	}
}

const (

	// MeteringPolicyAssignmentEntityTypeALL captures enum value "ALL"
	MeteringPolicyAssignmentEntityTypeALL string = "ALL"

	// MeteringPolicyAssignmentEntityTypePROJECT captures enum value "PROJECT"
	MeteringPolicyAssignmentEntityTypePROJECT string = "PROJECT"

	// MeteringPolicyAssignmentEntityTypeCLOUDZONE captures enum value "CLOUDZONE"
	MeteringPolicyAssignmentEntityTypeCLOUDZONE string = "CLOUDZONE"
)

// prop value enum
func (m *MeteringPolicyAssignment) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, meteringPolicyAssignmentTypeEntityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MeteringPolicyAssignment) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *MeteringPolicyAssignment) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MeteringPolicyAssignment) validateLastUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedAt", "body", "date-time", m.LastUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MeteringPolicyAssignment) validatePricingCardID(formats strfmt.Registry) error {

	if swag.IsZero(m.PricingCardID) { // not required
		return nil
	}

	if err := validate.FormatOf("pricingCardId", "body", "uuid", m.PricingCardID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MeteringPolicyAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeteringPolicyAssignment) UnmarshalBinary(b []byte) error {
	var res MeteringPolicyAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
