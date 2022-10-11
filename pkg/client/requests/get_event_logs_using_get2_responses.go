// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetEventLogsUsingGET2Reader is a Reader for the GetEventLogsUsingGET2 structure.
type GetEventLogsUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventLogsUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventLogsUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEventLogsUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEventLogsUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEventLogsUsingGET2OK creates a GetEventLogsUsingGET2OK with default headers values
func NewGetEventLogsUsingGET2OK() *GetEventLogsUsingGET2OK {
	return &GetEventLogsUsingGET2OK{}
}

/*
GetEventLogsUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetEventLogsUsingGET2OK struct {
	Payload *models.SliceOfEventLog
}

// IsSuccess returns true when this get event logs using g e t2 o k response has a 2xx status code
func (o *GetEventLogsUsingGET2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get event logs using g e t2 o k response has a 3xx status code
func (o *GetEventLogsUsingGET2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get event logs using g e t2 o k response has a 4xx status code
func (o *GetEventLogsUsingGET2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get event logs using g e t2 o k response has a 5xx status code
func (o *GetEventLogsUsingGET2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get event logs using g e t2 o k response a status code equal to that given
func (o *GetEventLogsUsingGET2OK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEventLogsUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}/events/{eventId}/logs][%d] getEventLogsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetEventLogsUsingGET2OK) String() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}/events/{eventId}/logs][%d] getEventLogsUsingGET2OK  %+v", 200, o.Payload)
}

func (o *GetEventLogsUsingGET2OK) GetPayload() *models.SliceOfEventLog {
	return o.Payload
}

func (o *GetEventLogsUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SliceOfEventLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventLogsUsingGET2Unauthorized creates a GetEventLogsUsingGET2Unauthorized with default headers values
func NewGetEventLogsUsingGET2Unauthorized() *GetEventLogsUsingGET2Unauthorized {
	return &GetEventLogsUsingGET2Unauthorized{}
}

/*
GetEventLogsUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEventLogsUsingGET2Unauthorized struct {
}

// IsSuccess returns true when this get event logs using g e t2 unauthorized response has a 2xx status code
func (o *GetEventLogsUsingGET2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get event logs using g e t2 unauthorized response has a 3xx status code
func (o *GetEventLogsUsingGET2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get event logs using g e t2 unauthorized response has a 4xx status code
func (o *GetEventLogsUsingGET2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get event logs using g e t2 unauthorized response has a 5xx status code
func (o *GetEventLogsUsingGET2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get event logs using g e t2 unauthorized response a status code equal to that given
func (o *GetEventLogsUsingGET2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetEventLogsUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}/events/{eventId}/logs][%d] getEventLogsUsingGET2Unauthorized ", 401)
}

func (o *GetEventLogsUsingGET2Unauthorized) String() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}/events/{eventId}/logs][%d] getEventLogsUsingGET2Unauthorized ", 401)
}

func (o *GetEventLogsUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEventLogsUsingGET2NotFound creates a GetEventLogsUsingGET2NotFound with default headers values
func NewGetEventLogsUsingGET2NotFound() *GetEventLogsUsingGET2NotFound {
	return &GetEventLogsUsingGET2NotFound{}
}

/*
GetEventLogsUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetEventLogsUsingGET2NotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this get event logs using g e t2 not found response has a 2xx status code
func (o *GetEventLogsUsingGET2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get event logs using g e t2 not found response has a 3xx status code
func (o *GetEventLogsUsingGET2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get event logs using g e t2 not found response has a 4xx status code
func (o *GetEventLogsUsingGET2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get event logs using g e t2 not found response has a 5xx status code
func (o *GetEventLogsUsingGET2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get event logs using g e t2 not found response a status code equal to that given
func (o *GetEventLogsUsingGET2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEventLogsUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}/events/{eventId}/logs][%d] getEventLogsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetEventLogsUsingGET2NotFound) String() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}/events/{eventId}/logs][%d] getEventLogsUsingGET2NotFound  %+v", 404, o.Payload)
}

func (o *GetEventLogsUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEventLogsUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}