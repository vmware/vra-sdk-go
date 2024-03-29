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

// ListTerraformVersionsUsingGETReader is a Reader for the ListTerraformVersionsUsingGET structure.
type ListTerraformVersionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTerraformVersionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTerraformVersionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListTerraformVersionsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListTerraformVersionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTerraformVersionsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTerraformVersionsUsingGETOK creates a ListTerraformVersionsUsingGETOK with default headers values
func NewListTerraformVersionsUsingGETOK() *ListTerraformVersionsUsingGETOK {
	return &ListTerraformVersionsUsingGETOK{}
}

/*
ListTerraformVersionsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type ListTerraformVersionsUsingGETOK struct {
	Payload *models.PageOfTerraformVersion
}

// IsSuccess returns true when this list terraform versions using g e t o k response has a 2xx status code
func (o *ListTerraformVersionsUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list terraform versions using g e t o k response has a 3xx status code
func (o *ListTerraformVersionsUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list terraform versions using g e t o k response has a 4xx status code
func (o *ListTerraformVersionsUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list terraform versions using g e t o k response has a 5xx status code
func (o *ListTerraformVersionsUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list terraform versions using g e t o k response a status code equal to that given
func (o *ListTerraformVersionsUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListTerraformVersionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *ListTerraformVersionsUsingGETOK) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *ListTerraformVersionsUsingGETOK) GetPayload() *models.PageOfTerraformVersion {
	return o.Payload
}

func (o *ListTerraformVersionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfTerraformVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTerraformVersionsUsingGETBadRequest creates a ListTerraformVersionsUsingGETBadRequest with default headers values
func NewListTerraformVersionsUsingGETBadRequest() *ListTerraformVersionsUsingGETBadRequest {
	return &ListTerraformVersionsUsingGETBadRequest{}
}

/*
ListTerraformVersionsUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListTerraformVersionsUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this list terraform versions using g e t bad request response has a 2xx status code
func (o *ListTerraformVersionsUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list terraform versions using g e t bad request response has a 3xx status code
func (o *ListTerraformVersionsUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list terraform versions using g e t bad request response has a 4xx status code
func (o *ListTerraformVersionsUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list terraform versions using g e t bad request response has a 5xx status code
func (o *ListTerraformVersionsUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list terraform versions using g e t bad request response a status code equal to that given
func (o *ListTerraformVersionsUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListTerraformVersionsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *ListTerraformVersionsUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *ListTerraformVersionsUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListTerraformVersionsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTerraformVersionsUsingGETUnauthorized creates a ListTerraformVersionsUsingGETUnauthorized with default headers values
func NewListTerraformVersionsUsingGETUnauthorized() *ListTerraformVersionsUsingGETUnauthorized {
	return &ListTerraformVersionsUsingGETUnauthorized{}
}

/*
ListTerraformVersionsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListTerraformVersionsUsingGETUnauthorized struct {
}

// IsSuccess returns true when this list terraform versions using g e t unauthorized response has a 2xx status code
func (o *ListTerraformVersionsUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list terraform versions using g e t unauthorized response has a 3xx status code
func (o *ListTerraformVersionsUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list terraform versions using g e t unauthorized response has a 4xx status code
func (o *ListTerraformVersionsUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list terraform versions using g e t unauthorized response has a 5xx status code
func (o *ListTerraformVersionsUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list terraform versions using g e t unauthorized response a status code equal to that given
func (o *ListTerraformVersionsUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListTerraformVersionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETUnauthorized ", 401)
}

func (o *ListTerraformVersionsUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETUnauthorized ", 401)
}

func (o *ListTerraformVersionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListTerraformVersionsUsingGETForbidden creates a ListTerraformVersionsUsingGETForbidden with default headers values
func NewListTerraformVersionsUsingGETForbidden() *ListTerraformVersionsUsingGETForbidden {
	return &ListTerraformVersionsUsingGETForbidden{}
}

/*
ListTerraformVersionsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListTerraformVersionsUsingGETForbidden struct {
}

// IsSuccess returns true when this list terraform versions using g e t forbidden response has a 2xx status code
func (o *ListTerraformVersionsUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list terraform versions using g e t forbidden response has a 3xx status code
func (o *ListTerraformVersionsUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list terraform versions using g e t forbidden response has a 4xx status code
func (o *ListTerraformVersionsUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list terraform versions using g e t forbidden response has a 5xx status code
func (o *ListTerraformVersionsUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list terraform versions using g e t forbidden response a status code equal to that given
func (o *ListTerraformVersionsUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListTerraformVersionsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETForbidden ", 403)
}

func (o *ListTerraformVersionsUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-integrations/terraform/versions][%d] listTerraformVersionsUsingGETForbidden ", 403)
}

func (o *ListTerraformVersionsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
