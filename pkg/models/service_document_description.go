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

// ServiceDocumentDescription service document description
// swagger:model ServiceDocumentDescription
type ServiceDocumentDescription struct {

	// description
	Description string `json:"description,omitempty"`

	// document indexing options
	DocumentIndexingOptions []string `json:"documentIndexingOptions"`

	// name
	Name string `json:"name,omitempty"`

	// property descriptions
	PropertyDescriptions map[string]PropertyDescription `json:"propertyDescriptions,omitempty"`

	// serialized state size limit
	SerializedStateSizeLimit int32 `json:"serializedStateSizeLimit,omitempty"`

	// service capabilities
	ServiceCapabilities []string `json:"serviceCapabilities"`

	// service request routes
	ServiceRequestRoutes map[string][]Route `json:"serviceRequestRoutes,omitempty"`

	// user interface resource path
	UserInterfaceResourcePath string `json:"userInterfaceResourcePath,omitempty"`

	// version retention floor
	VersionRetentionFloor int64 `json:"versionRetentionFloor,omitempty"`

	// version retention limit
	VersionRetentionLimit int64 `json:"versionRetentionLimit,omitempty"`
}

// Validate validates this service document description
func (m *ServiceDocumentDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocumentIndexingOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyDescriptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRequestRoutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceDocumentDescriptionDocumentIndexingOptionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INDEX_METADATA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDocumentDescriptionDocumentIndexingOptionsItemsEnum = append(serviceDocumentDescriptionDocumentIndexingOptionsItemsEnum, v)
	}
}

func (m *ServiceDocumentDescription) validateDocumentIndexingOptionsItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDocumentDescriptionDocumentIndexingOptionsItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDocumentDescription) validateDocumentIndexingOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.DocumentIndexingOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.DocumentIndexingOptions); i++ {

		// value enum
		if err := m.validateDocumentIndexingOptionsItemsEnum("documentIndexingOptions"+"."+strconv.Itoa(i), "body", m.DocumentIndexingOptions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ServiceDocumentDescription) validatePropertyDescriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyDescriptions) { // not required
		return nil
	}

	for k := range m.PropertyDescriptions {

		if err := validate.Required("propertyDescriptions"+"."+k, "body", m.PropertyDescriptions[k]); err != nil {
			return err
		}
		if val, ok := m.PropertyDescriptions[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var serviceDocumentDescriptionServiceCapabilitiesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INSTRUMENTATION","PERIODIC_MAINTENANCE","PERSISTENCE","REPLICATION","OWNER_SELECTION","STRICT_UPDATE_CHECKING","HTML_USER_INTERFACE","CONCURRENT_UPDATE_HANDLING","CONCURRENT_GET_HANDLING","IDEMPOTENT_POST","ON_DEMAND_LOAD","IMMUTABLE","WRAP_ERROR_RESPONSE","LIFO_QUEUE","URI_NAMESPACE_OWNER","CORE","UTILITY","FACTORY","FACTORY_ITEM","STATELESS","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDocumentDescriptionServiceCapabilitiesItemsEnum = append(serviceDocumentDescriptionServiceCapabilitiesItemsEnum, v)
	}
}

func (m *ServiceDocumentDescription) validateServiceCapabilitiesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDocumentDescriptionServiceCapabilitiesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDocumentDescription) validateServiceCapabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceCapabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceCapabilities); i++ {

		// value enum
		if err := m.validateServiceCapabilitiesItemsEnum("serviceCapabilities"+"."+strconv.Itoa(i), "body", m.ServiceCapabilities[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ServiceDocumentDescription) validateServiceRequestRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceRequestRoutes) { // not required
		return nil
	}

	for k := range m.ServiceRequestRoutes {

		if err := validate.Required("serviceRequestRoutes"+"."+k, "body", m.ServiceRequestRoutes[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.ServiceRequestRoutes[k]); i++ {

			if err := m.ServiceRequestRoutes[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceRequestRoutes" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDocumentDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDocumentDescription) UnmarshalBinary(b []byte) error {
	var res ServiceDocumentDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
