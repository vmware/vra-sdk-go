// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PatchPipelineGlobalByNameUsingPATCHReader is a Reader for the PatchPipelineGlobalByNameUsingPATCH structure.
type PatchPipelineGlobalByNameUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPipelineGlobalByNameUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchPipelineGlobalByNameUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchPipelineGlobalByNameUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchPipelineGlobalByNameUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchPipelineGlobalByNameUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchPipelineGlobalByNameUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchPipelineGlobalByNameUsingPATCHOK creates a PatchPipelineGlobalByNameUsingPATCHOK with default headers values
func NewPatchPipelineGlobalByNameUsingPATCHOK() *PatchPipelineGlobalByNameUsingPATCHOK {
	return &PatchPipelineGlobalByNameUsingPATCHOK{}
}

/*
PatchPipelineGlobalByNameUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with the updated Pipeline
*/
type PatchPipelineGlobalByNameUsingPATCHOK struct {
	Payload models.Pipeline
}

// IsSuccess returns true when this patch pipeline global by name using p a t c h o k response has a 2xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch pipeline global by name using p a t c h o k response has a 3xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch pipeline global by name using p a t c h o k response has a 4xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch pipeline global by name using p a t c h o k response has a 5xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch pipeline global by name using p a t c h o k response a status code equal to that given
func (o *PatchPipelineGlobalByNameUsingPATCHOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchPipelineGlobalByNameUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchPipelineGlobalByNameUsingPATCHOK) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHOK  %+v", 200, o.Payload)
}

func (o *PatchPipelineGlobalByNameUsingPATCHOK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *PatchPipelineGlobalByNameUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchPipelineGlobalByNameUsingPATCHUnauthorized creates a PatchPipelineGlobalByNameUsingPATCHUnauthorized with default headers values
func NewPatchPipelineGlobalByNameUsingPATCHUnauthorized() *PatchPipelineGlobalByNameUsingPATCHUnauthorized {
	return &PatchPipelineGlobalByNameUsingPATCHUnauthorized{}
}

/*
PatchPipelineGlobalByNameUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchPipelineGlobalByNameUsingPATCHUnauthorized struct {
}

// IsSuccess returns true when this patch pipeline global by name using p a t c h unauthorized response has a 2xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch pipeline global by name using p a t c h unauthorized response has a 3xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch pipeline global by name using p a t c h unauthorized response has a 4xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch pipeline global by name using p a t c h unauthorized response has a 5xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch pipeline global by name using p a t c h unauthorized response a status code equal to that given
func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHUnauthorized ", 401)
}

func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHUnauthorized ", 401)
}

func (o *PatchPipelineGlobalByNameUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchPipelineGlobalByNameUsingPATCHForbidden creates a PatchPipelineGlobalByNameUsingPATCHForbidden with default headers values
func NewPatchPipelineGlobalByNameUsingPATCHForbidden() *PatchPipelineGlobalByNameUsingPATCHForbidden {
	return &PatchPipelineGlobalByNameUsingPATCHForbidden{}
}

/*
PatchPipelineGlobalByNameUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchPipelineGlobalByNameUsingPATCHForbidden struct {
}

// IsSuccess returns true when this patch pipeline global by name using p a t c h forbidden response has a 2xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch pipeline global by name using p a t c h forbidden response has a 3xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch pipeline global by name using p a t c h forbidden response has a 4xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch pipeline global by name using p a t c h forbidden response has a 5xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch pipeline global by name using p a t c h forbidden response a status code equal to that given
func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHForbidden ", 403)
}

func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHForbidden ", 403)
}

func (o *PatchPipelineGlobalByNameUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchPipelineGlobalByNameUsingPATCHNotFound creates a PatchPipelineGlobalByNameUsingPATCHNotFound with default headers values
func NewPatchPipelineGlobalByNameUsingPATCHNotFound() *PatchPipelineGlobalByNameUsingPATCHNotFound {
	return &PatchPipelineGlobalByNameUsingPATCHNotFound{}
}

/*
PatchPipelineGlobalByNameUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchPipelineGlobalByNameUsingPATCHNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch pipeline global by name using p a t c h not found response has a 2xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch pipeline global by name using p a t c h not found response has a 3xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch pipeline global by name using p a t c h not found response has a 4xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch pipeline global by name using p a t c h not found response has a 5xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch pipeline global by name using p a t c h not found response a status code equal to that given
func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHNotFound  %+v", 404, o.Payload)
}

func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchPipelineGlobalByNameUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPipelineGlobalByNameUsingPATCHInternalServerError creates a PatchPipelineGlobalByNameUsingPATCHInternalServerError with default headers values
func NewPatchPipelineGlobalByNameUsingPATCHInternalServerError() *PatchPipelineGlobalByNameUsingPATCHInternalServerError {
	return &PatchPipelineGlobalByNameUsingPATCHInternalServerError{}
}

/*
PatchPipelineGlobalByNameUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchPipelineGlobalByNameUsingPATCHInternalServerError struct {
}

// IsSuccess returns true when this patch pipeline global by name using p a t c h internal server error response has a 2xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch pipeline global by name using p a t c h internal server error response has a 3xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch pipeline global by name using p a t c h internal server error response has a 4xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch pipeline global by name using p a t c h internal server error response has a 5xx status code
func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch pipeline global by name using p a t c h internal server error response a status code equal to that given
func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHInternalServerError ", 500)
}

func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}/global][%d] patchPipelineGlobalByNameUsingPATCHInternalServerError ", 500)
}

func (o *PatchPipelineGlobalByNameUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
