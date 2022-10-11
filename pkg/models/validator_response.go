// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ValidatorResponse ValidatorResponse
//
// Model used for validation responses.
//
// swagger:discriminator ValidatorResponse Model used for validation responses.
type ValidatorResponse interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Class which will be validated.
	// Example: Pipeline
	ClassType() string
	SetClassType(string)

	Message() ValidationMessage
	SetMessage(ValidationMessage)

	// Indicates whether validation was successful or not.
	// Example: false
	Success() bool
	SetSuccess(bool)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type validatorResponse struct {
	classTypeField string

	messageField ValidationMessage

	successField bool
}

// ClassType gets the class type of this polymorphic type
func (m *validatorResponse) ClassType() string {
	return m.classTypeField
}

// SetClassType sets the class type of this polymorphic type
func (m *validatorResponse) SetClassType(val string) {
	m.classTypeField = val
}

// Message gets the message of this polymorphic type
func (m *validatorResponse) Message() ValidationMessage {
	return m.messageField
}

// SetMessage sets the message of this polymorphic type
func (m *validatorResponse) SetMessage(val ValidationMessage) {
	m.messageField = val
}

// Success gets the success of this polymorphic type
func (m *validatorResponse) Success() bool {
	return m.successField
}

// SetSuccess sets the success of this polymorphic type
func (m *validatorResponse) SetSuccess(val bool) {
	m.successField = val
}

// UnmarshalValidatorResponseSlice unmarshals polymorphic slices of ValidatorResponse
func UnmarshalValidatorResponseSlice(reader io.Reader, consumer runtime.Consumer) ([]ValidatorResponse, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ValidatorResponse
	for _, element := range elements {
		obj, err := unmarshalValidatorResponse(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalValidatorResponse unmarshals polymorphic ValidatorResponse
func UnmarshalValidatorResponse(reader io.Reader, consumer runtime.Consumer) (ValidatorResponse, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalValidatorResponse(data, consumer)
}

func unmarshalValidatorResponse(data []byte, consumer runtime.Consumer) (ValidatorResponse, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Model used for validation responses. property.
	var getType struct {
		ModelUsedForValidationResponses string `json:"Model used for validation responses."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Model used for validation responses.", "body", getType.ModelUsedForValidationResponses); err != nil {
		return nil, err
	}

	// The value of Model used for validation responses. is used to determine which type to create and unmarshal the data into
	switch getType.ModelUsedForValidationResponses {
	case "ValidatorResponse":
		var result validatorResponse
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Model used for validation responses. value: %q", getType.ModelUsedForValidationResponses)
}

// Validate validates this validator response
func (m *validatorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *validatorResponse) validateMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.Message()) { // not required
		return nil
	}

	if err := m.Message().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("message")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("message")
		}
		return err
	}

	return nil
}

// ContextValidate validate this validator response based on the context it is used
func (m *validatorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *validatorResponse) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Message().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("message")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("message")
		}
		return err
	}

	return nil
}
