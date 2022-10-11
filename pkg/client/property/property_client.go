// Code generated by go-swagger; DO NOT EDIT.

package property

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new property API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for property API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteConfigurationProperty(params *DeleteConfigurationPropertyParams, opts ...ClientOption) (*DeleteConfigurationPropertyOK, error)

	GetConfigurationProperties(params *GetConfigurationPropertiesParams, opts ...ClientOption) (*GetConfigurationPropertiesOK, error)

	GetConfigurationProperty(params *GetConfigurationPropertyParams, opts ...ClientOption) (*GetConfigurationPropertyOK, error)

	PatchConfigurationProperty(params *PatchConfigurationPropertyParams, opts ...ClientOption) (*PatchConfigurationPropertyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteConfigurationProperty deletes a configuration property

Delete a configuration property
*/
func (a *Client) DeleteConfigurationProperty(params *DeleteConfigurationPropertyParams, opts ...ClientOption) (*DeleteConfigurationPropertyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteConfigurationPropertyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteConfigurationProperty",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/configuration-properties/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteConfigurationPropertyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteConfigurationPropertyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteConfigurationProperty: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetConfigurationProperties gets configuration properties

Get all configuration properties
*/
func (a *Client) GetConfigurationProperties(params *GetConfigurationPropertiesParams, opts ...ClientOption) (*GetConfigurationPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationPropertiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConfigurationProperties",
		Method:             "GET",
		PathPattern:        "/iaas/api/configuration-properties",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConfigurationPropertiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConfigurationPropertiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConfigurationProperties: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetConfigurationProperty gets single configuration property

Get single configuration property
*/
func (a *Client) GetConfigurationProperty(params *GetConfigurationPropertyParams, opts ...ClientOption) (*GetConfigurationPropertyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationPropertyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConfigurationProperty",
		Method:             "GET",
		PathPattern:        "/iaas/api/configuration-properties/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConfigurationPropertyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConfigurationPropertyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConfigurationProperty: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchConfigurationProperty updates or create configuration property

Update or create configuration property.
*/
func (a *Client) PatchConfigurationProperty(params *PatchConfigurationPropertyParams, opts ...ClientOption) (*PatchConfigurationPropertyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchConfigurationPropertyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchConfigurationProperty",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/configuration-properties",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchConfigurationPropertyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchConfigurationPropertyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchConfigurationProperty: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}