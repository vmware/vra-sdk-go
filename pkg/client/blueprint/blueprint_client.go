// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new blueprint API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for blueprint API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateBlueprintUsingPOST1 creates a blueprint
*/
func (a *Client) CreateBlueprintUsingPOST1(params *CreateBlueprintUsingPOST1Params) (*CreateBlueprintUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlueprintUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBlueprintUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBlueprintUsingPOST1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBlueprintUsingPOST1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlueprintUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateBlueprintVersionUsingPOST1 creates version for the given blueprint ID
*/
func (a *Client) CreateBlueprintVersionUsingPOST1(params *CreateBlueprintVersionUsingPOST1Params) (*CreateBlueprintVersionUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBlueprintVersionUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBlueprintVersionUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateBlueprintVersionUsingPOST1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBlueprintVersionUsingPOST1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBlueprintVersionUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteBlueprintUsingDELETE1 deletes a blueprint
*/
func (a *Client) DeleteBlueprintUsingDELETE1(params *DeleteBlueprintUsingDELETE1Params) (*DeleteBlueprintUsingDELETE1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBlueprintUsingDELETE1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBlueprintUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteBlueprintUsingDELETE1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBlueprintUsingDELETE1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBlueprintUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlueprintInputsSchemaUsingGET1 returns blueprint inputs schema
*/
func (a *Client) GetBlueprintInputsSchemaUsingGET1(params *GetBlueprintInputsSchemaUsingGET1Params) (*GetBlueprintInputsSchemaUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintInputsSchemaUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintInputsSchemaUsingGET_1",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/inputs-schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBlueprintInputsSchemaUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBlueprintInputsSchemaUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlueprintInputsSchemaUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlueprintUsingGET1 returns blueprint details
*/
func (a *Client) GetBlueprintUsingGET1(params *GetBlueprintUsingGET1Params) (*GetBlueprintUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintUsingGET_1",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBlueprintUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBlueprintUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlueprintUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlueprintVersionInputsSchemaUsingGET1 returns blueprint version inputs schema
*/
func (a *Client) GetBlueprintVersionInputsSchemaUsingGET1(params *GetBlueprintVersionInputsSchemaUsingGET1Params) (*GetBlueprintVersionInputsSchemaUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintVersionInputsSchemaUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintVersionInputsSchemaUsingGET_1",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/versions/{version}/inputs-schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBlueprintVersionInputsSchemaUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBlueprintVersionInputsSchemaUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlueprintVersionInputsSchemaUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBlueprintVersionUsingGET1 returns versioned blueprint details
*/
func (a *Client) GetBlueprintVersionUsingGET1(params *GetBlueprintVersionUsingGET1Params) (*GetBlueprintVersionUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlueprintVersionUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBlueprintVersionUsingGET_1",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBlueprintVersionUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBlueprintVersionUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBlueprintVersionUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListBlueprintVersionsUsingGET lists blueprint versions
*/
func (a *Client) ListBlueprintVersionsUsingGET(params *ListBlueprintVersionsUsingGETParams) (*ListBlueprintVersionsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBlueprintVersionsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBlueprintVersionsUsingGET",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListBlueprintVersionsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBlueprintVersionsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBlueprintVersionsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListBlueprintsUsingGET1 lists draft blueprint
*/
func (a *Client) ListBlueprintsUsingGET1(params *ListBlueprintsUsingGET1Params) (*ListBlueprintsUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBlueprintsUsingGET1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBlueprintsUsingGET_1",
		Method:             "GET",
		PathPattern:        "/blueprint/api/blueprints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListBlueprintsUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBlueprintsUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBlueprintsUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReleaseBlueprintVersionUsingPOST1 releases versioned blueprint to catalog
*/
func (a *Client) ReleaseBlueprintVersionUsingPOST1(params *ReleaseBlueprintVersionUsingPOST1Params) (*ReleaseBlueprintVersionUsingPOST1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReleaseBlueprintVersionUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "releaseBlueprintVersionUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/release",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReleaseBlueprintVersionUsingPOST1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReleaseBlueprintVersionUsingPOST1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for releaseBlueprintVersionUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestoreBlueprintVersionUsingPOST1 restores content of draft from versioned blueprint
*/
func (a *Client) RestoreBlueprintVersionUsingPOST1(params *RestoreBlueprintVersionUsingPOST1Params) (*RestoreBlueprintVersionUsingPOST1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreBlueprintVersionUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "restoreBlueprintVersionUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestoreBlueprintVersionUsingPOST1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreBlueprintVersionUsingPOST1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreBlueprintVersionUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UnReleaseBlueprintVersionUsingPOST1 uns release versioned blueprint from catalog
*/
func (a *Client) UnReleaseBlueprintVersionUsingPOST1(params *UnReleaseBlueprintVersionUsingPOST1Params) (*UnReleaseBlueprintVersionUsingPOST1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnReleaseBlueprintVersionUsingPOST1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unReleaseBlueprintVersionUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/unrelease",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnReleaseBlueprintVersionUsingPOST1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnReleaseBlueprintVersionUsingPOST1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unReleaseBlueprintVersionUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateBlueprintUsingPUT1 updates a blueprint
*/
func (a *Client) UpdateBlueprintUsingPUT1(params *UpdateBlueprintUsingPUT1Params) (*UpdateBlueprintUsingPUT1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBlueprintUsingPUT1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateBlueprintUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/blueprint/api/blueprints/{blueprintId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateBlueprintUsingPUT1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBlueprintUsingPUT1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBlueprintUsingPUT_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
