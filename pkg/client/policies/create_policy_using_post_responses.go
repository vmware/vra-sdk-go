// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// CreatePolicyUsingPOSTReader is a Reader for the CreatePolicyUsingPOST structure.
type CreatePolicyUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePolicyUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePolicyUsingPOSTOK creates a CreatePolicyUsingPOSTOK with default headers values
func NewCreatePolicyUsingPOSTOK() *CreatePolicyUsingPOSTOK {
	return &CreatePolicyUsingPOSTOK{}
}

/*CreatePolicyUsingPOSTOK handles this case with default header values.

OK
*/
type CreatePolicyUsingPOSTOK struct {
	Payload *models.Policy
}

func (o *CreatePolicyUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /policy/api/policies][%d] createPolicyUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreatePolicyUsingPOSTOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *CreatePolicyUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}