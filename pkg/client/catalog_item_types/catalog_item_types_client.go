// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog item types API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog item types API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetTypeByIDUsingGET2(params *GetTypeByIDUsingGET2Params, opts ...ClientOption) (*GetTypeByIDUsingGET2OK, error)

	GetTypesUsingGET4(params *GetTypesUsingGET4Params, opts ...ClientOption) (*GetTypesUsingGET4OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetTypeByIDUsingGET2 fetches catalog item type associated with the specified ID

Returns the Catalog Item Type with the specified ID.
*/
func (a *Client) GetTypeByIDUsingGET2(params *GetTypeByIDUsingGET2Params, opts ...ClientOption) (*GetTypeByIDUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTypeByIDUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTypeByIdUsingGET_2",
		Method:             "GET",
		PathPattern:        "/catalog/api/types/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTypeByIDUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetTypeByIDUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTypeByIdUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTypesUsingGET4 finds all catalog item types

Returns a paginated list of all available Catalog Item Types.
*/
func (a *Client) GetTypesUsingGET4(params *GetTypesUsingGET4Params, opts ...ClientOption) (*GetTypesUsingGET4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTypesUsingGET4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTypesUsingGET_4",
		Method:             "GET",
		PathPattern:        "/catalog/api/types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTypesUsingGET4Reader{formats: a.formats},
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
	success, ok := result.(*GetTypesUsingGET4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTypesUsingGET_4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
