// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CatalogItemResourceUpfrontPriceResponse CatalogItemResourceUpfrontPriceResponse
//
// # A response with upfront price for a resource in given catalog item
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

// ContextValidate validate this catalog item resource upfront price response based on the context it is used
func (m *CatalogItemResourceUpfrontPriceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDailyAdditionalPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDailyComputePrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDailyNetworkPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDailyStoragePrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDailyTotalPrice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceUpfrontPriceID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateDailyAdditionalPrice(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dailyAdditionalPrice", "body", float64(m.DailyAdditionalPrice)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateDailyComputePrice(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dailyComputePrice", "body", float64(m.DailyComputePrice)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateDailyNetworkPrice(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dailyNetworkPrice", "body", float64(m.DailyNetworkPrice)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateDailyStoragePrice(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dailyStoragePrice", "body", float64(m.DailyStoragePrice)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateDailyTotalPrice(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dailyTotalPrice", "body", float64(m.DailyTotalPrice)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateResourceUpfrontPriceID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "resourceUpfrontPriceId", "body", string(m.ResourceUpfrontPriceID)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateStatusDetails(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "statusDetails", "body", string(m.StatusDetails)); err != nil {
		return err
	}

	return nil
}

func (m *CatalogItemResourceUpfrontPriceResponse) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "unit", "body", string(m.Unit)); err != nil {
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
