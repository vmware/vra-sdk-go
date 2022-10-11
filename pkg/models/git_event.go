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

// GitEvent GitEvent
//
// Git Event Model.
//
// swagger:discriminator GitEvent Git Event Model.
type GitEvent interface {
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

	// CommitId of the event triggered.
	// Example: dvdsasfvaesvvffvgfbrrgeaerfare34fdgrt43
	CommitID() string
	SetCommitID(string)

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

	// Pipeline execution index.
	// Example: 2
	ExecutionIndex() int64
	SetExecutionIndex(int64)

	// Pipeline execution link.
	// Example: /codestream/api/executions/08c20cb8-17b6-4d9c-9f15-c77cf4913bc5
	ExecutionLink() string
	SetExecutionLink(string)

	// Pipeline execution status.
	// Example: STARTED
	ExecutionStatus() string
	SetExecutionStatus(string)

	// Url of the Git Repository.
	// Example: https://\u003cgit-server-url\u003e/\u003cusername\u003e/\u003crepo-name\u003e
	HTTPURL() string
	SetHTTPURL(string)

	// The id of this resource.
	// Example: 8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	ID() string
	SetID(string)

	// Message is populated when an error occurs.
	// Example: Invalid Oauth Refresh token. Verify webhook configuration.
	Message() string
	SetMessage(string)

	// A human-friendly name used as an identifier in APIs that support this option
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	// Owner Name who triggered the event.
	// Example: Demo User
	Owner() string
	SetOwner(string)

	// Pipeline to be executed when event is triggered.
	// Example: Demo Pipeline
	Pipeline() string
	SetPipeline(string)

	// The project this entity belongs to.
	// Example: My-Project
	Project() string
	SetProject(string)

	// RepoName corresponding to the event.
	// Example: MyRepository
	Repo() string
	SetRepo(string)

	// Url of the Git server.
	// Example: https://\u003cgit-server-url\u003e
	ServerURL() *URI
	SetServerURL(*URI)

	// Git webhook id.
	// Example: dvcsdweivcioekcxjhwiu4jd823idjekx
	ServerWebhookID() string
	SetServerWebhookID(string)

	// Subject for the commit.
	// Example: Updating codestream-service.yaml
	Subject() string
	SetSubject(string)

	// Branch name for which event is triggered.
	// Example: master
	TargetBranch() string
	SetTargetBranch(string)

	// Commit Time for the event.
	// Example: 2019-09-18T07:38:25.907+0000
	TimeStampInMicros() int64
	SetTimeStampInMicros(int64)

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

type gitEvent struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	commitIdField string

	createdAtField string

	createdByField string

	descriptionField *string

	executionIndexField int64

	executionLinkField string

	executionStatusField string

	httpUrlField string

	idField string

	messageField string

	nameField *string

	ownerField string

	pipelineField string

	projectField string

	repoField string

	serverUrlField *URI

	serverWebhookIdField string

	subjectField string

	targetBranchField string

	timeStampInMicrosField int64

	updatedAtField string

	updatedByField string

	versionField string
}

// CreateTimeInMicros gets the create time in micros of this polymorphic type
func (m *gitEvent) CreateTimeInMicros() int64 {
	return m.createTimeInMicrosField
}

// SetCreateTimeInMicros sets the create time in micros of this polymorphic type
func (m *gitEvent) SetCreateTimeInMicros(val int64) {
	m.createTimeInMicrosField = val
}

// Link gets the link of this polymorphic type
func (m *gitEvent) Link() string {
	return m.linkField
}

// SetLink sets the link of this polymorphic type
func (m *gitEvent) SetLink(val string) {
	m.linkField = val
}

// ProjectID gets the project Id of this polymorphic type
func (m *gitEvent) ProjectID() string {
	return m.projectIdField
}

// SetProjectID sets the project Id of this polymorphic type
func (m *gitEvent) SetProjectID(val string) {
	m.projectIdField = val
}

// UpdateTimeInMicros gets the update time in micros of this polymorphic type
func (m *gitEvent) UpdateTimeInMicros() int64 {
	return m.updateTimeInMicrosField
}

// SetUpdateTimeInMicros sets the update time in micros of this polymorphic type
func (m *gitEvent) SetUpdateTimeInMicros(val int64) {
	m.updateTimeInMicrosField = val
}

// CommitID gets the commit Id of this polymorphic type
func (m *gitEvent) CommitID() string {
	return m.commitIdField
}

// SetCommitID sets the commit Id of this polymorphic type
func (m *gitEvent) SetCommitID(val string) {
	m.commitIdField = val
}

// CreatedAt gets the created at of this polymorphic type
func (m *gitEvent) CreatedAt() string {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this polymorphic type
func (m *gitEvent) SetCreatedAt(val string) {
	m.createdAtField = val
}

// CreatedBy gets the created by of this polymorphic type
func (m *gitEvent) CreatedBy() string {
	return m.createdByField
}

// SetCreatedBy sets the created by of this polymorphic type
func (m *gitEvent) SetCreatedBy(val string) {
	m.createdByField = val
}

// Description gets the description of this polymorphic type
func (m *gitEvent) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *gitEvent) SetDescription(val *string) {
	m.descriptionField = val
}

// ExecutionIndex gets the execution index of this polymorphic type
func (m *gitEvent) ExecutionIndex() int64 {
	return m.executionIndexField
}

// SetExecutionIndex sets the execution index of this polymorphic type
func (m *gitEvent) SetExecutionIndex(val int64) {
	m.executionIndexField = val
}

// ExecutionLink gets the execution link of this polymorphic type
func (m *gitEvent) ExecutionLink() string {
	return m.executionLinkField
}

// SetExecutionLink sets the execution link of this polymorphic type
func (m *gitEvent) SetExecutionLink(val string) {
	m.executionLinkField = val
}

// ExecutionStatus gets the execution status of this polymorphic type
func (m *gitEvent) ExecutionStatus() string {
	return m.executionStatusField
}

// SetExecutionStatus sets the execution status of this polymorphic type
func (m *gitEvent) SetExecutionStatus(val string) {
	m.executionStatusField = val
}

// HTTPURL gets the http URL of this polymorphic type
func (m *gitEvent) HTTPURL() string {
	return m.httpUrlField
}

// SetHTTPURL sets the http URL of this polymorphic type
func (m *gitEvent) SetHTTPURL(val string) {
	m.httpUrlField = val
}

// ID gets the id of this polymorphic type
func (m *gitEvent) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *gitEvent) SetID(val string) {
	m.idField = val
}

// Message gets the message of this polymorphic type
func (m *gitEvent) Message() string {
	return m.messageField
}

// SetMessage sets the message of this polymorphic type
func (m *gitEvent) SetMessage(val string) {
	m.messageField = val
}

// Name gets the name of this polymorphic type
func (m *gitEvent) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *gitEvent) SetName(val *string) {
	m.nameField = val
}

// Owner gets the owner of this polymorphic type
func (m *gitEvent) Owner() string {
	return m.ownerField
}

// SetOwner sets the owner of this polymorphic type
func (m *gitEvent) SetOwner(val string) {
	m.ownerField = val
}

// Pipeline gets the pipeline of this polymorphic type
func (m *gitEvent) Pipeline() string {
	return m.pipelineField
}

// SetPipeline sets the pipeline of this polymorphic type
func (m *gitEvent) SetPipeline(val string) {
	m.pipelineField = val
}

// Project gets the project of this polymorphic type
func (m *gitEvent) Project() string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *gitEvent) SetProject(val string) {
	m.projectField = val
}

// Repo gets the repo of this polymorphic type
func (m *gitEvent) Repo() string {
	return m.repoField
}

// SetRepo sets the repo of this polymorphic type
func (m *gitEvent) SetRepo(val string) {
	m.repoField = val
}

// ServerURL gets the server URL of this polymorphic type
func (m *gitEvent) ServerURL() *URI {
	return m.serverUrlField
}

// SetServerURL sets the server URL of this polymorphic type
func (m *gitEvent) SetServerURL(val *URI) {
	m.serverUrlField = val
}

// ServerWebhookID gets the server webhook Id of this polymorphic type
func (m *gitEvent) ServerWebhookID() string {
	return m.serverWebhookIdField
}

// SetServerWebhookID sets the server webhook Id of this polymorphic type
func (m *gitEvent) SetServerWebhookID(val string) {
	m.serverWebhookIdField = val
}

// Subject gets the subject of this polymorphic type
func (m *gitEvent) Subject() string {
	return m.subjectField
}

// SetSubject sets the subject of this polymorphic type
func (m *gitEvent) SetSubject(val string) {
	m.subjectField = val
}

// TargetBranch gets the target branch of this polymorphic type
func (m *gitEvent) TargetBranch() string {
	return m.targetBranchField
}

// SetTargetBranch sets the target branch of this polymorphic type
func (m *gitEvent) SetTargetBranch(val string) {
	m.targetBranchField = val
}

// TimeStampInMicros gets the time stamp in micros of this polymorphic type
func (m *gitEvent) TimeStampInMicros() int64 {
	return m.timeStampInMicrosField
}

// SetTimeStampInMicros sets the time stamp in micros of this polymorphic type
func (m *gitEvent) SetTimeStampInMicros(val int64) {
	m.timeStampInMicrosField = val
}

// UpdatedAt gets the updated at of this polymorphic type
func (m *gitEvent) UpdatedAt() string {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this polymorphic type
func (m *gitEvent) SetUpdatedAt(val string) {
	m.updatedAtField = val
}

// UpdatedBy gets the updated by of this polymorphic type
func (m *gitEvent) UpdatedBy() string {
	return m.updatedByField
}

// SetUpdatedBy sets the updated by of this polymorphic type
func (m *gitEvent) SetUpdatedBy(val string) {
	m.updatedByField = val
}

// Version gets the version of this polymorphic type
func (m *gitEvent) Version() string {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *gitEvent) SetVersion(val string) {
	m.versionField = val
}

// UnmarshalGitEventSlice unmarshals polymorphic slices of GitEvent
func UnmarshalGitEventSlice(reader io.Reader, consumer runtime.Consumer) ([]GitEvent, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []GitEvent
	for _, element := range elements {
		obj, err := unmarshalGitEvent(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalGitEvent unmarshals polymorphic GitEvent
func UnmarshalGitEvent(reader io.Reader, consumer runtime.Consumer) (GitEvent, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalGitEvent(data, consumer)
}

func unmarshalGitEvent(data []byte, consumer runtime.Consumer) (GitEvent, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Git Event Model. property.
	var getType struct {
		GitEventModel string `json:"Git Event Model."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Git Event Model.", "body", getType.GitEventModel); err != nil {
		return nil, err
	}

	// The value of Git Event Model. is used to determine which type to create and unmarshal the data into
	switch getType.GitEventModel {
	case "GitEvent":
		var result gitEvent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Git Event Model. value: %q", getType.GitEventModel)
}

// Validate validates this git event
func (m *gitEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *gitEvent) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *gitEvent) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *gitEvent) validateServerURL(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerURL()) { // not required
		return nil
	}

	if m.ServerURL() != nil {
		if err := m.ServerURL().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverURL")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serverURL")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this git event based on the context it is used
func (m *gitEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *gitEvent) contextValidateServerURL(ctx context.Context, formats strfmt.Registry) error {

	if m.ServerURL() != nil {
		if err := m.ServerURL().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serverURL")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serverURL")
			}
			return err
		}
	}

	return nil
}