// Code generated by go-swagger; DO NOT EDIT.

package request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetRequestTrackerReader is a Reader for the GetRequestTracker structure.
type GetRequestTrackerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestTrackerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRequestTrackerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRequestTrackerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRequestTrackerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRequestTrackerOK creates a GetRequestTrackerOK with default headers values
func NewGetRequestTrackerOK() *GetRequestTrackerOK {
	return &GetRequestTrackerOK{}
}

/*
GetRequestTrackerOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRequestTrackerOK struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this get request tracker o k response has a 2xx status code
func (o *GetRequestTrackerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get request tracker o k response has a 3xx status code
func (o *GetRequestTrackerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get request tracker o k response has a 4xx status code
func (o *GetRequestTrackerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get request tracker o k response has a 5xx status code
func (o *GetRequestTrackerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get request tracker o k response a status code equal to that given
func (o *GetRequestTrackerOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRequestTrackerOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerOK  %+v", 200, o.Payload)
}

func (o *GetRequestTrackerOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerOK  %+v", 200, o.Payload)
}

func (o *GetRequestTrackerOK) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *GetRequestTrackerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestTrackerForbidden creates a GetRequestTrackerForbidden with default headers values
func NewGetRequestTrackerForbidden() *GetRequestTrackerForbidden {
	return &GetRequestTrackerForbidden{}
}

/*
GetRequestTrackerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRequestTrackerForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get request tracker forbidden response has a 2xx status code
func (o *GetRequestTrackerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get request tracker forbidden response has a 3xx status code
func (o *GetRequestTrackerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get request tracker forbidden response has a 4xx status code
func (o *GetRequestTrackerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get request tracker forbidden response has a 5xx status code
func (o *GetRequestTrackerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get request tracker forbidden response a status code equal to that given
func (o *GetRequestTrackerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRequestTrackerForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerForbidden  %+v", 403, o.Payload)
}

func (o *GetRequestTrackerForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerForbidden  %+v", 403, o.Payload)
}

func (o *GetRequestTrackerForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetRequestTrackerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestTrackerNotFound creates a GetRequestTrackerNotFound with default headers values
func NewGetRequestTrackerNotFound() *GetRequestTrackerNotFound {
	return &GetRequestTrackerNotFound{}
}

/*
GetRequestTrackerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRequestTrackerNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get request tracker not found response has a 2xx status code
func (o *GetRequestTrackerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get request tracker not found response has a 3xx status code
func (o *GetRequestTrackerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get request tracker not found response has a 4xx status code
func (o *GetRequestTrackerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get request tracker not found response has a 5xx status code
func (o *GetRequestTrackerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get request tracker not found response a status code equal to that given
func (o *GetRequestTrackerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRequestTrackerNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerNotFound  %+v", 404, o.Payload)
}

func (o *GetRequestTrackerNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerNotFound  %+v", 404, o.Payload)
}

func (o *GetRequestTrackerNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRequestTrackerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
