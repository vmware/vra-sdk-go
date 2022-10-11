// Code generated by go-swagger; DO NOT EDIT.

package onboarding_blueprints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new onboarding blueprints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for onboarding blueprints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBlueprint(params *CreateBlueprintParams, opts ...ClientOption) (*CreateBlueprintOK, error)

	CreateBlueprintByID(params *CreateBlueprintByIDParams, opts ...ClientOption) (*CreateBlueprintByIDOK, error)

	DeleteBlueprintByID(params *DeleteBlueprintByIDParams, opts ...ClientOption) (*DeleteBlueprintByIDOK, error)

	QueryBlueprints(params *QueryBlueprintsParams, opts ...ClientOption) (*QueryBlueprintsOK, error)

	UpdateBlueprintByID(params *UpdateBlueprintByIDParams, opts ...ClientOption) (*UpdateBlueprintByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBlueprint creates an onboarding blueprints
*/
func (a *Client) CreateBlueprint(params *CreateBlueprintParams, opts ...ClientOption) (*CreateBlueprintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlueprintParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBlueprint",
		Method:             "POST",
		PathPattern:        "/relocation/onboarding/blueprint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBlueprintReader{formats: a.formats},
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
	success, ok := result.(*CreateBlueprintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlueprint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateBlueprintByID gets the selected onboarding blueprint
*/
func (a *Client) CreateBlueprintByID(params *CreateBlueprintByIDParams, opts ...ClientOption) (*CreateBlueprintByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlueprintByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBlueprintById",
		Method:             "GET",
		PathPattern:        "/relocation/onboarding/blueprint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBlueprintByIDReader{formats: a.formats},
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
	success, ok := result.(*CreateBlueprintByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlueprintById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteBlueprintByID deletes the selected onboarding blueprint
*/
func (a *Client) DeleteBlueprintByID(params *DeleteBlueprintByIDParams, opts ...ClientOption) (*DeleteBlueprintByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlueprintByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBlueprintById",
		Method:             "DELETE",
		PathPattern:        "/relocation/onboarding/blueprint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBlueprintByIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteBlueprintByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBlueprintById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryBlueprints queries for onboarding blueprints
*/
func (a *Client) QueryBlueprints(params *QueryBlueprintsParams, opts ...ClientOption) (*QueryBlueprintsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryBlueprintsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryBlueprints",
		Method:             "GET",
		PathPattern:        "/relocation/onboarding/blueprint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryBlueprintsReader{formats: a.formats},
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
	success, ok := result.(*QueryBlueprintsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryBlueprints: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateBlueprintByID updates the selected onboarding blueprint
*/
func (a *Client) UpdateBlueprintByID(params *UpdateBlueprintByIDParams, opts ...ClientOption) (*UpdateBlueprintByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBlueprintByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBlueprintById",
		Method:             "PATCH",
		PathPattern:        "/relocation/onboarding/blueprint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBlueprintByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateBlueprintByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBlueprintById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
