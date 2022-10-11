// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloudAccountSpecification Specification for a generic cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.
//
// swagger:model CloudAccountSpecification
type CloudAccountSpecification struct {

	// Cloud accounts to associate with this cloud account
	// Example: [ \"42f3e0d199d134755684cd935435a\" ]
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Cloud Account IDs and directionalities create associations to other vSphere cloud accounts that can be used for workload mobility. ID refers to an associated cloud account, and directionality can be unidirectional or bidirectional. Only supported on vSphere cloud accounts.
	// Example: { \"42f3e0d199d134755684cd935435a\": \"BIDIRECTIONAL\" }
	AssociatedMobilityCloudAccountIds map[string]string `json:"associatedMobilityCloudAccountIds,omitempty"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Cloud Account specific properties supplied in as name value pairs
	// Example: {\"supportPublicImages\": \"true\", \"acceptSelfSignedCertificate\": \"true\" }
	// Required: true
	CloudAccountProperties map[string]string `json:"cloudAccountProperties"`

	// Cloud account type
	// Example: vsphere, aws, azure, nsxv, nsxt, vmc
	// Required: true
	CloudAccountType *string `json:"cloudAccountType"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Additional custom properties that may be used to extend the Cloud Account. In case of AAP, provide environment property here.Example: "customProperties": {
	//         "environment": "aap"
	//     }
	// Example: { \"sampleadapterProjectId\" : \"projectId\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Secret access key or password to be used to authenticate with the cloud account. In case of AAP pass a dummy value.
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	// Required: true
	PrivateKey *string `json:"privateKey"`

	// Access key id or username to be used to authenticate with the cloud account
	// Example: ACDC55DB4MFH6ADG75KK
	// Required: true
	PrivateKeyID *string `json:"privateKeyId"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// 'regionInfos' is a required parameter for AWS, AZURE, GCP, VSPHERE, VMC, VCF cloud account types.
	// Example: [{ \"name\": \"East Asia\",\"externalRegionId\": \"eastasia\"}]
	// Required: true
	Regions []*RegionSpecification `json:"regions"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this cloud account specification
func (m *CloudAccountSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssociatedMobilityCloudAccountIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCertificateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateKeyID(formats); err != nil {
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

// additional properties value enum
var cloudAccountSpecificationAssociatedMobilityCloudAccountIdsValueEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNIDIRECTIONAL","BIDIRECTIONAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudAccountSpecificationAssociatedMobilityCloudAccountIdsValueEnum = append(cloudAccountSpecificationAssociatedMobilityCloudAccountIdsValueEnum, v)
	}
}

func (m *CloudAccountSpecification) validateAssociatedMobilityCloudAccountIdsValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloudAccountSpecificationAssociatedMobilityCloudAccountIdsValueEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CloudAccountSpecification) validateAssociatedMobilityCloudAccountIds(formats strfmt.Registry) error {
	if swag.IsZero(m.AssociatedMobilityCloudAccountIds) { // not required
		return nil
	}

	for k := range m.AssociatedMobilityCloudAccountIds {

		// value enum
		if err := m.validateAssociatedMobilityCloudAccountIdsValueEnum("associatedMobilityCloudAccountIds"+"."+k, "body", m.AssociatedMobilityCloudAccountIds[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *CloudAccountSpecification) validateCertificateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CertificateInfo) { // not required
		return nil
	}

	if m.CertificateInfo != nil {
		if err := m.CertificateInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountSpecification) validateCloudAccountProperties(formats strfmt.Registry) error {

	if err := validate.Required("cloudAccountProperties", "body", m.CloudAccountProperties); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountSpecification) validateCloudAccountType(formats strfmt.Registry) error {

	if err := validate.Required("cloudAccountType", "body", m.CloudAccountType); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountSpecification) validatePrivateKey(formats strfmt.Registry) error {

	if err := validate.Required("privateKey", "body", m.PrivateKey); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountSpecification) validatePrivateKeyID(formats strfmt.Registry) error {

	if err := validate.Required("privateKeyId", "body", m.PrivateKeyID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountSpecification) validateRegions(formats strfmt.Registry) error {

	if err := validate.Required("regions", "body", m.Regions); err != nil {
		return err
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

func (m *CloudAccountSpecification) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this cloud account specification based on the context it is used
func (m *CloudAccountSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *CloudAccountSpecification) contextValidateCertificateInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CertificateInfo != nil {
		if err := m.CertificateInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificateInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificateInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CloudAccountSpecification) contextValidateRegions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CloudAccountSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudAccountSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
