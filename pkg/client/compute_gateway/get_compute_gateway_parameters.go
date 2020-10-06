// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

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

// NewGetComputeGatewayParams creates a new GetComputeGatewayParams object
// with the default values initialized.
func NewGetComputeGatewayParams() *GetComputeGatewayParams {
	var ()
	return &GetComputeGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetComputeGatewayParamsWithTimeout creates a new GetComputeGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetComputeGatewayParamsWithTimeout(timeout time.Duration) *GetComputeGatewayParams {
	var ()
	return &GetComputeGatewayParams{

		timeout: timeout,
	}
}

// NewGetComputeGatewayParamsWithContext creates a new GetComputeGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetComputeGatewayParamsWithContext(ctx context.Context) *GetComputeGatewayParams {
	var ()
	return &GetComputeGatewayParams{

		Context: ctx,
	}
}

// NewGetComputeGatewayParamsWithHTTPClient creates a new GetComputeGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetComputeGatewayParamsWithHTTPClient(client *http.Client) *GetComputeGatewayParams {
	var ()
	return &GetComputeGatewayParams{
		HTTPClient: client,
	}
}

/*GetComputeGatewayParams contains all the parameters to send to the API endpoint
for the get compute gateway operation typically these are written to a http.Request
*/
type GetComputeGatewayParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get compute gateway params
func (o *GetComputeGatewayParams) WithTimeout(timeout time.Duration) *GetComputeGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compute gateway params
func (o *GetComputeGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compute gateway params
func (o *GetComputeGatewayParams) WithContext(ctx context.Context) *GetComputeGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compute gateway params
func (o *GetComputeGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compute gateway params
func (o *GetComputeGatewayParams) WithHTTPClient(client *http.Client) *GetComputeGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compute gateway params
func (o *GetComputeGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get compute gateway params
func (o *GetComputeGatewayParams) WithAPIVersion(aPIVersion *string) *GetComputeGatewayParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get compute gateway params
func (o *GetComputeGatewayParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetComputeGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
