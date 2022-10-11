// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new blueprint terraform integrations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for blueprint terraform integrations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBlueprintFromMappingUsingPOST(params *CreateBlueprintFromMappingUsingPOSTParams, opts ...ClientOption) (*CreateBlueprintFromMappingUsingPOSTOK, error)

	CreateBlueprintMappingUsingPOST(params *CreateBlueprintMappingUsingPOSTParams, opts ...ClientOption) (*CreateBlueprintMappingUsingPOSTOK, error)

	CreateTerraformVersionUsingPOST(params *CreateTerraformVersionUsingPOSTParams, opts ...ClientOption) (*CreateTerraformVersionUsingPOSTCreated, error)

	DeleteTerraformVersionUsingDELETE(params *DeleteTerraformVersionUsingDELETEParams, opts ...ClientOption) (*DeleteTerraformVersionUsingDELETENoContent, error)

	GetConfigurationSourceTreeUsingGET(params *GetConfigurationSourceTreeUsingGETParams, opts ...ClientOption) (*GetConfigurationSourceTreeUsingGETOK, error)

	GetTerraformConfigurationSourceCommitListUsingGET(params *GetTerraformConfigurationSourceCommitListUsingGETParams, opts ...ClientOption) (*GetTerraformConfigurationSourceCommitListUsingGETOK, error)

	GetTerraformConfigurationSourcesUsingGET(params *GetTerraformConfigurationSourcesUsingGETParams, opts ...ClientOption) (*GetTerraformConfigurationSourcesUsingGETOK, error)

	GetTerraformVersionUsingGET(params *GetTerraformVersionUsingGETParams, opts ...ClientOption) (*GetTerraformVersionUsingGETOK, error)

	ListTerraformVersionsUsingGET(params *ListTerraformVersionsUsingGETParams, opts ...ClientOption) (*ListTerraformVersionsUsingGETOK, error)

	UpdateTerraformVersionUsingPATCH(params *UpdateTerraformVersionUsingPATCHParams, opts ...ClientOption) (*UpdateTerraformVersionUsingPATCHOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBlueprintFromMappingUsingPOST creates a blueprint from a terraform blueprint configuration obtained from create blueprint mapping
*/
func (a *Client) CreateBlueprintFromMappingUsingPOST(params *CreateBlueprintFromMappingUsingPOSTParams, opts ...ClientOption) (*CreateBlueprintFromMappingUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlueprintFromMappingUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBlueprintFromMappingUsingPOST",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/create-blueprint-from-mapping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBlueprintFromMappingUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreateBlueprintFromMappingUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlueprintFromMappingUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateBlueprintMappingUsingPOST retrieves and parses the specified terraform configuration file s and returns relevant information for blueprint construction
*/
func (a *Client) CreateBlueprintMappingUsingPOST(params *CreateBlueprintMappingUsingPOSTParams, opts ...ClientOption) (*CreateBlueprintMappingUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlueprintMappingUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBlueprintMappingUsingPOST",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/create-blueprint-mapping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBlueprintMappingUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreateBlueprintMappingUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlueprintMappingUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateTerraformVersionUsingPOST creates a version
*/
func (a *Client) CreateTerraformVersionUsingPOST(params *CreateTerraformVersionUsingPOSTParams, opts ...ClientOption) (*CreateTerraformVersionUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTerraformVersionUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createTerraformVersionUsingPOST",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTerraformVersionUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*CreateTerraformVersionUsingPOSTCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTerraformVersionUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTerraformVersionUsingDELETE deletes a terraform version
*/
func (a *Client) DeleteTerraformVersionUsingDELETE(params *DeleteTerraformVersionUsingDELETEParams, opts ...ClientOption) (*DeleteTerraformVersionUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTerraformVersionUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTerraformVersionUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/versions/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTerraformVersionUsingDELETEReader{formats: a.formats},
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
	success, ok := result.(*DeleteTerraformVersionUsingDELETENoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTerraformVersionUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetConfigurationSourceTreeUsingGET shows directories of the configuration source repository that correspond to terraform configurations
*/
func (a *Client) GetConfigurationSourceTreeUsingGET(params *GetConfigurationSourceTreeUsingGETParams, opts ...ClientOption) (*GetConfigurationSourceTreeUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigurationSourceTreeUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getConfigurationSourceTreeUsingGET",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetConfigurationSourceTreeUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetConfigurationSourceTreeUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getConfigurationSourceTreeUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTerraformConfigurationSourceCommitListUsingGET returns a paginated list of commits for a specified configuration source
*/
func (a *Client) GetTerraformConfigurationSourceCommitListUsingGET(params *GetTerraformConfigurationSourceCommitListUsingGETParams, opts ...ClientOption) (*GetTerraformConfigurationSourceCommitListUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTerraformConfigurationSourceCommitListUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTerraformConfigurationSourceCommitListUsingGET",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/get-configuration-source-commits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTerraformConfigurationSourceCommitListUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetTerraformConfigurationSourceCommitListUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTerraformConfigurationSourceCommitListUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTerraformConfigurationSourcesUsingGET returns a paginated list of configuration sources configured as storage for terraform configurations
*/
func (a *Client) GetTerraformConfigurationSourcesUsingGET(params *GetTerraformConfigurationSourcesUsingGETParams, opts ...ClientOption) (*GetTerraformConfigurationSourcesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTerraformConfigurationSourcesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTerraformConfigurationSourcesUsingGET",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/get-configuration-sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTerraformConfigurationSourcesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetTerraformConfigurationSourcesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTerraformConfigurationSourcesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTerraformVersionUsingGET returns terraform version details
*/
func (a *Client) GetTerraformVersionUsingGET(params *GetTerraformVersionUsingGETParams, opts ...ClientOption) (*GetTerraformVersionUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTerraformVersionUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTerraformVersionUsingGET",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/versions/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTerraformVersionUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetTerraformVersionUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTerraformVersionUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListTerraformVersionsUsingGET lists terraform versions
*/
func (a *Client) ListTerraformVersionsUsingGET(params *ListTerraformVersionsUsingGETParams, opts ...ClientOption) (*ListTerraformVersionsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTerraformVersionsUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTerraformVersionsUsingGET",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTerraformVersionsUsingGETReader{formats: a.formats},
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
	success, ok := result.(*ListTerraformVersionsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTerraformVersionsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateTerraformVersionUsingPATCH updates a terraform version
*/
func (a *Client) UpdateTerraformVersionUsingPATCH(params *UpdateTerraformVersionUsingPATCHParams, opts ...ClientOption) (*UpdateTerraformVersionUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTerraformVersionUsingPATCHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateTerraformVersionUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/blueprint/api/blueprint-integrations/terraform/versions/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTerraformVersionUsingPATCHReader{formats: a.formats},
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
	success, ok := result.(*UpdateTerraformVersionUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateTerraformVersionUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
