// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// AssignStatusUsingGETReader is a Reader for the AssignStatusUsingGET structure.
type AssignStatusUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignStatusUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignStatusUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAssignStatusUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignStatusUsingGETOK creates a AssignStatusUsingGETOK with default headers values
func NewAssignStatusUsingGETOK() *AssignStatusUsingGETOK {
	return &AssignStatusUsingGETOK{}
}

/*
AssignStatusUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type AssignStatusUsingGETOK struct {
	Payload *models.K8SClusterAssignResponseDTO
}

// IsSuccess returns true when this assign status using g e t o k response has a 2xx status code
func (o *AssignStatusUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign status using g e t o k response has a 3xx status code
func (o *AssignStatusUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign status using g e t o k response has a 4xx status code
func (o *AssignStatusUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign status using g e t o k response has a 5xx status code
func (o *AssignStatusUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign status using g e t o k response a status code equal to that given
func (o *AssignStatusUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *AssignStatusUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/clusters/assign/{requestId}][%d] assignStatusUsingGETOK  %+v", 200, o.Payload)
}

func (o *AssignStatusUsingGETOK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/clusters/assign/{requestId}][%d] assignStatusUsingGETOK  %+v", 200, o.Payload)
}

func (o *AssignStatusUsingGETOK) GetPayload() *models.K8SClusterAssignResponseDTO {
	return o.Payload
}

func (o *AssignStatusUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SClusterAssignResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignStatusUsingGETForbidden creates a AssignStatusUsingGETForbidden with default headers values
func NewAssignStatusUsingGETForbidden() *AssignStatusUsingGETForbidden {
	return &AssignStatusUsingGETForbidden{}
}

/*
AssignStatusUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type AssignStatusUsingGETForbidden struct {
}

// IsSuccess returns true when this assign status using g e t forbidden response has a 2xx status code
func (o *AssignStatusUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign status using g e t forbidden response has a 3xx status code
func (o *AssignStatusUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign status using g e t forbidden response has a 4xx status code
func (o *AssignStatusUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign status using g e t forbidden response has a 5xx status code
func (o *AssignStatusUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this assign status using g e t forbidden response a status code equal to that given
func (o *AssignStatusUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AssignStatusUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/clusters/assign/{requestId}][%d] assignStatusUsingGETForbidden ", 403)
}

func (o *AssignStatusUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/k8s/clusters/assign/{requestId}][%d] assignStatusUsingGETForbidden ", 403)
}

func (o *AssignStatusUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
