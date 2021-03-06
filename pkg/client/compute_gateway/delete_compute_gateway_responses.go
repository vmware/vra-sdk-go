// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteComputeGatewayReader is a Reader for the DeleteComputeGateway structure.
type DeleteComputeGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteComputeGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteComputeGatewayAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteComputeGatewayForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteComputeGatewayAccepted creates a DeleteComputeGatewayAccepted with default headers values
func NewDeleteComputeGatewayAccepted() *DeleteComputeGatewayAccepted {
	return &DeleteComputeGatewayAccepted{}
}

/*DeleteComputeGatewayAccepted handles this case with default header values.

successful operation
*/
type DeleteComputeGatewayAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteComputeGatewayAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/compute-gateways/{id}][%d] deleteComputeGatewayAccepted  %+v", 202, o.Payload)
}

func (o *DeleteComputeGatewayAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *DeleteComputeGatewayAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteComputeGatewayForbidden creates a DeleteComputeGatewayForbidden with default headers values
func NewDeleteComputeGatewayForbidden() *DeleteComputeGatewayForbidden {
	return &DeleteComputeGatewayForbidden{}
}

/*DeleteComputeGatewayForbidden handles this case with default header values.

Forbidden
*/
type DeleteComputeGatewayForbidden struct {
}

func (o *DeleteComputeGatewayForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/compute-gateways/{id}][%d] deleteComputeGatewayForbidden ", 403)
}

func (o *DeleteComputeGatewayForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
