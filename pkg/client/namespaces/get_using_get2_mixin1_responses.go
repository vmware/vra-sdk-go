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

// GetUsingGET2Mixin1Reader is a Reader for the GetUsingGET2Mixin1 structure.
type GetUsingGET2Mixin1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET2Mixin1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET2Mixin1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUsingGET2Mixin1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGET2Mixin1OK creates a GetUsingGET2Mixin1OK with default headers values
func NewGetUsingGET2Mixin1OK() *GetUsingGET2Mixin1OK {
	return &GetUsingGET2Mixin1OK{}
}

/*
GetUsingGET2Mixin1OK describes a response with status code 200, with default header values.

OK
*/
type GetUsingGET2Mixin1OK struct {
	Payload *models.K8SNamespace
}

// IsSuccess returns true when this get using g e t2 mixin1 o k response has a 2xx status code
func (o *GetUsingGET2Mixin1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get using g e t2 mixin1 o k response has a 3xx status code
func (o *GetUsingGET2Mixin1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get using g e t2 mixin1 o k response has a 4xx status code
func (o *GetUsingGET2Mixin1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get using g e t2 mixin1 o k response has a 5xx status code
func (o *GetUsingGET2Mixin1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get using g e t2 mixin1 o k response a status code equal to that given
func (o *GetUsingGET2Mixin1OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUsingGET2Mixin1OK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces/{id}][%d] getUsingGET2Mixin1OK  %+v", 200, o.Payload)
}

func (o *GetUsingGET2Mixin1OK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces/{id}][%d] getUsingGET2Mixin1OK  %+v", 200, o.Payload)
}

func (o *GetUsingGET2Mixin1OK) GetPayload() *models.K8SNamespace {
	return o.Payload
}

func (o *GetUsingGET2Mixin1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SNamespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGET2Mixin1Forbidden creates a GetUsingGET2Mixin1Forbidden with default headers values
func NewGetUsingGET2Mixin1Forbidden() *GetUsingGET2Mixin1Forbidden {
	return &GetUsingGET2Mixin1Forbidden{}
}

/*
GetUsingGET2Mixin1Forbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type GetUsingGET2Mixin1Forbidden struct {
}

// IsSuccess returns true when this get using g e t2 mixin1 forbidden response has a 2xx status code
func (o *GetUsingGET2Mixin1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get using g e t2 mixin1 forbidden response has a 3xx status code
func (o *GetUsingGET2Mixin1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get using g e t2 mixin1 forbidden response has a 4xx status code
func (o *GetUsingGET2Mixin1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get using g e t2 mixin1 forbidden response has a 5xx status code
func (o *GetUsingGET2Mixin1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get using g e t2 mixin1 forbidden response a status code equal to that given
func (o *GetUsingGET2Mixin1Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUsingGET2Mixin1Forbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces/{id}][%d] getUsingGET2Mixin1Forbidden ", 403)
}

func (o *GetUsingGET2Mixin1Forbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces/{id}][%d] getUsingGET2Mixin1Forbidden ", 403)
}

func (o *GetUsingGET2Mixin1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
