// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceDocumentQueryResult service document query result
//
// swagger:model ServiceDocumentQueryResult
type ServiceDocumentQueryResult struct {

	// document count
	DocumentCount int64 `json:"documentCount,omitempty"`

	// document links
	DocumentLinks []string `json:"documentLinks"`

	// documents
	Documents map[string]string `json:"documents,omitempty"`

	// next page link
	NextPageLink string `json:"nextPageLink,omitempty"`

	// next page links per group
	NextPageLinksPerGroup map[string]string `json:"nextPageLinksPerGroup,omitempty"`

	// prev page link
	PrevPageLink string `json:"prevPageLink,omitempty"`

	// selected documents
	SelectedDocuments map[string]string `json:"selectedDocuments,omitempty"`

	// selected links
	SelectedLinks []string `json:"selectedLinks"`

	// selected links per document
	SelectedLinksPerDocument map[string]string `json:"selectedLinksPerDocument,omitempty"`
}

// Validate validates this service document query result
func (m *ServiceDocumentQueryResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service document query result based on context it is used
func (m *ServiceDocumentQueryResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDocumentQueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDocumentQueryResult) UnmarshalBinary(b []byte) error {
	var res ServiceDocumentQueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
