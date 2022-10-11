// Code generated by go-swagger; DO NOT EDIT.

package fabric_vsphere_storage_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fabric v sphere storage policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric v sphere storage policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFabricVSphereStoragePolicies(params *GetFabricVSphereStoragePoliciesParams, opts ...ClientOption) (*GetFabricVSphereStoragePoliciesOK, error)

	GetFabricVSphereStoragePolicy(params *GetFabricVSphereStoragePolicyParams, opts ...ClientOption) (*GetFabricVSphereStoragePolicyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetFabricVSphereStoragePolicies gets fabric v sphere storage polices

Get all fabric vSphere storage polices.
*/
func (a *Client) GetFabricVSphereStoragePolicies(params *GetFabricVSphereStoragePoliciesParams, opts ...ClientOption) (*GetFabricVSphereStoragePoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricVSphereStoragePoliciesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFabricVSphereStoragePolicies",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-vsphere-storage-policies",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricVSphereStoragePoliciesReader{formats: a.formats},
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
	success, ok := result.(*GetFabricVSphereStoragePoliciesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricVSphereStoragePolicies: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFabricVSphereStoragePolicy gets fabric v sphere storage policy

Get fabric vSphere storage policy with a given id
*/
func (a *Client) GetFabricVSphereStoragePolicy(params *GetFabricVSphereStoragePolicyParams, opts ...ClientOption) (*GetFabricVSphereStoragePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricVSphereStoragePolicyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFabricVSphereStoragePolicy",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-vsphere-storage-policies/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricVSphereStoragePolicyReader{formats: a.formats},
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
	success, ok := result.(*GetFabricVSphereStoragePolicyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricVSphereStoragePolicy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
