// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObjectMeta ObjectMeta
//
// swagger:model ObjectMeta
type ObjectMeta struct {

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// creation timestamp
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// deletion timestamp
	DeletionTimestamp string `json:"deletionTimestamp,omitempty"`

	// generate name
	GenerateName string `json:"generateName,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// self link
	SelfLink string `json:"selfLink,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this object meta
func (m *ObjectMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this object meta based on context it is used
func (m *ObjectMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectMeta) UnmarshalBinary(b []byte) error {
	var res ObjectMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
