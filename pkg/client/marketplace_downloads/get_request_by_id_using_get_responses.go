// Code generated by go-swagger; DO NOT EDIT.

package marketplace_downloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetRequestByIDUsingGETReader is a Reader for the GetRequestByIDUsingGET structure.
type GetRequestByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRequestByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRequestByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRequestByIDUsingGETOK creates a GetRequestByIDUsingGETOK with default headers values
func NewGetRequestByIDUsingGETOK() *GetRequestByIDUsingGETOK {
	return &GetRequestByIDUsingGETOK{}
}

/*GetRequestByIDUsingGETOK handles this case with default header values.

Download request for id
*/
type GetRequestByIDUsingGETOK struct {
	Payload *models.MarketplaceContentDownloadRequest
}

func (o *GetRequestByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/download-requests/{id}][%d] getRequestByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetRequestByIDUsingGETOK) GetPayload() *models.MarketplaceContentDownloadRequest {
	return o.Payload
}

func (o *GetRequestByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MarketplaceContentDownloadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestByIDUsingGETNotFound creates a GetRequestByIDUsingGETNotFound with default headers values
func NewGetRequestByIDUsingGETNotFound() *GetRequestByIDUsingGETNotFound {
	return &GetRequestByIDUsingGETNotFound{}
}

/*GetRequestByIDUsingGETNotFound handles this case with default header values.

request not found
*/
type GetRequestByIDUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetRequestByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/download-requests/{id}][%d] getRequestByIdUsingGETNotFound  %+v", 404, o.Payload)
}

func (o *GetRequestByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRequestByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
