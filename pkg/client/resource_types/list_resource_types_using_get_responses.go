// Code generated by go-swagger; DO NOT EDIT.

package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// ListResourceTypesUsingGETReader is a Reader for the ListResourceTypesUsingGET structure.
type ListResourceTypesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResourceTypesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListResourceTypesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListResourceTypesUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListResourceTypesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListResourceTypesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListResourceTypesUsingGETOK creates a ListResourceTypesUsingGETOK with default headers values
func NewListResourceTypesUsingGETOK() *ListResourceTypesUsingGETOK {
	return &ListResourceTypesUsingGETOK{}
}

/*ListResourceTypesUsingGETOK handles this case with default header values.

OK
*/
type ListResourceTypesUsingGETOK struct {
	Payload *models.PageOfResourceType
}

func (o *ListResourceTypesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/resource-types][%d] listResourceTypesUsingGETOK  %+v", 200, o.Payload)
}

func (o *ListResourceTypesUsingGETOK) GetPayload() *models.PageOfResourceType {
	return o.Payload
}

func (o *ListResourceTypesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfResourceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourceTypesUsingGETBadRequest creates a ListResourceTypesUsingGETBadRequest with default headers values
func NewListResourceTypesUsingGETBadRequest() *ListResourceTypesUsingGETBadRequest {
	return &ListResourceTypesUsingGETBadRequest{}
}

/*ListResourceTypesUsingGETBadRequest handles this case with default header values.

Bad Request
*/
type ListResourceTypesUsingGETBadRequest struct {
}

func (o *ListResourceTypesUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/resource-types][%d] listResourceTypesUsingGETBadRequest ", 400)
}

func (o *ListResourceTypesUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListResourceTypesUsingGETUnauthorized creates a ListResourceTypesUsingGETUnauthorized with default headers values
func NewListResourceTypesUsingGETUnauthorized() *ListResourceTypesUsingGETUnauthorized {
	return &ListResourceTypesUsingGETUnauthorized{}
}

/*ListResourceTypesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ListResourceTypesUsingGETUnauthorized struct {
}

func (o *ListResourceTypesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/resource-types][%d] listResourceTypesUsingGETUnauthorized ", 401)
}

func (o *ListResourceTypesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListResourceTypesUsingGETForbidden creates a ListResourceTypesUsingGETForbidden with default headers values
func NewListResourceTypesUsingGETForbidden() *ListResourceTypesUsingGETForbidden {
	return &ListResourceTypesUsingGETForbidden{}
}

/*ListResourceTypesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ListResourceTypesUsingGETForbidden struct {
}

func (o *ListResourceTypesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/resource-types][%d] listResourceTypesUsingGETForbidden ", 403)
}

func (o *ListResourceTypesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
