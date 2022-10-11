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

// ResourceResponse ResourceResponse
//
// swagger:model ResourceResponse
type ResourceResponse struct {

	// failure message
	FailureMessage string `json:"failureMessage,omitempty"`

	// progress message
	ProgressMessage string `json:"progressMessage,omitempty"`

	// resource link
	ResourceLink string `json:"resourceLink,omitempty"`

	// resource properties
	ResourceProperties interface{} `json:"resourceProperties,omitempty"`

	// resource request Id
	ResourceRequestID string `json:"resourceRequestId,omitempty"`

	// status
	// Enum: [CREATED STARTED FINISHED FAILED CANCELLED UNKNOWN]
	Status string `json:"status,omitempty"`
}

// Validate validates this resource response
func (m *ResourceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var resourceResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATED","STARTED","FINISHED","FAILED","CANCELLED","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceResponseTypeStatusPropEnum = append(resourceResponseTypeStatusPropEnum, v)
	}
}

const (

	// ResourceResponseStatusCREATED captures enum value "CREATED"
	ResourceResponseStatusCREATED string = "CREATED"

	// ResourceResponseStatusSTARTED captures enum value "STARTED"
	ResourceResponseStatusSTARTED string = "STARTED"

	// ResourceResponseStatusFINISHED captures enum value "FINISHED"
	ResourceResponseStatusFINISHED string = "FINISHED"

	// ResourceResponseStatusFAILED captures enum value "FAILED"
	ResourceResponseStatusFAILED string = "FAILED"

	// ResourceResponseStatusCANCELLED captures enum value "CANCELLED"
	ResourceResponseStatusCANCELLED string = "CANCELLED"

	// ResourceResponseStatusUNKNOWN captures enum value "UNKNOWN"
	ResourceResponseStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *ResourceResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resourceResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResourceResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resource response based on context it is used
func (m *ResourceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceResponse) UnmarshalBinary(b []byte) error {
	var res ResourceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}