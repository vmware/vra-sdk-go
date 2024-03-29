// Code generated by go-swagger; DO NOT EDIT.

package onboarding_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new onboarding deployments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for onboarding deployments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeploymentMixin6(params *CreateDeploymentMixin6Params, opts ...ClientOption) (*CreateDeploymentMixin6OK, error)

	CreateDeploymentWithResources(params *CreateDeploymentWithResourcesParams, opts ...ClientOption) (*CreateDeploymentWithResourcesOK, error)

	CreateDeploymentsBulk(params *CreateDeploymentsBulkParams, opts ...ClientOption) (*CreateDeploymentsBulkOK, error)

	DeleteDeploymentBy(params *DeleteDeploymentByParams, opts ...ClientOption) (*DeleteDeploymentByOK, error)

	GetDeploymentByID(params *GetDeploymentByIDParams, opts ...ClientOption) (*GetDeploymentByIDOK, error)

	QueryDeployments(params *QueryDeploymentsParams, opts ...ClientOption) (*QueryDeploymentsOK, error)

	UpdateDeploymentByID(params *UpdateDeploymentByIDParams, opts ...ClientOption) (*UpdateDeploymentByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDeploymentMixin6 creates an onboarding deployment
*/
func (a *Client) CreateDeploymentMixin6(params *CreateDeploymentMixin6Params, opts ...ClientOption) (*CreateDeploymentMixin6OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentMixin6Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeploymentMixin6",
		Method:             "POST",
		PathPattern:        "/relocation/onboarding/deployment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeploymentMixin6Reader{formats: a.formats},
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
	success, ok := result.(*CreateDeploymentMixin6OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeploymentMixin6: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateDeploymentWithResources creates or updates an onboarding deployment and associated resources
*/
func (a *Client) CreateDeploymentWithResources(params *CreateDeploymentWithResourcesParams, opts ...ClientOption) (*CreateDeploymentWithResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentWithResourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeploymentWithResources",
		Method:             "POST",
		PathPattern:        "/relocation/onboarding/task/create-deployment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeploymentWithResourcesReader{formats: a.formats},
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
	success, ok := result.(*CreateDeploymentWithResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeploymentWithResources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateDeploymentsBulk creates or updates onboarding deployments in bulk
*/
func (a *Client) CreateDeploymentsBulk(params *CreateDeploymentsBulkParams, opts ...ClientOption) (*CreateDeploymentsBulkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentsBulkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeploymentsBulk",
		Method:             "POST",
		PathPattern:        "/relocation/onboarding/task/create-deployment-bulk",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeploymentsBulkReader{formats: a.formats},
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
	success, ok := result.(*CreateDeploymentsBulkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeploymentsBulk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDeploymentBy deletes the selected onboarding deployment
*/
func (a *Client) DeleteDeploymentBy(params *DeleteDeploymentByParams, opts ...ClientOption) (*DeleteDeploymentByOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentByParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDeploymentBy",
		Method:             "DELETE",
		PathPattern:        "/relocation/onboarding/deployment/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeploymentByReader{formats: a.formats},
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
	success, ok := result.(*DeleteDeploymentByOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDeploymentBy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeploymentByID gets the selected onboarding deployment
*/
func (a *Client) GetDeploymentByID(params *GetDeploymentByIDParams, opts ...ClientOption) (*GetDeploymentByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeploymentById",
		Method:             "GET",
		PathPattern:        "/relocation/onboarding/deployment/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentByIDReader{formats: a.formats},
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
	success, ok := result.(*GetDeploymentByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QueryDeployments queries for onboarding deployments
*/
func (a *Client) QueryDeployments(params *QueryDeploymentsParams, opts ...ClientOption) (*QueryDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryDeploymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "queryDeployments",
		Method:             "GET",
		PathPattern:        "/relocation/onboarding/deployment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QueryDeploymentsReader{formats: a.formats},
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
	success, ok := result.(*QueryDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryDeployments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDeploymentByID updates the selected onboarding deployment
*/
func (a *Client) UpdateDeploymentByID(params *UpdateDeploymentByIDParams, opts ...ClientOption) (*UpdateDeploymentByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeploymentByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDeploymentById",
		Method:             "PATCH",
		PathPattern:        "/relocation/onboarding/deployment/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeploymentByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateDeploymentByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDeploymentById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
