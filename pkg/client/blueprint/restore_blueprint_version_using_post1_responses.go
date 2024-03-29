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

// RestoreBlueprintVersionUsingPOST1Reader is a Reader for the RestoreBlueprintVersionUsingPOST1 structure.
type RestoreBlueprintVersionUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreBlueprintVersionUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreBlueprintVersionUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestoreBlueprintVersionUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestoreBlueprintVersionUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestoreBlueprintVersionUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestoreBlueprintVersionUsingPOST1OK creates a RestoreBlueprintVersionUsingPOST1OK with default headers values
func NewRestoreBlueprintVersionUsingPOST1OK() *RestoreBlueprintVersionUsingPOST1OK {
	return &RestoreBlueprintVersionUsingPOST1OK{}
}

/*
RestoreBlueprintVersionUsingPOST1OK describes a response with status code 200, with default header values.

OK
*/
type RestoreBlueprintVersionUsingPOST1OK struct {
	Payload *models.Blueprint
}

// IsSuccess returns true when this restore blueprint version using p o s t1 o k response has a 2xx status code
func (o *RestoreBlueprintVersionUsingPOST1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore blueprint version using p o s t1 o k response has a 3xx status code
func (o *RestoreBlueprintVersionUsingPOST1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore blueprint version using p o s t1 o k response has a 4xx status code
func (o *RestoreBlueprintVersionUsingPOST1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore blueprint version using p o s t1 o k response has a 5xx status code
func (o *RestoreBlueprintVersionUsingPOST1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore blueprint version using p o s t1 o k response a status code equal to that given
func (o *RestoreBlueprintVersionUsingPOST1OK) IsCode(code int) bool {
	return code == 200
}

func (o *RestoreBlueprintVersionUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *RestoreBlueprintVersionUsingPOST1OK) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *RestoreBlueprintVersionUsingPOST1OK) GetPayload() *models.Blueprint {
	return o.Payload
}

func (o *RestoreBlueprintVersionUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Blueprint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreBlueprintVersionUsingPOST1Unauthorized creates a RestoreBlueprintVersionUsingPOST1Unauthorized with default headers values
func NewRestoreBlueprintVersionUsingPOST1Unauthorized() *RestoreBlueprintVersionUsingPOST1Unauthorized {
	return &RestoreBlueprintVersionUsingPOST1Unauthorized{}
}

/*
RestoreBlueprintVersionUsingPOST1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RestoreBlueprintVersionUsingPOST1Unauthorized struct {
}

// IsSuccess returns true when this restore blueprint version using p o s t1 unauthorized response has a 2xx status code
func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore blueprint version using p o s t1 unauthorized response has a 3xx status code
func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore blueprint version using p o s t1 unauthorized response has a 4xx status code
func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore blueprint version using p o s t1 unauthorized response has a 5xx status code
func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this restore blueprint version using p o s t1 unauthorized response a status code equal to that given
func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1Unauthorized ", 401)
}

func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1Unauthorized ", 401)
}

func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestoreBlueprintVersionUsingPOST1Forbidden creates a RestoreBlueprintVersionUsingPOST1Forbidden with default headers values
func NewRestoreBlueprintVersionUsingPOST1Forbidden() *RestoreBlueprintVersionUsingPOST1Forbidden {
	return &RestoreBlueprintVersionUsingPOST1Forbidden{}
}

/*
RestoreBlueprintVersionUsingPOST1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RestoreBlueprintVersionUsingPOST1Forbidden struct {
}

// IsSuccess returns true when this restore blueprint version using p o s t1 forbidden response has a 2xx status code
func (o *RestoreBlueprintVersionUsingPOST1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore blueprint version using p o s t1 forbidden response has a 3xx status code
func (o *RestoreBlueprintVersionUsingPOST1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore blueprint version using p o s t1 forbidden response has a 4xx status code
func (o *RestoreBlueprintVersionUsingPOST1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore blueprint version using p o s t1 forbidden response has a 5xx status code
func (o *RestoreBlueprintVersionUsingPOST1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this restore blueprint version using p o s t1 forbidden response a status code equal to that given
func (o *RestoreBlueprintVersionUsingPOST1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RestoreBlueprintVersionUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1Forbidden ", 403)
}

func (o *RestoreBlueprintVersionUsingPOST1Forbidden) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1Forbidden ", 403)
}

func (o *RestoreBlueprintVersionUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestoreBlueprintVersionUsingPOST1NotFound creates a RestoreBlueprintVersionUsingPOST1NotFound with default headers values
func NewRestoreBlueprintVersionUsingPOST1NotFound() *RestoreBlueprintVersionUsingPOST1NotFound {
	return &RestoreBlueprintVersionUsingPOST1NotFound{}
}

/*
RestoreBlueprintVersionUsingPOST1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type RestoreBlueprintVersionUsingPOST1NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this restore blueprint version using p o s t1 not found response has a 2xx status code
func (o *RestoreBlueprintVersionUsingPOST1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore blueprint version using p o s t1 not found response has a 3xx status code
func (o *RestoreBlueprintVersionUsingPOST1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore blueprint version using p o s t1 not found response has a 4xx status code
func (o *RestoreBlueprintVersionUsingPOST1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore blueprint version using p o s t1 not found response has a 5xx status code
func (o *RestoreBlueprintVersionUsingPOST1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this restore blueprint version using p o s t1 not found response a status code equal to that given
func (o *RestoreBlueprintVersionUsingPOST1NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RestoreBlueprintVersionUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1NotFound  %+v", 404, o.Payload)
}

func (o *RestoreBlueprintVersionUsingPOST1NotFound) String() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1NotFound  %+v", 404, o.Payload)
}

func (o *RestoreBlueprintVersionUsingPOST1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RestoreBlueprintVersionUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
