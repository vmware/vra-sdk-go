// Code generated by go-swagger; DO NOT EDIT.

package vsphere_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v sphere endpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v sphere endpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFullVirtualMachineClassesUsingGET(params *GetFullVirtualMachineClassesUsingGETParams, opts ...ClientOption) (*GetFullVirtualMachineClassesUsingGETOK, error)

	GetStorageClassesUsingGET(params *GetStorageClassesUsingGETParams, opts ...ClientOption) (*GetStorageClassesUsingGETOK, error)

	GetTanzuKubernetesReleasesUsingGET(params *GetTanzuKubernetesReleasesUsingGETParams, opts ...ClientOption) (*GetTanzuKubernetesReleasesUsingGETOK, error)

	GetVirtualMachineClassesUsingGET(params *GetVirtualMachineClassesUsingGETParams, opts ...ClientOption) (*GetVirtualMachineClassesUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetFullVirtualMachineClassesUsingGET gets full information about all virtual machine classes for a v sphere endpoint

Get all virtual machine classes defined in all managed supervisor clusters in a particular vSphere instance. vSphere instance is identified by endpoint SelfLink Id
*/
func (a *Client) GetFullVirtualMachineClassesUsingGET(params *GetFullVirtualMachineClassesUsingGETParams, opts ...ClientOption) (*GetFullVirtualMachineClassesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFullVirtualMachineClassesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFullVirtualMachineClassesUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/vsphere/endpoints/{endpointSelfLinkId}/virtual-machine-classes-described",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFullVirtualMachineClassesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetFullVirtualMachineClassesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFullVirtualMachineClassesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetStorageClassesUsingGET gets all storage classes identifiers for a v sphere endpoint

Get all storage classes defined in all managed supervisor clusters in a particular vSphere instance. vSphere instance is identified by endpoint SelfLink Id
*/
func (a *Client) GetStorageClassesUsingGET(params *GetStorageClassesUsingGETParams, opts ...ClientOption) (*GetStorageClassesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStorageClassesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getStorageClassesUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/vsphere/endpoints/{endpointSelfLinkId}/storage-classes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStorageClassesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetStorageClassesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getStorageClassesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTanzuKubernetesReleasesUsingGET gets all tanzu kubernetes releases identifiers for a v sphere endpoint

Get all tanzu kubernetes releases supported in all managed supervisor clusters in a particular vSphere instance. vSphere instance is identified by endpoint SelfLink Id
*/
func (a *Client) GetTanzuKubernetesReleasesUsingGET(params *GetTanzuKubernetesReleasesUsingGETParams, opts ...ClientOption) (*GetTanzuKubernetesReleasesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTanzuKubernetesReleasesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTanzuKubernetesReleasesUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/vsphere/endpoints/{endpointSelfLinkId}/tanzu-kubernetes-releases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTanzuKubernetesReleasesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetTanzuKubernetesReleasesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTanzuKubernetesReleasesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVirtualMachineClassesUsingGET gets all virtual machine classes names for a v sphere endpoint

Get all virtual machine classes defined in all managed supervisor clusters in a particular vSphere instance. vSphere instance is identified by endpoint SelfLink Id
*/
func (a *Client) GetVirtualMachineClassesUsingGET(params *GetVirtualMachineClassesUsingGETParams, opts ...ClientOption) (*GetVirtualMachineClassesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVirtualMachineClassesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getVirtualMachineClassesUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/vsphere/endpoints/{endpointSelfLinkId}/virtual-machine-classes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVirtualMachineClassesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetVirtualMachineClassesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVirtualMachineClassesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
