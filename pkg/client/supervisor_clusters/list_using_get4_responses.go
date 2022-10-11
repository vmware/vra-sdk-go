// Code generated by go-swagger; DO NOT EDIT.

package supervisor_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ListUsingGET4Reader is a Reader for the ListUsingGET4 structure.
type ListUsingGET4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUsingGET4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUsingGET4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListUsingGET4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUsingGET4OK creates a ListUsingGET4OK with default headers values
func NewListUsingGET4OK() *ListUsingGET4OK {
	return &ListUsingGET4OK{}
}

/*
ListUsingGET4OK describes a response with status code 200, with default header values.

OK
*/
type ListUsingGET4OK struct {
	Payload *models.PageOfSupervisorCluster
}

// IsSuccess returns true when this list using g e t4 o k response has a 2xx status code
func (o *ListUsingGET4OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list using g e t4 o k response has a 3xx status code
func (o *ListUsingGET4OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list using g e t4 o k response has a 4xx status code
func (o *ListUsingGET4OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list using g e t4 o k response has a 5xx status code
func (o *ListUsingGET4OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list using g e t4 o k response a status code equal to that given
func (o *ListUsingGET4OK) IsCode(code int) bool {
	return code == 200
}

func (o *ListUsingGET4OK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-clusters][%d] listUsingGET4OK  %+v", 200, o.Payload)
}

func (o *ListUsingGET4OK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-clusters][%d] listUsingGET4OK  %+v", 200, o.Payload)
}

func (o *ListUsingGET4OK) GetPayload() *models.PageOfSupervisorCluster {
	return o.Payload
}

func (o *ListUsingGET4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfSupervisorCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsingGET4Forbidden creates a ListUsingGET4Forbidden with default headers values
func NewListUsingGET4Forbidden() *ListUsingGET4Forbidden {
	return &ListUsingGET4Forbidden{}
}

/*
ListUsingGET4Forbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type ListUsingGET4Forbidden struct {
}

// IsSuccess returns true when this list using g e t4 forbidden response has a 2xx status code
func (o *ListUsingGET4Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list using g e t4 forbidden response has a 3xx status code
func (o *ListUsingGET4Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list using g e t4 forbidden response has a 4xx status code
func (o *ListUsingGET4Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list using g e t4 forbidden response has a 5xx status code
func (o *ListUsingGET4Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list using g e t4 forbidden response a status code equal to that given
func (o *ListUsingGET4Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListUsingGET4Forbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-clusters][%d] listUsingGET4Forbidden ", 403)
}

func (o *ListUsingGET4Forbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-clusters][%d] listUsingGET4Forbidden ", 403)
}

func (o *ListUsingGET4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}