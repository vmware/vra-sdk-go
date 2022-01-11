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

// GerritTrigger GerritTrigger
//
// The configurations of a Gerrit Trigger decides which pipeline is required to be triggered depending on kind of event received.
//
// swagger:discriminator GerritTrigger The configurations of a Gerrit Trigger decides which pipeline is required to be triggered depending on kind of event received.
type GerritTrigger interface {
	runtime.Validatable
	runtime.ContextValidatable

	// This field is provided for backward compatibility. Contains the same value as the 'createdAt' field as a UNIX timestamp in microseconds
	// Example: 1568625938000000
	CreateTimeInMicros() int64
	SetCreateTimeInMicros(int64)

	// Partial URL that provides details of the resource.
	// Example: /codestream/api/\u003cprefix\u003e/8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	Link() string
	SetLink(string)

	// Contains project id of the entity
	// Example: abcd-abcd-abcd
	ProjectID() string
	SetProjectID(string)

	// This field is provided for backward compatibility. Contains the same value as the 'updatedAt' field as a UNIX timestamp in microseconds
	// Example: 1568625938000000
	UpdateTimeInMicros() int64
	SetUpdateTimeInMicros(int64)

	// Gerrit Project Branch on which the change has to be monitored.
	// Example: master
	Branch() string
	SetBranch(string)

	// The configuration of an Gerrit Event.
	// Example: [{"eventType":"change-merged","failureComment":"Pipeline Execution Failed","input":{},"pipeline":"rest_pipeline","successComment":"Pipeline Execution Successful"},{"eventType":"patchset-created","failureComment":"Pipeline Failed","input":{},"pipeline":"rest_pipeline","successComment":"Pipeline Completed","verifiedLabel":"Verified"}]
	Configurations() []*GerritTriggerGerritEventConfiguration
	SetConfigurations([]*GerritTriggerGerritEventConfiguration)

	// Date when the entity was created. The date is in ISO 8601 with time zone
	// Example: 2019-09-16 09:25:38.065065+00
	CreatedAt() string
	SetCreatedAt(string)

	// The user that created this entity
	// Example: exampleuser
	CreatedBy() string
	SetCreatedBy(string)

	// A human-friendly description.
	// Required: true
	Description() *string
	SetDescription(*string)

	// Indicates that this trigger will be addressed on receiving respective events.
	// Example: true
	Enabled() bool
	SetEnabled(bool)

	// Provide file exclusions as conditions for the trigger.
	// Example: [{"type":"PLAIN","value":"example.txt"}]
	Exclusions() []*GerritTriggerFileFilter
	SetExclusions([]*GerritTriggerFileFilter)

	// Gerrit Project which is to be monitored for the change.
	// Example: TestGerritProject
	GerritProject() string
	SetGerritProject(string)

	// The id of this resource.
	// Example: 8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	ID() string
	SetID(string)

	// Provide file inclusions as conditions for the trigger.
	// Example: [{"type":"PLAIN","value":"example.txt"}]
	Inclusions() []*GerritTriggerFileFilter
	SetInclusions([]*GerritTriggerFileFilter)

	// Gerrit Listener which will receive the events for this trigger.
	// Example: Gerrit-Listener
	Listener() string
	SetListener(string)

	// A human-friendly name used as an identifier in APIs that support this option
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	// Prioritize Exclusion ensures that Pipelines are not triggered even if any of the files in a commit match the specified files in the exclusion paths or regex.
	// Example: true
	PrioritizeExclusion() bool
	SetPrioritizeExclusion(bool)

	// The project this entity belongs to.
	// Example: My-Project
	Project() string
	SetProject(string)

	// Date when the entity was last updated. The date is in ISO 8601 with time zone.
	// Example: 2019-09-16 09:25:38.065065+00
	UpdatedAt() string
	SetUpdatedAt(string)

	// The user that last updated this entity
	// Example: exampleuser
	UpdatedBy() string
	SetUpdatedBy(string)

	// Version of the resource.
	// Example: v1
	Version() string
	SetVersion(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type gerritTrigger struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	branchField string

	configurationsField []*GerritTriggerGerritEventConfiguration

	createdAtField string

	createdByField string

	descriptionField *string

	enabledField bool

	exclusionsField []*GerritTriggerFileFilter

	gerritProjectField string

	idField string

	inclusionsField []*GerritTriggerFileFilter

	listenerField string

	nameField *string

	prioritizeExclusionField bool

	projectField string

	updatedAtField string

	updatedByField string

	versionField string
}

// CreateTimeInMicros gets the create time in micros of this polymorphic type
func (m *gerritTrigger) CreateTimeInMicros() int64 {
	return m.createTimeInMicrosField
}

// SetCreateTimeInMicros sets the create time in micros of this polymorphic type
func (m *gerritTrigger) SetCreateTimeInMicros(val int64) {
	m.createTimeInMicrosField = val
}

// Link gets the link of this polymorphic type
func (m *gerritTrigger) Link() string {
	return m.linkField
}

// SetLink sets the link of this polymorphic type
func (m *gerritTrigger) SetLink(val string) {
	m.linkField = val
}

// ProjectID gets the project Id of this polymorphic type
func (m *gerritTrigger) ProjectID() string {
	return m.projectIdField
}

// SetProjectID sets the project Id of this polymorphic type
func (m *gerritTrigger) SetProjectID(val string) {
	m.projectIdField = val
}

// UpdateTimeInMicros gets the update time in micros of this polymorphic type
func (m *gerritTrigger) UpdateTimeInMicros() int64 {
	return m.updateTimeInMicrosField
}

// SetUpdateTimeInMicros sets the update time in micros of this polymorphic type
func (m *gerritTrigger) SetUpdateTimeInMicros(val int64) {
	m.updateTimeInMicrosField = val
}

// Branch gets the branch of this polymorphic type
func (m *gerritTrigger) Branch() string {
	return m.branchField
}

// SetBranch sets the branch of this polymorphic type
func (m *gerritTrigger) SetBranch(val string) {
	m.branchField = val
}

// Configurations gets the configurations of this polymorphic type
func (m *gerritTrigger) Configurations() []*GerritTriggerGerritEventConfiguration {
	return m.configurationsField
}

// SetConfigurations sets the configurations of this polymorphic type
func (m *gerritTrigger) SetConfigurations(val []*GerritTriggerGerritEventConfiguration) {
	m.configurationsField = val
}

// CreatedAt gets the created at of this polymorphic type
func (m *gerritTrigger) CreatedAt() string {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this polymorphic type
func (m *gerritTrigger) SetCreatedAt(val string) {
	m.createdAtField = val
}

// CreatedBy gets the created by of this polymorphic type
func (m *gerritTrigger) CreatedBy() string {
	return m.createdByField
}

// SetCreatedBy sets the created by of this polymorphic type
func (m *gerritTrigger) SetCreatedBy(val string) {
	m.createdByField = val
}

// Description gets the description of this polymorphic type
func (m *gerritTrigger) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *gerritTrigger) SetDescription(val *string) {
	m.descriptionField = val
}

// Enabled gets the enabled of this polymorphic type
func (m *gerritTrigger) Enabled() bool {
	return m.enabledField
}

// SetEnabled sets the enabled of this polymorphic type
func (m *gerritTrigger) SetEnabled(val bool) {
	m.enabledField = val
}

// Exclusions gets the exclusions of this polymorphic type
func (m *gerritTrigger) Exclusions() []*GerritTriggerFileFilter {
	return m.exclusionsField
}

// SetExclusions sets the exclusions of this polymorphic type
func (m *gerritTrigger) SetExclusions(val []*GerritTriggerFileFilter) {
	m.exclusionsField = val
}

// GerritProject gets the gerrit project of this polymorphic type
func (m *gerritTrigger) GerritProject() string {
	return m.gerritProjectField
}

// SetGerritProject sets the gerrit project of this polymorphic type
func (m *gerritTrigger) SetGerritProject(val string) {
	m.gerritProjectField = val
}

// ID gets the id of this polymorphic type
func (m *gerritTrigger) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *gerritTrigger) SetID(val string) {
	m.idField = val
}

// Inclusions gets the inclusions of this polymorphic type
func (m *gerritTrigger) Inclusions() []*GerritTriggerFileFilter {
	return m.inclusionsField
}

// SetInclusions sets the inclusions of this polymorphic type
func (m *gerritTrigger) SetInclusions(val []*GerritTriggerFileFilter) {
	m.inclusionsField = val
}

// Listener gets the listener of this polymorphic type
func (m *gerritTrigger) Listener() string {
	return m.listenerField
}

// SetListener sets the listener of this polymorphic type
func (m *gerritTrigger) SetListener(val string) {
	m.listenerField = val
}

// Name gets the name of this polymorphic type
func (m *gerritTrigger) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *gerritTrigger) SetName(val *string) {
	m.nameField = val
}

// PrioritizeExclusion gets the prioritize exclusion of this polymorphic type
func (m *gerritTrigger) PrioritizeExclusion() bool {
	return m.prioritizeExclusionField
}

// SetPrioritizeExclusion sets the prioritize exclusion of this polymorphic type
func (m *gerritTrigger) SetPrioritizeExclusion(val bool) {
	m.prioritizeExclusionField = val
}

// Project gets the project of this polymorphic type
func (m *gerritTrigger) Project() string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *gerritTrigger) SetProject(val string) {
	m.projectField = val
}

// UpdatedAt gets the updated at of this polymorphic type
func (m *gerritTrigger) UpdatedAt() string {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this polymorphic type
func (m *gerritTrigger) SetUpdatedAt(val string) {
	m.updatedAtField = val
}

// UpdatedBy gets the updated by of this polymorphic type
func (m *gerritTrigger) UpdatedBy() string {
	return m.updatedByField
}

// SetUpdatedBy sets the updated by of this polymorphic type
func (m *gerritTrigger) SetUpdatedBy(val string) {
	m.updatedByField = val
}

// Version gets the version of this polymorphic type
func (m *gerritTrigger) Version() string {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *gerritTrigger) SetVersion(val string) {
	m.versionField = val
}

// UnmarshalGerritTriggerSlice unmarshals polymorphic slices of GerritTrigger
func UnmarshalGerritTriggerSlice(reader io.Reader, consumer runtime.Consumer) ([]GerritTrigger, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []GerritTrigger
	for _, element := range elements {
		obj, err := unmarshalGerritTrigger(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalGerritTrigger unmarshals polymorphic GerritTrigger
func UnmarshalGerritTrigger(reader io.Reader, consumer runtime.Consumer) (GerritTrigger, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalGerritTrigger(data, consumer)
}

func unmarshalGerritTrigger(data []byte, consumer runtime.Consumer) (GerritTrigger, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the The configurations of a Gerrit Trigger decides which pipeline is required to be triggered depending on kind of event received. property.
	var getType struct {
		TheConfigurationsOfaGerritTriggerDecidesWhichPipelineIsRequiredToBeTriggeredDependingOnKindOfEventReceived string `json:"The configurations of a Gerrit Trigger decides which pipeline is required to be triggered depending on kind of event received."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("The configurations of a Gerrit Trigger decides which pipeline is required to be triggered depending on kind of event received.", "body", getType.TheConfigurationsOfaGerritTriggerDecidesWhichPipelineIsRequiredToBeTriggeredDependingOnKindOfEventReceived); err != nil {
		return nil, err
	}

	// The value of The configurations of a Gerrit Trigger decides which pipeline is required to be triggered depending on kind of event received. is used to determine which type to create and unmarshal the data into
	switch getType.TheConfigurationsOfaGerritTriggerDecidesWhichPipelineIsRequiredToBeTriggeredDependingOnKindOfEventReceived {
	case "GerritTrigger":
		var result gerritTrigger
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid The configurations of a Gerrit Trigger decides which pipeline is required to be triggered depending on kind of event received. value: %q", getType.TheConfigurationsOfaGerritTriggerDecidesWhichPipelineIsRequiredToBeTriggeredDependingOnKindOfEventReceived)
}

// Validate validates this gerrit trigger
func (m *gerritTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExclusions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInclusions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *gerritTrigger) validateConfigurations(formats strfmt.Registry) error {
	if swag.IsZero(m.Configurations()) { // not required
		return nil
	}

	for i := 0; i < len(m.Configurations()); i++ {
		if swag.IsZero(m.configurationsField[i]) { // not required
			continue
		}

		if m.configurationsField[i] != nil {
			if err := m.configurationsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *gerritTrigger) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *gerritTrigger) validateExclusions(formats strfmt.Registry) error {
	if swag.IsZero(m.Exclusions()) { // not required
		return nil
	}

	for i := 0; i < len(m.Exclusions()); i++ {
		if swag.IsZero(m.exclusionsField[i]) { // not required
			continue
		}

		if m.exclusionsField[i] != nil {
			if err := m.exclusionsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exclusions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exclusions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *gerritTrigger) validateInclusions(formats strfmt.Registry) error {
	if swag.IsZero(m.Inclusions()) { // not required
		return nil
	}

	for i := 0; i < len(m.Inclusions()); i++ {
		if swag.IsZero(m.inclusionsField[i]) { // not required
			continue
		}

		if m.inclusionsField[i] != nil {
			if err := m.inclusionsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inclusions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inclusions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *gerritTrigger) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this gerrit trigger based on the context it is used
func (m *gerritTrigger) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigurations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExclusions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInclusions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *gerritTrigger) contextValidateConfigurations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Configurations()); i++ {

		if m.configurationsField[i] != nil {
			if err := m.configurationsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *gerritTrigger) contextValidateExclusions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Exclusions()); i++ {

		if m.exclusionsField[i] != nil {
			if err := m.exclusionsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exclusions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("exclusions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *gerritTrigger) contextValidateInclusions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Inclusions()); i++ {

		if m.inclusionsField[i] != nil {
			if err := m.inclusionsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inclusions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inclusions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
