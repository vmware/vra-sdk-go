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

// ResourceRequestResponse ResourceRequestResponse
//
// # Resource request response
//
// swagger:model ResourceRequestResponse
type ResourceRequestResponse struct {

	// Identifier of the requested deployment id to which the request applies to
	// Format: uuid
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`

	// Project identifier
	ProjectID string `json:"projectId,omitempty"`

	// Request identifier
	// Format: uuid
	RequestID strfmt.UUID `json:"requestId,omitempty"`

	// Resource ID
	ResourceID string `json:"resourceId,omitempty"`
}

// Validate validates this resource request response
func (m *ResourceRequestResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeploymentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceRequestResponse) validateDeploymentID(formats strfmt.Registry) error {
	if swag.IsZero(m.DeploymentID) { // not required
		return nil
	}

	if err := validate.FormatOf("deploymentId", "body", "uuid", m.DeploymentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResourceRequestResponse) validateRequestID(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("requestId", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resource request response based on context it is used
func (m *ResourceRequestResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceRequestResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceRequestResponse) UnmarshalBinary(b []byte) error {
	var res ResourceRequestResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
