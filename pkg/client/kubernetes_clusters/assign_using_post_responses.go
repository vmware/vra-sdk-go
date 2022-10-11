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

// AssignUsingPOSTReader is a Reader for the AssignUsingPOST structure.
type AssignUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAssignUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignUsingPOSTOK creates a AssignUsingPOSTOK with default headers values
func NewAssignUsingPOSTOK() *AssignUsingPOSTOK {
	return &AssignUsingPOSTOK{}
}

/*
AssignUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type AssignUsingPOSTOK struct {
	Payload *models.K8SClusterAssignResponseDTO
}

// IsSuccess returns true when this assign using p o s t o k response has a 2xx status code
func (o *AssignUsingPOSTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign using p o s t o k response has a 3xx status code
func (o *AssignUsingPOSTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign using p o s t o k response has a 4xx status code
func (o *AssignUsingPOSTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign using p o s t o k response has a 5xx status code
func (o *AssignUsingPOSTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign using p o s t o k response a status code equal to that given
func (o *AssignUsingPOSTOK) IsCode(code int) bool {
	return code == 200
}

func (o *AssignUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /cmx/api/resources/k8s/clusters/assign][%d] assignUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *AssignUsingPOSTOK) String() string {
	return fmt.Sprintf("[POST /cmx/api/resources/k8s/clusters/assign][%d] assignUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *AssignUsingPOSTOK) GetPayload() *models.K8SClusterAssignResponseDTO {
	return o.Payload
}

func (o *AssignUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SClusterAssignResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignUsingPOSTForbidden creates a AssignUsingPOSTForbidden with default headers values
func NewAssignUsingPOSTForbidden() *AssignUsingPOSTForbidden {
	return &AssignUsingPOSTForbidden{}
}

/*
AssignUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type AssignUsingPOSTForbidden struct {
}

// IsSuccess returns true when this assign using p o s t forbidden response has a 2xx status code
func (o *AssignUsingPOSTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this assign using p o s t forbidden response has a 3xx status code
func (o *AssignUsingPOSTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign using p o s t forbidden response has a 4xx status code
func (o *AssignUsingPOSTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this assign using p o s t forbidden response has a 5xx status code
func (o *AssignUsingPOSTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this assign using p o s t forbidden response a status code equal to that given
func (o *AssignUsingPOSTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AssignUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /cmx/api/resources/k8s/clusters/assign][%d] assignUsingPOSTForbidden ", 403)
}

func (o *AssignUsingPOSTForbidden) String() string {
	return fmt.Sprintf("[POST /cmx/api/resources/k8s/clusters/assign][%d] assignUsingPOSTForbidden ", 403)
}

func (o *AssignUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}