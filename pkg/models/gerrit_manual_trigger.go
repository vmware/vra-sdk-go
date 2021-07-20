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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GerritManualTrigger GerritManualTrigger
//
// Gerrit Manual Trigger feature is used to simulate an event corresponding to a change-id and match it against a Gerrit Trigger configuration to trigger pipeline(s) manually.
//
// swagger:discriminator GerritManualTrigger Gerrit Manual Trigger feature is used to simulate an event corresponding to a change-id and match it against a Gerrit Trigger configuration to trigger pipeline(s) manually.
type GerritManualTrigger interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The ChangeSet Id corresponding to which the manual trigger has to be performed.
	// Example: I522eab3b678dedd68dad3e2c04d3cccc0eb9fa00
	// Required: true
	ChangeID() *string
	SetChangeID(*string)

	// The tango project to which Gerrit Trigger belongs to.
	// Example: Gerrit-Trigger-Project
	// Required: true
	Project() *string
	SetProject(*string)

	// Name of the Gerrit Trigger for which event has to be analysed.
	// Example: Gerrit-Trigger
	// Required: true
	Trigger() *string
	SetTrigger(*string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type gerritManualTrigger struct {
	changeIdField *string

	projectField *string

	triggerField *string
}

// ChangeID gets the change Id of this polymorphic type
func (m *gerritManualTrigger) ChangeID() *string {
	return m.changeIdField
}

// SetChangeID sets the change Id of this polymorphic type
func (m *gerritManualTrigger) SetChangeID(val *string) {
	m.changeIdField = val
}

// Project gets the project of this polymorphic type
func (m *gerritManualTrigger) Project() *string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *gerritManualTrigger) SetProject(val *string) {
	m.projectField = val
}

// Trigger gets the trigger of this polymorphic type
func (m *gerritManualTrigger) Trigger() *string {
	return m.triggerField
}

// SetTrigger sets the trigger of this polymorphic type
func (m *gerritManualTrigger) SetTrigger(val *string) {
	m.triggerField = val
}

// UnmarshalGerritManualTriggerSlice unmarshals polymorphic slices of GerritManualTrigger
func UnmarshalGerritManualTriggerSlice(reader io.Reader, consumer runtime.Consumer) ([]GerritManualTrigger, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []GerritManualTrigger
	for _, element := range elements {
		obj, err := unmarshalGerritManualTrigger(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalGerritManualTrigger unmarshals polymorphic GerritManualTrigger
func UnmarshalGerritManualTrigger(reader io.Reader, consumer runtime.Consumer) (GerritManualTrigger, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalGerritManualTrigger(data, consumer)
}

func unmarshalGerritManualTrigger(data []byte, consumer runtime.Consumer) (GerritManualTrigger, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Gerrit Manual Trigger feature is used to simulate an event corresponding to a change-id and match it against a Gerrit Trigger configuration to trigger pipeline(s) manually. property.
	var getType struct {
		GerritManualTriggerFeatureIsUsedToSimulateAnEventCorrespondingToaChangeIDAndMatchItAgainstaGerritTriggerConfigurationToTriggerPipelinesManually string `json:"Gerrit Manual Trigger feature is used to simulate an event corresponding to a change-id and match it against a Gerrit Trigger configuration to trigger pipeline(s) manually."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Gerrit Manual Trigger feature is used to simulate an event corresponding to a change-id and match it against a Gerrit Trigger configuration to trigger pipeline(s) manually.", "body", getType.GerritManualTriggerFeatureIsUsedToSimulateAnEventCorrespondingToaChangeIDAndMatchItAgainstaGerritTriggerConfigurationToTriggerPipelinesManually); err != nil {
		return nil, err
	}

	// The value of Gerrit Manual Trigger feature is used to simulate an event corresponding to a change-id and match it against a Gerrit Trigger configuration to trigger pipeline(s) manually. is used to determine which type to create and unmarshal the data into
	switch getType.GerritManualTriggerFeatureIsUsedToSimulateAnEventCorrespondingToaChangeIDAndMatchItAgainstaGerritTriggerConfigurationToTriggerPipelinesManually {
	case "GerritManualTrigger":
		var result gerritManualTrigger
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Gerrit Manual Trigger feature is used to simulate an event corresponding to a change-id and match it against a Gerrit Trigger configuration to trigger pipeline(s) manually. value: %q", getType.GerritManualTriggerFeatureIsUsedToSimulateAnEventCorrespondingToaChangeIDAndMatchItAgainstaGerritTriggerConfigurationToTriggerPipelinesManually)
}

// Validate validates this gerrit manual trigger
func (m *gerritManualTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *gerritManualTrigger) validateChangeID(formats strfmt.Registry) error {

	if err := validate.Required("changeId", "body", m.ChangeID()); err != nil {
		return err
	}

	return nil
}

func (m *gerritManualTrigger) validateProject(formats strfmt.Registry) error {

	if err := validate.Required("project", "body", m.Project()); err != nil {
		return err
	}

	return nil
}

func (m *gerritManualTrigger) validateTrigger(formats strfmt.Registry) error {

	if err := validate.Required("trigger", "body", m.Trigger()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gerrit manual trigger based on context it is used
func (m *gerritManualTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
