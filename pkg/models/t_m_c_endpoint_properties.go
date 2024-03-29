// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TMCEndpointProperties TMCEndpointProperties
//
// swagger:model TMCEndpointProperties
type TMCEndpointProperties struct {

	// agent Id
	AgentID string `json:"agentId,omitempty"`

	// api endpoint
	APIEndpoint string `json:"apiEndpoint,omitempty"`

	// api token
	APIToken string `json:"apiToken,omitempty"`

	// ca certificate
	CaCertificate string `json:"caCertificate,omitempty"`

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// credentials type
	CredentialsType string `json:"credentialsType,omitempty"`

	// dc Id
	DcID string `json:"dcId,omitempty"`

	// default cluster group name
	DefaultClusterGroupName string `json:"defaultClusterGroupName,omitempty"`

	// default workspace name
	DefaultWorkspaceName string `json:"defaultWorkspaceName,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// private key
	PrivateKey string `json:"privateKey,omitempty"`

	// private key Id
	PrivateKeyID string `json:"privateKeyId,omitempty"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`

	// skip cert validation
	SkipCertValidation bool `json:"skipCertValidation,omitempty"`

	// user email
	UserEmail string `json:"userEmail,omitempty"`
}

// Validate validates this t m c endpoint properties
func (m *TMCEndpointProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this t m c endpoint properties based on context it is used
func (m *TMCEndpointProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TMCEndpointProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TMCEndpointProperties) UnmarshalBinary(b []byte) error {
	var res TMCEndpointProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
