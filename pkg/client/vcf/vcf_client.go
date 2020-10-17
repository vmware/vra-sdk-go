// Code generated by go-swagger; DO NOT EDIT.

package vcf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new v c f API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v c f API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateServiceCredentialsUsingPOST(params *CreateServiceCredentialsUsingPOSTParams) (*CreateServiceCredentialsUsingPOSTOK, error)

	DeleteServiceAccountUsingDELETE(params *DeleteServiceAccountUsingDELETEParams) (*DeleteServiceAccountUsingDELETENoContent, error)

	DeleteServiceCredentialUsingDELETE(params *DeleteServiceCredentialUsingDELETEParams) (*DeleteServiceCredentialUsingDELETENoContent, error)

	EnumerateDomainsUsingPOST(params *EnumerateDomainsUsingPOSTParams) (*EnumerateDomainsUsingPOSTOK, error)

	GetDomainUsingGET(params *GetDomainUsingGETParams) (*GetDomainUsingGETOK, error)

	GetDomainsUsingGET(params *GetDomainsUsingGETParams) (*GetDomainsUsingGETOK, error)

	PatchServiceAccountUsingPATCH(params *PatchServiceAccountUsingPATCHParams) (*PatchServiceAccountUsingPATCHOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateServiceCredentialsUsingPOST creates service credentials
*/
func (a *Client) CreateServiceCredentialsUsingPOST(params *CreateServiceCredentialsUsingPOSTParams) (*CreateServiceCredentialsUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServiceCredentialsUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createServiceCredentialsUsingPOST",
		Method:             "POST",
		PathPattern:        "/content/api/vcf/{integrationId}/domain/{domainId}/service-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateServiceCredentialsUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateServiceCredentialsUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createServiceCredentialsUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteServiceAccountUsingDELETE deletes a service account
*/
func (a *Client) DeleteServiceAccountUsingDELETE(params *DeleteServiceAccountUsingDELETEParams) (*DeleteServiceAccountUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceAccountUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServiceAccountUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/content/api/vcf/{integrationId}/domain/{domainId}/service-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceAccountUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceAccountUsingDELETENoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteServiceAccountUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteServiceCredentialUsingDELETE deletes a service credential
*/
func (a *Client) DeleteServiceCredentialUsingDELETE(params *DeleteServiceCredentialUsingDELETEParams) (*DeleteServiceCredentialUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServiceCredentialUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteServiceCredentialUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/content/api/vcf/{integrationId}/domain/{domainId}/service-accounts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceCredentialUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceCredentialUsingDELETENoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteServiceCredentialUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EnumerateDomainsUsingPOST enumerates domains
*/
func (a *Client) EnumerateDomainsUsingPOST(params *EnumerateDomainsUsingPOSTParams) (*EnumerateDomainsUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnumerateDomainsUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "enumerateDomainsUsingPOST",
		Method:             "POST",
		PathPattern:        "/content/api/vcf/domains-enumeration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnumerateDomainsUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnumerateDomainsUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enumerateDomainsUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomainUsingGET gets domain details
*/
func (a *Client) GetDomainUsingGET(params *GetDomainUsingGETParams) (*GetDomainUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDomainUsingGET",
		Method:             "GET",
		PathPattern:        "/content/api/vcf/{integrationId}/domain/{domainId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomainUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomainUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomainsUsingGET gets domains
*/
func (a *Client) GetDomainsUsingGET(params *GetDomainsUsingGETParams) (*GetDomainsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomainsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDomainsUsingGET",
		Method:             "GET",
		PathPattern:        "/content/api/vcf/{integrationId}/domains",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomainsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomainsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDomainsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchServiceAccountUsingPATCH patches service account
*/
func (a *Client) PatchServiceAccountUsingPATCH(params *PatchServiceAccountUsingPATCHParams) (*PatchServiceAccountUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchServiceAccountUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchServiceAccountUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/content/api/vcf/{integrationId}/domain/{domainId}/service-accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchServiceAccountUsingPATCHReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchServiceAccountUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchServiceAccountUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}