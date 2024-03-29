// Code generated by go-swagger; DO NOT EDIT.

package catalog_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PostUsingPOST2Reader is a Reader for the PostUsingPOST2 structure.
type PostUsingPOST2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsingPOST2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUsingPOST2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostUsingPOST2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsingPOST2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUsingPOST2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsingPOST2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUsingPOST2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUsingPOST2OK creates a PostUsingPOST2OK with default headers values
func NewPostUsingPOST2OK() *PostUsingPOST2OK {
	return &PostUsingPOST2OK{}
}

/*
PostUsingPOST2OK describes a response with status code 200, with default header values.

Validation is ok
*/
type PostUsingPOST2OK struct {
	Payload *models.CatalogSource
}

// IsSuccess returns true when this post using p o s t2 o k response has a 2xx status code
func (o *PostUsingPOST2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post using p o s t2 o k response has a 3xx status code
func (o *PostUsingPOST2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post using p o s t2 o k response has a 4xx status code
func (o *PostUsingPOST2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post using p o s t2 o k response has a 5xx status code
func (o *PostUsingPOST2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this post using p o s t2 o k response a status code equal to that given
func (o *PostUsingPOST2OK) IsCode(code int) bool {
	return code == 200
}

func (o *PostUsingPOST2OK) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2OK  %+v", 200, o.Payload)
}

func (o *PostUsingPOST2OK) String() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2OK  %+v", 200, o.Payload)
}

func (o *PostUsingPOST2OK) GetPayload() *models.CatalogSource {
	return o.Payload
}

func (o *PostUsingPOST2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsingPOST2Created creates a PostUsingPOST2Created with default headers values
func NewPostUsingPOST2Created() *PostUsingPOST2Created {
	return &PostUsingPOST2Created{}
}

/*
PostUsingPOST2Created describes a response with status code 201, with default header values.

Created
*/
type PostUsingPOST2Created struct {
	Payload *models.CatalogSource
}

// IsSuccess returns true when this post using p o s t2 created response has a 2xx status code
func (o *PostUsingPOST2Created) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post using p o s t2 created response has a 3xx status code
func (o *PostUsingPOST2Created) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post using p o s t2 created response has a 4xx status code
func (o *PostUsingPOST2Created) IsClientError() bool {
	return false
}

// IsServerError returns true when this post using p o s t2 created response has a 5xx status code
func (o *PostUsingPOST2Created) IsServerError() bool {
	return false
}

// IsCode returns true when this post using p o s t2 created response a status code equal to that given
func (o *PostUsingPOST2Created) IsCode(code int) bool {
	return code == 201
}

func (o *PostUsingPOST2Created) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2Created  %+v", 201, o.Payload)
}

func (o *PostUsingPOST2Created) String() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2Created  %+v", 201, o.Payload)
}

func (o *PostUsingPOST2Created) GetPayload() *models.CatalogSource {
	return o.Payload
}

func (o *PostUsingPOST2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsingPOST2BadRequest creates a PostUsingPOST2BadRequest with default headers values
func NewPostUsingPOST2BadRequest() *PostUsingPOST2BadRequest {
	return &PostUsingPOST2BadRequest{}
}

/*
PostUsingPOST2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostUsingPOST2BadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this post using p o s t2 bad request response has a 2xx status code
func (o *PostUsingPOST2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post using p o s t2 bad request response has a 3xx status code
func (o *PostUsingPOST2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post using p o s t2 bad request response has a 4xx status code
func (o *PostUsingPOST2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post using p o s t2 bad request response has a 5xx status code
func (o *PostUsingPOST2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post using p o s t2 bad request response a status code equal to that given
func (o *PostUsingPOST2BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUsingPOST2BadRequest) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2BadRequest  %+v", 400, o.Payload)
}

func (o *PostUsingPOST2BadRequest) String() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2BadRequest  %+v", 400, o.Payload)
}

func (o *PostUsingPOST2BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsingPOST2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsingPOST2Unauthorized creates a PostUsingPOST2Unauthorized with default headers values
func NewPostUsingPOST2Unauthorized() *PostUsingPOST2Unauthorized {
	return &PostUsingPOST2Unauthorized{}
}

/*
PostUsingPOST2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostUsingPOST2Unauthorized struct {
}

// IsSuccess returns true when this post using p o s t2 unauthorized response has a 2xx status code
func (o *PostUsingPOST2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post using p o s t2 unauthorized response has a 3xx status code
func (o *PostUsingPOST2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post using p o s t2 unauthorized response has a 4xx status code
func (o *PostUsingPOST2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post using p o s t2 unauthorized response has a 5xx status code
func (o *PostUsingPOST2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post using p o s t2 unauthorized response a status code equal to that given
func (o *PostUsingPOST2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUsingPOST2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2Unauthorized ", 401)
}

func (o *PostUsingPOST2Unauthorized) String() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2Unauthorized ", 401)
}

func (o *PostUsingPOST2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUsingPOST2Forbidden creates a PostUsingPOST2Forbidden with default headers values
func NewPostUsingPOST2Forbidden() *PostUsingPOST2Forbidden {
	return &PostUsingPOST2Forbidden{}
}

/*
PostUsingPOST2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PostUsingPOST2Forbidden struct {
}

// IsSuccess returns true when this post using p o s t2 forbidden response has a 2xx status code
func (o *PostUsingPOST2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post using p o s t2 forbidden response has a 3xx status code
func (o *PostUsingPOST2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post using p o s t2 forbidden response has a 4xx status code
func (o *PostUsingPOST2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post using p o s t2 forbidden response has a 5xx status code
func (o *PostUsingPOST2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post using p o s t2 forbidden response a status code equal to that given
func (o *PostUsingPOST2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUsingPOST2Forbidden) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2Forbidden ", 403)
}

func (o *PostUsingPOST2Forbidden) String() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2Forbidden ", 403)
}

func (o *PostUsingPOST2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUsingPOST2NotFound creates a PostUsingPOST2NotFound with default headers values
func NewPostUsingPOST2NotFound() *PostUsingPOST2NotFound {
	return &PostUsingPOST2NotFound{}
}

/*
PostUsingPOST2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostUsingPOST2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this post using p o s t2 not found response has a 2xx status code
func (o *PostUsingPOST2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post using p o s t2 not found response has a 3xx status code
func (o *PostUsingPOST2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post using p o s t2 not found response has a 4xx status code
func (o *PostUsingPOST2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post using p o s t2 not found response has a 5xx status code
func (o *PostUsingPOST2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post using p o s t2 not found response a status code equal to that given
func (o *PostUsingPOST2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostUsingPOST2NotFound) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2NotFound  %+v", 404, o.Payload)
}

func (o *PostUsingPOST2NotFound) String() string {
	return fmt.Sprintf("[POST /catalog/api/admin/sources][%d] postUsingPOST2NotFound  %+v", 404, o.Payload)
}

func (o *PostUsingPOST2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostUsingPOST2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
