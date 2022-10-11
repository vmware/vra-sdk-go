// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetExternalNetworkIPRangeReader is a Reader for the GetExternalNetworkIPRange structure.
type GetExternalNetworkIPRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalNetworkIPRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalNetworkIPRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetExternalNetworkIPRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalNetworkIPRangeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalNetworkIPRangeOK creates a GetExternalNetworkIPRangeOK with default headers values
func NewGetExternalNetworkIPRangeOK() *GetExternalNetworkIPRangeOK {
	return &GetExternalNetworkIPRangeOK{}
}

/*
GetExternalNetworkIPRangeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalNetworkIPRangeOK struct {
	Payload *models.ExternalNetworkIPRange
}

// IsSuccess returns true when this get external network Ip range o k response has a 2xx status code
func (o *GetExternalNetworkIPRangeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get external network Ip range o k response has a 3xx status code
func (o *GetExternalNetworkIPRangeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external network Ip range o k response has a 4xx status code
func (o *GetExternalNetworkIPRangeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get external network Ip range o k response has a 5xx status code
func (o *GetExternalNetworkIPRangeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get external network Ip range o k response a status code equal to that given
func (o *GetExternalNetworkIPRangeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalNetworkIPRangeOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/external-network-ip-ranges/{id}][%d] getExternalNetworkIpRangeOK  %+v", 200, o.Payload)
}

func (o *GetExternalNetworkIPRangeOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/external-network-ip-ranges/{id}][%d] getExternalNetworkIpRangeOK  %+v", 200, o.Payload)
}

func (o *GetExternalNetworkIPRangeOK) GetPayload() *models.ExternalNetworkIPRange {
	return o.Payload
}

func (o *GetExternalNetworkIPRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalNetworkIPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalNetworkIPRangeForbidden creates a GetExternalNetworkIPRangeForbidden with default headers values
func NewGetExternalNetworkIPRangeForbidden() *GetExternalNetworkIPRangeForbidden {
	return &GetExternalNetworkIPRangeForbidden{}
}

/*
GetExternalNetworkIPRangeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetExternalNetworkIPRangeForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get external network Ip range forbidden response has a 2xx status code
func (o *GetExternalNetworkIPRangeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external network Ip range forbidden response has a 3xx status code
func (o *GetExternalNetworkIPRangeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external network Ip range forbidden response has a 4xx status code
func (o *GetExternalNetworkIPRangeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external network Ip range forbidden response has a 5xx status code
func (o *GetExternalNetworkIPRangeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get external network Ip range forbidden response a status code equal to that given
func (o *GetExternalNetworkIPRangeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalNetworkIPRangeForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/external-network-ip-ranges/{id}][%d] getExternalNetworkIpRangeForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalNetworkIPRangeForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/external-network-ip-ranges/{id}][%d] getExternalNetworkIpRangeForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalNetworkIPRangeForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetExternalNetworkIPRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalNetworkIPRangeNotFound creates a GetExternalNetworkIPRangeNotFound with default headers values
func NewGetExternalNetworkIPRangeNotFound() *GetExternalNetworkIPRangeNotFound {
	return &GetExternalNetworkIPRangeNotFound{}
}

/*
GetExternalNetworkIPRangeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetExternalNetworkIPRangeNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get external network Ip range not found response has a 2xx status code
func (o *GetExternalNetworkIPRangeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get external network Ip range not found response has a 3xx status code
func (o *GetExternalNetworkIPRangeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external network Ip range not found response has a 4xx status code
func (o *GetExternalNetworkIPRangeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get external network Ip range not found response has a 5xx status code
func (o *GetExternalNetworkIPRangeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get external network Ip range not found response a status code equal to that given
func (o *GetExternalNetworkIPRangeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalNetworkIPRangeNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/external-network-ip-ranges/{id}][%d] getExternalNetworkIpRangeNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalNetworkIPRangeNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/external-network-ip-ranges/{id}][%d] getExternalNetworkIpRangeNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalNetworkIPRangeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExternalNetworkIPRangeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
