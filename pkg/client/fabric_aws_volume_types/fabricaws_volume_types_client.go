// Code generated by go-swagger; DO NOT EDIT.

package fabric_aws_volume_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fabric a w s volume types API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric a w s volume types API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFabricAwsVolumeTypes(params *GetFabricAwsVolumeTypesParams, opts ...ClientOption) (*GetFabricAwsVolumeTypesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetFabricAwsVolumeTypes gets fabric a w s volume types

Get all fabric AWS volume types.
*/
func (a *Client) GetFabricAwsVolumeTypes(params *GetFabricAwsVolumeTypesParams, opts ...ClientOption) (*GetFabricAwsVolumeTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricAwsVolumeTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFabricAwsVolumeTypes",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-aws-volume-types",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricAwsVolumeTypesReader{formats: a.formats},
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
	success, ok := result.(*GetFabricAwsVolumeTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricAwsVolumeTypes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
