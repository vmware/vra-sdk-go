// Code generated by go-swagger; DO NOT EDIT.

package network_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteNetworkProfileReader is a Reader for the DeleteNetworkProfile structure.
type DeleteNetworkProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteNetworkProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkProfileNoContent creates a DeleteNetworkProfileNoContent with default headers values
func NewDeleteNetworkProfileNoContent() *DeleteNetworkProfileNoContent {
	return &DeleteNetworkProfileNoContent{}
}

/*
DeleteNetworkProfileNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteNetworkProfileNoContent struct {
}

// IsSuccess returns true when this delete network profile no content response has a 2xx status code
func (o *DeleteNetworkProfileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete network profile no content response has a 3xx status code
func (o *DeleteNetworkProfileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network profile no content response has a 4xx status code
func (o *DeleteNetworkProfileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete network profile no content response has a 5xx status code
func (o *DeleteNetworkProfileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network profile no content response a status code equal to that given
func (o *DeleteNetworkProfileNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteNetworkProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-profiles/{id}][%d] deleteNetworkProfileNoContent ", 204)
}

func (o *DeleteNetworkProfileNoContent) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-profiles/{id}][%d] deleteNetworkProfileNoContent ", 204)
}

func (o *DeleteNetworkProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkProfileForbidden creates a DeleteNetworkProfileForbidden with default headers values
func NewDeleteNetworkProfileForbidden() *DeleteNetworkProfileForbidden {
	return &DeleteNetworkProfileForbidden{}
}

/*
DeleteNetworkProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteNetworkProfileForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this delete network profile forbidden response has a 2xx status code
func (o *DeleteNetworkProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete network profile forbidden response has a 3xx status code
func (o *DeleteNetworkProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete network profile forbidden response has a 4xx status code
func (o *DeleteNetworkProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete network profile forbidden response has a 5xx status code
func (o *DeleteNetworkProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete network profile forbidden response a status code equal to that given
func (o *DeleteNetworkProfileForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteNetworkProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-profiles/{id}][%d] deleteNetworkProfileForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNetworkProfileForbidden) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/network-profiles/{id}][%d] deleteNetworkProfileForbidden  %+v", 403, o.Payload)
}

func (o *DeleteNetworkProfileForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteNetworkProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
