// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// UserOpResponse UserOpResponse
//
// User operation response.
// swagger:discriminator UserOpResponse User operation response.
type UserOpResponse interface {
	runtime.Validatable

	// The response message which the responder would like to give.
	ResponseMessage() string
	SetResponseMessage(string)

	// The status of this entity.
	Status() string
	SetStatus(string)
}

type userOpResponse struct {
	responseMessageField string

	statusField string
}

// ResponseMessage gets the response message of this polymorphic type
func (m *userOpResponse) ResponseMessage() string {
	return m.responseMessageField
}

// SetResponseMessage sets the response message of this polymorphic type
func (m *userOpResponse) SetResponseMessage(val string) {
	m.responseMessageField = val
}

// Status gets the status of this polymorphic type
func (m *userOpResponse) Status() string {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *userOpResponse) SetStatus(val string) {
	m.statusField = val
}

// UnmarshalUserOpResponseSlice unmarshals polymorphic slices of UserOpResponse
func UnmarshalUserOpResponseSlice(reader io.Reader, consumer runtime.Consumer) ([]UserOpResponse, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []UserOpResponse
	for _, element := range elements {
		obj, err := unmarshalUserOpResponse(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalUserOpResponse unmarshals polymorphic UserOpResponse
func UnmarshalUserOpResponse(reader io.Reader, consumer runtime.Consumer) (UserOpResponse, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalUserOpResponse(data, consumer)
}

func unmarshalUserOpResponse(data []byte, consumer runtime.Consumer) (UserOpResponse, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the User operation response. property.
	var getType struct {
		UserOperationResponse string `json:"User operation response."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("User operation response.", "body", getType.UserOperationResponse); err != nil {
		return nil, err
	}

	// The value of User operation response. is used to determine which type to create and unmarshal the data into
	switch getType.UserOperationResponse {
	case "UserOpResponse":
		var result userOpResponse
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid User operation response. value: %q", getType.UserOperationResponse)

}

// Validate validates this user op response
func (m *userOpResponse) Validate(formats strfmt.Registry) error {
	return nil
}
