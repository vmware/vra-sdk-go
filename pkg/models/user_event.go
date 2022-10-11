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

// UserEvent UserEvent
//
// Represents events which require user input.
//
// swagger:model UserEvent
type UserEvent struct {

	// Time at which the event was created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Details of the user event. If the event is a user interaction the details contain the vRO Manual User Interaction ID.
	Details string `json:"details,omitempty"`

	// Id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// User event inputs. If user event is a user interaction the inputs form can be retrieved from Orchestrator Gateway API /vro/forms/user-interaction/{userInteractionId}
	Inputs interface{} `json:"inputs,omitempty"`

	// Short user-friendly label of the event
	// Required: true
	Name *string `json:"name"`

	// User event outputs
	Outputs interface{} `json:"outputs,omitempty"`

	// Request id
	// Required: true
	// Format: uuid
	RequestID *strfmt.UUID `json:"requestId"`

	// User that initiated the event
	// Required: true
	RequestedBy *string `json:"requestedBy"`

	// Resource id
	// Required: true
	ResourceIds []strfmt.UUID `json:"resourceIds"`

	// User event status.
	// Required: true
	// Enum: [SUCCESSFUL FAILED PENDING]
	Status *string `json:"status"`

	// Time at which the event was updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this user event
func (m *UserEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceIds(formats); err != nil {
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

func (m *UserEvent) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserEvent) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserEvent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UserEvent) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("requestId", "body", m.RequestID); err != nil {
		return err
	}

	if err := validate.FormatOf("requestId", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserEvent) validateRequestedBy(formats strfmt.Registry) error {

	if err := validate.Required("requestedBy", "body", m.RequestedBy); err != nil {
		return err
	}

	return nil
}

func (m *UserEvent) validateResourceIds(formats strfmt.Registry) error {

	if err := validate.Required("resourceIds", "body", m.ResourceIds); err != nil {
		return err
	}

	for i := 0; i < len(m.ResourceIds); i++ {

		if err := validate.FormatOf("resourceIds"+"."+strconv.Itoa(i), "body", "uuid", m.ResourceIds[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

var userEventTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESSFUL","FAILED","PENDING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userEventTypeStatusPropEnum = append(userEventTypeStatusPropEnum, v)
	}
}

const (

	// UserEventStatusSUCCESSFUL captures enum value "SUCCESSFUL"
	UserEventStatusSUCCESSFUL string = "SUCCESSFUL"

	// UserEventStatusFAILED captures enum value "FAILED"
	UserEventStatusFAILED string = "FAILED"

	// UserEventStatusPENDING captures enum value "PENDING"
	UserEventStatusPENDING string = "PENDING"
)

// prop value enum
func (m *UserEvent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userEventTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserEvent) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *UserEvent) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user event based on context it is used
func (m *UserEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserEvent) UnmarshalBinary(b []byte) error {
	var res UserEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
