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

// CreateUsingPOSTReader is a Reader for the CreateUsingPOST structure.
type CreateUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUsingPOSTConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUsingPOSTCreated creates a CreateUsingPOSTCreated with default headers values
func NewCreateUsingPOSTCreated() *CreateUsingPOSTCreated {
	return &CreateUsingPOSTCreated{}
}

/* CreateUsingPOSTCreated describes a response with status code 201, with default header values.

'Created' with the newly created project
*/
type CreateUsingPOSTCreated struct {
	Payload models.Project
}

func (o *CreateUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /project-service/api/projects][%d] createUsingPOSTCreated  %+v", 201, o.Payload)
}
func (o *CreateUsingPOSTCreated) GetPayload() models.Project {
	return o.Payload
}

func (o *CreateUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalProject(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateUsingPOSTBadRequest creates a CreateUsingPOSTBadRequest with default headers values
func NewCreateUsingPOSTBadRequest() *CreateUsingPOSTBadRequest {
	return &CreateUsingPOSTBadRequest{}
}

/* CreateUsingPOSTBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateUsingPOSTBadRequest struct {
	Payload *models.Error
}

func (o *CreateUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /project-service/api/projects][%d] createUsingPOSTBadRequest  %+v", 400, o.Payload)
}
func (o *CreateUsingPOSTBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUsingPOSTForbidden creates a CreateUsingPOSTForbidden with default headers values
func NewCreateUsingPOSTForbidden() *CreateUsingPOSTForbidden {
	return &CreateUsingPOSTForbidden{}
}

/* CreateUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type CreateUsingPOSTForbidden struct {
}

func (o *CreateUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /project-service/api/projects][%d] createUsingPOSTForbidden ", 403)
}

func (o *CreateUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOSTConflict creates a CreateUsingPOSTConflict with default headers values
func NewCreateUsingPOSTConflict() *CreateUsingPOSTConflict {
	return &CreateUsingPOSTConflict{}
}

/* CreateUsingPOSTConflict describes a response with status code 409, with default header values.

Conflict, project with this name already exists
*/
type CreateUsingPOSTConflict struct {
}

func (o *CreateUsingPOSTConflict) Error() string {
	return fmt.Sprintf("[POST /project-service/api/projects][%d] createUsingPOSTConflict ", 409)
}

func (o *CreateUsingPOSTConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
