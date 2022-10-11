// Code generated by go-swagger; DO NOT EDIT.

package vcf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDomainUsingGETReader is a Reader for the GetDomainUsingGET structure.
type GetDomainUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDomainUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDomainUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDomainUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDomainUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDomainUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDomainUsingGETOK creates a GetDomainUsingGETOK with default headers values
func NewGetDomainUsingGETOK() *GetDomainUsingGETOK {
	return &GetDomainUsingGETOK{}
}

/*
GetDomainUsingGETOK describes a response with status code 200, with default header values.

Get domain detail
*/
type GetDomainUsingGETOK struct {
	Payload *models.VcfDomain
}

// IsSuccess returns true when this get domain using g e t o k response has a 2xx status code
func (o *GetDomainUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get domain using g e t o k response has a 3xx status code
func (o *GetDomainUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain using g e t o k response has a 4xx status code
func (o *GetDomainUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get domain using g e t o k response has a 5xx status code
func (o *GetDomainUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain using g e t o k response a status code equal to that given
func (o *GetDomainUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDomainUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDomainUsingGETOK) String() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDomainUsingGETOK) GetPayload() *models.VcfDomain {
	return o.Payload
}

func (o *GetDomainUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcfDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainUsingGETBadRequest creates a GetDomainUsingGETBadRequest with default headers values
func NewGetDomainUsingGETBadRequest() *GetDomainUsingGETBadRequest {
	return &GetDomainUsingGETBadRequest{}
}

/*
GetDomainUsingGETBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetDomainUsingGETBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get domain using g e t bad request response has a 2xx status code
func (o *GetDomainUsingGETBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain using g e t bad request response has a 3xx status code
func (o *GetDomainUsingGETBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain using g e t bad request response has a 4xx status code
func (o *GetDomainUsingGETBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain using g e t bad request response has a 5xx status code
func (o *GetDomainUsingGETBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain using g e t bad request response a status code equal to that given
func (o *GetDomainUsingGETBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetDomainUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetDomainUsingGETBadRequest) String() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETBadRequest  %+v", 400, o.Payload)
}

func (o *GetDomainUsingGETBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDomainUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDomainUsingGETUnauthorized creates a GetDomainUsingGETUnauthorized with default headers values
func NewGetDomainUsingGETUnauthorized() *GetDomainUsingGETUnauthorized {
	return &GetDomainUsingGETUnauthorized{}
}

/*
GetDomainUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDomainUsingGETUnauthorized struct {
}

// IsSuccess returns true when this get domain using g e t unauthorized response has a 2xx status code
func (o *GetDomainUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain using g e t unauthorized response has a 3xx status code
func (o *GetDomainUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain using g e t unauthorized response has a 4xx status code
func (o *GetDomainUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain using g e t unauthorized response has a 5xx status code
func (o *GetDomainUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain using g e t unauthorized response a status code equal to that given
func (o *GetDomainUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetDomainUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETUnauthorized ", 401)
}

func (o *GetDomainUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETUnauthorized ", 401)
}

func (o *GetDomainUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDomainUsingGETForbidden creates a GetDomainUsingGETForbidden with default headers values
func NewGetDomainUsingGETForbidden() *GetDomainUsingGETForbidden {
	return &GetDomainUsingGETForbidden{}
}

/*
GetDomainUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDomainUsingGETForbidden struct {
}

// IsSuccess returns true when this get domain using g e t forbidden response has a 2xx status code
func (o *GetDomainUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get domain using g e t forbidden response has a 3xx status code
func (o *GetDomainUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get domain using g e t forbidden response has a 4xx status code
func (o *GetDomainUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get domain using g e t forbidden response has a 5xx status code
func (o *GetDomainUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get domain using g e t forbidden response a status code equal to that given
func (o *GetDomainUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetDomainUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETForbidden ", 403)
}

func (o *GetDomainUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /content/api/vcf/{integrationId}/domain/{domainId}][%d] getDomainUsingGETForbidden ", 403)
}

func (o *GetDomainUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
