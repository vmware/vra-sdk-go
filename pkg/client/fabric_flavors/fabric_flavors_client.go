// Code generated by go-swagger; DO NOT EDIT.

package fabric_flavors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fabric flavors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fabric flavors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetFabricFlavors gets fabric flavors

Get all fabric flavors
*/
func (a *Client) GetFabricFlavors(params *GetFabricFlavorsParams) (*GetFabricFlavorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFabricFlavorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getFabricFlavors",
		Method:             "GET",
		PathPattern:        "/iaas/api/fabric-flavors",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFabricFlavorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFabricFlavorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFabricFlavors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
