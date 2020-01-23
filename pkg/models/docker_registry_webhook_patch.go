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

// DockerRegistryWebhookPatch DockerRegistryWebhookPatch
//
// Docker Registry Webhook Patch details.
// swagger:discriminator DockerRegistryWebhookPatch Docker Registry Webhook Patch details.
type DockerRegistryWebhookPatch interface {
	runtime.Validatable

	// Indicates whether Docker webhook is enabled or not.
	Enabled() bool
	SetEnabled(bool)
}

type dockerRegistryWebhookPatch struct {
	enabledField bool
}

// Enabled gets the enabled of this polymorphic type
func (m *dockerRegistryWebhookPatch) Enabled() bool {
	return m.enabledField
}

// SetEnabled sets the enabled of this polymorphic type
func (m *dockerRegistryWebhookPatch) SetEnabled(val bool) {
	m.enabledField = val
}

// UnmarshalDockerRegistryWebhookPatchSlice unmarshals polymorphic slices of DockerRegistryWebhookPatch
func UnmarshalDockerRegistryWebhookPatchSlice(reader io.Reader, consumer runtime.Consumer) ([]DockerRegistryWebhookPatch, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []DockerRegistryWebhookPatch
	for _, element := range elements {
		obj, err := unmarshalDockerRegistryWebhookPatch(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalDockerRegistryWebhookPatch unmarshals polymorphic DockerRegistryWebhookPatch
func UnmarshalDockerRegistryWebhookPatch(reader io.Reader, consumer runtime.Consumer) (DockerRegistryWebhookPatch, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalDockerRegistryWebhookPatch(data, consumer)
}

func unmarshalDockerRegistryWebhookPatch(data []byte, consumer runtime.Consumer) (DockerRegistryWebhookPatch, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Docker Registry Webhook Patch details. property.
	var getType struct {
		DockerRegistryWebhookPatchDetails string `json:"Docker Registry Webhook Patch details."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Docker Registry Webhook Patch details.", "body", getType.DockerRegistryWebhookPatchDetails); err != nil {
		return nil, err
	}

	// The value of Docker Registry Webhook Patch details. is used to determine which type to create and unmarshal the data into
	switch getType.DockerRegistryWebhookPatchDetails {
	case "DockerRegistryWebhookPatch":
		var result dockerRegistryWebhookPatch
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid Docker Registry Webhook Patch details. value: %q", getType.DockerRegistryWebhookPatchDetails)

}

// Validate validates this docker registry webhook patch
func (m *dockerRegistryWebhookPatch) Validate(formats strfmt.Registry) error {
	return nil
}
