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

// OnboardUsingPOSTReader is a Reader for the OnboardUsingPOST structure.
type OnboardUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OnboardUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOnboardUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOnboardUsingPOSTOK creates a OnboardUsingPOSTOK with default headers values
func NewOnboardUsingPOSTOK() *OnboardUsingPOSTOK {
	return &OnboardUsingPOSTOK{}
}

/* OnboardUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type OnboardUsingPOSTOK struct {
	Payload *models.K8SCluster
}

func (o *OnboardUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /cmx/api/resources/k8s/clusters][%d] onboardUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *OnboardUsingPOSTOK) GetPayload() *models.K8SCluster {
	return o.Payload
}

func (o *OnboardUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
