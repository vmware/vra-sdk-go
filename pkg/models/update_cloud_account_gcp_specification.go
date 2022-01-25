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

// UpdateCloudAccountGcpSpecification Specification for a GCP cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region where the associated cloud account resources are hosted.
//
// swagger:model UpdateCloudAccountGcpSpecification
type UpdateCloudAccountGcpSpecification struct {

	// GCP Client email
	// Example: 321743978432-compute@developer.gserviceaccount.com
	// Required: true
	ClientEmail *string `json:"clientEmail"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// GCP Private key
	// Example: -----BEGIN PRIVATE KEY-----\nMIICXgIHAASBgSDHikastc8+I81zCg/qWW8dMr8mqvXQ3qbPAmu0RjxoZVI47tvs\nkYlFAXOf0sPrhO2nUuooJngnHV0639iTTEYG1vckNaW2R6U5QTdQ5Rq5u+uV3pMk\n7w7Vs4n3urQ4jnqt7rTXbC1DNa/PFeAZatbf7ffBBy0IGO0zc128IshYcwIDAQAB\nAoGBALTNl2JxTvq4SDW/3VH0fZkQXWH1MM10oeMbB2qO5beWb11FGaOO77nGKfWc\nbYgfp5Ogrql2yhBvLAXnxH8bcqqwORtFhlyV68U1y4R+8WxDNh0aevxH8hRS/1X5\n963DJm1JlU0E+vStiktN0tC3ebH5hE+1OxbIHSZ+WOWLYX7JAkEA5uigRgKp8ScG\nauUijvdOLZIhHWq9y5Wz+nOHUuDw8P7wOTKU34QJAoWEe771p9Pf/GTA/kr0BQnP\nQvWUDxGzJwJBAN05C6krwPeryFKrKtjOGJIbiIoY72wRnoNcdEEs3HDRhf48YWFo\nriRbZylzzzNFy/gmzT6XJQTfktGqq+FZD9UCQGIJaGrxHJgfmpDuAhMzGsUsYtTr\niRox0D1Iqa7dhE693t5aBG010OF6MLqdZA1CXrn5SRtuVVaCSLZEL/2J5UcCQQDA\nd3MXucNnN4NPuS/L9HMYJWD7lPoosaORcgyK77bSSNgk+u9WSjbH1uYIAIPSffUZ\nbti+jc2dUg5wb+aeZlgJAkEAurrpmpqj5vg087ZngKfFGR5rozDiTsK5DceTV97K\na1Y+Nzl+XWTxDBWk4YPh2ZlKv402hZEfWBYxUDn5ZkH/bw==\n-----END PRIVATE KEY-----\n
	// Required: true
	PrivateKey *string `json:"privateKey"`

	// GCP Private key ID
	// Example: 027f73d50a19452eedf5775a9b42c5083678abdf
	// Required: true
	PrivateKeyID *string `json:"privateKeyId"`

	// GCP Project ID
	// Example: example-gcp-project
	// Required: true
	ProjectID *string `json:"projectId"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"europe-west2\",\"externalRegionId\": \"europe-west2\"}]
	Regions []*RegionSpecification `json:"regions"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this update cloud account gcp specification
func (m *UpdateCloudAccountGcpSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
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

func (m *UpdateCloudAccountGcpSpecification) validateClientEmail(formats strfmt.Registry) error {

	if err := validate.Required("clientEmail", "body", m.ClientEmail); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCloudAccountGcpSpecification) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCloudAccountGcpSpecification) validatePrivateKeyID(formats strfmt.Registry) error {

	if err := validate.Required("privateKeyId", "body", m.PrivateKeyID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCloudAccountGcpSpecification) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCloudAccountGcpSpecification) validateRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	for i := 0; i < len(m.Regions); i++ {
		if swag.IsZero(m.Regions[i]) { // not required
			continue
		}

		if m.Regions[i] != nil {
			if err := m.Regions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateCloudAccountGcpSpecification) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this update cloud account gcp specification based on the context it is used
func (m *UpdateCloudAccountGcpSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegions(ctx, formats); err != nil {
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

func (m *UpdateCloudAccountGcpSpecification) contextValidateRegions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Regions); i++ {

		if m.Regions[i] != nil {
			if err := m.Regions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("regions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("regions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateCloudAccountGcpSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *UpdateCloudAccountGcpSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCloudAccountGcpSpecification) UnmarshalBinary(b []byte) error {
	var res UpdateCloudAccountGcpSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
