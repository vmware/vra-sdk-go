// Code generated by go-swagger; DO NOT EDIT.

package property_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPropertyGroupUsingGETReader is a Reader for the GetPropertyGroupUsingGET structure.
type GetPropertyGroupUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPropertyGroupUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPropertyGroupUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPropertyGroupUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPropertyGroupUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPropertyGroupUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPropertyGroupUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPropertyGroupUsingGETOK creates a GetPropertyGroupUsingGETOK with default headers values
func NewGetPropertyGroupUsingGETOK() *GetPropertyGroupUsingGETOK {
	return &GetPropertyGroupUsingGETOK{}
}

/*
GetPropertyGroupUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetPropertyGroupUsingGETOK struct {
	Payload *models.PropertyGroup
}

// IsSuccess returns true when this get property group using g e t o k response has a 2xx status code
func (o *GetPropertyGroupUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get property group using g e t o k response has a 3xx status code
func (o *GetPropertyGroupUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get property group using g e t o k response has a 4xx status code
func (o *GetPropertyGroupUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get property group using g e t o k response has a 5xx status code
func (o *GetPropertyGroupUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get property group using g e t o k response a status code equal to that given
func (o *GetPropertyGroupUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPropertyGroupUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPropertyGroupUsingGETOK) String() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPropertyGroupUsingGETOK) GetPayload() *models.PropertyGroup {
	return o.Payload
}

func (o *GetPropertyGroupUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPropertyGroupUsingGETBadRequest creates a GetPropertyGroupUsingGETBadRequest with default headers values
func NewGetPropertyGroupUsingGETBadRequest() *GetPropertyGroupUsingGETBadRequest {
	return &GetPropertyGroupUsingGETBadRequest{}
}

/*
GetPropertyGroupUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPropertyGroupUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get property group using g e t bad request response has a 2xx status code
func (o *GetPropertyGroupUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get property group using g e t bad request response has a 3xx status code
func (o *GetPropertyGroupUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get property group using g e t bad request response has a 4xx status code
func (o *GetPropertyGroupUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get property group using g e t bad request response has a 5xx status code
func (o *GetPropertyGroupUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get property group using g e t bad request response a status code equal to that given
func (o *GetPropertyGroupUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPropertyGroupUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetPropertyGroupUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetPropertyGroupUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPropertyGroupUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPropertyGroupUsingGETUnauthorized creates a GetPropertyGroupUsingGETUnauthorized with default headers values
func NewGetPropertyGroupUsingGETUnauthorized() *GetPropertyGroupUsingGETUnauthorized {
	return &GetPropertyGroupUsingGETUnauthorized{}
}

/*
GetPropertyGroupUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPropertyGroupUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get property group using g e t unauthorized response has a 2xx status code
func (o *GetPropertyGroupUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get property group using g e t unauthorized response has a 3xx status code
func (o *GetPropertyGroupUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get property group using g e t unauthorized response has a 4xx status code
func (o *GetPropertyGroupUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get property group using g e t unauthorized response has a 5xx status code
func (o *GetPropertyGroupUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get property group using g e t unauthorized response a status code equal to that given
func (o *GetPropertyGroupUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPropertyGroupUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETUnauthorized ", 401)
}

func (o *GetPropertyGroupUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETUnauthorized ", 401)
}

func (o *GetPropertyGroupUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyGroupUsingGETForbidden creates a GetPropertyGroupUsingGETForbidden with default headers values
func NewGetPropertyGroupUsingGETForbidden() *GetPropertyGroupUsingGETForbidden {
	return &GetPropertyGroupUsingGETForbidden{}
}

/*
GetPropertyGroupUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPropertyGroupUsingGETForbidden struct {
}

// IsSuccess returns true when this get property group using g e t forbidden response has a 2xx status code
func (o *GetPropertyGroupUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get property group using g e t forbidden response has a 3xx status code
func (o *GetPropertyGroupUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get property group using g e t forbidden response has a 4xx status code
func (o *GetPropertyGroupUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get property group using g e t forbidden response has a 5xx status code
func (o *GetPropertyGroupUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get property group using g e t forbidden response a status code equal to that given
func (o *GetPropertyGroupUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPropertyGroupUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETForbidden ", 403)
}

func (o *GetPropertyGroupUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETForbidden ", 403)
}

func (o *GetPropertyGroupUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyGroupUsingGETNotFound creates a GetPropertyGroupUsingGETNotFound with default headers values
func NewGetPropertyGroupUsingGETNotFound() *GetPropertyGroupUsingGETNotFound {
	return &GetPropertyGroupUsingGETNotFound{}
}

/*
GetPropertyGroupUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPropertyGroupUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get property group using g e t not found response has a 2xx status code
func (o *GetPropertyGroupUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get property group using g e t not found response has a 3xx status code
func (o *GetPropertyGroupUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get property group using g e t not found response has a 4xx status code
func (o *GetPropertyGroupUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get property group using g e t not found response has a 5xx status code
func (o *GetPropertyGroupUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get property group using g e t not found response a status code equal to that given
func (o *GetPropertyGroupUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPropertyGroupUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetPropertyGroupUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /properties/api/property-groups/{propertyGroupId}][%d] getPropertyGroupUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetPropertyGroupUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPropertyGroupUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
