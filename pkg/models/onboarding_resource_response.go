// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OnboardingResourceResponse onboarding resource response
//
// swagger:model OnboardingResourceResponse
type OnboardingResourceResponse struct {

	// Link to the parent deployment.
	DeploymentLink string `json:"deploymentLink,omitempty"`

	// Link to this resource.
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// Link to the parent plan.
	PlanLink string `json:"planLink,omitempty"`

	// Link to the compute resource (machine) in the Provisioning service inventory.
	ResourceLink string `json:"resourceLink,omitempty"`

	// Resource name, usually the machine name. Will be propagated to Cloud Assembly.
	ResourceName string `json:"resourceName,omitempty"`

	// Tag links in the Provisioning service inventory associated with the machine.
	TagLinks []string `json:"tagLinks"`

	// Link to the resource's tenant.
	TenantLink string `json:"tenantLink,omitempty"`
}

// Validate validates this onboarding resource response
func (m *OnboardingResourceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this onboarding resource response based on context it is used
func (m *OnboardingResourceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OnboardingResourceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnboardingResourceResponse) UnmarshalBinary(b []byte) error {
	var res OnboardingResourceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
