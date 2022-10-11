// Code generated by go-swagger; DO NOT EDIT.

package network_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new network profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNetworkProfile(params *CreateNetworkProfileParams, opts ...ClientOption) (*CreateNetworkProfileCreated, error)

	DeleteNetworkProfile(params *DeleteNetworkProfileParams, opts ...ClientOption) (*DeleteNetworkProfileNoContent, error)

	GetNetworkProfile(params *GetNetworkProfileParams, opts ...ClientOption) (*GetNetworkProfileOK, error)

	GetNetworkProfiles(params *GetNetworkProfilesParams, opts ...ClientOption) (*GetNetworkProfilesOK, error)

	UpdateNetworkProfile(params *UpdateNetworkProfileParams, opts ...ClientOption) (*UpdateNetworkProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateNetworkProfile creates network profile

Create network profile
*/
func (a *Client) CreateNetworkProfile(params *CreateNetworkProfileParams, opts ...ClientOption) (*CreateNetworkProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createNetworkProfile",
		Method:             "POST",
		PathPattern:        "/iaas/api/network-profiles",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNetworkProfileReader{formats: a.formats},
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
	success, ok := result.(*CreateNetworkProfileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createNetworkProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteNetworkProfile deletes network profile

Delete network profile with a given id
*/
func (a *Client) DeleteNetworkProfile(params *DeleteNetworkProfileParams, opts ...ClientOption) (*DeleteNetworkProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteNetworkProfile",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/network-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNetworkProfileReader{formats: a.formats},
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
	success, ok := result.(*DeleteNetworkProfileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNetworkProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNetworkProfile gets network profile

Get network profile with a given id
*/
func (a *Client) GetNetworkProfile(params *GetNetworkProfileParams, opts ...ClientOption) (*GetNetworkProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNetworkProfile",
		Method:             "GET",
		PathPattern:        "/iaas/api/network-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkProfileReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNetworkProfiles gets network profiles

Get all network profiles
*/
func (a *Client) GetNetworkProfiles(params *GetNetworkProfilesParams, opts ...ClientOption) (*GetNetworkProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNetworkProfiles",
		Method:             "GET",
		PathPattern:        "/iaas/api/network-profiles",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkProfilesReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateNetworkProfile updates network profile

Update network profile
*/
func (a *Client) UpdateNetworkProfile(params *UpdateNetworkProfileParams, opts ...ClientOption) (*UpdateNetworkProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateNetworkProfile",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/network-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkProfileReader{formats: a.formats},
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
	success, ok := result.(*UpdateNetworkProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
