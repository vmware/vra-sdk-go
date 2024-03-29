// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deployment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDeployment(params *CreateDeploymentParams, opts ...ClientOption) (*CreateDeploymentCreated, error)

	DeleteDeployment(params *DeleteDeploymentParams, opts ...ClientOption) (*DeleteDeploymentAccepted, error)

	GetDeployments(params *GetDeploymentsParams, opts ...ClientOption) (*GetDeploymentsOK, error)

	GetSingleDeployment(params *GetSingleDeploymentParams, opts ...ClientOption) (*GetSingleDeploymentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDeployment creates deployment

Create a new Deployment.
*/
func (a *Client) CreateDeployment(params *CreateDeploymentParams, opts ...ClientOption) (*CreateDeploymentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeployment",
		Method:             "POST",
		PathPattern:        "/iaas/api/deployments",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeploymentReader{formats: a.formats},
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
	success, ok := result.(*CreateDeploymentCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDeployment deletes a deployment

Delete a deployment.
*/
func (a *Client) DeleteDeployment(params *DeleteDeploymentParams, opts ...ClientOption) (*DeleteDeploymentAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDeployment",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/deployments/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDeploymentReader{formats: a.formats},
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
	success, ok := result.(*DeleteDeploymentAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDeployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeployments gets deployments

Get all deployments.
*/
func (a *Client) GetDeployments(params *GetDeploymentsParams, opts ...ClientOption) (*GetDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeployments",
		Method:             "GET",
		PathPattern:        "/iaas/api/deployments",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentsReader{formats: a.formats},
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
	success, ok := result.(*GetDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeployments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSingleDeployment gets a single deployment

Get a single deployment.
*/
func (a *Client) GetSingleDeployment(params *GetSingleDeploymentParams, opts ...ClientOption) (*GetSingleDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSingleDeploymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSingleDeployment",
		Method:             "GET",
		PathPattern:        "/iaas/api/deployments/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSingleDeploymentReader{formats: a.formats},
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
	success, ok := result.(*GetSingleDeploymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSingleDeployment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
