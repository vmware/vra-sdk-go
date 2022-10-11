// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetVersionByIDUsingGET2Reader is a Reader for the GetVersionByIDUsingGET2 structure.
type GetVersionByIDUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionByIDUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionByIDUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVersionByIDUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVersionByIDUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVersionByIDUsingGET2OK creates a GetVersionByIDUsingGET2OK with default headers values
func NewGetVersionByIDUsingGET2OK() *GetVersionByIDUsingGET2OK {
	return &GetVersionByIDUsingGET2OK{}
}

/*
GetVersionByIDUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetVersionByIDUsingGET2OK struct {
	Payload *models.CatalogItemVersion
}

// IsSuccess returns true when this get version by Id using g e t2 o k response has a 2xx status code
func (o *GetVersionByIDUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get version by Id using g e t2 o k response has a 3xx status code
func (o *GetVersionByIDUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version by Id using g e t2 o k response has a 4xx status code
func (o *GetVersionByIDUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get version by Id using g e t2 o k response has a 5xx status code
func (o *GetVersionByIDUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get version by Id using g e t2 o k response a status code equal to that given
func (o *GetVersionByIDUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVersionByIDUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions/{versionId}][%d] getVersionByIdUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetVersionByIDUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions/{versionId}][%d] getVersionByIdUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetVersionByIDUsingGET2OK) GetPayload() *models.CatalogItemVersion {
	return o.Payload
}

func (o *GetVersionByIDUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogItemVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionByIDUsingGET2Unauthorized creates a GetVersionByIDUsingGET2Unauthorized with default headers values
func NewGetVersionByIDUsingGET2Unauthorized() *GetVersionByIDUsingGET2Unauthorized {
	return &GetVersionByIDUsingGET2Unauthorized{}
}

/*
GetVersionByIDUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVersionByIDUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get version by Id using g e t2 unauthorized response has a 2xx status code
func (o *GetVersionByIDUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get version by Id using g e t2 unauthorized response has a 3xx status code
func (o *GetVersionByIDUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version by Id using g e t2 unauthorized response has a 4xx status code
func (o *GetVersionByIDUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get version by Id using g e t2 unauthorized response has a 5xx status code
func (o *GetVersionByIDUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get version by Id using g e t2 unauthorized response a status code equal to that given
func (o *GetVersionByIDUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetVersionByIDUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions/{versionId}][%d] getVersionByIdUsingGET2Unauthorized ", 401)
}

func (o *GetVersionByIDUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions/{versionId}][%d] getVersionByIdUsingGET2Unauthorized ", 401)
}

func (o *GetVersionByIDUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionByIDUsingGET2NotFound creates a GetVersionByIDUsingGET2NotFound with default headers values
func NewGetVersionByIDUsingGET2NotFound() *GetVersionByIDUsingGET2NotFound {
	return &GetVersionByIDUsingGET2NotFound{}
}

/*
GetVersionByIDUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVersionByIDUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get version by Id using g e t2 not found response has a 2xx status code
func (o *GetVersionByIDUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get version by Id using g e t2 not found response has a 3xx status code
func (o *GetVersionByIDUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get version by Id using g e t2 not found response has a 4xx status code
func (o *GetVersionByIDUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get version by Id using g e t2 not found response has a 5xx status code
func (o *GetVersionByIDUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get version by Id using g e t2 not found response a status code equal to that given
func (o *GetVersionByIDUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVersionByIDUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions/{versionId}][%d] getVersionByIdUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetVersionByIDUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions/{versionId}][%d] getVersionByIdUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetVersionByIDUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVersionByIDUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}