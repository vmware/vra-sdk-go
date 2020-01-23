// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateUsingPOST2Reader is a Reader for the CreateUsingPOST2 structure.
type CreateUsingPOST2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingPOST2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUsingPOST2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUsingPOST2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUsingPOST2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUsingPOST2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUsingPOST2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUsingPOST2OK creates a CreateUsingPOST2OK with default headers values
func NewCreateUsingPOST2OK() *CreateUsingPOST2OK {
	return &CreateUsingPOST2OK{}
}

/*CreateUsingPOST2OK handles this case with default header values.

'Success' with Gerrit Trigger Creation
*/
type CreateUsingPOST2OK struct {
	Payload models.GerritTrigger
}

func (o *CreateUsingPOST2OK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers][%d] createUsingPOST2OK  %+v", 200, o.Payload)
}

func (o *CreateUsingPOST2OK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *CreateUsingPOST2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateUsingPOST2Unauthorized creates a CreateUsingPOST2Unauthorized with default headers values
func NewCreateUsingPOST2Unauthorized() *CreateUsingPOST2Unauthorized {
	return &CreateUsingPOST2Unauthorized{}
}

/*CreateUsingPOST2Unauthorized handles this case with default header values.

Unauthorized Request
*/
type CreateUsingPOST2Unauthorized struct {
}

func (o *CreateUsingPOST2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers][%d] createUsingPOST2Unauthorized ", 401)
}

func (o *CreateUsingPOST2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOST2Forbidden creates a CreateUsingPOST2Forbidden with default headers values
func NewCreateUsingPOST2Forbidden() *CreateUsingPOST2Forbidden {
	return &CreateUsingPOST2Forbidden{}
}

/*CreateUsingPOST2Forbidden handles this case with default header values.

Forbidden
*/
type CreateUsingPOST2Forbidden struct {
}

func (o *CreateUsingPOST2Forbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers][%d] createUsingPOST2Forbidden ", 403)
}

func (o *CreateUsingPOST2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOST2NotFound creates a CreateUsingPOST2NotFound with default headers values
func NewCreateUsingPOST2NotFound() *CreateUsingPOST2NotFound {
	return &CreateUsingPOST2NotFound{}
}

/*CreateUsingPOST2NotFound handles this case with default header values.

Not Found
*/
type CreateUsingPOST2NotFound struct {
}

func (o *CreateUsingPOST2NotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers][%d] createUsingPOST2NotFound ", 404)
}

func (o *CreateUsingPOST2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOST2InternalServerError creates a CreateUsingPOST2InternalServerError with default headers values
func NewCreateUsingPOST2InternalServerError() *CreateUsingPOST2InternalServerError {
	return &CreateUsingPOST2InternalServerError{}
}

/*CreateUsingPOST2InternalServerError handles this case with default header values.

Server Error
*/
type CreateUsingPOST2InternalServerError struct {
}

func (o *CreateUsingPOST2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers][%d] createUsingPOST2InternalServerError ", 500)
}

func (o *CreateUsingPOST2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
