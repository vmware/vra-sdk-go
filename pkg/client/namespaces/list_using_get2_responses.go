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

// ListUsingGET2Reader is a Reader for the ListUsingGET2 structure.
type ListUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListUsingGET2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUsingGET2OK creates a ListUsingGET2OK with default headers values
func NewListUsingGET2OK() *ListUsingGET2OK {
	return &ListUsingGET2OK{}
}

/*
ListUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type ListUsingGET2OK struct {
	Payload *models.PageOfK8SNamespace
}

// IsSuccess returns true when this list using g e t2 o k response has a 2xx status code
func (o *ListUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list using g e t2 o k response has a 3xx status code
func (o *ListUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list using g e t2 o k response has a 4xx status code
func (o *ListUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list using g e t2 o k response has a 5xx status code
func (o *ListUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list using g e t2 o k response a status code equal to that given
func (o *ListUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces][%d] listUsingGET2OK  %+v", 200, o.Payload)
}

func (o *ListUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces][%d] listUsingGET2OK  %+v", 200, o.Payload)
}

func (o *ListUsingGET2OK) GetPayload() *models.PageOfK8SNamespace {
	return o.Payload
}

func (o *ListUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfK8SNamespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsingGET2Forbidden creates a ListUsingGET2Forbidden with default headers values
func NewListUsingGET2Forbidden() *ListUsingGET2Forbidden {
	return &ListUsingGET2Forbidden{}
}

/*
ListUsingGET2Forbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type ListUsingGET2Forbidden struct {
}

// IsSuccess returns true when this list using g e t2 forbidden response has a 2xx status code
func (o *ListUsingGET2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list using g e t2 forbidden response has a 3xx status code
func (o *ListUsingGET2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list using g e t2 forbidden response has a 4xx status code
func (o *ListUsingGET2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list using g e t2 forbidden response has a 5xx status code
func (o *ListUsingGET2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list using g e t2 forbidden response a status code equal to that given
func (o *ListUsingGET2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListUsingGET2Forbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces][%d] listUsingGET2Forbidden ", 403)
}

func (o *ListUsingGET2Forbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/namespaces][%d] listUsingGET2Forbidden ", 403)
}

func (o *ListUsingGET2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
