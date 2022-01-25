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

// MachineSpecification Specification for a cloud agnostic machine.
//
// swagger:model MachineSpecification
type MachineSpecification struct {

	// A valid cloud config data in json-escaped yaml syntax
	BootConfig *MachineBootConfig `json:"bootConfig,omitempty"`

	// A set of settings that specify how the provided Boot Config should be handled
	BootConfigSettings *MachineBootConfigSettings `json:"bootConfigSettings,omitempty"`

	// Constraints that are used to drive placement policies for the virtual machine that is produced from this specification. Constraint expressions are matched against tags on existing placement targets.
	// Example: [{\"mandatory\" : \"true\", \"expression\": \"environment\":\"prod\"}, {\"mandatory\" : \"false\", \"expression\": \"pci\"}]
	Constraints []*Constraint `json:"constraints"`

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud
	Description string `json:"description,omitempty"`

	// A set of disk specifications for this machine.
	Disks []*DiskAttachmentSpecification `json:"disks"`

	// Flavor of machine instance.
	// Example: small, medium, large
	// Required: true
	Flavor *string `json:"flavor"`

	// Provider specific flavor reference. Valid if no flavor property is provided
	// Example: t2.micro
	// Required: true
	FlavorRef *string `json:"flavorRef"`

	// Type of image used for this machine.
	// Example: vmware-gold-master, ubuntu-latest, rhel-compliant, windows
	// Required: true
	Image *string `json:"image"`

	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets.
	// Example: [{\"mandatory\" : \"true\", \"expression\": \"environment:prod\"}, {\"mandatory\" : \"false\", \"expression\": \"pci\"}]
	ImageDiskConstraints []*Constraint `json:"imageDiskConstraints"`

	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	// Example: ami-f6795a8c
	// Required: true
	ImageRef *string `json:"imageRef"`

	// Number of machines to provision - default 1.
	// Example: 3
	MachineCount int32 `json:"machineCount,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics []*NetworkInterfaceSpecification `json:"nics"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`

	// Settings to remotely connect to the provisioned machine, by public/private key pair or username/password authentication. AWS and vSphere support key pair. Azure supports key pair or username/password.
	RemoteAccess *RemoteAccessSpecification `json:"remoteAccess,omitempty"`

	// Settings to apply salt configuration on the provisioned machine.
	SaltConfiguration *SaltConfiguration `json:"saltConfiguration,omitempty"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this machine specification
func (m *MachineSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBootConfigSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlavor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlavorRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageDiskConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaltConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineSpecification) validateBootConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.BootConfig) { // not required
		return nil
	}

	if m.BootConfig != nil {
		if err := m.BootConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootConfig")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) validateBootConfigSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.BootConfigSettings) { // not required
		return nil
	}

	if m.BootConfigSettings != nil {
		if err := m.BootConfigSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootConfigSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootConfigSettings")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) validateConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateFlavor(formats strfmt.Registry) error {

	if err := validate.Required("flavor", "body", m.Flavor); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateFlavorRef(formats strfmt.Registry) error {

	if err := validate.Required("flavorRef", "body", m.FlavorRef); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateImageDiskConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageDiskConstraints) { // not required
		return nil
	}

	for i := 0; i < len(m.ImageDiskConstraints); i++ {
		if swag.IsZero(m.ImageDiskConstraints[i]) { // not required
			continue
		}

		if m.ImageDiskConstraints[i] != nil {
			if err := m.ImageDiskConstraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imageDiskConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("imageDiskConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateImageRef(formats strfmt.Registry) error {

	if err := validate.Required("imageRef", "body", m.ImageRef); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateNics(formats strfmt.Registry) error {
	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateRemoteAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteAccess) { // not required
		return nil
	}

	if m.RemoteAccess != nil {
		if err := m.RemoteAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remoteAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remoteAccess")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) validateSaltConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.SaltConfiguration) { // not required
		return nil
	}

	if m.SaltConfiguration != nil {
		if err := m.SaltConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saltConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saltConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this machine specification based on the context it is used
func (m *MachineSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBootConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBootConfigSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImageDiskConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSaltConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineSpecification) contextValidateBootConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.BootConfig != nil {
		if err := m.BootConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootConfig")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) contextValidateBootConfigSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.BootConfigSettings != nil {
		if err := m.BootConfigSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootConfigSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bootConfigSettings")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) contextValidateConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Constraints); i++ {

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) contextValidateImageDiskConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImageDiskConstraints); i++ {

		if m.ImageDiskConstraints[i] != nil {
			if err := m.ImageDiskConstraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imageDiskConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("imageDiskConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) contextValidateNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Nics); i++ {

		if m.Nics[i] != nil {
			if err := m.Nics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) contextValidateRemoteAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.RemoteAccess != nil {
		if err := m.RemoteAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remoteAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("remoteAccess")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) contextValidateSaltConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.SaltConfiguration != nil {
		if err := m.SaltConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saltConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saltConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineSpecification) UnmarshalBinary(b []byte) error {
	var res MachineSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
