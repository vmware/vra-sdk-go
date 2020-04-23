// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
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
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRestoreBlueprintVersionUsingPOST1OK creates a RestoreBlueprintVersionUsingPOST1OK with default headers values
func NewRestoreBlueprintVersionUsingPOST1OK() *RestoreBlueprintVersionUsingPOST1OK {
	return &RestoreBlueprintVersionUsingPOST1OK{}
}

/*RestoreBlueprintVersionUsingPOST1OK handles this case with default header values.

OK
*/
type RestoreBlueprintVersionUsingPOST1OK struct {
	Payload *models.Blueprint
}

func (o *RestoreBlueprintVersionUsingPOST1OK) Error() string {
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

/*RestoreBlueprintVersionUsingPOST1Unauthorized handles this case with default header values.

Unauthorized
*/
type RestoreBlueprintVersionUsingPOST1Unauthorized struct {
}

func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1Unauthorized ", 401)
}

func (o *RestoreBlueprintVersionUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestoreBlueprintVersionUsingPOST1Forbidden creates a RestoreBlueprintVersionUsingPOST1Forbidden with default headers values
func NewRestoreBlueprintVersionUsingPOST1Forbidden() *RestoreBlueprintVersionUsingPOST1Forbidden {
	return &RestoreBlueprintVersionUsingPOST1Forbidden{}
}

/*RestoreBlueprintVersionUsingPOST1Forbidden handles this case with default header values.

Forbidden
*/
type RestoreBlueprintVersionUsingPOST1Forbidden struct {
}

func (o *RestoreBlueprintVersionUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /blueprint/api/blueprints/{blueprintId}/versions/{version}/actions/restore][%d] restoreBlueprintVersionUsingPOST1Forbidden ", 403)
}

func (o *RestoreBlueprintVersionUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestoreBlueprintVersionUsingPOST1NotFound creates a RestoreBlueprintVersionUsingPOST1NotFound with default headers values
func NewRestoreBlueprintVersionUsingPOST1NotFound() *RestoreBlueprintVersionUsingPOST1NotFound {
	return &RestoreBlueprintVersionUsingPOST1NotFound{}
}

/*RestoreBlueprintVersionUsingPOST1NotFound handles this case with default header values.

Not Found
*/
type RestoreBlueprintVersionUsingPOST1NotFound struct {
	Payload *models.Error
}

func (o *RestoreBlueprintVersionUsingPOST1NotFound) Error() string {
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
