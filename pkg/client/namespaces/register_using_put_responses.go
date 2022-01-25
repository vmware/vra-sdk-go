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

// RegisterUsingPUTReader is a Reader for the RegisterUsingPUT structure.
type RegisterUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterUsingPUTOK creates a RegisterUsingPUTOK with default headers values
func NewRegisterUsingPUTOK() *RegisterUsingPUTOK {
	return &RegisterUsingPUTOK{}
}

/* RegisterUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type RegisterUsingPUTOK struct {
	Payload *models.K8SNamespace
}

func (o *RegisterUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/k8s/namespaces/{id}/register][%d] registerUsingPUTOK  %+v", 200, o.Payload)
}
func (o *RegisterUsingPUTOK) GetPayload() *models.K8SNamespace {
	return o.Payload
}

func (o *RegisterUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SNamespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
