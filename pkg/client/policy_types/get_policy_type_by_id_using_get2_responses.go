// Code generated by go-swagger; DO NOT EDIT.

package policy_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPolicyTypeByIDUsingGET2Reader is a Reader for the GetPolicyTypeByIDUsingGET2 structure.
type GetPolicyTypeByIDUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyTypeByIDUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyTypeByIDUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPolicyTypeByIDUsingGET2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPolicyTypeByIDUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPolicyTypeByIDUsingGET2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPolicyTypeByIDUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPolicyTypeByIDUsingGET2OK creates a GetPolicyTypeByIDUsingGET2OK with default headers values
func NewGetPolicyTypeByIDUsingGET2OK() *GetPolicyTypeByIDUsingGET2OK {
	return &GetPolicyTypeByIDUsingGET2OK{}
}

/*
GetPolicyTypeByIDUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetPolicyTypeByIDUsingGET2OK struct {
	Payload *models.PolicyType
}

// IsSuccess returns true when this get policy type by Id using g e t2 o k response has a 2xx status code
func (o *GetPolicyTypeByIDUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policy type by Id using g e t2 o k response has a 3xx status code
func (o *GetPolicyTypeByIDUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy type by Id using g e t2 o k response has a 4xx status code
func (o *GetPolicyTypeByIDUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policy type by Id using g e t2 o k response has a 5xx status code
func (o *GetPolicyTypeByIDUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy type by Id using g e t2 o k response a status code equal to that given
func (o *GetPolicyTypeByIDUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPolicyTypeByIDUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetPolicyTypeByIDUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetPolicyTypeByIDUsingGET2OK) GetPayload() *models.PolicyType {
	return o.Payload
}

func (o *GetPolicyTypeByIDUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyTypeByIDUsingGET2BadRequest creates a GetPolicyTypeByIDUsingGET2BadRequest with default headers values
func NewGetPolicyTypeByIDUsingGET2BadRequest() *GetPolicyTypeByIDUsingGET2BadRequest {
	return &GetPolicyTypeByIDUsingGET2BadRequest{}
}

/*
GetPolicyTypeByIDUsingGET2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetPolicyTypeByIDUsingGET2BadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get policy type by Id using g e t2 bad request response has a 2xx status code
func (o *GetPolicyTypeByIDUsingGET2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy type by Id using g e t2 bad request response has a 3xx status code
func (o *GetPolicyTypeByIDUsingGET2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy type by Id using g e t2 bad request response has a 4xx status code
func (o *GetPolicyTypeByIDUsingGET2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy type by Id using g e t2 bad request response has a 5xx status code
func (o *GetPolicyTypeByIDUsingGET2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy type by Id using g e t2 bad request response a status code equal to that given
func (o *GetPolicyTypeByIDUsingGET2BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPolicyTypeByIDUsingGET2BadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2BadRequest  %+v", 400, o.Payload)
}

func (o *GetPolicyTypeByIDUsingGET2BadRequest) String() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2BadRequest  %+v", 400, o.Payload)
}

func (o *GetPolicyTypeByIDUsingGET2BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPolicyTypeByIDUsingGET2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyTypeByIDUsingGET2Unauthorized creates a GetPolicyTypeByIDUsingGET2Unauthorized with default headers values
func NewGetPolicyTypeByIDUsingGET2Unauthorized() *GetPolicyTypeByIDUsingGET2Unauthorized {
	return &GetPolicyTypeByIDUsingGET2Unauthorized{}
}

/*
GetPolicyTypeByIDUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPolicyTypeByIDUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get policy type by Id using g e t2 unauthorized response has a 2xx status code
func (o *GetPolicyTypeByIDUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy type by Id using g e t2 unauthorized response has a 3xx status code
func (o *GetPolicyTypeByIDUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy type by Id using g e t2 unauthorized response has a 4xx status code
func (o *GetPolicyTypeByIDUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy type by Id using g e t2 unauthorized response has a 5xx status code
func (o *GetPolicyTypeByIDUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy type by Id using g e t2 unauthorized response a status code equal to that given
func (o *GetPolicyTypeByIDUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPolicyTypeByIDUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2Unauthorized ", 401)
}

func (o *GetPolicyTypeByIDUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2Unauthorized ", 401)
}

func (o *GetPolicyTypeByIDUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyTypeByIDUsingGET2Forbidden creates a GetPolicyTypeByIDUsingGET2Forbidden with default headers values
func NewGetPolicyTypeByIDUsingGET2Forbidden() *GetPolicyTypeByIDUsingGET2Forbidden {
	return &GetPolicyTypeByIDUsingGET2Forbidden{}
}

/*
GetPolicyTypeByIDUsingGET2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPolicyTypeByIDUsingGET2Forbidden struct {
}

// IsSuccess returns true when this get policy type by Id using g e t2 forbidden response has a 2xx status code
func (o *GetPolicyTypeByIDUsingGET2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy type by Id using g e t2 forbidden response has a 3xx status code
func (o *GetPolicyTypeByIDUsingGET2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy type by Id using g e t2 forbidden response has a 4xx status code
func (o *GetPolicyTypeByIDUsingGET2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy type by Id using g e t2 forbidden response has a 5xx status code
func (o *GetPolicyTypeByIDUsingGET2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy type by Id using g e t2 forbidden response a status code equal to that given
func (o *GetPolicyTypeByIDUsingGET2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPolicyTypeByIDUsingGET2Forbidden) Error() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2Forbidden ", 403)
}

func (o *GetPolicyTypeByIDUsingGET2Forbidden) String() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2Forbidden ", 403)
}

func (o *GetPolicyTypeByIDUsingGET2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPolicyTypeByIDUsingGET2NotFound creates a GetPolicyTypeByIDUsingGET2NotFound with default headers values
func NewGetPolicyTypeByIDUsingGET2NotFound() *GetPolicyTypeByIDUsingGET2NotFound {
	return &GetPolicyTypeByIDUsingGET2NotFound{}
}

/*
GetPolicyTypeByIDUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPolicyTypeByIDUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get policy type by Id using g e t2 not found response has a 2xx status code
func (o *GetPolicyTypeByIDUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policy type by Id using g e t2 not found response has a 3xx status code
func (o *GetPolicyTypeByIDUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policy type by Id using g e t2 not found response has a 4xx status code
func (o *GetPolicyTypeByIDUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policy type by Id using g e t2 not found response has a 5xx status code
func (o *GetPolicyTypeByIDUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get policy type by Id using g e t2 not found response a status code equal to that given
func (o *GetPolicyTypeByIDUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPolicyTypeByIDUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetPolicyTypeByIDUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes/{id}][%d] getPolicyTypeByIdUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetPolicyTypeByIDUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPolicyTypeByIDUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
