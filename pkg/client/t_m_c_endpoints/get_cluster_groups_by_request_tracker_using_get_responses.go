// Code generated by go-swagger; DO NOT EDIT.

package t_m_c_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetClusterGroupsByRequestTrackerUsingGETReader is a Reader for the GetClusterGroupsByRequestTrackerUsingGET structure.
type GetClusterGroupsByRequestTrackerUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterGroupsByRequestTrackerUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterGroupsByRequestTrackerUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetClusterGroupsByRequestTrackerUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterGroupsByRequestTrackerUsingGETOK creates a GetClusterGroupsByRequestTrackerUsingGETOK with default headers values
func NewGetClusterGroupsByRequestTrackerUsingGETOK() *GetClusterGroupsByRequestTrackerUsingGETOK {
	return &GetClusterGroupsByRequestTrackerUsingGETOK{}
}

/*
GetClusterGroupsByRequestTrackerUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetClusterGroupsByRequestTrackerUsingGETOK struct {
	Payload *models.PageOfClusterGroup
}

// IsSuccess returns true when this get cluster groups by request tracker using g e t o k response has a 2xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster groups by request tracker using g e t o k response has a 3xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster groups by request tracker using g e t o k response has a 4xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster groups by request tracker using g e t o k response has a 5xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster groups by request tracker using g e t o k response a status code equal to that given
func (o *GetClusterGroupsByRequestTrackerUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClusterGroupsByRequestTrackerUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/tmc/endpoints/clustergroups][%d] getClusterGroupsByRequestTrackerUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetClusterGroupsByRequestTrackerUsingGETOK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/tmc/endpoints/clustergroups][%d] getClusterGroupsByRequestTrackerUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetClusterGroupsByRequestTrackerUsingGETOK) GetPayload() *models.PageOfClusterGroup {
	return o.Payload
}

func (o *GetClusterGroupsByRequestTrackerUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterGroupsByRequestTrackerUsingGETForbidden creates a GetClusterGroupsByRequestTrackerUsingGETForbidden with default headers values
func NewGetClusterGroupsByRequestTrackerUsingGETForbidden() *GetClusterGroupsByRequestTrackerUsingGETForbidden {
	return &GetClusterGroupsByRequestTrackerUsingGETForbidden{}
}

/*
GetClusterGroupsByRequestTrackerUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type GetClusterGroupsByRequestTrackerUsingGETForbidden struct {
}

// IsSuccess returns true when this get cluster groups by request tracker using g e t forbidden response has a 2xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get cluster groups by request tracker using g e t forbidden response has a 3xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster groups by request tracker using g e t forbidden response has a 4xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get cluster groups by request tracker using g e t forbidden response has a 5xx status code
func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster groups by request tracker using g e t forbidden response a status code equal to that given
func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/tmc/endpoints/clustergroups][%d] getClusterGroupsByRequestTrackerUsingGETForbidden ", 403)
}

func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/tmc/endpoints/clustergroups][%d] getClusterGroupsByRequestTrackerUsingGETForbidden ", 403)
}

func (o *GetClusterGroupsByRequestTrackerUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}