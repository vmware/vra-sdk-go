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

// ModifyProjectCostsUsingPATCHReader is a Reader for the ModifyProjectCostsUsingPATCH structure.
type ModifyProjectCostsUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyProjectCostsUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyProjectCostsUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewModifyProjectCostsUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewModifyProjectCostsUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifyProjectCostsUsingPATCHOK creates a ModifyProjectCostsUsingPATCHOK with default headers values
func NewModifyProjectCostsUsingPATCHOK() *ModifyProjectCostsUsingPATCHOK {
	return &ModifyProjectCostsUsingPATCHOK{}
}

/* ModifyProjectCostsUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with the Project
*/
type ModifyProjectCostsUsingPATCHOK struct {
	Payload *models.Project
}

func (o *ModifyProjectCostsUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}/cost][%d] modifyProjectCostsUsingPATCHOK  %+v", 200, o.Payload)
}
func (o *ModifyProjectCostsUsingPATCHOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *ModifyProjectCostsUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyProjectCostsUsingPATCHForbidden creates a ModifyProjectCostsUsingPATCHForbidden with default headers values
func NewModifyProjectCostsUsingPATCHForbidden() *ModifyProjectCostsUsingPATCHForbidden {
	return &ModifyProjectCostsUsingPATCHForbidden{}
}

/* ModifyProjectCostsUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type ModifyProjectCostsUsingPATCHForbidden struct {
}

func (o *ModifyProjectCostsUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}/cost][%d] modifyProjectCostsUsingPATCHForbidden ", 403)
}

func (o *ModifyProjectCostsUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyProjectCostsUsingPATCHNotFound creates a ModifyProjectCostsUsingPATCHNotFound with default headers values
func NewModifyProjectCostsUsingPATCHNotFound() *ModifyProjectCostsUsingPATCHNotFound {
	return &ModifyProjectCostsUsingPATCHNotFound{}
}

/* ModifyProjectCostsUsingPATCHNotFound describes a response with status code 404, with default header values.

'Not found' if no project with the provided id
*/
type ModifyProjectCostsUsingPATCHNotFound struct {
	Payload *models.Error
}

func (o *ModifyProjectCostsUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}/cost][%d] modifyProjectCostsUsingPATCHNotFound  %+v", 404, o.Payload)
}
func (o *ModifyProjectCostsUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModifyProjectCostsUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
