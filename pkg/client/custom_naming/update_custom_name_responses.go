// Code generated by go-swagger; DO NOT EDIT.

package custom_naming

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateCustomNameReader is a Reader for the UpdateCustomName structure.
type UpdateCustomNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateCustomNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCustomNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCustomNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCustomNameNoContent creates a UpdateCustomNameNoContent with default headers values
func NewUpdateCustomNameNoContent() *UpdateCustomNameNoContent {
	return &UpdateCustomNameNoContent{}
}

/* UpdateCustomNameNoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateCustomNameNoContent struct {
}

func (o *UpdateCustomNameNoContent) Error() string {
	return fmt.Sprintf("[PUT /iaas/api/naming][%d] updateCustomNameNoContent ", 204)
}

func (o *UpdateCustomNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCustomNameBadRequest creates a UpdateCustomNameBadRequest with default headers values
func NewUpdateCustomNameBadRequest() *UpdateCustomNameBadRequest {
	return &UpdateCustomNameBadRequest{}
}

/* UpdateCustomNameBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type UpdateCustomNameBadRequest struct {
	Payload *models.Error
}

func (o *UpdateCustomNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /iaas/api/naming][%d] updateCustomNameBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCustomNameBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomNameForbidden creates a UpdateCustomNameForbidden with default headers values
func NewUpdateCustomNameForbidden() *UpdateCustomNameForbidden {
	return &UpdateCustomNameForbidden{}
}

/* UpdateCustomNameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCustomNameForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *UpdateCustomNameForbidden) Error() string {
	return fmt.Sprintf("[PUT /iaas/api/naming][%d] updateCustomNameForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCustomNameForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateCustomNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomNameNotFound creates a UpdateCustomNameNotFound with default headers values
func NewUpdateCustomNameNotFound() *UpdateCustomNameNotFound {
	return &UpdateCustomNameNotFound{}
}

/* UpdateCustomNameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateCustomNameNotFound struct {
	Payload *models.Error
}

func (o *UpdateCustomNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /iaas/api/naming][%d] updateCustomNameNotFound  %+v", 404, o.Payload)
}
func (o *UpdateCustomNameNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateCustomNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
