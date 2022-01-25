// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteIntegrationReader is a Reader for the DeleteIntegration structure.
type DeleteIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteIntegrationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteIntegrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIntegrationAccepted creates a DeleteIntegrationAccepted with default headers values
func NewDeleteIntegrationAccepted() *DeleteIntegrationAccepted {
	return &DeleteIntegrationAccepted{}
}

/* DeleteIntegrationAccepted describes a response with status code 202, with default header values.

successful operation
*/
type DeleteIntegrationAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteIntegrationAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/integrations/{id}][%d] deleteIntegrationAccepted  %+v", 202, o.Payload)
}
func (o *DeleteIntegrationAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *DeleteIntegrationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationForbidden creates a DeleteIntegrationForbidden with default headers values
func NewDeleteIntegrationForbidden() *DeleteIntegrationForbidden {
	return &DeleteIntegrationForbidden{}
}

/* DeleteIntegrationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteIntegrationForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *DeleteIntegrationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/integrations/{id}][%d] deleteIntegrationForbidden  %+v", 403, o.Payload)
}
func (o *DeleteIntegrationForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteIntegrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
