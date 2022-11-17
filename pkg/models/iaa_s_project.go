// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IaaSProject Projects link users and cloud zones, thus controlling who can use what cloud resources.<br>**HATEOAS** links:<br>**self** - Project - Self link to this project
//
// swagger:model IaaSProject
type IaaSProject struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// List of administrator users associated with the project. Only administrators can manage project's configuration.
	// Example: [ { \"email\":\"administrator@vmware.com\" } ]
	Administrators []*User `json:"administrators"`

	// List of storage, network and extensibility constraints to be applied when provisioning through this project.
	// Example: {\"network\":[{\"mandatory\": \"true\", \"expression\": \"env:dev\"}],\"storage\":[{\"mandatory\": \"false\", \"expression\": \"gold\"}],\"extensibility\":[{\"mandatory\": \"false\", \"expression\": \"key:value\"}]}
	Constraints map[string][]Constraint `json:"constraints,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// The project custom properties which are added to all requests in this project
	// Example: { \"property\" : \"value\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// The naming template to be used for machines provisioned in this project
	// Example: ${project.name}-test-${####}
	MachineNamingTemplate string `json:"machineNamingTemplate,omitempty"`

	// List of member users associated with the project.
	// Example: [ { \"email\":\"member@vmware.com\" } ]
	Members []*User `json:"members"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The timeout that should be used for Blueprint operations and Provisioning tasks. The timeout is in seconds
	OperationTimeout int64 `json:"operationTimeout,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user or display name of the group that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Type of a owner(user/ad_group) that owns the entity.
	// Example: ad_group
	OwnerType string `json:"ownerType,omitempty"`

	// Placement policy for the project. Determines how a zone will be selected for provisioning. DEFAULT or SPREAD.
	// Example: DEFAULT
	PlacementPolicy string `json:"placementPolicy,omitempty"`

	// Specifies whether the resources in this projects are shared or not.
	SharedResources bool `json:"sharedResources,omitempty"`

	// List of supervisor users associated with the project.
	// Example: [ { \"email\":\"supervisor@vmware.com\" } ]
	Supervisors []*User `json:"supervisors"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`

	// List of viewer users associated with the project.
	// Example: [ { \"email\":\"viewer@vmware.com\" } ]
	Viewers []*User `json:"viewers"`

	// List of Cloud Zones assigned to this project. You can limit deployment to a single region or allow multi-region placement by adding more than one cloud zone to a project. A cloud zone lists available resources. Use tags on resources to control workload placement.
	Zones []*ZoneAssignment `json:"zones"`
}

// Validate validates this iaas  project
func (m *IaaSProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdministrators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupervisors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IaaSProject) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_links" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_links" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateAdministrators(formats strfmt.Registry) error {
	if swag.IsZero(m.Administrators) { // not required
		return nil
	}

	for i := 0; i < len(m.Administrators); i++ {
		if swag.IsZero(m.Administrators[i]) { // not required
			continue
		}

		if m.Administrators[i] != nil {
			if err := m.Administrators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("administrators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("administrators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for k := range m.Constraints {

		if err := validate.Required("constraints"+"."+k, "body", m.Constraints[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.Constraints[k]); i++ {

			if err := m.Constraints[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *IaaSProject) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IaaSProject) validateMembers(formats strfmt.Registry) error {
	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateSupervisors(formats strfmt.Registry) error {
	if swag.IsZero(m.Supervisors) { // not required
		return nil
	}

	for i := 0; i < len(m.Supervisors); i++ {
		if swag.IsZero(m.Supervisors[i]) { // not required
			continue
		}

		if m.Supervisors[i] != nil {
			if err := m.Supervisors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supervisors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supervisors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateViewers(formats strfmt.Registry) error {
	if swag.IsZero(m.Viewers) { // not required
		return nil
	}

	for i := 0; i < len(m.Viewers); i++ {
		if swag.IsZero(m.Viewers[i]) { // not required
			continue
		}

		if m.Viewers[i] != nil {
			if err := m.Viewers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("viewers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("viewers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) validateZones(formats strfmt.Registry) error {
	if swag.IsZero(m.Zones) { // not required
		return nil
	}

	for i := 0; i < len(m.Zones); i++ {
		if swag.IsZero(m.Zones[i]) { // not required
			continue
		}

		if m.Zones[i] != nil {
			if err := m.Zones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this iaas  project based on the context it is used
func (m *IaaSProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAdministrators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMembers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupervisors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IaaSProject) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if val, ok := m.Links[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateAdministrators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Administrators); i++ {

		if m.Administrators[i] != nil {
			if err := m.Administrators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("administrators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("administrators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Constraints {

		for i := 0; i < len(m.Constraints[k]); i++ {

			if err := m.Constraints[k][i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *IaaSProject) contextValidateMembers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Members); i++ {

		if m.Members[i] != nil {
			if err := m.Members[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateSupervisors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Supervisors); i++ {

		if m.Supervisors[i] != nil {
			if err := m.Supervisors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supervisors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supervisors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateViewers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Viewers); i++ {

		if m.Viewers[i] != nil {
			if err := m.Viewers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("viewers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("viewers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IaaSProject) contextValidateZones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Zones); i++ {

		if m.Zones[i] != nil {
			if err := m.Zones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IaaSProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IaaSProject) UnmarshalBinary(b []byte) error {
	var res IaaSProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
