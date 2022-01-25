// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ResumeOrRerunGitEventUsingPOSTReader is a Reader for the ResumeOrRerunGitEventUsingPOST structure.
type ResumeOrRerunGitEventUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeOrRerunGitEventUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeOrRerunGitEventUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResumeOrRerunGitEventUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResumeOrRerunGitEventUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResumeOrRerunGitEventUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResumeOrRerunGitEventUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResumeOrRerunGitEventUsingPOSTOK creates a ResumeOrRerunGitEventUsingPOSTOK with default headers values
func NewResumeOrRerunGitEventUsingPOSTOK() *ResumeOrRerunGitEventUsingPOSTOK {
	return &ResumeOrRerunGitEventUsingPOSTOK{}
}

/* ResumeOrRerunGitEventUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with Re-run/Resume Git Event
*/
type ResumeOrRerunGitEventUsingPOSTOK struct {
	Payload models.GitEvent
}

func (o *ResumeOrRerunGitEventUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunGitEventUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *ResumeOrRerunGitEventUsingPOSTOK) GetPayload() models.GitEvent {
	return o.Payload
}

func (o *ResumeOrRerunGitEventUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewResumeOrRerunGitEventUsingPOSTUnauthorized creates a ResumeOrRerunGitEventUsingPOSTUnauthorized with default headers values
func NewResumeOrRerunGitEventUsingPOSTUnauthorized() *ResumeOrRerunGitEventUsingPOSTUnauthorized {
	return &ResumeOrRerunGitEventUsingPOSTUnauthorized{}
}

/* ResumeOrRerunGitEventUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ResumeOrRerunGitEventUsingPOSTUnauthorized struct {
}

func (o *ResumeOrRerunGitEventUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunGitEventUsingPOSTUnauthorized ", 401)
}

func (o *ResumeOrRerunGitEventUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeOrRerunGitEventUsingPOSTForbidden creates a ResumeOrRerunGitEventUsingPOSTForbidden with default headers values
func NewResumeOrRerunGitEventUsingPOSTForbidden() *ResumeOrRerunGitEventUsingPOSTForbidden {
	return &ResumeOrRerunGitEventUsingPOSTForbidden{}
}

/* ResumeOrRerunGitEventUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ResumeOrRerunGitEventUsingPOSTForbidden struct {
}

func (o *ResumeOrRerunGitEventUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunGitEventUsingPOSTForbidden ", 403)
}

func (o *ResumeOrRerunGitEventUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeOrRerunGitEventUsingPOSTNotFound creates a ResumeOrRerunGitEventUsingPOSTNotFound with default headers values
func NewResumeOrRerunGitEventUsingPOSTNotFound() *ResumeOrRerunGitEventUsingPOSTNotFound {
	return &ResumeOrRerunGitEventUsingPOSTNotFound{}
}

/* ResumeOrRerunGitEventUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ResumeOrRerunGitEventUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *ResumeOrRerunGitEventUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunGitEventUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *ResumeOrRerunGitEventUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResumeOrRerunGitEventUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeOrRerunGitEventUsingPOSTInternalServerError creates a ResumeOrRerunGitEventUsingPOSTInternalServerError with default headers values
func NewResumeOrRerunGitEventUsingPOSTInternalServerError() *ResumeOrRerunGitEventUsingPOSTInternalServerError {
	return &ResumeOrRerunGitEventUsingPOSTInternalServerError{}
}

/* ResumeOrRerunGitEventUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ResumeOrRerunGitEventUsingPOSTInternalServerError struct {
}

func (o *ResumeOrRerunGitEventUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunGitEventUsingPOSTInternalServerError ", 500)
}

func (o *ResumeOrRerunGitEventUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
