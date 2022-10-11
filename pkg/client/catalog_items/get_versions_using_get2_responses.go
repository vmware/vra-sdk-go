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

// GetVersionsUsingGET2Reader is a Reader for the GetVersionsUsingGET2 structure.
type GetVersionsUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionsUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionsUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVersionsUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVersionsUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVersionsUsingGET2OK creates a GetVersionsUsingGET2OK with default headers values
func NewGetVersionsUsingGET2OK() *GetVersionsUsingGET2OK {
	return &GetVersionsUsingGET2OK{}
}

/*
GetVersionsUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetVersionsUsingGET2OK struct {
	Payload *models.PageOfCatalogItemVersion
}

// IsSuccess returns true when this get versions using g e t2 o k response has a 2xx status code
func (o *GetVersionsUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get versions using g e t2 o k response has a 3xx status code
func (o *GetVersionsUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions using g e t2 o k response has a 4xx status code
func (o *GetVersionsUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get versions using g e t2 o k response has a 5xx status code
func (o *GetVersionsUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions using g e t2 o k response a status code equal to that given
func (o *GetVersionsUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVersionsUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions][%d] getVersionsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetVersionsUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions][%d] getVersionsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetVersionsUsingGET2OK) GetPayload() *models.PageOfCatalogItemVersion {
	return o.Payload
}

func (o *GetVersionsUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfCatalogItemVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionsUsingGET2Unauthorized creates a GetVersionsUsingGET2Unauthorized with default headers values
func NewGetVersionsUsingGET2Unauthorized() *GetVersionsUsingGET2Unauthorized {
	return &GetVersionsUsingGET2Unauthorized{}
}

/*
GetVersionsUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVersionsUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get versions using g e t2 unauthorized response has a 2xx status code
func (o *GetVersionsUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versions using g e t2 unauthorized response has a 3xx status code
func (o *GetVersionsUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions using g e t2 unauthorized response has a 4xx status code
func (o *GetVersionsUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versions using g e t2 unauthorized response has a 5xx status code
func (o *GetVersionsUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions using g e t2 unauthorized response a status code equal to that given
func (o *GetVersionsUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetVersionsUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions][%d] getVersionsUsingGET2Unauthorized ", 401)
}

func (o *GetVersionsUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions][%d] getVersionsUsingGET2Unauthorized ", 401)
}

func (o *GetVersionsUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersionsUsingGET2NotFound creates a GetVersionsUsingGET2NotFound with default headers values
func NewGetVersionsUsingGET2NotFound() *GetVersionsUsingGET2NotFound {
	return &GetVersionsUsingGET2NotFound{}
}

/*
GetVersionsUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetVersionsUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get versions using g e t2 not found response has a 2xx status code
func (o *GetVersionsUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versions using g e t2 not found response has a 3xx status code
func (o *GetVersionsUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versions using g e t2 not found response has a 4xx status code
func (o *GetVersionsUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versions using g e t2 not found response has a 5xx status code
func (o *GetVersionsUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get versions using g e t2 not found response a status code equal to that given
func (o *GetVersionsUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVersionsUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions][%d] getVersionsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetVersionsUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /catalog/api/items/{id}/versions][%d] getVersionsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetVersionsUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVersionsUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
