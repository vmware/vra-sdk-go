// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAboutPageReader is a Reader for the GetAboutPage structure.
type GetAboutPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAboutPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAboutPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAboutPageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAboutPageOK creates a GetAboutPageOK with default headers values
func NewGetAboutPageOK() *GetAboutPageOK {
	return &GetAboutPageOK{}
}

/*
GetAboutPageOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAboutPageOK struct {
	Payload *models.IaaSAbout
}

// IsSuccess returns true when this get about page o k response has a 2xx status code
func (o *GetAboutPageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get about page o k response has a 3xx status code
func (o *GetAboutPageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get about page o k response has a 4xx status code
func (o *GetAboutPageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get about page o k response has a 5xx status code
func (o *GetAboutPageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get about page o k response a status code equal to that given
func (o *GetAboutPageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAboutPageOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/about][%d] getAboutPageOK  %+v", 200, o.Payload)
}

func (o *GetAboutPageOK) String() string {
	return fmt.Sprintf("[GET /iaas/api/about][%d] getAboutPageOK  %+v", 200, o.Payload)
}

func (o *GetAboutPageOK) GetPayload() *models.IaaSAbout {
	return o.Payload
}

func (o *GetAboutPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaaSAbout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAboutPageForbidden creates a GetAboutPageForbidden with default headers values
func NewGetAboutPageForbidden() *GetAboutPageForbidden {
	return &GetAboutPageForbidden{}
}

/*
GetAboutPageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAboutPageForbidden struct {
	Payload *models.ServiceErrorResponse
}

// IsSuccess returns true when this get about page forbidden response has a 2xx status code
func (o *GetAboutPageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get about page forbidden response has a 3xx status code
func (o *GetAboutPageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get about page forbidden response has a 4xx status code
func (o *GetAboutPageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get about page forbidden response has a 5xx status code
func (o *GetAboutPageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get about page forbidden response a status code equal to that given
func (o *GetAboutPageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAboutPageForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/about][%d] getAboutPageForbidden  %+v", 403, o.Payload)
}

func (o *GetAboutPageForbidden) String() string {
	return fmt.Sprintf("[GET /iaas/api/about][%d] getAboutPageForbidden  %+v", 403, o.Payload)
}

func (o *GetAboutPageForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetAboutPageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
