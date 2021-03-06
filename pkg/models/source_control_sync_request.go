// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SourceControlSyncRequest SourceControlSyncRequest
//
// swagger:model SourceControlSyncRequest
type SourceControlSyncRequest struct {

	// Created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Last Updated at
	// Read Only: true
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Message
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Project Id
	// Read Only: true
	ProjectID string `json:"projectId,omitempty"`

	// Request Id
	// Read Only: true
	// Format: uuid
	RequestID strfmt.UUID `json:"requestId,omitempty"`

	// Content Source Id
	// Format: uuid
	SourceID strfmt.UUID `json:"sourceId,omitempty"`

	// Status
	// Read Only: true
	// Enum: [REQUESTED STARTED PROCESSING COMPLETED FAILED SKIPPED]
	Status string `json:"status,omitempty"`
}

// Validate validates this source control sync request
func (m *SourceControlSyncRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceControlSyncRequest) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceControlSyncRequest) validateLastUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedAt", "body", "date-time", m.LastUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceControlSyncRequest) validateRequestID(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("requestId", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceControlSyncRequest) validateSourceID(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceID) { // not required
		return nil
	}

	if err := validate.FormatOf("sourceId", "body", "uuid", m.SourceID.String(), formats); err != nil {
		return err
	}

	return nil
}

var sourceControlSyncRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","STARTED","PROCESSING","COMPLETED","FAILED","SKIPPED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sourceControlSyncRequestTypeStatusPropEnum = append(sourceControlSyncRequestTypeStatusPropEnum, v)
	}
}

const (

	// SourceControlSyncRequestStatusREQUESTED captures enum value "REQUESTED"
	SourceControlSyncRequestStatusREQUESTED string = "REQUESTED"

	// SourceControlSyncRequestStatusSTARTED captures enum value "STARTED"
	SourceControlSyncRequestStatusSTARTED string = "STARTED"

	// SourceControlSyncRequestStatusPROCESSING captures enum value "PROCESSING"
	SourceControlSyncRequestStatusPROCESSING string = "PROCESSING"

	// SourceControlSyncRequestStatusCOMPLETED captures enum value "COMPLETED"
	SourceControlSyncRequestStatusCOMPLETED string = "COMPLETED"

	// SourceControlSyncRequestStatusFAILED captures enum value "FAILED"
	SourceControlSyncRequestStatusFAILED string = "FAILED"

	// SourceControlSyncRequestStatusSKIPPED captures enum value "SKIPPED"
	SourceControlSyncRequestStatusSKIPPED string = "SKIPPED"
)

// prop value enum
func (m *SourceControlSyncRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sourceControlSyncRequestTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SourceControlSyncRequest) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourceControlSyncRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceControlSyncRequest) UnmarshalBinary(b []byte) error {
	var res SourceControlSyncRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
