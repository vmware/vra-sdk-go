// Code generated by go-swagger; DO NOT EDIT.

package provider_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new provider requests API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for provider requests API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PatchUsingPATCH(params *PatchUsingPATCHParams, opts ...ClientOption) (*PatchUsingPATCHOK, error)

	PostUsingPOST(params *PostUsingPOSTParams, opts ...ClientOption) (*PostUsingPOSTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PatchUsingPATCH handles unfinished requests

Handles unfinished requests that are stopped because of service shutdown or a service outage.
*/
func (a *Client) PatchUsingPATCH(params *PatchUsingPATCHParams, opts ...ClientOption) (*PatchUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchUsingPATCHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/cmx/api/requests/blueprint/blueprint-provider-request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchUsingPATCHReader{formats: a.formats},
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
	success, ok := result.(*PatchUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostUsingPOST handles blueprint resources requests

Handles blueprint requests for resource operations - allocation, provisioning, removal
*/
func (a *Client) PostUsingPOST(params *PostUsingPOSTParams, opts ...ClientOption) (*PostUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postUsingPOST",
		Method:             "POST",
		PathPattern:        "/cmx/api/requests/blueprint/blueprint-provider-request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*PostUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
