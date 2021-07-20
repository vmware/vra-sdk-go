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

// ResumeOrRerunEventUsingPOST2Reader is a Reader for the ResumeOrRerunEventUsingPOST2 structure.
type ResumeOrRerunEventUsingPOST2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeOrRerunEventUsingPOST2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeOrRerunEventUsingPOST2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResumeOrRerunEventUsingPOST2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResumeOrRerunEventUsingPOST2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResumeOrRerunEventUsingPOST2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResumeOrRerunEventUsingPOST2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResumeOrRerunEventUsingPOST2OK creates a ResumeOrRerunEventUsingPOST2OK with default headers values
func NewResumeOrRerunEventUsingPOST2OK() *ResumeOrRerunEventUsingPOST2OK {
	return &ResumeOrRerunEventUsingPOST2OK{}
}

/* ResumeOrRerunEventUsingPOST2OK describes a response with status code 200, with default header values.

'Success' with Re-run/Resume Git Event
*/
type ResumeOrRerunEventUsingPOST2OK struct {
	Payload models.GitEvent
}

func (o *ResumeOrRerunEventUsingPOST2OK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunEventUsingPOST2OK  %+v", 200, o.Payload)
}
func (o *ResumeOrRerunEventUsingPOST2OK) GetPayload() models.GitEvent {
	return o.Payload
}

func (o *ResumeOrRerunEventUsingPOST2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewResumeOrRerunEventUsingPOST2Unauthorized creates a ResumeOrRerunEventUsingPOST2Unauthorized with default headers values
func NewResumeOrRerunEventUsingPOST2Unauthorized() *ResumeOrRerunEventUsingPOST2Unauthorized {
	return &ResumeOrRerunEventUsingPOST2Unauthorized{}
}

/* ResumeOrRerunEventUsingPOST2Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ResumeOrRerunEventUsingPOST2Unauthorized struct {
}

func (o *ResumeOrRerunEventUsingPOST2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunEventUsingPOST2Unauthorized ", 401)
}

func (o *ResumeOrRerunEventUsingPOST2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeOrRerunEventUsingPOST2Forbidden creates a ResumeOrRerunEventUsingPOST2Forbidden with default headers values
func NewResumeOrRerunEventUsingPOST2Forbidden() *ResumeOrRerunEventUsingPOST2Forbidden {
	return &ResumeOrRerunEventUsingPOST2Forbidden{}
}

/* ResumeOrRerunEventUsingPOST2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ResumeOrRerunEventUsingPOST2Forbidden struct {
}

func (o *ResumeOrRerunEventUsingPOST2Forbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunEventUsingPOST2Forbidden ", 403)
}

func (o *ResumeOrRerunEventUsingPOST2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeOrRerunEventUsingPOST2NotFound creates a ResumeOrRerunEventUsingPOST2NotFound with default headers values
func NewResumeOrRerunEventUsingPOST2NotFound() *ResumeOrRerunEventUsingPOST2NotFound {
	return &ResumeOrRerunEventUsingPOST2NotFound{}
}

/* ResumeOrRerunEventUsingPOST2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type ResumeOrRerunEventUsingPOST2NotFound struct {
	Payload *models.Error
}

func (o *ResumeOrRerunEventUsingPOST2NotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunEventUsingPOST2NotFound  %+v", 404, o.Payload)
}
func (o *ResumeOrRerunEventUsingPOST2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResumeOrRerunEventUsingPOST2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeOrRerunEventUsingPOST2InternalServerError creates a ResumeOrRerunEventUsingPOST2InternalServerError with default headers values
func NewResumeOrRerunEventUsingPOST2InternalServerError() *ResumeOrRerunEventUsingPOST2InternalServerError {
	return &ResumeOrRerunEventUsingPOST2InternalServerError{}
}

/* ResumeOrRerunEventUsingPOST2InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ResumeOrRerunEventUsingPOST2InternalServerError struct {
}

func (o *ResumeOrRerunEventUsingPOST2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-events/{id}][%d] resumeOrRerunEventUsingPOST2InternalServerError ", 500)
}

func (o *ResumeOrRerunEventUsingPOST2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
