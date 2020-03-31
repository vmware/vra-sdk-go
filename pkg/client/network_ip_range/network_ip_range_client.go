// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new network ip range API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network ip range API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateNetworkIPRange creates network IP range

Creates a network IP range.
*/
func (a *Client) CreateNetworkIPRange(params *CreateNetworkIPRangeParams) (*CreateNetworkIPRangeCreated, *CreateNetworkIPRangeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkIPRangeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNetworkIPRange",
		Method:             "POST",
		PathPattern:        "/iaas/api/network-ip-ranges",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNetworkIPRangeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateNetworkIPRangeCreated:
		return value, nil, nil
	case *CreateNetworkIPRangeAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for network_ip_range: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteNetworkIPRange deletes network IP range

Delete network IP range with a given id
*/
func (a *Client) DeleteNetworkIPRange(params *DeleteNetworkIPRangeParams) (*DeleteNetworkIPRangeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkIPRangeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNetworkIPRange",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNetworkIPRangeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworkIPRangeNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNetworkIPRange gets network IP range

Get network IP range with a given id
*/
func (a *Client) GetNetworkIPRange(params *GetNetworkIPRangeParams) (*GetNetworkIPRangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkIPRangeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkIPRange",
		Method:             "GET",
		PathPattern:        "/iaas/api/network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkIPRangeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkIPRangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNetworkIPRanges gets network IP ranges

Get all network IP ranges
*/
func (a *Client) GetNetworkIPRanges(params *GetNetworkIPRangesParams) (*GetNetworkIPRangesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkIPRangesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkIPRanges",
		Method:             "GET",
		PathPattern:        "/iaas/api/network-ip-ranges",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkIPRangesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkIPRangesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkIPRanges: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateNetworkIPRange updates network IP range

Update network IP range.
*/
func (a *Client) UpdateNetworkIPRange(params *UpdateNetworkIPRangeParams) (*UpdateNetworkIPRangeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkIPRangeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkIPRange",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/network-ip-ranges/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkIPRangeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkIPRangeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkIPRange: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
