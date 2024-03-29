// Code generated by go-swagger; DO NOT EDIT.

package supervisor_namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// SetNamespaceQuotasUsingPUTReader is a Reader for the SetNamespaceQuotasUsingPUT structure.
type SetNamespaceQuotasUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetNamespaceQuotasUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetNamespaceQuotasUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSetNamespaceQuotasUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetNamespaceQuotasUsingPUTOK creates a SetNamespaceQuotasUsingPUTOK with default headers values
func NewSetNamespaceQuotasUsingPUTOK() *SetNamespaceQuotasUsingPUTOK {
	return &SetNamespaceQuotasUsingPUTOK{}
}

/*
SetNamespaceQuotasUsingPUTOK describes a response with status code 200, with default header values.

OK
*/
type SetNamespaceQuotasUsingPUTOK struct {
	Payload []*models.SupervisorNamespaceQuota
}

// IsSuccess returns true when this set namespace quotas using p u t o k response has a 2xx status code
func (o *SetNamespaceQuotasUsingPUTOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set namespace quotas using p u t o k response has a 3xx status code
func (o *SetNamespaceQuotasUsingPUTOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set namespace quotas using p u t o k response has a 4xx status code
func (o *SetNamespaceQuotasUsingPUTOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set namespace quotas using p u t o k response has a 5xx status code
func (o *SetNamespaceQuotasUsingPUTOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set namespace quotas using p u t o k response a status code equal to that given
func (o *SetNamespaceQuotasUsingPUTOK) IsCode(code int) bool {
	return code == 200
}

func (o *SetNamespaceQuotasUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/supervisor-namespaces/{selfLinkId}/resource-quotas][%d] setNamespaceQuotasUsingPUTOK  %+v", 200, o.Payload)
}

func (o *SetNamespaceQuotasUsingPUTOK) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/supervisor-namespaces/{selfLinkId}/resource-quotas][%d] setNamespaceQuotasUsingPUTOK  %+v", 200, o.Payload)
}

func (o *SetNamespaceQuotasUsingPUTOK) GetPayload() []*models.SupervisorNamespaceQuota {
	return o.Payload
}

func (o *SetNamespaceQuotasUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetNamespaceQuotasUsingPUTForbidden creates a SetNamespaceQuotasUsingPUTForbidden with default headers values
func NewSetNamespaceQuotasUsingPUTForbidden() *SetNamespaceQuotasUsingPUTForbidden {
	return &SetNamespaceQuotasUsingPUTForbidden{}
}

/*
SetNamespaceQuotasUsingPUTForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type SetNamespaceQuotasUsingPUTForbidden struct {
}

// IsSuccess returns true when this set namespace quotas using p u t forbidden response has a 2xx status code
func (o *SetNamespaceQuotasUsingPUTForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set namespace quotas using p u t forbidden response has a 3xx status code
func (o *SetNamespaceQuotasUsingPUTForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set namespace quotas using p u t forbidden response has a 4xx status code
func (o *SetNamespaceQuotasUsingPUTForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set namespace quotas using p u t forbidden response has a 5xx status code
func (o *SetNamespaceQuotasUsingPUTForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set namespace quotas using p u t forbidden response a status code equal to that given
func (o *SetNamespaceQuotasUsingPUTForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SetNamespaceQuotasUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/supervisor-namespaces/{selfLinkId}/resource-quotas][%d] setNamespaceQuotasUsingPUTForbidden ", 403)
}

func (o *SetNamespaceQuotasUsingPUTForbidden) String() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/supervisor-namespaces/{selfLinkId}/resource-quotas][%d] setNamespaceQuotasUsingPUTForbidden ", 403)
}

func (o *SetNamespaceQuotasUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
