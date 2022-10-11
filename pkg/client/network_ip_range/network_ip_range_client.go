// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new network ip range API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network ip range API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInternalNetworkIPRange(params *CreateInternalNetworkIPRangeParams, opts ...ClientOption) (*CreateInternalNetworkIPRangeCreated, error)

	DeleteInternalNetworkIPRange(params *DeleteInternalNetworkIPRangeParams, opts ...ClientOption) (*DeleteInternalNetworkIPRangeNoContent, error)

	GetExternalIPBlock(params *GetExternalIPBlockParams, opts ...ClientOption) (*GetExternalIPBlockOK, error)

	GetExternalIPBlocks(params *GetExternalIPBlocksParams, opts ...ClientOption) (*GetExternalIPBlocksOK, error)

	GetExternalNetworkIPRange(params *GetExternalNetworkIPRangeParams, opts ...ClientOption) (*GetExternalNetworkIPRangeOK, error)

	GetExternalNetworkIPRanges(params *GetExternalNetworkIPRangesParams, opts ...ClientOption) (*GetExternalNetworkIPRangesOK, error)

	GetInternalNetworkIPRange(params *GetInternalNetworkIPRangeParams, opts ...ClientOption) (*GetInternalNetworkIPRangeOK, error)

	GetInternalNetworkIPRanges(params *GetInternalNetworkIPRangesParams, opts ...ClientOption) (*GetInternalNetworkIPRangesOK, error)

	UpdateExternalNetworkIPRange(params *UpdateExternalNetworkIPRangeParams, opts ...ClientOption) (*UpdateExternalNetworkIPRangeOK, error)

	UpdateInternalNetworkIPRange(params *UpdateInternalNetworkIPRangeParams, opts ...ClientOption) (*UpdateInternalNetworkIPRangeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateInternalNetworkIPRange creates internal network IP range

Creates an internal network IP range.
*/
func (a *Client) CreateInternalNetworkIPRange(params *CreateInternalNetworkIPRangeParams, opts ...ClientOption) (*CreateInternalNetworkIPRangeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInternalNetworkIPRangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createInternalNetworkIPRange",
		Method:             "POST",
		PathPattern:        "/iaas/api/network-ip-ranges",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInternalNetworkIPRangeReader{formats: a.formats},
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
	success, ok := result.(*CreateInternalNetworkIPRangeCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createInternalNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteInternalNetworkIPRange deletes internal network IP range

Delete internal network IP range with a given id
*/
func (a *Client) DeleteInternalNetworkIPRange(params *DeleteInternalNetworkIPRangeParams, opts ...ClientOption) (*DeleteInternalNetworkIPRangeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInternalNetworkIPRangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInternalNetworkIPRange",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInternalNetworkIPRangeReader{formats: a.formats},
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
	success, ok := result.(*DeleteInternalNetworkIPRangeNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInternalNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExternalIPBlock gets specific external IP block by id

An external IP block is network coming from external IPAM provider that can be used to create subnetworks inside it
*/
func (a *Client) GetExternalIPBlock(params *GetExternalIPBlockParams, opts ...ClientOption) (*GetExternalIPBlockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalIPBlockParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExternalIpBlock",
		Method:             "GET",
		PathPattern:        "/iaas/api/external-ip-blocks/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalIPBlockReader{formats: a.formats},
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
	success, ok := result.(*GetExternalIPBlockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExternalIpBlock: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExternalIPBlocks gets all external IP blocks

An external IP block is network coming from external IPAM provider that can be used to create subnetworks inside it
*/
func (a *Client) GetExternalIPBlocks(params *GetExternalIPBlocksParams, opts ...ClientOption) (*GetExternalIPBlocksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalIPBlocksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExternalIpBlocks",
		Method:             "GET",
		PathPattern:        "/iaas/api/external-ip-blocks",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalIPBlocksReader{formats: a.formats},
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
	success, ok := result.(*GetExternalIPBlocksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExternalIpBlocks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExternalNetworkIPRange gets external IP a m network IP range

Get external IPAM network IP range with a given id
*/
func (a *Client) GetExternalNetworkIPRange(params *GetExternalNetworkIPRangeParams, opts ...ClientOption) (*GetExternalNetworkIPRangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalNetworkIPRangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExternalNetworkIPRange",
		Method:             "GET",
		PathPattern:        "/iaas/api/external-network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalNetworkIPRangeReader{formats: a.formats},
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
	success, ok := result.(*GetExternalNetworkIPRangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExternalNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExternalNetworkIPRanges gets external IP a m network IP ranges

Get all external IPAM network IP ranges
*/
func (a *Client) GetExternalNetworkIPRanges(params *GetExternalNetworkIPRangesParams, opts ...ClientOption) (*GetExternalNetworkIPRangesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalNetworkIPRangesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExternalNetworkIPRanges",
		Method:             "GET",
		PathPattern:        "/iaas/api/external-network-ip-ranges",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExternalNetworkIPRangesReader{formats: a.formats},
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
	success, ok := result.(*GetExternalNetworkIPRangesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExternalNetworkIPRanges: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInternalNetworkIPRange gets internal IP a m network IP range

Get internal IPAM network IP range with a given id
*/
func (a *Client) GetInternalNetworkIPRange(params *GetInternalNetworkIPRangeParams, opts ...ClientOption) (*GetInternalNetworkIPRangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternalNetworkIPRangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInternalNetworkIPRange",
		Method:             "GET",
		PathPattern:        "/iaas/api/network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternalNetworkIPRangeReader{formats: a.formats},
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
	success, ok := result.(*GetInternalNetworkIPRangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInternalNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInternalNetworkIPRanges gets internal IP a m network IP ranges

Get all internal IPAM network IP ranges
*/
func (a *Client) GetInternalNetworkIPRanges(params *GetInternalNetworkIPRangesParams, opts ...ClientOption) (*GetInternalNetworkIPRangesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternalNetworkIPRangesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInternalNetworkIPRanges",
		Method:             "GET",
		PathPattern:        "/iaas/api/network-ip-ranges",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternalNetworkIPRangesReader{formats: a.formats},
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
	success, ok := result.(*GetInternalNetworkIPRangesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInternalNetworkIPRanges: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateExternalNetworkIPRange updates external IP a m network IP range

Assign the external IPAM network IP range to a different network and/or change the tags of the external IPAM network IP range.
*/
func (a *Client) UpdateExternalNetworkIPRange(params *UpdateExternalNetworkIPRangeParams, opts ...ClientOption) (*UpdateExternalNetworkIPRangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExternalNetworkIPRangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateExternalNetworkIPRange",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/external-network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExternalNetworkIPRangeReader{formats: a.formats},
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
	success, ok := result.(*UpdateExternalNetworkIPRangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateExternalNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateInternalNetworkIPRange updates internal network IP range

Update internal network IP range.
*/
func (a *Client) UpdateInternalNetworkIPRange(params *UpdateInternalNetworkIPRangeParams, opts ...ClientOption) (*UpdateInternalNetworkIPRangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateInternalNetworkIPRangeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateInternalNetworkIPRange",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateInternalNetworkIPRangeReader{formats: a.formats},
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
	success, ok := result.(*UpdateInternalNetworkIPRangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateInternalNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
