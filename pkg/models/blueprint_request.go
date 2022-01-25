// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BlueprintRequest BlueprintRequest
//
// swagger:model BlueprintRequest
type BlueprintRequest struct {

	// Blueprint Id
	// Format: uuid
	BlueprintID strfmt.UUID `json:"blueprintId,omitempty"`

	// Blueprint version
	BlueprintVersion string `json:"blueprintVersion,omitempty"`

	// Cancel request time
	// Read Only: true
	// Format: date-time
	CancelRequestedAt strfmt.DateTime `json:"cancelRequestedAt,omitempty"`

	// Cancel requested by
	// Read Only: true
	CancelRequestedBy string `json:"cancelRequestedBy,omitempty"`

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Existing deployment Id
	DeploymentID string `json:"deploymentId,omitempty"`

	// Name for the new deployment
	DeploymentName string `json:"deploymentName,omitempty"`

	// Description for the new request
	Description string `json:"description,omitempty"`

	// Destroy existing deployment
	Destroy bool `json:"destroy,omitempty"`

	// Failure message
	// Read Only: true
	FailureMessage string `json:"failureMessage,omitempty"`

	// Flow execution Id
	// Read Only: true
	FlowExecutionID string `json:"flowExecutionId,omitempty"`

	// Flow Id
	// Read Only: true
	FlowID string `json:"flowId,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Ignore delete failures in blueprint request
	IgnoreDeleteFailures bool `json:"ignoreDeleteFailures,omitempty"`

	// Blueprint request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Plan only without affecting existing deployment
	Plan bool `json:"plan,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Reason for requesting a blueprint
	Reason string `json:"reason,omitempty"`

	// Request tracker Id
	// Read Only: true
	RequestTrackerID string `json:"requestTrackerId,omitempty"`

	// Simulate blueprint request with providers
	Simulate bool `json:"simulate,omitempty"`

	// Status
	// Read Only: true
	// Enum: [CREATED STARTED FINISHED FAILED CANCELLED]
	Status string `json:"status,omitempty"`

	// Target resources
	// Read Only: true
	TargetResources []string `json:"targetResources"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Validation messages
	// Read Only: true
	ValidationMessages []*BlueprintValidationMessage `json:"validationMessages"`
}

// Validate validates this blueprint request
func (m *BlueprintRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlueprintID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCancelRequestedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintRequest) validateBlueprintID(formats strfmt.Registry) error {
	if swag.IsZero(m.BlueprintID) { // not required
		return nil
	}

	if err := validate.FormatOf("blueprintId", "body", "uuid", m.BlueprintID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) validateCancelRequestedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CancelRequestedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("cancelRequestedAt", "body", "date-time", m.CancelRequestedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var blueprintRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATED","STARTED","FINISHED","FAILED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blueprintRequestTypeStatusPropEnum = append(blueprintRequestTypeStatusPropEnum, v)
	}
}

const (

	// BlueprintRequestStatusCREATED captures enum value "CREATED"
	BlueprintRequestStatusCREATED string = "CREATED"

	// BlueprintRequestStatusSTARTED captures enum value "STARTED"
	BlueprintRequestStatusSTARTED string = "STARTED"

	// BlueprintRequestStatusFINISHED captures enum value "FINISHED"
	BlueprintRequestStatusFINISHED string = "FINISHED"

	// BlueprintRequestStatusFAILED captures enum value "FAILED"
	BlueprintRequestStatusFAILED string = "FAILED"

	// BlueprintRequestStatusCANCELLED captures enum value "CANCELLED"
	BlueprintRequestStatusCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *BlueprintRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, blueprintRequestTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BlueprintRequest) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) validateValidationMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationMessages) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidationMessages); i++ {
		if swag.IsZero(m.ValidationMessages[i]) { // not required
			continue
		}

		if m.ValidationMessages[i] != nil {
			if err := m.ValidationMessages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationMessages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this blueprint request based on the context it is used
func (m *BlueprintRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCancelRequestedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCancelRequestedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFailureMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlowExecutionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlowID(ctx, formats); err != nil {
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

	if err := m.contextValidateRequestTrackerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTargetResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidationMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintRequest) contextValidateCancelRequestedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cancelRequestedAt", "body", strfmt.DateTime(m.CancelRequestedAt)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateCancelRequestedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cancelRequestedBy", "body", string(m.CancelRequestedBy)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateFailureMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "failureMessage", "body", string(m.FailureMessage)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateFlowExecutionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "flowExecutionId", "body", string(m.FlowExecutionID)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateFlowID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "flowId", "body", string(m.FlowID)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateOrgID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "orgId", "body", string(m.OrgID)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateProjectName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "projectName", "body", string(m.ProjectName)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateRequestTrackerID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "requestTrackerId", "body", string(m.RequestTrackerID)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateTargetResources(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "targetResources", "body", []string(m.TargetResources)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateUpdatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedBy", "body", string(m.UpdatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintRequest) contextValidateValidationMessages(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "validationMessages", "body", []*BlueprintValidationMessage(m.ValidationMessages)); err != nil {
		return err
	}

	for i := 0; i < len(m.ValidationMessages); i++ {

		if m.ValidationMessages[i] != nil {
			if err := m.ValidationMessages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validationMessages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validationMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlueprintRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintRequest) UnmarshalBinary(b []byte) error {
	var res BlueprintRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
