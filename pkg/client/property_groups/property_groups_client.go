// Code generated by go-swagger; DO NOT EDIT.

package property_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new property groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for property groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePropertyGroupUsingPOST(params *CreatePropertyGroupUsingPOSTParams, opts ...ClientOption) (*CreatePropertyGroupUsingPOSTCreated, error)

	DeletePropertyGroupUsingDELETE(params *DeletePropertyGroupUsingDELETEParams, opts ...ClientOption) (*DeletePropertyGroupUsingDELETENoContent, error)

	GetPropertyGroupUsingGET(params *GetPropertyGroupUsingGETParams, opts ...ClientOption) (*GetPropertyGroupUsingGETOK, error)

	ListPropertyGroupsUsingGET(params *ListPropertyGroupsUsingGETParams, opts ...ClientOption) (*ListPropertyGroupsUsingGETOK, error)

	UpdatePropertyGroupUsingPUT(params *UpdatePropertyGroupUsingPUTParams, opts ...ClientOption) (*UpdatePropertyGroupUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePropertyGroupUsingPOST creates a property group
*/
func (a *Client) CreatePropertyGroupUsingPOST(params *CreatePropertyGroupUsingPOSTParams, opts ...ClientOption) (*CreatePropertyGroupUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePropertyGroupUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPropertyGroupUsingPOST",
		Method:             "POST",
		PathPattern:        "/properties/api/property-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePropertyGroupUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreatePropertyGroupUsingPOSTCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPropertyGroupUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePropertyGroupUsingDELETE deletes a property group
*/
func (a *Client) DeletePropertyGroupUsingDELETE(params *DeletePropertyGroupUsingDELETEParams, opts ...ClientOption) (*DeletePropertyGroupUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePropertyGroupUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePropertyGroupUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/properties/api/property-groups/{propertyGroupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePropertyGroupUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeletePropertyGroupUsingDELETENoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePropertyGroupUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPropertyGroupUsingGET returns property group details
*/
func (a *Client) GetPropertyGroupUsingGET(params *GetPropertyGroupUsingGETParams, opts ...ClientOption) (*GetPropertyGroupUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPropertyGroupUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPropertyGroupUsingGET",
		Method:             "GET",
		PathPattern:        "/properties/api/property-groups/{propertyGroupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPropertyGroupUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetPropertyGroupUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPropertyGroupUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPropertyGroupsUsingGET lists property groups
*/
func (a *Client) ListPropertyGroupsUsingGET(params *ListPropertyGroupsUsingGETParams, opts ...ClientOption) (*ListPropertyGroupsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPropertyGroupsUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPropertyGroupsUsingGET",
		Method:             "GET",
		PathPattern:        "/properties/api/property-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListPropertyGroupsUsingGETReader{formats: a.formats},
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
	success, ok := result.(*ListPropertyGroupsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPropertyGroupsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePropertyGroupUsingPUT updates a property group
*/
func (a *Client) UpdatePropertyGroupUsingPUT(params *UpdatePropertyGroupUsingPUTParams, opts ...ClientOption) (*UpdatePropertyGroupUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePropertyGroupUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePropertyGroupUsingPUT",
		Method:             "PUT",
		PathPattern:        "/properties/api/property-groups/{propertyGroupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePropertyGroupUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdatePropertyGroupUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePropertyGroupUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}