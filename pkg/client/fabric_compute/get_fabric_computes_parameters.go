// Code generated by go-swagger; DO NOT EDIT.

package fabric_compute

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
)

// NewGetFabricComputesParams creates a new GetFabricComputesParams object
// with the default values initialized.
func NewGetFabricComputesParams() *GetFabricComputesParams {
	var ()
	return &GetFabricComputesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricComputesParamsWithTimeout creates a new GetFabricComputesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFabricComputesParamsWithTimeout(timeout time.Duration) *GetFabricComputesParams {
	var ()
	return &GetFabricComputesParams{

		timeout: timeout,
	}
}

// NewGetFabricComputesParamsWithContext creates a new GetFabricComputesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFabricComputesParamsWithContext(ctx context.Context) *GetFabricComputesParams {
	var ()
	return &GetFabricComputesParams{

		Context: ctx,
	}
}

// NewGetFabricComputesParamsWithHTTPClient creates a new GetFabricComputesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFabricComputesParamsWithHTTPClient(client *http.Client) *GetFabricComputesParams {
	var ()
	return &GetFabricComputesParams{
		HTTPClient: client,
	}
}

/*GetFabricComputesParams contains all the parameters to send to the API endpoint
for the get fabric computes operation typically these are written to a http.Request
*/
type GetFabricComputesParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fabric computes params
func (o *GetFabricComputesParams) WithTimeout(timeout time.Duration) *GetFabricComputesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric computes params
func (o *GetFabricComputesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric computes params
func (o *GetFabricComputesParams) WithContext(ctx context.Context) *GetFabricComputesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric computes params
func (o *GetFabricComputesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric computes params
func (o *GetFabricComputesParams) WithHTTPClient(client *http.Client) *GetFabricComputesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric computes params
func (o *GetFabricComputesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get fabric computes params
func (o *GetFabricComputesParams) WithAPIVersion(aPIVersion *string) *GetFabricComputesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric computes params
func (o *GetFabricComputesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricComputesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}