// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new compute gateway API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for compute gateway API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateComputeGateway(params *CreateComputeGatewayParams, opts ...ClientOption) (*CreateComputeGatewayAccepted, error)

	DeleteComputeGateway(params *DeleteComputeGatewayParams, opts ...ClientOption) (*DeleteComputeGatewayAccepted, error)

	GetComputeGateway(params *GetComputeGatewayParams, opts ...ClientOption) (*GetComputeGatewayOK, error)

	GetComputeGateways(params *GetComputeGatewaysParams, opts ...ClientOption) (*GetComputeGatewaysOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateComputeGateway creates a compute gateway

Create a new compute gateway.
*/
func (a *Client) CreateComputeGateway(params *CreateComputeGatewayParams, opts ...ClientOption) (*CreateComputeGatewayAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateComputeGatewayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createComputeGateway",
		Method:             "POST",
		PathPattern:        "/iaas/api/compute-gateways",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateComputeGatewayReader{formats: a.formats},
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
	success, ok := result.(*CreateComputeGatewayAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createComputeGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteComputeGateway deletes a compute gateway

Delete compute gateway with a given id
*/
func (a *Client) DeleteComputeGateway(params *DeleteComputeGatewayParams, opts ...ClientOption) (*DeleteComputeGatewayAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComputeGatewayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteComputeGateway",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/compute-gateways/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteComputeGatewayReader{formats: a.formats},
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
	success, ok := result.(*DeleteComputeGatewayAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteComputeGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetComputeGateway gets a compute gateway

Get compute gateway with a given id
*/
func (a *Client) GetComputeGateway(params *GetComputeGatewayParams, opts ...ClientOption) (*GetComputeGatewayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputeGatewayParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComputeGateway",
		Method:             "GET",
		PathPattern:        "/iaas/api/compute-gateways/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputeGatewayReader{formats: a.formats},
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
	success, ok := result.(*GetComputeGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComputeGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetComputeGateways gets compute gateways

Get all compute gateways
*/
func (a *Client) GetComputeGateways(params *GetComputeGatewaysParams, opts ...ClientOption) (*GetComputeGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputeGatewaysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComputeGateways",
		Method:             "GET",
		PathPattern:        "/iaas/api/compute-gateways",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputeGatewaysReader{formats: a.formats},
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
	success, ok := result.(*GetComputeGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComputeGateways: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
