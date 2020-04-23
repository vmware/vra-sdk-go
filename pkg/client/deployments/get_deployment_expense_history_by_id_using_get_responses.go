// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDeploymentExpenseHistoryByIDUsingGETReader is a Reader for the GetDeploymentExpenseHistoryByIDUsingGET structure.
type GetDeploymentExpenseHistoryByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentExpenseHistoryByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentExpenseHistoryByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentExpenseHistoryByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentExpenseHistoryByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeploymentExpenseHistoryByIDUsingGETOK creates a GetDeploymentExpenseHistoryByIDUsingGETOK with default headers values
func NewGetDeploymentExpenseHistoryByIDUsingGETOK() *GetDeploymentExpenseHistoryByIDUsingGETOK {
	return &GetDeploymentExpenseHistoryByIDUsingGETOK{}
}

/*GetDeploymentExpenseHistoryByIDUsingGETOK handles this case with default header values.

OK
*/
type GetDeploymentExpenseHistoryByIDUsingGETOK struct {
	Payload *models.DeploymentExpenseHistory
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{depId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETOK) GetPayload() *models.DeploymentExpenseHistory {
	return o.Payload
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentExpenseHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentExpenseHistoryByIDUsingGETUnauthorized creates a GetDeploymentExpenseHistoryByIDUsingGETUnauthorized with default headers values
func NewGetDeploymentExpenseHistoryByIDUsingGETUnauthorized() *GetDeploymentExpenseHistoryByIDUsingGETUnauthorized {
	return &GetDeploymentExpenseHistoryByIDUsingGETUnauthorized{}
}

/*GetDeploymentExpenseHistoryByIDUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeploymentExpenseHistoryByIDUsingGETUnauthorized struct {
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{depId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGETUnauthorized ", 401)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentExpenseHistoryByIDUsingGETNotFound creates a GetDeploymentExpenseHistoryByIDUsingGETNotFound with default headers values
func NewGetDeploymentExpenseHistoryByIDUsingGETNotFound() *GetDeploymentExpenseHistoryByIDUsingGETNotFound {
	return &GetDeploymentExpenseHistoryByIDUsingGETNotFound{}
}

/*GetDeploymentExpenseHistoryByIDUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetDeploymentExpenseHistoryByIDUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{depId}/expense-history][%d] getDeploymentExpenseHistoryByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeploymentExpenseHistoryByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
