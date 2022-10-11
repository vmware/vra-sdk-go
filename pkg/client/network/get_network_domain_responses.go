// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetNetworkDomainReader is a Reader for the GetNetworkDomain structure.
type GetNetworkDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkDomainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetNetworkDomainForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNetworkDomainNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkDomainOK creates a GetNetworkDomainOK with default headers values
func NewGetNetworkDomainOK() *GetNetworkDomainOK {
	return &GetNetworkDomainOK{}
}

/*
GetNetworkDomainOK describes a response with status code 200, with default header values.

successful operation
*/
type GetNetworkDomainOK struct {
	Payload *models.NetworkDomain
}

// IsSuccess returns true when this get network domain o k response has a 2xx status code
func (o *GetNetworkDomainOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network domain o k response has a 3xx status code
func (o *GetNetworkDomainOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network domain o k response has a 4xx status code
func (o *GetNetworkDomainOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network domain o k response has a 5xx status code
func (o *GetNetworkDomainOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network domain o k response a status code equal to that given
func (o *GetNetworkDomainOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkDomainOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains/{id}][%d] getNetworkDomainOK  %+v", 200, o.Payload)
}

func (o *GetNetworkDomainOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains/{id}][%d] getNetworkDomainOK  %+v", 200, o.Payload)
}

func (o *GetNetworkDomainOK) GetPayload() *models.NetworkDomain {
	return o.Payload
}

func (o *GetNetworkDomainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDomainForbidden creates a GetNetworkDomainForbidden with default headers values
func NewGetNetworkDomainForbidden() *GetNetworkDomainForbidden {
	return &GetNetworkDomainForbidden{}
}

/*
GetNetworkDomainForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNetworkDomainForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get network domain forbidden response has a 2xx status code
func (o *GetNetworkDomainForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network domain forbidden response has a 3xx status code
func (o *GetNetworkDomainForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network domain forbidden response has a 4xx status code
func (o *GetNetworkDomainForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network domain forbidden response has a 5xx status code
func (o *GetNetworkDomainForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get network domain forbidden response a status code equal to that given
func (o *GetNetworkDomainForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetNetworkDomainForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains/{id}][%d] getNetworkDomainForbidden  %+v", 403, o.Payload)
}

func (o *GetNetworkDomainForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains/{id}][%d] getNetworkDomainForbidden  %+v", 403, o.Payload)
}

func (o *GetNetworkDomainForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetNetworkDomainForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDomainNotFound creates a GetNetworkDomainNotFound with default headers values
func NewGetNetworkDomainNotFound() *GetNetworkDomainNotFound {
	return &GetNetworkDomainNotFound{}
}

/*
GetNetworkDomainNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetNetworkDomainNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get network domain not found response has a 2xx status code
func (o *GetNetworkDomainNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get network domain not found response has a 3xx status code
func (o *GetNetworkDomainNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network domain not found response has a 4xx status code
func (o *GetNetworkDomainNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get network domain not found response has a 5xx status code
func (o *GetNetworkDomainNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get network domain not found response a status code equal to that given
func (o *GetNetworkDomainNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetNetworkDomainNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains/{id}][%d] getNetworkDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworkDomainNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains/{id}][%d] getNetworkDomainNotFound  %+v", 404, o.Payload)
}

func (o *GetNetworkDomainNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworkDomainNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
