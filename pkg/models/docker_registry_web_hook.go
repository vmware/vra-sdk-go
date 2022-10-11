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
	"github.com/go-openapi/validate"
)

// DockerRegistryWebHook DockerRegistryWebHook
//
// Docker Registry Webhook details.
//
// swagger:discriminator DockerRegistryWebHook Docker Registry Webhook details.
type DockerRegistryWebHook interface {
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

	// Indicates whether Docker webhook is enabled or not.
	// Example: true
	Enabled() bool
	SetEnabled(bool)

	// Name of the Docker Endpoint.
	// Example: Docker Endpoint
	Endpoint() string
	SetEndpoint(string)

	// Docker webhook listener link.
	// Example: codestream/api/registry-webhooks/sdscfvefw34rfrdsvdsat43erwfdf
	ExternalListenerLink() string
	SetExternalListenerLink(string)

	// The id of this resource.
	// Example: 8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	ID() string
	SetID(string)

	// If provided then the pipeline execution is triggered only when the given image name regex matches the image name in the received payload.
	// Example: admin/wordpress
	ImageNameRegExPattern() string
	SetImageNameRegExPattern(string)

	// Pipeline execution input properties.
	// Example: [{"ip":"10.5.23.84","script":"testScript.sh"}]
	Input() interface{}
	SetInput(interface{})

	// A human-friendly name used as an identifier in APIs that support this option
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	// Pipeline name which is meant to be triggered when a docker event occur.
	// Example: DemoPipeline
	Pipeline() string
	SetPipeline(string)

	// The project this entity belongs to.
	// Example: My-Project
	Project() string
	SetProject(string)

	// This token is used to authenticate when calling VMware Cloud Services APIs. These tokens are scoped within the organization.
	// Example: cKNNVCSHijnaxlrfnVsxUYr6wM2g5Bg11tfnotmWb9XdA5kpUCvI2ubJojTIGp9g
	RefreshToken() string
	SetRefreshToken(string)

	// Docker Repo Name.
	// Example: MyRepository
	RepoName() string
	SetRepoName(string)

	// Secret token to validate received payloads.
	// Example: dvcdefrffsdvdfvfdgfdvbfdbvrefg3ff=
	SecretToken() string
	SetSecretToken(string)

	// Docker server type.
	// Example: DockerHub/Docker_Trusted_Registry
	ServerType() string
	SetServerType(string)

	// Docker webhook name.
	// Example: docker-webhook
	Slug() string
	SetSlug(string)

	// If provided then the pipeline execution is triggered only when the given tag name regex matches the tag name(s) in the received payload.
	// Example: tag1
	TagNamePattern() string
	SetTagNamePattern(string)

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

type dockerRegistryWebHook struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	createdAtField string

	createdByField string

	descriptionField *string

	enabledField bool

	endpointField string

	externalListenerLinkField string

	idField string

	imageNameRegExPatternField string

	inputField interface{}

	nameField *string

	pipelineField string

	projectField string

	refreshTokenField string

	repoNameField string

	secretTokenField string

	serverTypeField string

	slugField string

	tagNamePatternField string

	updatedAtField string

	updatedByField string

	versionField string
}

// CreateTimeInMicros gets the create time in micros of this polymorphic type
func (m *dockerRegistryWebHook) CreateTimeInMicros() int64 {
	return m.createTimeInMicrosField
}

// SetCreateTimeInMicros sets the create time in micros of this polymorphic type
func (m *dockerRegistryWebHook) SetCreateTimeInMicros(val int64) {
	m.createTimeInMicrosField = val
}

// Link gets the link of this polymorphic type
func (m *dockerRegistryWebHook) Link() string {
	return m.linkField
}

// SetLink sets the link of this polymorphic type
func (m *dockerRegistryWebHook) SetLink(val string) {
	m.linkField = val
}

// ProjectID gets the project Id of this polymorphic type
func (m *dockerRegistryWebHook) ProjectID() string {
	return m.projectIdField
}

// SetProjectID sets the project Id of this polymorphic type
func (m *dockerRegistryWebHook) SetProjectID(val string) {
	m.projectIdField = val
}

// UpdateTimeInMicros gets the update time in micros of this polymorphic type
func (m *dockerRegistryWebHook) UpdateTimeInMicros() int64 {
	return m.updateTimeInMicrosField
}

// SetUpdateTimeInMicros sets the update time in micros of this polymorphic type
func (m *dockerRegistryWebHook) SetUpdateTimeInMicros(val int64) {
	m.updateTimeInMicrosField = val
}

// CreatedAt gets the created at of this polymorphic type
func (m *dockerRegistryWebHook) CreatedAt() string {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this polymorphic type
func (m *dockerRegistryWebHook) SetCreatedAt(val string) {
	m.createdAtField = val
}

// CreatedBy gets the created by of this polymorphic type
func (m *dockerRegistryWebHook) CreatedBy() string {
	return m.createdByField
}

// SetCreatedBy sets the created by of this polymorphic type
func (m *dockerRegistryWebHook) SetCreatedBy(val string) {
	m.createdByField = val
}

// Description gets the description of this polymorphic type
func (m *dockerRegistryWebHook) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *dockerRegistryWebHook) SetDescription(val *string) {
	m.descriptionField = val
}

// Enabled gets the enabled of this polymorphic type
func (m *dockerRegistryWebHook) Enabled() bool {
	return m.enabledField
}

// SetEnabled sets the enabled of this polymorphic type
func (m *dockerRegistryWebHook) SetEnabled(val bool) {
	m.enabledField = val
}

// Endpoint gets the endpoint of this polymorphic type
func (m *dockerRegistryWebHook) Endpoint() string {
	return m.endpointField
}

// SetEndpoint sets the endpoint of this polymorphic type
func (m *dockerRegistryWebHook) SetEndpoint(val string) {
	m.endpointField = val
}

// ExternalListenerLink gets the external listener link of this polymorphic type
func (m *dockerRegistryWebHook) ExternalListenerLink() string {
	return m.externalListenerLinkField
}

// SetExternalListenerLink sets the external listener link of this polymorphic type
func (m *dockerRegistryWebHook) SetExternalListenerLink(val string) {
	m.externalListenerLinkField = val
}

// ID gets the id of this polymorphic type
func (m *dockerRegistryWebHook) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *dockerRegistryWebHook) SetID(val string) {
	m.idField = val
}

// ImageNameRegExPattern gets the image name reg ex pattern of this polymorphic type
func (m *dockerRegistryWebHook) ImageNameRegExPattern() string {
	return m.imageNameRegExPatternField
}

// SetImageNameRegExPattern sets the image name reg ex pattern of this polymorphic type
func (m *dockerRegistryWebHook) SetImageNameRegExPattern(val string) {
	m.imageNameRegExPatternField = val
}

// Input gets the input of this polymorphic type
func (m *dockerRegistryWebHook) Input() interface{} {
	return m.inputField
}

// SetInput sets the input of this polymorphic type
func (m *dockerRegistryWebHook) SetInput(val interface{}) {
	m.inputField = val
}

// Name gets the name of this polymorphic type
func (m *dockerRegistryWebHook) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *dockerRegistryWebHook) SetName(val *string) {
	m.nameField = val
}

// Pipeline gets the pipeline of this polymorphic type
func (m *dockerRegistryWebHook) Pipeline() string {
	return m.pipelineField
}

// SetPipeline sets the pipeline of this polymorphic type
func (m *dockerRegistryWebHook) SetPipeline(val string) {
	m.pipelineField = val
}

// Project gets the project of this polymorphic type
func (m *dockerRegistryWebHook) Project() string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *dockerRegistryWebHook) SetProject(val string) {
	m.projectField = val
}

// RefreshToken gets the refresh token of this polymorphic type
func (m *dockerRegistryWebHook) RefreshToken() string {
	return m.refreshTokenField
}

// SetRefreshToken sets the refresh token of this polymorphic type
func (m *dockerRegistryWebHook) SetRefreshToken(val string) {
	m.refreshTokenField = val
}

// RepoName gets the repo name of this polymorphic type
func (m *dockerRegistryWebHook) RepoName() string {
	return m.repoNameField
}

// SetRepoName sets the repo name of this polymorphic type
func (m *dockerRegistryWebHook) SetRepoName(val string) {
	m.repoNameField = val
}

// SecretToken gets the secret token of this polymorphic type
func (m *dockerRegistryWebHook) SecretToken() string {
	return m.secretTokenField
}

// SetSecretToken sets the secret token of this polymorphic type
func (m *dockerRegistryWebHook) SetSecretToken(val string) {
	m.secretTokenField = val
}

// ServerType gets the server type of this polymorphic type
func (m *dockerRegistryWebHook) ServerType() string {
	return m.serverTypeField
}

// SetServerType sets the server type of this polymorphic type
func (m *dockerRegistryWebHook) SetServerType(val string) {
	m.serverTypeField = val
}

// Slug gets the slug of this polymorphic type
func (m *dockerRegistryWebHook) Slug() string {
	return m.slugField
}

// SetSlug sets the slug of this polymorphic type
func (m *dockerRegistryWebHook) SetSlug(val string) {
	m.slugField = val
}

// TagNamePattern gets the tag name pattern of this polymorphic type
func (m *dockerRegistryWebHook) TagNamePattern() string {
	return m.tagNamePatternField
}

// SetTagNamePattern sets the tag name pattern of this polymorphic type
func (m *dockerRegistryWebHook) SetTagNamePattern(val string) {
	m.tagNamePatternField = val
}

// UpdatedAt gets the updated at of this polymorphic type
func (m *dockerRegistryWebHook) UpdatedAt() string {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this polymorphic type
func (m *dockerRegistryWebHook) SetUpdatedAt(val string) {
	m.updatedAtField = val
}

// UpdatedBy gets the updated by of this polymorphic type
func (m *dockerRegistryWebHook) UpdatedBy() string {
	return m.updatedByField
}

// SetUpdatedBy sets the updated by of this polymorphic type
func (m *dockerRegistryWebHook) SetUpdatedBy(val string) {
	m.updatedByField = val
}

// Version gets the version of this polymorphic type
func (m *dockerRegistryWebHook) Version() string {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *dockerRegistryWebHook) SetVersion(val string) {
	m.versionField = val
}

// UnmarshalDockerRegistryWebHookSlice unmarshals polymorphic slices of DockerRegistryWebHook
func UnmarshalDockerRegistryWebHookSlice(reader io.Reader, consumer runtime.Consumer) ([]DockerRegistryWebHook, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []DockerRegistryWebHook
	for _, element := range elements {
		obj, err := unmarshalDockerRegistryWebHook(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalDockerRegistryWebHook unmarshals polymorphic DockerRegistryWebHook
func UnmarshalDockerRegistryWebHook(reader io.Reader, consumer runtime.Consumer) (DockerRegistryWebHook, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalDockerRegistryWebHook(data, consumer)
}

func unmarshalDockerRegistryWebHook(data []byte, consumer runtime.Consumer) (DockerRegistryWebHook, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Docker Registry Webhook details. property.
	var getType struct {
		DockerRegistryWebhookDetails string `json:"Docker Registry Webhook details."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Docker Registry Webhook details.", "body", getType.DockerRegistryWebhookDetails); err != nil {
		return nil, err
	}

	// The value of Docker Registry Webhook details. is used to determine which type to create and unmarshal the data into
	switch getType.DockerRegistryWebhookDetails {
	case "DockerRegistryWebHook":
		var result dockerRegistryWebHook
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Docker Registry Webhook details. value: %q", getType.DockerRegistryWebhookDetails)
}

// Validate validates this docker registry web hook
func (m *dockerRegistryWebHook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
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

func (m *dockerRegistryWebHook) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *dockerRegistryWebHook) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this docker registry web hook based on context it is used
func (m *dockerRegistryWebHook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
