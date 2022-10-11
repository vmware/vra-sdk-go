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

// DeleteRequestReader is a Reader for the DeleteRequest structure.
type DeleteRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRequestNoContent creates a DeleteRequestNoContent with default headers values
func NewDeleteRequestNoContent() *DeleteRequestNoContent {
	return &DeleteRequestNoContent{}
}

/*
DeleteRequestNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteRequestNoContent struct {
}

// IsSuccess returns true when this delete request no content response has a 2xx status code
func (o *DeleteRequestNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete request no content response has a 3xx status code
func (o *DeleteRequestNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete request no content response has a 4xx status code
func (o *DeleteRequestNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete request no content response has a 5xx status code
func (o *DeleteRequestNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete request no content response a status code equal to that given
func (o *DeleteRequestNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteRequestNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/request-tracker/{id}][%d] deleteRequestNoContent ", 204)
}

func (o *DeleteRequestNoContent) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/request-tracker/{id}][%d] deleteRequestNoContent ", 204)
}

func (o *DeleteRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRequestForbidden creates a DeleteRequestForbidden with default headers values
func NewDeleteRequestForbidden() *DeleteRequestForbidden {
	return &DeleteRequestForbidden{}
}

/*
DeleteRequestForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRequestForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this delete request forbidden response has a 2xx status code
func (o *DeleteRequestForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete request forbidden response has a 3xx status code
func (o *DeleteRequestForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete request forbidden response has a 4xx status code
func (o *DeleteRequestForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete request forbidden response has a 5xx status code
func (o *DeleteRequestForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete request forbidden response a status code equal to that given
func (o *DeleteRequestForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRequestForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/request-tracker/{id}][%d] deleteRequestForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRequestForbidden) String() string {
	return fmt.Sprintf("[DELETE /iaas/api/request-tracker/{id}][%d] deleteRequestForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRequestForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
