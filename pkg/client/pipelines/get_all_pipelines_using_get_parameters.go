// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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
)

// NewGetAllPipelinesUsingGETParams creates a new GetAllPipelinesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAllPipelinesUsingGETParams() *GetAllPipelinesUsingGETParams {
	return &GetAllPipelinesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllPipelinesUsingGETParamsWithTimeout creates a new GetAllPipelinesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetAllPipelinesUsingGETParamsWithTimeout(timeout time.Duration) *GetAllPipelinesUsingGETParams {
	return &GetAllPipelinesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetAllPipelinesUsingGETParamsWithContext creates a new GetAllPipelinesUsingGETParams object
// with the ability to set a context for a request.
func NewGetAllPipelinesUsingGETParamsWithContext(ctx context.Context) *GetAllPipelinesUsingGETParams {
	return &GetAllPipelinesUsingGETParams{
		Context: ctx,
	}
}

// NewGetAllPipelinesUsingGETParamsWithHTTPClient creates a new GetAllPipelinesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAllPipelinesUsingGETParamsWithHTTPClient(client *http.Client) *GetAllPipelinesUsingGETParams {
	return &GetAllPipelinesUsingGETParams{
		HTTPClient: client,
	}
}

/* GetAllPipelinesUsingGETParams contains all the parameters to send to the API endpoint
   for the get all pipelines using g e t operation.

   Typically these are written to a http.Request.
*/
type GetAllPipelinesUsingGETParams struct {

	/* DollarFilter.

	   To list with OData like filter
	*/
	DollarFilter *string

	/* DollarOrderby.

	   Order by attribute

	   Default: "name desc"
	*/
	DollarOrderby *string

	/* DollarSkip.

	   To skip 'n' Pipelines for listing

	   Default: "0"
	*/
	DollarSkip *string

	/* DollarTop.

	   To list top 'n' Pipelines for listing

	   Default: "25"
	*/
	DollarTop *string

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Page.

	   To select 'n'th page for listing

	   Default: "0"
	*/
	Page *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get all pipelines using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllPipelinesUsingGETParams) WithDefaults() *GetAllPipelinesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get all pipelines using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAllPipelinesUsingGETParams) SetDefaults() {
	var (
		dollarOrderbyDefault = string("name desc")

		dollarSkipDefault = string("0")

		dollarTopDefault = string("25")

		pageDefault = string("0")
	)

	val := GetAllPipelinesUsingGETParams{
		DollarOrderby: &dollarOrderbyDefault,
		DollarSkip:    &dollarSkipDefault,
		DollarTop:     &dollarTopDefault,
		Page:          &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithTimeout(timeout time.Duration) *GetAllPipelinesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithContext(ctx context.Context) *GetAllPipelinesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithHTTPClient(client *http.Client) *GetAllPipelinesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithDollarFilter(dollarFilter *string) *GetAllPipelinesUsingGETParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarOrderby adds the dollarOrderby to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithDollarOrderby(dollarOrderby *string) *GetAllPipelinesUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetDollarOrderby(dollarOrderby *string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithDollarSkip(dollarSkip *string) *GetAllPipelinesUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetDollarSkip(dollarSkip *string) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithDollarTop(dollarTop *string) *GetAllPipelinesUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetDollarTop(dollarTop *string) {
	o.DollarTop = dollarTop
}

// WithAuthorization adds the authorization to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithAuthorization(authorization string) *GetAllPipelinesUsingGETParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithAPIVersion(aPIVersion string) *GetAllPipelinesUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithPage adds the page to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) WithPage(page *string) *GetAllPipelinesUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get all pipelines using get params
func (o *GetAllPipelinesUsingGETParams) SetPage(page *string) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllPipelinesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.DollarOrderby != nil {

		// query param $orderby
		var qrDollarOrderby string

		if o.DollarOrderby != nil {
			qrDollarOrderby = *o.DollarOrderby
		}
		qDollarOrderby := qrDollarOrderby
		if qDollarOrderby != "" {

			if err := r.SetQueryParam("$orderby", qDollarOrderby); err != nil {
				return err
			}
		}
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip string

		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := qrDollarSkip
		if qDollarSkip != "" {

			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}
	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop string

		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := qrDollarTop
		if qDollarTop != "" {

			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage string

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
