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

// UserOpResource UserOpResource
//
// swagger:model UserOpResource
type UserOpResource struct {

	// approver groups
	ApproverGroups []string `json:"approverGroups"`

	// approvers
	Approvers []string `json:"approvers"`

	// cancel previous pending user op
	CancelPreviousPendingUserOp bool `json:"cancelPreviousPendingUserOp,omitempty"`

	// change log
	ChangeLog string `json:"changeLog,omitempty"`

	// comments
	Comments string `json:"comments,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// execution Id
	ExecutionID string `json:"executionId,omitempty"`

	// execution index
	ExecutionIndex int64 `json:"executionIndex,omitempty"`

	// expiration
	Expiration int32 `json:"expiration,omitempty"`

	// expiration in days
	ExpirationInDays int32 `json:"expirationInDays,omitempty"`

	// expiration in seconds
	ExpirationInSeconds int64 `json:"expirationInSeconds,omitempty"`

	// expiration unit
	// Enum: [MINUTES HOURS DAYS]
	ExpirationUnit string `json:"expirationUnit,omitempty"`

	// expires on in micros
	ExpiresOnInMicros int64 `json:"expiresOnInMicros,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// index
	Index string `json:"index,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// parent Id
	ParentID string `json:"parentId,omitempty"`

	// pipeline
	Pipeline string `json:"pipeline,omitempty"`

	// pipeline Id
	PipelineID string `json:"pipelineId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// requested by
	RequestedBy string `json:"requestedBy,omitempty"`

	// requested on in micros
	RequestedOnInMicros int64 `json:"requestedOnInMicros,omitempty"`

	// responded by
	RespondedBy string `json:"respondedBy,omitempty"`

	// responded by email
	RespondedByEmail string `json:"respondedByEmail,omitempty"`

	// responded on in micros
	RespondedOnInMicros int64 `json:"respondedOnInMicros,omitempty"`

	// responder roles
	ResponderRoles []string `json:"responderRoles"`

	// responder token
	ResponderToken string `json:"responderToken,omitempty"`

	// resume count
	ResumeCount int32 `json:"resumeCount,omitempty"`

	// sendemail
	Sendemail bool `json:"sendemail,omitempty"`

	// stage key
	StageKey string `json:"stageKey,omitempty"`

	// status
	// Enum: [ACTIVE APPROVED CANCELED EXPIRED REJECTED]
	Status string `json:"status,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// task execution Id
	TaskExecutionID string `json:"taskExecutionId,omitempty"`

	// task key
	TaskKey string `json:"taskKey,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this user op resource
func (m *UserOpResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserOpResource) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var userOpResourceTypeExpirationUnitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MINUTES","HOURS","DAYS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userOpResourceTypeExpirationUnitPropEnum = append(userOpResourceTypeExpirationUnitPropEnum, v)
	}
}

const (

	// UserOpResourceExpirationUnitMINUTES captures enum value "MINUTES"
	UserOpResourceExpirationUnitMINUTES string = "MINUTES"

	// UserOpResourceExpirationUnitHOURS captures enum value "HOURS"
	UserOpResourceExpirationUnitHOURS string = "HOURS"

	// UserOpResourceExpirationUnitDAYS captures enum value "DAYS"
	UserOpResourceExpirationUnitDAYS string = "DAYS"
)

// prop value enum
func (m *UserOpResource) validateExpirationUnitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userOpResourceTypeExpirationUnitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserOpResource) validateExpirationUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpirationUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateExpirationUnitEnum("expirationUnit", "body", m.ExpirationUnit); err != nil {
		return err
	}

	return nil
}

var userOpResourceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","APPROVED","CANCELED","EXPIRED","REJECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userOpResourceTypeStatusPropEnum = append(userOpResourceTypeStatusPropEnum, v)
	}
}

const (

	// UserOpResourceStatusACTIVE captures enum value "ACTIVE"
	UserOpResourceStatusACTIVE string = "ACTIVE"

	// UserOpResourceStatusAPPROVED captures enum value "APPROVED"
	UserOpResourceStatusAPPROVED string = "APPROVED"

	// UserOpResourceStatusCANCELED captures enum value "CANCELED"
	UserOpResourceStatusCANCELED string = "CANCELED"

	// UserOpResourceStatusEXPIRED captures enum value "EXPIRED"
	UserOpResourceStatusEXPIRED string = "EXPIRED"

	// UserOpResourceStatusREJECTED captures enum value "REJECTED"
	UserOpResourceStatusREJECTED string = "REJECTED"
)

// prop value enum
func (m *UserOpResource) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userOpResourceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserOpResource) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *UserOpResource) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user op resource based on context it is used
func (m *UserOpResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserOpResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserOpResource) UnmarshalBinary(b []byte) error {
	var res UserOpResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
