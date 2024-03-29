// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog entitlements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog entitlements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateEntitlementUsingPOST2(params *CreateEntitlementUsingPOST2Params, opts ...ClientOption) (*CreateEntitlementUsingPOST2OK, *CreateEntitlementUsingPOST2Created, error)

	DeleteEntitlementUsingDELETE2(params *DeleteEntitlementUsingDELETE2Params, opts ...ClientOption) (*DeleteEntitlementUsingDELETE2NoContent, error)

	GetEntitlementsUsingGET2(params *GetEntitlementsUsingGET2Params, opts ...ClientOption) (*GetEntitlementsUsingGET2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateEntitlementUsingPOST2 creates an entitlement

Creates an entitlement for a given project.
*/
func (a *Client) CreateEntitlementUsingPOST2(params *CreateEntitlementUsingPOST2Params, opts ...ClientOption) (*CreateEntitlementUsingPOST2OK, *CreateEntitlementUsingPOST2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEntitlementUsingPOST2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createEntitlementUsingPOST_2",
		Method:             "POST",
		PathPattern:        "/catalog/api/admin/entitlements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEntitlementUsingPOST2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateEntitlementUsingPOST2OK:
		return value, nil, nil
	case *CreateEntitlementUsingPOST2Created:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for catalog_entitlements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteEntitlementUsingDELETE2 deletes an entitlement

Deletes the entitlement with the specified id.
*/
func (a *Client) DeleteEntitlementUsingDELETE2(params *DeleteEntitlementUsingDELETE2Params, opts ...ClientOption) (*DeleteEntitlementUsingDELETE2NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntitlementUsingDELETE2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteEntitlementUsingDELETE_2",
		Method:             "DELETE",
		PathPattern:        "/catalog/api/admin/entitlements/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEntitlementUsingDELETE2Reader{formats: a.formats},
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
	success, ok := result.(*DeleteEntitlementUsingDELETE2NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEntitlementUsingDELETE_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEntitlementsUsingGET2 returns all entitlements filtered by project Id

Returns all entitlements (filtered by projectId).
*/
func (a *Client) GetEntitlementsUsingGET2(params *GetEntitlementsUsingGET2Params, opts ...ClientOption) (*GetEntitlementsUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitlementsUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEntitlementsUsingGET_2",
		Method:             "GET",
		PathPattern:        "/catalog/api/admin/entitlements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitlementsUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetEntitlementsUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEntitlementsUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
