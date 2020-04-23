// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetComputesReader is a Reader for the GetComputes structure.
type GetComputesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComputesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComputesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetComputesOK creates a GetComputesOK with default headers values
func NewGetComputesOK() *GetComputesOK {
	return &GetComputesOK{}
}

/*GetComputesOK handles this case with default header values.

successful operation
*/
type GetComputesOK struct {
	Payload *models.FabricComputeResult
}

func (o *GetComputesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones/{id}/computes][%d] getComputesOK  %+v", 200, o.Payload)
}

func (o *GetComputesOK) GetPayload() *models.FabricComputeResult {
	return o.Payload
}

func (o *GetComputesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricComputeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputesForbidden creates a GetComputesForbidden with default headers values
func NewGetComputesForbidden() *GetComputesForbidden {
	return &GetComputesForbidden{}
}

/*GetComputesForbidden handles this case with default header values.

Forbidden
*/
type GetComputesForbidden struct {
}

func (o *GetComputesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones/{id}/computes][%d] getComputesForbidden ", 403)
}

func (o *GetComputesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComputesNotFound creates a GetComputesNotFound with default headers values
func NewGetComputesNotFound() *GetComputesNotFound {
	return &GetComputesNotFound{}
}

/*GetComputesNotFound handles this case with default header values.

Not Found
*/
type GetComputesNotFound struct {
	Payload *models.Error
}

func (o *GetComputesNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones/{id}/computes][%d] getComputesNotFound  %+v", 404, o.Payload)
}

func (o *GetComputesNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetComputesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
