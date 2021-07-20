// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BatchUserOperationResponse BatchUserOperationResponse
//
// Batch user operation response.
//
// swagger:discriminator BatchUserOperationResponse Batch user operation response.
type BatchUserOperationResponse interface {
	runtime.Validatable
	runtime.ContextValidatable

	Errors() []BatchUserOperationStatus
	SetErrors([]BatchUserOperationStatus)

	// The number of approval requests failed.
	// Example: 5
	Failure() int32
	SetFailure(int32)

	// The number of approval requests succeeded.
	// Example: 10
	Success() int32
	SetSuccess(int32)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type batchUserOperationResponse struct {
	errorsField []BatchUserOperationStatus

	failureField int32

	successField int32
}

// Errors gets the errors of this polymorphic type
func (m *batchUserOperationResponse) Errors() []BatchUserOperationStatus {
	return m.errorsField
}

// SetErrors sets the errors of this polymorphic type
func (m *batchUserOperationResponse) SetErrors(val []BatchUserOperationStatus) {
	m.errorsField = val
}

// Failure gets the failure of this polymorphic type
func (m *batchUserOperationResponse) Failure() int32 {
	return m.failureField
}

// SetFailure sets the failure of this polymorphic type
func (m *batchUserOperationResponse) SetFailure(val int32) {
	m.failureField = val
}

// Success gets the success of this polymorphic type
func (m *batchUserOperationResponse) Success() int32 {
	return m.successField
}

// SetSuccess sets the success of this polymorphic type
func (m *batchUserOperationResponse) SetSuccess(val int32) {
	m.successField = val
}

// UnmarshalBatchUserOperationResponseSlice unmarshals polymorphic slices of BatchUserOperationResponse
func UnmarshalBatchUserOperationResponseSlice(reader io.Reader, consumer runtime.Consumer) ([]BatchUserOperationResponse, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []BatchUserOperationResponse
	for _, element := range elements {
		obj, err := unmarshalBatchUserOperationResponse(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalBatchUserOperationResponse unmarshals polymorphic BatchUserOperationResponse
func UnmarshalBatchUserOperationResponse(reader io.Reader, consumer runtime.Consumer) (BatchUserOperationResponse, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalBatchUserOperationResponse(data, consumer)
}

func unmarshalBatchUserOperationResponse(data []byte, consumer runtime.Consumer) (BatchUserOperationResponse, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Batch user operation response. property.
	var getType struct {
		BatchUserOperationResponse string `json:"Batch user operation response."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Batch user operation response.", "body", getType.BatchUserOperationResponse); err != nil {
		return nil, err
	}

	// The value of Batch user operation response. is used to determine which type to create and unmarshal the data into
	switch getType.BatchUserOperationResponse {
	case "BatchUserOperationResponse":
		var result batchUserOperationResponse
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Batch user operation response. value: %q", getType.BatchUserOperationResponse)
}

// Validate validates this batch user operation response
func (m *batchUserOperationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *batchUserOperationResponse) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors()) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors()); i++ {

		if err := m.errorsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errors" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this batch user operation response based on the context it is used
func (m *batchUserOperationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *batchUserOperationResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors()); i++ {

		if err := m.errorsField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errors" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}
