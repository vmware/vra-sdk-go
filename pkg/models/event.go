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

// Event Event
//
// Represents request events.
//
// swagger:model Event
type Event struct {

	// Longer user-friendly details of the event.
	Details string `json:"details,omitempty"`

	// Indicates whether the event has logs or not.
	HasLogs bool `json:"hasLogs,omitempty"`

	// Event identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Short user-friendly label of the event (e.g. 'shutting down myVM')
	// Required: true
	Name *string `json:"name"`

	// Optional resource name to which the event applies to
	ResourceName string `json:"resourceName,omitempty"`

	// Optional resource type to which the event applies to
	ResourceType string `json:"resourceType,omitempty"`

	// Timestamp of the Event (e.g. date format '2019-07-13T23:16:49.310Z').
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`

	// Indicates if the event represents user input.
	UserEvent bool `json:"userEvent,omitempty"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Event) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event based on context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
