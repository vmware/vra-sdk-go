// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ListBlueprintsUsingGET1Reader is a Reader for the ListBlueprintsUsingGET1 structure.
type ListBlueprintsUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBlueprintsUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBlueprintsUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListBlueprintsUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListBlueprintsUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListBlueprintsUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListBlueprintsUsingGET1OK creates a ListBlueprintsUsingGET1OK with default headers values
func NewListBlueprintsUsingGET1OK() *ListBlueprintsUsingGET1OK {
	return &ListBlueprintsUsingGET1OK{}
}

/*
ListBlueprintsUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type ListBlueprintsUsingGET1OK struct {
	Payload *models.PageOfBlueprint
}

// IsSuccess returns true when this list blueprints using g e t1 o k response has a 2xx status code
func (o *ListBlueprintsUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list blueprints using g e t1 o k response has a 3xx status code
func (o *ListBlueprintsUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list blueprints using g e t1 o k response has a 4xx status code
func (o *ListBlueprintsUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list blueprints using g e t1 o k response has a 5xx status code
func (o *ListBlueprintsUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list blueprints using g e t1 o k response a status code equal to that given
func (o *ListBlueprintsUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListBlueprintsUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1OK  %+v", 200, o.Payload)
}

func (o *ListBlueprintsUsingGET1OK) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1OK  %+v", 200, o.Payload)
}

func (o *ListBlueprintsUsingGET1OK) GetPayload() *models.PageOfBlueprint {
	return o.Payload
}

func (o *ListBlueprintsUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfBlueprint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBlueprintsUsingGET1BadRequest creates a ListBlueprintsUsingGET1BadRequest with default headers values
func NewListBlueprintsUsingGET1BadRequest() *ListBlueprintsUsingGET1BadRequest {
	return &ListBlueprintsUsingGET1BadRequest{}
}

/*
ListBlueprintsUsingGET1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListBlueprintsUsingGET1BadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this list blueprints using g e t1 bad request response has a 2xx status code
func (o *ListBlueprintsUsingGET1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list blueprints using g e t1 bad request response has a 3xx status code
func (o *ListBlueprintsUsingGET1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list blueprints using g e t1 bad request response has a 4xx status code
func (o *ListBlueprintsUsingGET1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list blueprints using g e t1 bad request response has a 5xx status code
func (o *ListBlueprintsUsingGET1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list blueprints using g e t1 bad request response a status code equal to that given
func (o *ListBlueprintsUsingGET1BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListBlueprintsUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1BadRequest  %+v", 400, o.Payload)
}

func (o *ListBlueprintsUsingGET1BadRequest) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1BadRequest  %+v", 400, o.Payload)
}

func (o *ListBlueprintsUsingGET1BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListBlueprintsUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBlueprintsUsingGET1Unauthorized creates a ListBlueprintsUsingGET1Unauthorized with default headers values
func NewListBlueprintsUsingGET1Unauthorized() *ListBlueprintsUsingGET1Unauthorized {
	return &ListBlueprintsUsingGET1Unauthorized{}
}

/*
ListBlueprintsUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListBlueprintsUsingGET1Unauthorized struct {
}

// IsSuccess returns true when this list blueprints using g e t1 unauthorized response has a 2xx status code
func (o *ListBlueprintsUsingGET1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list blueprints using g e t1 unauthorized response has a 3xx status code
func (o *ListBlueprintsUsingGET1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list blueprints using g e t1 unauthorized response has a 4xx status code
func (o *ListBlueprintsUsingGET1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list blueprints using g e t1 unauthorized response has a 5xx status code
func (o *ListBlueprintsUsingGET1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list blueprints using g e t1 unauthorized response a status code equal to that given
func (o *ListBlueprintsUsingGET1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListBlueprintsUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1Unauthorized ", 401)
}

func (o *ListBlueprintsUsingGET1Unauthorized) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1Unauthorized ", 401)
}

func (o *ListBlueprintsUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListBlueprintsUsingGET1Forbidden creates a ListBlueprintsUsingGET1Forbidden with default headers values
func NewListBlueprintsUsingGET1Forbidden() *ListBlueprintsUsingGET1Forbidden {
	return &ListBlueprintsUsingGET1Forbidden{}
}

/*
ListBlueprintsUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListBlueprintsUsingGET1Forbidden struct {
}

// IsSuccess returns true when this list blueprints using g e t1 forbidden response has a 2xx status code
func (o *ListBlueprintsUsingGET1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list blueprints using g e t1 forbidden response has a 3xx status code
func (o *ListBlueprintsUsingGET1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list blueprints using g e t1 forbidden response has a 4xx status code
func (o *ListBlueprintsUsingGET1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list blueprints using g e t1 forbidden response has a 5xx status code
func (o *ListBlueprintsUsingGET1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list blueprints using g e t1 forbidden response a status code equal to that given
func (o *ListBlueprintsUsingGET1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListBlueprintsUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1Forbidden ", 403)
}

func (o *ListBlueprintsUsingGET1Forbidden) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints][%d] listBlueprintsUsingGET1Forbidden ", 403)
}

func (o *ListBlueprintsUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
