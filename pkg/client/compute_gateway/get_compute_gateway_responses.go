// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetComputeGatewayReader is a Reader for the GetComputeGateway structure.
type GetComputeGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeGatewayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComputeGatewayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComputeGatewayNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComputeGatewayOK creates a GetComputeGatewayOK with default headers values
func NewGetComputeGatewayOK() *GetComputeGatewayOK {
	return &GetComputeGatewayOK{}
}

/*
GetComputeGatewayOK describes a response with status code 200, with default header values.

successful operation
*/
type GetComputeGatewayOK struct {
	Payload *models.ComputeGateway
}

// IsSuccess returns true when this get compute gateway o k response has a 2xx status code
func (o *GetComputeGatewayOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compute gateway o k response has a 3xx status code
func (o *GetComputeGatewayOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compute gateway o k response has a 4xx status code
func (o *GetComputeGatewayOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compute gateway o k response has a 5xx status code
func (o *GetComputeGatewayOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compute gateway o k response a status code equal to that given
func (o *GetComputeGatewayOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetComputeGatewayOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getComputeGatewayOK  %+v", 200, o.Payload)
}

func (o *GetComputeGatewayOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getComputeGatewayOK  %+v", 200, o.Payload)
}

func (o *GetComputeGatewayOK) GetPayload() *models.ComputeGateway {
	return o.Payload
}

func (o *GetComputeGatewayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeGatewayForbidden creates a GetComputeGatewayForbidden with default headers values
func NewGetComputeGatewayForbidden() *GetComputeGatewayForbidden {
	return &GetComputeGatewayForbidden{}
}

/*
GetComputeGatewayForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetComputeGatewayForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get compute gateway forbidden response has a 2xx status code
func (o *GetComputeGatewayForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get compute gateway forbidden response has a 3xx status code
func (o *GetComputeGatewayForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compute gateway forbidden response has a 4xx status code
func (o *GetComputeGatewayForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get compute gateway forbidden response has a 5xx status code
func (o *GetComputeGatewayForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get compute gateway forbidden response a status code equal to that given
func (o *GetComputeGatewayForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetComputeGatewayForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getComputeGatewayForbidden  %+v", 403, o.Payload)
}

func (o *GetComputeGatewayForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getComputeGatewayForbidden  %+v", 403, o.Payload)
}

func (o *GetComputeGatewayForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetComputeGatewayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeGatewayNotFound creates a GetComputeGatewayNotFound with default headers values
func NewGetComputeGatewayNotFound() *GetComputeGatewayNotFound {
	return &GetComputeGatewayNotFound{}
}

/*
GetComputeGatewayNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetComputeGatewayNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get compute gateway not found response has a 2xx status code
func (o *GetComputeGatewayNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get compute gateway not found response has a 3xx status code
func (o *GetComputeGatewayNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compute gateway not found response has a 4xx status code
func (o *GetComputeGatewayNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get compute gateway not found response has a 5xx status code
func (o *GetComputeGatewayNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get compute gateway not found response a status code equal to that given
func (o *GetComputeGatewayNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetComputeGatewayNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getComputeGatewayNotFound  %+v", 404, o.Payload)
}

func (o *GetComputeGatewayNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/compute-gateways/{id}][%d] getComputeGatewayNotFound  %+v", 404, o.Payload)
}

func (o *GetComputeGatewayNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetComputeGatewayNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
