// Code generated by go-swagger; DO NOT EDIT.

package flavor_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new flavor profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for flavor profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFlavorProfile(params *CreateFlavorProfileParams, opts ...ClientOption) (*CreateFlavorProfileCreated, error)

	DeleteFlavorProfile(params *DeleteFlavorProfileParams, opts ...ClientOption) (*DeleteFlavorProfileNoContent, error)

	GetFlavorProfile(params *GetFlavorProfileParams, opts ...ClientOption) (*GetFlavorProfileOK, error)

	GetFlavorProfiles(params *GetFlavorProfilesParams, opts ...ClientOption) (*GetFlavorProfilesOK, error)

	UpdateFlavorProfile(params *UpdateFlavorProfileParams, opts ...ClientOption) (*UpdateFlavorProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateFlavorProfile creates flavor profile

Create flavor profile
*/
func (a *Client) CreateFlavorProfile(params *CreateFlavorProfileParams, opts ...ClientOption) (*CreateFlavorProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFlavorProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFlavorProfile",
		Method:             "POST",
		PathPattern:        "/iaas/api/flavor-profiles",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFlavorProfileReader{formats: a.formats},
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
	success, ok := result.(*CreateFlavorProfileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createFlavorProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteFlavorProfile deletes flavor profile

Delete flavor profile with a given id
*/
func (a *Client) DeleteFlavorProfile(params *DeleteFlavorProfileParams, opts ...ClientOption) (*DeleteFlavorProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFlavorProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteFlavorProfile",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/flavor-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFlavorProfileReader{formats: a.formats},
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
	success, ok := result.(*DeleteFlavorProfileNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteFlavorProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFlavorProfile gets flavor profile

Get flavor profile with a given id
*/
func (a *Client) GetFlavorProfile(params *GetFlavorProfileParams, opts ...ClientOption) (*GetFlavorProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFlavorProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFlavorProfile",
		Method:             "GET",
		PathPattern:        "/iaas/api/flavor-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFlavorProfileReader{formats: a.formats},
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
	success, ok := result.(*GetFlavorProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFlavorProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFlavorProfiles gets flavor profile

Get all flavor profile
*/
func (a *Client) GetFlavorProfiles(params *GetFlavorProfilesParams, opts ...ClientOption) (*GetFlavorProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFlavorProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFlavorProfiles",
		Method:             "GET",
		PathPattern:        "/iaas/api/flavor-profiles",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFlavorProfilesReader{formats: a.formats},
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
	success, ok := result.(*GetFlavorProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFlavorProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateFlavorProfile updates flavor profile

Update flavor profile
*/
func (a *Client) UpdateFlavorProfile(params *UpdateFlavorProfileParams, opts ...ClientOption) (*UpdateFlavorProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateFlavorProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateFlavorProfile",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/flavor-profiles/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateFlavorProfileReader{formats: a.formats},
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
	success, ok := result.(*UpdateFlavorProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateFlavorProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
