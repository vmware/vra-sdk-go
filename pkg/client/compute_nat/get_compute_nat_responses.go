// Code generated by go-swagger; DO NOT EDIT.

package compute_nat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetComputeNatReader is a Reader for the GetComputeNat structure.
type GetComputeNatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeNatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeNatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComputeNatForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComputeNatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComputeNatOK creates a GetComputeNatOK with default headers values
func NewGetComputeNatOK() *GetComputeNatOK {
	return &GetComputeNatOK{}
}

/*
GetComputeNatOK describes a response with status code 200, with default header values.

successful operation
*/
type GetComputeNatOK struct {
	Payload *models.ComputeNat
}

// IsSuccess returns true when this get compute nat o k response has a 2xx status code
func (o *GetComputeNatOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compute nat o k response has a 3xx status code
func (o *GetComputeNatOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compute nat o k response has a 4xx status code
func (o *GetComputeNatOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compute nat o k response has a 5xx status code
func (o *GetComputeNatOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compute nat o k response a status code equal to that given
func (o *GetComputeNatOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetComputeNatOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats/{id}][%d] getComputeNatOK  %+v", 200, o.Payload)
}

func (o *GetComputeNatOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats/{id}][%d] getComputeNatOK  %+v", 200, o.Payload)
}

func (o *GetComputeNatOK) GetPayload() *models.ComputeNat {
	return o.Payload
}

func (o *GetComputeNatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeNat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeNatForbidden creates a GetComputeNatForbidden with default headers values
func NewGetComputeNatForbidden() *GetComputeNatForbidden {
	return &GetComputeNatForbidden{}
}

/*
GetComputeNatForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetComputeNatForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get compute nat forbidden response has a 2xx status code
func (o *GetComputeNatForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get compute nat forbidden response has a 3xx status code
func (o *GetComputeNatForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compute nat forbidden response has a 4xx status code
func (o *GetComputeNatForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get compute nat forbidden response has a 5xx status code
func (o *GetComputeNatForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get compute nat forbidden response a status code equal to that given
func (o *GetComputeNatForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetComputeNatForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats/{id}][%d] getComputeNatForbidden  %+v", 403, o.Payload)
}

func (o *GetComputeNatForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats/{id}][%d] getComputeNatForbidden  %+v", 403, o.Payload)
}

func (o *GetComputeNatForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetComputeNatForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeNatNotFound creates a GetComputeNatNotFound with default headers values
func NewGetComputeNatNotFound() *GetComputeNatNotFound {
	return &GetComputeNatNotFound{}
}

/*
GetComputeNatNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetComputeNatNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get compute nat not found response has a 2xx status code
func (o *GetComputeNatNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get compute nat not found response has a 3xx status code
func (o *GetComputeNatNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compute nat not found response has a 4xx status code
func (o *GetComputeNatNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get compute nat not found response has a 5xx status code
func (o *GetComputeNatNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get compute nat not found response a status code equal to that given
func (o *GetComputeNatNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetComputeNatNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats/{id}][%d] getComputeNatNotFound  %+v", 404, o.Payload)
}

func (o *GetComputeNatNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats/{id}][%d] getComputeNatNotFound  %+v", 404, o.Payload)
}

func (o *GetComputeNatNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetComputeNatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
