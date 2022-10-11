// Code generated by go-swagger; DO NOT EDIT.

package onboarding_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new onboarding resources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for onboarding resources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateResource(params *CreateResourceParams, opts ...ClientOption) (*CreateResourceOK, error)

	DeleteResourceByID(params *DeleteResourceByIDParams, opts ...ClientOption) (*DeleteResourceByIDOK, error)

	GetResourceByID(params *GetResourceByIDParams, opts ...ClientOption) (*GetResourceByIDOK, error)

	QueryResources(params *QueryResourcesParams, opts ...ClientOption) (*QueryResourcesOK, error)

	UpdateResourceByID(params *UpdateResourceByIDParams, opts ...ClientOption) (*UpdateResourceByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateResource creates an onboarding resource
*/
func (a *Client) CreateResource(params *CreateResourceParams, opts ...ClientOption) (*CreateResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createResource",
		Method:             "POST",
		PathPattern:        "/relocation/onboarding/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateResourceReader{formats: a.formats},
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
	success, ok := result.(*CreateResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteResourceByID deletes the selected onboarding resource
*/
func (a *Client) DeleteResourceByID(params *DeleteResourceByIDParams, opts ...ClientOption) (*DeleteResourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteResourceByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteResourceById",
		Method:             "DELETE",
		PathPattern:        "/relocation/onboarding/resource/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteResourceByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteResourceByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteResourceById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetResourceByID gets the selected onboarding resource
*/
func (a *Client) GetResourceByID(params *GetResourceByIDParams, opts ...ClientOption) (*GetResourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getResourceById",
		Method:             "GET",
		PathPattern:        "/relocation/onboarding/resource/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceByIDReader{formats: a.formats},
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
	success, ok := result.(*GetResourceByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourceById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryResources queries for onboarding resources
*/
func (a *Client) QueryResources(params *QueryResourcesParams, opts ...ClientOption) (*QueryResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryResourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryResources",
		Method:             "GET",
		PathPattern:        "/relocation/onboarding/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryResourcesReader{formats: a.formats},
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
	success, ok := result.(*QueryResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateResourceByID updates the selected onboarding resource
*/
func (a *Client) UpdateResourceByID(params *UpdateResourceByIDParams, opts ...ClientOption) (*UpdateResourceByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateResourceByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateResourceById",
		Method:             "PATCH",
		PathPattern:        "/relocation/onboarding/resource/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateResourceByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateResourceByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateResourceById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}