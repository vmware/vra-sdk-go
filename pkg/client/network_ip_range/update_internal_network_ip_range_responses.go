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

// UpdateInternalNetworkIPRangeReader is a Reader for the UpdateInternalNetworkIPRange structure.
type UpdateInternalNetworkIPRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInternalNetworkIPRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInternalNetworkIPRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateInternalNetworkIPRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateInternalNetworkIPRangeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateInternalNetworkIPRangeOK creates a UpdateInternalNetworkIPRangeOK with default headers values
func NewUpdateInternalNetworkIPRangeOK() *UpdateInternalNetworkIPRangeOK {
	return &UpdateInternalNetworkIPRangeOK{}
}

/*
UpdateInternalNetworkIPRangeOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateInternalNetworkIPRangeOK struct {
	Payload *models.NetworkIPRange
}

// IsSuccess returns true when this update internal network Ip range o k response has a 2xx status code
func (o *UpdateInternalNetworkIPRangeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update internal network Ip range o k response has a 3xx status code
func (o *UpdateInternalNetworkIPRangeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update internal network Ip range o k response has a 4xx status code
func (o *UpdateInternalNetworkIPRangeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update internal network Ip range o k response has a 5xx status code
func (o *UpdateInternalNetworkIPRangeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update internal network Ip range o k response a status code equal to that given
func (o *UpdateInternalNetworkIPRangeOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateInternalNetworkIPRangeOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeOK  %+v", 200, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeOK) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeOK  %+v", 200, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeOK) GetPayload() *models.NetworkIPRange {
	return o.Payload
}

func (o *UpdateInternalNetworkIPRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkIPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInternalNetworkIPRangeForbidden creates a UpdateInternalNetworkIPRangeForbidden with default headers values
func NewUpdateInternalNetworkIPRangeForbidden() *UpdateInternalNetworkIPRangeForbidden {
	return &UpdateInternalNetworkIPRangeForbidden{}
}

/*
UpdateInternalNetworkIPRangeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateInternalNetworkIPRangeForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this update internal network Ip range forbidden response has a 2xx status code
func (o *UpdateInternalNetworkIPRangeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update internal network Ip range forbidden response has a 3xx status code
func (o *UpdateInternalNetworkIPRangeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update internal network Ip range forbidden response has a 4xx status code
func (o *UpdateInternalNetworkIPRangeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update internal network Ip range forbidden response has a 5xx status code
func (o *UpdateInternalNetworkIPRangeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update internal network Ip range forbidden response a status code equal to that given
func (o *UpdateInternalNetworkIPRangeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateInternalNetworkIPRangeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeForbidden) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeForbidden  %+v", 403, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateInternalNetworkIPRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInternalNetworkIPRangeNotFound creates a UpdateInternalNetworkIPRangeNotFound with default headers values
func NewUpdateInternalNetworkIPRangeNotFound() *UpdateInternalNetworkIPRangeNotFound {
	return &UpdateInternalNetworkIPRangeNotFound{}
}

/*
UpdateInternalNetworkIPRangeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateInternalNetworkIPRangeNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update internal network Ip range not found response has a 2xx status code
func (o *UpdateInternalNetworkIPRangeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update internal network Ip range not found response has a 3xx status code
func (o *UpdateInternalNetworkIPRangeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update internal network Ip range not found response has a 4xx status code
func (o *UpdateInternalNetworkIPRangeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update internal network Ip range not found response has a 5xx status code
func (o *UpdateInternalNetworkIPRangeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update internal network Ip range not found response a status code equal to that given
func (o *UpdateInternalNetworkIPRangeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateInternalNetworkIPRangeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeNotFound) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-ip-ranges/{id}][%d] updateInternalNetworkIpRangeNotFound  %+v", 404, o.Payload)
}

func (o *UpdateInternalNetworkIPRangeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInternalNetworkIPRangeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
