// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudAccountAzureRegionEnumerationSpecification Specification for a region enumeration of azure cloud account.
//
// swagger:model CloudAccountAzureRegionEnumerationSpecification
type CloudAccountAzureRegionEnumerationSpecification struct {

	// Azure Client Application ID. Either provide clientApplicationId or provide a cloudAccountId of an existing account.
	// Example: 3287dd6e-76d8-41b7-9856-2584969e7739
	ClientApplicationID string `json:"clientApplicationId,omitempty"`

	// Azure Client Application Secret Key. Either provide clientApplicationSecretKey or provide a cloudAccountId of an existing account.
	// Example: GDfdasDasdASFas321das32cas2x3dsXCSA76xdcasg=
	ClientApplicationSecretKey string `json:"clientApplicationSecretKey,omitempty"`

	// Existing cloud account id. Either provide id of existing account, or cloud account credentials: clientApplicationId, clientApplicationSecretKey and tenantId.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Azure Subscribtion ID. Either provide subscriptionId or provide a cloudAccountId of an existing account.
	// Example: 064865b2-e914-4717-b415-8806d17948f7
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// Azure Tenant ID. Either provide tenantId or provide a cloudAccountId of an existing account.
	// Example: 9a13d920-4691-4e2d-b5d5-9c4c1279bc9a
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this cloud account azure region enumeration specification
func (m *CloudAccountAzureRegionEnumerationSpecification) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cloud account azure region enumeration specification based on context it is used
func (m *CloudAccountAzureRegionEnumerationSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountAzureRegionEnumerationSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountAzureRegionEnumerationSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountAzureRegionEnumerationSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
