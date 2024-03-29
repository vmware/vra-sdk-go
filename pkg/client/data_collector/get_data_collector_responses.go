// Code generated by go-swagger; DO NOT EDIT.

package data_collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDataCollectorReader is a Reader for the GetDataCollector structure.
type GetDataCollectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataCollectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataCollectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetDataCollectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDataCollectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDataCollectorOK creates a GetDataCollectorOK with default headers values
func NewGetDataCollectorOK() *GetDataCollectorOK {
	return &GetDataCollectorOK{}
}

/*
GetDataCollectorOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDataCollectorOK struct {
	Payload *models.DataCollector
}

// IsSuccess returns true when this get data collector o k response has a 2xx status code
func (o *GetDataCollectorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get data collector o k response has a 3xx status code
func (o *GetDataCollectorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data collector o k response has a 4xx status code
func (o *GetDataCollectorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data collector o k response has a 5xx status code
func (o *GetDataCollectorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get data collector o k response a status code equal to that given
func (o *GetDataCollectorOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDataCollectorOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors/{id}][%d] getDataCollectorOK  %+v", 200, o.Payload)
}

func (o *GetDataCollectorOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors/{id}][%d] getDataCollectorOK  %+v", 200, o.Payload)
}

func (o *GetDataCollectorOK) GetPayload() *models.DataCollector {
	return o.Payload
}

func (o *GetDataCollectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataCollector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataCollectorForbidden creates a GetDataCollectorForbidden with default headers values
func NewGetDataCollectorForbidden() *GetDataCollectorForbidden {
	return &GetDataCollectorForbidden{}
}

/*
GetDataCollectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDataCollectorForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get data collector forbidden response has a 2xx status code
func (o *GetDataCollectorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data collector forbidden response has a 3xx status code
func (o *GetDataCollectorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data collector forbidden response has a 4xx status code
func (o *GetDataCollectorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data collector forbidden response has a 5xx status code
func (o *GetDataCollectorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get data collector forbidden response a status code equal to that given
func (o *GetDataCollectorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDataCollectorForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors/{id}][%d] getDataCollectorForbidden  %+v", 403, o.Payload)
}

func (o *GetDataCollectorForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors/{id}][%d] getDataCollectorForbidden  %+v", 403, o.Payload)
}

func (o *GetDataCollectorForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetDataCollectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataCollectorNotFound creates a GetDataCollectorNotFound with default headers values
func NewGetDataCollectorNotFound() *GetDataCollectorNotFound {
	return &GetDataCollectorNotFound{}
}

/*
GetDataCollectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDataCollectorNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get data collector not found response has a 2xx status code
func (o *GetDataCollectorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get data collector not found response has a 3xx status code
func (o *GetDataCollectorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data collector not found response has a 4xx status code
func (o *GetDataCollectorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get data collector not found response has a 5xx status code
func (o *GetDataCollectorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get data collector not found response a status code equal to that given
func (o *GetDataCollectorNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDataCollectorNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors/{id}][%d] getDataCollectorNotFound  %+v", 404, o.Payload)
}

func (o *GetDataCollectorNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors/{id}][%d] getDataCollectorNotFound  %+v", 404, o.Payload)
}

func (o *GetDataCollectorNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDataCollectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
