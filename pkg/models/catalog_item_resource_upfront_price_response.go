// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogItemResourceUpfrontPriceResponse CatalogItemResourceUpfrontPriceResponse
//
// A response with upfront price for a resource in given catalog item
//
// swagger:model CatalogItemResourceUpfrontPriceResponse
type CatalogItemResourceUpfrontPriceResponse struct {

	// Additional Price incurred for the catalog.
	// Read Only: true
	DailyAdditionalPrice float64 `json:"dailyAdditionalPrice,omitempty"`

	// Compute Price of the catalog resource.
	// Read Only: true
	DailyComputePrice float64 `json:"dailyComputePrice,omitempty"`

	// Network Price of the catalog resource.
	// Read Only: true
	DailyNetworkPrice float64 `json:"dailyNetworkPrice,omitempty"`

	// Storage Price of the catalog resource.
	// Read Only: true
	DailyStoragePrice float64 `json:"dailyStoragePrice,omitempty"`

	// Total Price of the catalog resource.
	// Read Only: true
	DailyTotalPrice float64 `json:"dailyTotalPrice,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// Id
	// Read Only: true
	ResourceUpfrontPriceID string `json:"resourceUpfrontPriceId,omitempty"`

	// Upfront price sync status
	// Read Only: true
	// Enum: [SUCCESS ERROR DATA_NOT_AVAILABLE CURRENCY_NOT_SET]
	Status string `json:"status,omitempty"`

	// Upfront price status detail.
	// Read Only: true
	StatusDetails string `json:"statusDetails,omitempty"`

	// Monetary unit.
	// Read Only: true
	Unit string `json:"unit,omitempty"`
}

// Validate validates this catalog item resource upfront price response
func (m *CatalogItemResourceUpfrontPriceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var catalogItemResourceUpfrontPriceResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","ERROR","DATA_NOT_AVAILABLE","CURRENCY_NOT_SET"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogItemResourceUpfrontPriceResponseTypeStatusPropEnum = append(catalogItemResourceUpfrontPriceResponseTypeStatusPropEnum, v)
	}
}

const (

	// CatalogItemResourceUpfrontPriceResponseStatusSUCCESS captures enum value "SUCCESS"
	CatalogItemResourceUpfrontPriceResponseStatusSUCCESS string = "SUCCESS"

	// CatalogItemResourceUpfrontPriceResponseStatusERROR captures enum value "ERROR"
	CatalogItemResourceUpfrontPriceResponseStatusERROR string = "ERROR"

	// CatalogItemResourceUpfrontPriceResponseStatusDATANOTAVAILABLE captures enum value "DATA_NOT_AVAILABLE"
	CatalogItemResourceUpfrontPriceResponseStatusDATANOTAVAILABLE string = "DATA_NOT_AVAILABLE"

	// CatalogItemResourceUpfrontPriceResponseStatusCURRENCYNOTSET captures enum value "CURRENCY_NOT_SET"
	CatalogItemResourceUpfrontPriceResponseStatusCURRENCYNOTSET string = "CURRENCY_NOT_SET"
)

// prop value enum
func (m *CatalogItemResourceUpfrontPriceResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, catalogItemResourceUpfrontPriceResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *CatalogItemResourceUpfrontPriceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemResourceUpfrontPriceResponse) UnmarshalBinary(b []byte) error {
	var res CatalogItemResourceUpfrontPriceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
