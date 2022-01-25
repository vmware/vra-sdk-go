// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateResourceUsingPOST2Reader is a Reader for the CreateResourceUsingPOST2 structure.
type CreateResourceUsingPOST2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResourceUsingPOST2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateResourceUsingPOST2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateResourceUsingPOST2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateResourceUsingPOST2OK creates a CreateResourceUsingPOST2OK with default headers values
func NewCreateResourceUsingPOST2OK() *CreateResourceUsingPOST2OK {
	return &CreateResourceUsingPOST2OK{}
}

/* CreateResourceUsingPOST2OK describes a response with status code 200, with default header values.

OK
*/
type CreateResourceUsingPOST2OK struct {
	Payload *models.ResourceRequestResponse
}

func (o *CreateResourceUsingPOST2OK) Error() string {
	return fmt.Sprintf("[POST /deployment/api/resources][%d] createResourceUsingPOST2OK  %+v", 200, o.Payload)
}
func (o *CreateResourceUsingPOST2OK) GetPayload() *models.ResourceRequestResponse {
	return o.Payload
}

func (o *CreateResourceUsingPOST2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateResourceUsingPOST2Unauthorized creates a CreateResourceUsingPOST2Unauthorized with default headers values
func NewCreateResourceUsingPOST2Unauthorized() *CreateResourceUsingPOST2Unauthorized {
	return &CreateResourceUsingPOST2Unauthorized{}
}

/* CreateResourceUsingPOST2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateResourceUsingPOST2Unauthorized struct {
}

func (o *CreateResourceUsingPOST2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /deployment/api/resources][%d] createResourceUsingPOST2Unauthorized ", 401)
}

func (o *CreateResourceUsingPOST2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
