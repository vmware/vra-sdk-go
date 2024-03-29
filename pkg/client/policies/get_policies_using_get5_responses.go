// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetPoliciesUsingGET5Reader is a Reader for the GetPoliciesUsingGET5 structure.
type GetPoliciesUsingGET5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPoliciesUsingGET5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPoliciesUsingGET5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPoliciesUsingGET5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPoliciesUsingGET5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPoliciesUsingGET5OK creates a GetPoliciesUsingGET5OK with default headers values
func NewGetPoliciesUsingGET5OK() *GetPoliciesUsingGET5OK {
	return &GetPoliciesUsingGET5OK{}
}

/*
GetPoliciesUsingGET5OK describes a response with status code 200, with default header values.

OK
*/
type GetPoliciesUsingGET5OK struct {
	Payload *models.PageOfPolicy
}

// IsSuccess returns true when this get policies using g e t5 o k response has a 2xx status code
func (o *GetPoliciesUsingGET5OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get policies using g e t5 o k response has a 3xx status code
func (o *GetPoliciesUsingGET5OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policies using g e t5 o k response has a 4xx status code
func (o *GetPoliciesUsingGET5OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get policies using g e t5 o k response has a 5xx status code
func (o *GetPoliciesUsingGET5OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get policies using g e t5 o k response a status code equal to that given
func (o *GetPoliciesUsingGET5OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPoliciesUsingGET5OK) Error() string {
	return fmt.Sprintf("[GET /policy/api/policies][%d] getPoliciesUsingGET5OK  %+v", 200, o.Payload)
}

func (o *GetPoliciesUsingGET5OK) String() string {
	return fmt.Sprintf("[GET /policy/api/policies][%d] getPoliciesUsingGET5OK  %+v", 200, o.Payload)
}

func (o *GetPoliciesUsingGET5OK) GetPayload() *models.PageOfPolicy {
	return o.Payload
}

func (o *GetPoliciesUsingGET5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPoliciesUsingGET5Unauthorized creates a GetPoliciesUsingGET5Unauthorized with default headers values
func NewGetPoliciesUsingGET5Unauthorized() *GetPoliciesUsingGET5Unauthorized {
	return &GetPoliciesUsingGET5Unauthorized{}
}

/*
GetPoliciesUsingGET5Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPoliciesUsingGET5Unauthorized struct {
}

// IsSuccess returns true when this get policies using g e t5 unauthorized response has a 2xx status code
func (o *GetPoliciesUsingGET5Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policies using g e t5 unauthorized response has a 3xx status code
func (o *GetPoliciesUsingGET5Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policies using g e t5 unauthorized response has a 4xx status code
func (o *GetPoliciesUsingGET5Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policies using g e t5 unauthorized response has a 5xx status code
func (o *GetPoliciesUsingGET5Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get policies using g e t5 unauthorized response a status code equal to that given
func (o *GetPoliciesUsingGET5Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPoliciesUsingGET5Unauthorized) Error() string {
	return fmt.Sprintf("[GET /policy/api/policies][%d] getPoliciesUsingGET5Unauthorized ", 401)
}

func (o *GetPoliciesUsingGET5Unauthorized) String() string {
	return fmt.Sprintf("[GET /policy/api/policies][%d] getPoliciesUsingGET5Unauthorized ", 401)
}

func (o *GetPoliciesUsingGET5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPoliciesUsingGET5Forbidden creates a GetPoliciesUsingGET5Forbidden with default headers values
func NewGetPoliciesUsingGET5Forbidden() *GetPoliciesUsingGET5Forbidden {
	return &GetPoliciesUsingGET5Forbidden{}
}

/*
GetPoliciesUsingGET5Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPoliciesUsingGET5Forbidden struct {
}

// IsSuccess returns true when this get policies using g e t5 forbidden response has a 2xx status code
func (o *GetPoliciesUsingGET5Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get policies using g e t5 forbidden response has a 3xx status code
func (o *GetPoliciesUsingGET5Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get policies using g e t5 forbidden response has a 4xx status code
func (o *GetPoliciesUsingGET5Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get policies using g e t5 forbidden response has a 5xx status code
func (o *GetPoliciesUsingGET5Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get policies using g e t5 forbidden response a status code equal to that given
func (o *GetPoliciesUsingGET5Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPoliciesUsingGET5Forbidden) Error() string {
	return fmt.Sprintf("[GET /policy/api/policies][%d] getPoliciesUsingGET5Forbidden ", 403)
}

func (o *GetPoliciesUsingGET5Forbidden) String() string {
	return fmt.Sprintf("[GET /policy/api/policies][%d] getPoliciesUsingGET5Forbidden ", 403)
}

func (o *GetPoliciesUsingGET5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
