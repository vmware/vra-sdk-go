// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// RegisterUsingPUTReader is a Reader for the RegisterUsingPUT structure.
type RegisterUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRegisterUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterUsingPUTOK creates a RegisterUsingPUTOK with default headers values
func NewRegisterUsingPUTOK() *RegisterUsingPUTOK {
	return &RegisterUsingPUTOK{}
}

/*
RegisterUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type RegisterUsingPUTOK struct {
	Payload *models.K8SNamespace
}

// IsSuccess returns true when this register using p u t o k response has a 2xx status code
func (o *RegisterUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register using p u t o k response has a 3xx status code
func (o *RegisterUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register using p u t o k response has a 4xx status code
func (o *RegisterUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this register using p u t o k response has a 5xx status code
func (o *RegisterUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this register using p u t o k response a status code equal to that given
func (o *RegisterUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

func (o *RegisterUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/namespaces/{id}/register][%d] registerUsingPUTOK  %+v", 200, o.Payload)
}

func (o *RegisterUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/namespaces/{id}/register][%d] registerUsingPUTOK  %+v", 200, o.Payload)
}

func (o *RegisterUsingPUTOK) GetPayload() *models.K8SNamespace {
	return o.Payload
}

func (o *RegisterUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SNamespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterUsingPUTForbidden creates a RegisterUsingPUTForbidden with default headers values
func NewRegisterUsingPUTForbidden() *RegisterUsingPUTForbidden {
	return &RegisterUsingPUTForbidden{}
}

/*
RegisterUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type RegisterUsingPUTForbidden struct {
}

// IsSuccess returns true when this register using p u t forbidden response has a 2xx status code
func (o *RegisterUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this register using p u t forbidden response has a 3xx status code
func (o *RegisterUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register using p u t forbidden response has a 4xx status code
func (o *RegisterUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this register using p u t forbidden response has a 5xx status code
func (o *RegisterUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this register using p u t forbidden response a status code equal to that given
func (o *RegisterUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RegisterUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/namespaces/{id}/register][%d] registerUsingPUTForbidden ", 403)
}

func (o *RegisterUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/namespaces/{id}/register][%d] registerUsingPUTForbidden ", 403)
}

func (o *RegisterUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}