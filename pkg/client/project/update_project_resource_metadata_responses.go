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

// UpdateProjectResourceMetadataReader is a Reader for the UpdateProjectResourceMetadata structure.
type UpdateProjectResourceMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectResourceMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectResourceMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectResourceMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectResourceMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectResourceMetadataOK creates a UpdateProjectResourceMetadataOK with default headers values
func NewUpdateProjectResourceMetadataOK() *UpdateProjectResourceMetadataOK {
	return &UpdateProjectResourceMetadataOK{}
}

/*
UpdateProjectResourceMetadataOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateProjectResourceMetadataOK struct {
	Payload *models.IaaSProject
}

// IsSuccess returns true when this update project resource metadata o k response has a 2xx status code
func (o *UpdateProjectResourceMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project resource metadata o k response has a 3xx status code
func (o *UpdateProjectResourceMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project resource metadata o k response has a 4xx status code
func (o *UpdateProjectResourceMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project resource metadata o k response has a 5xx status code
func (o *UpdateProjectResourceMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project resource metadata o k response a status code equal to that given
func (o *UpdateProjectResourceMetadataOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProjectResourceMetadataOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectResourceMetadataOK) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectResourceMetadataOK) GetPayload() *models.IaaSProject {
	return o.Payload
}

func (o *UpdateProjectResourceMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaaSProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectResourceMetadataBadRequest creates a UpdateProjectResourceMetadataBadRequest with default headers values
func NewUpdateProjectResourceMetadataBadRequest() *UpdateProjectResourceMetadataBadRequest {
	return &UpdateProjectResourceMetadataBadRequest{}
}

/*
UpdateProjectResourceMetadataBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type UpdateProjectResourceMetadataBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update project resource metadata bad request response has a 2xx status code
func (o *UpdateProjectResourceMetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project resource metadata bad request response has a 3xx status code
func (o *UpdateProjectResourceMetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project resource metadata bad request response has a 4xx status code
func (o *UpdateProjectResourceMetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project resource metadata bad request response has a 5xx status code
func (o *UpdateProjectResourceMetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update project resource metadata bad request response a status code equal to that given
func (o *UpdateProjectResourceMetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateProjectResourceMetadataBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectResourceMetadataBadRequest) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectResourceMetadataBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateProjectResourceMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectResourceMetadataForbidden creates a UpdateProjectResourceMetadataForbidden with default headers values
func NewUpdateProjectResourceMetadataForbidden() *UpdateProjectResourceMetadataForbidden {
	return &UpdateProjectResourceMetadataForbidden{}
}

/*
UpdateProjectResourceMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateProjectResourceMetadataForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this update project resource metadata forbidden response has a 2xx status code
func (o *UpdateProjectResourceMetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project resource metadata forbidden response has a 3xx status code
func (o *UpdateProjectResourceMetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project resource metadata forbidden response has a 4xx status code
func (o *UpdateProjectResourceMetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project resource metadata forbidden response has a 5xx status code
func (o *UpdateProjectResourceMetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update project resource metadata forbidden response a status code equal to that given
func (o *UpdateProjectResourceMetadataForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateProjectResourceMetadataForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectResourceMetadataForbidden) String() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectResourceMetadataForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateProjectResourceMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
