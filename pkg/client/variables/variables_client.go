// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new variables API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for variables API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVariableUsingPOST(params *CreateVariableUsingPOSTParams, opts ...ClientOption) (*CreateVariableUsingPOSTOK, error)

	DeleteVariableByIDUsingDELETE(params *DeleteVariableByIDUsingDELETEParams, opts ...ClientOption) (*DeleteVariableByIDUsingDELETEOK, error)

	DeleteVariableByNameUsingDELETE(params *DeleteVariableByNameUsingDELETEParams, opts ...ClientOption) (*DeleteVariableByNameUsingDELETEOK, error)

	GetAllVariablesUsingGET(params *GetAllVariablesUsingGETParams, opts ...ClientOption) (*GetAllVariablesUsingGETOK, error)

	GetVariableByIDUsingGET(params *GetVariableByIDUsingGETParams, opts ...ClientOption) (*GetVariableByIDUsingGETOK, error)

	GetVariableByNameUsingGET(params *GetVariableByNameUsingGETParams, opts ...ClientOption) (*GetVariableByNameUsingGETOK, error)

	UpdateVariableByIDUsingPUT(params *UpdateVariableByIDUsingPUTParams, opts ...ClientOption) (*UpdateVariableByIDUsingPUTOK, error)

	UpdateVariableByNameUsingPUT(params *UpdateVariableByNameUsingPUTParams, opts ...ClientOption) (*UpdateVariableByNameUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVariableUsingPOST creates a variable

  Creates a Variable based on project name
*/
func (a *Client) CreateVariableUsingPOST(params *CreateVariableUsingPOSTParams, opts ...ClientOption) (*CreateVariableUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVariableUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createVariableUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/variables",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateVariableUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreateVariableUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createVariableUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVariableByIDUsingDELETE deletes a variable by Id

  Deletes a Variable with the given Id
*/
func (a *Client) DeleteVariableByIDUsingDELETE(params *DeleteVariableByIDUsingDELETEParams, opts ...ClientOption) (*DeleteVariableByIDUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVariableByIDUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVariableByIdUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/variables/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVariableByIDUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeleteVariableByIDUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVariableByIdUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVariableByNameUsingDELETE deletes a variable by project and name

  Deletes a Variable with the given name
*/
func (a *Client) DeleteVariableByNameUsingDELETE(params *DeleteVariableByNameUsingDELETEParams, opts ...ClientOption) (*DeleteVariableByNameUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVariableByNameUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteVariableByNameUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/variables/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVariableByNameUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeleteVariableByNameUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVariableByNameUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllVariablesUsingGET gets all variables

  Get all Variables with specified paging and filter parameters.
*/
func (a *Client) GetAllVariablesUsingGET(params *GetAllVariablesUsingGETParams, opts ...ClientOption) (*GetAllVariablesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllVariablesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllVariablesUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/variables",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllVariablesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllVariablesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllVariablesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVariableByIDUsingGET gets a variable

  Gets a Variable with the given id
*/
func (a *Client) GetVariableByIDUsingGET(params *GetVariableByIDUsingGETParams, opts ...ClientOption) (*GetVariableByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariableByIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVariableByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/variables/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVariableByIDUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetVariableByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariableByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVariableByNameUsingGET gets a variable by project and name

  Get an Variable with the given project and name
*/
func (a *Client) GetVariableByNameUsingGET(params *GetVariableByNameUsingGETParams, opts ...ClientOption) (*GetVariableByNameUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVariableByNameUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVariableByNameUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/variables/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVariableByNameUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetVariableByNameUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVariableByNameUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVariableByIDUsingPUT updates a variable by id

  Updates a Variable with the given id
*/
func (a *Client) UpdateVariableByIDUsingPUT(params *UpdateVariableByIDUsingPUTParams, opts ...ClientOption) (*UpdateVariableByIDUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVariableByIDUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVariableByIdUsingPUT",
		Method:             "PUT",
		PathPattern:        "/codestream/api/variables/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVariableByIDUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdateVariableByIDUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVariableByIdUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVariableByNameUsingPUT updates a variable by project and name

  Update an Variable with the given project and name
*/
func (a *Client) UpdateVariableByNameUsingPUT(params *UpdateVariableByNameUsingPUTParams, opts ...ClientOption) (*UpdateVariableByNameUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVariableByNameUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateVariableByNameUsingPUT",
		Method:             "PUT",
		PathPattern:        "/codestream/api/variables/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVariableByNameUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdateVariableByNameUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVariableByNameUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
