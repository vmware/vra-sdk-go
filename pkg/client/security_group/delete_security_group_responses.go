// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteSecurityGroupReader is a Reader for the DeleteSecurityGroup structure.
type DeleteSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteSecurityGroupAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteSecurityGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteSecurityGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSecurityGroupAccepted creates a DeleteSecurityGroupAccepted with default headers values
func NewDeleteSecurityGroupAccepted() *DeleteSecurityGroupAccepted {
	return &DeleteSecurityGroupAccepted{}
}

/*DeleteSecurityGroupAccepted handles this case with default header values.

successful operation
*/
type DeleteSecurityGroupAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteSecurityGroupAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/security-groups/{id}][%d] deleteSecurityGroupAccepted  %+v", 202, o.Payload)
}

func (o *DeleteSecurityGroupAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *DeleteSecurityGroupAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSecurityGroupNoContent creates a DeleteSecurityGroupNoContent with default headers values
func NewDeleteSecurityGroupNoContent() *DeleteSecurityGroupNoContent {
	return &DeleteSecurityGroupNoContent{}
}

/*DeleteSecurityGroupNoContent handles this case with default header values.

No Content
*/
type DeleteSecurityGroupNoContent struct {
}

func (o *DeleteSecurityGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/security-groups/{id}][%d] deleteSecurityGroupNoContent ", 204)
}

func (o *DeleteSecurityGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSecurityGroupForbidden creates a DeleteSecurityGroupForbidden with default headers values
func NewDeleteSecurityGroupForbidden() *DeleteSecurityGroupForbidden {
	return &DeleteSecurityGroupForbidden{}
}

/*DeleteSecurityGroupForbidden handles this case with default header values.

Forbidden
*/
type DeleteSecurityGroupForbidden struct {
}

func (o *DeleteSecurityGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/security-groups/{id}][%d] deleteSecurityGroupForbidden ", 403)
}

func (o *DeleteSecurityGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
