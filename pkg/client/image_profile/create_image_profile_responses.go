// Code generated by go-swagger; DO NOT EDIT.

package image_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateImageProfileReader is a Reader for the CreateImageProfile structure.
type CreateImageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateImageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateImageProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateImageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateImageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateImageProfileCreated creates a CreateImageProfileCreated with default headers values
func NewCreateImageProfileCreated() *CreateImageProfileCreated {
	return &CreateImageProfileCreated{}
}

/*
CreateImageProfileCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateImageProfileCreated struct {
	Payload *models.ImageProfile
}

// IsSuccess returns true when this create image profile created response has a 2xx status code
func (o *CreateImageProfileCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create image profile created response has a 3xx status code
func (o *CreateImageProfileCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create image profile created response has a 4xx status code
func (o *CreateImageProfileCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create image profile created response has a 5xx status code
func (o *CreateImageProfileCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create image profile created response a status code equal to that given
func (o *CreateImageProfileCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateImageProfileCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/image-profiles][%d] createImageProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateImageProfileCreated) String() string {
	return fmt.Sprintf("[POST /iaas/api/image-profiles][%d] createImageProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateImageProfileCreated) GetPayload() *models.ImageProfile {
	return o.Payload
}

func (o *CreateImageProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImageProfileBadRequest creates a CreateImageProfileBadRequest with default headers values
func NewCreateImageProfileBadRequest() *CreateImageProfileBadRequest {
	return &CreateImageProfileBadRequest{}
}

/*
CreateImageProfileBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateImageProfileBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create image profile bad request response has a 2xx status code
func (o *CreateImageProfileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create image profile bad request response has a 3xx status code
func (o *CreateImageProfileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create image profile bad request response has a 4xx status code
func (o *CreateImageProfileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create image profile bad request response has a 5xx status code
func (o *CreateImageProfileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create image profile bad request response a status code equal to that given
func (o *CreateImageProfileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateImageProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/image-profiles][%d] createImageProfileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateImageProfileBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/image-profiles][%d] createImageProfileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateImageProfileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateImageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateImageProfileForbidden creates a CreateImageProfileForbidden with default headers values
func NewCreateImageProfileForbidden() *CreateImageProfileForbidden {
	return &CreateImageProfileForbidden{}
}

/*
CreateImageProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateImageProfileForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this create image profile forbidden response has a 2xx status code
func (o *CreateImageProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create image profile forbidden response has a 3xx status code
func (o *CreateImageProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create image profile forbidden response has a 4xx status code
func (o *CreateImageProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create image profile forbidden response has a 5xx status code
func (o *CreateImageProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create image profile forbidden response a status code equal to that given
func (o *CreateImageProfileForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateImageProfileForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/image-profiles][%d] createImageProfileForbidden  %+v", 403, o.Payload)
}

func (o *CreateImageProfileForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/image-profiles][%d] createImageProfileForbidden  %+v", 403, o.Payload)
}

func (o *CreateImageProfileForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateImageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
