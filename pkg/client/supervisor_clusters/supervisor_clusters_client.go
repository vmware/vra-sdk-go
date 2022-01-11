// Code generated by go-swagger; DO NOT EDIT.

package supervisor_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new supervisor clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for supervisor clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetClusterUsingGET1(params *GetClusterUsingGET1Params, opts ...ClientOption) (*GetClusterUsingGET1OK, error)

	GetClusterUsingGET2(params *GetClusterUsingGET2Params, opts ...ClientOption) (*GetClusterUsingGET2OK, error)

	ListClustersOnEndpointUsingGET(params *ListClustersOnEndpointUsingGETParams, opts ...ClientOption) (*ListClustersOnEndpointUsingGETOK, error)

	ListUsingGET4(params *ListUsingGET4Params, opts ...ClientOption) (*ListUsingGET4OK, error)

	RegisterUsingPUT1(params *RegisterUsingPUT1Params, opts ...ClientOption) (*RegisterUsingPUT1OK, error)

	UnregisterUsingDELETE(params *UnregisterUsingDELETEParams, opts ...ClientOption) (*UnregisterUsingDELETEOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetClusterUsingGET1 finds a supervisor cluster by v sphere moref and v sphere endpoint id

  Retrieve a Supervisor Cluster by vSphere moref and id from the endpoint self link of the vSphere endpoint this cluster is associated to
*/
func (a *Client) GetClusterUsingGET1(params *GetClusterUsingGET1Params, opts ...ClientOption) (*GetClusterUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterUsingGET_1",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-clusters/endpoint/{endpointSelfLinkId}/cluster/{moref}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetClusterUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClusterUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClusterUsingGET2 finds a supervisor cluster by the id from document self link

  Retrieve a Supervisor Cluster by id from documentSelfLink
*/
func (a *Client) GetClusterUsingGET2(params *GetClusterUsingGET2Params, opts ...ClientOption) (*GetClusterUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClusterUsingGET_2",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-clusters/{selfLinkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetClusterUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClusterUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListClustersOnEndpointUsingGET gets all supervisor clusters on a v sphere endpoint

  Get all Supervisor Clusters on a vSphere endpoint by provided id from the endpoint self link
*/
func (a *Client) ListClustersOnEndpointUsingGET(params *ListClustersOnEndpointUsingGETParams, opts ...ClientOption) (*ListClustersOnEndpointUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersOnEndpointUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listClustersOnEndpointUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-clusters/endpoint/{endpointSelfLinkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListClustersOnEndpointUsingGETReader{formats: a.formats},
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
	success, ok := result.(*ListClustersOnEndpointUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listClustersOnEndpointUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUsingGET4 gets all managed supervisor clusters

  Get all managed Supervisor Clusters
*/
func (a *Client) ListUsingGET4(params *ListUsingGET4Params, opts ...ClientOption) (*ListUsingGET4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUsingGET4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listUsingGET_4",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUsingGET4Reader{formats: a.formats},
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
	success, ok := result.(*ListUsingGET4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUsingGET_4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RegisterUsingPUT1 makes a supervisor cluster a managed entity

  A valid document self link id shall be provided.
*/
func (a *Client) RegisterUsingPUT1(params *RegisterUsingPUT1Params, opts ...ClientOption) (*RegisterUsingPUT1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterUsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "registerUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/cmx/api/resources/supervisor-clusters/{clusterSelfLinkId}/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegisterUsingPUT1Reader{formats: a.formats},
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
	success, ok := result.(*RegisterUsingPUT1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for registerUsingPUT_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnregisterUsingDELETE makes a supervisor cluster an unmanaged entity

  A valid document self link id shall be provided.
*/
func (a *Client) UnregisterUsingDELETE(params *UnregisterUsingDELETEParams, opts ...ClientOption) (*UnregisterUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnregisterUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "unregisterUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/cmx/api/resources/supervisor-clusters/{clusterSelfLinkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnregisterUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*UnregisterUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unregisterUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
