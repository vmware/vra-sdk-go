// Code generated by go-swagger; DO NOT EDIT.

package fabric_network

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

// NewGetFabricNetworksParams creates a new GetFabricNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFabricNetworksParams() *GetFabricNetworksParams {
	return &GetFabricNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricNetworksParamsWithTimeout creates a new GetFabricNetworksParams object
// with the ability to set a timeout on a request.
func NewGetFabricNetworksParamsWithTimeout(timeout time.Duration) *GetFabricNetworksParams {
	return &GetFabricNetworksParams{
		timeout: timeout,
	}
}

// NewGetFabricNetworksParamsWithContext creates a new GetFabricNetworksParams object
// with the ability to set a context for a request.
func NewGetFabricNetworksParamsWithContext(ctx context.Context) *GetFabricNetworksParams {
	return &GetFabricNetworksParams{
		Context: ctx,
	}
}

// NewGetFabricNetworksParamsWithHTTPClient creates a new GetFabricNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFabricNetworksParamsWithHTTPClient(client *http.Client) *GetFabricNetworksParams {
	return &GetFabricNetworksParams{
		HTTPClient: client,
	}
}

/*
GetFabricNetworksParams contains all the parameters to send to the API endpoint

	for the get fabric networks operation.

	Typically these are written to a http.Request.
*/
type GetFabricNetworksParams struct {

	/* DollarCount.

	   Flag which when specified, regardless of the assigned value, shows the total number of records. If the collection has a filter it shows the number of records matching the filter.
	*/
	DollarCount *bool

	/* DollarFilter.

	   Filter the results by a specified predicate expression. Operators: eq, ne, and, or.
	*/
	DollarFilter *string

	/* DollarSelect.

	   Select a subset of properties to include in the response.
	*/
	DollarSelect *string

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

// WithDefaults hydrates default values in the get fabric networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFabricNetworksParams) WithDefaults() *GetFabricNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get fabric networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFabricNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get fabric networks params
func (o *GetFabricNetworksParams) WithTimeout(timeout time.Duration) *GetFabricNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric networks params
func (o *GetFabricNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric networks params
func (o *GetFabricNetworksParams) WithContext(ctx context.Context) *GetFabricNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric networks params
func (o *GetFabricNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric networks params
func (o *GetFabricNetworksParams) WithHTTPClient(client *http.Client) *GetFabricNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric networks params
func (o *GetFabricNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the get fabric networks params
func (o *GetFabricNetworksParams) WithDollarCount(dollarCount *bool) *GetFabricNetworksParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the get fabric networks params
func (o *GetFabricNetworksParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the get fabric networks params
func (o *GetFabricNetworksParams) WithDollarFilter(dollarFilter *string) *GetFabricNetworksParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get fabric networks params
func (o *GetFabricNetworksParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarSelect adds the dollarSelect to the get fabric networks params
func (o *GetFabricNetworksParams) WithDollarSelect(dollarSelect *string) *GetFabricNetworksParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get fabric networks params
func (o *GetFabricNetworksParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the get fabric networks params
func (o *GetFabricNetworksParams) WithDollarSkip(dollarSkip *int64) *GetFabricNetworksParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get fabric networks params
func (o *GetFabricNetworksParams) SetDollarSkip(dollarSkip *int64) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get fabric networks params
func (o *GetFabricNetworksParams) WithDollarTop(dollarTop *int64) *GetFabricNetworksParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get fabric networks params
func (o *GetFabricNetworksParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get fabric networks params
func (o *GetFabricNetworksParams) WithAPIVersion(aPIVersion *string) *GetFabricNetworksParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric networks params
func (o *GetFabricNetworksParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarCount != nil {

		// query param $count
		var qrDollarCount bool

		if o.DollarCount != nil {
			qrDollarCount = *o.DollarCount
		}
		qDollarCount := swag.FormatBool(qrDollarCount)
		if qDollarCount != "" {

			if err := r.SetQueryParam("$count", qDollarCount); err != nil {
				return err
			}
		}
	}

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string

		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {

			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}
	}

	if o.DollarSelect != nil {

		// query param $select
		var qrDollarSelect string

		if o.DollarSelect != nil {
			qrDollarSelect = *o.DollarSelect
		}
		qDollarSelect := qrDollarSelect
		if qDollarSelect != "" {

			if err := r.SetQueryParam("$select", qDollarSelect); err != nil {
				return err
			}
		}
	}

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
