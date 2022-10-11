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

// SupervisorNamespaceSyncResponseDTO SupervisorNamespaceSyncResponseDTO
//
// swagger:model SupervisorNamespaceSyncResponseDTO
type SupervisorNamespaceSyncResponseDTO struct {

	// message
	Message string `json:"message,omitempty"`

	// operation tracker link
	OperationTrackerLink string `json:"operationTrackerLink,omitempty"`

	// status
	// Enum: [IN_PROGRESS COMPLETED FAILED]
	Status string `json:"status,omitempty"`
}

// Validate validates this supervisor namespace sync response d t o
func (m *SupervisorNamespaceSyncResponseDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var supervisorNamespaceSyncResponseDTOTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","COMPLETED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supervisorNamespaceSyncResponseDTOTypeStatusPropEnum = append(supervisorNamespaceSyncResponseDTOTypeStatusPropEnum, v)
	}
}

const (

	// SupervisorNamespaceSyncResponseDTOStatusINPROGRESS captures enum value "IN_PROGRESS"
	SupervisorNamespaceSyncResponseDTOStatusINPROGRESS string = "IN_PROGRESS"

	// SupervisorNamespaceSyncResponseDTOStatusCOMPLETED captures enum value "COMPLETED"
	SupervisorNamespaceSyncResponseDTOStatusCOMPLETED string = "COMPLETED"

	// SupervisorNamespaceSyncResponseDTOStatusFAILED captures enum value "FAILED"
	SupervisorNamespaceSyncResponseDTOStatusFAILED string = "FAILED"
)

// prop value enum
func (m *SupervisorNamespaceSyncResponseDTO) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, supervisorNamespaceSyncResponseDTOTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SupervisorNamespaceSyncResponseDTO) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this supervisor namespace sync response d t o based on context it is used
func (m *SupervisorNamespaceSyncResponseDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupervisorNamespaceSyncResponseDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupervisorNamespaceSyncResponseDTO) UnmarshalBinary(b []byte) error {
	var res SupervisorNamespaceSyncResponseDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}