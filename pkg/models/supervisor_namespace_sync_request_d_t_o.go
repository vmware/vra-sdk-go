// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SupervisorNamespaceSyncRequestDTO SupervisorNamespaceSyncRequestDTO
//
// swagger:model SupervisorNamespaceSyncRequestDTO
type SupervisorNamespaceSyncRequestDTO struct {

	// edit groups
	EditGroups string `json:"editGroups,omitempty"`

	// edit users
	EditUsers string `json:"editUsers,omitempty"`

	// owner groups
	OwnerGroups string `json:"ownerGroups,omitempty"`

	// owner users
	OwnerUsers string `json:"ownerUsers,omitempty"`

	// view groups
	ViewGroups string `json:"viewGroups,omitempty"`

	// view users
	ViewUsers string `json:"viewUsers,omitempty"`
}

// Validate validates this supervisor namespace sync request d t o
func (m *SupervisorNamespaceSyncRequestDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this supervisor namespace sync request d t o based on context it is used
func (m *SupervisorNamespaceSyncRequestDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupervisorNamespaceSyncRequestDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupervisorNamespaceSyncRequestDTO) UnmarshalBinary(b []byte) error {
	var res SupervisorNamespaceSyncRequestDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
