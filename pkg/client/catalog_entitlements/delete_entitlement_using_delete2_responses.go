// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEntitlementUsingDELETE2Reader is a Reader for the DeleteEntitlementUsingDELETE2 structure.
type DeleteEntitlementUsingDELETE2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEntitlementUsingDELETE2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEntitlementUsingDELETE2NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEntitlementUsingDELETE2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEntitlementUsingDELETE2NoContent creates a DeleteEntitlementUsingDELETE2NoContent with default headers values
func NewDeleteEntitlementUsingDELETE2NoContent() *DeleteEntitlementUsingDELETE2NoContent {
	return &DeleteEntitlementUsingDELETE2NoContent{}
}

/* DeleteEntitlementUsingDELETE2NoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteEntitlementUsingDELETE2NoContent struct {
}

func (o *DeleteEntitlementUsingDELETE2NoContent) Error() string {
	return fmt.Sprintf("[DELETE /catalog/api/admin/entitlements/{id}][%d] deleteEntitlementUsingDELETE2NoContent ", 204)
}

func (o *DeleteEntitlementUsingDELETE2NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEntitlementUsingDELETE2Unauthorized creates a DeleteEntitlementUsingDELETE2Unauthorized with default headers values
func NewDeleteEntitlementUsingDELETE2Unauthorized() *DeleteEntitlementUsingDELETE2Unauthorized {
	return &DeleteEntitlementUsingDELETE2Unauthorized{}
}

/* DeleteEntitlementUsingDELETE2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteEntitlementUsingDELETE2Unauthorized struct {
}

func (o *DeleteEntitlementUsingDELETE2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /catalog/api/admin/entitlements/{id}][%d] deleteEntitlementUsingDELETE2Unauthorized ", 401)
}

func (o *DeleteEntitlementUsingDELETE2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
