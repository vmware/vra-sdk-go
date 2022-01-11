// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePolicyUsingDELETE5(params *DeletePolicyUsingDELETE5Params, opts ...ClientOption) (*DeletePolicyUsingDELETE5NoContent, error)

	DryRunPolicyUsingPOST2(params *DryRunPolicyUsingPOST2Params, opts ...ClientOption) (*DryRunPolicyUsingPOST2OK, error)

	GetPoliciesUsingGET5(params *GetPoliciesUsingGET5Params, opts ...ClientOption) (*GetPoliciesUsingGET5OK, error)

	GetPolicyUsingGET5(params *GetPolicyUsingGET5Params, opts ...ClientOption) (*GetPolicyUsingGET5OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePolicyUsingDELETE5 deletes a policy

  Delete a specified policy corresponding to its unique id.
*/
func (a *Client) DeletePolicyUsingDELETE5(params *DeletePolicyUsingDELETE5Params, opts ...ClientOption) (*DeletePolicyUsingDELETE5NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePolicyUsingDELETE5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePolicyUsingDELETE_5",
		Method:             "DELETE",
		PathPattern:        "/policy/api/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePolicyUsingDELETE5Reader{formats: a.formats},
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
	success, ok := result.(*DeletePolicyUsingDELETE5NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePolicyUsingDELETE_5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DryRunPolicyUsingPOST2 triggers a policy dry run

  Dry-run an existing policy to rehearse actual policy effect on application.
*/
func (a *Client) DryRunPolicyUsingPOST2(params *DryRunPolicyUsingPOST2Params, opts ...ClientOption) (*DryRunPolicyUsingPOST2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDryRunPolicyUsingPOST2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "dryRunPolicyUsingPOST_2",
		Method:             "POST",
		PathPattern:        "/policy/api/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DryRunPolicyUsingPOST2Reader{formats: a.formats},
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
	success, ok := result.(*DryRunPolicyUsingPOST2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for dryRunPolicyUsingPOST_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPoliciesUsingGET5 returns a paginated list of policies

  Find all the policies associated with current org.
*/
func (a *Client) GetPoliciesUsingGET5(params *GetPoliciesUsingGET5Params, opts ...ClientOption) (*GetPoliciesUsingGET5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoliciesUsingGET5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPoliciesUsingGET_5",
		Method:             "GET",
		PathPattern:        "/policy/api/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoliciesUsingGET5Reader{formats: a.formats},
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
	success, ok := result.(*GetPoliciesUsingGET5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPoliciesUsingGET_5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPolicyUsingGET5 returns a specified policy

  Find a specific policy based on the input policy id.
*/
func (a *Client) GetPolicyUsingGET5(params *GetPolicyUsingGET5Params, opts ...ClientOption) (*GetPolicyUsingGET5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyUsingGET5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPolicyUsingGET_5",
		Method:             "GET",
		PathPattern:        "/policy/api/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPolicyUsingGET5Reader{formats: a.formats},
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
	success, ok := result.(*GetPolicyUsingGET5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPolicyUsingGET_5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
