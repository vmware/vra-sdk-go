// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ListResourceRequestsUsingGET2Reader is a Reader for the ListResourceRequestsUsingGET2 structure.
type ListResourceRequestsUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResourceRequestsUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListResourceRequestsUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListResourceRequestsUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListResourceRequestsUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListResourceRequestsUsingGET2OK creates a ListResourceRequestsUsingGET2OK with default headers values
func NewListResourceRequestsUsingGET2OK() *ListResourceRequestsUsingGET2OK {
	return &ListResourceRequestsUsingGET2OK{}
}

/*
ListResourceRequestsUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type ListResourceRequestsUsingGET2OK struct {
	Payload *models.PageOfRequest
}

// IsSuccess returns true when this list resource requests using g e t2 o k response has a 2xx status code
func (o *ListResourceRequestsUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list resource requests using g e t2 o k response has a 3xx status code
func (o *ListResourceRequestsUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list resource requests using g e t2 o k response has a 4xx status code
func (o *ListResourceRequestsUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list resource requests using g e t2 o k response has a 5xx status code
func (o *ListResourceRequestsUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list resource requests using g e t2 o k response a status code equal to that given
func (o *ListResourceRequestsUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListResourceRequestsUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/requests][%d] listResourceRequestsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *ListResourceRequestsUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/requests][%d] listResourceRequestsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *ListResourceRequestsUsingGET2OK) GetPayload() *models.PageOfRequest {
	return o.Payload
}

func (o *ListResourceRequestsUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResourceRequestsUsingGET2Unauthorized creates a ListResourceRequestsUsingGET2Unauthorized with default headers values
func NewListResourceRequestsUsingGET2Unauthorized() *ListResourceRequestsUsingGET2Unauthorized {
	return &ListResourceRequestsUsingGET2Unauthorized{}
}

/*
ListResourceRequestsUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListResourceRequestsUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this list resource requests using g e t2 unauthorized response has a 2xx status code
func (o *ListResourceRequestsUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list resource requests using g e t2 unauthorized response has a 3xx status code
func (o *ListResourceRequestsUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list resource requests using g e t2 unauthorized response has a 4xx status code
func (o *ListResourceRequestsUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list resource requests using g e t2 unauthorized response has a 5xx status code
func (o *ListResourceRequestsUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list resource requests using g e t2 unauthorized response a status code equal to that given
func (o *ListResourceRequestsUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListResourceRequestsUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/requests][%d] listResourceRequestsUsingGET2Unauthorized ", 401)
}

func (o *ListResourceRequestsUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/requests][%d] listResourceRequestsUsingGET2Unauthorized ", 401)
}

func (o *ListResourceRequestsUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListResourceRequestsUsingGET2NotFound creates a ListResourceRequestsUsingGET2NotFound with default headers values
func NewListResourceRequestsUsingGET2NotFound() *ListResourceRequestsUsingGET2NotFound {
	return &ListResourceRequestsUsingGET2NotFound{}
}

/*
ListResourceRequestsUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListResourceRequestsUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this list resource requests using g e t2 not found response has a 2xx status code
func (o *ListResourceRequestsUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list resource requests using g e t2 not found response has a 3xx status code
func (o *ListResourceRequestsUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list resource requests using g e t2 not found response has a 4xx status code
func (o *ListResourceRequestsUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list resource requests using g e t2 not found response has a 5xx status code
func (o *ListResourceRequestsUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list resource requests using g e t2 not found response a status code equal to that given
func (o *ListResourceRequestsUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListResourceRequestsUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/requests][%d] listResourceRequestsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *ListResourceRequestsUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/requests][%d] listResourceRequestsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *ListResourceRequestsUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListResourceRequestsUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
