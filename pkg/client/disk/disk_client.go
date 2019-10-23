// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new disk API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for disk API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AttachMachineDisk attaches machine disk

Attach a disk to a machine.
*/
func (a *Client) AttachMachineDisk(params *AttachMachineDiskParams) (*AttachMachineDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachMachineDiskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attachMachineDisk",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/disks",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AttachMachineDiskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AttachMachineDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for attachMachineDisk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateBlockDevice creates block device

Following disk custom properties can be passed while creating a block device:

    1. dataStore: Defines name of the datastore in which the disk has to be provisioned.
    2. storagePolicy: Defines name of the storage policy in which the disk has to be provisioned. If name of the datastore is specified in the custom properties then, datastore takes precedence.
    3. provisioningType: Defines the type of provisioning. For eg. thick/thin.
*/
func (a *Client) CreateBlockDevice(params *CreateBlockDeviceParams) (*CreateBlockDeviceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlockDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBlockDevice",
		Method:             "POST",
		PathPattern:        "/iaas/api/block-devices",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBlockDeviceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBlockDeviceAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlockDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteBlockDevice deletes a block device

Delete a BlockDevice
*/
func (a *Client) DeleteBlockDevice(params *DeleteBlockDeviceParams) (*DeleteBlockDeviceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlockDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBlockDevice",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/block-devices/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBlockDeviceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBlockDeviceAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBlockDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMachineDisk deletes machine disk

Remove a disk from a given machine.
*/
func (a *Client) DeleteMachineDisk(params *DeleteMachineDiskParams) (*DeleteMachineDiskAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMachineDiskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMachineDisk",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/machines/{id}/disks/{id1}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMachineDiskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteMachineDiskAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMachineDisk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlockDevice gets block device

Get a single BlockDevice
*/
func (a *Client) GetBlockDevice(params *GetBlockDeviceParams) (*GetBlockDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlockDeviceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlockDevice",
		Method:             "GET",
		PathPattern:        "/iaas/api/block-devices/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBlockDeviceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBlockDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlockDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlockDevices gets block devices

Get all BlockDevices
*/
func (a *Client) GetBlockDevices(params *GetBlockDevicesParams) (*GetBlockDevicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlockDevicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlockDevices",
		Method:             "GET",
		PathPattern:        "/iaas/api/block-devices",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBlockDevicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBlockDevicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlockDevices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMachineDisk gets a machine disk

Get disk with a given id for specific machine
*/
func (a *Client) GetMachineDisk(params *GetMachineDiskParams) (*GetMachineDiskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMachineDiskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMachineDisk",
		Method:             "GET",
		PathPattern:        "/iaas/api/machines/{id}/disks/{id1}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMachineDiskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMachineDiskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMachineDisk: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMachineDisks gets machine disks

Get all machine disks
*/
func (a *Client) GetMachineDisks(params *GetMachineDisksParams) (*GetMachineDisksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMachineDisksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMachineDisks",
		Method:             "GET",
		PathPattern:        "/iaas/api/machines/{id}/disks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMachineDisksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMachineDisksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMachineDisks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
