// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PlanExecutionResponse plan execution response
//
// swagger:model PlanExecutionResponse
type PlanExecutionResponse struct {

	// document self link
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// failure message
	FailureMessage string `json:"failureMessage,omitempty"`

	// plan link
	PlanLink string `json:"planLink,omitempty"`

	// task info
	TaskInfo *TaskState `json:"taskInfo,omitempty"`

	// Link to a tenant
	TenantLink string `json:"tenantLink,omitempty"`
}

// Validate validates this plan execution response
func (m *PlanExecutionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaskInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanExecutionResponse) validateTaskInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskInfo) { // not required
		return nil
	}

	if m.TaskInfo != nil {
		if err := m.TaskInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taskInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan execution response based on the context it is used
func (m *PlanExecutionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaskInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlanExecutionResponse) contextValidateTaskInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.TaskInfo != nil {
		if err := m.TaskInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taskInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlanExecutionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlanExecutionResponse) UnmarshalBinary(b []byte) error {
	var res PlanExecutionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
