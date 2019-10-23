// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogItemUpfrontPriceResponse CatalogItemUpfrontPriceResponse
//
// The response to a catalog item upfront price request
// swagger:model CatalogItemUpfrontPriceResponse
type CatalogItemUpfrontPriceResponse struct {

	// Additional Price incurred for the catalog.
	// Read Only: true
	DailyAdditionalPrice float64 `json:"dailyAdditionalPrice,omitempty"`

	// Compute Price of the catalog.
	// Read Only: true
	DailyComputePrice float64 `json:"dailyComputePrice,omitempty"`

	// Network Price of the catalog.
	// Read Only: true
	DailyNetworkPrice float64 `json:"dailyNetworkPrice,omitempty"`

	// Storage Price of the catalog.
	// Read Only: true
	DailyStoragePrice float64 `json:"dailyStoragePrice,omitempty"`

	// Total Price of the catalog.
	// Read Only: true
	DailyTotalPrice float64 `json:"dailyTotalPrice,omitempty"`

	// resource price details
	ResourcePriceDetails []*CatalogItemResourceUpfrontPriceResponse `json:"resourcePriceDetails"`

	// Upfront price sync status
	// Read Only: true
	// Enum: [STARTED PREPARING_COST_ESTIMATION IN_PROGRESS SUCCESS ERROR DATA_NOT_AVAILABLE CURRENCY_NOT_SET PUBLIC_CLOUD_NOT_SUPPORTED]
	Status string `json:"status,omitempty"`

	// Upfront price status detail.
	// Read Only: true
	StatusDetails string `json:"statusDetails,omitempty"`

	// Monetary unit.
	// Read Only: true
	Unit string `json:"unit,omitempty"`

	// Id
	// Read Only: true
	UpfrontPriceID string `json:"upfrontPriceId,omitempty"`
}

// Validate validates this catalog item upfront price response
func (m *CatalogItemUpfrontPriceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourcePriceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogItemUpfrontPriceResponse) validateResourcePriceDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourcePriceDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourcePriceDetails); i++ {
		if swag.IsZero(m.ResourcePriceDetails[i]) { // not required
			continue
		}

		if m.ResourcePriceDetails[i] != nil {
			if err := m.ResourcePriceDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resourcePriceDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var catalogItemUpfrontPriceResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STARTED","PREPARING_COST_ESTIMATION","IN_PROGRESS","SUCCESS","ERROR","DATA_NOT_AVAILABLE","CURRENCY_NOT_SET","PUBLIC_CLOUD_NOT_SUPPORTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogItemUpfrontPriceResponseTypeStatusPropEnum = append(catalogItemUpfrontPriceResponseTypeStatusPropEnum, v)
	}
}

const (

	// CatalogItemUpfrontPriceResponseStatusSTARTED captures enum value "STARTED"
	CatalogItemUpfrontPriceResponseStatusSTARTED string = "STARTED"

	// CatalogItemUpfrontPriceResponseStatusPREPARINGCOSTESTIMATION captures enum value "PREPARING_COST_ESTIMATION"
	CatalogItemUpfrontPriceResponseStatusPREPARINGCOSTESTIMATION string = "PREPARING_COST_ESTIMATION"

	// CatalogItemUpfrontPriceResponseStatusINPROGRESS captures enum value "IN_PROGRESS"
	CatalogItemUpfrontPriceResponseStatusINPROGRESS string = "IN_PROGRESS"

	// CatalogItemUpfrontPriceResponseStatusSUCCESS captures enum value "SUCCESS"
	CatalogItemUpfrontPriceResponseStatusSUCCESS string = "SUCCESS"

	// CatalogItemUpfrontPriceResponseStatusERROR captures enum value "ERROR"
	CatalogItemUpfrontPriceResponseStatusERROR string = "ERROR"

	// CatalogItemUpfrontPriceResponseStatusDATANOTAVAILABLE captures enum value "DATA_NOT_AVAILABLE"
	CatalogItemUpfrontPriceResponseStatusDATANOTAVAILABLE string = "DATA_NOT_AVAILABLE"

	// CatalogItemUpfrontPriceResponseStatusCURRENCYNOTSET captures enum value "CURRENCY_NOT_SET"
	CatalogItemUpfrontPriceResponseStatusCURRENCYNOTSET string = "CURRENCY_NOT_SET"

	// CatalogItemUpfrontPriceResponseStatusPUBLICCLOUDNOTSUPPORTED captures enum value "PUBLIC_CLOUD_NOT_SUPPORTED"
	CatalogItemUpfrontPriceResponseStatusPUBLICCLOUDNOTSUPPORTED string = "PUBLIC_CLOUD_NOT_SUPPORTED"
)

// prop value enum
func (m *CatalogItemUpfrontPriceResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, catalogItemUpfrontPriceResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CatalogItemUpfrontPriceResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogItemUpfrontPriceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemUpfrontPriceResponse) UnmarshalBinary(b []byte) error {
	var res CatalogItemUpfrontPriceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
