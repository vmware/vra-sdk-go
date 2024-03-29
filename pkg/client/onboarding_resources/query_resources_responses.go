// Code generated by go-swagger; DO NOT EDIT.

package onboarding_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// QueryResourcesReader is a Reader for the QueryResources structure.
type QueryResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewQueryResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryResourcesOK creates a QueryResourcesOK with default headers values
func NewQueryResourcesOK() *QueryResourcesOK {
	return &QueryResourcesOK{}
}

/*
QueryResourcesOK describes a response with status code 200, with default header values.

Success
*/
type QueryResourcesOK struct {
	Payload *models.ServiceDocumentQueryResult
}

// IsSuccess returns true when this query resources o k response has a 2xx status code
func (o *QueryResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query resources o k response has a 3xx status code
func (o *QueryResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resources o k response has a 4xx status code
func (o *QueryResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query resources o k response has a 5xx status code
func (o *QueryResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query resources o k response a status code equal to that given
func (o *QueryResourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *QueryResourcesOK) Error() string {
	return fmt.Sprintf("[GET /relocation/onboarding/resource][%d] queryResourcesOK  %+v", 200, o.Payload)
}

func (o *QueryResourcesOK) String() string {
	return fmt.Sprintf("[GET /relocation/onboarding/resource][%d] queryResourcesOK  %+v", 200, o.Payload)
}

func (o *QueryResourcesOK) GetPayload() *models.ServiceDocumentQueryResult {
	return o.Payload
}

func (o *QueryResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDocumentQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryResourcesUnauthorized creates a QueryResourcesUnauthorized with default headers values
func NewQueryResourcesUnauthorized() *QueryResourcesUnauthorized {
	return &QueryResourcesUnauthorized{}
}

/*
QueryResourcesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type QueryResourcesUnauthorized struct {
}

// IsSuccess returns true when this query resources unauthorized response has a 2xx status code
func (o *QueryResourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query resources unauthorized response has a 3xx status code
func (o *QueryResourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query resources unauthorized response has a 4xx status code
func (o *QueryResourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this query resources unauthorized response has a 5xx status code
func (o *QueryResourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this query resources unauthorized response a status code equal to that given
func (o *QueryResourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *QueryResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /relocation/onboarding/resource][%d] queryResourcesUnauthorized ", 401)
}

func (o *QueryResourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /relocation/onboarding/resource][%d] queryResourcesUnauthorized ", 401)
}

func (o *QueryResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
