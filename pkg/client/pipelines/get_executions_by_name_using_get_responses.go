// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetExecutionsByNameUsingGETReader is a Reader for the GetExecutionsByNameUsingGET structure.
type GetExecutionsByNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExecutionsByNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExecutionsByNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExecutionsByNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExecutionsByNameUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExecutionsByNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExecutionsByNameUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExecutionsByNameUsingGETOK creates a GetExecutionsByNameUsingGETOK with default headers values
func NewGetExecutionsByNameUsingGETOK() *GetExecutionsByNameUsingGETOK {
	return &GetExecutionsByNameUsingGETOK{}
}

/* GetExecutionsByNameUsingGETOK describes a response with status code 200, with default header values.

'Success' with Executions on pages
*/
type GetExecutionsByNameUsingGETOK struct {
	Payload models.Executions
}

func (o *GetExecutionsByNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetExecutionsByNameUsingGETOK) GetPayload() models.Executions {
	return o.Payload
}

func (o *GetExecutionsByNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutions(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetExecutionsByNameUsingGETUnauthorized creates a GetExecutionsByNameUsingGETUnauthorized with default headers values
func NewGetExecutionsByNameUsingGETUnauthorized() *GetExecutionsByNameUsingGETUnauthorized {
	return &GetExecutionsByNameUsingGETUnauthorized{}
}

/* GetExecutionsByNameUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetExecutionsByNameUsingGETUnauthorized struct {
}

func (o *GetExecutionsByNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETUnauthorized ", 401)
}

func (o *GetExecutionsByNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionsByNameUsingGETForbidden creates a GetExecutionsByNameUsingGETForbidden with default headers values
func NewGetExecutionsByNameUsingGETForbidden() *GetExecutionsByNameUsingGETForbidden {
	return &GetExecutionsByNameUsingGETForbidden{}
}

/* GetExecutionsByNameUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetExecutionsByNameUsingGETForbidden struct {
}

func (o *GetExecutionsByNameUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETForbidden ", 403)
}

func (o *GetExecutionsByNameUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionsByNameUsingGETNotFound creates a GetExecutionsByNameUsingGETNotFound with default headers values
func NewGetExecutionsByNameUsingGETNotFound() *GetExecutionsByNameUsingGETNotFound {
	return &GetExecutionsByNameUsingGETNotFound{}
}

/* GetExecutionsByNameUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetExecutionsByNameUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetExecutionsByNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetExecutionsByNameUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExecutionsByNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionsByNameUsingGETInternalServerError creates a GetExecutionsByNameUsingGETInternalServerError with default headers values
func NewGetExecutionsByNameUsingGETInternalServerError() *GetExecutionsByNameUsingGETInternalServerError {
	return &GetExecutionsByNameUsingGETInternalServerError{}
}

/* GetExecutionsByNameUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetExecutionsByNameUsingGETInternalServerError struct {
}

func (o *GetExecutionsByNameUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{project}/{name}/executions][%d] getExecutionsByNameUsingGETInternalServerError ", 500)
}

func (o *GetExecutionsByNameUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
