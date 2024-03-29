// Code generated by go-swagger; DO NOT EDIT.

package resource_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetResourceActionUsingGET5Reader is a Reader for the GetResourceActionUsingGET5 structure.
type GetResourceActionUsingGET5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceActionUsingGET5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceActionUsingGET5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceActionUsingGET5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceActionUsingGET5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceActionUsingGET5OK creates a GetResourceActionUsingGET5OK with default headers values
func NewGetResourceActionUsingGET5OK() *GetResourceActionUsingGET5OK {
	return &GetResourceActionUsingGET5OK{}
}

/*
GetResourceActionUsingGET5OK describes a response with status code 200, with default header values.

OK
*/
type GetResourceActionUsingGET5OK struct {
	Payload *models.ResourceAction
}

// IsSuccess returns true when this get resource action using g e t5 o k response has a 2xx status code
func (o *GetResourceActionUsingGET5OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource action using g e t5 o k response has a 3xx status code
func (o *GetResourceActionUsingGET5OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource action using g e t5 o k response has a 4xx status code
func (o *GetResourceActionUsingGET5OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource action using g e t5 o k response has a 5xx status code
func (o *GetResourceActionUsingGET5OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource action using g e t5 o k response a status code equal to that given
func (o *GetResourceActionUsingGET5OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResourceActionUsingGET5OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/actions/{actionId}][%d] getResourceActionUsingGET5OK  %+v", 200, o.Payload)
}

func (o *GetResourceActionUsingGET5OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/actions/{actionId}][%d] getResourceActionUsingGET5OK  %+v", 200, o.Payload)
}

func (o *GetResourceActionUsingGET5OK) GetPayload() *models.ResourceAction {
	return o.Payload
}

func (o *GetResourceActionUsingGET5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceAction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceActionUsingGET5Unauthorized creates a GetResourceActionUsingGET5Unauthorized with default headers values
func NewGetResourceActionUsingGET5Unauthorized() *GetResourceActionUsingGET5Unauthorized {
	return &GetResourceActionUsingGET5Unauthorized{}
}

/*
GetResourceActionUsingGET5Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourceActionUsingGET5Unauthorized struct {
}

// IsSuccess returns true when this get resource action using g e t5 unauthorized response has a 2xx status code
func (o *GetResourceActionUsingGET5Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource action using g e t5 unauthorized response has a 3xx status code
func (o *GetResourceActionUsingGET5Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource action using g e t5 unauthorized response has a 4xx status code
func (o *GetResourceActionUsingGET5Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource action using g e t5 unauthorized response has a 5xx status code
func (o *GetResourceActionUsingGET5Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource action using g e t5 unauthorized response a status code equal to that given
func (o *GetResourceActionUsingGET5Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResourceActionUsingGET5Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/actions/{actionId}][%d] getResourceActionUsingGET5Unauthorized ", 401)
}

func (o *GetResourceActionUsingGET5Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/actions/{actionId}][%d] getResourceActionUsingGET5Unauthorized ", 401)
}

func (o *GetResourceActionUsingGET5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceActionUsingGET5NotFound creates a GetResourceActionUsingGET5NotFound with default headers values
func NewGetResourceActionUsingGET5NotFound() *GetResourceActionUsingGET5NotFound {
	return &GetResourceActionUsingGET5NotFound{}
}

/*
GetResourceActionUsingGET5NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetResourceActionUsingGET5NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get resource action using g e t5 not found response has a 2xx status code
func (o *GetResourceActionUsingGET5NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resource action using g e t5 not found response has a 3xx status code
func (o *GetResourceActionUsingGET5NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource action using g e t5 not found response has a 4xx status code
func (o *GetResourceActionUsingGET5NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resource action using g e t5 not found response has a 5xx status code
func (o *GetResourceActionUsingGET5NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource action using g e t5 not found response a status code equal to that given
func (o *GetResourceActionUsingGET5NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResourceActionUsingGET5NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/actions/{actionId}][%d] getResourceActionUsingGET5NotFound  %+v", 404, o.Payload)
}

func (o *GetResourceActionUsingGET5NotFound) String() string {
	return fmt.Sprintf("[GET /deployment/api/resources/{resourceId}/actions/{actionId}][%d] getResourceActionUsingGET5NotFound  %+v", 404, o.Payload)
}

func (o *GetResourceActionUsingGET5NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResourceActionUsingGET5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
