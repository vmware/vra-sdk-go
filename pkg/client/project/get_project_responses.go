// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/*
GetProjectOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProjectOK struct {
	Payload *models.IaaSProject
}

// IsSuccess returns true when this get project o k response has a 2xx status code
func (o *GetProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project o k response has a 3xx status code
func (o *GetProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project o k response has a 4xx status code
func (o *GetProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project o k response has a 5xx status code
func (o *GetProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project o k response a status code equal to that given
func (o *GetProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) GetPayload() *models.IaaSProject {
	return o.Payload
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaaSProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectForbidden creates a GetProjectForbidden with default headers values
func NewGetProjectForbidden() *GetProjectForbidden {
	return &GetProjectForbidden{}
}

/*
GetProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get project forbidden response has a 2xx status code
func (o *GetProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project forbidden response has a 3xx status code
func (o *GetProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project forbidden response has a 4xx status code
func (o *GetProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project forbidden response has a 5xx status code
func (o *GetProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get project forbidden response a status code equal to that given
func (o *GetProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectNotFound creates a GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {
	return &GetProjectNotFound{}
}

/*
GetProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetProjectNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get project not found response has a 2xx status code
func (o *GetProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project not found response has a 3xx status code
func (o *GetProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project not found response has a 4xx status code
func (o *GetProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project not found response has a 5xx status code
func (o *GetProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get project not found response a status code equal to that given
func (o *GetProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) String() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
