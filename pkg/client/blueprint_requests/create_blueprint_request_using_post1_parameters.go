// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

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

// NewCreateBlueprintRequestUsingPOST1Params creates a new CreateBlueprintRequestUsingPOST1Params object
// with the default values initialized.
func NewCreateBlueprintRequestUsingPOST1Params() *CreateBlueprintRequestUsingPOST1Params {
	var ()
	return &CreateBlueprintRequestUsingPOST1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlueprintRequestUsingPOST1ParamsWithTimeout creates a new CreateBlueprintRequestUsingPOST1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBlueprintRequestUsingPOST1ParamsWithTimeout(timeout time.Duration) *CreateBlueprintRequestUsingPOST1Params {
	var ()
	return &CreateBlueprintRequestUsingPOST1Params{

		timeout: timeout,
	}
}

// NewCreateBlueprintRequestUsingPOST1ParamsWithContext creates a new CreateBlueprintRequestUsingPOST1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBlueprintRequestUsingPOST1ParamsWithContext(ctx context.Context) *CreateBlueprintRequestUsingPOST1Params {
	var ()
	return &CreateBlueprintRequestUsingPOST1Params{

		Context: ctx,
	}
}

// NewCreateBlueprintRequestUsingPOST1ParamsWithHTTPClient creates a new CreateBlueprintRequestUsingPOST1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBlueprintRequestUsingPOST1ParamsWithHTTPClient(client *http.Client) *CreateBlueprintRequestUsingPOST1Params {
	var ()
	return &CreateBlueprintRequestUsingPOST1Params{
		HTTPClient: client,
	}
}

/*CreateBlueprintRequestUsingPOST1Params contains all the parameters to send to the API endpoint
for the create blueprint request using p o s t 1 operation typically these are written to a http.Request
*/
type CreateBlueprintRequestUsingPOST1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about

	*/
	APIVersion *string
	/*Request
	  request

	*/
	Request *models.BlueprintRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) WithTimeout(timeout time.Duration) *CreateBlueprintRequestUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) WithContext(ctx context.Context) *CreateBlueprintRequestUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) WithHTTPClient(client *http.Client) *CreateBlueprintRequestUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) WithAPIVersion(aPIVersion *string) *CreateBlueprintRequestUsingPOST1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithRequest adds the request to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) WithRequest(request *models.BlueprintRequest) *CreateBlueprintRequestUsingPOST1Params {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create blueprint request using p o s t 1 params
func (o *CreateBlueprintRequestUsingPOST1Params) SetRequest(request *models.BlueprintRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlueprintRequestUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string
		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {
			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}

	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
