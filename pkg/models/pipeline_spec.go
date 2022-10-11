// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PipelineSpec PipelineSpec
//
// # Pipeline Specification
//
// swagger:discriminator PipelineSpec Pipeline Specification
type PipelineSpec interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Additional information about Input Properties
	// Example: {"test":{"description":"test-value","mandatory":false}}
	InputMeta() map[string]PropertyMetaData
	SetInputMeta(map[string]PropertyMetaData)

	// Number of Executions of the Pipeline that can run concurrently.
	// Example: 10
	Concurrency() int32
	SetConcurrency(int32)

	// A human-friendly description.
	// Required: true
	Description() *string
	SetDescription(*string)

	// Indicates if the Pipeline is in enabled state.
	// Example: true
	Enabled() bool
	SetEnabled(bool)

	// String description of the icon used for this Pipeline.
	// Example: tools,,is-success is-solid
	Icon() string
	SetIcon(string)

	// Map representing the Input properties for the Pipeline.
	// Example: [{"ip":"10.5.23.84","script":"testScript.sh"}]
	Input() interface{}
	SetInput(interface{})

	// A human-friendly name used as an identifier in APIs that support this option
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	Notifications() NotificationConfiguration
	SetNotifications(NotificationConfiguration)

	// Represents the different options to trigger a Pipeline. Selecting an option auto injects the Input properties needed to execute a Pipeline with that trigger.
	// Example: ["DOCKER_TRIGGER"]
	Options() []string
	SetOptions([]string)

	// Map representing the Output properties for the Pipeline.
	// Example: [{"deployedMachineIP":"10.108.35.54","result":"true"}]
	Output() interface{}
	SetOutput(interface{})

	// The project this entity belongs to.
	// Example: My-Project
	Project() string
	SetProject(string)

	Rollbacks() []RollbackConfiguration
	SetRollbacks([]RollbackConfiguration)

	// Represents the order in which Stages will be executed.
	// Example: ["AcquireToken","Upgrade","E2E"]
	StageOrder() []string
	SetStageOrder([]string)

	// Map representing the details of the various Stages of the Pipeline.
	Stages() map[string]Stage
	SetStages(map[string]Stage)

	Starred() PipelineStarredProperty
	SetStarred(PipelineStarredProperty)

	// Indicates if the Pipeline is enabled/disabled/released to catalog.
	// Example: RELEASED
	State() string
	SetState(string)

	// A set of tag keys and optional values that were set on on the resource.
	// Example: [{"key":"env","value":"dev"}]
	Tags() []string
	SetTags([]string)

	// Represents the configuration to be used for CI and Custom tasks.
	Workspace() *Workspace
	SetWorkspace(*Workspace)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type pipelineSpec struct {
	inputMetaField map[string]PropertyMetaData

	concurrencyField int32

	descriptionField *string

	enabledField bool

	iconField string

	inputField interface{}

	nameField *string

	notificationsField NotificationConfiguration

	optionsField []string

	outputField interface{}

	projectField string

	rollbacksField []RollbackConfiguration

	stageOrderField []string

	stagesField map[string]Stage

	starredField PipelineStarredProperty

	stateField string

	tagsField []string

	workspaceField *Workspace
}

// InputMeta gets the input meta of this polymorphic type
func (m *pipelineSpec) InputMeta() map[string]PropertyMetaData {
	return m.inputMetaField
}

// SetInputMeta sets the input meta of this polymorphic type
func (m *pipelineSpec) SetInputMeta(val map[string]PropertyMetaData) {
	m.inputMetaField = val
}

// Concurrency gets the concurrency of this polymorphic type
func (m *pipelineSpec) Concurrency() int32 {
	return m.concurrencyField
}

// SetConcurrency sets the concurrency of this polymorphic type
func (m *pipelineSpec) SetConcurrency(val int32) {
	m.concurrencyField = val
}

// Description gets the description of this polymorphic type
func (m *pipelineSpec) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *pipelineSpec) SetDescription(val *string) {
	m.descriptionField = val
}

// Enabled gets the enabled of this polymorphic type
func (m *pipelineSpec) Enabled() bool {
	return m.enabledField
}

// SetEnabled sets the enabled of this polymorphic type
func (m *pipelineSpec) SetEnabled(val bool) {
	m.enabledField = val
}

// Icon gets the icon of this polymorphic type
func (m *pipelineSpec) Icon() string {
	return m.iconField
}

// SetIcon sets the icon of this polymorphic type
func (m *pipelineSpec) SetIcon(val string) {
	m.iconField = val
}

// Input gets the input of this polymorphic type
func (m *pipelineSpec) Input() interface{} {
	return m.inputField
}

// SetInput sets the input of this polymorphic type
func (m *pipelineSpec) SetInput(val interface{}) {
	m.inputField = val
}

// Name gets the name of this polymorphic type
func (m *pipelineSpec) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *pipelineSpec) SetName(val *string) {
	m.nameField = val
}

// Notifications gets the notifications of this polymorphic type
func (m *pipelineSpec) Notifications() NotificationConfiguration {
	return m.notificationsField
}

// SetNotifications sets the notifications of this polymorphic type
func (m *pipelineSpec) SetNotifications(val NotificationConfiguration) {
	m.notificationsField = val
}

// Options gets the options of this polymorphic type
func (m *pipelineSpec) Options() []string {
	return m.optionsField
}

// SetOptions sets the options of this polymorphic type
func (m *pipelineSpec) SetOptions(val []string) {
	m.optionsField = val
}

// Output gets the output of this polymorphic type
func (m *pipelineSpec) Output() interface{} {
	return m.outputField
}

// SetOutput sets the output of this polymorphic type
func (m *pipelineSpec) SetOutput(val interface{}) {
	m.outputField = val
}

// Project gets the project of this polymorphic type
func (m *pipelineSpec) Project() string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *pipelineSpec) SetProject(val string) {
	m.projectField = val
}

// Rollbacks gets the rollbacks of this polymorphic type
func (m *pipelineSpec) Rollbacks() []RollbackConfiguration {
	return m.rollbacksField
}

// SetRollbacks sets the rollbacks of this polymorphic type
func (m *pipelineSpec) SetRollbacks(val []RollbackConfiguration) {
	m.rollbacksField = val
}

// StageOrder gets the stage order of this polymorphic type
func (m *pipelineSpec) StageOrder() []string {
	return m.stageOrderField
}

// SetStageOrder sets the stage order of this polymorphic type
func (m *pipelineSpec) SetStageOrder(val []string) {
	m.stageOrderField = val
}

// Stages gets the stages of this polymorphic type
func (m *pipelineSpec) Stages() map[string]Stage {
	return m.stagesField
}

// SetStages sets the stages of this polymorphic type
func (m *pipelineSpec) SetStages(val map[string]Stage) {
	m.stagesField = val
}

// Starred gets the starred of this polymorphic type
func (m *pipelineSpec) Starred() PipelineStarredProperty {
	return m.starredField
}

// SetStarred sets the starred of this polymorphic type
func (m *pipelineSpec) SetStarred(val PipelineStarredProperty) {
	m.starredField = val
}

// State gets the state of this polymorphic type
func (m *pipelineSpec) State() string {
	return m.stateField
}

// SetState sets the state of this polymorphic type
func (m *pipelineSpec) SetState(val string) {
	m.stateField = val
}

// Tags gets the tags of this polymorphic type
func (m *pipelineSpec) Tags() []string {
	return m.tagsField
}

// SetTags sets the tags of this polymorphic type
func (m *pipelineSpec) SetTags(val []string) {
	m.tagsField = val
}

// Workspace gets the workspace of this polymorphic type
func (m *pipelineSpec) Workspace() *Workspace {
	return m.workspaceField
}

// SetWorkspace sets the workspace of this polymorphic type
func (m *pipelineSpec) SetWorkspace(val *Workspace) {
	m.workspaceField = val
}

// UnmarshalPipelineSpecSlice unmarshals polymorphic slices of PipelineSpec
func UnmarshalPipelineSpecSlice(reader io.Reader, consumer runtime.Consumer) ([]PipelineSpec, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []PipelineSpec
	for _, element := range elements {
		obj, err := unmarshalPipelineSpec(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPipelineSpec unmarshals polymorphic PipelineSpec
func UnmarshalPipelineSpec(reader io.Reader, consumer runtime.Consumer) (PipelineSpec, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPipelineSpec(data, consumer)
}

func unmarshalPipelineSpec(data []byte, consumer runtime.Consumer) (PipelineSpec, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Pipeline Specification property.
	var getType struct {
		PipelineSpecification string `json:"Pipeline Specification"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Pipeline Specification", "body", getType.PipelineSpecification); err != nil {
		return nil, err
	}

	// The value of Pipeline Specification is used to determine which type to create and unmarshal the data into
	switch getType.PipelineSpecification {
	case "PipelineSpec":
		var result pipelineSpec
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Pipeline Specification value: %q", getType.PipelineSpecification)
}

// Validate validates this pipeline spec
func (m *pipelineSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRollbacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarred(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *pipelineSpec) validateInputMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.InputMeta()) { // not required
		return nil
	}

	for k := range m.InputMeta() {

		if err := validate.Required("_inputMeta"+"."+k, "body", m.InputMeta()[k]); err != nil {
			return err
		}
		if val, ok := m.InputMeta()[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_inputMeta" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_inputMeta" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *pipelineSpec) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *pipelineSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *pipelineSpec) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications()) { // not required
		return nil
	}

	if err := m.Notifications().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notifications")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("notifications")
		}
		return err
	}

	return nil
}

func (m *pipelineSpec) validateRollbacks(formats strfmt.Registry) error {
	if swag.IsZero(m.Rollbacks()) { // not required
		return nil
	}

	for i := 0; i < len(m.Rollbacks()); i++ {

		if err := m.rollbacksField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollbacks" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollbacks" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *pipelineSpec) validateStages(formats strfmt.Registry) error {
	if swag.IsZero(m.Stages()) { // not required
		return nil
	}

	for k := range m.Stages() {

		if err := validate.Required("stages"+"."+k, "body", m.Stages()[k]); err != nil {
			return err
		}
		if val, ok := m.Stages()[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stages" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("stages" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *pipelineSpec) validateStarred(formats strfmt.Registry) error {
	if swag.IsZero(m.Starred()) { // not required
		return nil
	}

	if err := m.Starred().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("starred")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("starred")
		}
		return err
	}

	return nil
}

func (m *pipelineSpec) validateWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspace()) { // not required
		return nil
	}

	if m.Workspace() != nil {
		if err := m.Workspace().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pipeline spec based on the context it is used
func (m *pipelineSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRollbacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStarred(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *pipelineSpec) contextValidateInputMeta(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.InputMeta() {

		if val, ok := m.InputMeta()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *pipelineSpec) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Notifications().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("notifications")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("notifications")
		}
		return err
	}

	return nil
}

func (m *pipelineSpec) contextValidateRollbacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rollbacks()); i++ {

		if err := m.rollbacksField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollbacks" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rollbacks" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *pipelineSpec) contextValidateStages(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Stages() {

		if val, ok := m.Stages()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *pipelineSpec) contextValidateStarred(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Starred().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("starred")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("starred")
		}
		return err
	}

	return nil
}

func (m *pipelineSpec) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspace() != nil {
		if err := m.Workspace().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}
