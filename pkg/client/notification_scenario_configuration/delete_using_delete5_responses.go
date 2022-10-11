// Code generated by go-swagger; DO NOT EDIT.

package notification_scenario_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUsingDELETE5Reader is a Reader for the DeleteUsingDELETE5 structure.
type DeleteUsingDELETE5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETE5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUsingDELETE5NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETE5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsingDELETE5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUsingDELETE5NoContent creates a DeleteUsingDELETE5NoContent with default headers values
func NewDeleteUsingDELETE5NoContent() *DeleteUsingDELETE5NoContent {
	return &DeleteUsingDELETE5NoContent{}
}

/*
DeleteUsingDELETE5NoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteUsingDELETE5NoContent struct {
}

// IsSuccess returns true when this delete using d e l e t e5 no content response has a 2xx status code
func (o *DeleteUsingDELETE5NoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete using d e l e t e5 no content response has a 3xx status code
func (o *DeleteUsingDELETE5NoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete using d e l e t e5 no content response has a 4xx status code
func (o *DeleteUsingDELETE5NoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete using d e l e t e5 no content response has a 5xx status code
func (o *DeleteUsingDELETE5NoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete using d e l e t e5 no content response a status code equal to that given
func (o *DeleteUsingDELETE5NoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteUsingDELETE5NoContent) Error() string {
	return fmt.Sprintf("[DELETE /notification/api/scenario-configs/{scenarioId}][%d] deleteUsingDELETE5NoContent ", 204)
}

func (o *DeleteUsingDELETE5NoContent) String() string {
	return fmt.Sprintf("[DELETE /notification/api/scenario-configs/{scenarioId}][%d] deleteUsingDELETE5NoContent ", 204)
}

func (o *DeleteUsingDELETE5NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE5Unauthorized creates a DeleteUsingDELETE5Unauthorized with default headers values
func NewDeleteUsingDELETE5Unauthorized() *DeleteUsingDELETE5Unauthorized {
	return &DeleteUsingDELETE5Unauthorized{}
}

/*
DeleteUsingDELETE5Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteUsingDELETE5Unauthorized struct {
}

// IsSuccess returns true when this delete using d e l e t e5 unauthorized response has a 2xx status code
func (o *DeleteUsingDELETE5Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete using d e l e t e5 unauthorized response has a 3xx status code
func (o *DeleteUsingDELETE5Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete using d e l e t e5 unauthorized response has a 4xx status code
func (o *DeleteUsingDELETE5Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete using d e l e t e5 unauthorized response has a 5xx status code
func (o *DeleteUsingDELETE5Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete using d e l e t e5 unauthorized response a status code equal to that given
func (o *DeleteUsingDELETE5Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteUsingDELETE5Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /notification/api/scenario-configs/{scenarioId}][%d] deleteUsingDELETE5Unauthorized ", 401)
}

func (o *DeleteUsingDELETE5Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /notification/api/scenario-configs/{scenarioId}][%d] deleteUsingDELETE5Unauthorized ", 401)
}

func (o *DeleteUsingDELETE5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE5Forbidden creates a DeleteUsingDELETE5Forbidden with default headers values
func NewDeleteUsingDELETE5Forbidden() *DeleteUsingDELETE5Forbidden {
	return &DeleteUsingDELETE5Forbidden{}
}

/*
DeleteUsingDELETE5Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUsingDELETE5Forbidden struct {
}

// IsSuccess returns true when this delete using d e l e t e5 forbidden response has a 2xx status code
func (o *DeleteUsingDELETE5Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete using d e l e t e5 forbidden response has a 3xx status code
func (o *DeleteUsingDELETE5Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete using d e l e t e5 forbidden response has a 4xx status code
func (o *DeleteUsingDELETE5Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete using d e l e t e5 forbidden response has a 5xx status code
func (o *DeleteUsingDELETE5Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete using d e l e t e5 forbidden response a status code equal to that given
func (o *DeleteUsingDELETE5Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteUsingDELETE5Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /notification/api/scenario-configs/{scenarioId}][%d] deleteUsingDELETE5Forbidden ", 403)
}

func (o *DeleteUsingDELETE5Forbidden) String() string {
	return fmt.Sprintf("[DELETE /notification/api/scenario-configs/{scenarioId}][%d] deleteUsingDELETE5Forbidden ", 403)
}

func (o *DeleteUsingDELETE5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}