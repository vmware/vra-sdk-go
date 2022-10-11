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

// RollbackResponse RollbackResponse
//
// Represents the response after rolling back a Stage or Task.
//
// swagger:discriminator RollbackResponse Represents the response after rolling back a Stage or Task.
type RollbackResponse interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Execution link of the rollback Pipeline.
	// Example: /codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337
	Link() string
	SetLink(string)

	// Execution index of the rollback Pipeline.
	// Example: 4
	Index() int64
	SetIndex(int64)

	// Name of the rollback Pipeline.
	// Example: My-Rolled-Back-Pipeline
	Name() string
	SetName(string)

	// Output properties of a rollback Pipeline.
	// Example: [{"key":"env","value":"dev"}]
	Output() interface{}
	SetOutput(interface{})

	// Execution status of a rollback Pipeline.
	// Example: COMPLETED
	// Enum: [NOT_STARTED STARTED RUNNING CANCELING WAITING RESUMING PAUSING PAUSED CANCELED COMPLETED FAILED SKIPPED QUEUED FAILED_CONTINUE ROLLING_BACK ROLLBACK_FAILED PREPARING_WORKSPACE ROLLBACK_COMPLETED]
	Status() string
	SetStatus(string)

	// Execution status message of a rollback Pipeline.
	// Example: Executing Stage0
	StatusMessage() string
	SetStatusMessage(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type rollbackResponse struct {
	linkField string

	indexField int64

	nameField string

	outputField interface{}

	statusField string

	statusMessageField string
}

// Link gets the link of this polymorphic type
func (m *rollbackResponse) Link() string {
	return m.linkField
}

// SetLink sets the link of this polymorphic type
func (m *rollbackResponse) SetLink(val string) {
	m.linkField = val
}

// Index gets the index of this polymorphic type
func (m *rollbackResponse) Index() int64 {
	return m.indexField
}

// SetIndex sets the index of this polymorphic type
func (m *rollbackResponse) SetIndex(val int64) {
	m.indexField = val
}

// Name gets the name of this polymorphic type
func (m *rollbackResponse) Name() string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *rollbackResponse) SetName(val string) {
	m.nameField = val
}

// Output gets the output of this polymorphic type
func (m *rollbackResponse) Output() interface{} {
	return m.outputField
}

// SetOutput sets the output of this polymorphic type
func (m *rollbackResponse) SetOutput(val interface{}) {
	m.outputField = val
}

// Status gets the status of this polymorphic type
func (m *rollbackResponse) Status() string {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *rollbackResponse) SetStatus(val string) {
	m.statusField = val
}

// StatusMessage gets the status message of this polymorphic type
func (m *rollbackResponse) StatusMessage() string {
	return m.statusMessageField
}

// SetStatusMessage sets the status message of this polymorphic type
func (m *rollbackResponse) SetStatusMessage(val string) {
	m.statusMessageField = val
}

// UnmarshalRollbackResponseSlice unmarshals polymorphic slices of RollbackResponse
func UnmarshalRollbackResponseSlice(reader io.Reader, consumer runtime.Consumer) ([]RollbackResponse, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []RollbackResponse
	for _, element := range elements {
		obj, err := unmarshalRollbackResponse(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalRollbackResponse unmarshals polymorphic RollbackResponse
func UnmarshalRollbackResponse(reader io.Reader, consumer runtime.Consumer) (RollbackResponse, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalRollbackResponse(data, consumer)
}

func unmarshalRollbackResponse(data []byte, consumer runtime.Consumer) (RollbackResponse, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Represents the response after rolling back a Stage or Task. property.
	var getType struct {
		RepresentsTheResponseAfterRollingBackaStageOrTask string `json:"Represents the response after rolling back a Stage or Task."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Represents the response after rolling back a Stage or Task.", "body", getType.RepresentsTheResponseAfterRollingBackaStageOrTask); err != nil {
		return nil, err
	}

	// The value of Represents the response after rolling back a Stage or Task. is used to determine which type to create and unmarshal the data into
	switch getType.RepresentsTheResponseAfterRollingBackaStageOrTask {
	case "RollbackResponse":
		var result rollbackResponse
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Represents the response after rolling back a Stage or Task. value: %q", getType.RepresentsTheResponseAfterRollingBackaStageOrTask)
}

// Validate validates this rollback response
func (m *rollbackResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rollbackResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_STARTED","STARTED","RUNNING","CANCELING","WAITING","RESUMING","PAUSING","PAUSED","CANCELED","COMPLETED","FAILED","SKIPPED","QUEUED","FAILED_CONTINUE","ROLLING_BACK","ROLLBACK_FAILED","PREPARING_WORKSPACE","ROLLBACK_COMPLETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rollbackResponseTypeStatusPropEnum = append(rollbackResponseTypeStatusPropEnum, v)
	}
}

const (

	// RollbackResponseStatusNOTSTARTED captures enum value "NOT_STARTED"
	RollbackResponseStatusNOTSTARTED string = "NOT_STARTED"

	// RollbackResponseStatusSTARTED captures enum value "STARTED"
	RollbackResponseStatusSTARTED string = "STARTED"

	// RollbackResponseStatusRUNNING captures enum value "RUNNING"
	RollbackResponseStatusRUNNING string = "RUNNING"

	// RollbackResponseStatusCANCELING captures enum value "CANCELING"
	RollbackResponseStatusCANCELING string = "CANCELING"

	// RollbackResponseStatusWAITING captures enum value "WAITING"
	RollbackResponseStatusWAITING string = "WAITING"

	// RollbackResponseStatusRESUMING captures enum value "RESUMING"
	RollbackResponseStatusRESUMING string = "RESUMING"

	// RollbackResponseStatusPAUSING captures enum value "PAUSING"
	RollbackResponseStatusPAUSING string = "PAUSING"

	// RollbackResponseStatusPAUSED captures enum value "PAUSED"
	RollbackResponseStatusPAUSED string = "PAUSED"

	// RollbackResponseStatusCANCELED captures enum value "CANCELED"
	RollbackResponseStatusCANCELED string = "CANCELED"

	// RollbackResponseStatusCOMPLETED captures enum value "COMPLETED"
	RollbackResponseStatusCOMPLETED string = "COMPLETED"

	// RollbackResponseStatusFAILED captures enum value "FAILED"
	RollbackResponseStatusFAILED string = "FAILED"

	// RollbackResponseStatusSKIPPED captures enum value "SKIPPED"
	RollbackResponseStatusSKIPPED string = "SKIPPED"

	// RollbackResponseStatusQUEUED captures enum value "QUEUED"
	RollbackResponseStatusQUEUED string = "QUEUED"

	// RollbackResponseStatusFAILEDCONTINUE captures enum value "FAILED_CONTINUE"
	RollbackResponseStatusFAILEDCONTINUE string = "FAILED_CONTINUE"

	// RollbackResponseStatusROLLINGBACK captures enum value "ROLLING_BACK"
	RollbackResponseStatusROLLINGBACK string = "ROLLING_BACK"

	// RollbackResponseStatusROLLBACKFAILED captures enum value "ROLLBACK_FAILED"
	RollbackResponseStatusROLLBACKFAILED string = "ROLLBACK_FAILED"

	// RollbackResponseStatusPREPARINGWORKSPACE captures enum value "PREPARING_WORKSPACE"
	RollbackResponseStatusPREPARINGWORKSPACE string = "PREPARING_WORKSPACE"

	// RollbackResponseStatusROLLBACKCOMPLETED captures enum value "ROLLBACK_COMPLETED"
	RollbackResponseStatusROLLBACKCOMPLETED string = "ROLLBACK_COMPLETED"
)

// prop value enum
func (m *rollbackResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rollbackResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *rollbackResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status()) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rollback response based on context it is used
func (m *rollbackResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
