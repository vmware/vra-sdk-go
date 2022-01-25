// Code generated by go-swagger; DO NOT EDIT.

package supervisor_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// RegisterUsingPUT1Reader is a Reader for the RegisterUsingPUT1 structure.
type RegisterUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterUsingPUT1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterUsingPUT1OK creates a RegisterUsingPUT1OK with default headers values
func NewRegisterUsingPUT1OK() *RegisterUsingPUT1OK {
	return &RegisterUsingPUT1OK{}
}

/* RegisterUsingPUT1OK describes a response with status code 200, with default header values.

OK
*/
type RegisterUsingPUT1OK struct {
	Payload *models.SupervisorCluster
}

func (o *RegisterUsingPUT1OK) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/supervisor-clusters/{clusterSelfLinkId}/register][%d] registerUsingPUT1OK  %+v", 200, o.Payload)
}
func (o *RegisterUsingPUT1OK) GetPayload() *models.SupervisorCluster {
	return o.Payload
}

func (o *RegisterUsingPUT1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupervisorCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
