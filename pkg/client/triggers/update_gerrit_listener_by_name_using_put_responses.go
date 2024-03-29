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

// UpdateGerritListenerByNameUsingPUTReader is a Reader for the UpdateGerritListenerByNameUsingPUT structure.
type UpdateGerritListenerByNameUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGerritListenerByNameUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGerritListenerByNameUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateGerritListenerByNameUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGerritListenerByNameUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGerritListenerByNameUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateGerritListenerByNameUsingPUTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGerritListenerByNameUsingPUTOK creates a UpdateGerritListenerByNameUsingPUTOK with default headers values
func NewUpdateGerritListenerByNameUsingPUTOK() *UpdateGerritListenerByNameUsingPUTOK {
	return &UpdateGerritListenerByNameUsingPUTOK{}
}

/*
UpdateGerritListenerByNameUsingPUTOK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener Update
*/
type UpdateGerritListenerByNameUsingPUTOK struct {
	Payload models.GerritListener
}

// IsSuccess returns true when this update gerrit listener by name using p u t o k response has a 2xx status code
func (o *UpdateGerritListenerByNameUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update gerrit listener by name using p u t o k response has a 3xx status code
func (o *UpdateGerritListenerByNameUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gerrit listener by name using p u t o k response has a 4xx status code
func (o *UpdateGerritListenerByNameUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update gerrit listener by name using p u t o k response has a 5xx status code
func (o *UpdateGerritListenerByNameUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update gerrit listener by name using p u t o k response a status code equal to that given
func (o *UpdateGerritListenerByNameUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateGerritListenerByNameUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateGerritListenerByNameUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateGerritListenerByNameUsingPUTOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *UpdateGerritListenerByNameUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateGerritListenerByNameUsingPUTUnauthorized creates a UpdateGerritListenerByNameUsingPUTUnauthorized with default headers values
func NewUpdateGerritListenerByNameUsingPUTUnauthorized() *UpdateGerritListenerByNameUsingPUTUnauthorized {
	return &UpdateGerritListenerByNameUsingPUTUnauthorized{}
}

/*
UpdateGerritListenerByNameUsingPUTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateGerritListenerByNameUsingPUTUnauthorized struct {
}

// IsSuccess returns true when this update gerrit listener by name using p u t unauthorized response has a 2xx status code
func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gerrit listener by name using p u t unauthorized response has a 3xx status code
func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gerrit listener by name using p u t unauthorized response has a 4xx status code
func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update gerrit listener by name using p u t unauthorized response has a 5xx status code
func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update gerrit listener by name using p u t unauthorized response a status code equal to that given
func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTUnauthorized ", 401)
}

func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) String() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTUnauthorized ", 401)
}

func (o *UpdateGerritListenerByNameUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGerritListenerByNameUsingPUTForbidden creates a UpdateGerritListenerByNameUsingPUTForbidden with default headers values
func NewUpdateGerritListenerByNameUsingPUTForbidden() *UpdateGerritListenerByNameUsingPUTForbidden {
	return &UpdateGerritListenerByNameUsingPUTForbidden{}
}

/*
UpdateGerritListenerByNameUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateGerritListenerByNameUsingPUTForbidden struct {
}

// IsSuccess returns true when this update gerrit listener by name using p u t forbidden response has a 2xx status code
func (o *UpdateGerritListenerByNameUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gerrit listener by name using p u t forbidden response has a 3xx status code
func (o *UpdateGerritListenerByNameUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gerrit listener by name using p u t forbidden response has a 4xx status code
func (o *UpdateGerritListenerByNameUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update gerrit listener by name using p u t forbidden response has a 5xx status code
func (o *UpdateGerritListenerByNameUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update gerrit listener by name using p u t forbidden response a status code equal to that given
func (o *UpdateGerritListenerByNameUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateGerritListenerByNameUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTForbidden ", 403)
}

func (o *UpdateGerritListenerByNameUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTForbidden ", 403)
}

func (o *UpdateGerritListenerByNameUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGerritListenerByNameUsingPUTNotFound creates a UpdateGerritListenerByNameUsingPUTNotFound with default headers values
func NewUpdateGerritListenerByNameUsingPUTNotFound() *UpdateGerritListenerByNameUsingPUTNotFound {
	return &UpdateGerritListenerByNameUsingPUTNotFound{}
}

/*
UpdateGerritListenerByNameUsingPUTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateGerritListenerByNameUsingPUTNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update gerrit listener by name using p u t not found response has a 2xx status code
func (o *UpdateGerritListenerByNameUsingPUTNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gerrit listener by name using p u t not found response has a 3xx status code
func (o *UpdateGerritListenerByNameUsingPUTNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gerrit listener by name using p u t not found response has a 4xx status code
func (o *UpdateGerritListenerByNameUsingPUTNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update gerrit listener by name using p u t not found response has a 5xx status code
func (o *UpdateGerritListenerByNameUsingPUTNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update gerrit listener by name using p u t not found response a status code equal to that given
func (o *UpdateGerritListenerByNameUsingPUTNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateGerritListenerByNameUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGerritListenerByNameUsingPUTNotFound) String() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGerritListenerByNameUsingPUTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGerritListenerByNameUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGerritListenerByNameUsingPUTInternalServerError creates a UpdateGerritListenerByNameUsingPUTInternalServerError with default headers values
func NewUpdateGerritListenerByNameUsingPUTInternalServerError() *UpdateGerritListenerByNameUsingPUTInternalServerError {
	return &UpdateGerritListenerByNameUsingPUTInternalServerError{}
}

/*
UpdateGerritListenerByNameUsingPUTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateGerritListenerByNameUsingPUTInternalServerError struct {
}

// IsSuccess returns true when this update gerrit listener by name using p u t internal server error response has a 2xx status code
func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update gerrit listener by name using p u t internal server error response has a 3xx status code
func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update gerrit listener by name using p u t internal server error response has a 4xx status code
func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update gerrit listener by name using p u t internal server error response has a 5xx status code
func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update gerrit listener by name using p u t internal server error response a status code equal to that given
func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTInternalServerError ", 500)
}

func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) String() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateGerritListenerByNameUsingPUTInternalServerError ", 500)
}

func (o *UpdateGerritListenerByNameUsingPUTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
