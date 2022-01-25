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

// Resource Resource
//
// A resource
//
// swagger:model Resource
type Resource struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Current ongoing request on the resource
	CurrentRequest *Request `json:"currentRequest,omitempty"`

	// Deployment to which resource belongs
	Deployment *DeploymentReference `json:"deployment,omitempty"`

	// Resource deployment id
	// Format: uuid
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`

	// A description of the resource
	Description string `json:"description,omitempty"`

	// Expense associated with the deployment.
	// Read Only: true
	Expense *Expense `json:"expense,omitempty"`

	// Unique identifier of the resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the resource
	// Required: true
	Name *string `json:"name"`

	// Resource org id
	OrgID string `json:"orgId,omitempty"`

	// Origin of the resource
	// Enum: [DISCOVERED ONBOARDED MIGRATED]
	Origin string `json:"origin,omitempty"`

	// Project to which resource's deployment belongs
	Project *ResourceReference `json:"project,omitempty"`

	// Resource project id
	ProjectID string `json:"projectId,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`

	// The current sync status
	// Enum: [SUCCESS MISSING STALE]
	SyncStatus string `json:"syncStatus,omitempty"`

	// Type of the resource
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeployment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeploymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateCurrentRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentRequest) { // not required
		return nil
	}

	if m.CurrentRequest != nil {
		if err := m.CurrentRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentRequest")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateDeployment(formats strfmt.Registry) error {
	if swag.IsZero(m.Deployment) { // not required
		return nil
	}

	if m.Deployment != nil {
		if err := m.Deployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateDeploymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentID) { // not required
		return nil
	}

	if err := validate.FormatOf("deploymentId", "body", "uuid", m.DeploymentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateExpense(formats strfmt.Registry) error {
	if swag.IsZero(m.Expense) { // not required
		return nil
	}

	if m.Expense != nil {
		if err := m.Expense.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expense")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expense")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var resourceTypeOriginPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISCOVERED","ONBOARDED","MIGRATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeOriginPropEnum = append(resourceTypeOriginPropEnum, v)
	}
}

const (

	// ResourceOriginDISCOVERED captures enum value "DISCOVERED"
	ResourceOriginDISCOVERED string = "DISCOVERED"

	// ResourceOriginONBOARDED captures enum value "ONBOARDED"
	ResourceOriginONBOARDED string = "ONBOARDED"

	// ResourceOriginMIGRATED captures enum value "MIGRATED"
	ResourceOriginMIGRATED string = "MIGRATED"
)

// prop value enum
func (m *Resource) validateOriginEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resourceTypeOriginPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Resource) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	// value enum
	if err := m.validateOriginEnum("origin", "body", m.Origin); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateProject(formats strfmt.Registry) error {
	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

var resourceTypeSyncStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","MISSING","STALE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeSyncStatusPropEnum = append(resourceTypeSyncStatusPropEnum, v)
	}
}

const (

	// ResourceSyncStatusSUCCESS captures enum value "SUCCESS"
	ResourceSyncStatusSUCCESS string = "SUCCESS"

	// ResourceSyncStatusMISSING captures enum value "MISSING"
	ResourceSyncStatusMISSING string = "MISSING"

	// ResourceSyncStatusSTALE captures enum value "STALE"
	ResourceSyncStatusSTALE string = "STALE"
)

// prop value enum
func (m *Resource) validateSyncStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resourceTypeSyncStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Resource) validateSyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSyncStatusEnum("syncStatus", "body", m.SyncStatus); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this resource based on the context it is used
func (m *Resource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrentRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeployment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExpense(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resource) contextValidateCurrentRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentRequest != nil {
		if err := m.CurrentRequest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentRequest")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) contextValidateDeployment(ctx context.Context, formats strfmt.Registry) error {

	if m.Deployment != nil {
		if err := m.Deployment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deployment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deployment")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) contextValidateExpense(ctx context.Context, formats strfmt.Registry) error {

	if m.Expense != nil {
		if err := m.Expense.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expense")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("expense")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) contextValidateProject(ctx context.Context, formats strfmt.Registry) error {

	if m.Project != nil {
		if err := m.Project.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
