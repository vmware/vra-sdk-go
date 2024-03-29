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

// NewGetAwsCloudAccountsParams creates a new GetAwsCloudAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAwsCloudAccountsParams() *GetAwsCloudAccountsParams {
	return &GetAwsCloudAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAwsCloudAccountsParamsWithTimeout creates a new GetAwsCloudAccountsParams object
// with the ability to set a timeout on a request.
func NewGetAwsCloudAccountsParamsWithTimeout(timeout time.Duration) *GetAwsCloudAccountsParams {
	return &GetAwsCloudAccountsParams{
		timeout: timeout,
	}
}

// NewGetAwsCloudAccountsParamsWithContext creates a new GetAwsCloudAccountsParams object
// with the ability to set a context for a request.
func NewGetAwsCloudAccountsParamsWithContext(ctx context.Context) *GetAwsCloudAccountsParams {
	return &GetAwsCloudAccountsParams{
		Context: ctx,
	}
}

// NewGetAwsCloudAccountsParamsWithHTTPClient creates a new GetAwsCloudAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAwsCloudAccountsParamsWithHTTPClient(client *http.Client) *GetAwsCloudAccountsParams {
	return &GetAwsCloudAccountsParams{
		HTTPClient: client,
	}
}

/*
GetAwsCloudAccountsParams contains all the parameters to send to the API endpoint

	for the get aws cloud accounts operation.

	Typically these are written to a http.Request.
*/
type GetAwsCloudAccountsParams struct {

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

// WithDefaults hydrates default values in the get aws cloud accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsCloudAccountsParams) WithDefaults() *GetAwsCloudAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get aws cloud accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsCloudAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) WithTimeout(timeout time.Duration) *GetAwsCloudAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) WithContext(ctx context.Context) *GetAwsCloudAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) WithHTTPClient(client *http.Client) *GetAwsCloudAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSkip adds the dollarSkip to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) WithDollarSkip(dollarSkip *int64) *GetAwsCloudAccountsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) SetDollarSkip(dollarSkip *int64) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) WithDollarTop(dollarTop *int64) *GetAwsCloudAccountsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) WithAPIVersion(aPIVersion *string) *GetAwsCloudAccountsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get aws cloud accounts params
func (o *GetAwsCloudAccountsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetAwsCloudAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
