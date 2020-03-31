// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog entitlements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog entitlements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateEntitlementUsingPOST creates an entitlement

Creates an entitlement for a given project.
*/
func (a *Client) CreateEntitlementUsingPOST(params *CreateEntitlementUsingPOSTParams) (*CreateEntitlementUsingPOSTOK, *CreateEntitlementUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEntitlementUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEntitlementUsingPOST",
		Method:             "POST",
		PathPattern:        "/catalog/api/admin/entitlements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateEntitlementUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateEntitlementUsingPOSTOK:
		return value, nil, nil
	case *CreateEntitlementUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for catalog_entitlements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteEntitlementUsingDELETE deletes an entitlement

Deletes the entitlement with the specified id.
*/
func (a *Client) DeleteEntitlementUsingDELETE(params *DeleteEntitlementUsingDELETEParams) (*DeleteEntitlementUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEntitlementUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEntitlementUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/catalog/api/admin/entitlements/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteEntitlementUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteEntitlementUsingDELETENoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteEntitlementUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEntitlementsUsingGET returns all entitlements filtered by project Id

Returns all entitlements (filtered by projectId).
*/
func (a *Client) GetEntitlementsUsingGET(params *GetEntitlementsUsingGETParams) (*GetEntitlementsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEntitlementsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEntitlementsUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/admin/entitlements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEntitlementsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEntitlementsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEntitlementsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
