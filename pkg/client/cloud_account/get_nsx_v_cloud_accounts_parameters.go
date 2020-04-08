// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewGetNsxVCloudAccountsParams creates a new GetNsxVCloudAccountsParams object
// with the default values initialized.
func NewGetNsxVCloudAccountsParams() *GetNsxVCloudAccountsParams {
	var ()
	return &GetNsxVCloudAccountsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNsxVCloudAccountsParamsWithTimeout creates a new GetNsxVCloudAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNsxVCloudAccountsParamsWithTimeout(timeout time.Duration) *GetNsxVCloudAccountsParams {
	var ()
	return &GetNsxVCloudAccountsParams{

		timeout: timeout,
	}
}

// NewGetNsxVCloudAccountsParamsWithContext creates a new GetNsxVCloudAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNsxVCloudAccountsParamsWithContext(ctx context.Context) *GetNsxVCloudAccountsParams {
	var ()
	return &GetNsxVCloudAccountsParams{

		Context: ctx,
	}
}

// NewGetNsxVCloudAccountsParamsWithHTTPClient creates a new GetNsxVCloudAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNsxVCloudAccountsParamsWithHTTPClient(client *http.Client) *GetNsxVCloudAccountsParams {
	var ()
	return &GetNsxVCloudAccountsParams{
		HTTPClient: client,
	}
}

/*GetNsxVCloudAccountsParams contains all the parameters to send to the API endpoint
for the get nsx v cloud accounts operation typically these are written to a http.Request
*/
type GetNsxVCloudAccountsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) WithTimeout(timeout time.Duration) *GetNsxVCloudAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) WithContext(ctx context.Context) *GetNsxVCloudAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) WithHTTPClient(client *http.Client) *GetNsxVCloudAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) WithAPIVersion(aPIVersion *string) *GetNsxVCloudAccountsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get nsx v cloud accounts params
func (o *GetNsxVCloudAccountsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetNsxVCloudAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
