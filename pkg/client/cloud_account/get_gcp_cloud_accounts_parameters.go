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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetGcpCloudAccountsParams creates a new GetGcpCloudAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGcpCloudAccountsParams() *GetGcpCloudAccountsParams {
	return &GetGcpCloudAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGcpCloudAccountsParamsWithTimeout creates a new GetGcpCloudAccountsParams object
// with the ability to set a timeout on a request.
func NewGetGcpCloudAccountsParamsWithTimeout(timeout time.Duration) *GetGcpCloudAccountsParams {
	return &GetGcpCloudAccountsParams{
		timeout: timeout,
	}
}

// NewGetGcpCloudAccountsParamsWithContext creates a new GetGcpCloudAccountsParams object
// with the ability to set a context for a request.
func NewGetGcpCloudAccountsParamsWithContext(ctx context.Context) *GetGcpCloudAccountsParams {
	return &GetGcpCloudAccountsParams{
		Context: ctx,
	}
}

// NewGetGcpCloudAccountsParamsWithHTTPClient creates a new GetGcpCloudAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGcpCloudAccountsParamsWithHTTPClient(client *http.Client) *GetGcpCloudAccountsParams {
	return &GetGcpCloudAccountsParams{
		HTTPClient: client,
	}
}

/*
GetGcpCloudAccountsParams contains all the parameters to send to the API endpoint

	for the get gcp cloud accounts operation.

	Typically these are written to a http.Request.
*/
type GetGcpCloudAccountsParams struct {

	/* DollarSkip.

	   Number of records you want to skip.
	*/
	DollarSkip *int64

	/* DollarTop.

	   Number of records you want to get.
	*/
	DollarTop *int64

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gcp cloud accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGcpCloudAccountsParams) WithDefaults() *GetGcpCloudAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gcp cloud accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGcpCloudAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) WithTimeout(timeout time.Duration) *GetGcpCloudAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) WithContext(ctx context.Context) *GetGcpCloudAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) WithHTTPClient(client *http.Client) *GetGcpCloudAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSkip adds the dollarSkip to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) WithDollarSkip(dollarSkip *int64) *GetGcpCloudAccountsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) SetDollarSkip(dollarSkip *int64) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) WithDollarTop(dollarTop *int64) *GetGcpCloudAccountsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) WithAPIVersion(aPIVersion *string) *GetGcpCloudAccountsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get gcp cloud accounts params
func (o *GetGcpCloudAccountsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetGcpCloudAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int64

		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt64(qrDollarSkip)
		if qDollarSkip != "" {

			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}
	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int64

		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt64(qrDollarTop)
		if qDollarTop != "" {

			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}
	}

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
