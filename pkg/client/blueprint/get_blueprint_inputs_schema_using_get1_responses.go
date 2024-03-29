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

// GetBlueprintInputsSchemaUsingGET1Reader is a Reader for the GetBlueprintInputsSchemaUsingGET1 structure.
type GetBlueprintInputsSchemaUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintInputsSchemaUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlueprintInputsSchemaUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBlueprintInputsSchemaUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBlueprintInputsSchemaUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBlueprintInputsSchemaUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBlueprintInputsSchemaUsingGET1OK creates a GetBlueprintInputsSchemaUsingGET1OK with default headers values
func NewGetBlueprintInputsSchemaUsingGET1OK() *GetBlueprintInputsSchemaUsingGET1OK {
	return &GetBlueprintInputsSchemaUsingGET1OK{}
}

/*
GetBlueprintInputsSchemaUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetBlueprintInputsSchemaUsingGET1OK struct {
	Payload *models.PropertyDefinition
}

// IsSuccess returns true when this get blueprint inputs schema using g e t1 o k response has a 2xx status code
func (o *GetBlueprintInputsSchemaUsingGET1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get blueprint inputs schema using g e t1 o k response has a 3xx status code
func (o *GetBlueprintInputsSchemaUsingGET1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint inputs schema using g e t1 o k response has a 4xx status code
func (o *GetBlueprintInputsSchemaUsingGET1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get blueprint inputs schema using g e t1 o k response has a 5xx status code
func (o *GetBlueprintInputsSchemaUsingGET1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint inputs schema using g e t1 o k response a status code equal to that given
func (o *GetBlueprintInputsSchemaUsingGET1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetBlueprintInputsSchemaUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetBlueprintInputsSchemaUsingGET1OK) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetBlueprintInputsSchemaUsingGET1OK) GetPayload() *models.PropertyDefinition {
	return o.Payload
}

func (o *GetBlueprintInputsSchemaUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprintInputsSchemaUsingGET1Unauthorized creates a GetBlueprintInputsSchemaUsingGET1Unauthorized with default headers values
func NewGetBlueprintInputsSchemaUsingGET1Unauthorized() *GetBlueprintInputsSchemaUsingGET1Unauthorized {
	return &GetBlueprintInputsSchemaUsingGET1Unauthorized{}
}

/*
GetBlueprintInputsSchemaUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetBlueprintInputsSchemaUsingGET1Unauthorized struct {
}

// IsSuccess returns true when this get blueprint inputs schema using g e t1 unauthorized response has a 2xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blueprint inputs schema using g e t1 unauthorized response has a 3xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint inputs schema using g e t1 unauthorized response has a 4xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blueprint inputs schema using g e t1 unauthorized response has a 5xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint inputs schema using g e t1 unauthorized response a status code equal to that given
func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1Unauthorized ", 401)
}

func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1Unauthorized ", 401)
}

func (o *GetBlueprintInputsSchemaUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintInputsSchemaUsingGET1Forbidden creates a GetBlueprintInputsSchemaUsingGET1Forbidden with default headers values
func NewGetBlueprintInputsSchemaUsingGET1Forbidden() *GetBlueprintInputsSchemaUsingGET1Forbidden {
	return &GetBlueprintInputsSchemaUsingGET1Forbidden{}
}

/*
GetBlueprintInputsSchemaUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetBlueprintInputsSchemaUsingGET1Forbidden struct {
}

// IsSuccess returns true when this get blueprint inputs schema using g e t1 forbidden response has a 2xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blueprint inputs schema using g e t1 forbidden response has a 3xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint inputs schema using g e t1 forbidden response has a 4xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blueprint inputs schema using g e t1 forbidden response has a 5xx status code
func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint inputs schema using g e t1 forbidden response a status code equal to that given
func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1Forbidden ", 403)
}

func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1Forbidden ", 403)
}

func (o *GetBlueprintInputsSchemaUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintInputsSchemaUsingGET1NotFound creates a GetBlueprintInputsSchemaUsingGET1NotFound with default headers values
func NewGetBlueprintInputsSchemaUsingGET1NotFound() *GetBlueprintInputsSchemaUsingGET1NotFound {
	return &GetBlueprintInputsSchemaUsingGET1NotFound{}
}

/*
GetBlueprintInputsSchemaUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetBlueprintInputsSchemaUsingGET1NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get blueprint inputs schema using g e t1 not found response has a 2xx status code
func (o *GetBlueprintInputsSchemaUsingGET1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blueprint inputs schema using g e t1 not found response has a 3xx status code
func (o *GetBlueprintInputsSchemaUsingGET1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blueprint inputs schema using g e t1 not found response has a 4xx status code
func (o *GetBlueprintInputsSchemaUsingGET1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blueprint inputs schema using g e t1 not found response has a 5xx status code
func (o *GetBlueprintInputsSchemaUsingGET1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get blueprint inputs schema using g e t1 not found response a status code equal to that given
func (o *GetBlueprintInputsSchemaUsingGET1NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetBlueprintInputsSchemaUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1NotFound  %+v", 404, o.Payload)
}

func (o *GetBlueprintInputsSchemaUsingGET1NotFound) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGET1NotFound  %+v", 404, o.Payload)
}

func (o *GetBlueprintInputsSchemaUsingGET1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBlueprintInputsSchemaUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
