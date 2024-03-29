// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OnboardingDeploymentResponse onboarding deployment response
//
// swagger:model OnboardingDeploymentResponse
type OnboardingDeploymentResponse struct {

	// Link to the onboarded deployment in Cloud Assembly.
	ConsumerDeploymentLink string `json:"consumerDeploymentLink,omitempty"`

	// Deployment description. Will be propagated to Cloud Assembly deployment.
	Description string `json:"description,omitempty"`

	// Link to this deployment.
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// Deployment name. Will be propagated to Cloud Assembly deployment.
	Name string `json:"name,omitempty"`

	// Link to the parent plan.
	PlanLink string `json:"planLink,omitempty"`

	// Link to the deployment's tenant.
	TenantLink string `json:"tenantLink,omitempty"`
}

// Validate validates this onboarding deployment response
func (m *OnboardingDeploymentResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this onboarding deployment response based on context it is used
func (m *OnboardingDeploymentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OnboardingDeploymentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnboardingDeploymentResponse) UnmarshalBinary(b []byte) error {
	var res OnboardingDeploymentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
