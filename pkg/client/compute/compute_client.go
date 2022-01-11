// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new compute API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for compute API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateMachine(params *CreateMachineParams, opts ...ClientOption) (*CreateMachineAccepted, error)

	CreateMachineSnapshot(params *CreateMachineSnapshotParams, opts ...ClientOption) (*CreateMachineSnapshotAccepted, error)

	DeleteMachine(params *DeleteMachineParams, opts ...ClientOption) (*DeleteMachineAccepted, error)

	DeleteMachineSnapshot(params *DeleteMachineSnapshotParams, opts ...ClientOption) (*DeleteMachineSnapshotAccepted, error)

	GetMachine(params *GetMachineParams, opts ...ClientOption) (*GetMachineOK, error)

	GetMachineSnapshot(params *GetMachineSnapshotParams, opts ...ClientOption) (*GetMachineSnapshotOK, error)

	GetMachineSnapshots(params *GetMachineSnapshotsParams, opts ...ClientOption) (*GetMachineSnapshotsOK, error)

	GetMachines(params *GetMachinesParams, opts ...ClientOption) (*GetMachinesOK, error)

	PowerOffMachine(params *PowerOffMachineParams, opts ...ClientOption) (*PowerOffMachineAccepted, error)

	PowerOnMachine(params *PowerOnMachineParams, opts ...ClientOption) (*PowerOnMachineAccepted, error)

	RebootMachine(params *RebootMachineParams, opts ...ClientOption) (*RebootMachineAccepted, error)

	ResetMachine(params *ResetMachineParams, opts ...ClientOption) (*ResetMachineAccepted, error)

	ResizeMachine(params *ResizeMachineParams, opts ...ClientOption) (*ResizeMachineAccepted, error)

	RestartMachine(params *RestartMachineParams, opts ...ClientOption) (*RestartMachineAccepted, error)

	RevertMachineSnapshot(params *RevertMachineSnapshotParams, opts ...ClientOption) (*RevertMachineSnapshotAccepted, error)

	ShutdownMachine(params *ShutdownMachineParams, opts ...ClientOption) (*ShutdownMachineAccepted, error)

	SuspendMachine(params *SuspendMachineParams, opts ...ClientOption) (*SuspendMachineAccepted, error)

	UpdateMachine(params *UpdateMachineParams, opts ...ClientOption) (*UpdateMachineOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateMachine creates machine

  Create machine
*/
func (a *Client) CreateMachine(params *CreateMachineParams, opts ...ClientOption) (*CreateMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMachineReader{formats: a.formats},
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
	success, ok := result.(*CreateMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateMachineSnapshot creates snapshot operation for machine

  Second day create snapshot operation for machine
*/
func (a *Client) CreateMachineSnapshot(params *CreateMachineSnapshotParams, opts ...ClientOption) (*CreateMachineSnapshotAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMachineSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createMachineSnapshot",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/snapshots",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMachineSnapshotReader{formats: a.formats},
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
	success, ok := result.(*CreateMachineSnapshotAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createMachineSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteMachine deletes machine

  Delete Machine with a given id
*/
func (a *Client) DeleteMachine(params *DeleteMachineParams, opts ...ClientOption) (*DeleteMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMachine",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/machines/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMachineReader{formats: a.formats},
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
	success, ok := result.(*DeleteMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteMachineSnapshot deletes snapshot operation for machine

  Second day delete snapshot operation for machine
*/
func (a *Client) DeleteMachineSnapshot(params *DeleteMachineSnapshotParams, opts ...ClientOption) (*DeleteMachineSnapshotAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMachineSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteMachineSnapshot",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/machines/{id}/snapshots/{snapshotId}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMachineSnapshotReader{formats: a.formats},
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
	success, ok := result.(*DeleteMachineSnapshotAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMachineSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMachine gets machine

  Get machine with a given id
*/
func (a *Client) GetMachine(params *GetMachineParams, opts ...ClientOption) (*GetMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMachine",
		Method:             "GET",
		PathPattern:        "/iaas/api/machines/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMachineReader{formats: a.formats},
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
	success, ok := result.(*GetMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMachineSnapshot gets machine snapshot

  Get snapshot with a given id for specific machine
*/
func (a *Client) GetMachineSnapshot(params *GetMachineSnapshotParams, opts ...ClientOption) (*GetMachineSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMachineSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMachineSnapshot",
		Method:             "GET",
		PathPattern:        "/iaas/api/machines/{id}/snapshots/{snapshotId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMachineSnapshotReader{formats: a.formats},
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
	success, ok := result.(*GetMachineSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMachineSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMachineSnapshots gets machine snapshots information

  Get machine snapshots information
*/
func (a *Client) GetMachineSnapshots(params *GetMachineSnapshotsParams, opts ...ClientOption) (*GetMachineSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMachineSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMachineSnapshots",
		Method:             "GET",
		PathPattern:        "/iaas/api/machines/{id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMachineSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*GetMachineSnapshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMachineSnapshots: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMachines gets machines

  Get all machines
*/
func (a *Client) GetMachines(params *GetMachinesParams, opts ...ClientOption) (*GetMachinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMachinesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMachines",
		Method:             "GET",
		PathPattern:        "/iaas/api/machines",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMachinesReader{formats: a.formats},
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
	success, ok := result.(*GetMachinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMachines: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PowerOffMachine powers off operation for machine

  Second day power-off operation for machine
*/
func (a *Client) PowerOffMachine(params *PowerOffMachineParams, opts ...ClientOption) (*PowerOffMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPowerOffMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "powerOffMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/power-off",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PowerOffMachineReader{formats: a.formats},
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
	success, ok := result.(*PowerOffMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for powerOffMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PowerOnMachine powers on operation for machine

  Second day power-on operation for machine
*/
func (a *Client) PowerOnMachine(params *PowerOnMachineParams, opts ...ClientOption) (*PowerOnMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPowerOnMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "powerOnMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/power-on",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PowerOnMachineReader{formats: a.formats},
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
	success, ok := result.(*PowerOnMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for powerOnMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RebootMachine reboots operation for machine

  Second day reboot operation for machine
*/
func (a *Client) RebootMachine(params *RebootMachineParams, opts ...ClientOption) (*RebootMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "rebootMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/reboot",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RebootMachineReader{formats: a.formats},
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
	success, ok := result.(*RebootMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rebootMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResetMachine resets operation for machine

  Second day reset operation for machine
*/
func (a *Client) ResetMachine(params *ResetMachineParams, opts ...ClientOption) (*ResetMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/reset",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResetMachineReader{formats: a.formats},
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
	success, ok := result.(*ResetMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResizeMachine resizes operation for machine

  Second day resize operation for machine
*/
func (a *Client) ResizeMachine(params *ResizeMachineParams, opts ...ClientOption) (*ResizeMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResizeMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resizeMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/resize",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResizeMachineReader{formats: a.formats},
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
	success, ok := result.(*ResizeMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resizeMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RestartMachine restarts operation for machine

  Second day restart operation for machine
*/
func (a *Client) RestartMachine(params *RestartMachineParams, opts ...ClientOption) (*RestartMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restartMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/restart",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestartMachineReader{formats: a.formats},
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
	success, ok := result.(*RestartMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restartMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RevertMachineSnapshot reverts snapshot operation for machine

  Second day revert snapshot operation for machine
*/
func (a *Client) RevertMachineSnapshot(params *RevertMachineSnapshotParams, opts ...ClientOption) (*RevertMachineSnapshotAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevertMachineSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "revertMachineSnapshot",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/revert/{snapshotId}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevertMachineSnapshotReader{formats: a.formats},
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
	success, ok := result.(*RevertMachineSnapshotAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revertMachineSnapshot: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ShutdownMachine shuts down operation for machine

  Second day shut down operation machine
*/
func (a *Client) ShutdownMachine(params *ShutdownMachineParams, opts ...ClientOption) (*ShutdownMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShutdownMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "shutdownMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/shutdown",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShutdownMachineReader{formats: a.formats},
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
	success, ok := result.(*ShutdownMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for shutdownMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SuspendMachine suspends operation for machine

  Second day suspend operation for machine
*/
func (a *Client) SuspendMachine(params *SuspendMachineParams, opts ...ClientOption) (*SuspendMachineAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSuspendMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "suspendMachine",
		Method:             "POST",
		PathPattern:        "/iaas/api/machines/{id}/operations/suspend",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SuspendMachineReader{formats: a.formats},
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
	success, ok := result.(*SuspendMachineAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for suspendMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateMachine updates machine

  Update machine. Only description, tag and custom property updates are supported. All other properties in the MachineSpecification body are ignored.
*/
func (a *Client) UpdateMachine(params *UpdateMachineParams, opts ...ClientOption) (*UpdateMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateMachineParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateMachine",
		Method:             "PATCH",
		PathPattern:        "/iaas/api/machines/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateMachineReader{formats: a.formats},
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
	success, ok := result.(*UpdateMachineOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateMachine: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
