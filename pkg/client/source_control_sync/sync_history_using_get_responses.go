// Code generated by go-swagger; DO NOT EDIT.

package source_control_sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// SyncHistoryUsingGETReader is a Reader for the SyncHistoryUsingGET structure.
type SyncHistoryUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncHistoryUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncHistoryUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSyncHistoryUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSyncHistoryUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyncHistoryUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSyncHistoryUsingGETOK creates a SyncHistoryUsingGETOK with default headers values
func NewSyncHistoryUsingGETOK() *SyncHistoryUsingGETOK {
	return &SyncHistoryUsingGETOK{}
}

/*
SyncHistoryUsingGETOK describes a response with status code 200, with default header values.

Sync history
*/
type SyncHistoryUsingGETOK struct {
	Payload *models.SourceControlSyncHistory
}

// IsSuccess returns true when this sync history using g e t o k response has a 2xx status code
func (o *SyncHistoryUsingGETOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync history using g e t o k response has a 3xx status code
func (o *SyncHistoryUsingGETOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync history using g e t o k response has a 4xx status code
func (o *SyncHistoryUsingGETOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync history using g e t o k response has a 5xx status code
func (o *SyncHistoryUsingGETOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync history using g e t o k response a status code equal to that given
func (o *SyncHistoryUsingGETOK) IsCode(code int) bool {
	return code == 200
}

func (o *SyncHistoryUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETOK  %+v", 200, o.Payload)
}

func (o *SyncHistoryUsingGETOK) String() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETOK  %+v", 200, o.Payload)
}

func (o *SyncHistoryUsingGETOK) GetPayload() *models.SourceControlSyncHistory {
	return o.Payload
}

func (o *SyncHistoryUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SourceControlSyncHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncHistoryUsingGETUnauthorized creates a SyncHistoryUsingGETUnauthorized with default headers values
func NewSyncHistoryUsingGETUnauthorized() *SyncHistoryUsingGETUnauthorized {
	return &SyncHistoryUsingGETUnauthorized{}
}

/*
SyncHistoryUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SyncHistoryUsingGETUnauthorized struct {
}

// IsSuccess returns true when this sync history using g e t unauthorized response has a 2xx status code
func (o *SyncHistoryUsingGETUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync history using g e t unauthorized response has a 3xx status code
func (o *SyncHistoryUsingGETUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync history using g e t unauthorized response has a 4xx status code
func (o *SyncHistoryUsingGETUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync history using g e t unauthorized response has a 5xx status code
func (o *SyncHistoryUsingGETUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this sync history using g e t unauthorized response a status code equal to that given
func (o *SyncHistoryUsingGETUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SyncHistoryUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETUnauthorized ", 401)
}

func (o *SyncHistoryUsingGETUnauthorized) String() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETUnauthorized ", 401)
}

func (o *SyncHistoryUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncHistoryUsingGETForbidden creates a SyncHistoryUsingGETForbidden with default headers values
func NewSyncHistoryUsingGETForbidden() *SyncHistoryUsingGETForbidden {
	return &SyncHistoryUsingGETForbidden{}
}

/*
SyncHistoryUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SyncHistoryUsingGETForbidden struct {
}

// IsSuccess returns true when this sync history using g e t forbidden response has a 2xx status code
func (o *SyncHistoryUsingGETForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync history using g e t forbidden response has a 3xx status code
func (o *SyncHistoryUsingGETForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync history using g e t forbidden response has a 4xx status code
func (o *SyncHistoryUsingGETForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync history using g e t forbidden response has a 5xx status code
func (o *SyncHistoryUsingGETForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this sync history using g e t forbidden response a status code equal to that given
func (o *SyncHistoryUsingGETForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SyncHistoryUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETForbidden ", 403)
}

func (o *SyncHistoryUsingGETForbidden) String() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETForbidden ", 403)
}

func (o *SyncHistoryUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncHistoryUsingGETNotFound creates a SyncHistoryUsingGETNotFound with default headers values
func NewSyncHistoryUsingGETNotFound() *SyncHistoryUsingGETNotFound {
	return &SyncHistoryUsingGETNotFound{}
}

/*
SyncHistoryUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SyncHistoryUsingGETNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this sync history using g e t not found response has a 2xx status code
func (o *SyncHistoryUsingGETNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this sync history using g e t not found response has a 3xx status code
func (o *SyncHistoryUsingGETNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync history using g e t not found response has a 4xx status code
func (o *SyncHistoryUsingGETNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this sync history using g e t not found response has a 5xx status code
func (o *SyncHistoryUsingGETNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this sync history using g e t not found response a status code equal to that given
func (o *SyncHistoryUsingGETNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SyncHistoryUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *SyncHistoryUsingGETNotFound) String() string {
	return fmt.Sprintf("[GET /content/api/sourcecontrol/sync-history][%d] syncHistoryUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *SyncHistoryUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SyncHistoryUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
