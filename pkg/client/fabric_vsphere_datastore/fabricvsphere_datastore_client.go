// Code generated by go-swagger; DO NOT EDIT.

package fabric_vsphere_datastore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fabric v sphere datastore API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric v sphere datastore API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFabricVSphereDatastore gets fabric v sphere datastore

Get fabric vSphere datastore with a given id
*/
func (a *Client) GetFabricVSphereDatastore(params *GetFabricVSphereDatastoreParams) (*GetFabricVSphereDatastoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricVSphereDatastoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricVSphereDatastore",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-vsphere-datastores/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricVSphereDatastoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFabricVSphereDatastoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricVSphereDatastore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFabricVSphereDatastores gets fabric v sphere datastores

Get all fabric vSphere datastores.
1. If regionId is not provided as a parameter then all datastores are returned
2. If regionId is provided as a parameter then datastores iin that region are returned
*/
func (a *Client) GetFabricVSphereDatastores(params *GetFabricVSphereDatastoresParams) (*GetFabricVSphereDatastoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricVSphereDatastoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricVSphereDatastores",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-vsphere-datastores",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricVSphereDatastoresReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFabricVSphereDatastoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricVSphereDatastores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
