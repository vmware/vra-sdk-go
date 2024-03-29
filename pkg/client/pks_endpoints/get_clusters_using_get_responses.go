// Code generated by go-swagger; DO NOT EDIT.

package pks_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetClustersUsingGETReader is a Reader for the GetClustersUsingGET structure.
type GetClustersUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClustersUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClustersUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetClustersUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClustersUsingGETOK creates a GetClustersUsingGETOK with default headers values
func NewGetClustersUsingGETOK() *GetClustersUsingGETOK {
	return &GetClustersUsingGETOK{}
}

/*
GetClustersUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetClustersUsingGETOK struct {
	Payload []*models.PKSCluster
}

// IsSuccess returns true when this get clusters using g e t o k response has a 2xx status code
func (o *GetClustersUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get clusters using g e t o k response has a 3xx status code
func (o *GetClustersUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters using g e t o k response has a 4xx status code
func (o *GetClustersUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get clusters using g e t o k response has a 5xx status code
func (o *GetClustersUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters using g e t o k response a status code equal to that given
func (o *GetClustersUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetClustersUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/pks/endpoints/{id}/clusters][%d] getClustersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetClustersUsingGETOK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/pks/endpoints/{id}/clusters][%d] getClustersUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetClustersUsingGETOK) GetPayload() []*models.PKSCluster {
	return o.Payload
}

func (o *GetClustersUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClustersUsingGETForbidden creates a GetClustersUsingGETForbidden with default headers values
func NewGetClustersUsingGETForbidden() *GetClustersUsingGETForbidden {
	return &GetClustersUsingGETForbidden{}
}

/*
GetClustersUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type GetClustersUsingGETForbidden struct {
}

// IsSuccess returns true when this get clusters using g e t forbidden response has a 2xx status code
func (o *GetClustersUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clusters using g e t forbidden response has a 3xx status code
func (o *GetClustersUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clusters using g e t forbidden response has a 4xx status code
func (o *GetClustersUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clusters using g e t forbidden response has a 5xx status code
func (o *GetClustersUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get clusters using g e t forbidden response a status code equal to that given
func (o *GetClustersUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetClustersUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/pks/endpoints/{id}/clusters][%d] getClustersUsingGETForbidden ", 403)
}

func (o *GetClustersUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/pks/endpoints/{id}/clusters][%d] getClustersUsingGETForbidden ", 403)
}

func (o *GetClustersUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
