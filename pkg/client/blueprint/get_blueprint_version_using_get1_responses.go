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

// GetBlueprintVersionUsingGET1Reader is a Reader for the GetBlueprintVersionUsingGET1 structure.
type GetBlueprintVersionUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintVersionUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlueprintVersionUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBlueprintVersionUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBlueprintVersionUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBlueprintVersionUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBlueprintVersionUsingGET1OK creates a GetBlueprintVersionUsingGET1OK with default headers values
func NewGetBlueprintVersionUsingGET1OK() *GetBlueprintVersionUsingGET1OK {
	return &GetBlueprintVersionUsingGET1OK{}
}

/*
GetBlueprintVersionUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetBlueprintVersionUsingGET1OK struct {
	Payload *models.BlueprintVersion
}

// IsSuccess returns true when this get blueprint version using g e t1 o k response has a 2xx status code
func (o *GetBlueprintVersionUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get blueprint version using g e t1 o k response has a 3xx status code
func (o *GetBlueprintVersionUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint version using g e t1 o k response has a 4xx status code
func (o *GetBlueprintVersionUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get blueprint version using g e t1 o k response has a 5xx status code
func (o *GetBlueprintVersionUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint version using g e t1 o k response a status code equal to that given
func (o *GetBlueprintVersionUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBlueprintVersionUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetBlueprintVersionUsingGET1OK) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetBlueprintVersionUsingGET1OK) GetPayload() *models.BlueprintVersion {
	return o.Payload
}

func (o *GetBlueprintVersionUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlueprintVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprintVersionUsingGET1Unauthorized creates a GetBlueprintVersionUsingGET1Unauthorized with default headers values
func NewGetBlueprintVersionUsingGET1Unauthorized() *GetBlueprintVersionUsingGET1Unauthorized {
	return &GetBlueprintVersionUsingGET1Unauthorized{}
}

/*
GetBlueprintVersionUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetBlueprintVersionUsingGET1Unauthorized struct {
}

// IsSuccess returns true when this get blueprint version using g e t1 unauthorized response has a 2xx status code
func (o *GetBlueprintVersionUsingGET1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blueprint version using g e t1 unauthorized response has a 3xx status code
func (o *GetBlueprintVersionUsingGET1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint version using g e t1 unauthorized response has a 4xx status code
func (o *GetBlueprintVersionUsingGET1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blueprint version using g e t1 unauthorized response has a 5xx status code
func (o *GetBlueprintVersionUsingGET1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint version using g e t1 unauthorized response a status code equal to that given
func (o *GetBlueprintVersionUsingGET1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetBlueprintVersionUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1Unauthorized ", 401)
}

func (o *GetBlueprintVersionUsingGET1Unauthorized) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1Unauthorized ", 401)
}

func (o *GetBlueprintVersionUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintVersionUsingGET1Forbidden creates a GetBlueprintVersionUsingGET1Forbidden with default headers values
func NewGetBlueprintVersionUsingGET1Forbidden() *GetBlueprintVersionUsingGET1Forbidden {
	return &GetBlueprintVersionUsingGET1Forbidden{}
}

/*
GetBlueprintVersionUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetBlueprintVersionUsingGET1Forbidden struct {
}

// IsSuccess returns true when this get blueprint version using g e t1 forbidden response has a 2xx status code
func (o *GetBlueprintVersionUsingGET1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blueprint version using g e t1 forbidden response has a 3xx status code
func (o *GetBlueprintVersionUsingGET1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint version using g e t1 forbidden response has a 4xx status code
func (o *GetBlueprintVersionUsingGET1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blueprint version using g e t1 forbidden response has a 5xx status code
func (o *GetBlueprintVersionUsingGET1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint version using g e t1 forbidden response a status code equal to that given
func (o *GetBlueprintVersionUsingGET1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetBlueprintVersionUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1Forbidden ", 403)
}

func (o *GetBlueprintVersionUsingGET1Forbidden) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1Forbidden ", 403)
}

func (o *GetBlueprintVersionUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintVersionUsingGET1NotFound creates a GetBlueprintVersionUsingGET1NotFound with default headers values
func NewGetBlueprintVersionUsingGET1NotFound() *GetBlueprintVersionUsingGET1NotFound {
	return &GetBlueprintVersionUsingGET1NotFound{}
}

/*
GetBlueprintVersionUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetBlueprintVersionUsingGET1NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get blueprint version using g e t1 not found response has a 2xx status code
func (o *GetBlueprintVersionUsingGET1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blueprint version using g e t1 not found response has a 3xx status code
func (o *GetBlueprintVersionUsingGET1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint version using g e t1 not found response has a 4xx status code
func (o *GetBlueprintVersionUsingGET1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blueprint version using g e t1 not found response has a 5xx status code
func (o *GetBlueprintVersionUsingGET1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint version using g e t1 not found response a status code equal to that given
func (o *GetBlueprintVersionUsingGET1NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetBlueprintVersionUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1NotFound  %+v", 404, o.Payload)
}

func (o *GetBlueprintVersionUsingGET1NotFound) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}][%d] getBlueprintVersionUsingGET1NotFound  %+v", 404, o.Payload)
}

func (o *GetBlueprintVersionUsingGET1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBlueprintVersionUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
