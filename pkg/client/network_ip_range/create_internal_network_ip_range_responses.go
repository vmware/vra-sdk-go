// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateInternalNetworkIPRangeReader is a Reader for the CreateInternalNetworkIPRange structure.
type CreateInternalNetworkIPRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInternalNetworkIPRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInternalNetworkIPRangeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateInternalNetworkIPRangeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInternalNetworkIPRangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInternalNetworkIPRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInternalNetworkIPRangeCreated creates a CreateInternalNetworkIPRangeCreated with default headers values
func NewCreateInternalNetworkIPRangeCreated() *CreateInternalNetworkIPRangeCreated {
	return &CreateInternalNetworkIPRangeCreated{}
}

/*CreateInternalNetworkIPRangeCreated handles this case with default header values.

successful operation
*/
type CreateInternalNetworkIPRangeCreated struct {
	Payload *models.RequestTracker
}

func (o *CreateInternalNetworkIPRangeCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/network-ip-ranges][%d] createInternalNetworkIpRangeCreated  %+v", 201, o.Payload)
}

func (o *CreateInternalNetworkIPRangeCreated) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateInternalNetworkIPRangeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternalNetworkIPRangeAccepted creates a CreateInternalNetworkIPRangeAccepted with default headers values
func NewCreateInternalNetworkIPRangeAccepted() *CreateInternalNetworkIPRangeAccepted {
	return &CreateInternalNetworkIPRangeAccepted{}
}

/*CreateInternalNetworkIPRangeAccepted handles this case with default header values.

successful operation
*/
type CreateInternalNetworkIPRangeAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateInternalNetworkIPRangeAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/network-ip-ranges][%d] createInternalNetworkIpRangeAccepted  %+v", 202, o.Payload)
}

func (o *CreateInternalNetworkIPRangeAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *CreateInternalNetworkIPRangeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternalNetworkIPRangeBadRequest creates a CreateInternalNetworkIPRangeBadRequest with default headers values
func NewCreateInternalNetworkIPRangeBadRequest() *CreateInternalNetworkIPRangeBadRequest {
	return &CreateInternalNetworkIPRangeBadRequest{}
}

/*CreateInternalNetworkIPRangeBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateInternalNetworkIPRangeBadRequest struct {
	Payload *models.Error
}

func (o *CreateInternalNetworkIPRangeBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/network-ip-ranges][%d] createInternalNetworkIpRangeBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInternalNetworkIPRangeBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateInternalNetworkIPRangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternalNetworkIPRangeForbidden creates a CreateInternalNetworkIPRangeForbidden with default headers values
func NewCreateInternalNetworkIPRangeForbidden() *CreateInternalNetworkIPRangeForbidden {
	return &CreateInternalNetworkIPRangeForbidden{}
}

/*CreateInternalNetworkIPRangeForbidden handles this case with default header values.

Forbidden
*/
type CreateInternalNetworkIPRangeForbidden struct {
}

func (o *CreateInternalNetworkIPRangeForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/network-ip-ranges][%d] createInternalNetworkIpRangeForbidden ", 403)
}

func (o *CreateInternalNetworkIPRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
