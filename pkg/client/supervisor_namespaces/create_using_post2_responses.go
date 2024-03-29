// Code generated by go-swagger; DO NOT EDIT.

package supervisor_namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateUsingPOST2Reader is a Reader for the CreateUsingPOST2 structure.
type CreateUsingPOST2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingPOST2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUsingPOST2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateUsingPOST2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUsingPOST2OK creates a CreateUsingPOST2OK with default headers values
func NewCreateUsingPOST2OK() *CreateUsingPOST2OK {
	return &CreateUsingPOST2OK{}
}

/*
CreateUsingPOST2OK describes a response with status code 200, with default header values.

OK
*/
type CreateUsingPOST2OK struct {
	Payload *models.SupervisorNamespace
}

// IsSuccess returns true when this create using p o s t2 o k response has a 2xx status code
func (o *CreateUsingPOST2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create using p o s t2 o k response has a 3xx status code
func (o *CreateUsingPOST2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create using p o s t2 o k response has a 4xx status code
func (o *CreateUsingPOST2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create using p o s t2 o k response has a 5xx status code
func (o *CreateUsingPOST2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this create using p o s t2 o k response a status code equal to that given
func (o *CreateUsingPOST2OK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateUsingPOST2OK) Error() string {
	return fmt.Sprintf("[POST /cmx/api/resources/supervisor-namespaces][%d] createUsingPOST2OK  %+v", 200, o.Payload)
}

func (o *CreateUsingPOST2OK) String() string {
	return fmt.Sprintf("[POST /cmx/api/resources/supervisor-namespaces][%d] createUsingPOST2OK  %+v", 200, o.Payload)
}

func (o *CreateUsingPOST2OK) GetPayload() *models.SupervisorNamespace {
	return o.Payload
}

func (o *CreateUsingPOST2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupervisorNamespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUsingPOST2Forbidden creates a CreateUsingPOST2Forbidden with default headers values
func NewCreateUsingPOST2Forbidden() *CreateUsingPOST2Forbidden {
	return &CreateUsingPOST2Forbidden{}
}

/*
CreateUsingPOST2Forbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type CreateUsingPOST2Forbidden struct {
}

// IsSuccess returns true when this create using p o s t2 forbidden response has a 2xx status code
func (o *CreateUsingPOST2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create using p o s t2 forbidden response has a 3xx status code
func (o *CreateUsingPOST2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create using p o s t2 forbidden response has a 4xx status code
func (o *CreateUsingPOST2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create using p o s t2 forbidden response has a 5xx status code
func (o *CreateUsingPOST2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create using p o s t2 forbidden response a status code equal to that given
func (o *CreateUsingPOST2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateUsingPOST2Forbidden) Error() string {
	return fmt.Sprintf("[POST /cmx/api/resources/supervisor-namespaces][%d] createUsingPOST2Forbidden ", 403)
}

func (o *CreateUsingPOST2Forbidden) String() string {
	return fmt.Sprintf("[POST /cmx/api/resources/supervisor-namespaces][%d] createUsingPOST2Forbidden ", 403)
}

func (o *CreateUsingPOST2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
