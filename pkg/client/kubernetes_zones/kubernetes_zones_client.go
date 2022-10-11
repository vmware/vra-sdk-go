// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_zones

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new kubernetes zones API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kubernetes zones API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateZoneUsingPOST(params *CreateZoneUsingPOSTParams, opts ...ClientOption) (*CreateZoneUsingPOSTOK, error)

	DeleteZoneUsingDELETE(params *DeleteZoneUsingDELETEParams, opts ...ClientOption) (*DeleteZoneUsingDELETEOK, error)

	GetZoneUsingGET(params *GetZoneUsingGETParams, opts ...ClientOption) (*GetZoneUsingGETOK, error)

	ListZonesUsingGET(params *ListZonesUsingGETParams, opts ...ClientOption) (*ListZonesUsingGETOK, error)

	UpdateZoneProjectsUsingPUT(params *UpdateZoneProjectsUsingPUTParams, opts ...ClientOption) (*UpdateZoneProjectsUsingPUTOK, error)

	UpdateZoneUsingPUT(params *UpdateZoneUsingPUTParams, opts ...ClientOption) (*UpdateZoneUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateZoneUsingPOST creates a k8 s zone

Create a K8S Zone entity
*/
func (a *Client) CreateZoneUsingPOST(params *CreateZoneUsingPOSTParams, opts ...ClientOption) (*CreateZoneUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateZoneUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createZoneUsingPOST",
		Method:             "POST",
		PathPattern:        "/cmx/api/resources/k8s-zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateZoneUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreateZoneUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createZoneUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteZoneUsingDELETE deletes a k8 s zone

Remove a K8S Zone
*/
func (a *Client) DeleteZoneUsingDELETE(params *DeleteZoneUsingDELETEParams, opts ...ClientOption) (*DeleteZoneUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteZoneUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteZoneUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/cmx/api/resources/k8s-zones/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteZoneUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeleteZoneUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteZoneUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetZoneUsingGET gets a k8 s zone

Get a K8S Zone by Id
*/
func (a *Client) GetZoneUsingGET(params *GetZoneUsingGETParams, opts ...ClientOption) (*GetZoneUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZoneUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getZoneUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/k8s-zones/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZoneUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetZoneUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getZoneUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListZonesUsingGET gets all k8 s zones

List of all K8S Zones
*/
func (a *Client) ListZonesUsingGET(params *ListZonesUsingGETParams, opts ...ClientOption) (*ListZonesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListZonesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listZonesUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/k8s-zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListZonesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*ListZonesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listZonesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateZoneProjectsUsingPUT updates a k8 s zone project assignments

Assignment of projects to K8S Zone
*/
func (a *Client) UpdateZoneProjectsUsingPUT(params *UpdateZoneProjectsUsingPUTParams, opts ...ClientOption) (*UpdateZoneProjectsUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateZoneProjectsUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateZoneProjectsUsingPUT",
		Method:             "PUT",
		PathPattern:        "/cmx/api/resources/k8s-zones/{id}/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateZoneProjectsUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdateZoneProjectsUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateZoneProjectsUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateZoneUsingPUT updates a k8 s zone

Modify a K8S Zone
*/
func (a *Client) UpdateZoneUsingPUT(params *UpdateZoneUsingPUTParams, opts ...ClientOption) (*UpdateZoneUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateZoneUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateZoneUsingPUT",
		Method:             "PUT",
		PathPattern:        "/cmx/api/resources/k8s-zones/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateZoneUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdateZoneUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateZoneUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}