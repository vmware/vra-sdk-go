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

// DeploymentExpenseHistory DeploymentExpenseHistory
//
// Expense history of a deployment.
//
// swagger:model DeploymentExpenseHistory
type DeploymentExpenseHistory struct {

	// The currency code of the expense history.
	// Read Only: true
	Currency string `json:"currency,omitempty"`

	// The list of expense history metric.
	// Read Only: true
	Data []*TimeSeriesValue `json:"data"`

	// The requested interval type.
	// Read Only: true
	// Enum: [daily weekly monthly]
	Interval string `json:"interval,omitempty"`
}

// Validate validates this deployment expense history
func (m *DeploymentExpenseHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentExpenseHistory) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var deploymentExpenseHistoryTypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["daily","weekly","monthly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentExpenseHistoryTypeIntervalPropEnum = append(deploymentExpenseHistoryTypeIntervalPropEnum, v)
	}
}

const (

	// DeploymentExpenseHistoryIntervalDaily captures enum value "daily"
	DeploymentExpenseHistoryIntervalDaily string = "daily"

	// DeploymentExpenseHistoryIntervalWeekly captures enum value "weekly"
	DeploymentExpenseHistoryIntervalWeekly string = "weekly"

	// DeploymentExpenseHistoryIntervalMonthly captures enum value "monthly"
	DeploymentExpenseHistoryIntervalMonthly string = "monthly"
)

// prop value enum
func (m *DeploymentExpenseHistory) validateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentExpenseHistoryTypeIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentExpenseHistory) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	// value enum
	if err := m.validateIntervalEnum("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this deployment expense history based on the context it is used
func (m *DeploymentExpenseHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterval(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentExpenseHistory) contextValidateCurrency(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "currency", "body", string(m.Currency)); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentExpenseHistory) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "data", "body", []*TimeSeriesValue(m.Data)); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeploymentExpenseHistory) contextValidateInterval(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interval", "body", string(m.Interval)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentExpenseHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentExpenseHistory) UnmarshalBinary(b []byte) error {
	var res DeploymentExpenseHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
