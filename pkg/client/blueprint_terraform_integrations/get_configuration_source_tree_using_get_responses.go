// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetConfigurationSourceTreeUsingGETReader is a Reader for the GetConfigurationSourceTreeUsingGET structure.
type GetConfigurationSourceTreeUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationSourceTreeUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationSourceTreeUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationSourceTreeUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConfigurationSourceTreeUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationSourceTreeUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigurationSourceTreeUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationSourceTreeUsingGETOK creates a GetConfigurationSourceTreeUsingGETOK with default headers values
func NewGetConfigurationSourceTreeUsingGETOK() *GetConfigurationSourceTreeUsingGETOK {
	return &GetConfigurationSourceTreeUsingGETOK{}
}

/*
GetConfigurationSourceTreeUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetConfigurationSourceTreeUsingGETOK struct {
	Payload *models.FileTree
}

// IsSuccess returns true when this get configuration source tree using g e t o k response has a 2xx status code
func (o *GetConfigurationSourceTreeUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configuration source tree using g e t o k response has a 3xx status code
func (o *GetConfigurationSourceTreeUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration source tree using g e t o k response has a 4xx status code
func (o *GetConfigurationSourceTreeUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration source tree using g e t o k response has a 5xx status code
func (o *GetConfigurationSourceTreeUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration source tree using g e t o k response a status code equal to that given
func (o *GetConfigurationSourceTreeUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConfigurationSourceTreeUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationSourceTreeUsingGETOK) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationSourceTreeUsingGETOK) GetPayload() *models.FileTree {
	return o.Payload
}

func (o *GetConfigurationSourceTreeUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileTree)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSourceTreeUsingGETBadRequest creates a GetConfigurationSourceTreeUsingGETBadRequest with default headers values
func NewGetConfigurationSourceTreeUsingGETBadRequest() *GetConfigurationSourceTreeUsingGETBadRequest {
	return &GetConfigurationSourceTreeUsingGETBadRequest{}
}

/*
GetConfigurationSourceTreeUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetConfigurationSourceTreeUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get configuration source tree using g e t bad request response has a 2xx status code
func (o *GetConfigurationSourceTreeUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration source tree using g e t bad request response has a 3xx status code
func (o *GetConfigurationSourceTreeUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration source tree using g e t bad request response has a 4xx status code
func (o *GetConfigurationSourceTreeUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration source tree using g e t bad request response has a 5xx status code
func (o *GetConfigurationSourceTreeUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration source tree using g e t bad request response a status code equal to that given
func (o *GetConfigurationSourceTreeUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConfigurationSourceTreeUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationSourceTreeUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationSourceTreeUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConfigurationSourceTreeUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSourceTreeUsingGETUnauthorized creates a GetConfigurationSourceTreeUsingGETUnauthorized with default headers values
func NewGetConfigurationSourceTreeUsingGETUnauthorized() *GetConfigurationSourceTreeUsingGETUnauthorized {
	return &GetConfigurationSourceTreeUsingGETUnauthorized{}
}

/*
GetConfigurationSourceTreeUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetConfigurationSourceTreeUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get configuration source tree using g e t unauthorized response has a 2xx status code
func (o *GetConfigurationSourceTreeUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration source tree using g e t unauthorized response has a 3xx status code
func (o *GetConfigurationSourceTreeUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration source tree using g e t unauthorized response has a 4xx status code
func (o *GetConfigurationSourceTreeUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration source tree using g e t unauthorized response has a 5xx status code
func (o *GetConfigurationSourceTreeUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration source tree using g e t unauthorized response a status code equal to that given
func (o *GetConfigurationSourceTreeUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConfigurationSourceTreeUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETUnauthorized ", 401)
}

func (o *GetConfigurationSourceTreeUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETUnauthorized ", 401)
}

func (o *GetConfigurationSourceTreeUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationSourceTreeUsingGETForbidden creates a GetConfigurationSourceTreeUsingGETForbidden with default headers values
func NewGetConfigurationSourceTreeUsingGETForbidden() *GetConfigurationSourceTreeUsingGETForbidden {
	return &GetConfigurationSourceTreeUsingGETForbidden{}
}

/*
GetConfigurationSourceTreeUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetConfigurationSourceTreeUsingGETForbidden struct {
}

// IsSuccess returns true when this get configuration source tree using g e t forbidden response has a 2xx status code
func (o *GetConfigurationSourceTreeUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration source tree using g e t forbidden response has a 3xx status code
func (o *GetConfigurationSourceTreeUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration source tree using g e t forbidden response has a 4xx status code
func (o *GetConfigurationSourceTreeUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration source tree using g e t forbidden response has a 5xx status code
func (o *GetConfigurationSourceTreeUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration source tree using g e t forbidden response a status code equal to that given
func (o *GetConfigurationSourceTreeUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConfigurationSourceTreeUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETForbidden ", 403)
}

func (o *GetConfigurationSourceTreeUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETForbidden ", 403)
}

func (o *GetConfigurationSourceTreeUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationSourceTreeUsingGETNotFound creates a GetConfigurationSourceTreeUsingGETNotFound with default headers values
func NewGetConfigurationSourceTreeUsingGETNotFound() *GetConfigurationSourceTreeUsingGETNotFound {
	return &GetConfigurationSourceTreeUsingGETNotFound{}
}

/*
GetConfigurationSourceTreeUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetConfigurationSourceTreeUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get configuration source tree using g e t not found response has a 2xx status code
func (o *GetConfigurationSourceTreeUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration source tree using g e t not found response has a 3xx status code
func (o *GetConfigurationSourceTreeUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration source tree using g e t not found response has a 4xx status code
func (o *GetConfigurationSourceTreeUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration source tree using g e t not found response has a 5xx status code
func (o *GetConfigurationSourceTreeUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration source tree using g e t not found response a status code equal to that given
func (o *GetConfigurationSourceTreeUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConfigurationSourceTreeUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigurationSourceTreeUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/get-configuration-source-tree][%d] getConfigurationSourceTreeUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigurationSourceTreeUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConfigurationSourceTreeUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}