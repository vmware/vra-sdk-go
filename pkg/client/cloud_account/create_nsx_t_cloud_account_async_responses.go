// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateNsxTCloudAccountAsyncReader is a Reader for the CreateNsxTCloudAccountAsync structure.
type CreateNsxTCloudAccountAsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNsxTCloudAccountAsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewCreateNsxTCloudAccountAsyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateNsxTCloudAccountAsyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNsxTCloudAccountAsyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNsxTCloudAccountAsyncAccepted creates a CreateNsxTCloudAccountAsyncAccepted with default headers values
func NewCreateNsxTCloudAccountAsyncAccepted() *CreateNsxTCloudAccountAsyncAccepted {
	return &CreateNsxTCloudAccountAsyncAccepted{}
}

/*
CreateNsxTCloudAccountAsyncAccepted describes a response with status code 202, with default header values.

successful operation
*/
type CreateNsxTCloudAccountAsyncAccepted struct {
	Payload *models.RequestTracker
}

// IsSuccess returns true when this create nsx t cloud account async accepted response has a 2xx status code
func (o *CreateNsxTCloudAccountAsyncAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create nsx t cloud account async accepted response has a 3xx status code
func (o *CreateNsxTCloudAccountAsyncAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nsx t cloud account async accepted response has a 4xx status code
func (o *CreateNsxTCloudAccountAsyncAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create nsx t cloud account async accepted response has a 5xx status code
func (o *CreateNsxTCloudAccountAsyncAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create nsx t cloud account async accepted response a status code equal to that given
func (o *CreateNsxTCloudAccountAsyncAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *CreateNsxTCloudAccountAsyncAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *CreateNsxTCloudAccountAsyncAccepted) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountAsyncAccepted  %+v", 202, o.Payload)
}

func (o *CreateNsxTCloudAccountAsyncAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateNsxTCloudAccountAsyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNsxTCloudAccountAsyncBadRequest creates a CreateNsxTCloudAccountAsyncBadRequest with default headers values
func NewCreateNsxTCloudAccountAsyncBadRequest() *CreateNsxTCloudAccountAsyncBadRequest {
	return &CreateNsxTCloudAccountAsyncBadRequest{}
}

/*
CreateNsxTCloudAccountAsyncBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateNsxTCloudAccountAsyncBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this create nsx t cloud account async bad request response has a 2xx status code
func (o *CreateNsxTCloudAccountAsyncBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create nsx t cloud account async bad request response has a 3xx status code
func (o *CreateNsxTCloudAccountAsyncBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nsx t cloud account async bad request response has a 4xx status code
func (o *CreateNsxTCloudAccountAsyncBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create nsx t cloud account async bad request response has a 5xx status code
func (o *CreateNsxTCloudAccountAsyncBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create nsx t cloud account async bad request response a status code equal to that given
func (o *CreateNsxTCloudAccountAsyncBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateNsxTCloudAccountAsyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountAsyncBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNsxTCloudAccountAsyncBadRequest) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountAsyncBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNsxTCloudAccountAsyncBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateNsxTCloudAccountAsyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNsxTCloudAccountAsyncForbidden creates a CreateNsxTCloudAccountAsyncForbidden with default headers values
func NewCreateNsxTCloudAccountAsyncForbidden() *CreateNsxTCloudAccountAsyncForbidden {
	return &CreateNsxTCloudAccountAsyncForbidden{}
}

/*
CreateNsxTCloudAccountAsyncForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateNsxTCloudAccountAsyncForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this create nsx t cloud account async forbidden response has a 2xx status code
func (o *CreateNsxTCloudAccountAsyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create nsx t cloud account async forbidden response has a 3xx status code
func (o *CreateNsxTCloudAccountAsyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create nsx t cloud account async forbidden response has a 4xx status code
func (o *CreateNsxTCloudAccountAsyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create nsx t cloud account async forbidden response has a 5xx status code
func (o *CreateNsxTCloudAccountAsyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create nsx t cloud account async forbidden response a status code equal to that given
func (o *CreateNsxTCloudAccountAsyncForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateNsxTCloudAccountAsyncForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *CreateNsxTCloudAccountAsyncForbidden) String() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-nsx-t][%d] createNsxTCloudAccountAsyncForbidden  %+v", 403, o.Payload)
}

func (o *CreateNsxTCloudAccountAsyncForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateNsxTCloudAccountAsyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}