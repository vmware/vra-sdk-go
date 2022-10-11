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

// SyncStatusUsingGETReader is a Reader for the SyncStatusUsingGET structure.
type SyncStatusUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncStatusUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncStatusUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSyncStatusUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSyncStatusUsingGETOK creates a SyncStatusUsingGETOK with default headers values
func NewSyncStatusUsingGETOK() *SyncStatusUsingGETOK {
	return &SyncStatusUsingGETOK{}
}

/*
SyncStatusUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type SyncStatusUsingGETOK struct {
	Payload *models.SupervisorNamespaceSyncResponseDTO
}

// IsSuccess returns true when this sync status using g e t o k response has a 2xx status code
func (o *SyncStatusUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync status using g e t o k response has a 3xx status code
func (o *SyncStatusUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync status using g e t o k response has a 4xx status code
func (o *SyncStatusUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync status using g e t o k response has a 5xx status code
func (o *SyncStatusUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync status using g e t o k response a status code equal to that given
func (o *SyncStatusUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *SyncStatusUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-namespaces/{namespaceSelfLinkId}/principals/{requestId}][%d] syncStatusUsingGETOK  %+v", 200, o.Payload)
}

func (o *SyncStatusUsingGETOK) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-namespaces/{namespaceSelfLinkId}/principals/{requestId}][%d] syncStatusUsingGETOK  %+v", 200, o.Payload)
}

func (o *SyncStatusUsingGETOK) GetPayload() *models.SupervisorNamespaceSyncResponseDTO {
	return o.Payload
}

func (o *SyncStatusUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupervisorNamespaceSyncResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncStatusUsingGETForbidden creates a SyncStatusUsingGETForbidden with default headers values
func NewSyncStatusUsingGETForbidden() *SyncStatusUsingGETForbidden {
	return &SyncStatusUsingGETForbidden{}
}

/*
SyncStatusUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type SyncStatusUsingGETForbidden struct {
}

// IsSuccess returns true when this sync status using g e t forbidden response has a 2xx status code
func (o *SyncStatusUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync status using g e t forbidden response has a 3xx status code
func (o *SyncStatusUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync status using g e t forbidden response has a 4xx status code
func (o *SyncStatusUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync status using g e t forbidden response has a 5xx status code
func (o *SyncStatusUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this sync status using g e t forbidden response a status code equal to that given
func (o *SyncStatusUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SyncStatusUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-namespaces/{namespaceSelfLinkId}/principals/{requestId}][%d] syncStatusUsingGETForbidden ", 403)
}

func (o *SyncStatusUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-namespaces/{namespaceSelfLinkId}/principals/{requestId}][%d] syncStatusUsingGETForbidden ", 403)
}

func (o *SyncStatusUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
