// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateUsingPOST1Params creates a new CreateUsingPOST1Params object
// with the default values initialized.
func NewCreateUsingPOST1Params() *CreateUsingPOST1Params {
	var ()
	return &CreateUsingPOST1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPOST1ParamsWithTimeout creates a new CreateUsingPOST1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateUsingPOST1Params {
	var ()
	return &CreateUsingPOST1Params{

		timeout: timeout,
	}
}

// NewCreateUsingPOST1ParamsWithContext creates a new CreateUsingPOST1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateUsingPOST1ParamsWithContext(ctx context.Context) *CreateUsingPOST1Params {
	var ()
	return &CreateUsingPOST1Params{

		Context: ctx,
	}
}

// NewCreateUsingPOST1ParamsWithHTTPClient creates a new CreateUsingPOST1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateUsingPOST1Params {
	var ()
	return &CreateUsingPOST1Params{
		HTTPClient: client,
	}
}

/*CreateUsingPOST1Params contains all the parameters to send to the API endpoint
for the create using p o s t 1 operation typically these are written to a http.Request
*/
type CreateUsingPOST1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Body
	  Endpoint specification

	*/
	Body models.EndpointSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithContext(ctx context.Context) *CreateUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithAPIVersion(aPIVersion string) *CreateUsingPOST1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) WithBody(body models.EndpointSpec) *CreateUsingPOST1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create using p o s t 1 params
func (o *CreateUsingPOST1Params) SetBody(body models.EndpointSpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
