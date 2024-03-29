// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new requests API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for requests API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ActionDeploymentRequestUsingPOST2(params *ActionDeploymentRequestUsingPOST2Params, opts ...ClientOption) (*ActionDeploymentRequestUsingPOST2OK, error)

	GetDeploymentRequestsUsingGET2(params *GetDeploymentRequestsUsingGET2Params, opts ...ClientOption) (*GetDeploymentRequestsUsingGET2OK, error)

	GetEventLogsContentUsingGET2(params *GetEventLogsContentUsingGET2Params, opts ...ClientOption) (*GetEventLogsContentUsingGET2OK, error)

	GetEventLogsUsingGET2(params *GetEventLogsUsingGET2Params, opts ...ClientOption) (*GetEventLogsUsingGET2OK, error)

	GetRequestEventsUsingGET2(params *GetRequestEventsUsingGET2Params, opts ...ClientOption) (*GetRequestEventsUsingGET2OK, error)

	GetRequestUsingGET2(params *GetRequestUsingGET2Params, opts ...ClientOption) (*GetRequestUsingGET2OK, error)

	ListResourceRequestsUsingGET2(params *ListResourceRequestsUsingGET2Params, opts ...ClientOption) (*ListResourceRequestsUsingGET2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ActionDeploymentRequestUsingPOST2 submits action on requests allowable values cancel dismiss

Cancel can be submitted on In-progress requests and Dismiss can be submitted on Failed requests.
*/
func (a *Client) ActionDeploymentRequestUsingPOST2(params *ActionDeploymentRequestUsingPOST2Params, opts ...ClientOption) (*ActionDeploymentRequestUsingPOST2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActionDeploymentRequestUsingPOST2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "actionDeploymentRequestUsingPOST_2",
		Method:             "POST",
		PathPattern:        "/deployment/api/requests/{requestId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActionDeploymentRequestUsingPOST2Reader{formats: a.formats},
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
	success, ok := result.(*ActionDeploymentRequestUsingPOST2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for actionDeploymentRequestUsingPOST_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeploymentRequestsUsingGET2 fetches deployment requests

Returns the requests for the deployment.
*/
func (a *Client) GetDeploymentRequestsUsingGET2(params *GetDeploymentRequestsUsingGET2Params, opts ...ClientOption) (*GetDeploymentRequestsUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentRequestsUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeploymentRequestsUsingGET_2",
		Method:             "GET",
		PathPattern:        "/deployment/api/deployments/{deploymentId}/requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentRequestsUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetDeploymentRequestsUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeploymentRequestsUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEventLogsContentUsingGET2 fetches event logs content as a file

Returns the log file for an event.
*/
func (a *Client) GetEventLogsContentUsingGET2(params *GetEventLogsContentUsingGET2Params, opts ...ClientOption) (*GetEventLogsContentUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventLogsContentUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEventLogsContentUsingGET_2",
		Method:             "GET",
		PathPattern:        "/deployment/api/requests/{requestId}/events/{eventId}/logs/download",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventLogsContentUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetEventLogsContentUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEventLogsContentUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEventLogsUsingGET2 fetches event logs

Returns the logs for an event.
*/
func (a *Client) GetEventLogsUsingGET2(params *GetEventLogsUsingGET2Params, opts ...ClientOption) (*GetEventLogsUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventLogsUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEventLogsUsingGET_2",
		Method:             "GET",
		PathPattern:        "/deployment/api/requests/{requestId}/events/{eventId}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventLogsUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetEventLogsUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEventLogsUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRequestEventsUsingGET2 fetches request events

Returns all the events for a request.
*/
func (a *Client) GetRequestEventsUsingGET2(params *GetRequestEventsUsingGET2Params, opts ...ClientOption) (*GetRequestEventsUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRequestEventsUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRequestEventsUsingGET_2",
		Method:             "GET",
		PathPattern:        "/deployment/api/requests/{requestId}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRequestEventsUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetRequestEventsUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRequestEventsUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetRequestUsingGET2 gets the request

Returns the request with the given ID.
*/
func (a *Client) GetRequestUsingGET2(params *GetRequestUsingGET2Params, opts ...ClientOption) (*GetRequestUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRequestUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRequestUsingGET_2",
		Method:             "GET",
		PathPattern:        "/deployment/api/requests/{requestId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRequestUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetRequestUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRequestUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListResourceRequestsUsingGET2 gets all requests for a resource
*/
func (a *Client) ListResourceRequestsUsingGET2(params *ListResourceRequestsUsingGET2Params, opts ...ClientOption) (*ListResourceRequestsUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListResourceRequestsUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listResourceRequestsUsingGET_2",
		Method:             "GET",
		PathPattern:        "/deployment/api/resources/{resourceId}/requests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListResourceRequestsUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*ListResourceRequestsUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listResourceRequestsUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
