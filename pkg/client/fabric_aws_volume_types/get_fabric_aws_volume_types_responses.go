// Code generated by go-swagger; DO NOT EDIT.

package fabric_aws_volume_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricAwsVolumeTypesReader is a Reader for the GetFabricAwsVolumeTypes structure.
type GetFabricAwsVolumeTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricAwsVolumeTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricAwsVolumeTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricAwsVolumeTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFabricAwsVolumeTypesOK creates a GetFabricAwsVolumeTypesOK with default headers values
func NewGetFabricAwsVolumeTypesOK() *GetFabricAwsVolumeTypesOK {
	return &GetFabricAwsVolumeTypesOK{}
}

/*
GetFabricAwsVolumeTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFabricAwsVolumeTypesOK struct {
	Payload *models.VolumeTypeList
}

// IsSuccess returns true when this get fabric aws volume types o k response has a 2xx status code
func (o *GetFabricAwsVolumeTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get fabric aws volume types o k response has a 3xx status code
func (o *GetFabricAwsVolumeTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric aws volume types o k response has a 4xx status code
func (o *GetFabricAwsVolumeTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get fabric aws volume types o k response has a 5xx status code
func (o *GetFabricAwsVolumeTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric aws volume types o k response a status code equal to that given
func (o *GetFabricAwsVolumeTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFabricAwsVolumeTypesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-aws-volume-types][%d] getFabricAwsVolumeTypesOK  %+v", 200, o.Payload)
}

func (o *GetFabricAwsVolumeTypesOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-aws-volume-types][%d] getFabricAwsVolumeTypesOK  %+v", 200, o.Payload)
}

func (o *GetFabricAwsVolumeTypesOK) GetPayload() *models.VolumeTypeList {
	return o.Payload
}

func (o *GetFabricAwsVolumeTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeTypeList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricAwsVolumeTypesForbidden creates a GetFabricAwsVolumeTypesForbidden with default headers values
func NewGetFabricAwsVolumeTypesForbidden() *GetFabricAwsVolumeTypesForbidden {
	return &GetFabricAwsVolumeTypesForbidden{}
}

/*
GetFabricAwsVolumeTypesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFabricAwsVolumeTypesForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get fabric aws volume types forbidden response has a 2xx status code
func (o *GetFabricAwsVolumeTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get fabric aws volume types forbidden response has a 3xx status code
func (o *GetFabricAwsVolumeTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get fabric aws volume types forbidden response has a 4xx status code
func (o *GetFabricAwsVolumeTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get fabric aws volume types forbidden response has a 5xx status code
func (o *GetFabricAwsVolumeTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get fabric aws volume types forbidden response a status code equal to that given
func (o *GetFabricAwsVolumeTypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFabricAwsVolumeTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-aws-volume-types][%d] getFabricAwsVolumeTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetFabricAwsVolumeTypesForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-aws-volume-types][%d] getFabricAwsVolumeTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetFabricAwsVolumeTypesForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetFabricAwsVolumeTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
