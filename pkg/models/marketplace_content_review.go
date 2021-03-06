// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MarketplaceContentReview MarketplaceContentReview
//
// swagger:model MarketplaceContentReview
type MarketplaceContentReview struct {

	// Comment
	Comment string `json:"comment,omitempty"`

	// Content ID
	ContentID string `json:"contentId,omitempty"`

	// Review ID
	ID string `json:"id,omitempty"`

	// Rating
	Rating string `json:"rating,omitempty"`

	// Title
	Title string `json:"title,omitempty"`

	// User
	User string `json:"user,omitempty"`
}

// Validate validates this marketplace content review
func (m *MarketplaceContentReview) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceContentReview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceContentReview) UnmarshalBinary(b []byte) error {
	var res MarketplaceContentReview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
