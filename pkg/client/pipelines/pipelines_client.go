// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pipelines API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pipelines API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ActOnPipelineUsingPOST(params *ActOnPipelineUsingPOSTParams, opts ...ClientOption) (*ActOnPipelineUsingPOSTOK, error)

	ClonePipelineByNameUsingPOST(params *ClonePipelineByNameUsingPOSTParams, opts ...ClientOption) (*ClonePipelineByNameUsingPOSTOK, error)

	CreatePipelineUsingPOST(params *CreatePipelineUsingPOSTParams, opts ...ClientOption) (*CreatePipelineUsingPOSTOK, error)

	DeletePipelineByIDUsingDELETE(params *DeletePipelineByIDUsingDELETEParams, opts ...ClientOption) (*DeletePipelineByIDUsingDELETEOK, error)

	DeletePipelineByNameUsingDELETE(params *DeletePipelineByNameUsingDELETEParams, opts ...ClientOption) (*DeletePipelineByNameUsingDELETEOK, error)

	ExecutePipelineByIDUsingPOST(params *ExecutePipelineByIDUsingPOSTParams, opts ...ClientOption) (*ExecutePipelineByIDUsingPOSTOK, *ExecutePipelineByIDUsingPOSTAccepted, error)

	ExecutePipelineByNameUsingPOST(params *ExecutePipelineByNameUsingPOSTParams, opts ...ClientOption) (*ExecutePipelineByNameUsingPOSTOK, *ExecutePipelineByNameUsingPOSTAccepted, error)

	ExportUsingGET(params *ExportUsingGETParams, opts ...ClientOption) (*ExportUsingGETOK, error)

	GetAllPipelinesUsingGET(params *GetAllPipelinesUsingGETParams, opts ...ClientOption) (*GetAllPipelinesUsingGETOK, error)

	GetExecutionByIndexAndPipelineIDUsingGET(params *GetExecutionByIndexAndPipelineIDUsingGETParams, opts ...ClientOption) (*GetExecutionByIndexAndPipelineIDUsingGETOK, error)

	GetExecutionByNameAndIndexUsingGET(params *GetExecutionByNameAndIndexUsingGETParams, opts ...ClientOption) (*GetExecutionByNameAndIndexUsingGETOK, error)

	GetExecutionsByIDUsingGET(params *GetExecutionsByIDUsingGETParams, opts ...ClientOption) (*GetExecutionsByIDUsingGETOK, error)

	GetExecutionsByNameUsingGET(params *GetExecutionsByNameUsingGETParams, opts ...ClientOption) (*GetExecutionsByNameUsingGETOK, error)

	GetPipelineByIDUsingGET(params *GetPipelineByIDUsingGETParams, opts ...ClientOption) (*GetPipelineByIDUsingGETOK, error)

	GetPipelineByNameUsingGET(params *GetPipelineByNameUsingGETParams, opts ...ClientOption) (*GetPipelineByNameUsingGETOK, error)

	GetPipelineTilesUsingGET(params *GetPipelineTilesUsingGETParams, opts ...ClientOption) (*GetPipelineTilesUsingGETOK, error)

	ImportYamlUsingPOST(params *ImportYamlUsingPOSTParams, opts ...ClientOption) (*ImportYamlUsingPOSTOK, error)

	PatchPipelineByNameUsingPATCH(params *PatchPipelineByNameUsingPATCHParams, opts ...ClientOption) (*PatchPipelineByNameUsingPATCHOK, error)

	PatchPipelineUsingPATCH(params *PatchPipelineUsingPATCHParams, opts ...ClientOption) (*PatchPipelineUsingPATCHOK, error)

	UpdatePipelineByNameUsingPUT(params *UpdatePipelineByNameUsingPUTParams, opts ...ClientOption) (*UpdatePipelineByNameUsingPUTOK, error)

	UpdatePipelineUsingPUT(params *UpdatePipelineUsingPUTParams, opts ...ClientOption) (*UpdatePipelineUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ActOnPipelineUsingPOST clones a pipeline

Clone a Pipeline with the given id
*/
func (a *Client) ActOnPipelineUsingPOST(params *ActOnPipelineUsingPOSTParams, opts ...ClientOption) (*ActOnPipelineUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActOnPipelineUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "actOnPipelineUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/pipelines/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActOnPipelineUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*ActOnPipelineUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for actOnPipelineUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ClonePipelineByNameUsingPOST clones a pipeline by project and name

Clone a Pipeline with the given project and name
*/
func (a *Client) ClonePipelineByNameUsingPOST(params *ClonePipelineByNameUsingPOSTParams, opts ...ClientOption) (*ClonePipelineByNameUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClonePipelineByNameUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "clonePipelineByNameUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ClonePipelineByNameUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*ClonePipelineByNameUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for clonePipelineByNameUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreatePipelineUsingPOST creates a pipeline

Create a Pipeline based on the given project
*/
func (a *Client) CreatePipelineUsingPOST(params *CreatePipelineUsingPOSTParams, opts ...ClientOption) (*CreatePipelineUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePipelineUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPipelineUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/pipelines",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePipelineUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreatePipelineUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPipelineUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePipelineByIDUsingDELETE deletes a pipeline by id

Delete a Pipeline with the given id
*/
func (a *Client) DeletePipelineByIDUsingDELETE(params *DeletePipelineByIDUsingDELETEParams, opts ...ClientOption) (*DeletePipelineByIDUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePipelineByIDUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePipelineByIdUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/pipelines/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePipelineByIDUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeletePipelineByIDUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePipelineByIdUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePipelineByNameUsingDELETE deletes a pipeline by project and name

Delete a Pipeline with the given project and name
*/
func (a *Client) DeletePipelineByNameUsingDELETE(params *DeletePipelineByNameUsingDELETEParams, opts ...ClientOption) (*DeletePipelineByNameUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePipelineByNameUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePipelineByNameUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePipelineByNameUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeletePipelineByNameUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePipelineByNameUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecutePipelineByIDUsingPOST executes a pipeline

Execute a Pipeline based on the given id
*/
func (a *Client) ExecutePipelineByIDUsingPOST(params *ExecutePipelineByIDUsingPOSTParams, opts ...ClientOption) (*ExecutePipelineByIDUsingPOSTOK, *ExecutePipelineByIDUsingPOSTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecutePipelineByIDUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "executePipelineByIdUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/pipelines/{id}/executions",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecutePipelineByIDUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ExecutePipelineByIDUsingPOSTOK:
		return value, nil, nil
	case *ExecutePipelineByIDUsingPOSTAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pipelines: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExecutePipelineByNameUsingPOST executes a pipeline

Execute a Pipeline based on the given project and name
*/
func (a *Client) ExecutePipelineByNameUsingPOST(params *ExecutePipelineByNameUsingPOSTParams, opts ...ClientOption) (*ExecutePipelineByNameUsingPOSTOK, *ExecutePipelineByNameUsingPOSTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecutePipelineByNameUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "executePipelineByNameUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}/executions",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExecutePipelineByNameUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ExecutePipelineByNameUsingPOSTOK:
		return value, nil, nil
	case *ExecutePipelineByNameUsingPOSTAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pipelines: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ExportUsingGET exports

Exports a single pipeline (and endpoints referred in that pipeline) or a list of pipelines/endpoints in a given project or a custom integration with given list of versions as 'YAML' in a given project
*/
func (a *Client) ExportUsingGET(params *ExportUsingGETParams, opts ...ClientOption) (*ExportUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "exportUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/export",
		ProducesMediaTypes: []string{"application/x-yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportUsingGETReader{formats: a.formats},
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
	success, ok := result.(*ExportUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllPipelinesUsingGET gets all pipelines

Get all Pipelines with specified paging and filter parameters.
*/
func (a *Client) GetAllPipelinesUsingGET(params *GetAllPipelinesUsingGETParams, opts ...ClientOption) (*GetAllPipelinesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPipelinesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllPipelinesUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipelines",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllPipelinesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllPipelinesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllPipelinesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExecutionByIndexAndPipelineIDUsingGET gets an execution

Get an Execution based on the given pipeline id and execution index
*/
func (a *Client) GetExecutionByIndexAndPipelineIDUsingGET(params *GetExecutionByIndexAndPipelineIDUsingGETParams, opts ...ClientOption) (*GetExecutionByIndexAndPipelineIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExecutionByIndexAndPipelineIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExecutionByIndexAndPipelineIdUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipelines/{id}/executions/{index}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExecutionByIndexAndPipelineIDUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetExecutionByIndexAndPipelineIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExecutionByIndexAndPipelineIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExecutionByNameAndIndexUsingGET gets an execution

Get an Execution based on the given pipeline id and execution index
*/
func (a *Client) GetExecutionByNameAndIndexUsingGET(params *GetExecutionByNameAndIndexUsingGETParams, opts ...ClientOption) (*GetExecutionByNameAndIndexUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExecutionByNameAndIndexUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExecutionByNameAndIndexUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}/executions/{index}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExecutionByNameAndIndexUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetExecutionByNameAndIndexUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExecutionByNameAndIndexUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExecutionsByIDUsingGET gets all executions of a pipeline

Get all Executions of a Pipeline with specified paging and filter parameters.
*/
func (a *Client) GetExecutionsByIDUsingGET(params *GetExecutionsByIDUsingGETParams, opts ...ClientOption) (*GetExecutionsByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExecutionsByIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExecutionsByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipelines/{id}/executions",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExecutionsByIDUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetExecutionsByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExecutionsByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetExecutionsByNameUsingGET gets all executions of a pipeline

Get all Executions of a Pipeline with specified paging and filter parameters.
*/
func (a *Client) GetExecutionsByNameUsingGET(params *GetExecutionsByNameUsingGETParams, opts ...ClientOption) (*GetExecutionsByNameUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExecutionsByNameUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getExecutionsByNameUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}/executions",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExecutionsByNameUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetExecutionsByNameUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getExecutionsByNameUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPipelineByIDUsingGET gets a pipeline

Get a Pipeline with the given id
*/
func (a *Client) GetPipelineByIDUsingGET(params *GetPipelineByIDUsingGETParams, opts ...ClientOption) (*GetPipelineByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineByIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPipelineByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipelines/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelineByIDUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetPipelineByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPipelineByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPipelineByNameUsingGET gets a pipeline by project and name

Get a Pipeline with the given project and name
*/
func (a *Client) GetPipelineByNameUsingGET(params *GetPipelineByNameUsingGETParams, opts ...ClientOption) (*GetPipelineByNameUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineByNameUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPipelineByNameUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelineByNameUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetPipelineByNameUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPipelineByNameUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPipelineTilesUsingGET gets pipeline tiles

Get Pipeline Tiles
*/
func (a *Client) GetPipelineTilesUsingGET(params *GetPipelineTilesUsingGETParams, opts ...ClientOption) (*GetPipelineTilesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPipelineTilesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPipelineTilesUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/pipeline-tiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPipelineTilesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetPipelineTilesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPipelineTilesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImportYamlUsingPOST imports

Imports pipeline(s)/endpoint(s) into Code Stream.
*/
func (a *Client) ImportYamlUsingPOST(params *ImportYamlUsingPOSTParams, opts ...ClientOption) (*ImportYamlUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportYamlUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "importYamlUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/import",
		ProducesMediaTypes: []string{"application/x-yaml"},
		ConsumesMediaTypes: []string{"application/x-yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportYamlUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*ImportYamlUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importYamlUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchPipelineByNameUsingPATCH updates a pipeline by project and name

Update a Pipeline with the given project and name
*/
func (a *Client) PatchPipelineByNameUsingPATCH(params *PatchPipelineByNameUsingPATCHParams, opts ...ClientOption) (*PatchPipelineByNameUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPipelineByNameUsingPATCHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchPipelineByNameUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPipelineByNameUsingPATCHReader{formats: a.formats},
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
	success, ok := result.(*PatchPipelineByNameUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchPipelineByNameUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchPipelineUsingPATCH updates a pipeline by id

Update a Pipeline with the given id
*/
func (a *Client) PatchPipelineUsingPATCH(params *PatchPipelineUsingPATCHParams, opts ...ClientOption) (*PatchPipelineUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPipelineUsingPATCHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchPipelineUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/codestream/api/pipelines/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPipelineUsingPATCHReader{formats: a.formats},
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
	success, ok := result.(*PatchPipelineUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchPipelineUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePipelineByNameUsingPUT updates a pipeline by project and name

Update a Pipeline with the given project and name
*/
func (a *Client) UpdatePipelineByNameUsingPUT(params *UpdatePipelineByNameUsingPUTParams, opts ...ClientOption) (*UpdatePipelineByNameUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePipelineByNameUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePipelineByNameUsingPUT",
		Method:             "PUT",
		PathPattern:        "/codestream/api/pipelines/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePipelineByNameUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdatePipelineByNameUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePipelineByNameUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePipelineUsingPUT updates a pipeline by id

Update a Pipeline with the given id
*/
func (a *Client) UpdatePipelineUsingPUT(params *UpdatePipelineUsingPUTParams, opts ...ClientOption) (*UpdatePipelineUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePipelineUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePipelineUsingPUT",
		Method:             "PUT",
		PathPattern:        "/codestream/api/pipelines/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePipelineUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*UpdatePipelineUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePipelineUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}