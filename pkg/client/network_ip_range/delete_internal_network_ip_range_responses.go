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

// DeleteInternalNetworkIPRangeReader is a Reader for the DeleteInternalNetworkIPRange structure.
type DeleteInternalNetworkIPRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInternalNetworkIPRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteInternalNetworkIPRangeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteInternalNetworkIPRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteInternalNetworkIPRangeNoContent creates a DeleteInternalNetworkIPRangeNoContent with default headers values
func NewDeleteInternalNetworkIPRangeNoContent() *DeleteInternalNetworkIPRangeNoContent {
	return &DeleteInternalNetworkIPRangeNoContent{}
}

/*
DeleteInternalNetworkIPRangeNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteInternalNetworkIPRangeNoContent struct {
}

// IsSuccess returns true when this delete internal network Ip range no content response has a 2xx status code
func (o *DeleteInternalNetworkIPRangeNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete internal network Ip range no content response has a 3xx status code
func (o *DeleteInternalNetworkIPRangeNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete internal network Ip range no content response has a 4xx status code
func (o *DeleteInternalNetworkIPRangeNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete internal network Ip range no content response has a 5xx status code
func (o *DeleteInternalNetworkIPRangeNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete internal network Ip range no content response a status code equal to that given
func (o *DeleteInternalNetworkIPRangeNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteInternalNetworkIPRangeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-ip-ranges/{id}][%d] deleteInternalNetworkIpRangeNoContent ", 204)
}

func (o *DeleteInternalNetworkIPRangeNoContent) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-ip-ranges/{id}][%d] deleteInternalNetworkIpRangeNoContent ", 204)
}

func (o *DeleteInternalNetworkIPRangeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInternalNetworkIPRangeForbidden creates a DeleteInternalNetworkIPRangeForbidden with default headers values
func NewDeleteInternalNetworkIPRangeForbidden() *DeleteInternalNetworkIPRangeForbidden {
	return &DeleteInternalNetworkIPRangeForbidden{}
}

/*
DeleteInternalNetworkIPRangeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteInternalNetworkIPRangeForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this delete internal network Ip range forbidden response has a 2xx status code
func (o *DeleteInternalNetworkIPRangeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete internal network Ip range forbidden response has a 3xx status code
func (o *DeleteInternalNetworkIPRangeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete internal network Ip range forbidden response has a 4xx status code
func (o *DeleteInternalNetworkIPRangeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete internal network Ip range forbidden response has a 5xx status code
func (o *DeleteInternalNetworkIPRangeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete internal network Ip range forbidden response a status code equal to that given
func (o *DeleteInternalNetworkIPRangeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteInternalNetworkIPRangeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-ip-ranges/{id}][%d] deleteInternalNetworkIpRangeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInternalNetworkIPRangeForbidden) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-ip-ranges/{id}][%d] deleteInternalNetworkIpRangeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInternalNetworkIPRangeForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteInternalNetworkIPRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
