// Code generated by go-swagger; DO NOT EDIT.

package p_k_s_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new p k s endpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p k s endpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateClusterUsingPOST(params *CreateClusterUsingPOSTParams, opts ...ClientOption) (*CreateClusterUsingPOSTOK, error)

	DestroyClusterUsingDELETE1(params *DestroyClusterUsingDELETE1Params, opts ...ClientOption) (*DestroyClusterUsingDELETE1OK, error)

	GetClustersUsingGET(params *GetClustersUsingGETParams, opts ...ClientOption) (*GetClustersUsingGETOK, error)

	GetPlansUsingGET(params *GetPlansUsingGETParams, opts ...ClientOption) (*GetPlansUsingGETOK, error)

	UpdateClusterUsingPUT(params *UpdateClusterUsingPUTParams, opts ...ClientOption) (*UpdateClusterUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateClusterUsingPOST creates a k8 s cluster

  Create a K8S cluster on PKS endpoint specified by endpoint id
*/
func (a *Client) CreateClusterUsingPOST(params *CreateClusterUsingPOSTParams, opts ...ClientOption) (*CreateClusterUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createClusterUsingPOST",
		Method:             "POST",
		PathPattern:        "/cmx/api/resources/pks/endpoints/{id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreateClusterUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createClusterUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DestroyClusterUsingDELETE1 destroys a k8 s cluster on a specific p k s endpoint

  Destroy and unregister a K8S cluster on PKS endpoint
*/
func (a *Client) DestroyClusterUsingDELETE1(params *DestroyClusterUsingDELETE1Params, opts ...ClientOption) (*DestroyClusterUsingDELETE1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDestroyClusterUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "destroyClusterUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/cmx/api/resources/pks/endpoints/{id}/clusters/{clusterId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DestroyClusterUsingDELETE1Reader{formats: a.formats},
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
	success, ok := result.(*DestroyClusterUsingDELETE1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for destroyClusterUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClustersUsingGET gets all k8 s clusters for a p k s endpoint

  Get all K8S clusters for a PKS endpoint by provided PKS endpoint id
*/
func (a *Client) GetClustersUsingGET(params *GetClustersUsingGETParams, opts ...ClientOption) (*GetClustersUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClustersUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClustersUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/pks/endpoints/{id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClustersUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetClustersUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getClustersUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPlansUsingGET gets supported plans for a p k s endpoint

  Get supported plans by providing PKS endpoint id
*/
func (a *Client) GetPlansUsingGET(params *GetPlansUsingGETParams, opts ...ClientOption) (*GetPlansUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPlansUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPlansUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/pks/endpoints/{id}/plans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPlansUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetPlansUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPlansUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateClusterUsingPUT updates a k8 s cluster on a specific endpoint id

  Update a K8S cluster on a PKS endpoint identified by id
*/
func (a *Client) UpdateClusterUsingPUT(params *UpdateClusterUsingPUTParams, opts ...ClientOption) (*UpdateClusterUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateClusterUsingPUT",
		Method:             "PUT",
		PathPattern:        "/cmx/api/resources/pks/endpoints/{id}/clusters/{clusterId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdateClusterUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateClusterUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
